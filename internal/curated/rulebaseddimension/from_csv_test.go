package rulebaseddimension

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestNewFromCSVCommand(t *testing.T) {
	cmd := NewFromCSVCommand()
	if cmd.Name() != "from_csv" {
		t.Fatalf("unexpected command name %q", cmd.Name())
	}
	for _, name := range []string{"generate", "create", "update"} {
		if _, _, err := cmd.Find([]string{name}); err != nil {
			t.Errorf("missing %s command: %v", name, err)
		}
	}
}

func TestCSVFlagsOptions(t *testing.T) {
	flags := csvFlags{
		separatorHeader: "SPLIT",
		idTemplate:      "rbd_{{.ColumnSlug}}",
		nameTemplate:    "{{.ColumnHeader}}",
		effectiveAt:     "2025-01",
		columnConfig:    `{"Department":{"id":"rbd_department","skip":false}}`,
		caseInsensitive: false,
	}
	opts, err := flags.options()
	if err != nil {
		t.Fatalf("options: %v", err)
	}
	if opts.SeparatorHeader != "SPLIT" || opts.EffectiveAt != "2025-01" || opts.CaseInsensitive {
		t.Fatalf("unexpected options: %+v", opts)
	}
	if opts.ColumnConfig["Department"].ID != "rbd_department" {
		t.Fatalf("unexpected column config: %+v", opts.ColumnConfig)
	}
}

func TestApplyCommandsExposeCSVFlags(t *testing.T) {
	cmd := NewFromCSVCommand()
	for _, name := range []string{"create", "update"} {
		child, _, err := cmd.Find([]string{name})
		if err != nil {
			t.Fatal(err)
		}
		if child.Flags().Lookup("file") == nil || child.Flags().Lookup("dry-run") == nil {
			t.Errorf("%s is missing CSV/apply flags", name)
		}
	}
}

func TestCommandCanBeAddedToRoot(t *testing.T) {
	root := &cobra.Command{Use: "root"}
	generated := &cobra.Command{Use: "rule-based-dimension"}
	root.AddCommand(generated)
	if err := Attach(root); err != nil {
		t.Fatalf("attach: %v", err)
	}
	if _, _, err := root.Find([]string{"rule-based-dimension", "from_csv", "generate"}); err != nil {
		t.Fatalf("find command: %v", err)
	}
}
