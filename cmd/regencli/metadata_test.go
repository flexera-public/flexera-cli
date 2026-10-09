package main

import (
	"bytes"
	"encoding/json"
	"path/filepath"
	"testing"

	"github.com/flexera-public/flexera-cli/internal/catalog"
)

func TestCatalogOperationMetadata(t *testing.T) {
	stage := t.TempDir()
	e := catalog.Entry{OperationID: "test", Command: []string{"test", "create"}, Method: "POST", Path: "/test", ResponseEnvelope: "none"}
	data, _ := json.Marshal(map[string]catalog.Entry{"test": e})
	writeTestFile(t, filepath.Join(stage, "test", "metadata.json"), string(data))
	spec := []byte(`{"paths":{"/test":{"post":{
		"operationId":"test","deprecated":true,"externalDocs":{"url":"https://example.invalid/docs"},
		"parameters":[{"name":"Api-Version","in":"header","required":true,"description":"Version guidance","schema":{"type":"string","enum":["1.0"]},"example":"1.0"}],
		"requestBody":{"required":true,"description":"Body guidance","content":{"application/json":{},"application/octet-stream":{}}}
	}}},"components":{"schemas":{}}}`)
	raw, err := buildCatalog(stage, []genTag{{Pkg: "test", Cmd: "test"}}, spec)
	if err != nil {
		t.Fatal(err)
	}
	var doc catalog.Document
	if err := json.Unmarshal(raw, &doc); err != nil {
		t.Fatal(err)
	}
	e = doc.Entries[0]
	if !e.Deprecated || !e.RequestRequired || e.RequestDescription != "Body guidance" || len(e.ExternalDocs) == 0 || len(e.RequestMediaTypes) != 2 || e.RequestMediaTypes[0] != "application/json" {
		t.Fatalf("operation metadata lost: %+v", e)
	}
	if len(e.Headers) != 1 || e.Headers[0].Name != "Api-Version" || !e.Headers[0].Required || string(e.Headers[0].Example) != `"1.0"` {
		t.Fatalf("header metadata lost: %+v", e.Headers)
	}
}

func TestNewMetadataSampleSafety(t *testing.T) {
	doc := catalog.Document{Schemas: map[string]json.RawMessage{}, Entries: []catalog.Entry{{
		OperationID: "test", Command: []string{"test"}, Method: "POST", Path: "/test", ResponseEnvelope: "none",
		Params: []catalog.Param{
			{Flag: "key", Name: "api_key", In: "query", Source: "flag", Schema: json.RawMessage(`{"type":"string"}`), Example: json.RawMessage(`"hidden-param"`), Examples: json.RawMessage(`{"named":{"value":"hidden-param"}}`)},
			{Flag: "filter", In: "query", Source: "flag", Schema: json.RawMessage(`{"type":"object","properties":{"value":{"type":"string","writeOnly":true}}}`), Examples: json.RawMessage(`{"named":{"value":{"value":"hidden-named"}}}`)},
			{Flag: "id", In: "query", Source: "flag", Schema: json.RawMessage(`{"type":"integer"}`), Example: json.RawMessage(`9007199254740993`), Examples: json.RawMessage(`{"named":{"value":9007199254740993}}`)},
		},
		BodyFlags:  []string{"credential"},
		BodyFields: []catalog.BodyField{{Flag: "credential", Property: "client_secret", Schema: json.RawMessage(`{"type":"string","example":"hidden-body"}`)}},
		Headers:    []catalog.Header{{Name: "Authorization", Schema: json.RawMessage(`{"type":"string","example":"hidden-header"}`), Example: json.RawMessage(`"hidden-header"`), Examples: json.RawMessage(`{"named":{"value":"hidden-header"}}`)}},
	}}}
	if err := sanitizeCatalogExamples(&doc); err != nil {
		t.Fatal(err)
	}
	raw, err := json.Marshal(doc)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Contains(raw, []byte("hidden-")) || bytes.Count(raw, []byte("9007199254740993")) != 2 {
		t.Fatalf("unsafe or lossy sample publication: %s", raw)
	}
}
