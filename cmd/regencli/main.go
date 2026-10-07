// Command regencli regenerates every per-tag cobra command package from the
// annotated unified OpenAPI spec, plus the RegisterAll wiring.
//
// It enumerates tags in the spec, invokes ./cmd/gencli for each into
// internal/commands/<pkg>/cmd_gen.go, prunes stale packages, and writes
// internal/commands/register_gen.go containing:
//
//	func RegisterAll(root *cobra.Command) { root.AddCommand(<tag>.NewCmd(), ...) }
//
// Usage:
//
//	go run ./cmd/regencli
//
// Invoked transitively by `go generate ./...` via the //go:generate
// directive in cli/flexera-cli/generate.go.
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

const (
	specPath    = "unified-openapi/openapi3.json"
	commandsDir = "internal/commands"
	modulePath  = "github.com/flexera-public/flexera-cli"
	// cmdSuffix is appended to every generated root command name. Empty now
	// that the generated commands are canonical (the hand-written CRUD they
	// replaced has been deleted).
	cmdSuffix = ""
)

// curatedCanonical lists canonical command names owned by a hand-written
// curated command. The matching generated tag is skipped so it does not
// collide with (or duplicate) the curated command of the same name.
var curatedCanonical = map[string]bool{
	"user-orgs": true, // IAM user-orgs is exposed via the curated user-orgs command
	// Project-scoped policy tags (paths include project_id): owned by the
	// curated `policy` super-command, which adds GRS project auto-resolution.
	// The org-scoped mechanical policy tags are NOT excluded — they generate
	// as top-level canonical commands.
	"action-status":     true,
	"applied-policy":    true,
	"archived-incident": true,
	"policy-template":   true,
}

var supportedActions = map[string]bool{
	"list": true, "get": true, "create": true,
	"replace": true, "update": true, "delete": true,
	"action": true,
}

type genTag struct {
	Tag string // OpenAPI tag
	Pkg string // Go package / dir name
	Cmd string // CLI command name (with suffix)
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "regencli:", err)
		os.Exit(1)
	}
}

func run() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	if filepath.Base(wd) != "flexera-cli" {
		return fmt.Errorf("regencli must run from cli/flexera-cli (cwd: %s)", wd)
	}

	tags, err := collectTags()
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "regencli: %d tags discovered\n", len(tags))
	return regenerate(commandsDir, tags, generationHooks{
		generate: generateTag,
		verify:   verifyGenerated,
		rename:   os.Rename,
		prepare:  preparePublication,
	})
}

// Hooks keep failure tests independent of subprocesses and the real SDK.
type generationHooks struct {
	generate func(genTag, string) error
	verify   func(string, []genTag) error
	rename   func(string, string) error
	prepare  func(string, []genTag) ([]publicationArtifact, error)
}

func generateTag(tag genTag, out string) error {
	c := exec.Command("go", "run", "./cmd/gencli",
		"-spec", specPath, "-tag", tag.Tag, "-pkg", tag.Pkg, "-cmd", tag.Cmd, "-out", out, "-metadata", filepath.Join(filepath.Dir(out), "metadata.json"))
	c.Stderr = os.Stderr
	return c.Run()
}

