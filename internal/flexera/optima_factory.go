package flexera

import (
	"fmt"
	"strings"

	cliconfig "github.com/flexera-public/flexera-cli/internal/config"
	flexera "github.com/flexera-public/unified-go-client"
)

// EnvOptimaBaseURL overrides the auto-derived Optima base URL.
const EnvOptimaBaseURL = "FLEXERA_CLI_OPTIMA_BASE_URL"

// resolveOptimaBaseURL picks the effective Optima host.
// Precedence: explicit flag > env var > zone default (via the unified pkg).
func resolveOptimaBaseURL(cfg cliconfig.CommonConfig, getenv func(string) string, flagValue string) (string, error) {
	if v := strings.TrimSpace(flagValue); v != "" {
		return v, nil
	}
	if getenv != nil {
		if v := strings.TrimSpace(getenv(EnvOptimaBaseURL)); v != "" {
			return v, nil
		}
	}
	if base := flexera.OptimaBaseURL(cfg.Zone); base != "" {
		return base, nil
	}
	return "", fmt.Errorf("could not determine Optima base URL for zone %q (set --optima-base-url or %s)", cfg.Zone, EnvOptimaBaseURL)
}

// NewOptimaClient constructs a root client whose Optima operations route to
// the resolved Optima host and share the CLI's authenticated HTTP client.
func (f Factory) NewOptimaClient(cfg cliconfig.CommonConfig, getenv func(string) string, baseURLOverride string) (*flexera.ClientWithResponses, error) {
	baseURL, err := resolveOptimaBaseURL(cfg, getenv, baseURLOverride)
	if err != nil {
		return nil, err
	}
	if !cfg.HasAccessToken() {
		if err := cfg.ValidateAPIAuth(); err != nil {
			return nil, err
		}
	}

	auth := f.clientAuth(cfg)
	auth.OptimaBaseURL = baseURL
	if !cfg.HasAccessToken() {
		src, err := f.tokenSourceFor(cfg)
		if err != nil {
			return nil, err
		}
		auth.SharedTokenSource = src
	}
	return flexera.NewClientWithResponsesForAuth(auth)
}
