// Package userorgs provides the hand-written "user-orgs" command: a thin,
// opinionated wrapper over the IAM endpoint GET /iam/v1/users/{id}/orgs that
// auto-detects the caller's user ID from the access-token JWT when --id is
// omitted. This is the preferred replacement for the legacy GRS user-orgs
// command (GRS is deprecated; IAM endpoints are preferred where available).
package userorgs

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	clipkg "github.com/flexera-public/flexera-cli/internal/cli"
	flexera "github.com/flexera-public/unified-go-client"
)

// NewCmd builds the "user-orgs" command tree.
func NewCmd() *cobra.Command {
	c := &cobra.Command{
		Use:     "user-orgs",
		Short:   "List organizations a user can access (IAM)",
		Example: "flexera-cli user-orgs list --refresh-token <token>   # discover orgs your credentials can access",
	}
	c.AddCommand(newListCmd())
	return c
}

func newListCmd() *cobra.Command {
	var userID int
	c := &cobra.Command{
		Use:   "list",
		Short: "List organizations the authenticated user (or --id) can access",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			deps := clipkg.DepsFrom(cmd.Context())
			cfg := deps.Config
			ctx := cmd.Context()

			// Resolve the user identity. With a static access token we parse
			// the JWT directly; with refresh-token / client-credentials we
			// first mint a token (and reject service-account tokens).
			accessToken := strings.TrimSpace(cfg.AccessToken)
			if userID == 0 {
				if accessToken == "" {
					if err := cfg.ValidateAPIAuth(); err != nil {
						return err
					}
					var tok *flexera.AuthTokenResponseBody
					var tokErr error
					if cfg.HasRefreshToken() {
						tok, tokErr = deps.Factory().TokenWithRefreshToken(ctx, cfg)
					} else {
						tok, tokErr = deps.Factory().TokenWithClientCredentials(ctx, cfg)
					}
					if tokErr != nil {
						return fmt.Errorf("failed to mint access token to identify user: %w", tokErr)
					}
					accessToken = tok.AccessToken
				}
				uid, err := flexera.UserIDFromToken(accessToken)
				if err != nil {
					return err
				}
				userID = uid
			}

			// Build a client; if we minted a token above, use it so the
			// request matches the user we just identified.
			clientCfg := cfg
			if accessToken != "" {
				clientCfg.AccessToken = accessToken
			}
			client, err := deps.Factory().NewAPIClient(clientCfg)
			if err != nil {
				return err
			}
			resp, err := client.IamUserMembershipsIndexWithResponse(ctx, userID)
			if err != nil {
				return err
			}
			if resp.JSON200 == nil {
				return flexera.ResponseError(resp.StatusCode(), resp.Body)
			}
			return deps.Printer.Render(deps.Stdout, deps.Config.Output, resp.JSON200)
		},
	}
	c.Flags().IntVar(&userID, "id", 0, "Flexera user ID (auto-detected from the access-token JWT when omitted)")
	return c
}
