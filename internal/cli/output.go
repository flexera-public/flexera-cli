package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

const OutputAnnotation = "flexera.output"

// ValidateOutputOptions checks command contracts, not response bytes. Text that
// happens to look like JSON must never become a shaped API response.
func ValidateOutputOptions(cmd *cobra.Command) error {
	flags := cmd.Root().PersistentFlags()
	if flags.Changed(FlagOutput) {
		format, _ := flags.GetString(FlagOutput)
		if err := validateTableShaping(cmd, format); err != nil {
			return err
		}
	}
	shaping := flags.Changed(FlagOutJQ) || flags.Changed(FlagOutFields) || flags.Changed(FlagRawOutput)
	if !shaping {
		return nil
	}
	contract := cmd.Annotations[OutputAnnotation]
	if cmd.HasSubCommands() {
		contract = "text"
	}
	for p := cmd; p != nil; p = p.Parent() {
		if p.Name() == "completion" || p.Name() == "help" || p.Name() == "__complete" || p.Name() == "__completeNoDesc" {
			contract = "text"
			break
		}
	}
	if help, _ := cmd.Flags().GetBool("help"); help {
		contract = "text"
	}
	if version, _ := cmd.Flags().GetBool("version"); version {
		contract = "text"
	}
	switch contract {
	case "text", "binary", "mixed":
		return Exit(2, fmt.Errorf("%s preserves %s output and does not support --out-jq, --out-fields or --raw-output", cmd.CommandPath(), contract))
	}
	return nil
}

func validateTableShaping(cmd *cobra.Command, format string) error {
	flags := cmd.Root().PersistentFlags()
	if strings.EqualFold(strings.TrimSpace(format), "table") && (flags.Changed(FlagOutJQ) || flags.Changed(FlagOutFields)) {
		return Exit(2, fmt.Errorf("--output table cannot be combined with --out-jq or --out-fields"))
	}
	return nil
}
