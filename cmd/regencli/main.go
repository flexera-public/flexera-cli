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
	wd, err := os.Getwd()
	check(err)
	if filepath.Base(wd) != "flexera-cli" {
		fmt.Fprintln(os.Stderr, "regencli must run from cli/flexera-cli (cwd:", wd, ")")
		os.Exit(2)
	}

	tags := collectTags()
	fmt.Fprintf(os.Stderr, "regencli: %d tags discovered\n", len(tags))

	// Remove the entire generated tree so dropped tags get pruned.
	check(os.RemoveAll(commandsDir))
	check(os.MkdirAll(commandsDir, 0o755))

	// Remove any legacy package-main generated files from the prior
	// switch-tree generator.
	for _, glob := range []string{"cmd_*_gen.go", "cmd_*_gen.go.raw", "commands_gen.go"} {
		matches, _ := filepath.Glob(glob)
		for _, m := range matches {
			_ = os.Remove(m)
		}
	}

	used := map[string]bool{}
	var generated []genTag
	for _, tag := range tags {
		cmd := kebab(cleanIdent(tag)) + cmdSuffix
		if curatedCanonical[cmd] {
			fmt.Fprintf(os.Stderr, "regencli: skip tag %q (curated command owns %q)\n", tag, cmd)
			continue
		}
		pkg := uniquePkg(cleanIdent(tag), used)
		out := filepath.Join(commandsDir, pkg, "cmd_gen.go")
		c := exec.Command("go", "run", "./cmd/gencli",
			"-spec", specPath, "-tag", tag, "-pkg", pkg, "-cmd", cmd, "-out", out)
		c.Stderr = os.Stderr
		if err := c.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "regencli: skip tag %q (%v)\n", tag, err)
			// Drop the empty dir if gencli created none.
			_ = os.Remove(filepath.Join(commandsDir, pkg))
			delete(used, pkg)
			continue
		}
		generated = append(generated, genTag{Tag: tag, Pkg: pkg, Cmd: cmd})
	}

	verifyGenerated(generated)

	writeRegister(generated)
	fmt.Fprintf(os.Stderr, "regencli: wrote %d generated tag packages\n", len(generated))
}

// verifyGenerated type-checks every flexera.* / client.* symbol referenced
// by the freshly generated command files against the real unified-go-client
// package (loaded once via go/packages + go/types). This turns a spec/client
// naming drift (e.g. an operationId or generated type renamed upstream)
// into a hard regen-time failure instead of a downstream `go build` error —
// or worse, a silent mismatch that only trips at runtime.
func verifyGenerated(tags []genTag) {
	sv, err := newSymbolVerifier()
	check(err)

	var issues []string
	for _, t := range tags {
		out := filepath.Join(commandsDir, t.Pkg, "cmd_gen.go")
		issues = append(issues, sv.verifyFile(out)...)
	}
	if len(issues) > 0 {
		fmt.Fprintln(os.Stderr, "regencli: symbol verification against unified-go-client failed:")
		for _, msg := range issues {
			fmt.Fprintln(os.Stderr, "  "+msg)
		}
		fmt.Fprintf(os.Stderr, "regencli: %d symbol mismatch(es); aborting before register_gen.go is written\n", len(issues))
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "regencli: verified %d generated file(s) against unified-go-client symbols\n", len(tags))
}

type spec struct {
	Paths map[string]map[string]json.RawMessage `json:"paths"`
}

func collectTags() []string {
	data, err := os.ReadFile(specPath)
	check(err)
	var s spec
	check(json.Unmarshal(data, &s))
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
			action, _ := op["x-flexera-action"].(string)
			if !supportedActions[action] {
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
	return tags
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

func writeRegister(tags []genTag) {
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

	check(os.WriteFile(filepath.Join(commandsDir, "register_gen.go"), []byte(b.String()), 0o644))
}

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "regencli:", err)
		os.Exit(1)
	}
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
