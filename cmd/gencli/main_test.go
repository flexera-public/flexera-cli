package main

import (
	"encoding/json"
	"fmt"
	"go/format"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestGeneratedOutputContracts(t *testing.T) {
	for _, tc := range []struct {
		name, want string
		responses  []successResponse
	}{
		{"json", "structured", []successResponse{{Code: "200", HasJSON: true, HasSchema: true, MediaType: "application/json"}}},
		{"no-content", "text", []successResponse{{Code: "204"}}},
		{"text-schema", "text", []successResponse{{Code: "200", HasSchema: true, MediaType: "text/plain"}}},
		{"mixed", "mixed", []successResponse{{Code: "200", HasJSON: true, HasSchema: true, MediaType: "application/json"}, {Code: "204"}}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			op := operation{Method: "get", Path: "/test", OperationID: "Test_show", Action: "get", Success2xx: tc.responses[0].Code, SuccessResponses: tc.responses}
			src, err := render("Test", "test", "test", []operation{op})
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(string(src), `"flexera.output": "`+tc.want+`"`) {
				t.Fatalf("missing output contract %s: %s", tc.want, src)
			}
			if _, err := format.Source(src); err != nil {
				t.Fatal(err)
			}
			if tc.name == "text-schema" && !strings.Contains(string(src), "deps.Stdout.Write(resp.Body)") {
				t.Fatal("nonJSON bytes decoded instead of preserved")
			}
		})
	}
}

func TestExtractParamsFormattedQuery(t *testing.T) {
	for _, format := range []string{"date", "date-time"} {
		for _, required := range []bool{false, true} {
			op := map[string]interface{}{"parameters": []interface{}{map[string]interface{}{
				"name": "when", "in": "query", "required": required,
				"schema": map[string]interface{}{"type": "string", "format": format},
			}}}
			_, query, _, ok := extractParams(op)
			want := "time.Time"
			if format == "date" {
				want = "openapi_types.Date"
			}
			if !ok || len(query) != 1 || query[0].GoType != want || query[0].TimeLayout == "" || query[0].Required != required {
				t.Fatalf("%s required=%v: %+v, ok=%v", format, required, query, ok)
			}
			p := op["parameters"].([]interface{})[0].(map[string]interface{})
			p["in"] = "path"
			if _, _, _, ok := extractParams(op); ok {
				t.Fatal("formatted path unexpectedly supported")
			}
			p["in"] = "query"
			p["schema"] = map[string]interface{}{"type": "array", "items": map[string]interface{}{"type": "string", "format": format}}
			if _, _, _, ok := extractParams(op); ok {
				t.Fatal("formatted array unexpectedly supported")
			}
		}
	}
}

func TestCollectOpsUnsupportedAudit(t *testing.T) {
	// Existing silent gaps: unvisited methods and ignored ref/cookie params.
	// This audit records evidence; it does not expand date-query scope.
	for _, tc := range []struct {
		name, method, parameter string
		wantOps                 int
	}{
		{"unvisited HEAD", "head", "", 0},
		{"ignored parameter ref", "get", `{"$ref":"#/components/parameters/When"}`, 1},
		{"ignored cookie", "get", `{"name":"session","in":"cookie","schema":{"type":"string"}}`, 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			parameters := "[]"
			if tc.parameter != "" {
				parameters = "[" + tc.parameter + "]"
			}
			raw := json.RawMessage(fmt.Sprintf(`{%q:{"tags":["Audit"],"operationId":"Audit_get","x-flexera-action":"get","parameters":%s,"responses":{"200":{}}}}`, tc.method, parameters))
			ops, drops := collectOps(&spec{Paths: map[string]json.RawMessage{"/audit": raw}}, "Audit")
			if len(ops) != tc.wantOps || len(drops) != 0 {
				t.Fatalf("ops=%+v drops=%+v", ops, drops)
			}
			if len(ops) > 0 && len(ops[0].QueryParams)+len(ops[0].PathParams) != 0 {
				t.Fatal("ignored parameter unexpectedly emitted")
			}
		})
	}
}

