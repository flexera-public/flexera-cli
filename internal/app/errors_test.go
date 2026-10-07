package app_test

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/flexera-public/flexera-cli/internal/app"
	"github.com/flexera-public/flexera-cli/internal/cli"
)

func TestUnknownCommandSuggestionsAreOffline(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	if err := os.MkdirAll(filepath.Join(os.Getenv("HOME"), ".flexera"), 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(os.Getenv("HOME"), ".flexera", "config.yaml"), []byte("broken: ["), 0600); err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		args                []string
		parent, token, want string
	}{
		{[]string{"delete", "budget"}, "flexera-cli", "delete", "flexera-cli budget delete"},
		{[]string{"policy", "applied-policy", "run"}, "flexera-cli policy applied-policy", "run", "flexera-cli policy applied-policy evaluate"},
		{[]string{"finops", "cost", `unknown"quoted`}, "flexera-cli finops cost", `unknown"quoted`, ""},
		{[]string{"zzzznonexistenttoken"}, "flexera-cli", "zzzznonexistenttoken", ""},
	} {
		t.Run(strings.Join(tc.args, "/"), func(t *testing.T) {
			var out, errOut bytes.Buffer
			root, deps := app.NewRootCmd(&out, &errOut, os.Getenv, offlineDoer{t}, "test")
			code := cli.Execute(context.Background(), root, deps, tc.args)
			var value struct {
				Error, Parent       string
				Tokens, Suggestions []string
			}
			if err := json.Unmarshal(errOut.Bytes(), &value); err != nil {
				t.Fatalf("invalid JSON: %v %s", err, errOut.String())
			}
			if code != 2 || out.Len() != 0 || value.Parent != tc.parent || len(value.Tokens) == 0 || value.Tokens[0] != tc.token || value.Suggestions == nil || len(value.Suggestions) > 3 {
				t.Fatalf("bad execution error: %d %+v", code, value)
			}
			if !strings.Contains(value.Error, "unknown command") || strings.Contains(value.Error, "reading config") {
				t.Fatalf("preinit error replaced: %s", value.Error)
			}
			if tc.want != "" {
				found := false
				for _, suggestion := range value.Suggestions {
					found = found || suggestion == tc.want
				}
				if !found {
					t.Errorf("missing expected suggestion %s: %v", tc.want, value.Suggestions)
				}
			}
			if tc.parent == "flexera-cli" && tc.want == "" && len(value.Suggestions) != 0 {
				t.Fatalf("no match should have empty suggestions: %v", value.Suggestions)
			}
		})
	}
}

func TestEarlyErrorsRetainCobraText(t *testing.T) {
	var out, errOut bytes.Buffer
	root, deps := app.NewRootCmd(&out, &errOut, func(string) string { return "" }, offlineDoer{t}, "test")
	args := []string{"budge"}
	_, _, original := root.Find(args)
	if original == nil {
		t.Fatal("fixture should be an unknown command")
	}
	if code := cli.Execute(context.Background(), root, deps, args); code != 2 {
		t.Fatalf("exit=%d %s", code, errOut.String())
	}
	var value struct{ Error string }
	if err := json.Unmarshal(errOut.Bytes(), &value); err != nil {
		t.Fatal(err)
	}
	if value.Error != original.Error() {
		t.Fatalf("Cobra error text changed: %q vs %q", value.Error, original.Error())
	}
}

func TestMalformedFlagsBeforeInitialization(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	for _, args := range [][]string{{"budget", "list", "--missing"}, {"budget", "list", "--org-id", "invalid"}, {"cli", "search", "budget", "--limit"}, {"cli", "schema", "budget", "create", "--depth", "wrong"}} {
		var out, errOut bytes.Buffer
		root, deps := app.NewRootCmd(&out, &errOut, os.Getenv, offlineDoer{t}, "test")
		if code := cli.Execute(context.Background(), root, deps, args); code != 2 || out.Len() != 0 || !json.Valid(errOut.Bytes()) {
			t.Errorf("parse error %v: %d %s", args, code, errOut.String())
		}
		if deps.Stdout != nil {
			t.Errorf("root initialized for parse failure %v", args)
		}
	}
}

func TestBareHelpAndVersionShaping(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	if err := os.MkdirAll(filepath.Join(os.Getenv("HOME"), ".flexera"), 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(os.Getenv("HOME"), ".flexera", "config.yaml"), []byte("broken: ["), 0600); err != nil {
		t.Fatal(err)
	}
	for _, args := range [][]string{{}, {"budget"}, {"policy"}, {"cli"}} {
		code, out, stderr := schemaRun(t, args...)
		if code != 0 || out == "" {
			t.Errorf("bare help blocked %v: %d %s", args, code, stderr)
		}
	}
	code, out, stderr := schemaRun(t, "--version", "--out-jq", ".")
	if code != 2 || out != "" || !strings.Contains(stderr, "preserves text output") {
		t.Fatalf("version silently ignored shaping: %d %s %s", code, out, stderr)
	}
}
