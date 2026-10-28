package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// FindCommand returns the registered descendant of root at path (canonical
// names only). Curated commands use it to attach beneath generated service
// and tag commands.
func FindCommand(root *cobra.Command, path ...string) (*cobra.Command, error) {
	current := root
	for i, name := range path {
		var next *cobra.Command
		for _, child := range current.Commands() {
			if child.Name() == name {
				next = child
				break
			}
		}
		if next == nil {
			return nil, fmt.Errorf("command %q is not registered", strings.Join(path[:i+1], " "))
		}
		current = next
	}
	return current, nil
}

// AttachCommands adds children beneath the registered command at path.
func AttachCommands(root *cobra.Command, path []string, children ...*cobra.Command) error {
	parent, err := FindCommand(root, path...)
	if err != nil {
		return err
	}
	parent.AddCommand(children...)
	return nil
}

// AddCommandAt adds leaf beneath root, creating any missing group commands
// along groups. Generated tag commands use it to nest sub-resource commands
// (saas license fees list).
func AddCommandAt(root *cobra.Command, groups []string, leaf *cobra.Command) {
	parent := root
	for _, name := range groups {
		group, err := FindCommand(parent, name)
		if err != nil {
			group = &cobra.Command{Use: name, Short: name + " operations"}
			parent.AddCommand(group)
		}
		parent = group
	}
	parent.AddCommand(leaf)
}
