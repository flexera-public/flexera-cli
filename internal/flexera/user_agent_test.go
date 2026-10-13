package flexera

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	cliconfig "github.com/flexera-public/flexera-cli/internal/config"
	unified "github.com/flexera-public/unified-go-client"
)

type userAgentCaptureDoer struct {
	userAgent string
}

func (d *userAgentCaptureDoer) Do(req *http.Request) (*http.Response, error) {
	d.userAgent = req.Header.Get("User-Agent")
	return &http.Response{
		StatusCode: http.StatusNoContent,
		Header:     make(http.Header),
		Body:       io.NopCloser(strings.NewReader("")),
		Request:    req,
	}, nil
}

func TestFactoryAddsProductUserAgent(t *testing.T) {
	tests := []struct {
		name      string
		newClient func(Factory, cliconfig.CommonConfig) (*unified.ClientWithResponses, error)
		request   func(*unified.ClientWithResponses) error
	}{
		{
			name: "API",
			newClient: func(factory Factory, cfg cliconfig.CommonConfig) (*unified.ClientWithResponses, error) {
				return factory.NewAPIClient(cfg)
			},
			request: func(client *unified.ClientWithResponses) error {
				_, err := client.CredCredentialIndexProjectWithResponse(context.Background(), 1, nil)
				return err
			},
		},
		{
			name: "Optima",
			newClient: func(factory Factory, cfg cliconfig.CommonConfig) (*unified.ClientWithResponses, error) {
				return factory.NewOptimaClient(cfg, nil, "https://optima.example.test")
			},
			request: func(client *unified.ClientWithResponses) error {
				_, err := client.BillAnalysisAdjustmentDefinitionShowWithResponse(context.Background(), 1)
				return err
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			doer := &userAgentCaptureDoer{}
			factory := Factory{HTTPClient: doer, UserAgentVersion: "1.2.3"}
			cfg := cliconfig.CommonConfig{Zone: unified.ZoneNAM, AccessToken: "test-token"}
			client, err := tc.newClient(factory, cfg)
			if err != nil {
				t.Fatalf("create client: %v", err)
			}
			if err := tc.request(client); err != nil {
				t.Fatalf("make request: %v", err)
			}
			if !strings.HasPrefix(doer.userAgent, "flexera-cli/1.2.3 flexera-unified-go-client/") {
				t.Fatalf("unexpected User-Agent %q", doer.userAgent)
			}
		})
	}
}
