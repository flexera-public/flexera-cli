package catalog

import (
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestSearchCorpusSafetyAndFilters(t *testing.T) {
	root := &cobra.Command{Use: "flexera-cli"}
	group := &cobra.Command{Use: "budget"}
	root.AddCommand(group)
	for _, leaf := range []struct{ verb, id string }{{"list", "Budget_Budget_index"}, {"delete", "Budget_Budget_delete"}} {
		// Use the real catalog identity rather than guessing IDs.
		all, err := All()
		if err != nil {
			t.Fatal(err)
		}
		id := ""
		for _, e := range all {
			if strings.Join(e.Command, " ") == "budget "+leaf.verb {
				id = e.OperationID
			}
		}
		group.AddCommand(&cobra.Command{Use: leaf.verb, Aliases: []string{leaf.verb + "-alias"}, Short: leaf.verb + " budgets", Annotations: map[string]string{"flexera.operationId": id}, RunE: func(*cobra.Command, []string) error { t.Fatal("index executed a leaf"); return nil }})
	}
	root.AddCommand(&cobra.Command{Use: "inspect", Short: "Inspect local budgets", Annotations: map[string]string{"flexera.readOnly": "true"}, RunE: func(*cobra.Command, []string) error { return nil }})
	root.AddCommand(&cobra.Command{Use: "unknown", Short: "Unknown budget side effects", RunE: func(*cobra.Command, []string) error { return nil }})
	root.AddCommand(&cobra.Command{Use: "hidden", Hidden: true, Short: "budget", RunE: func(*cobra.Command, []string) error { return nil }})
	index, err := NewSearchIndex(root)
	if err != nil {
		t.Fatal(err)
	}
	results := index.Search("destroy budget", SearchOptions{Limit: 10})
	if len(results) == 0 || results[0].Command != "flexera-cli budget delete" || !results[0].Destructive || !strings.Contains(results[0].Usage, "--org-id ORG_ID") {
		t.Fatalf("bad results: %+v", results)
	}
	for _, result := range index.Search("budget", SearchOptions{ReadOnly: true, Limit: 10}) {
		if result.Command == "flexera-cli budget delete" || result.Command == "flexera-cli unknown" {
			t.Fatalf("unsafe read-only result: %+v", result)
		}
		if result.Command == "flexera-cli inspect" && result.Schema != "" {
			t.Fatal("curated-only schema link")
		}
	}
	filtered := index.Search("budget", SearchOptions{Tag: "Budget", Action: "list", Limit: 1})
	if len(filtered) != 1 || filtered[0].Command != "flexera-cli budget list" {
		t.Fatalf("bad filters: %+v", filtered)
	}
	if got := index.Search("zzzznonexistent", SearchOptions{}); len(got) != 0 || got == nil {
		t.Fatalf("no match should be []: %+v", got)
	}
}

func TestSearchNormalization(t *testing.T) {
	joined := strings.Join(tokens("cloud account removeBudget spend fetch"), " ")
	for _, word := range []string{"connector", "delete", "budget", "cost", "get"} {
		if !strings.Contains(joined, word) {
			t.Errorf("missing %s: %s", word, joined)
		}
	}
}
