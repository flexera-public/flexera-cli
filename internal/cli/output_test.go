package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestOutputGuardBeforeInputAndInitialization(t *testing.T) {
	for _, contract := range []string{"text", "binary", "mixed"} {
		for _, option := range [][]string{{"--out-jq", "."}, {"--out-fields", "id"}, {"--raw-output=false"}} {
			t.Run(contract+strings.Join(option, ""), func(t *testing.T) {
				t.Setenv("HOME", t.TempDir())
				if err := os.MkdirAll(filepath.Join(os.Getenv("HOME"), ".flexera"), 0700); err != nil {
					t.Fatal(err)
				}
				if err := os.WriteFile(filepath.Join(os.Getenv("HOME"), ".flexera", "config.yaml"), []byte("broken: ["), 0600); err != nil {
					t.Fatal(err)
				}
				var out, stderr bytes.Buffer
				root, deps := NewRootCmd(RootOptions{})
				root.SetOut(&out)
				root.SetErr(&stderr)
				root.AddCommand(&cobra.Command{Use: "raw", Annotations: map[string]string{OutputAnnotation: contract}, RunE: func(*cobra.Command, []string) error { t.Fatal("unsupported shaping executed leaf"); return nil }})
				args := append([]string{"raw"}, option...)
				if code := Execute(context.Background(), root, deps, args); code != 2 || out.Len() != 0 || !json.Valid(stderr.Bytes()) || !strings.Contains(stderr.String(), "does not support") {
					t.Fatalf("guard failed: %d %s %s", code, out.String(), stderr.String())
				}
				if deps.Stdout != nil {
					t.Fatal("guard initialized dependencies")
				}
			})
		}
	}
}

func TestTextOutputIgnoresJSONStyle(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	var out, stderr bytes.Buffer
	root, deps := NewRootCmd(RootOptions{})
	root.SetOut(&out)
	root.SetErr(&stderr)
	root.AddCommand(&cobra.Command{Use: "text", Annotations: map[string]string{OutputAnnotation: "text"}, RunE: func(cmd *cobra.Command, _ []string) error {
		_, err := cmd.OutOrStdout().Write([]byte("raw\x00bytes\n"))
		return err
	}})
	if code := Execute(context.Background(), root, deps, []string{"text", "--json-style", "pretty"}); code != 0 || out.String() != "raw\x00bytes\n" {
		t.Fatalf("text changed %d %q %s", code, out.String(), stderr.String())
	}
}

func TestTableOutputRejectsJSONShapingBeforeLeaf(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	configDir := filepath.Join(home, ".flexera")
	if err := os.MkdirAll(configDir, 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(configDir, "config.yaml"), []byte("broken: ["), 0600); err != nil {
		t.Fatal(err)
	}
	for _, option := range [][]string{{"--out-jq", "."}, {"--out-fields", "id"}} {
		t.Run(option[0], func(t *testing.T) {
			var out, stderr bytes.Buffer
			root, deps := NewRootCmd(RootOptions{})
			root.SetOut(&out)
			root.SetErr(&stderr)
			root.AddCommand(&cobra.Command{
				Use:         "leaf",
				Annotations: map[string]string{OutputAnnotation: "structured"},
				RunE:        func(*cobra.Command, []string) error { t.Fatal("incompatible output ran leaf"); return nil },
			})
			args := []string{"--access-token", "test-token", "--output", "table"}
			args = append(args, option...)
			args = append(args, "leaf")
			if code := Execute(context.Background(), root, deps, args); code != 2 || out.Len() != 0 || !strings.Contains(stderr.String(), "cannot be combined") {
				t.Fatalf("guard failed: code=%d stdout=%q stderr=%q", code, out.String(), stderr.String())
			}
			if deps.Stdout != nil {
				t.Fatal("guard initialized dependencies")
			}
		})
	}
}

func TestHelpCompletionOutputGuards(t *testing.T) {
	for _, args := range [][]string{{"--help", "--out-jq", "."}, {"leaf", "--help", "--out-fields", "id"}, {"completion", "bash", "--out-jq", "."}, {"help", "leaf", "--out-jq", "."}} {
		var out, stderr bytes.Buffer
		root, deps := NewRootCmd(RootOptions{})
		root.SetOut(&out)
		root.SetErr(&stderr)
		root.AddCommand(&cobra.Command{Use: "leaf", RunE: func(*cobra.Command, []string) error { t.Fatal("help ran leaf"); return nil }})
		if code := Execute(context.Background(), root, deps, args); code != 2 || out.Len() != 0 || !json.Valid(stderr.Bytes()) {
			t.Errorf("help/completion guard %v: %d %s %s", args, code, out.String(), stderr.String())
		}
	}
}

func TestDryRunDoesNotApplyResponseShaping(t *testing.T) {
	code, err := CompileJQ(`error("must not shape plan")`)
	if err != nil {
		t.Fatal(err)
	}
	printer := Printer{Style: JSONStyleCompact, JQ: code, Fields: []FieldPath{{"discard-plan"}}, RawOutput: true}
	var out bytes.Buffer
	done, err := ConfirmWrite(true, false, true, &out, map[string]any{"method": "DELETE /test"}, printer)
	if err != nil || !done {
		t.Fatalf("plan failed %v", err)
	}
	var plan map[string]any
	if err := json.Unmarshal(out.Bytes(), &plan); err != nil {
		t.Fatal(err)
	}
	if plan["dryRun"] != true || plan["destructive"] != true || plan["plan"] == nil {
		t.Fatalf("shaped dryrun: %s", out.String())
	}
}
