// Package userorgs provides `iam user-memberships orgs`: a wrapper over
// GET /iam/v1/users/{id}/orgs (Iam_User_Memberships_index) that auto-detects
// the caller's user ID from the access-token JWT when --id is omitted.
package userorgs

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	clipkg "github.com/flexera-public/flexera-cli/internal/cli"
	flexera "github.com/flexera-public/unified-go-client"
)

// Attach adds `orgs` beneath the generated `iam user-memberships` command.
func Attach(root *cobra.Command) error {
	return clipkg.AttachCommands(root, []string{"iam", "user-memberships"}, newOrgsCmd())
}

func newOrgsCmd() *cobra.Command {
	var userID int
	c := &cobra.Command{
		Use:         "orgs",
		Short:       "List organizations the authenticated user (or --id) can access",
		Example:     "flexera-cli iam user-memberships orgs --refresh-token <token>   # discover orgs your credentials can access",
		Annotations: map[string]string{"flexera.operationId": "Iam_User_Memberships_index", "flexera.readOnly": "true", clipkg.OutputAnnotation: "structured"},
		Args:        cobra.NoArgs,
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
