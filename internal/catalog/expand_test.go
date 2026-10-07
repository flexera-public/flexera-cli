package catalog

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestExpandDepthAndCycles(t *testing.T) {
	doc := testDocument()
	doc.Schemas["Widget"] = json.RawMessage(`{"type":"object","nullable":true,"required":["id"],"additionalProperties":false,"properties":{"$ref":{"type":"string"},"left":{"$ref":"#/components/schemas/Child"},"right":{"$ref":"#/components/schemas/Child"}},"example":{"$ref":"literal"}}`)
	doc.Schemas["Child"] = json.RawMessage(`{"allOf":[{"type":"object","properties":{"parent":{"$ref":"#/components/schemas/Widget"}}}]}`)
	data, _ := json.Marshal(doc)
	c, err := Parse(data)
	if err != nil {
		t.Fatal(err)
	}
	ref := json.RawMessage(`{"$ref":"#/components/schemas/Widget"}`)
	zero, err := c.Expand(ref, 0)
	if err != nil || string(zero) != string(ref) {
		t.Fatalf("depth0: %s %v", zero, err)
	}
	one, err := c.Expand(ref, 1)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Count(string(one), `#/components/schemas/Child`) != 2 {
		t.Fatalf("unexpected depth1: %s", one)
	}
	two, err := c.Expand(ref, 2)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Count(string(two), `"allOf"`) != 2 || strings.Count(string(two), `#/components/schemas/Widget`) != 2 {
		t.Fatalf("siblings/cycles not preserved: %s", two)
	}
	for _, fragment := range []string{`"nullable":true`, `"required":["id"]`, `"additionalProperties":false`, `"$ref":{"type":"string"}`, `"example":{"$ref":"literal"}`} {
		if !strings.Contains(string(two), fragment) {
			t.Errorf("lost %s in %s", fragment, two)
		}
	}
	if _, err := c.Expand(ref, -1); err == nil {
		t.Fatal("accepted negative depth")
	}
}

func TestValidExample(t *testing.T) {
	for _, tc := range []struct {
		name, schema, example string
		valid                 bool
	}{
		{"valid", `{"type":"object","required":["name"],"properties":{"name":{"type":"string","enum":["sample"]}}}`, `{"name":"sample"}`, true},
		{"missing", `{"type":"object","required":["name"]}`, `{}`, false},
		{"enum", `{"type":"string","enum":["one"]}`, `"two"`, false},
		{"type", `{"type":"string"}`, `42`, false},
		{"nullable", `{"type":"string","nullable":true}`, `null`, true},
		{"secret-key", `{"type":"object"}`, `{"client_secret":"never print"}`, false},
		{"writeOnly", `{"type":"object","properties":{"value":{"type":"string","writeOnly":true}}}`, `{"value":"never print"}`, false},
		{"password", `{"type":"string","format":"password"}`, `"never print"`, false},
		{"missing-example", `{"type":"object"}`, ``, false},
		{"large-id", `{"type":"integer"}`, `9007199254740993`, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			doc := testDocument()
			data, _ := json.Marshal(doc)
			c, err := Parse(data)
			if err != nil {
				t.Fatal(err)
			}
			e := Entry{RequestSchema: json.RawMessage(tc.schema), RequestExample: json.RawMessage(tc.example)}
			body, err := c.ValidExample(e)
			if (err == nil) != tc.valid {
				t.Fatalf("valid=%v err=%v", tc.valid, err)
			}
			if tc.valid && string(body) != tc.example {
				t.Fatalf("body changed: %s", body)
			}
			if err != nil && strings.Contains(err.Error(), "never print") {
				t.Fatal("secret leaked in error")
			}
		})
	}
}
