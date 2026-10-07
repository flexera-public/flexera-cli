package catalog

import (
	"encoding/json"
	"strings"
	"testing"
)

func testDocument() Document {
	return Document{Entries: []Entry{{OperationID: "Widget_get", Command: []string{"widget", "get"}, Method: "GET", Path: "/widgets/{id}", ResponseEnvelope: "none", ResponseSchema: json.RawMessage(`{"$ref":"#/components/schemas/Widget"}`)}}, Schemas: map[string]json.RawMessage{"Widget": json.RawMessage(`{"type":"object","nullable":true,"properties":{"child":{"$ref":"#/components/schemas/Widget"}}}`)}}
}

func TestParseLookupAndCopies(t *testing.T) {
	data, _ := json.Marshal(testDocument())
	c, err := Parse(data)
	if err != nil {
		t.Fatal(err)
	}
	e, ok := c.Lookup("Widget_get")
	if !ok {
		t.Fatal("entry missing")
	}
	e.Command[0] = "modified"
	if c.All()[0].Command[0] != "widget" {
		t.Fatal("shared catalog mutated")
	}
	schema, ok := c.Schema("#/components/schemas/Widget")
	if !ok || !strings.Contains(string(schema), `"nullable":true`) {
		t.Fatal(string(schema))
	}
	schema[0] = 'x'
	again, _ := c.Schema("#/components/schemas/Widget")
	if again[0] != '{' {
		t.Fatal("shared schema mutated")
	}
	if _, ok := c.Lookup("missing"); ok {
		t.Fatal("unexpected lookup")
	}
}

func TestParseRejectsInvalidDocuments(t *testing.T) {
	for _, scenario := range []string{"duplicate-id", "duplicate-command", "dangling-ref", "nested-ref", "envelope", "method"} {
		t.Run(scenario, func(t *testing.T) {
			doc := testDocument()
			switch scenario {
			case "duplicate-id":
				doc.Entries = append(doc.Entries, doc.Entries[0])
			case "duplicate-command":
				e := doc.Entries[0]
				e.OperationID = "other"
				doc.Entries = append(doc.Entries, e)
			case "dangling-ref":
				delete(doc.Schemas, "Widget")
			case "nested-ref":
				doc.Schemas["Widget"] = json.RawMessage(`{"allOf":[{"$ref":"#/components/schemas/Missing"}]}`)
			case "envelope":
				doc.Entries[0].ResponseEnvelope = "values"
			case "method":
				doc.Entries[0].Method = "get"
			}
			data, _ := json.Marshal(doc)
			if _, err := Parse(data); err == nil {
				t.Fatal("invalid document accepted")
			}
		})
	}
	for _, raw := range []string{`{}`, `null`, `{"entries":[],"schemas":null}`, `invalid`} {
		if _, err := Parse([]byte(raw)); err == nil {
			t.Fatalf("accepted %s", raw)
		}
	}
}

func TestSchemaPropertyNamedRef(t *testing.T) {
	doc := testDocument()
	doc.Schemas["Widget"] = json.RawMessage(`{"type":"object","properties":{"$ref":{"type":"string"},"child":{"$ref":"#/components/schemas/Widget"}}}`)
	data, _ := json.Marshal(doc)
	if _, err := Parse(data); err != nil {
		t.Fatal(err)
	}
}

func TestPublishedCatalogLoads(t *testing.T) {
	c, err := Load()
	if err != nil || c == nil || len(c.All()) == 0 {
		t.Fatalf("published catalog failed to load: %v", err)
	}
	if len(artifact) > 5*1024*1024 {
		t.Fatalf("catalog exceeds approved budget: %d", len(artifact))
	}
}
