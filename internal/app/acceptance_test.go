package app_test

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/flexera-public/flexera-cli/internal/app"
	"github.com/flexera-public/flexera-cli/internal/catalog"
	"github.com/flexera-public/flexera-cli/internal/commands"
	"github.com/spf13/cobra"
)

// Identify generated operations from the registry, rather than treating every
// annotated command as generated (the live tree also contains curated owners).
func acceptanceGeneratedCommands(t *testing.T) (*cobra.Command, map[string]*cobra.Command) {
	t.Helper()
	collect := func(root *cobra.Command) map[string]*cobra.Command {
		result := map[string]*cobra.Command{}
		var walk func(*cobra.Command)
		walk = func(cmd *cobra.Command) {
			if id := cmd.Annotations["flexera.operationId"]; id != "" {
				if previous := result[id]; previous != nil {
					t.Fatalf("duplicate operation %s: %s and %s", id, previous.CommandPath(), cmd.CommandPath())
				}
				result[id] = cmd
			}
			for _, child := range cmd.Commands() {
				walk(child)
			}
		}
		walk(root)
		return result
	}
	registry := &cobra.Command{Use: "flexera-cli"}
	commands.RegisterAll(registry)
	generated := collect(registry)
	if len(generated) == 0 {
		t.Fatal("generated registry contains no operations")
	}
	root, _ := app.NewRootCmd(io.Discard, io.Discard, func(string) string { return "" }, nil, "test")
	live := collect(root)
	for id := range generated {
		cmd := live[id]
		if cmd == nil {
			t.Fatalf("generated operation %s is missing from the live tree", id)
		}
		generated[id] = cmd
	}
	return root, generated
}

func TestAcceptanceGeneratedExamplesResolveLiveCommands(t *testing.T) {
	root, generated := acceptanceGeneratedCommands(t)
	nestedVendors := map[string]bool{}
	for id, cmd := range generated {
		t.Run(id, func(t *testing.T) {
			invocations := 0
			for _, line := range strings.Split(cmd.Example, "\n") {
				line = strings.TrimSpace(line)
				if !strings.HasPrefix(line, "flexera-cli ") {
					continue
				}
				invocations++
				if strings.HasPrefix(line, "flexera-cli cli schema ") {
					parts := strings.Fields(line)
					end := 3
					for end < len(parts) && !strings.HasPrefix(parts[end], "--") {
						end++
					}
					found, remaining, err := root.Find(parts[3:end])
					if err != nil || len(remaining) != 0 || found != cmd {
						t.Errorf("schema export points at wrong command: %s", line)
					}
					continue
				}
				if acceptanceUnquotedLessThan(line) {
					t.Errorf("example contains an unquoted shell placeholder: %s", line)
				}
				tokens := strings.Fields(line)
				for _, token := range tokens[1:] {
					if token == "--yes" || strings.HasPrefix(token, "--yes=") {
						t.Errorf("example pre-approves a write: %s", line)
					}
				}
				// Do not parse or execute placeholder flag values: UUID, number,
				// and required-parameter validation are not relevant to routing.
				end := 1
				for end < len(tokens) && !strings.HasPrefix(tokens[end], "-") {
					end++
				}
				found, remaining, err := root.Find(tokens[1:end])
				if err != nil {
					t.Fatalf("resolving example %q: %v", line, err)
				}
				if found != cmd || found.Annotations["flexera.operationId"] != id || len(remaining) != 0 {
					t.Errorf("example %q resolves to %s (%s), remaining %v; want %s (%s)", line, found.CommandPath(), found.Annotations["flexera.operationId"], remaining, cmd.CommandPath(), id)
				}
				for _, token := range tokens[end:] {
					if !strings.HasPrefix(token, "--") {
						continue
					}
					name := strings.SplitN(strings.TrimPrefix(token, "--"), "=", 2)[0]
					if cmd.Flags().Lookup(name) == nil && cmd.InheritedFlags().Lookup(name) == nil {
						t.Errorf("example documents nonexistent flag --%s: %s", name, line)
					}
				}
			}
			if invocations == 0 {
				t.Error("generated operation has no flexera-cli invocation example")
			}
		})
		path := strings.Fields(cmd.CommandPath())
		if len(path) == 4 && path[1] == "bill-connect" {
			nestedVendors[path[2]] = true
		}
	}
	// These vendors are deliberately nested by RegisterAll, not published as
	// top-level bill-connect-<vendor> command paths.
	for _, vendor := range []string{"aws", "azure-csp", "azure-ea-management", "azure-mca", "common-bill-ingestion", "databricks", "gcp"} {
		if !nestedVendors[vendor] {
			t.Errorf("no generated examples exercised bill-connect %s", vendor)
		}
	}
}

