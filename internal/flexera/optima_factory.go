package flexera

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	cliconfig "github.com/flexera-public/flexera-cli/internal/config"
	flexera "github.com/flexera-public/unified-go-client"
	"github.com/flexera-public/unified-go-client/service/optima"
)

// EnvOptimaBaseURL overrides the auto-derived Optima base URL.
const EnvOptimaBaseURL = "FLEXERA_CLI_OPTIMA_BASE_URL"

// OptimaClients is a type alias for the canonical bundle in the optima
// sub-package, kept so existing CLI callers don't need import changes.
type OptimaClients = optima.Clients

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

// NewOptimaClients constructs the three Optima-host clients sharing one
// authenticated HTTP roundtripper. When cfg carries OAuth credentials,
// the per-process tokenSourceFor cache is consulted so the unified API
// client and the Optima clients share a single OAuth2Doer; long-running
// pagination walks survive mid-walk token expiry transparently.
func (f Factory) NewOptimaClients(ctx context.Context, cfg cliconfig.CommonConfig, getenv func(string) string, baseURLOverride string) (*OptimaClients, error) {
	baseURL, err := resolveOptimaBaseURL(cfg, getenv, baseURLOverride)
	if err != nil {
		return nil, err
	}
	_ = ctx // optima.Config.TokenSource receives the per-request ctx; this top-level ctx is unused

	var httpDoer optima.HttpRequestDoer
	if d, ok := f.HTTPClient.(optima.HttpRequestDoer); ok && d != nil {
		httpDoer = d
	} else {
		httpDoer = &http.Client{}
	}

	optCfg := optima.Config{
		BaseURL:    baseURL,
		HTTPClient: httpDoer,
	}
	if cfg.HasAccessToken() {
		optCfg.AccessToken = cfg.AccessToken
	} else {
		src, err := f.tokenSourceFor(cfg)
		if err != nil {
			return nil, err
		}
		// src.Token is single-flight + lazy-refresh; satisfies the
		// Optima auth editor's per-request token need with shared
		// refresh against the unified API client's doer.
		optCfg.TokenSource = src.Token
	}
	return optima.NewClients(optCfg)
}
