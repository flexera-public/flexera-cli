package config

import (
	"fmt"
	"strconv"
	"strings"

	flexera "github.com/flexera-public/unified-go-client"
)

const (
	EnvZone         = "FLEXERA_CLI_ZONE"
	EnvAPIBaseURL   = "FLEXERA_CLI_API_BASE_URL"
	EnvLoginBaseURL = "FLEXERA_CLI_LOGIN_BASE_URL"
	EnvOutput       = "FLEXERA_CLI_OUTPUT"
	EnvAccessToken  = "FLEXERA_CLI_ACCESS_TOKEN"
	EnvClientID     = "FLEXERA_CLI_CLIENT_ID"
	EnvClientSecret = "FLEXERA_CLI_CLIENT_SECRET"
	EnvRefreshToken = "FLEXERA_CLI_REFRESH_TOKEN"
	EnvOrgID        = "FLEXERA_CLI_ORG_ID"
)

type LookupEnv func(string) string

type CommonOptions struct {
	Zone         string
	APIBaseURL   string
	LoginBaseURL string
	Output       string
	AccessToken  string
	ClientID     string
	ClientSecret string
	RefreshToken string
	OrgID        int
}

type CommonConfig struct {
	Zone         flexera.Zone
	APIBaseURL   string
	LoginBaseURL string
	Output       string
	AccessToken  string
	ClientID     string
	ClientSecret string
	RefreshToken string
	OrgID        int
}

func ResolveCommon(opts CommonOptions, getenv LookupEnv) (CommonConfig, error) {
	if getenv == nil {
		getenv = func(string) string { return "" }
	}

	zone, err := parseZone(resolveString(opts.Zone, getenv(EnvZone), "nam"))
	if err != nil {
		return CommonConfig{}, err
	}

	output := strings.ToLower(resolveString(opts.Output, getenv(EnvOutput), "json"))
	if output != "json" && output != "table" {
		return CommonConfig{}, fmt.Errorf("unsupported output format %q; supported formats: json, table", output)
	}

	orgID, err := resolveInt(opts.OrgID, getenv(EnvOrgID))
	if err != nil {
		return CommonConfig{}, err
	}

	return CommonConfig{
		Zone:         zone,
		APIBaseURL:   strings.TrimSpace(resolveString(opts.APIBaseURL, getenv(EnvAPIBaseURL), "")),
		LoginBaseURL: strings.TrimSpace(resolveString(opts.LoginBaseURL, getenv(EnvLoginBaseURL), "")),
		Output:       output,
		AccessToken:  strings.TrimSpace(resolveString(opts.AccessToken, getenv(EnvAccessToken), "")),
		ClientID:     strings.TrimSpace(resolveString(opts.ClientID, getenv(EnvClientID), "")),
		ClientSecret: strings.TrimSpace(resolveString(opts.ClientSecret, getenv(EnvClientSecret), "")),
		RefreshToken: strings.TrimSpace(resolveString(opts.RefreshToken, getenv(EnvRefreshToken), "")),
		OrgID:        orgID,
	}, nil
}

func (c CommonConfig) RequireOrgID() error {
	if c.OrgID <= 0 {
		return fmt.Errorf("org ID is required; use --org-id or %s", EnvOrgID)
	}
	return nil
}

func (c CommonConfig) HasAccessToken() bool {
	return c.AccessToken != ""
}

func (c CommonConfig) HasClientCredentials() bool {
	return c.ClientID != "" && c.ClientSecret != ""
}

func (c CommonConfig) HasRefreshToken() bool {
	return c.RefreshToken != ""
}

func (c CommonConfig) HasOAuthCredentials() bool {
	return c.HasRefreshToken() || c.HasClientCredentials()

}

func (c CommonConfig) ValidateAPIAuth() error {
	if c.HasAccessToken() || c.HasOAuthCredentials() {
		return nil
	}

	if c.ClientID != "" || c.ClientSecret != "" {
		return fmt.Errorf("both client ID and client secret are required; use --client-id and --client-secret together")
	}

	return fmt.Errorf("authentication is required; provide --access-token (or %s), or --client-id and --client-secret", EnvAccessToken)
}

func parseZone(raw string) (flexera.Zone, error) {
	return flexera.ParseZone(raw)
}

func resolveString(flagValue, envValue, defaultValue string) string {
	if strings.TrimSpace(flagValue) != "" {
		return flagValue
	}
	if strings.TrimSpace(envValue) != "" {
		return envValue
	}
	return defaultValue
}

func resolveInt(flagValue int, envValue string) (int, error) {
	if flagValue > 0 {
		return flagValue, nil
	}
	trimmed := strings.TrimSpace(envValue)
	if trimmed == "" {
		return 0, nil
	}
	parsed, err := strconv.Atoi(trimmed)
	if err != nil {
		return 0, fmt.Errorf("invalid org ID %q in %s", envValue, EnvOrgID)
	}
	if parsed <= 0 {
		return 0, fmt.Errorf("invalid org ID %q in %s", envValue, EnvOrgID)
	}
	return parsed, nil
}
