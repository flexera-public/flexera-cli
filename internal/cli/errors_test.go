package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestExecutionErrorContracts(t *testing.T) {
	for _, tc := range []struct {
		name     string
		args     []string
		runErr   error
		code     int
		fragment string
	}{
		{"unknown", []string{"delete", "budget"}, nil, 2, "unknown command"},
		{"flag", []string{"leaf", "--missing"}, nil, 2, "unknown flag"},
		{"argument", []string{"leaf", "--value"}, nil, 2, "flag needs an argument"},
		{"bad-bool", []string{"leaf", "--enabled=bad"}, nil, 2, "invalid argument"},
		{"runtime-like-parse", []string{"leaf"}, errors.New("unknown flag: API response text"), 1, "API response text"},
		{"runtime", []string{"leaf"}, errors.New("runtime failure"), 1, "runtime failure"},
		{"pretty-early", []string{"--json-style", "pretty", "leaf", "--missing"}, nil, 2, "unknown flag"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv("HOME", t.TempDir())
			var out, errOut bytes.Buffer
			root, deps := NewRootCmd(RootOptions{})
			root.SetOut(&out)
			root.SetErr(&errOut)
			leaf := &cobra.Command{Use: "leaf", Args: cobra.NoArgs, RunE: func(*cobra.Command, []string) error { return tc.runErr }}
			leaf.Flags().String("value", "", "")
			leaf.Flags().Bool("enabled", false, "")
			root.AddCommand(leaf)
			code := Execute(context.Background(), root, deps, tc.args)
			var value map[string]any
			if err := json.Unmarshal(errOut.Bytes(), &value); err != nil {
				t.Fatalf("invalid error JSON %s: %v", errOut.String(), err)
			}
			if code != tc.code || !strings.Contains(value["error"].(string), tc.fragment) || out.Len() != 0 {
				t.Fatalf("code=%d err=%s out=%s", code, errOut.String(), out.String())
			}
			if tc.name == "unknown" {
				if _, ok := value["suggestions"].([]any); !ok {
					t.Fatal("missing suggestions array")
				}
			}
			if tc.name == "pretty-early" && !strings.Contains(errOut.String(), "\n  ") {
				t.Fatal("early pretty flag ignored")
			}
		})
	}
}

func TestStructuredValidationSurvivesWrapping(t *testing.T) {
	detail := &ValidationError{Details: []ValidationDetail{{Path: "/name", Message: `property "name" is missing`}}, Schema: "flexera-cli cli schema budget create"}
	err := Exit(2, fmt.Errorf("wrapper: %w", detail))
	var found *ValidationError
	if !errors.As(err, &found) {
		t.Fatal("lost typed error")
	}
	var out bytes.Buffer
	code, compileErr := CompileJQ(`error("must not execute")`)
	if compileErr != nil {
		t.Fatal(compileErr)
	}
	if err := (Printer{JQ: code, Fields: []FieldPath{{"anything"}}}).RenderError(&out, err); err != nil {
		t.Fatal(err)
	}
	var value map[string]any
	if err := json.Unmarshal(out.Bytes(), &value); err != nil {
		t.Fatal(err)
	}
	if value["schema"] != detail.Schema || len(value["details"].([]any)) != 1 {
		t.Fatalf("lost structured details: %s", out.String())
	}
}

func TestInvalidConfigAndJQAreUsageErrors(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	path := filepath.Join(t.TempDir(), "config.yaml")
	if err := os.WriteFile(path, []byte("broken: ["), 0600); err != nil {
		t.Fatal(err)
	}
	for _, args := range [][]string{{"leaf", "--config", path}, {"leaf", "--out-jq", "["}, {"leaf", "--json-style", "invalid"}} {
		var out, errOut bytes.Buffer
		root, deps := NewRootCmd(RootOptions{})
		root.SetOut(&out)
		root.SetErr(&errOut)
		root.AddCommand(&cobra.Command{Use: "leaf", RunE: func(*cobra.Command, []string) error { t.Fatal("invalid configuration ran command"); return nil }})
		if code := Execute(context.Background(), root, deps, args); code != 2 || !json.Valid(errOut.Bytes()) || out.Len() != 0 {
			t.Fatalf("invalid usage output: %d %s", code, errOut.String())
		}
		if len(args) > 1 && args[1] == "--out-jq" && !strings.Contains(errOut.String(), `"position"`) {
			t.Fatalf("missing jq position: %s", errOut.String())
		}
	}
}
