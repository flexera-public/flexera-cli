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
	"github.com/spf13/cobra"
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

func TestAppliedPolicyListPaginatesByDefaultWithLimit(t *testing.T) {
	var calls atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		call := calls.Add(1)
		if r.URL.Path != "/policy/v1/orgs/123/projects/456/applied-policies" {
			t.Errorf("unexpected path %s", r.URL.Path)
		}
		if got := r.URL.Query().Get("limit"); got != "1" {
			t.Errorf("limit=%q", got)
		}
		if got := r.URL.Query().Get("orderBy"); got != "createdAt" {
			t.Errorf("orderBy=%q", got)
		}
		if got := r.URL.Query().Get("filter"); got != `name co "child"` {
			t.Errorf("filter=%q", got)
		}
		if call == 2 && r.URL.Query().Get("skipToken") != "second" {
			t.Errorf("page 2 skip token=%q", r.URL.Query().Get("skipToken"))
		}
		w.Header().Set("Content-Type", "application/json")
		if call == 1 {
			_, _ = w.Write([]byte(`{"kind":"policy#applied_policy_list","values":[{"id":"one"}],"count":1,"nextPage":"https://example.invalid/next?skipToken=second"}`))
		} else {
			_, _ = w.Write([]byte(`{"kind":"policy#applied_policy_list","values":[{"id":"two"}],"count":1}`))
		}
	}))
	defer srv.Close()
	deps := depsForServer(t, srv)
	var out strings.Builder
	deps.Stdout = &out
	cmd := newAppliedPolicyListCmd()
	cmd.SetContext(clipkg.WithDeps(context.Background(), deps))
	cmd.SetArgs([]string{"--project-id", "456", "--limit", "1", "--order-by", "createdAt", "--filter", `name co "child"`})
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}
	if got := calls.Load(); got != 2 {
		t.Fatalf("applied-policy list made %d requests; want first and nextPage", got)
	}
	if !strings.Contains(out.String(), `"one"`) || !strings.Contains(out.String(), `"two"`) || !strings.Contains(out.String(), `"count":2`) {
		t.Fatalf("missing merged page results %s", out.String())
	}
}

func TestAppliedPolicyListNoPaginateStopsAfterOnePage(t *testing.T) {
	var calls atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"kind":"policy#applied_policy_list","values":[{"id":"one"}],"nextPage":"https://example.invalid/next?skipToken=second"}`))
	}))
	defer srv.Close()
	deps := depsForServer(t, srv)
	var out strings.Builder
	deps.Stdout = &out
	cmd := newAppliedPolicyListCmd()
	cmd.SetContext(clipkg.WithDeps(context.Background(), deps))
	cmd.SetArgs([]string{"--project-id", "456", "--limit", "1", "--no-paginate"})
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}
	if calls.Load() != 1 || !strings.Contains(out.String(), `"one"`) {
		t.Fatalf("no-paginate requests=%d output=%s", calls.Load(), out.String())
	}
}

func TestAppliedPolicyListPaginatesTable(t *testing.T) {
	var calls atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		call := calls.Add(1)
		w.Header().Set("Content-Type", "application/json")
		if call == 1 {
			_, _ = w.Write([]byte(`{"kind":"policy:applied-policy-list","values":[{"id":"one","name":"First"}],"nextPage":"https://example.invalid/next?skipToken=second"}`))
		} else {
			_, _ = w.Write([]byte(`{"kind":"policy:applied-policy-list","values":[{"id":"two","name":"Second"}]}`))
		}
	}))
	defer srv.Close()
	deps := depsForServer(t, srv)
	deps.Config.Output = "table"
	var out strings.Builder
	deps.Stdout = &out
	cmd := newAppliedPolicyListCmd()
	cmd.SetContext(clipkg.WithDeps(context.Background(), deps))
	cmd.SetArgs([]string{"--project-id", "456"})
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}
	for _, value := range []string{"First", "Second"} {
		if !strings.Contains(out.String(), value) {
			t.Errorf("missing paginated table row %q: %s", value, out.String())
		}
	}
	if calls.Load() != 2 {
		t.Fatalf("requests=%d", calls.Load())
	}
}

func TestOtherCuratedPolicyListsPaginateByDefault(t *testing.T) {
	tests := []struct {
		name, path string
		newCommand func() *cobra.Command
	}{
		{"action-status", "/policy/v1/orgs/123/projects/456/action-statuses", newActionStatusListCmd},
		{"archived-incident", "/policy/v1/orgs/123/projects/456/archived-incidents", newArchivedIncidentListCmd},
		{"policy-template", "/policy/v1/orgs/123/projects/456/policy-templates", newPolicyTemplateListCmd},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for _, noPaginate := range []bool{false, true} {
				var calls atomic.Int32
				srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					call := calls.Add(1)
					if r.URL.Path != tc.path {
						t.Errorf("path=%s want %s", r.URL.Path, tc.path)
					}
					if call == 2 && r.URL.Query().Get("skipToken") != "second" {
						t.Errorf("skipToken=%q", r.URL.Query().Get("skipToken"))
					}
					w.Header().Set("Content-Type", "application/json")
					if call == 1 {
						_, _ = w.Write([]byte(`{"values":[{"id":"first"}],"nextPage":"https://example.invalid/next?skipToken=second"}`))
					} else {
						_, _ = w.Write([]byte(`{"values":[{"id":"second"}]}`))
					}
				}))
				deps := depsForServer(t, srv)
				var out strings.Builder
				deps.Stdout = &out
				cmd := tc.newCommand()
				cmd.SetContext(clipkg.WithDeps(context.Background(), deps))
				args := []string{"--project-id", "456"}
				if noPaginate {
					args = append(args, "--no-paginate")
				}
				cmd.SetArgs(args)
				if err := cmd.Execute(); err != nil {
					t.Fatal(err)
				}
				wantCalls := int32(2)
				if noPaginate {
					wantCalls = 1
				}
				if calls.Load() != wantCalls {
					t.Fatalf("calls=%d want=%d", calls.Load(), wantCalls)
				}
				if !strings.Contains(out.String(), "first") || !noPaginate && !strings.Contains(out.String(), "second") {
					t.Fatalf("missing values: %s", out.String())
				}
				srv.Close()
			}
		})
	}
}
