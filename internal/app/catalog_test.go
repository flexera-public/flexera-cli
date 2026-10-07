package app_test

import (
	"encoding/json"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/flexera-public/flexera-cli/internal/app"
	"github.com/flexera-public/flexera-cli/internal/catalog"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func TestCatalogMatchesLiveTreeBothDirections(t *testing.T) {
	root, _ := app.NewRootCmd(io.Discard, io.Discard, func(string) string { return "" }, nil, "test")
	entries, err := catalog.All()
	if err != nil {
		t.Fatal(err)
	}
	expected := map[string]catalog.Entry{}
	for _, e := range entries {
		expected[e.OperationID] = e
	}
	seen := map[string]bool{}
	var walk func(*cobra.Command)
	walk = func(c *cobra.Command) {
		if id := c.Annotations["flexera.operationId"]; id != "" {
			e, ok := expected[id]
			if !ok {
				t.Errorf("annotated command has no catalog: %s", c.CommandPath())
			}
			path := strings.TrimPrefix(c.CommandPath(), root.Name()+" ")
			if seen[id] || path != strings.Join(e.Command, " ") {
				t.Errorf("catalog/tree mismatch for %s: %s vs %v", id, path, e.Command)
			}
			seen[id] = true
			for _, p := range e.Params {
				if c.Flags().Lookup(p.Flag) == nil && c.InheritedFlags().Lookup(p.Flag) == nil {
					t.Errorf("%s catalog flag --%s does not exist", path, p.Flag)
				}
			}
		}
		for _, child := range c.Commands() {
			walk(child)
		}
	}
	walk(root)
	for id, e := range expected {
		if !seen[id] {
			t.Errorf("catalog has no command: %s %v", id, e.Command)
		}
	}
}

func TestNoLocalFlagsShadowInheritedFlags(t *testing.T) {
	root, _ := app.NewRootCmd(io.Discard, io.Discard, func(string) string { return "" }, nil, "test")
	var walk func(*cobra.Command)
	walk = func(c *cobra.Command) {
		if c != root {
			c.LocalNonPersistentFlags().VisitAll(func(f *pflag.Flag) {
				for p := c.Parent(); p != nil; p = p.Parent() {
					p.PersistentFlags().VisitAll(func(inherited *pflag.Flag) {
						if f.Name == inherited.Name || f.Shorthand != "" && f.Shorthand == inherited.Shorthand {
							t.Errorf("%s --%s shadows %s --%s", c.CommandPath(), f.Name, p.CommandPath(), inherited.Name)
						}
					})
				}
			})
		}
		for _, child := range c.Commands() {
			walk(child)
		}
	}
	walk(root)
}

func TestAllAnnotatedSpecOperationsHaveCatalogEntries(t *testing.T) {
	data, err := os.ReadFile("../../unified-openapi/openapi3.json")
	if err != nil {
		t.Fatal(err)
	}
	var spec struct {
		Paths map[string]map[string]json.RawMessage
	}
	if err := json.Unmarshal(data, &spec); err != nil {
		t.Fatal(err)
	}
	entries, err := catalog.All()
	if err != nil {
		t.Fatal(err)
	}
	ids := map[string]bool{}
	for _, e := range entries {
		ids[e.OperationID] = true
	}
	count := 0
	for _, pi := range spec.Paths {
		for _, raw := range pi {
			var op map[string]json.RawMessage
			if json.Unmarshal(raw, &op) != nil {
				continue
			}
			if _, ok := op["x-flexera-action"]; !ok {
				continue
			}
			var id string
			_ = json.Unmarshal(op["operationId"], &id)
			if !ids[id] {
				t.Errorf("spec operation missing from catalog: %s", id)
			}
			count++
		}
	}
	if count != len(entries) {
		t.Errorf("spec/catalog counts differ: %d/%d", count, len(entries))
	}
}
