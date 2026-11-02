package cli

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestAttachCommands(t *testing.T) {
	root := &cobra.Command{Use: "root"}
	service := &cobra.Command{Use: "svc", Aliases: []string{"s"}}
	tag := &cobra.Command{Use: "tag"}
	service.AddCommand(tag)
	root.AddCommand(service)

	leaf := &cobra.Command{Use: "leaf"}
	if err := AttachCommands(root, []string{"svc", "tag"}, leaf); err != nil {
		t.Fatal(err)
	}
	if leaf.Parent() != tag {
		t.Fatal("leaf not attached under svc tag")
	}
	if _, err := FindCommand(root, "s", "tag"); err == nil {
		t.Fatal("aliases must not resolve; attach paths are canonical")
	}
	if err := AttachCommands(root, []string{"svc", "missing"}, &cobra.Command{Use: "x"}); err == nil || err.Error() != `command "svc missing" is not registered` {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestAddCommandAtCreatesAndReusesGroups(t *testing.T) {
	root := &cobra.Command{Use: "license"}
	list := &cobra.Command{Use: "list"}
	get := &cobra.Command{Use: "get"}
	AddCommandAt(root, []string{"fees"}, list)
	AddCommandAt(root, []string{"fees"}, get)
	AddCommandAt(root, nil, &cobra.Command{Use: "delete"})

	if list.Parent() != get.Parent() || list.Parent().Name() != "fees" || list.Parent().Parent() != root {
		t.Fatal("leaves must share one fees group under root")
	}
	if got := len(root.Commands()); got != 2 {
		t.Fatalf("root has %d children, want fees and delete", got)
	}
}
