package app_test

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/flexera-public/flexera-cli/internal/app"
	"github.com/flexera-public/flexera-cli/internal/catalog"
)

func TestSearchGoldenQueries(t *testing.T) {
	root, _ := app.NewRootCmd(io.Discard, io.Discard, func(string) string { return "" }, offlineDoer{t}, "test")
	index, err := catalog.NewSearchIndex(root)
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct{ query, want string }{{"delete a budget", "flexera-cli budget delete"}, {"list cloud bill connects", "flexera-cli finops-onboarding bill-connect list"}, {"who has access to org", "flexera-cli iam access-policy users"}, {"run a policy", "flexera-cli policy applied-policy evaluate"}} {
		results := index.Search(tc.query, catalog.SearchOptions{Limit: 3})
		found := false
		for _, r := range results {
			found = found || r.Command == tc.want
		}
		if !found {
			t.Errorf("%q expected %s in top3: %+v", tc.query, tc.want, results)
		}
	}
}

func TestOfflineSearchOutput(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	t.Setenv("FLEXERA_CLI_CONFIG", "")
	t.Setenv("FLEXERA_CLI_ZONE", "invalid")
	t.Setenv("FLEXERA_CLI_OUTPUT", "json")
	t.Setenv("FLEXERA_CLI_JSON_STYLE", "compact")
	if err := os.MkdirAll(filepath.Join(os.Getenv("HOME"), ".flexera"), 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(os.Getenv("HOME"), ".flexera", "config.yaml"), []byte("invalid: ["), 0600); err != nil {
		t.Fatal(err)
	}
	code, out, stderr := schemaRun(t, "cli", "search", "delete budget", "--limit", "1")
	if code != 0 {
		t.Fatalf("%d %s", code, stderr)
	}
	var results []catalog.SearchResult
	if err := json.Unmarshal([]byte(out), &results); err != nil || len(results) != 1 || results[0].Command != "flexera-cli budget delete" {
		t.Fatalf("invalid output %s %v", out, err)
	}
	code, out, stderr = schemaRun(t, "cli", "search", "delete budget", "--limit", "1", "-o", "table")
	if code != 0 || !strings.Contains(out, "COMMAND") || !strings.Contains(out, "budget delete") {
		t.Fatalf("table: %d %s %s", code, out, stderr)
	}
	code, out, stderr = schemaRun(t, "cli", "search", "zzzznonexistent")
	if code != 0 || strings.TrimSpace(out) != "[]" {
		t.Fatalf("empty: %d %s %s", code, out, stderr)
	}
	code, out, stderr = schemaRun(t, "cli", "search", "delete budget", "--limit", "1", "--out-jq", ".[0].command", "--raw-output")
	if code != 0 || strings.TrimSpace(out) != "flexera-cli budget delete" {
		t.Fatalf("shaped: %d %s %s", code, out, stderr)
	}
	for _, args := range [][]string{{"cli", "search"}, {"cli", "search", "budget", "--limit", "0"}} {
		code, out, _ := schemaRun(t, args...)
		if code != 2 || out != "" {
			t.Errorf("invalid args accepted: %v", args)
		}
	}
}

func BenchmarkSearchIndexBuild(b *testing.B) {
	root, _ := app.NewRootCmd(io.Discard, io.Discard, func(string) string { return "" }, nil, "benchmark")
	if _, err := catalog.Load(); err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for b.Loop() {
		if _, err := catalog.NewSearchIndex(root); err != nil {
			b.Fatal(err)
		}
	}
}

func TestSearchReadOnlyAndSynopsis(t *testing.T) {
	root, _ := app.NewRootCmd(io.Discard, io.Discard, func(string) string { return "" }, nil, "test")
	index, err := catalog.NewSearchIndex(root)
	if err != nil {
		t.Fatal(err)
	}
	for _, result := range index.Search("policy", catalog.SearchOptions{ReadOnly: true, Limit: 1000}) {
		if strings.Contains(result.Command, " evaluate") || strings.Contains(result.Command, " create") || strings.Contains(result.Command, " delete") || strings.Contains(result.Command, " update") {
			t.Errorf("write included in read-only: %s", result.Command)
		}
	}
	results := index.Search("budget create", catalog.SearchOptions{Tag: "Budget", Action: "create"})
	if len(results) != 1 || !strings.Contains(results[0].Usage, "--org-id ORG_ID") || !strings.Contains(results[0].Usage, "--body BODY_JSON_OR_@FILE_OR_@-") {
		t.Fatalf("missing required synopsis input: %+v", results)
	}
	for _, result := range index.Search("verify CSV", catalog.SearchOptions{ReadOnly: true, Limit: 1000}) {
		if result.Command == "flexera-cli bill-upload verify" && result.Schema != "" {
			t.Fatal("curated-only result has schema")
		}
	}
}

func TestSearchServiceFilter(t *testing.T) {
	root, _ := app.NewRootCmd(io.Discard, io.Discard, func(string) string { return "" }, offlineDoer{t}, "test")
	index, err := catalog.NewSearchIndex(root)
	if err != nil {
		t.Fatal(err)
	}
	for _, service := range []string{"bill-analysis", "BA", "bill_analysis"} {
		results := index.Search("list costs", catalog.SearchOptions{Limit: 20, Service: service})
		if len(results) == 0 {
			t.Fatalf("%s: no results", service)
		}
		for _, r := range results {
			if !strings.HasPrefix(r.Command, "flexera-cli bill-analysis ") {
				t.Errorf("%s: result outside service: %s", service, r.Command)
			}
		}
	}
	if results := index.Search("list projects", catalog.SearchOptions{Limit: 20, Service: "grs"}); len(results) < 2 {
		t.Fatalf("curated and generated grs commands not both indexed: %+v", results)
	}
}