func acceptanceUnquotedLessThan(line string) bool {
	var quote rune
	escaped := false
	for _, ch := range line {
		if escaped {
			escaped = false
			continue
		}
		if ch == '\\' && quote != '\'' {
			escaped = true
			continue
		}
		if quote != 0 {
			if ch == quote {
				quote = 0
			}
			continue
		}
		if ch == '\'' || ch == '"' {
			quote = ch
		} else if ch == '<' {
			return true
		}
	}
	return false
}

func TestAcceptanceGeneratedInteractiveFlags(t *testing.T) {
	_, generated := acceptanceGeneratedCommands(t)
	data, err := os.ReadFile("../../unified-openapi/openapi3.json")
	if err != nil {
		t.Fatal(err)
	}
	type operation struct {
		ID          string `json:"operationId"`
		RequestBody *struct {
			Ref     string                     `json:"$ref"`
			Content map[string]json.RawMessage `json:"content"`
		} `json:"requestBody"`
	}
	var spec struct {
		Paths map[string]map[string]json.RawMessage `json:"paths"`
	}
	if err := json.Unmarshal(data, &spec); err != nil {
		t.Fatal(err)
	}
	eligible := map[string]bool{}
	for _, path := range spec.Paths {
		for method, raw := range path {
			var op operation
			if json.Unmarshal(raw, &op) != nil || op.ID == "" {
				continue
			}
			if op.RequestBody != nil && op.RequestBody.Ref != "" {
				t.Fatalf("resolve requestBody reference before classifying %s", op.ID)
			}
			write := method == "post" || method == "put" || method == "patch" || method == "delete"
			eligible[op.ID] = write && op.RequestBody != nil && op.RequestBody.Content["application/json"] != nil
		}
	}
	writes, others := 0, 0
	for id, cmd := range generated {
		t.Run(id, func(t *testing.T) {
			want, exists := eligible[id]
			if !exists {
				t.Fatal("generated operation is missing from the spec")
			}
			local := cmd.LocalNonPersistentFlags()
			interactive := local.Lookup("interactive")
			shorthand := local.ShorthandLookup("i")
			if want {
				writes++
				if interactive == nil {
					t.Fatal("JSON write is missing local --interactive")
				}
				if interactive.Shorthand != "i" || shorthand != interactive || interactive.Value.Type() != "bool" || interactive.DefValue != "false" {
					t.Error("--interactive must be a local boolean, default false, with shorthand -i")
				}
			} else {
				others++
				if interactive != nil || shorthand != nil || cmd.InheritedFlags().Lookup("interactive") != nil || cmd.InheritedFlags().ShorthandLookup("i") != nil {
					t.Error("non-JSON-write command exposes --interactive or -i")
				}
			}
		})
	}
	if writes == 0 || others == 0 {
		t.Fatalf("interactive coverage is empty: %d JSON writes, %d other operations", writes, others)
	}
	// Ancestor long/shorthand collisions are covered separately by
	// TestNoLocalFlagsShadowInheritedFlags in catalog_test.go.
}

func TestAcceptancePublishedRequestExamplesAreValid(t *testing.T) {
	index, err := catalog.Load()
	if err != nil {
		t.Fatal(err)
	}
	checked := 0
	for _, entry := range index.All() {
		if len(entry.RequestExample) == 0 {
			continue
		}
		checked++
		t.Run(entry.OperationID, func(t *testing.T) {
			valid, err := index.ValidExample(entry)
			if err != nil {
				t.Fatalf("published request example is invalid or unsafe: %v", err)
			}
			if !bytes.Equal(valid, entry.RequestExample) {
				t.Error("ValidExample changed the published request example")
			}
		})
	}
	if checked == 0 {
		t.Fatal("published catalog contains no request examples")
	}
	// Schema-sensitive sample sanitization is already exercised by the
	// publication tests in cmd/regencli/example_safety_test.go.
}
