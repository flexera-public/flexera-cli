package app_test

import (
	"io"
	"testing"

	"github.com/flexera-public/flexera-cli/internal/app"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func TestOutputShapingFlagsDoNotCollideWithLocalFlags(t *testing.T) {
	root, _ := app.NewRootCmd(io.Discard, io.Discard, func(string) string { return "" }, nil, "test")
	reserved := map[string]bool{"out-jq": true, "out-fields": true, "raw-output": true}
	var walk func(*cobra.Command)
	walk = func(cmd *cobra.Command) {
		cmd.LocalNonPersistentFlags().VisitAll(func(flag *pflag.Flag) {
			if reserved[flag.Name] {
				t.Errorf("%s local --%s collides with a new persistent output flag", cmd.CommandPath(), flag.Name)
			}
		})
		for _, child := range cmd.Commands() {
			walk(child)
		}
	}
	walk(root)
}
