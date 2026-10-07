package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/spf13/cobra"
)

func TestPlanRedactsWithoutChangingRequest(t *testing.T) {
	body := json.RawMessage(`{"name":"test","client_secret":"private-secret","nested":{"value":"private-value"}}`)
	var schema openapi3.Schema
	if err := json.Unmarshal([]byte(`{"type":"object","properties":{"nested":{"type":"object","properties":{"value":{"type":"string","writeOnly":true}}}}}`), &schema); err != nil {
		t.Fatal(err)
	}
	plan := Plan{Command: "budget create", Method: "POST", Path: "/budgets", OrgID: 123, Params: map[string]any{"id": "resource", "access-token": "private-token"}, Body: body, Validation: &ValidationResult{Status: "ok"}, RequestSchema: &schema}
	code, err := CompileJQ(`error("not a response")`)
	if err != nil {
		t.Fatal(err)
	}
	var out bytes.Buffer
	if err := plan.RenderJSON(&out, Printer{JQ: code, Style: JSONStyleCompact}); err != nil {
		t.Fatal(err)
	}
	if strings.Contains(out.String(), "private-") || !strings.Contains(out.String(), `"redacted":true`) || !strings.Contains(out.String(), `"status":"ok"`) {
		t.Fatalf("unsafe preview %s", out.String())
	}
	if string(plan.Body) != string(body) || !strings.Contains(string(body), "private-secret") {
		t.Fatal("request body mutated")
	}
	var value map[string]any
	if err := json.Unmarshal(out.Bytes(), &value); err != nil {
		t.Fatal(err)
	}
	if value["dryRun"] != true || value["plan"] == nil || value["destructive"] != false {
		t.Fatal(value)
	}
}

func TestPlanRawUploadAndSkippedStatus(t *testing.T) {
	for _, raw := range []bool{false, true} {
		var out bytes.Buffer
		plan := Plan{Method: "POST", Path: "/test", Validation: &ValidationResult{Status: "skipped"}, Body: json.RawMessage(`{"name":"test"}`), RawUpload: raw}
		if raw {
			plan.Body = []byte("\x00private-bytes")
			plan.Validation.Status = "unsupported"
		}
		if err := plan.RenderJSON(&out, Printer{}); err != nil {
			t.Fatal(err)
		}
		if raw && strings.Contains(out.String(), "private-bytes") {
			t.Fatal("raw bytes leaked")
		}
		if !strings.Contains(out.String(), plan.Validation.Status) {
			t.Fatal(out.String())
		}
	}
}

func TestNoValidateIsCLIOnly(t *testing.T) {
	for _, key := range persistentKeys {
		if key == FlagNoValidate {
			t.Fatal("no-validate config bound")
		}
	}
}

func TestNoValidateRejectedUntilCommandIntegration(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	var out, stderr bytes.Buffer
	root, deps := NewRootCmd(RootOptions{})
	root.SetOut(&out)
	root.SetErr(&stderr)
	root.AddCommand(&cobra.Command{Use: "legacy", RunE: func(*cobra.Command, []string) error { t.Fatal("unintegrated command ran"); return nil }})
	if code := Execute(context.Background(), root, deps, []string{"legacy", "--no-validate"}); code != 2 || out.Len() != 0 || !strings.Contains(stderr.String(), "does not support --no-validate") {
		t.Fatalf("silent unsupported flag: %d %s", code, stderr.String())
	}
}

func TestComposedSchemaRedaction(t *testing.T) {
	var schema openapi3.Schema
	if err := json.Unmarshal([]byte(`{"allOf":[{"type":"string","writeOnly":true}]}`), &schema); err != nil {
		t.Fatal(err)
	}
	value, changed := redactPreview("private-value", []*openapi3.Schema{&schema})
	if !changed || value != "[REDACTED]" {
		t.Fatalf("composition sensitive value leaked: %v", value)
	}
}

func TestHumanPlanShowsEveryEffectiveInputSafely(t *testing.T) {
	plan := Plan{Command: "flexera-cli budget create", Method: "POST", Path: "/budgets", OrgID: 123, Params: map[string]any{"org-id": 123, "enabled": false, "limit": 0, "client-secret": "private"}, Body: json.RawMessage(`{"name":"sample","nested":{"a":1,"b":2},"password":"private"}`), Validation: &ValidationResult{Status: "ok"}}
	var out bytes.Buffer
	if err := plan.RenderHuman(&out); err != nil {
		t.Fatal(err)
	}
	text := out.String()
	for _, want := range []string{`"enabled" = false`, `"limit" = 0`, `"name" = "sample"`, `"nested" = {"a":1,"b":2}`, `[REDACTED]`, `not a server-side diff`, `not replayable`, `No reusable command or body file has been exported`} {
		if !strings.Contains(text, want) {
			t.Errorf("missing %s: %s", want, text)
		}
	}
	if strings.Contains(text, "private") && strings.Contains(text, `"private"`) {
		t.Fatal("secret exposed")
	}
	if strings.Contains(text, "--body @request.json") {
		t.Fatal("invented nonexistent request file")
	}
	if strings.Contains(text, "... ") {
		t.Fatal("plan omitted inputs")
	}
}