func regenerate(destination string, tags []string, hooks generationHooks) (err error) {
	parent := filepath.Dir(destination)
	if err := os.MkdirAll(parent, 0o755); err != nil {
		return err
	}
	// A sibling staging directory ensures publication stays on one filesystem.
	stage, err := os.MkdirTemp(parent, ".commands-stage-")
	if err != nil {
		return err
	}
	defer func() { err = errors.Join(err, os.RemoveAll(stage)) }()
	// MkdirTemp defaults to 0700; published commands retain the usual mode.
	if err := os.Chmod(stage, 0o755); err != nil {
		return err
	}

	used := map[string]bool{}
	var generated []genTag
	for _, tag := range tags {
		cmd := kebab(cleanIdent(tag)) + cmdSuffix
		if curatedCanonical[cmd] || tag == "Authentication" {
			fmt.Fprintf(os.Stderr, "regencli: skip tag %q (curated command owns %q)\n", tag, cmd)
			continue
		}
		pkg := uniquePkg(cleanIdent(tag), used)
		t := genTag{Tag: tag, Pkg: pkg, Cmd: cmd}
		out := filepath.Join(stage, pkg, "cmd_gen.go")
		if err := hooks.generate(t, out); err != nil {
			return fmt.Errorf("generate tag %q: %w", tag, err)
		}
		generated = append(generated, t)
	}

	if err := hooks.verify(stage, generated); err != nil {
		return err
	}
	if err := writeRegister(stage, generated); err != nil {
		return fmt.Errorf("write staged registration: %w", err)
	}
	if err := rewriteStagedHelp(stage, generated); err != nil {
		return err
	}
	if hooks.prepare != nil {
		artifacts, err := hooks.prepare(stage, generated)
		if err != nil {
			return err
		}
		defer func() {
			for _, artifact := range artifacts {
				err = errors.Join(err, os.RemoveAll(artifact.Staged))
			}
		}()
		artifacts = append([]publicationArtifact{{Staged: stage, Destination: destination}}, artifacts...)
		return publishArtifacts(artifacts, hooks.rename)
	}
	if err := publish(stage, destination, hooks.rename); err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "regencli: wrote %d generated tag packages\n", len(generated))
	return nil
}

func preparePublication(stage string, tags []genTag) (artifacts []publicationArtifact, err error) {
	defer func() {
		if err != nil {
			for _, artifact := range artifacts {
				_ = os.RemoveAll(artifact.Staged)
			}
		}
	}()
	specData, err := os.ReadFile(specPath)
	if err != nil {
		return nil, err
	}
	curated, err := curatedEntries(specData)
	if err != nil {
		return nil, err
	}
	data, err := buildCatalog(stage, tags, specData, curated...)
	if err != nil {
		return nil, err
	}
	if len(data) > 5*1024*1024 {
		return nil, fmt.Errorf("catalog size %d exceeds the approved 5 MiB budget", len(data))
	}
	second, err := buildCatalog(stage, tags, specData, curated...)
	if err != nil {
		return nil, err
	}
	if string(data) != string(second) {
		return nil, fmt.Errorf("catalog generation is nondeterministic")
	}
	report, err := verifyCoverage(specData, data)
	if err != nil {
		return nil, err
	}
	fmt.Fprintf(os.Stderr, "regencli: catalog %d bytes, %d exact curated operations\n", len(data), len(curated))
	if err := auditStagedTree(stage, data); err != nil {
		return nil, err
	}
	for _, tag := range tags {
		if err := os.Remove(filepath.Join(stage, tag.Pkg, "metadata.json")); err != nil {
			return nil, err
		}
	}
	for _, item := range []struct {
		name, destination string
		data              []byte
	}{{"catalog_gen.json", "internal/catalog/catalog_gen.json", data}, {"coverage_gen.json", "internal/catalog/coverage_gen.json", report}} {
		// Files are staged beside the commands tree; publication uses sibling
		// backups and rollback for all destinations as one transaction.
		path := filepath.Join(stage, item.name)
		if err := os.WriteFile(path, item.data, 0o644); err != nil {
			return nil, err
		}
		// Keep publication sources outside the directory renamed first.
		file, err := os.CreateTemp(filepath.Dir(stage), ".catalog-stage-*")
		if err != nil {
			return nil, err
		}
		name := file.Name()
		artifacts = append(artifacts, publicationArtifact{Staged: name, Destination: item.destination})
		if err := file.Close(); err != nil {
			return nil, err
		}
		if err := os.Rename(path, name); err != nil {
			return nil, err
		}
	}
	return artifacts, nil
}

