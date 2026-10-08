package cli

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	cliconfig "github.com/flexera-public/flexera-cli/internal/config"
	cliflexera "github.com/flexera-public/flexera-cli/internal/flexera"
)

func TestAPIClientHonorsOptimaBaseURLEnv(t *testing.T) {
	var gotPath string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	deps := &Deps{
		HTTP:   server.Client(),
		Config: cliconfig.CommonConfig{AccessToken: "token", Zone: "com", OrgID: 7},
		Getenv: func(key string) string {
			if key == cliflexera.EnvOptimaBaseURL {
				return server.URL
			}
			return ""
		},
	}
	client, err := deps.APIClient()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := client.BillAnalysisCostsDimensionsWithResponse(context.Background(), 7); err != nil {
		t.Fatal(err)
	}
	if gotPath == "" {
		t.Fatal("Optima operation did not reach the overridden host")
	}
}
