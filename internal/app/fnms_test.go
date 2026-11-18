package app_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/flexera-public/flexera-cli/internal/app"
)

func TestFNMSRootContracts(t *testing.T) {
	payload := []byte{'P', 'K', 0, 255, '\n'}
	external := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "" {
			t.Error("download leaked auth")
		}
		w.Write(payload)
	}))
	defer external.Close()
	requests := 0
	api := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests++
		if r.Header.Get("Authorization") != "Bearer offline-token" {
			t.Error("API missing auth")
		}
		if strings.Contains(r.URL.Path, "/files/") {
			w.Header().Set("Location", external.URL)
			w.WriteHeader(303)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(202)
		fmt.Fprint(w, `{"jobId":"job","status":"Created"}`)
	}))
	defer api.Close()
	execute := func(args ...string) (string, error) {
		var out, stderr bytes.Buffer
		root, _ := app.NewRootCmd(&out, &stderr, func(string) string { return "" }, api.Client(), "test")
		root.SetArgs(append(args, "--config", os.DevNull, "--access-token", "offline-token", "--org-id", "7", "--api-base-url", api.URL))
		err := root.Execute()
		return out.String(), err
	}
	out, err := execute("fnms", "file", "get", "--filename", "export.zip")
	if err != nil || !bytes.Equal([]byte(out), payload) {
		t.Fatalf("raw output %q err=%v", out, err)
	}
	before := requests
	for _, flag := range []string{"--out-jq", "--out-fields"} {
		_, err = execute("fnms", "file", "get", "--filename", "export.zip", flag, ".")
		if err == nil {
			t.Errorf("binary accepted %s", flag)
		}
	}
	out, err = execute("fnms", "report", "execute-async", "create", "--id", "42", "--dry-run", "--filter-params", `{"Name":"dev"}`)
	if err != nil || !strings.Contains(out, "dryRun") || requests != before {
		t.Fatalf("preview %s %v calls=%d before=%d", out, err, requests, before)
	}
	out, err = execute("fnms", "report", "execute-async", "create", "--id", "42")
	if err != nil || !strings.Contains(out, "job") {
		t.Fatalf("non-destructive report confirmation %s %v", out, err)
	}
	out, err = execute("cli", "schema", "fnms", "report", "exports", "create")
	if err != nil || !strings.Contains(out, "filter-params") || !strings.Contains(out, "Accept") {
		t.Fatalf("schema flags/headers %s %v", out, err)
	}
}

func TestFNMSOAuthDebugDownload(t *testing.T) {
	var tokenRequests int
	external := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "" {
			t.Error("OAuth credentials leaked")
		}
		fmt.Fprint(w, "binary\x00bytes")
	}))
	defer external.Close()
	api := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "token") {
			tokenRequests++
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{"access_token":"offline-oauth","token_type":"Bearer","expires_in":3600}`)
			return
		}
		if r.Header.Get("Authorization") != "Bearer offline-oauth" {
			t.Error("missing OAuth API auth")
		}
		w.Header().Set("Location", external.URL)
		w.WriteHeader(303)
	}))
	defer api.Close()
	var out, stderr bytes.Buffer
	root, _ := app.NewRootCmd(&out, &stderr, func(string) string { return "" }, api.Client(), "test")
	root.SetArgs([]string{"fnms", "file", "get", "--filename", "export.zip", "--config", os.DevNull, "--org-id", "7", "--api-base-url", api.URL, "--login-base-url", api.URL, "--client-id", "offline-client", "--refresh-token", "offline-refresh", "--debug"})
	err := root.Execute()
	if err != nil || out.String() != "binary\x00bytes" || tokenRequests != 1 {
		t.Fatalf("output %q err=%v token requests=%d", out.String(), err, tokenRequests)
	}
	if strings.Contains(stderr.String(), "offline-oauth") || strings.Contains(stderr.String(), "offline-refresh") {
		t.Fatal("debug disclosed credentials")
	}
}

func TestFNMSExactOperationsAndWorkflowMetadata(t *testing.T) {
	root, _ := app.NewRootCmd(io.Discard, io.Discard, func(string) string { return "" }, nil, "test")
	for path, id := range map[string]string{"fnms file get": "Fnms_File_show", "fnms report execute-async create": "Fnms_Report_execute_async", "fnms report exports create": "Fnms_Report_export_report", "fnms report execute": "", "fnms report export": "", "fnms report wait-export": "", "fnms report wait-execution": "", "fnms report download-export": ""} {
		cmd, remaining, err := root.Find(strings.Fields(path))
		if err != nil || len(remaining) > 0 || cmd.Annotations["flexera.operationId"] != id {
			t.Fatalf("metadata %s %v remaining=%v annotations=%v", path, err, remaining, cmd.Annotations)
		}
	}
}
