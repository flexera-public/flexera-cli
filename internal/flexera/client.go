package flexera

import (
	"context"
	"fmt"
	"strings"
	"sync"

	cliconfig "github.com/flexera-public/flexera-cli/internal/config"
	flexera "github.com/flexera-public/unified-go-client"
)

type Factory struct {
	HTTPClient flexera.HttpRequestDoer
}

// tokenSources caches *flexera.SharedTokenSource per credential tuple
// for the lifetime of the CLI process. The cache lets concurrent client
// constructions (unified API + Optima within one command invocation)
// share a single OAuth2Doer and avoid duplicate /oidc/token round
// trips. Static-token configs do not participate.
var (
	tokenSourcesMu sync.Mutex
	tokenSources   = map[tokenSourceKey]*flexera.SharedTokenSource{}
)

// tokenSourceKey discriminates cached sources. ClientSecret and
// RefreshToken are both included so two configs sharing
// (loginBaseURL, ClientID) but differing in either secret cannot
// cross-pollute the cached doer (e.g. mid-rotation, integration test
// fixtures, alternating client-credentials vs refresh-token grants).
type tokenSourceKey struct {
	loginBaseURL string
	clientID     string
	clientSecret string
	refreshToken string
}

// resetTokenSourcesForTest drops the package-level cache. Test-only.
func resetTokenSourcesForTest() {
	tokenSourcesMu.Lock()
	defer tokenSourcesMu.Unlock()
	tokenSources = map[tokenSourceKey]*flexera.SharedTokenSource{}
}

func (f Factory) NewAuthHelper(cfg cliconfig.CommonConfig) (*flexera.AuthHelper, error) {
	return flexera.NewAuthHelper(flexera.AuthHelperConfig{
		Zone:         cfg.Zone,
		HTTPClient:   f.HTTPClient,
		APIBaseURL:   cfg.APIBaseURL,
		LoginBaseURL: cfg.LoginBaseURL,
	})
}

// clientAuth maps the CLI's resolved config to the unified-go-client ClientAuth used by
// flexera.NewClientWithResponsesForAuth.
func (f Factory) clientAuth(cfg cliconfig.CommonConfig) flexera.ClientAuth {
	return flexera.ClientAuth{
		Zone:         cfg.Zone,
		HTTPClient:   f.HTTPClient,
		APIBaseURL:   cfg.APIBaseURL,
		LoginBaseURL: cfg.LoginBaseURL,
		AccessToken:  cfg.AccessToken,
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		RefreshToken: cfg.RefreshToken,
	}
}

// tokenSourceFor returns the cached SharedTokenSource for cfg's
// credentials, constructing one (and caching it) on first use. Returns
// nil when cfg carries a static access token (no OAuth flow needed) or
// when validation fails.
func (f Factory) tokenSourceFor(cfg cliconfig.CommonConfig) (*flexera.SharedTokenSource, error) {
	if cfg.HasAccessToken() {
		return nil, nil
	}
	if !cfg.HasOAuthCredentials() {
		return nil, cfg.ValidateAPIAuth()
	}
	helper, err := f.NewAuthHelper(cfg)
	if err != nil {
		return nil, err
	}
	key := tokenSourceKey{
		loginBaseURL: helper.LoginBaseURL(),
		clientID:     cfg.ClientID,
		clientSecret: cfg.ClientSecret,
		refreshToken: cfg.RefreshToken,
	}
	tokenSourcesMu.Lock()
	defer tokenSourcesMu.Unlock()
	if src, ok := tokenSources[key]; ok {
		return src, nil
	}
	doer, err := helper.NewOAuth2Doer(flexera.OAuthConfig{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		RefreshToken: cfg.RefreshToken,
	})
	if err != nil {
		return nil, err
	}
	src := flexera.NewSharedTokenSource(doer)
	tokenSources[key] = src
	return src, nil
}

func (f Factory) NewAPIClient(cfg cliconfig.CommonConfig) (*flexera.ClientWithResponses, error) {
	return f.NewAPIClientForBaseURL(cfg, "")
}

// NewAPIClientForBaseURL builds an authed client optionally pointed at a
// non-gateway host (e.g. the GRS grs-front host). An empty baseURL uses the
// zone's default gateway. When cfg carries OAuth credentials the
// per-process tokenSourceFor cache is consulted so multiple sibling
// clients share one OAuth doer.
func (f Factory) NewAPIClientForBaseURL(cfg cliconfig.CommonConfig, baseURL string) (*flexera.ClientWithResponses, error) {
	if !cfg.HasAccessToken() {
		if err := cfg.ValidateAPIAuth(); err != nil {
			return nil, err
		}
	}
	auth := f.clientAuth(cfg)
	if !cfg.HasAccessToken() {
		src, err := f.tokenSourceFor(cfg)
		if err != nil {
			return nil, err
		}
		auth.SharedTokenSource = src
	}
	if strings.TrimSpace(baseURL) != "" {
		return flexera.NewClientWithResponsesForAuth(auth, flexera.WithBaseURL(baseURL))
	}
	return flexera.NewClientWithResponsesForAuth(auth)
}

func (f Factory) TokenWithClientCredentials(ctx context.Context, cfg cliconfig.CommonConfig) (*flexera.AuthTokenResponseBody, error) {
	helper, err := f.NewAuthHelper(cfg)
	if err != nil {
		return nil, err
	}
	if strings.TrimSpace(cfg.ClientID) == "" || strings.TrimSpace(cfg.ClientSecret) == "" {
		return nil, fmt.Errorf("client credentials are required; use --client-id and --client-secret")
	}
	return helper.TokenWithClientCredentials(ctx, cfg.ClientID, cfg.ClientSecret)
}

func (f Factory) TokenWithRefreshToken(ctx context.Context, cfg cliconfig.CommonConfig) (*flexera.AuthTokenResponseBody, error) {
	helper, err := f.NewAuthHelper(cfg)
	if err != nil {
		return nil, err
	}
	if strings.TrimSpace(cfg.RefreshToken) == "" {
		return nil, fmt.Errorf("refresh token is required; use --refresh-token")
	}
	return helper.TokenWithRefreshToken(ctx, cfg.RefreshToken)
}
