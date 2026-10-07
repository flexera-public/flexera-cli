package main

import (
	"bytes"
	"encoding/json"
	"go/format"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/flexera-public/flexera-cli/internal/catalog"
)

func TestFinalHelpPrefix(t *testing.T) {
	parent := genTag{Pkg: billConnectParent, Cmd: "billing"}
	for _, child := range billConnectChildren {
		t.Run(child.Pkg, func(t *testing.T) {
			tag := genTag{Pkg: child.Pkg, Cmd: "provisional-" + child.Pkg}
			if got := finalHelpPrefix(tag, []genTag{tag, parent}); got != "flexera-cli billing "+child.Use+" " {
				t.Fatalf("final prefix: %q", got)
			}
			if got := finalHelpPrefix(tag, []genTag{tag}); got != "flexera-cli "+tag.Cmd+" " {
				t.Fatalf("parent absent: %q", got)
			}
		})
	}
	if got := finalHelpPrefix(parent, []genTag{parent}); got != "flexera-cli billing " {
		t.Fatal(got)
	}
}

func TestRewriteCobraExamplesOnly(t *testing.T) {
	source := "package fixture\nimport (cb \"github.com/spf13/cobra\"; other \"example.org/other\")\n" +
		"// flexera-cli bill-connect-aws get\n" +
		"var unrelated = \"flexera-cli bill-connect-aws get\"\n" +
		"var _ = other.Command{Example: \"flexera-cli bill-connect-aws get\"}\n" +
		"var _ = &cb.Command{Use: \"bill-connect-aws\", Short: \"flexera-cli bill-connect-aws get\", Long: \"flexera-cli bill-connect-aws get\", Example: `flexera-cli bill-connect-aws get\nflexera-cli bill-connect-aws delete --yes\nflexera-cli bill-connect-aws-extra get\nflexera-cli budget get`}\n" +
		"var _ = cb.Command{Example: unrelated}\n" +
		"var _ = cb.Command{Example: \"flexera-cli bill-connect-aws get\\n  flexera-cli bill-connect-aws get --help\"}\n"
	before, err := format.Source([]byte(source))
	if err != nil {
		t.Fatal(err)
	}
	got, err := rewriteCobraExamples(before, "flexera-cli bill-connect-aws ", "flexera-cli bill-connect aws ")
	if err != nil {
		t.Fatal(err)
	}
	want := strings.Replace(string(before), "`flexera-cli bill-connect-aws get\nflexera-cli bill-connect-aws delete --yes\nflexera-cli bill-connect-aws-extra get\nflexera-cli budget get`", strconv.Quote("flexera-cli bill-connect aws get\nflexera-cli bill-connect aws delete --yes\nflexera-cli bill-connect-aws-extra get\nflexera-cli budget get"), 1)
	want = strings.Replace(want, `Example: "flexera-cli bill-connect-aws get\n  flexera-cli bill-connect-aws get --help"`, `Example: "flexera-cli bill-connect aws get\n  flexera-cli bill-connect aws get --help"`, 1)
	wantBytes, err := format.Source([]byte(want))
	if err != nil || !bytes.Equal(got, wantBytes) {
		t.Fatalf("changed unrelated source: %v\ngot:\n%s\nwant:\n%s", err, got, wantBytes)
	}
	second, err := rewriteCobraExamples(got, "flexera-cli bill-connect-aws ", "flexera-cli bill-connect aws ")
	if err != nil || !bytes.Equal(got, second) {
		t.Fatalf("not idempotent: %v", err)
	}
	if _, err := rewriteCobraExamples([]byte("not Go"), "a", "b"); err == nil {
		t.Fatal("accepted malformed Go")
	}
}

