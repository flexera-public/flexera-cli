package billupload

import (
	"testing"

	"github.com/spf13/cobra"
)

// TestNewPushCmdFlags verifies the curated push command exposes the flags
// the workflow depends on and is registered as a distinct command from the
// generated "files" subcommand.
func TestNewPushCmdFlags(t *testing.T) {
	cmd := NewPushCmd()
	if cmd.Use != "push" {
		t.Fatalf("expected Use=push, got %q", cmd.Use)
	}
	for _, name := range []string{"bill-connect-id", "billing-period", "file"} {
		if cmd.Flags().Lookup(name) == nil {
			t.Errorf("expected --%s flag to be registered", name)
		}
	}
}

// TestAttachRejectsMissingGeneratedCommand verifies Attach fails loudly
// rather than silently no-op-ing when the generated "bill-upload" command
// tree is not present (e.g. if upstream regeneration ever renames/removes
// it, as happened once already -- see UPSTREAM-BILL-UPLOAD.md).
func TestAttachRejectsMissingGeneratedCommand(t *testing.T) {
	root := &cobra.Command{Use: "root"}
	if err := Attach(root); err == nil {
		t.Fatal("expected error when bill-upload generated command is absent")
	}
}

// TestAttachAddsPushUnderBillUpload verifies Attach wires the curated push
// command as a child of the generated bill-upload command, alongside (not
// replacing) any generated subcommands such as "files".
func TestAttachAddsPushUnderBillUpload(t *testing.T) {
	root := &cobra.Command{Use: "root"}
	billUpload := &cobra.Command{Use: "bill-upload"}
	billUpload.AddCommand(&cobra.Command{Use: "files"})
	root.AddCommand(billUpload)

	if err := Attach(root); err != nil {
		t.Fatalf("Attach: %v", err)
	}

	var names []string
	for _, c := range billUpload.Commands() {
		names = append(names, c.Name())
	}
	hasPush, hasFiles := false, false
	for _, n := range names {
		hasPush = hasPush || n == "push"
		hasFiles = hasFiles || n == "files"
	}
	if !hasPush {
		t.Errorf("expected push command to be attached, got children %v", names)
	}
	if !hasFiles {
		t.Errorf("expected generated files command to remain attached, got children %v", names)
	}
}