func TestGeneratedUsageMessageQuerySDK(t *testing.T) {
	data, err := os.ReadFile("../../unified-openapi/openapi3.json")
	if err != nil {
		t.Fatal(err)
	}
	var s spec
	if err := json.Unmarshal(data, &s); err != nil {
		t.Fatal(err)
	}
	ops, drops := collectOps(&s, "Usage Message Query")
	if len(drops) != 0 || len(ops) != 1 || ops[0].OperationID != "Saas_Usage_Message_Query_index" {
		t.Fatalf("actual tag operations=%+v drops=%+v", ops, drops)
	}
	if wrong, _ := collectOps(&s, "Saas Usage Message Query"); len(wrong) != 0 {
		t.Fatal("incorrect tag matched")
	}
	src, metadata, err := renderWithMetadata("Usage Message Query", "datequerytest", "usage-message-query", ops)
	if err != nil {
		t.Fatal(err)
	}
	if metadata[ops[0].OperationID] == nil {
		t.Fatal("usage query metadata missing")
	}
	for _, text := range []string{"flexera.SaasUsageMessageQueryIndexParams", "client.SaasUsageMessageQueryIndexWithResponse", "params.StartedAt = &startedAtValue", "params.EndedAt = &endedAtValue"} {
		if !strings.Contains(string(src), text) {
			t.Fatalf("missing %s", text)
		}
	}
	// There are no date-only query params in the pinned spec. Exercise their
	// generated parsing/assignment with a fixture using the SDK's concrete Date
	// type and query serializer; the real operation above compiles unchanged.
	dateOp := operation{OperationID: "Date_Fixture", Method: "get", Path: "/dates", Action: "get", Success2xx: "204"}
	_, dateOp.QueryParams, _, _ = extractParams(map[string]interface{}{"parameters": []interface{}{
		map[string]interface{}{"name": "optional", "in": "query", "schema": map[string]interface{}{"type": "string", "format": "date"}},
		map[string]interface{}{"name": "required", "in": "query", "required": true, "schema": map[string]interface{}{"type": "string", "format": "date"}},
	}})
	dateSrc, err := render("Date", "datequerytest", "date", []operation{dateOp})
	if err != nil {
		t.Fatal(err)
	}
	dateSrc = []byte(strings.ReplaceAll(string(dateSrc), "NewCmd", "NewDateCmd"))
	dateSrc = []byte(strings.ReplaceAll(string(dateSrc), "flexera.DateFixtureParams", "fixtureDateParams"))
	dateSrc = []byte(strings.ReplaceAll(string(dateSrc), "client.DateFixtureWithResponse", "fixtureDateWithResponse"))
	dateSrc = []byte(strings.ReplaceAll(string(dateSrc), `clipkg.ValidateCommandParams(cmd, "Date_Fixture")`, `fixtureDateParamsValidation(cmd)`))
	dateSrc = []byte(strings.ReplaceAll(string(dateSrc), "client, err := deps.APIClient()", "_, err = deps.APIClient()"))
	dir, err := os.MkdirTemp("../..", ".gencli-date-test-")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(dir) })
	for name, content := range map[string][]byte{"usage.go": src, "date.go": dateSrc, "query_test.go": []byte(generatedDateQueryTests)} {
		formatted, err := format.Source(content)
		if err != nil {
			t.Fatalf("%s: %v", name, err)
		}
		if err := os.WriteFile(filepath.Join(dir, name), formatted, 0600); err != nil {
			t.Fatal(err)
		}
	}
	cmd := exec.Command("go", "test", "-count=1", ".")
	cmd.Dir = dir
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("generated package: %v\n%s", err, out)
	}
}

