// Package auth provides the hand-written "auth" command tree: OAuth token
// acquisition flows that are not modeled in the unified OpenAPI spec.
package auth

import (
	"github.com/spf13/cobra"

	clipkg "github.com/flexera-public/flexera-cli/internal/cli"
)

// NewCmd builds the "auth" command tree.
func NewCmd() *cobra.Command {
	c := &cobra.Command{
		Use:     "auth",
		Short:   "Acquire OAuth tokens from Flexera login",
		Example: "flexera-cli auth token client-credentials --client-id <id> --client-secret <secret>",
	}
	token := &cobra.Command{
		Use:   "token",
		Short: "Mint an access token",
	}
	token.AddCommand(newClientCredentialsCmd(), newRefreshCmd())
	c.AddCommand(token)
	return c
}

func newClientCredentialsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "client-credentials",
		Short: "Exchange --client-id/--client-secret for an access token",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			deps := clipkg.DepsFrom(cmd.Context())
			value, err := deps.Factory().TokenWithClientCredentials(cmd.Context(), deps.Config)
			if err != nil {
				return err
			}
			return deps.Printer.Render(deps.Stdout, deps.Config.Output, value)
		},
	}
}

func newRefreshCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "refresh",
		Short: "Exchange a --refresh-token for an access token",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			deps := clipkg.DepsFrom(cmd.Context())
			value, err := deps.Factory().TokenWithRefreshToken(cmd.Context(), deps.Config)
			if err != nil {
				return err
			}
			return deps.Printer.Render(deps.Stdout, deps.Config.Output, value)
		},
	}
}
