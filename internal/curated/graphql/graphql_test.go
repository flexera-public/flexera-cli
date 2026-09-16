package graphql

import (
	"testing"

	"github.com/spf13/cobra"
)

// TestNewQueryCmdFlags verifies the raw /explore/graphql query command
// exposes the flags the workflow depends on.
func TestNewQueryCmdFlags(t *testing.T) {
	cmd := NewQueryCmd()
	if cmd.Use != "query" {
		t.Fatalf("expected Use=query, got %q", cmd.Use)
	}
	for _, name := range []string{"body", "path"} {
		if cmd.Flags().Lookup(name) == nil {
			t.Errorf("expected --%s flag to be registered", name)
		}
	}
}

// TestNewGenerateCmdFlags verifies the GenerateQuery command exposes the
// flags the workflow depends on.
func TestNewGenerateCmdFlags(t *testing.T) {
	cmd := NewGenerateCmd()
	if cmd.Use != "generate" {
		t.Fatalf("expected Use=generate, got %q", cmd.Use)
	}
	for _, name := range []string{"prompt", "query", "modify-prompt", "indent"} {
		if cmd.Flags().Lookup(name) == nil {
			t.Errorf("expected --%s flag to be registered", name)
		}
	}
}

// TestAttachRejectsMissingGeneratedCommand verifies Attach fails loudly
// rather than silently no-op-ing when the generated "graphql" command tree
// is not present.
func TestAttachRejectsMissingGeneratedCommand(t *testing.T) {
	root := &cobra.Command{Use: "root"}
	if err := Attach(root); err == nil {
		t.Fatal("expected error when graphql generated command is absent")
	}
}

// TestAttachAddsQueryAndGenerate verifies Attach wires both curated
// subcommands as children of the generated graphql command.
func TestAttachAddsQueryAndGenerate(t *testing.T) {
	root := &cobra.Command{Use: "root"}
	graphqlCmd := &cobra.Command{Use: "graphql"}
	root.AddCommand(graphqlCmd)

	if err := Attach(root); err != nil {
		t.Fatalf("Attach: %v", err)
	}

	hasQuery, hasGenerate := false, false
	for _, c := range graphqlCmd.Commands() {
		hasQuery = hasQuery || c.Name() == "query"
		hasGenerate = hasGenerate || c.Name() == "generate"
	}
	if !hasQuery {
		t.Error("expected query command to be attached")
	}
	if !hasGenerate {
		t.Error("expected generate command to be attached")
	}
}
