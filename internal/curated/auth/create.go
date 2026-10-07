package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	clipkg "github.com/flexera-public/flexera-cli/internal/cli"
	flexera "github.com/flexera-public/unified-go-client"
	"github.com/spf13/cobra"
)

// These are request-body fields, never inherited authentication config.
var tokenBodyFlags = []struct{ flag, key, description string }{
	{"body-client-id", "client_id", "Service app client ID (required for client_credentials)"},
	{"body-client-secret", "client_secret", "Service app client secret (required for client_credentials)"},
	{"code", "code", "Authorization code (required for authorization_code)"},
	{"grant-type", "grant_type", "Required: authorization_code, client_credentials, or refresh_token"},
	{"redirect-uri", "redirect_uri", "Redirect URI (required for authorization_code)"},
	{"body-refresh-token", "refresh_token", "Refresh token from a previous offline_access authorization code exchange"},
}

func newCreateCmd() *cobra.Command {
	var bodyRaw string
	var dryRun bool
	values := make([]string, len(tokenBodyFlags))
	c := &cobra.Command{
		Use:         "create",
		Short:       "Generate access token",
		Long:        "Generate a Flexera One access token using POST /oidc/token. JSON input is encoded as application/x-www-form-urlencoded. No prior authentication is required; inherited authentication flags are not request-body fields. Grant-specific field requirements are enforced by the login service.",
		Example:     "flexera-cli auth token create --body @request.json --dry-run",
		Args:        cobra.NoArgs,
		Annotations: map[string]string{"flexera.operationId": "Auth_Token_token"},
		RunE: func(cmd *cobra.Command, _ []string) error {
			fields := map[string]any{}
			for i, f := range tokenBodyFlags {
				if cmd.Flags().Changed(f.flag) {
					fields[f.key] = values[i]
				}
			}
			raw, err := clipkg.ResolveBody(bodyRaw, fields, cmd.InOrStdin())
			if err != nil {
				return fmt.Errorf("cannot read token request body; provide a JSON object via --body or body field flags")
			}
			body, err := decodeTokenBody(raw)
			if err != nil {
				return err
			}
			deps := clipkg.DepsFrom(cmd.Context())
			if deps == nil {
				return fmt.Errorf("token creation requires CLI dependencies")
			}
			if dryRun {
				// Do not claim catalog validation: the form adapter currently
				// checks only JSON shape, supported fields, and the grant enum.
				preview := map[string]any{"grant_type": body.GrantType}
				for _, f := range tokenBodyFlags {
					if f.key != "grant_type" && fieldsPresent(raw, f.key) {
						preview[f.key] = "[REDACTED]"
					}
				}
				return clipkg.WriteJSON(deps.Stdout, map[string]any{
					"dryRun": true, "destructive": false, "redacted": true,
					"validation": map[string]string{"status": "unsupported"},
					"plan":       map[string]any{"method": "POST /oidc/token", "contentType": "application/x-www-form-urlencoded", "body": preview},
				}, deps.Printer.Style, nil)
			}
			// The shared SDK debug doer dumps unredacted form/token bodies.
			// Reject tracing rather than disclose credentials on stderr.
			if debug, _ := cmd.Flags().GetBool(clipkg.FlagDebug); debug {
				return fmt.Errorf("--debug is not supported for token creation because token bodies contain secrets")
			}
			helper, err := deps.Factory().NewAuthHelper(deps.Config)
			if err != nil {
				return fmt.Errorf("cannot resolve login endpoint")
			}
			client, err := flexera.NewClientWithResponses(helper.LoginBaseURL(), flexera.WithHTTPClient(deps.HTTP))
			if err != nil {
				return fmt.Errorf("cannot construct token client")
			}
			resp, err := client.AuthTokenTokenWithFormdataBodyWithResponse(cmd.Context(), body)
			if err != nil {
				// Transport/parser errors can contain credentials or token values.
				return fmt.Errorf("token creation failed while sending or decoding the response")
			}
			if resp.JSON200 == nil {
				// Auth_Error.message and arbitrary response bodies may echo secrets.
				return fmt.Errorf("token creation failed (HTTP %d)", resp.StatusCode())
			}
			return deps.Printer.Render(deps.Stdout, deps.Config.Output, resp.JSON200)
		},
	}
	for i, f := range tokenBodyFlags {
		c.Flags().StringVar(&values[i], f.flag, "", f.description+" (body)")
	}
	c.Flags().StringVar(&bodyRaw, "body", "", "JSON body (inline | @file | @-); overrides body field flags")
	c.Flags().BoolVar(&dryRun, "dry-run", false, "preview the request with sensitive values redacted; no HTTP request")
	return c
}

func decodeTokenBody(raw []byte) (flexera.AuthTokenTokenFormdataRequestBody, error) {
	var body flexera.AuthTokenTokenFormdataRequestBody
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(raw, &fields); err != nil || fields == nil {
		return body, fmt.Errorf("token request body must be a JSON object")
	}
	for key, value := range fields {
		known := false
		for _, f := range tokenBodyFlags {
			known = known || key == f.key
		}
		var text string
		if !known || bytes.Equal(bytes.TrimSpace(value), []byte("null")) || json.Unmarshal(value, &text) != nil {
			return body, fmt.Errorf("token request body must contain only supported string properties")
		}
	}
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&body); err != nil {
		return body, fmt.Errorf("invalid token request body")
	}
	if err := decoder.Decode(new(any)); err != io.EOF {
		return body, fmt.Errorf("token request body must contain exactly one JSON object")
	}
	switch body.GrantType {
	case flexera.AuthTokenRequestBodyGrantTypeAuthorizationCode,
		flexera.AuthTokenRequestBodyGrantTypeClientCredentials,
		flexera.AuthTokenRequestBodyGrantTypeRefreshToken:
		return body, nil
	default:
		return body, fmt.Errorf("grant_type is required; use --grant-type or grant_type in --body with authorization_code, client_credentials, or refresh_token")
	}
}

func fieldsPresent(raw []byte, key string) bool {
	var fields map[string]json.RawMessage
	_ = json.Unmarshal(raw, &fields)
	_, ok := fields[key]
	return ok
}
