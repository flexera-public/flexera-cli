package cli

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/flexera-public/flexera-cli/internal/catalog"
	cliconfig "github.com/flexera-public/flexera-cli/internal/config"
	flexera "github.com/flexera-public/unified-go-client"
)

// RootOptions configures NewRootCmd. All fields are optional; zero values
// fall back to production defaults.
type RootOptions struct {
	// Use is the root command name (default "flexera-cli").
	Use string
	// Short/Long describe the command in help output.
	Short string
	Long  string
	// BaseHTTP is the underlying request doer before any debug wrapping
	// (default &http.Client{}). Injected for tests.
	BaseHTTP flexera.HttpRequestDoer
	// Version string reported by `--version`.
	Version string
	// Getenv reads environment variables (default os.Getenv). Injected for
	// tests so curated commands honoring ad-hoc env vars stay testable.
	Getenv func(string) string
}

// NewRootCmd builds the cobra root command: it declares the persistent
// configuration flags, wires viper in PersistentPreRunE (precedence
// flag > env > config file > default), constructs the authenticated HTTP
// client (debug-wrapped when --debug/FLEXERA_CLI_DEBUG is set), and stashes
// the resulting *Deps on the command context for every subcommand's RunE.
//
// The returned *Deps pointer is the same instance attached to the context;
// callers (and tests) may inspect it after Execute.
func NewRootCmd(opts RootOptions) (*cobra.Command, *Deps) {
	use := opts.Use
	if use == "" {
		use = "flexera-cli"
	}
	deps := &Deps{}

	root := &cobra.Command{
		Use:           use,
		Short:         opts.Short,
		Long:          opts.Long,
		Version:       opts.Version,
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	RegisterPersistentFlags(root.PersistentFlags())

	root.PersistentPreRunE = func(cmd *cobra.Command, _ []string) error {
		if cmd.HasSubCommands() {
			return ValidateOutputOptions(cmd)
		}
		if cmd.Root().PersistentFlags().Changed(FlagNoValidate) && cmd.Annotations["flexera.validation"] == "" {
			return Exit(2, fmt.Errorf("%s does not support --no-validate; request validation is not enabled for this command", cmd.CommandPath()))
		}
		if err := ValidateOutputOptions(cmd); err != nil {
			return err
		}
		offline := cmd.Annotations["flexera.offline"] == "true"
		cfgFile, _ := cmd.Flags().GetString(FlagConfig)
		getenv := opts.Getenv
		if getenv == nil {
			getenv = os.Getenv
		}
		if offline && cfgFile == "" {
			cfgFile = getenv("FLEXERA_CLI_CONFIG")
		}
		v, err := newViper(cmd.Root().PersistentFlags(), cfgFile, !offline)
		if err != nil {
			return Exit(2, err)
		}
		if offline {
			for _, key := range []string{FlagOutput, FlagJSONStyle} {
				if value := getenv("FLEXERA_CLI_" + strings.ToUpper(strings.ReplaceAll(key, "-", "_"))); value != "" && !cmd.Root().PersistentFlags().Changed(key) {
					v.Set(key, value)
				}
			}
		}
		var cfg cliconfig.CommonConfig
		if offline {
			cfg.Output = v.GetString(FlagOutput)
			if cfg.Output == "" {
				cfg.Output = "json"
			}
			if cfg.Output != "json" && cfg.Output != "table" {
				return Exit(2, fmt.Errorf("unsupported output format %q", cfg.Output))
			}
		} else {
			cfg, err = Resolve(v)
		}
		if err != nil {
			return Exit(2, err)
		}
		if cmd.Annotations[OutputAnnotation] == "binary" && cfg.Output == "table" {
			return Exit(2, fmt.Errorf("%s preserves binary output and does not support table output", cmd.CommandPath()))
		}
		if err := validateTableShaping(cmd, cfg.Output); err != nil {
			return err
		}
		style, err := ResolveJSONStyle(v)
		if err != nil {
			return Exit(2, err)
		}
		jqExpression, err := cmd.Flags().GetString(FlagOutJQ)
		if err != nil {
			return Exit(2, err)
		}
		fieldExpression, err := cmd.Flags().GetString(FlagOutFields)
		if err != nil {
			return Exit(2, err)
		}
		rawOutput, err := cmd.Flags().GetBool(FlagRawOutput)
		if err != nil {
			return Exit(2, err)
		}
		printer := Printer{Style: style, Context: cmd.Context(), RawOutput: rawOutput}
		rootFlags := cmd.Root().PersistentFlags()
		jqSet := rootFlags.Changed(FlagOutJQ)
		fieldsSet := rootFlags.Changed(FlagOutFields)
		if rawOutput && (!jqSet || strings.TrimSpace(jqExpression) == "") {
			return Exit(2, fmt.Errorf("--raw-output requires --out-jq"))
		}
		if jqSet {
			printer.JQ, err = CompileJQ(jqExpression)
			if err != nil {
				return err
			}
		}
		if fieldsSet {
			printer.Fields, err = ParseFields(fieldExpression)
			if err != nil {
				return err
			}
			if operationID := cmd.Annotations["flexera.operationId"]; operationID != "" {
				entry, found, catalogErr := catalog.Lookup(operationID)
				if catalogErr != nil {
					return Exit(2, fmt.Errorf("cannot determine response envelope for --out-fields: %w", catalogErr))
				}
				if found {
					printer.Envelope = entry.ResponseEnvelope
				}
			}
		}

		base := opts.BaseHTTP
		debug, _ := cmd.Flags().GetBool(FlagDebug)
		doer := base
		if debug && base != nil && !offline {
			doer = NewSafeDebugDoer(base, cmd.ErrOrStderr())
			fmt.Fprintln(cmd.ErrOrStderr(), use+": safe debug HTTP tracing enabled (no headers, queries or bodies)")
		}

		deps.Config = cfg
		deps.OrgIDPresent = v.IsSet(FlagOrgID)
		deps.HTTP = doer
		deps.Stdout = cmd.OutOrStdout()
		deps.Stderr = cmd.ErrOrStderr()
		deps.Printer = printer
		if opts.Getenv != nil {
			deps.Getenv = opts.Getenv
		} else {
			deps.Getenv = os.Getenv
		}

		cmd.SetContext(WithDeps(cmd.Context(), deps))
		return nil
	}

	return root, deps
}

// requireOutput is a tiny helper so curated commands can resolve the writer
// pair without re-reading cobra streams.
func (d *Deps) Out() io.Writer { return d.Stdout }
func (d *Deps) Err() io.Writer { return d.Stderr }