const generatedDateQueryTests = `package datequerytest
import (
 "bytes"
 "context"
 "errors"
 "net/http"
 "net/url"
 "testing"
 "time"
 "io"
 "strings"
 clipkg "github.com/flexera-public/flexera-cli/internal/cli"
 cfg "github.com/flexera-public/flexera-cli/internal/config"
 flexera "github.com/flexera-public/unified-go-client"
 "github.com/oapi-codegen/runtime"
 "github.com/spf13/cobra"
 openapi_types "github.com/oapi-codegen/runtime/types"
)
func fixtureDateParamsValidation(cmd *cobra.Command)(map[string]any,error){return map[string]any{},nil}
// Compile-time checks use the active module pin, not the sibling SDK checkout.
var _ *time.Time = (flexera.SaasUsageMessageQueryIndexParams{}).StartedAt
var _ *time.Time = (flexera.SaasUsageMessageQueryIndexParams{}).EndedAt
var _ = (*flexera.ClientWithResponses).SaasUsageMessageQueryIndexWithResponse
type doer struct { calls int; query url.Values }
func (d *doer) Do(r *http.Request) (*http.Response, error) {
 d.calls++; d.query = r.URL.Query()
 return &http.Response{StatusCode:200, Header:http.Header{"Content-Type":[]string{"application/json"}}, Body:io.NopCloser(strings.NewReader("{}"))}, nil
}
func TestUsageWire(t *testing.T) {
 for _, tc := range []struct{ name string; args []string; start, end string; bad bool }{
  {name:"omitted"},
  {name:"UTC",args:[]string{"--started-at","2021-06-28T00:00:00Z"},start:"2021-06-28T00:00:00Z"},
  {name:"offsets",args:[]string{"--started-at","2021-06-28T01:02:03.123456789+05:30","--ended-at","2022-06-28T01:02:03-07:00"},start:"2021-06-28T01:02:03.123456789+05:30",end:"2022-06-28T01:02:03-07:00"},
  {name:"empty",args:[]string{"--started-at="},bad:true},
  {name:"missing zone",args:[]string{"--started-at","2021-06-28T00:00:00"},bad:true},
  {name:"invalid end",args:[]string{"--ended-at","not-a-date"},bad:true},
  {name:"invalid calendar",args:[]string{"--started-at","2021-02-30T00:00:00Z"},bad:true},
 } { t.Run(tc.name,func(t *testing.T) {
  d:=&doer{}; var out bytes.Buffer
  deps:=&clipkg.Deps{Config:cfg.CommonConfig{OrgID:100, AccessToken:"test", Output:"json", APIBaseURL:"https://example.invalid"},HTTP:d,Stdout:&out}
  if tc.bad { deps.Config.AccessToken="" } // parsing must win over auth errors
  c:=NewCmd(); c.SetArgs(append([]string{"list"},tc.args...)); c.SilenceErrors=true; c.SilenceUsage=true
  err:=c.ExecuteContext(clipkg.WithDeps(context.Background(),deps))
  if tc.bad { var ee *clipkg.ExitError; if !errors.As(err,&ee)||ee.Code!=2||!strings.Contains(err.Error(),"RFC3339")||d.calls!=0 { t.Fatalf("err=%v calls=%d",err,d.calls) }; return }
  if err!=nil {t.Fatal(err)}
  if d.calls!=1||d.query.Get("startedAt")!=tc.start||d.query.Get("endedAt")!=tc.end {t.Fatalf("calls=%d query=%v",d.calls,d.query)}
  if tc.start=="" && d.query.Has("startedAt") {t.Fatal("omitted timestamp serialized")}
  if tc.end=="" && d.query.Has("endedAt") {t.Fatal("omitted end serialized")}
 }) }
}
type fixtureDateParams struct { Optional *openapi_types.Date; Required openapi_types.Date }
var dateQuery url.Values
func fixtureDateWithResponse(_ context.Context, p *fixtureDateParams) (*flexera.SaasUsageMessageQueryIndexResponse,error) {
 dateQuery=url.Values{}
 values:=map[string]openapi_types.Date{"required":p.Required}
 if p.Optional!=nil { values["optional"]=*p.Optional }
 for k,v:=range values {s,err:=runtime.StyleParamWithLocation("form",true,k,runtime.ParamLocationQuery,v);if err!=nil{return nil,err}; q,err:=url.ParseQuery(s);if err!=nil{return nil,err};dateQuery[k]=q[k]}
 return &flexera.SaasUsageMessageQueryIndexResponse{HTTPResponse:&http.Response{StatusCode:204}},nil
}
func TestDateWire(t *testing.T) {
 for _,tc:=range []struct{args []string; bad bool}{
  {args:[]string{"--required","2024-02-29"}},
  {args:[]string{"--required","2024-02-29","--optional","2021-06-28"}},
  {bad:true}, {args:[]string{"--required="},bad:true},
  {args:[]string{"--required","2023-02-29"},bad:true},
  {args:[]string{"--required","2024-02-29","--optional="},bad:true},
  {args:[]string{"--required","2024-02-29T00:00:00Z"},bad:true},
 } {
  dateQuery=nil; var out bytes.Buffer
  deps:=&clipkg.Deps{Config:cfg.CommonConfig{AccessToken:"test", APIBaseURL:"https://example.invalid"},Stdout:&out}
  c:=NewDateCmd();c.SetArgs(append([]string{"get"},tc.args...));c.SilenceErrors=true;c.SilenceUsage=true
  err:=c.ExecuteContext(clipkg.WithDeps(context.Background(),deps))
  if tc.bad {var ee *clipkg.ExitError;if !errors.As(err,&ee)||ee.Code!=2||!strings.Contains(err.Error(),"YYYY-MM-DD")||dateQuery!=nil {t.Fatalf("err=%v query=%v",err,dateQuery)};continue}
  if err!=nil {t.Fatal(err)}
  if dateQuery.Get("required")!="2024-02-29" {t.Fatal(dateQuery)}
  if len(tc.args)==2 && dateQuery.Has("optional") {t.Fatal("optional date serialized")}
  if len(tc.args)>2 && dateQuery.Get("optional")!="2021-06-28" {t.Fatal(dateQuery)}
 }
}
`

