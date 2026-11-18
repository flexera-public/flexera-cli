package fnms

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	clipkg "github.com/flexera-public/flexera-cli/internal/cli"
	cliconfig "github.com/flexera-public/flexera-cli/internal/config"
	"github.com/spf13/cobra"
)

func run(t *testing.T, c *cobra.Command, server *httptest.Server, args ...string) (string, error) {
	t.Helper()
	var out bytes.Buffer
	d := &clipkg.Deps{Config: cliconfig.CommonConfig{OrgID: 7, AccessToken: "offline-token", APIBaseURL: server.URL, Output: "json"}, HTTP: server.Client(), Stdout: &out, Stderr: io.Discard, Printer: clipkg.Printer{}, Getenv: func(string) string { return "" }}
	c.SetContext(clipkg.WithDeps(context.Background(), d))
	c.SetArgs(args)
	c.SetOut(&out)
	c.SetErr(io.Discard)
	c.SilenceUsage = true
	c.SilenceErrors = true
	err := c.Execute()
	return out.String(), err
}
func respond(w http.ResponseWriter, code int, body string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprint(w, body)
}

func TestReportFiltersAndDeduplication(t *testing.T) {
	for _, code := range []int{202, 409} {
		t.Run(fmt.Sprint(code), func(t *testing.T) {
			s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "POST" || r.URL.Path != "/fnms/v1/orgs/7/reports/42/execute-async" {
					t.Errorf("request %s %s", r.Method, r.URL.Path)
				}
				q := r.URL.Query()
				if q.Get("filterParams[DeviceName]") != "a & b" || q.Get("limit") != "5000" || q.Get("searchText") != "dev & test" || q.Get("forceNewExecution") != "false" {
					t.Errorf("query %v", q)
				}
				respond(w, code, `{"jobId":"existing","status":"Created"}`)
			}))
			defer s.Close()
			out, err := run(t, reportCmd("create", "Fnms_Report_execute_async", false, false), s, "--id", "42", "--filter-params", `{"DeviceName":"a & b"}`, "--limit", "5000", "--search-text", "dev & test", "--force-new-execution=false")
			if err != nil || !strings.Contains(out, `"jobId":"existing"`) {
				t.Fatalf("out=%s err=%v", out, err)
			}
		})
	}
}

func TestReportPreviewAndValidation(t *testing.T) {
	calls := 0
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { calls++; respond(w, 202, `{"jobId":"job"}`) }))
	defer s.Close()
	for _, args := range [][]string{{"--dry-run", "--filter-params", `{"password":"hide","Name":"dev"}`}, {"--filter-params", "null"}, {"--filter-params", "[]"}, {"--filter-params", "{"}, {"--limit", "4999"}, {"--limit", "10001"}} {
		out, err := run(t, reportCmd("create", "", false, false), s, append([]string{"--id", "42"}, args...)...)
		if args[0] == "--dry-run" {
			if err != nil || !strings.Contains(out, "dryRun") || strings.Contains(out, "hide") {
				t.Fatalf("preview %s %v", out, err)
			}
		} else if err == nil {
			t.Fatalf("accepted invalid args %v", args)
		}
	}
	if calls != 0 {
		t.Fatalf("validation/preview called API %d times", calls)
	}
	out, err := run(t, reportCmd("create", "", false, false), s, "--id", "42", "--yes")
	if err != nil || !strings.Contains(out, "job") || calls != 1 {
		t.Fatalf("confirmation %s %v %d", out, err, calls)
	}
}

func TestReportExportAndWait(t *testing.T) {
	polls := 0
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if r.Header.Get("Accept") != "text/csv" || r.URL.Query().Get("filterParams[Owner]") != "owner" {
				t.Errorf("export params %v %v", r.Header, r.URL.Query())
			}
			respond(w, 201, `{"data":{"exportId":"exp"}}`)
			return
		}
		polls++
		status := "In progress"
		if polls > 1 {
			status = "Completed"
		}
		respond(w, 200, fmt.Sprintf(`{"data":{"exportId":"exp","status":%q,"path":"/download"}}`, status))
	}))
	defer s.Close()
	out, err := run(t, reportCmd("export", "", true, true), s, "--id", "42", "--accept", "text/csv", "--filter-params", `{"Owner":"owner"}`, "--poll-interval", "1ms")
	if err != nil || polls != 2 || !strings.Contains(out, "/download") {
		t.Fatalf("%s %v polls=%d", out, err, polls)
	}
	for _, body := range []string{`{"data":{"status":"No data found"}}`, `{"data":{"status":"Errored"}}`, `{"data":{"status":"Completed"}}`, `{"data":{"status":"unknown"}}`} {
		s.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { respond(w, 200, body) })
		out, err = run(t, waitCmd(true), s, "--export-id", "exp", "--poll-interval", "1ms")
		if strings.Contains(body, "No data found") {
			if err != nil {
				t.Fatal(err)
			}
		} else if err == nil {
			t.Fatalf("accepted bad terminal state %s %s", body, out)
		}
	}
}