func TestRegenerateMapsHelpBeforePreparation(t *testing.T) {
	destination := filepath.Join(t.TempDir(), "commands")
	hooks := testHooks(t)
	hooks.generate = func(tag genTag, out string) error {
		writeTestFile(t, out, "package "+tag.Pkg+"\nimport \"github.com/spf13/cobra\"\nvar _ = cobra.Command{Example: "+strconv.Quote("flexera-cli "+tag.Cmd+" get")+"}\n")
		return nil
	}
	hooks.prepare = func(stage string, tags []genTag) ([]publicationArtifact, error) {
		data, err := os.ReadFile(filepath.Join(stage, "billconnectaws", "cmd_gen.go"))
		if err != nil || !strings.Contains(string(data), "flexera-cli bill-connect aws get") {
			t.Fatalf("preparation saw provisional help: %s, %v", data, err)
		}
		if _, err := os.Stat(filepath.Join(stage, "register_gen.go")); err != nil {
			t.Fatal(err)
		}
		return nil, nil
	}
	if err := regenerate(destination, []string{"Bill Connect", "Bill Connect AWS"}, hooks); err != nil {
		t.Fatal(err)
	}
}

func TestStagedHelpWithoutParentLeavesBytesUntouched(t *testing.T) {
	stage := t.TempDir()
	tag := genTag{Pkg: "billconnectaws", Cmd: "bill-connect-aws"}
	// Deliberately not gofmt-formatted: a no-op must not rewrite the file.
	source := "package fixture\nimport \"github.com/spf13/cobra\"\nvar _=cobra.Command{Example:\"flexera-cli bill-connect-aws get\"}\n"
	path := filepath.Join(stage, tag.Pkg, "cmd_gen.go")
	writeTestFile(t, path, source)
	if err := rewriteStagedHelp(stage, []genTag{tag}); err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(path)
	if err != nil || string(got) != source {
		t.Fatalf("ungrouped file changed: %q, %v", got, err)
	}
}

func TestStagedHelpFailureDoesNotPublish(t *testing.T) {
	parent := t.TempDir()
	destination := filepath.Join(parent, "commands")
	writeTestFile(t, filepath.Join(destination, "register_gen.go"), "old registry")
	hooks := testHooks(t)
	hooks.verify = func(stage string, tags []genTag) error {
		// Inject invalid staged source after the separate symbol gate.
		writeTestFile(t, filepath.Join(stage, "billconnectaws", "cmd_gen.go"), "not Go")
		return nil
	}
	hooks.prepare = func(string, []genTag) ([]publicationArtifact, error) {
		t.Fatal("preparation ran after failed help transform")
		return nil, nil
	}
	if err := regenerate(destination, []string{"Bill Connect", "Bill Connect AWS"}, hooks); err == nil || !strings.Contains(err.Error(), "staged help") {
		t.Fatalf("expected contextual transform error: %v", err)
	}
	files := treeContents(t, destination)
	if len(files) != 1 || files["register_gen.go"] != "old registry" {
		t.Fatalf("published bytes changed: %v", files)
	}
	assertNoTemporaryTrees(t, parent)
}