func TestCollectOps_DropsNonJSONBody(t *testing.T) {
	post := map[string]interface{}{
		"tags":             []interface{}{"sample"},
		"x-flexera-action": "create",
		"operationId":      "createWidget",
		"requestBody": map[string]interface{}{
			"content": map[string]interface{}{
				"multipart/form-data": map[string]interface{}{},
			},
		},
		"responses": map[string]interface{}{
			"200": map[string]interface{}{"description": "ok"},
		},
	}
	pathItem := map[string]interface{}{"post": post}
	rawPath, err := json.Marshal(pathItem)
	if err != nil {
		t.Fatal(err)
	}
	s := &spec{Paths: map[string]json.RawMessage{"/foo": rawPath}}

	ops, drops := collectOps(s, "sample")
	if len(ops) != 0 {
		t.Fatalf("expected 0 ops, got %d", len(ops))
	}
	if len(drops) != 1 {
		t.Fatalf("expected 1 drop, got %d", len(drops))
	}
	if !strings.Contains(drops[0].Reason, "non-JSON request body") {
		t.Fatalf("expected non-JSON drop reason, got %q", drops[0].Reason)
	}
	if drops[0].Method != "post" || drops[0].Path != "/foo" {
		t.Fatalf("unexpected drop coords: %+v", drops[0])
	}
}

