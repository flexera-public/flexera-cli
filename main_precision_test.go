package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"runtime/debug"
	"strings"
	"testing"

	flexera "github.com/flexera-public/unified-go-client"
)

// The template ID itself is a string. Use the principal's schema-backed int64
// ID instead: PolicyPublishedTemplateList -> Values -> CreatedBy -> Id.
// These are release-blocking regressions, not expected-failure tests: the SDK
// pin must be updated (or unsafe responses rejected) before shipping.
func TestPrecisionGeneratedPagination(t *testing.T) {
	precisionLogSDK(t)
	var principal flexera.PolicyApplicationVndFlexeraPolicyPrincipal
	var list flexera.PolicyPublishedTemplateList
	var template flexera.PolicyFlexeraPolicyPublishedTemplate
	var integerID *int64 = &principal.Id
	var integerCount *int64 = list.Count
	var integerTotal *int64 = list.Total
	var createdBy *flexera.PolicyApplicationVndFlexeraPolicyPrincipal = template.CreatedBy
	_, _, _, _ = integerID, integerCount, integerTotal, createdBy

	const first = `{"kind":"policy#published_template_list","count":9007199254740993,"total":9007199254740997,"nextPage":"https://precision.invalid/policy/v1/orgs/123/published-templates?skipToken=second","values":[{"id":"template-1","kind":"policy#published_template","name":"first","fingerprint":"fp-1","ref":"ref-1","createdBy":{"id":9007199254740993,"kind":"user","name":"First"}}]}`
	const second = `{"kind":"policy#published_template_list","count":2,"total":9007199254740999,"prevPage":"https://precision.invalid/previous","values":[{"id":"template-2","kind":"policy#published_template","name":"second","fingerprint":"fp-2","ref":"ref-2","createdBy":{"id":9007199254740995,"kind":"user","name":"Second"}}]}`
	for _, tc := range []struct {
		name      string
		flags     []string
		pages     int
		count     string
		total     string
		projected bool
		scalar    string
	}{
		{name: "merged", pages: 2, count: "9007199254740995", total: "9007199254740999"},
		{name: "merged_fields_then_jq", flags: []string{"--out-fields", "createdBy.id", "--out-jq", "."}, pages: 2, count: "9007199254740995", total: "9007199254740999", projected: true},
		{name: "first_page", flags: []string{"--no-paginate"}, pages: 1, count: "9007199254740993", total: "9007199254740997"},
		{name: "first_page_fields_then_jq", flags: []string{"--no-paginate", "--out-fields", "createdBy.id", "--out-jq", "."}, pages: 1, count: "9007199254740993", total: "9007199254740997", projected: true},
		{name: "merged_scalar", flags: []string{"--out-fields", "createdBy.id", "--out-jq", ".values[1].createdBy.id"}, pages: 2, scalar: "9007199254740995"},
		{name: "first_page_scalar", flags: []string{"--no-paginate", "--out-jq", ".values[0].createdBy.id"}, pages: 1, scalar: "9007199254740993"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			calls := 0
			doer := precisionDoer(func(r *http.Request) (*http.Response, error) {
				calls++
				precisionCheckRequest(t, r, http.MethodGet, "/policy/v1/orgs/123/published-templates")
				wantToken, body := "", first
				if calls == 2 {
					wantToken, body = "second", second
				}
				if calls > tc.pages || r.URL.Query().Get("skipToken") != wantToken {
					t.Fatalf("unexpected pagination request %d: %s", calls, r.URL)
				}
				return precisionResponse(r, body), nil
			})
			args := append([]string{"policy", "published-template", "list", "--org-id", "123", "--access-token", "precision-token", "--json-style", "compact"}, tc.flags...)
			var out, stderr bytes.Buffer
			code := run(context.Background(), args, &out, &stderr, func(string) string { return "" }, doer)
			if calls != tc.pages {
				t.Errorf("requests = %d, want %d", calls, tc.pages)
			}
			if code != 0 || stderr.Len() != 0 {
				t.Fatalf("exit=%d stdout=%s stderr=%s", code, &out, &stderr)
			}
			if tc.scalar != "" {
				precisionEqualJSON(t, out.Bytes(), []byte(tc.scalar))
				return
			}
			value := precisionDecode(t, out.Bytes()).(map[string]any)
			precisionEqualJSON(t, mustPrecisionJSON(t, value["count"]), []byte(tc.count))
			precisionEqualJSON(t, mustPrecisionJSON(t, value["total"]), []byte(tc.total))
			if value["kind"] != "policy#published_template_list" {
				t.Errorf("first-page metadata changed: %v", value["kind"])
			}
			rows, ok := value["values"].([]any)
			if !ok || len(rows) != tc.pages {
				t.Fatalf("values = %v, want %d rows", value["values"], tc.pages)
			}
			for i, row := range rows {
				object := row.(map[string]any)
				wantID := []string{"9007199254740993", "9007199254740995"}[i]
				precisionEqualJSON(t, mustPrecisionJSON(t, object["createdBy"].(map[string]any)["id"]), []byte(wantID))
				if tc.projected && len(object) != 1 {
					t.Errorf("field projection retained unselected keys: %v", object)
				}
			}
			if tc.pages == 1 {
				if _, ok := value["nextPage"].(string); !ok {
					t.Error("--no-paginate removed the resume cursor")
				}
			} else {
				for _, key := range []string{"nextPage", "prevPage"} {
					if _, ok := value[key]; ok {
						t.Errorf("merged response retained %s", key)
					}
				}
			}
		})
	}
}

