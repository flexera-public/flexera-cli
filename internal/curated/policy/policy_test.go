package policy

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"

	clipkg "github.com/flexera-public/flexera-cli/internal/cli"
	cliconfig "github.com/flexera-public/flexera-cli/internal/config"
)

// fakeJWT builds an unsigned, JWT-shaped token (header.payload.signature)
// carrying the given `sub` claim, sufficient for flexera.UserIDFromToken to
// decode without verifying a signature.
func fakeJWT(t *testing.T, sub string) string {
	t.Helper()
	enc := base64.RawURLEncoding.EncodeToString
	header := enc([]byte(`{"alg":"none","typ":"JWT"}`))
	payload, err := json.Marshal(map[string]string{"sub": sub})
	if err != nil {
		t.Fatalf("marshal JWT payload: %v", err)
	}
	return header + "." + enc(payload) + ".sig"
}

// grsServer returns an httptest server that serves the GRS user-projects
// index for the test user (u-999, see fakeJWT) with the given JSON array
// body, and a counter of requests received.
func grsServer(t *testing.T, body string) (*httptest.Server, *int32) {
	t.Helper()
	var hits int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&hits, 1)
		if r.URL.Path != "/grs/users/999/projects" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(body))
	}))
	t.Cleanup(srv.Close)
	return srv, &hits
}

func depsForServer(t *testing.T, srv *httptest.Server) *clipkg.Deps {
	t.Helper()
	cfg, err := cliconfig.ResolveCommon(cliconfig.CommonOptions{
		Zone:        "nam",
		OrgID:       123,
		AccessToken: fakeJWT(t, "u-999"),
		APIBaseURL:  srv.URL,
	}, func(string) string { return "" })
	if err != nil {
		t.Fatalf("ResolveCommon: %v", err)
	}
	return &clipkg.Deps{Config: cfg, HTTP: srv.Client(), Getenv: func(string) string { return "" }}
}

func TestResolvePolicyProjectID_ExplicitNoHTTP(t *testing.T) {
	srv, hits := grsServer(t, `[]`)
	deps := depsForServer(t, srv)
	got, err := resolvePolicyProjectID(context.Background(), deps, 777)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 777 {
		t.Fatalf("expected explicit project id 777, got %d", got)
	}
	if n := atomic.LoadInt32(hits); n != 0 {
		t.Fatalf("expected no HTTP calls for explicit --project-id, got %d", n)
	}
}

func TestResolvePolicyProjectID_SingleAutoResolves(t *testing.T) {
	srv, _ := grsServer(t, `[{"id":42,"name":"Solo"}]`)
	deps := depsForServer(t, srv)
	got, err := resolvePolicyProjectID(context.Background(), deps, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 42 {
		t.Fatalf("expected auto-resolved project id 42, got %d", got)
	}
}

func TestResolvePolicyProjectID_MultipleErrors(t *testing.T) {
	srv, _ := grsServer(t, `[{"id":1,"name":"A"},{"id":2,"name":"B"}]`)
	deps := depsForServer(t, srv)
	_, err := resolvePolicyProjectID(context.Background(), deps, 0)
	if err == nil {
		t.Fatal("expected error for multiple projects")
	}
	if !strings.Contains(err.Error(), "--project-id") {
		t.Fatalf("expected error to mention --project-id, got: %v", err)
	}
}

func TestResolvePolicyProjectID_ZeroErrors(t *testing.T) {
	srv, _ := grsServer(t, `[]`)
	deps := depsForServer(t, srv)
	_, err := resolvePolicyProjectID(context.Background(), deps, 0)
	if err == nil {
		t.Fatal("expected error for zero projects")
	}
	if !strings.Contains(err.Error(), "no projects") {
		t.Fatalf("expected 'no projects' error, got: %v", err)
	}
}