func TestSuccessResponsesOAS3(t *testing.T) {
	op := map[string]interface{}{"responses": map[string]interface{}{
		"202":     map[string]interface{}{"content": map[string]interface{}{"application/json": map[string]interface{}{"schema": map[string]interface{}{"$ref": "#/components/schemas/Accepted"}}}},
		"201":     map[string]interface{}{"content": map[string]interface{}{"application/json": map[string]interface{}{"schema": map[string]interface{}{"type": "object"}}}},
		"203":     map[string]interface{}{"content": map[string]interface{}{"application/json": map[string]interface{}{}}},
		"200":     map[string]interface{}{"schema": map[string]interface{}{"type": "object"}},
		"default": map[string]interface{}{}, "2XX": map[string]interface{}{},
	}}
	code, jsonBody, schema := successCode(op)
	if code != "200" || jsonBody || schema {
		t.Fatalf("OAS2 schema must not be selected: %s %v %v", code, jsonBody, schema)
	}
	delete(op["responses"].(map[string]interface{}), "200")
	code, jsonBody, schema = successCode(op)
	if code != "201" || !jsonBody || !schema {
		t.Fatalf("unexpected primary: %s %v %v", code, jsonBody, schema)
	}
	responses := successResponses(op)
	if len(responses) != 3 || responses[1].Schema["$ref"] != "#/components/schemas/Accepted" || !responses[2].HasJSON || responses[2].HasSchema {
		t.Fatalf("responses: %+v", responses)
	}
	src, _, err := renderWithMetadata("Sample", "sample", "sample", []operation{{Method: "post", Path: "/widgets", OperationID: "Sample_create", Action: "create", Success2xx: code, SuccessResponses: responses}})
	if err != nil {
		t.Fatal(err)
	}
	for _, text := range []string{"case 201:", "case 202:", "case 203:", "resp.JSON201", "resp.JSON202", "clipkg.DecodeResponseJSON(resp.Body)"} {
		if !strings.Contains(string(src), text) {
			t.Errorf("missing %q", text)
		}
	}
	if strings.Contains(string(src), "resp.JSON203") {
		t.Error("schema-less JSON response must not use a generated typed response field")
	}
	if _, err := format.Source(src); err != nil {
		t.Fatal(err)
	}
}

func TestMetadataFlagsAndExamples(t *testing.T) {
	org := param{Name: "orgId", GoName: "orgID", FlagName: "org-id", In: "query", GoType: "int64", Required: true, Schema: map[string]interface{}{"type": "integer"}}
	client := param{Name: "clientId", GoName: "clientID", FlagName: "client-id", In: "path", GoType: "string", Required: true, Schema: map[string]interface{}{"type": "string"}}
	src, metadata, err := renderWithMetadata("Service Account Client", "sample", "service-account-client", []operation{{
		Method: "delete", Path: "/clients/{clientId}/revoke", OperationID: "Sample_revoke", Tag: "Service Account Client", Action: "action", Summary: "Revoke", Description: "Description",
		PathParams: []param{client}, QueryParams: []param{org}, HasBody: true, BodyTypeName: "SampleRevokeJSONRequestBody", Success2xx: "204",
		BodyFields:    []bodyField{{JSONKey: "orgId", FlagName: "org-id", GoVar: "fOrgID", GoType: "int"}},
		RequestSchema: map[string]interface{}{"type": "object"}, RequestExample: map[string]interface{}{},
	}})
	if err != nil {
		t.Fatal(err)
	}
	got := string(src)
	for _, want := range []string{"\"target-client-id\"", "\"body-org-id\"", "params.OrgId = int64(deps.Config.OrgID)", "flexera.operationId", "--body @request.json", "--org-id ORG_ID", "--dry-run"} {
		if !strings.Contains(got, want) {
			t.Errorf("missing %q", want)
		}
	}
	for _, bad := range []string{"Int64Var(&orgID", "StringVar(&clientID, \"client-id\"", "--yes", "<int>"} {
		if strings.Contains(got, bad) {
			t.Errorf("unexpected %q", bad)
		}
	}
	e := metadata["Sample_revoke"].(map[string]interface{})
	if e["destructive"] != true || e["description"] != "Description" {
		t.Fatalf("metadata: %+v", e)
	}
	command := e["command"].([]string)
	if strings.Join(command, " ") != "service-account-client revoke" {
		t.Fatal(command)
	}
	params := e["params"].([]interface{})
	if params[1].(map[string]interface{})["source"] != "config" {
		t.Fatal(params)
	}
	if _, err := format.Source(src); err != nil {
		t.Fatal(err)
	}
}