// publish replaces the entire tree (including registration), pruning stale
// packages only after generation and verification succeed. If restoration also
// fails, preserve the backup and report its location for manual recovery.
func publish(stage, destination string, rename func(string, string) error) (err error) {
	if _, err := os.Lstat(destination); errors.Is(err, os.ErrNotExist) {
		return rename(stage, destination)
	} else if err != nil {
		return err
	}
	backupDir, err := os.MkdirTemp(filepath.Dir(destination), ".commands-backup-")
	if err != nil {
		return err
	}
	backup := filepath.Join(backupDir, "commands")
	keepBackup := false
	defer func() {
		if !keepBackup {
			err = errors.Join(err, os.RemoveAll(backupDir))
		}
	}()
	if err := rename(destination, backup); err != nil {
		return fmt.Errorf("backup generated commands: %w", err)
	}
	if err := rename(stage, destination); err != nil {
		if rollbackErr := rename(backup, destination); rollbackErr != nil {
			keepBackup = true
			return errors.Join(fmt.Errorf("publish generated commands: %w", err),
				fmt.Errorf("rollback failed; old commands preserved at %s: %w", backup, rollbackErr))
		}
		return fmt.Errorf("publish generated commands (rolled back): %w", err)
	}
	return nil
}

// verifyGenerated type-checks every flexera.* / client.* symbol referenced
// by the freshly generated command files against the real unified-go-client
// package (loaded once via go/packages + go/types). This turns a spec/client
// naming drift (e.g. an operationId or generated type renamed upstream)
// into a hard regen-time failure instead of a downstream `go build` error —
// or worse, a silent mismatch that only trips at runtime.
func verifyGenerated(directory string, tags []genTag) error {
	sv, err := newSymbolVerifier()
	if err != nil {
		return err
	}
	return verifyFiles(sv, directory, tags)
}

func verifyFiles(sv *symbolVerifier, directory string, tags []genTag) error {
	var issues []string
	for _, t := range tags {
		out := filepath.Join(directory, t.Pkg, "cmd_gen.go")
		issues = append(issues, sv.verifyFile(out)...)
	}
	if len(issues) > 0 {
		return fmt.Errorf("symbol verification against unified-go-client failed (%d mismatch(es)):\n%s", len(issues), strings.Join(issues, "\n"))
	}
	fmt.Fprintf(os.Stderr, "regencli: verified %d generated file(s) against unified-go-client symbols\n", len(tags))
	return nil
}

type spec struct {
	Paths      map[string]map[string]json.RawMessage `json:"paths"`
	Components struct {
		Schemas map[string]json.RawMessage `json:"schemas"`
	} `json:"components"`
}

func collectTags() ([]string, error) {
	data, err := os.ReadFile(specPath)
	if err != nil {
		return nil, err
	}
	var s spec
	if err := json.Unmarshal(data, &s); err != nil {
		return nil, err
	}
	seen := map[string]bool{}
	for _, pi := range s.Paths {
		for m, raw := range pi {
			if m == "parameters" || m == "summary" || m == "description" {
				continue
			}
			var op map[string]interface{}
			if err := json.Unmarshal(raw, &op); err != nil {
				continue
			}
			if _, annotated := op["x-flexera-action"]; !annotated {
				continue
			}
			tags, _ := op["tags"].([]interface{})
			for _, t := range tags {
				if ts, ok := t.(string); ok {
					seen[ts] = true
				}
			}
		}
	}
	tags := make([]string, 0, len(seen))
	for t := range seen {
		tags = append(tags, t)
	}
	sort.Strings(tags)
	return tags, nil
}

// billConnectParent / billConnectChildren describe how the generated
// bill-connect tags are nested under a single top-level `bill-connect` command
// (instead of 8 separate top-level commands). The vendor packages are still
// generated and imported — they are just nested, not registered top-level.
var billConnectParent = "billconnect"

var billConnectChildren = []struct{ Pkg, Use string }{
	{"billconnectaws", "aws"},
	{"billconnectazurecsp", "azure-csp"},
	{"billconnectazureeamanagement", "azure-ea-management"},
	{"billconnectazuremca", "azure-mca"},
	{"billconnectcommonbillingestion", "common-bill-ingestion"},
	{"billconnectdatabricks", "databricks"},
	{"billconnectgcp", "gcp"},
}

