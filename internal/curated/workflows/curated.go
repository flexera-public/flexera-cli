// Package workflows provides the hand-written "curated" command tree:
// opinionated, multi-step workflows that compose several API calls. The CLI
// surface is intentionally small and hand-maintained (these are not
// spec-generated CRUD).
package workflows

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	clipkg "github.com/flexera-public/flexera-cli/internal/cli"
	flexera "github.com/flexera-public/unified-go-client"
	"github.com/flexera-public/unified-go-client/anomaly"
)

type tool struct {
	name        string
	description string
}

// tools is the curated catalogue surfaced by `curated list`.
var tools = []tool{
	{"anomaly-investigation", "AI-powered cost anomaly investigation across service/usage/region/compute dimensions"},
}

// NewCmd builds the "curated" command tree.
func NewCmd() *cobra.Command {
	c := &cobra.Command{
		Use:     "curated",
		Short:   "Run curated multi-step workflows (e.g. anomaly-investigation)",
		Example: "flexera-cli curated list",
	}
	c.AddCommand(newListCmd(), newAnomalyInvestigationCmd())
	return c
}

func newListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List available curated tools",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			out := cmd.OutOrStdout()
			fmt.Fprintln(out, "Available tools:")
			for _, t := range tools {
				fmt.Fprintf(out, "  %-28s %s\n", t.name, t.description)
			}
			return nil
		},
	}
}

func newAnomalyInvestigationCmd() *cobra.Command {
	var optimaBaseURL, input, file string
	c := &cobra.Command{
		Use:   "anomaly-investigation",
		Short: "AI-powered cost anomaly investigation",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			deps := clipkg.DepsFrom(cmd.Context())
			ctx := cmd.Context()

			raw, err := resolveInput(input, file, cmd.InOrStdin())
			if err != nil {
				return err
			}
			var in anomaly.Input
			if len(raw) > 0 && string(raw) != "null" {
				if err := json.Unmarshal(raw, &in); err != nil {
					return fmt.Errorf("invalid anomaly input: %w", err)
				}
			}

			cfg := deps.Config
			if in.OrgID == 0 {
				if err := cfg.RequireOrgID(); err != nil {
					return errors.New("org_id is required (set it in input JSON or pass --org-id)")
				}
				in.OrgID = cfg.OrgID
			}

			clients, err := deps.Factory().NewOptimaClients(ctx, cfg, deps.Getenv, optimaBaseURL)
			if err != nil {
				return err
			}
			resolver := flexera.NewBillingCenterResolver(clients.BillingCenterService)

			var dbg anomaly.DebugLogger
			if debug, _ := cmd.Flags().GetBool(clipkg.FlagDebug); debug {
				dbg = func(_ context.Context, format string, args ...any) {
					fmt.Fprintf(deps.Stderr, "[anomaly] "+strings.TrimRight(format, "\n")+"\n", args...)
				}
			}

			inv := anomaly.New(clients.BillAnalysis, resolver, dbg)
			out, err := inv.Invoke(ctx, in)
			if err != nil {
				return fmt.Errorf("anomaly-investigation: %w", err)
			}
			return deps.Printer.Render(deps.Stdout, "json", out)
		},
	}
	c.Flags().StringVar(&optimaBaseURL, "optima-base-url", "", "Override the Optima base URL")
	c.Flags().StringVar(&input, "input", "", "Inline JSON input (mutually exclusive with --file)")
	c.Flags().StringVar(&file, "file", "", "Path to JSON input, or - for stdin")
	return c
}

// resolveInput mirrors the legacy curated input resolution: --input (inline
// JSON) and --file (path or - for stdin) are mutually exclusive.
func resolveInput(inline, file string, stdin interface{ Read([]byte) (int, error) }) (json.RawMessage, error) {
	inline = strings.TrimSpace(inline)
	file = strings.TrimSpace(file)
	if inline != "" && file != "" {
		return nil, errors.New("--input and --file are mutually exclusive")
	}
	if inline != "" {
		if !json.Valid([]byte(inline)) {
			return nil, errors.New("--input is not valid JSON")
		}
		return json.RawMessage(inline), nil
	}
	if file != "" {
		arg := file
		if arg != "-" {
			arg = "@" + arg
		} else {
			arg = "@-"
		}
		return clipkg.ResolveBody(arg, nil, stdin)
	}
	return nil, nil
}