func TestBodyFlagFallbackAndExamplePrecedence(t *testing.T) {
	s := &spec{}
	media := map[string]interface{}{"schema": map[string]interface{}{"type": "object", "required": []interface{}{"orgId", "name", "readonly"}, "properties": map[string]interface{}{
		"orgId": map[string]interface{}{"type": "integer"}, "bodyOrgId": map[string]interface{}{"type": "integer"},
		"name":     map[string]interface{}{"type": "string", "enum": []interface{}{"example"}},
		"readonly": map[string]interface{}{"type": "string", "readOnly": true},
	}}}
	fields := extractBodyFields(s, media)
	for _, f := range fields {
		if f.JSONKey == "orgId" && f.FlagName != "body-body-org-id" {
			t.Fatalf("fallback stole natural spelling: %+v", f)
		}
		if f.JSONKey == "readonly" {
			t.Fatal("readOnly flag emitted")
		}
	}
	example := bodyExample(s, media).(map[string]interface{})
	if len(example) != 2 || example["name"] != "example" {
		t.Fatal(example)
	}
	media["examples"] = map[string]interface{}{"z": map[string]interface{}{"value": "last"}, "a": map[string]interface{}{"value": "first"}}
	if bodyExample(s, media) != "first" {
		t.Fatal("examples are not sorted")
	}
	media["example"] = "explicit"
	if bodyExample(s, media) != "explicit" {
		t.Fatal("wrong media example precedence")
	}
}

func TestSpecMetadataDeterministic(t *testing.T) {
	data, err := os.ReadFile("../../unified-openapi/openapi3.json")
	if err != nil {
		t.Fatal(err)
	}
	var s spec
	if err := json.Unmarshal(data, &s); err != nil {
		t.Fatal(err)
	}
	for _, tag := range []string{"Budget", "Service Account Client", "Bill Analysis"} {
		ops, _ := collectOps(&s, tag)
		if len(ops) == 0 {
			continue
		}
		src, metadata, err := renderWithMetadata(tag, "sample", kebab(tag), ops)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := format.Source(src); err != nil {
			t.Fatalf("%s: %v", tag, err)
		}
		_, again, err := renderWithMetadata(tag, "sample", kebab(tag), ops)
		if err != nil {
			t.Fatal(err)
		}
		a, _ := json.Marshal(metadata)
		b, _ := json.Marshal(again)
		if string(a) != string(b) || len(metadata) != len(ops) {
			t.Fatalf("%s: non-deterministic or missing entries", tag)
		}
	}
}

func TestMetadataAssignedVerbsAndPagination(t *testing.T) {
	ops := []operation{
		{Method: "get", Path: "/orgs/{orgId}/widgets", OperationID: "Org_list", Action: "list", Success2xx: "200", HasJSONResp: true, Paginated: true,
			PathParams:  []param{{Name: "orgId", FlagName: "org-id", In: "path", GoType: "int", Required: true}},
			QueryParams: []param{{Name: "skipToken", FlagName: "skip-token", In: "query", GoType: "string"}}},
		{Method: "get", Path: "/projects/{projectId}/widgets", OperationID: "Project_list", Action: "list", Success2xx: "200"},
	}
	src, metadata, err := renderWithMetadata("Widget", "widget", "widget", ops)
	if err != nil {
		t.Fatal(err)
	}
	org := metadata["Org_list"].(map[string]interface{})
	project := metadata["Project_list"].(map[string]interface{})
	if org["command"].([]string)[1] != "list" || project["command"].([]string)[1] != "list-project" {
		t.Fatal(metadata)
	}
	if org["paginated"] != true || org["responseEnvelope"] != "values" || project["responseEnvelope"] != "none" {
		t.Fatal(metadata)
	}
	params := org["params"].([]interface{})
	if params[0].(map[string]interface{})["source"] != "config" || params[1].(map[string]interface{})["source"] != "pagination" {
		t.Fatal(params)
	}
	if !strings.Contains(string(src), "noPaginate bool") {
		t.Fatal("pagination-only variables missing")
	}
	if _, err := format.Source(src); err != nil {
		t.Fatal(err)
	}
}