func writeRegister(directory string, tags []genTag) error {
	sort.Slice(tags, func(i, j int) bool { return tags[i].Pkg < tags[j].Pkg })

	present := map[string]bool{}
	for _, t := range tags {
		present[t.Pkg] = true
	}

	// Bill-connect grouping: only when the base package is present.
	grouped := map[string]bool{}
	var bcChildren []struct{ Pkg, Use string }
	if present[billConnectParent] {
		grouped[billConnectParent] = true
		for _, c := range billConnectChildren {
			if present[c.Pkg] {
				grouped[c.Pkg] = true
				bcChildren = append(bcChildren, c)
			}
		}
	}

	var b strings.Builder
	b.WriteString("// Code generated by cmd/regencli. DO NOT EDIT.\n\n")
	b.WriteString("package commands\n\n")
	b.WriteString("import (\n")
	b.WriteString("\t\"github.com/spf13/cobra\"\n\n")
	for _, t := range tags {
		fmt.Fprintf(&b, "\t%s %q\n", t.Pkg, modulePath+"/"+commandsDir+"/"+t.Pkg)
	}
	b.WriteString(")\n\n")
	b.WriteString("// RegisterAll adds every generated tag command to root.\n")
	b.WriteString("func RegisterAll(root *cobra.Command) {\n")
	b.WriteString("\troot.AddCommand(\n")
	for _, t := range tags {
		if grouped[t.Pkg] {
			continue
		}
		fmt.Fprintf(&b, "\t\t%s.NewCmd(),\n", t.Pkg)
	}
	b.WriteString("\t)\n")
	if len(grouped) > 0 {
		b.WriteString("\troot.AddCommand(billConnectGroup())\n")
	}
	b.WriteString("}\n")

	if len(grouped) > 0 {
		b.WriteString("\n// billConnectGroup nests the per-vendor bill-connect commands under a\n")
		b.WriteString("// single `bill-connect` parent (vendors are not top-level commands).\n")
		b.WriteString("func billConnectGroup() *cobra.Command {\n")
		b.WriteString("\tnest := func(c *cobra.Command, use string) *cobra.Command { c.Use = use; return c }\n")
		fmt.Fprintf(&b, "\tparent := %s.NewCmd()\n", billConnectParent)
		b.WriteString("\tparent.AddCommand(\n")
		for _, c := range bcChildren {
			fmt.Fprintf(&b, "\t\tnest(%s.NewCmd(), %q),\n", c.Pkg, c.Use)
		}
		b.WriteString("\t)\n")
		b.WriteString("\treturn parent\n")
		b.WriteString("}\n")
	}

	return os.WriteFile(filepath.Join(directory, "register_gen.go"), []byte(b.String()), 0o644)
}

// cleanIdent turns a tag into a PascalCase alphanumeric identifier
// ("API Event" -> "APIEvent", "Bill Connect AWS" -> "BillConnectAWS").
func cleanIdent(t string) string {
	tmp := strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			return r
		}
		return ' '
	}, t)
	parts := strings.Fields(tmp)
	for i, p := range parts {
		parts[i] = strings.ToUpper(p[:1]) + p[1:]
	}
	return strings.Join(parts, "")
}

// uniquePkg lowercases ident into a Go package name, ensuring uniqueness and
// a non-digit first rune.
func uniquePkg(ident string, used map[string]bool) string {
	base := strings.ToLower(ident)
	if base == "" {
		base = "tag"
	}
	if base[0] >= '0' && base[0] <= '9' {
		base = "t" + base
	}
	pkg := base
	for n := 2; used[pkg]; n++ {
		pkg = fmt.Sprintf("%s%d", base, n)
	}
	used[pkg] = true
	return pkg
}

func kebab(s string) string {
	runes := []rune(s)
	var b strings.Builder
	for i, r := range runes {
		isUpper := r >= 'A' && r <= 'Z'
		if i > 0 && isUpper {
			prev := runes[i-1]
			prevUpper := prev >= 'A' && prev <= 'Z'
			nextLower := i+1 < len(runes) && runes[i+1] >= 'a' && runes[i+1] <= 'z'
			if !prevUpper || (prevUpper && nextLower) {
				b.WriteByte('-')
			}
		}
		if isUpper {
			b.WriteRune(r + 32)
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}
