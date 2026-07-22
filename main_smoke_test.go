package main

import (
	"bytes"
	"context"
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
	for _, want := range []string{"budget", "policy", "finops", "user-orgs"} {
		if !strings.Contains(out, want) {
			t.Errorf("--help output missing %q", want)
		}
	}
}

// TestUnknownCommand verifies a non-existent command is a non-zero exit.
func TestUnknownCommand(t *testing.T) {
	var stdout, stderr bytes.Buffer
	exit := run(context.Background(), []string{"definitely-not-a-command"}, &stdout, &stderr, noEnv, &http.Client{})
	if exit == 0 {
		t.Fatal("expected non-zero exit for unknown command")
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
