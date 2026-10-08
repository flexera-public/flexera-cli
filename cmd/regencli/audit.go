package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
)

func copyTree(source, destination string) error {
	return filepath.WalkDir(source, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}
		target := filepath.Join(destination, rel)
		if entry.IsDir() {
			return os.MkdirAll(target, 0o755)
		}
		if !entry.Type().IsRegular() {
			return fmt.Errorf("verification workspace: unsupported file %s", path)
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(target, data, 0o644)
	})
}

func auditStagedTree(stage string, catalogData []byte) error {
	workspace, err := os.MkdirTemp("", "flexera-cli-verify-*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(workspace)
	for _, name := range []string{"go.mod", "go.sum"} {
		data, err := os.ReadFile(name)
		if err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(workspace, name), data, 0o644); err != nil {
			return err
		}
	}
	if err := absolutizeLocalReplaces(workspace); err != nil {
		return err
	}
	entries, err := os.ReadDir("internal")
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if !entry.IsDir() || entry.Name() == "commands" {
			continue
		}
		if err := copyTree(filepath.Join("internal", entry.Name()), filepath.Join(workspace, "internal", entry.Name())); err != nil {
			return err
		}
	}
	if err := copyTree(stage, filepath.Join(workspace, "internal", "commands")); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(workspace, "internal", "catalog", "catalog_gen.json"), catalogData, 0o644); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(workspace, "audit.go"), []byte(stagedAuditSource), 0o644); err != nil {
		return err
	}
	c := exec.Command("go", "run", "./audit.go")
	c.Dir = workspace
	output, err := c.CombinedOutput()
	if err != nil {
		return fmt.Errorf("staged tree audit: %w\n%s", err, output)
	}
	return nil
}

// absolutizeLocalReplaces rewrites relative filesystem `replace` targets in the
// workspace go.mod (e.g. ../unified-go-client) so they still resolve from the
// temporary verification directory.
func absolutizeLocalReplaces(workspace string) error {
	out, err := exec.Command("go", "mod", "edit", "-json").Output()
	if err != nil {
		return fmt.Errorf("read go.mod replaces: %w", err)
	}
	var mod struct {
		Replace []struct {
			Old struct{ Path, Version string }
			New struct{ Path, Version string }
		}
	}
	if err := json.Unmarshal(out, &mod); err != nil {
		return fmt.Errorf("parse go.mod replaces: %w", err)
	}
	for _, r := range mod.Replace {
		if r.New.Version != "" || filepath.IsAbs(r.New.Path) {
			continue
		}
		abs, err := filepath.Abs(r.New.Path)
		if err != nil {
			return err
		}
		old := r.Old.Path
		if r.Old.Version != "" {
			old += "@" + r.Old.Version
		}
		c := exec.Command("go", "mod", "edit", "-replace", old+"="+abs)
		c.Dir = workspace
		if output, err := c.CombinedOutput(); err != nil {
			return fmt.Errorf("rewrite replace %s: %w\n%s", old, err, output)
		}
	}
	return nil
}

const stagedAuditSource = `package main
import (
 "fmt"
 "io"
 "os"
 "strings"
 "github.com/flexera-public/flexera-cli/internal/app"
 "github.com/flexera-public/flexera-cli/internal/catalog"
 "github.com/spf13/cobra"
 "github.com/spf13/pflag"
)
func main() {
 root,_:=app.NewRootCmd(io.Discard,io.Discard,func(string)string{return ""},nil,"audit")
 entries,err:=catalog.All();if err!=nil{panic(err)}
 expected:=map[string]string{};metadata:=map[string]catalog.Entry{};for _,e:=range entries {expected[e.OperationID]=strings.Join(e.Command," ");metadata[e.OperationID]=e}
 seen:=map[string]bool{};issues:=[]string{}
 var walk func(*cobra.Command)
 walk=func(c *cobra.Command) {
	if c.Parent()==root && (c.Name()=="cli" || contains(c.Aliases,"cli")) && c.Annotations["flexera.meta"]!="true" {issues=append(issues,"reserved top-level cli collision")}
	if id:=c.Annotations["flexera.operationId"];id!="" {path:=strings.TrimPrefix(c.CommandPath(),root.Name()+" ");if seen[id]||expected[id]!=path{issues=append(issues,fmt.Sprintf("catalog/tree mismatch %s: tree=%s catalog=%s",id,path,expected[id]))};seen[id]=true;for _,p:=range metadata[id].Params{if c.Flags().Lookup(p.Flag)==nil && c.InheritedFlags().Lookup(p.Flag)==nil{issues=append(issues,"catalog flag missing: "+path+" --"+p.Flag)}};for _,flag:=range metadata[id].BodyFlags{if c.Flags().Lookup(flag)==nil{issues=append(issues,"catalog body flag missing: "+path+" --"+flag)}}}
  if c!=root {c.LocalNonPersistentFlags().VisitAll(func(f *pflag.Flag){for p:=c.Parent();p!=nil;p=p.Parent(){p.PersistentFlags().VisitAll(func(inherited *pflag.Flag){if f.Name==inherited.Name || f.Shorthand!="" && f.Shorthand==inherited.Shorthand {issues=append(issues,fmt.Sprintf("flag collision: %s --%s with %s --%s",c.CommandPath(),f.Name,p.CommandPath(),inherited.Name))}})}})}
  names:=map[string]string{};for _,child:=range c.Commands(){for _,n:=range append([]string{child.Name()},child.Aliases...){if other,dup:=names[n];dup{issues=append(issues,fmt.Sprintf("sibling name collision under %s: %q used by %s and %s",c.CommandPath(),n,other,child.Name()))};names[n]=child.Name()}}
  for _,child:=range c.Commands(){walk(child)}
 }
 walk(root);for id,path:=range expected{if !seen[id]{issues=append(issues,"catalog entry without annotated command: "+id+" "+path)}}
 if len(issues)>0{for _,issue:=range issues{fmt.Fprintln(os.Stderr,issue)};os.Exit(1)}
}
func contains(values []string,want string)bool{for _,v:=range values{if v==want{return true}};return false}
`
