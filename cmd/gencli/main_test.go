package main

import (
	"encoding/json"
	"strings"
	"testing"
)

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