func TestExecutionWaitFailures(t *testing.T) {
	for _, tc := range []struct {
		body string
		code int
		want string
	}{{`{"jobId":"j","status":"Completed"}`, 200, ""}, {`{"jobId":"j","status":"Errored","errorMessage":"failed"}`, 200, "failed"}, {`{"jobId":"j","status":"unknown"}`, 200, "unknown"}, {`{"jobId":"j","status":"In progress"}`, 200, "deadline"}, {`{}`, 403, "403"}} {
		s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { respond(w, tc.code, tc.body) }))
		out, err := run(t, waitCmd(false), s, "--id", "42", "--job-id", "j", "--poll-interval", "1ms", "--poll-timeout", "10ms")
		s.Close()
		if tc.want == "" {
			if err != nil || !strings.Contains(out, "Completed") {
				t.Fatalf("%s %v", out, err)
			}
		} else if err == nil || !strings.Contains(err.Error(), tc.want) {
			t.Fatalf("want %s got %v", tc.want, err)
		}
	}
}

func TestDownloadAuthenticationAndBytes(t *testing.T) {
	payload := []byte{'P', 'K', 0, 255, '\n', '{', '}'}
	external := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "" || r.Header.Get("Cookie") != "" {
			t.Error("credentials forwarded to download")
		}
		if r.URL.Path == "/redirect" {
			w.Header().Set("Location", "/bytes")
			w.WriteHeader(307)
			return
		}
		w.Write(payload)
	}))
	defer external.Close()
	api := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer offline-token" {
			t.Error("missing API authentication")
		}
		w.Header().Set("Location", external.URL+"/redirect")
		w.WriteHeader(303)
	}))
	defer api.Close()
	out, err := run(t, downloadCmd("get", "Fnms_File_show"), api, "--filename", "export.zip")
	if err != nil || !bytes.Equal([]byte(out), payload) {
		t.Fatalf("bytes %q err=%v", out, err)
	}
	dir, err := os.MkdirTemp(".", ".fnms-test-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)
	path := filepath.Join(dir, "export.zip")
	out, err = run(t, downloadCmd("get", "Fnms_File_show"), api, "--filename", "export.zip", "--destination", path)
	data, readErr := os.ReadFile(path)
	if err != nil || readErr != nil || out != "" || !bytes.Equal(data, payload) {
		t.Fatalf("file bytes %q output=%s err=%v read=%v", data, out, err, readErr)
	}
	_, err = run(t, downloadCmd("get", "Fnms_File_show"), api, "--filename", "export.zip", "--destination", path)
	if err == nil {
		t.Fatal("overwrote existing file")
	}
	out, err = run(t, downloadCmd("download-export", ""), api, "--path", external.URL+"/bytes")
	if err != nil || !bytes.Equal([]byte(out), payload) {
		t.Fatalf("external export %q %v", out, err)
	}
	out, err = run(t, downloadCmd("download-export", ""), api, "--path", "/download")
	if err != nil || !bytes.Equal([]byte(out), payload) {
		t.Fatalf("same-origin export %q %v", out, err)
	}
	api.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(403) })
	failed := filepath.Join(dir, "failed.zip")
	_, err = run(t, downloadCmd("get", "Fnms_File_show"), api, "--filename", "bad.zip", "--destination", failed)
	if err == nil {
		t.Fatal("accepted API failure")
	}
	if _, statErr := os.Stat(failed); !os.IsNotExist(statErr) {
		t.Fatalf("failed file retained: %v", statErr)
	}
}

func TestExecuteWorkflowAndStartFailures(t *testing.T) {
	calls := 0
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		if r.Method == "POST" {
			respond(w, 202, `{"jobId":"j","status":"Created"}`)
			return
		}
		respond(w, 200, `{"jobId":"j","status":"Completed"}`)
	}))
	defer s.Close()
	out, err := run(t, reportCmd("execute", "", false, true), s, "--id", "42", "--poll-interval", "1ms")
	if err != nil || calls != 2 || !strings.Contains(out, "Completed") {
		t.Fatalf("execute %s %v calls=%d", out, err, calls)
	}
	for _, tc := range []struct {
		export bool
		code   int
		body   string
	}{{false, 202, `{}`}, {false, 409, `{}`}, {false, 403, `{}`}, {true, 201, `{"data":{}}`}, {true, 400, `{}`}} {
		s.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { respond(w, tc.code, tc.body) })
		_, err = run(t, reportCmd("create", "", tc.export, false), s, "--id", "42")
		if err == nil {
			t.Fatalf("accepted invalid start response %+v", tc)
		}
	}
}