// Audit every invocation against an actual tree compiled from writeRegister,
// including operation aliases. Catalog paths must match the resolved canonical
// leaf, not just share a plausible textual prefix.
func TestCatalogHelpInvocationsMatchRegisteredTree(t *testing.T) {
	workspace := t.TempDir()
	stage := filepath.Join(workspace, "internal", "commands")
	tags := []genTag{{Pkg: billConnectParent, Cmd: "bill-connect"}, {Pkg: "budget", Cmd: "budget"}}
	for _, child := range billConnectChildren {
		tags = append(tags, genTag{Pkg: child.Pkg, Cmd: "bill-connect-" + child.Use})
	}
	paths := map[string]any{}
	for _, tag := range tags {
		entry := catalog.Entry{OperationID: tag.Pkg, Command: []string{tag.Cmd, "get"}, Method: "GET", Path: "/" + tag.Pkg, ResponseEnvelope: "none"}
		metadata, err := json.Marshal(map[string]catalog.Entry{tag.Pkg: entry})
		if err != nil {
			t.Fatal(err)
		}
		writeTestFile(t, filepath.Join(stage, tag.Pkg, "metadata.json"), string(metadata))
		paths[entry.Path] = map[string]any{"get": map[string]string{"operationId": tag.Pkg}}
		example := "  flexera-cli " + tag.Cmd + " get --help\n  flexera-cli " + tag.Cmd + " show --help"
		writeTestFile(t, filepath.Join(stage, tag.Pkg, "cmd_gen.go"), "package "+tag.Pkg+"\nimport \"github.com/spf13/cobra\"\nfunc NewCmd() *cobra.Command { c := &cobra.Command{Use: "+strconv.Quote(tag.Cmd)+"}; c.AddCommand(&cobra.Command{Use: \"get\", Aliases: []string{\"show\"}, Annotations: map[string]string{\"flexera.operationId\": "+strconv.Quote(tag.Pkg)+"}, Example: "+strconv.Quote(example)+"}); return c }\n")
	}
	if err := writeRegister(stage, tags); err != nil {
		t.Fatal(err)
	}
	if err := rewriteStagedHelp(stage, tags); err != nil {
		t.Fatal(err)
	}
	specData, err := json.Marshal(map[string]any{"paths": paths})
	if err != nil {
		t.Fatal(err)
	}
	data, err := buildCatalog(stage, tags, specData)
	if err != nil {
		t.Fatal(err)
	}
	writeTestFile(t, filepath.Join(workspace, "catalog.json"), string(data))
	for _, name := range []string{"go.mod", "go.sum"} {
		data, err := os.ReadFile(filepath.Join("..", "..", name))
		if err != nil {
			t.Fatal(err)
		}
		writeTestFile(t, filepath.Join(workspace, name), string(data))
	}
	writeTestFile(t, filepath.Join(workspace, "audit.go"), helpAuditTestSource)
	command := exec.Command("go", "run", "-mod=readonly", "./audit.go")
	command.Dir = workspace
	command.Env = append(os.Environ(), "GOWORK=off")
	if output, err := command.CombinedOutput(); err != nil {
		t.Fatalf("catalog/help audit: %v\n%s", err, output)
	}
}

const helpAuditTestSource = `package main
import (
 "encoding/json"
 "fmt"
 "os"
 "strings"
 "github.com/flexera-public/flexera-cli/internal/commands"
 "github.com/spf13/cobra"
)
func main() {
 data,err:=os.ReadFile("catalog.json");if err!=nil{panic(err)}
 var doc struct{Entries []struct{OperationID string;Command []string}}
 if err:=json.Unmarshal(data,&doc);err!=nil{panic(err)}
 expected:=map[string]string{};for _,e:=range doc.Entries{expected[e.OperationID]="flexera-cli "+strings.Join(e.Command," ")}
 root:=&cobra.Command{Use:"flexera-cli"};commands.RegisterAll(root)
 seen:=map[string]bool{}
 var walk func(*cobra.Command)
 walk=func(c *cobra.Command){
  id:=c.Annotations["flexera.operationId"]
  if id!=""{
   seen[id]=true;want:=expected[id];if c.CommandPath()!=want{panic(fmt.Sprintf("catalog path %s != %s",want,c.CommandPath()))}
   count:=0
   for _,line:=range strings.Split(c.Example,"\n"){
    words:=strings.Fields(line);if len(words)==0||words[0]!="flexera-cli"{continue};count++
    args:=[]string{};for _,word:=range words[1:]{if strings.HasPrefix(word,"-"){break};args=append(args,word)}
    leaf,rest,err:=root.Find(args);if err!=nil||len(rest)!=0||leaf!=c{panic(fmt.Sprintf("invalid invocation %q: %v %v",line,rest,err))}
    if !strings.HasPrefix(line,"  "+strings.TrimSuffix(want," get")+" "){panic("wrong final prefix: "+line)}
   }
   if count!=2{panic("missing help invocations: "+id)}
  }
  for _,child:=range c.Commands(){walk(child)}
 };walk(root)
 for id:=range expected{if !seen[id]{panic("unregistered catalog operation: "+id)}}
}
`
