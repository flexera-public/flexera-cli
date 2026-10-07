package app_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/flexera-public/flexera-cli/internal/app"
	"github.com/flexera-public/flexera-cli/internal/catalog"
	"github.com/flexera-public/flexera-cli/internal/cli"
)

type interactiveScript struct {
	approval  string
	approvals int
	edit      bool
	titles    []string
}

func (p *interactiveScript) Ask(_ context.Context, f cli.PromptField) (any, error) {
	p.titles = append(p.titles, f.Title)
	if f.Title == "Required parameter: org-id" {
		return json.Number("123"), nil
	}
	if f.Title == "Request body.name" && p.edit {
		return "edited budget", nil
	}
	if f.Value != nil {
		return f.Value, nil
	}
	if f.Schema.Type != nil && f.Schema.Type.Is("boolean") {
		return false, nil
	}
	if f.Schema.Type != nil && (f.Schema.Type.Is("number") || f.Schema.Type.Is("integer")) {
		return json.Number("130.7"), nil
	}
	return "sample", nil
}
func (p *interactiveScript) SelectFields(_ context.Context, title string, options, selected []string) ([]string, error) {
	return nil, nil
}
func (p *interactiveScript) Approve(context.Context) (string, error) {
	p.approvals++
	return p.approval, nil
}

type interactiveHTTP struct {
	t     *testing.T
	calls int
	body  []byte
	allow bool
}

func (d *interactiveHTTP) Do(r *http.Request) (*http.Response, error) {
	d.calls++
	if !d.allow {
		d.t.Fatal("interactive guard/preview/cancel attempted HTTP")
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	d.body = body
	return &http.Response{StatusCode: 201, Header: http.Header{"Content-Type": []string{"application/json"}}, Body: io.NopCloser(strings.NewReader("{}")), Request: r}, nil
}

func budgetFormBody(t *testing.T) string {
	t.Helper()
	index, err := catalog.Load()
	if err != nil {
		t.Fatal(err)
	}
	entry, found := index.Lookup("Budget_Budget_create")
	if !found {
		t.Fatal("missing budget")
	}
	var body map[string]json.RawMessage
	if err := json.Unmarshal(entry.RequestExample, &body); err != nil {
		t.Fatal(err)
	}
	delete(body, "segments")
	delete(body, "filter")
	delete(body, "dimensions")
	raw, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	return string(raw)
}

func TestLiveInteractivePlanApplyAndCancel(t *testing.T) {
	t.Setenv("EDITOR", "")
	t.Setenv("HOME", t.TempDir())
	t.Setenv("FLEXERA_CLI_ORG_ID", "")
	t.Setenv("FLEXERA_CLI_ACCESS_TOKEN", "")
	t.Setenv("FLEXERA_CLI_CLIENT_ID", "")
	t.Setenv("FLEXERA_CLI_CLIENT_SECRET", "")
	for _, tc := range []struct {
		name, answer    string
		dry, yes        bool
		code, approvals int
	}{{"preview", "", true, false, 0, 0}, {"auto-approve", "", false, true, 0, 0}, {"approve", "yes", false, false, 0, 1}, {"cancel", "Yes", false, false, 1, 1}} {
		t.Run(tc.name, func(t *testing.T) {
			var out, stderr bytes.Buffer
			script := &interactiveScript{approval: tc.answer, edit: true}
			doer := &interactiveHTTP{t: t, allow: !tc.dry && tc.code == 0}
			root, deps := app.NewRootCmd(&out, &stderr, os.Getenv, doer, "test")
			deps.Prompter = script
			deps.IsTerminal = func(any) bool { return true }
			args := []string{"budget", "create", "--interactive", "--body", budgetFormBody(t)}
			if tc.dry {
				args = append(args, "--dry-run")
			} else {
				args = append(args, "--access-token", "fixture-token")
			}
			if tc.yes {
				args = append(args, "--yes")
			}
			code := cli.Execute(context.Background(), root, deps, args)
			if code != tc.code {
				t.Fatalf("exit=%d stderr=%s", code, stderr.String())
			}
			if script.approvals != tc.approvals || len(script.titles) == 0 || script.titles[0] != "Required parameter: org-id" {
				t.Fatalf("wrong prompt order/approval: %+v", script)
			}
			if !strings.Contains(stderr.String(), "edited budget") {
				t.Fatal("plan does not use edited body")
			}
			if tc.dry {
				var plan map[string]any
				if err := json.Unmarshal(out.Bytes(), &plan); err != nil {
					t.Fatal(err)
				}
				body := plan["plan"].(map[string]any)["body"].(map[string]any)
				if body["name"] != "edited budget" || doer.calls != 0 {
					t.Fatalf("invalid dryrun %v", plan)
				}
			} else if tc.code == 0 {
				var body map[string]any
				if err := json.Unmarshal(doer.body, &body); err != nil {
					t.Fatal(err)
				}
				if body["name"] != "edited budget" || doer.calls != 1 {
					t.Fatalf("wire did not use edited input: %v calls=%d", body, doer.calls)
				}
			} else if doer.calls != 0 || !strings.Contains(stderr.String(), "Apply cancelled.") {
				t.Fatal("cancel applied or omitted notice")
			}
		})
	}
}

func TestLiveInteractiveTTYAndInputGuards(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	for _, tc := range []struct {
		name          string
		stdin, stderr bool
		body          string
		fragment      string
	}{{"stdin-piped", false, true, "", "terminal input and stderr"}, {"stderr-redirected", true, false, "", "terminal input and stderr"}, {"body-stdin", true, true, "@-", "cannot use --body @-"}} {
		t.Run(tc.name, func(t *testing.T) {
			var out, errOut bytes.Buffer
			input := strings.NewReader("must not read")
			script := &interactiveScript{}
			doer := &interactiveHTTP{t: t}
			root, deps := app.NewRootCmd(&out, &errOut, os.Getenv, doer, "test")
			root.SetIn(input)
			deps.Prompter = script
			deps.IsTerminal = func(value any) bool {
				if value == input {
					return tc.stdin
				}
				if value == &errOut {
					return tc.stderr
				}
				return false
			}
			args := []string{"budget", "create", "-i", "--dry-run"}
			if tc.body != "" {
				args = append(args, "--body", tc.body)
			}
			if code := cli.Execute(context.Background(), root, deps, args); code != 2 || !strings.Contains(errOut.String(), tc.fragment) || out.Len() != 0 || len(script.titles) != 0 || doer.calls != 0 {
				t.Fatalf("guard failed %d %s", code, errOut.String())
			}
			if input.Len() != len("must not read") {
				t.Fatal("guard consumed stdin")
			}
		})
	}
}

func TestNoImplicitInteractivePrompt(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	var out, stderr bytes.Buffer
	script := &interactiveScript{}
	root, deps := app.NewRootCmd(&out, &stderr, os.Getenv, &interactiveHTTP{t: t}, "test")
	deps.Prompter = script
	deps.IsTerminal = func(any) bool { return true }
	if code := cli.Execute(context.Background(), root, deps, []string{"budget", "create", "--org-id", "123", "--dry-run"}); code != 2 || len(script.titles) != 0 {
		t.Fatalf("implicit prompt: %d %v", code, script.titles)
	}
}