// This generated operation has JSON200 *interface{}, so wire numbers can be
// lost before the CLI printer sees them. The pagination-only SDK fix does not
// fix this path. Either preserve the number or explicitly reject it; silently
// rounded success is never accepted, including for scalar JSON responses.
func TestPrecisionGeneratedUntypedResponse(t *testing.T) {
	precisionLogSDK(t)
	for _, tc := range []struct {
		name, body, want string
		flags            []string
	}{
		{"safe_object", `{"id":42}`, `{"id":42}`, nil},
		{"large_object", `{"id":9007199254740993}`, `{"id":9007199254740993}`, nil},
		{"large_object_fields_then_jq", `{"id":9007199254740993,"discard":true}`, `9007199254740993`, []string{"--out-fields", "id", "--out-jq", ".id"}},
		{"large_scalar", `9007199254740993`, `9007199254740993`, nil},
		{"large_scalar_jq", `9007199254740993`, `9007199254740993`, []string{"--out-jq", "."}},
		{"negative_scalar_jq", `-9007199254740993`, `-9007199254740993`, []string{"--out-jq", "."}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			calls := 0
			doer := precisionDoer(func(r *http.Request) (*http.Response, error) {
				calls++
				precisionCheckRequest(t, r, http.MethodPost, "/risk/v1/orgs/123/regulatory-compliance/asset-statistics")
				if calls != 1 {
					t.Fatal("nonpaginated operation made multiple requests")
				}
				return precisionResponse(r, tc.body), nil
			})
			args := append([]string{"risk", "regulatory-compliance", "asset-statistics", "--body", `{"metrics":[]}`, "--org-id", "123", "--access-token", "precision-token", "--json-style", "compact"}, tc.flags...)
			var out, stderr bytes.Buffer
			code := run(context.Background(), args, &out, &stderr, func(string) string { return "" }, doer)
			if calls != 1 {
				t.Fatalf("requests=%d exit=%d stderr=%s", calls, code, &stderr)
			}
			if code != 0 {
				if tc.name == "safe_object" || out.Len() != 0 {
					t.Fatalf("unexpected rejection/partial output: exit=%d stdout=%s stderr=%s", code, &out, &stderr)
				}
				problem, ok := precisionDecode(t, stderr.Bytes()).(map[string]any)
				message, _ := problem["error"].(string)
				if !ok || !strings.Contains(strings.ToLower(message), "precision") {
					t.Fatalf("rejection must explicitly identify precision loss: exit=%d stderr=%s", code, &stderr)
				}
				t.Logf("unsafe response explicitly rejected: exit=%d stderr=%s", code, &stderr)
				return
			}
			if stderr.Len() != 0 {
				t.Fatalf("successful response emitted stderr: %s", &stderr)
			}
			precisionEqualJSON(t, out.Bytes(), []byte(tc.want))
		})
	}
}

type precisionDoer func(*http.Request) (*http.Response, error)

func (d precisionDoer) Do(r *http.Request) (*http.Response, error) { return d(r) }

func precisionCheckRequest(t *testing.T, r *http.Request, method, path string) {
	t.Helper()
	if r.Method != method || r.URL.Path != path {
		t.Fatalf("unexpected request: %s %s", r.Method, r.URL)
	}
	if got := r.Header.Get("Authorization"); got != "Bearer precision-token" {
		t.Fatalf("Authorization = %q", got)
	}
}

func precisionResponse(r *http.Request, body string) *http.Response {
	return &http.Response{StatusCode: http.StatusOK, Header: http.Header{"Content-Type": {"application/json"}}, Body: io.NopCloser(strings.NewReader(body)), Request: r}
}

func precisionDecode(t *testing.T, raw []byte) any {
	t.Helper()
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.UseNumber() // Never round the assertion's expected or actual value.
	var value any
	if err := decoder.Decode(&value); err != nil {
		t.Fatalf("decode %q: %v", raw, err)
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		t.Fatalf("expected exactly one JSON value: %q (extra=%v, err=%v)", raw, extra, err)
	}
	return value
}

func precisionEqualJSON(t *testing.T, got, want []byte) {
	t.Helper()
	if !reflect.DeepEqual(precisionDecode(t, got), precisionDecode(t, want)) {
		t.Errorf("precision changed: got %s, want %s", got, want)
	}
}

func mustPrecisionJSON(t *testing.T, value any) []byte {
	t.Helper()
	raw, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	return raw
}

func precisionLogSDK(t *testing.T) {
	t.Helper()
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, dep := range info.Deps {
			if dep.Path == "github.com/flexera-public/unified-go-client" {
				description := dep.Path + "@" + dep.Version
				if dep.Replace != nil {
					description += fmt.Sprintf(" => %s@%s", dep.Replace.Path, dep.Replace.Version)
				}
				t.Log(description)
				return
			}
		}
	}
}
