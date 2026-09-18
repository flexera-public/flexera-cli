package rulebaseddimension

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
)

func TestNewBulkCmdFlags(t *testing.T) {
	cmd := NewBulkCmd()
	for _, name := range []string{"input", "file", "dry-run", "continue-on-error"} {
		if cmd.Flags().Lookup(name) == nil {
			t.Errorf("expected --%s flag", name)
		}
	}
}

func TestResolveInput(t *testing.T) {
	raw, err := resolveInput(`{"dimensions":[]}`, "", bytes.NewReader(nil))
	if err != nil {
		t.Fatalf("resolveInput: %v", err)
	}
	if string(raw) != `{"dimensions":[]}` {
		t.Fatalf("unexpected input: %s", raw)
	}
}

func TestAttach(t *testing.T) {
	root := &cobra.Command{Use: "root"}
	generated := &cobra.Command{Use: "rule-based-dimension"}
	root.AddCommand(generated)
	if err := Attach(root); err != nil {
		t.Fatalf("Attach: %v", err)
	}
	if generated.Commands()[0].Name() != "bulk" {
		t.Fatalf("expected bulk command, got %q", generated.Commands()[0].Name())
	}
	if len(generated.Aliases) != 1 || generated.Aliases[0] != "rule-base-dimension" {
		t.Fatalf("expected compatibility alias, got %v", generated.Aliases)
	}
}
