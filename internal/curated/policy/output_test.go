package policy

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	clipkg "github.com/flexera-public/flexera-cli/internal/cli"
)

func TestAppliedPolicyLogPreservesMarkdown(t *testing.T) {
	for _, body := range []string{"# Evaluation\n\nSuccess\n", `{"looks":"like JSON"}`, ""} {
		t.Run(body, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "text/markdown")
				_, _ = w.Write([]byte(body))
			}))
			defer srv.Close()
			deps := depsForServer(t, srv)
			var out bytes.Buffer
			deps.Stdout = &out
			deps.Printer = clipkg.Printer{Style: clipkg.JSONStylePretty}
			cmd := newAppliedPolicyLogCmd()
			cmd.SetContext(clipkg.WithDeps(context.Background(), deps))
			cmd.SetArgs([]string{"--id", "policy-1", "--project-id", "42"})
			if err := cmd.Execute(); err != nil {
				t.Fatal(err)
			}
			want := body
			if !strings.HasSuffix(want, "\n") {
				want += "\n"
			}
			if out.String() != want {
				t.Fatalf("got %q, want %q", out.String(), want)
			}
		})
	}
}

func TestAppliedPolicyLogRejectsTableBeforeHTTP(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		t.Error("unsupported table performed an HTTP request")
	}))
	defer srv.Close()
	deps := depsForServer(t, srv)
	deps.Config.Output = "table"
	cmd := newAppliedPolicyLogCmd()
	cmd.SetContext(clipkg.WithDeps(context.Background(), deps))
	cmd.SetArgs([]string{"--id", "policy-1", "--project-id", "42"})
	if err := cmd.Execute(); err == nil || !strings.Contains(err.Error(), "table output is not supported") {
		t.Fatalf("got %v", err)
	}
}