func TestExampleReferenceCycle(t *testing.T) {
	s := &spec{}
	s.Components.Schemas = map[string]json.RawMessage{
		"Node": json.RawMessage(`{"type":"object","required":["child","name"],"properties":{"child":{"$ref":"#/components/schemas/Node"},"name":{"type":"string"}}}`),
	}
	example := bodyExample(s, map[string]interface{}{"schema": map[string]interface{}{"$ref": "#/components/schemas/Node"}})
	if example != nil {
		t.Fatalf("required cycle must not advertise a partial body: %v", example)
	}
}

func TestCollectOps_AcceptsOctetStreamBody(t *testing.T) {
	post := map[string]interface{}{
		"tags":             []interface{}{"sample"},
		"x-flexera-action": "action",
		"operationId":      "uploadFile",
		"requestBody": map[string]interface{}{
			"content": map[string]interface{}{
				"application/octet-stream": map[string]interface{}{},
			},
		},
		"responses": map[string]interface{}{
			"204": map[string]interface{}{"description": "ok"},
		},
	}
	rawPath, err := json.Marshal(map[string]interface{}{"post": post})
	if err != nil {
		t.Fatal(err)
	}
	ops, drops := collectOps(&spec{Paths: map[string]json.RawMessage{"/foo/{id}/files/{name}": rawPath}}, "sample")
	if len(drops) != 0 {
		t.Fatalf("expected no drops, got %v", drops)
	}
	if len(ops) != 1 || !ops[0].HasRawBody || ops[0].RawBodyType != "application/octet-stream" {
		t.Fatalf("unexpected raw-body operation: %+v", ops)
	}
}

// TestSupportedActionsCount asserts the supportedActions map carries
// exactly the seven verbs gencli emits cobra leaves for. "query" was
// reserved in the design but dropped per S3 plan-review since no
// annotateForCLI branch emits it.
func TestSupportedActionsCount(t *testing.T) {
	if got, want := len(supportedActions), 7; got != want {
		t.Fatalf("len(supportedActions) = %d, want %d", got, want)
	}
	for _, v := range []string{"list", "get", "create", "update", "replace", "delete", "action"} {
		if !supportedActions[v] {
			t.Errorf("supportedActions missing %q", v)
		}
	}
	if supportedActions["query"] {
		t.Errorf(`supportedActions still carries "query"; expected dropped`)
	}
}

func TestRender_ListsUntypedJSONResponses(t *testing.T) {
	src, err := render("Sample", "sample", "sample", []operation{{
		Method:        "get",
		Path:          "/orgs/{orgId}/widgets",
		OperationID:   "Sample_Widget_index",
		Action:        "list",
		Success2xx:    "200",
		HasSchemaResp: true,
	}})
	if err != nil {
		t.Fatal(err)
	}

	got := string(src)
	for _, want := range []string{
		"switch resp.StatusCode()",
		"case 200:",
		"clipkg.DecodeResponseJSON(resp.Body)",
		"return deps.Printer.Render(deps.Stdout, deps.Config.Output, result)",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("generated command does not contain %q:\n%s", want, got)
		}
	}
}
