// Package workflows provides opinionated, multi-step workflows that compose
// several API calls (e.g. `bill-analysis anomalies investigate`). They are
// hand-maintained and attached beneath the generated service commands.
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

// Attach adds the multi-step workflows beneath their generated service
// commands.
func Attach(root *cobra.Command) error {
	return clipkg.AttachCommands(root, []string{"bill-analysis", "anomalies"}, newAnomalyInvestigationCmd())
}

func newAnomalyInvestigationCmd() *cobra.Command {
	var optimaBaseURL, input, file string
	c := &cobra.Command{
		Use:     "investigate",
		Short:   "AI-powered cost anomaly investigation (multi-step workflow)",
		Example: "flexera-cli bill-analysis anomalies investigate --org-id 123 --input '{\"granularity\":\"day\",\"recency\":\"P7D\",\"increase_only\":true}'",
		Args:    cobra.NoArgs,
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

			client, err := deps.Factory().NewOptimaClient(cfg, deps.Getenv, optimaBaseURL)
			if err != nil {
				return err
			}
			resolver := flexera.NewBillingCenterResolver(client)

			var dbg anomaly.DebugLogger
			if debug, _ := cmd.Flags().GetBool(clipkg.FlagDebug); debug {
				dbg = func(_ context.Context, format string, args ...any) {
					fmt.Fprintf(deps.Stderr, "[anomaly] "+strings.TrimRight(format, "\n")+"\n", args...)
				}
			}

			inv := anomaly.New(client, resolver, dbg)
			out, err := inv.Invoke(ctx, in)
			if err != nil {
				return fmt.Errorf("anomaly investigation: %w", err)
			}
			return deps.Printer.Render(deps.Stdout, deps.Config.Output, out)
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
