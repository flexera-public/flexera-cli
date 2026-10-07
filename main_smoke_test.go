package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// noEnv is an empty environment lookup for tests.
func noEnv(string) string { return "" }

// TestRootHelp verifies the root command builds and lists commands.
func TestRootHelp(t *testing.T) {
	var stdout, stderr bytes.Buffer
	exit := run(context.Background(), []string{"--help"}, &stdout, &stderr, noEnv, &http.Client{})
	if exit != 0 {
		t.Fatalf("--help exit=%d stderr=%s", exit, stderr.String())
	}
	out := stdout.String() + stderr.String()
	for _, want := range []string{"budget", "policy", "finops", "user-orgs", "Find commands for a task: flexera-cli cli search"} {
		if !strings.Contains(out, want) {
			t.Errorf("--help output missing %q", want)
		}
	}
}

func TestOfflineSearchSmoke(t *testing.T) {
	var stdout, stderr bytes.Buffer
	exit := run(context.Background(), []string{"cli", "search", "delete budget", "--limit", "1", "--out-jq", ".[0].command", "-r"}, &stdout, &stderr, noEnv, &http.Client{})
	if exit != 0 || strings.TrimSpace(stdout.String()) != "flexera-cli budget delete" {
		t.Fatalf("offline search: %d %s %s", exit, stdout.String(), stderr.String())
	}
}

func TestBinaryDownloadPreservesBytes(t *testing.T) {
	body := []byte("id,name\n1,example\n\x00\xff")
	doer := smokeDoer(func(r *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 200, Header: http.Header{"Content-Type": []string{"text/plain; charset=utf-8"}}, Body: io.NopCloser(bytes.NewReader(body)), Request: r}, nil
	})
	var out, stderr bytes.Buffer
	args := []string{"bill-months", "download", "--download-token", "test-token", "--org-id", "123", "--access-token", "smoke-token", "--json-style", "pretty"}
	if code := run(context.Background(), args, &out, &stderr, noEnv, doer); code != 0 || !bytes.Equal(out.Bytes(), body) {
		t.Fatalf("download bytes changed: %d %q %s", code, out.Bytes(), stderr.String())
	}
}

type smokeDoer func(*http.Request) (*http.Response, error)

func (d smokeDoer) Do(r *http.Request) (*http.Response, error) { return d(r) }

// TestUnknownCommand verifies a non-existent command is a non-zero exit.
func TestUnknownCommand(t *testing.T) {
	var stdout, stderr bytes.Buffer
	exit := run(context.Background(), []string{"definitely-not-a-command"}, &stdout, &stderr, noEnv, &http.Client{})
	if exit != 2 {
		t.Fatalf("expected usage exit 2 for unknown command, got %d", exit)
	}
	var value struct {
		Error       string   `json:"error"`
		Suggestions []string `json:"suggestions"`
	}
	if err := json.Unmarshal(stderr.Bytes(), &value); err != nil || value.Error == "" || value.Suggestions == nil {
		t.Fatalf("expected JSON command error with suggestions, got %s", stderr.String())
	}
}

func TestRuleBasedDimensionCompatibilityAlias(t *testing.T) {
	var stdout, stderr bytes.Buffer
	exit := run(context.Background(), []string{"rule-base-dimension", "bulk", "--help"}, &stdout, &stderr, noEnv, &http.Client{})
	if exit != 0 {
		t.Fatalf("alias help exit=%d stderr=%s", exit, stderr.String())
	}
	if !strings.Contains(stdout.String(), "Create or update rule-based dimensions") {
		t.Fatalf("expected bulk help, got %s", stdout.String())
	}
}

// TestGeneratedCommandSmoke exercises a spec-generated command end-to-end
// against a mock server: flag parsing, authed client construction, the
// generated client call, and JSON rendering.
func TestGeneratedCommandSmoke(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Authorization"); got != "Bearer smoke-token" {
			t.Errorf("unexpected auth header: %q", got)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"kind":"finops:budget-list","values":[{"id":"b-1","name":"Smoke Budget"}]}`))
	}))
	defer server.Close()

	var stdout, stderr bytes.Buffer
	exit := run(context.Background(), []string{
		"budget", "list",
		"--api-base-url", server.URL,
		"--org-id", "123",
		"--access-token", "smoke-token",
	}, &stdout, &stderr, noEnv, server.Client())
	if exit != 0 {
		t.Fatalf("budget list exit=%d stderr=%s", exit, stderr.String())
	}
	if !strings.Contains(stdout.String(), "Smoke Budget") {
		t.Fatalf("expected budget JSON, got %s", stdout.String())
	}
}

func TestGeneratedTypedTableSmoke(t *testing.T) {
	doer := smokeDoer(func(r *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 200, Header: http.Header{"Content-Type": []string{"application/json"}}, Body: io.NopCloser(strings.NewReader(`{"values":[{"id":"b-1","name":"Table Budget"}]}`)), Request: r}, nil
	})
	var out, stderr bytes.Buffer
	if code := run(context.Background(), []string{"budget", "list", "--org-id", "123", "--access-token", "fixture-token", "-o", "table"}, &out, &stderr, noEnv, doer); code != 0 || !strings.Contains(out.String(), "Table Budget") {
		t.Fatalf("typed table broke: %d %s %s", code, out.String(), stderr.String())
	}
}

// TestCuratedCommandSmoke exercises a curated command (policy, which resolves
// the project) end-to-end against a mock server. Uses an explicit
// --project-id so no GRS project-resolution call is needed.
func TestCuratedCommandSmoke(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/policy/v1/orgs/123/projects/456/applied-policies" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"kind":"policy:applied-policy-list","values":[]}`))
	}))
	defer server.Close()

	var stdout, stderr bytes.Buffer
	exit := run(context.Background(), []string{
		"policy", "applied-policy", "list",
		"--project-id", "456",
		"--api-base-url", server.URL,
		"--org-id", "123",
		"--access-token", "smoke-token",
	}, &stdout, &stderr, noEnv, server.Client())
	if exit != 0 {
		t.Fatalf("policy applied-policy list exit=%d stderr=%s", exit, stderr.String())
	}
}

// TestGraphQLQuerySmoke exercises the curated graphql query command backed
// by flexera.(*Client).GraphQL, which POSTs to /explore/graphql -- a path
// with no generated OpenAPI operation.
func TestGraphQLQuerySmoke(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/explore/graphql" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("Authorization"); got == "" {
			t.Errorf("expected an Authorization header, got none")
		}
		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Errorf("decode request body: %v", err)
		}
		if body["query"] != "{ viewer { id } }" {
			t.Errorf("unexpected query in request body: %v", body["query"])
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"data":{"viewer":{"id":"42"}}}`))
	}))
	defer server.Close()

	var stdout, stderr bytes.Buffer
	exit := run(context.Background(), []string{
		"graphql", "query",
		"--body", `{"query":"{ viewer { id } }"}`,
		"--api-base-url", server.URL,
		"--org-id", "123",
		"--access-token", "smoke-token",
	}, &stdout, &stderr, noEnv, server.Client())
	if exit != 0 {
		t.Fatalf("graphql query exit=%d stderr=%s", exit, stderr.String())
	}
	if !strings.Contains(stdout.String(), `"id"`) {
		t.Fatalf("expected graphql response JSON, got %s", stdout.String())
	}
}
