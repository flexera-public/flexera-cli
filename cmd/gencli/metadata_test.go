package main

import (
	"encoding/json"
	"testing"

	"github.com/flexera-public/flexera-cli/internal/catalog"
)

func TestInputMetadataFidelity(t *testing.T) {
	var s spec
	if err := json.Unmarshal([]byte(`{
		"paths":{"/widgets/{clientId}":{"post":{
			"tags":["Widget"],"operationId":"Widget_create","x-flexera-action":"create",
			"parameters":[
				{"name":"clientId","in":"path","required":true,"description":"Target client","example":9007199254740993,"schema":{"type":"string"}},
				{"name":"filter","in":"query","description":"Filter guidance","deprecated":true,"style":"form","explode":false,"allowReserved":true,
					"examples":{"named":{"summary":"Named guidance","value":"name eq 'sample'"}},"schema":{"type":"string"}}
			],
			"requestBody":{"required":true,"content":{"application/json":{"example":{"name":"sample"},"schema":{
				"type":"object","required":["name","orgId"],"properties":{
					"name":{"type":"string","description":"Display name","minLength":1},
					"orgId":{"type":"integer","description":"Target organization"}
				}
			}}}},
			"responses":{"201":{"description":"Created","content":{"application/json":{"schema":{"type":"object"}}}}}
		}}}
	}`), &s); err != nil {
		t.Fatal(err)
	}
	ops, drops := collectOps(&s, "Widget")
	if len(ops) != 1 || len(drops) != 0 {
		t.Fatalf("ops=%v drops=%v", ops, drops)
	}
	_, metadata, err := renderWithMetadata("Widget", "widget", "widget", ops)
	if err != nil {
		t.Fatal(err)
	}
	raw, err := json.Marshal(metadata["Widget_create"])
	if err != nil {
		t.Fatal(err)
	}
	var e catalog.Entry
	if err := json.Unmarshal(raw, &e); err != nil {
		t.Fatal(err)
	}
	if len(e.Params) != 2 || e.Params[0].Name != "clientId" || e.Params[0].Flag != "target-client-id" || string(e.Params[0].Example) != `9007199254740993` {
		t.Fatalf("path metadata lost: %+v", e.Params)
	}
	p := e.Params[1]
	if p.Name != "filter" || !p.Deprecated || p.Style != "form" || p.Explode == nil || *p.Explode || p.AllowReserved == nil || !*p.AllowReserved || len(p.Examples) == 0 {
		t.Fatalf("query metadata lost: %+v", p)
	}
	if len(e.BodyFields) != 2 {
		t.Fatalf("body mappings missing: %+v", e)
	}
	for _, field := range e.BodyFields {
		if !field.Required || len(field.Schema) == 0 {
			t.Fatalf("field metadata lost: %+v", field)
		}
		if field.Property == "orgId" && field.Flag != "body-org-id" {
			t.Fatalf("renamed body flag not mapped: %+v", field)
		}
	}
	if e.RequestExampleSource != "upstream" {
		t.Fatalf("example provenance lost: %+v", e)
	}
	if bodyExampleSource(map[string]interface{}{"schema": map[string]interface{}{"type": "string"}}) != "synthesized" {
		t.Fatal("synthesized example mislabeled")
	}
}
