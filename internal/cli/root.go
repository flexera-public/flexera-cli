package cli

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"

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
		cfgFile, _ := cmd.Flags().GetString(FlagConfig)
		v, err := NewViper(cmd.Root().PersistentFlags(), cfgFile)
		if err != nil {
			return err
		}
		cfg, err := Resolve(v)
		if err != nil {
			return err
		}

		base := opts.BaseHTTP
		debug, _ := cmd.Flags().GetBool(FlagDebug)
		doer := base
		if debug && base != nil {
			doer = flexera.NewDebugDoer(base, cmd.ErrOrStderr(), flexera.DebugOptions{
				Label:       use,
				IncludeBody: true,
			})
			fmt.Fprintln(cmd.ErrOrStderr(), use+": debug HTTP tracing enabled (Authorization redacted)")
		}

		deps.Config = cfg
		deps.HTTP = doer
		deps.Stdout = cmd.OutOrStdout()
		deps.Stderr = cmd.ErrOrStderr()
		deps.Printer = Printer{}
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
