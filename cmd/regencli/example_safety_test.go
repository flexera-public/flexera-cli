package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/flexera-public/flexera-cli/internal/catalog"
)

func TestBuildCatalogExampleSafety(t *testing.T) {
	for _, tc := range []struct {
		name, schema, components, example string
		valid                             bool
	}{
		{"valid", `{"type":"object","required":["id"],"additionalProperties":false,"properties":{"id":{"type":"integer"}},"example":{"id":9007199254740993}}`, `{}`, `{"id":9007199254740993}`, true},
		{"invalid-required", `{"type":"object","required":["id"],"properties":{"id":{"type":"integer"}}}`, `{}`, `{}`, false},
		{"invalid-enum", `{"type":"string","enum":["allowed"],"example":"other"}`, `{}`, `"other"`, false},
		{"invalid-additional", `{"type":"object","additionalProperties":false}`, `{}`, `{"extra":1}`, false},
		{"invalid-ref", `{"$ref":"#/components/schemas/Value"}`, `{"Value":{"type":"string","minLength":5}}`, `"x"`, false},
		{"name", `{"type":"object","properties":{"client_secret":{"type":"string","example":"hidden-name","default":"hidden-name","examples":["hidden-name"]}},"example":{"client_secret":"hidden-name"}}`, `{}`, `{"client_secret":"hidden-name"}`, false},
		{"authorization", `{"type":"object","example":{"Authorization":"hidden-auth"}}`, `{}`, `{"Authorization":"hidden-auth"}`, false},
		{"additional-ref", `{"type":"object","additionalProperties":{"$ref":"#/components/schemas/Value"},"example":{"custom":"hidden-additional"},"default":{"custom":"hidden-additional"},"examples":[{"custom":"hidden-additional"}]}`, `{"Value":{"type":"string","writeOnly":true,"example":"hidden-additional","default":"hidden-additional","examples":["hidden-additional"]}}`, `{"custom":"hidden-additional"}`, false},
		{"allOf", `{"allOf":[{"$ref":"#/components/schemas/Value"}],"example":{"value":"hidden-allof"}}`, `{"Value":{"type":"object","properties":{"value":{"type":"string","format":"password","example":"hidden-allof"}}}}`, `{"value":"hidden-allof"}`, false},
		{"anyOf", `{"anyOf":[{"type":"string"},{"$ref":"#/components/schemas/Value"}],"example":"hidden-anyof"}`, `{"Value":{"type":"string","format":"password","default":"hidden-anyof"}}`, `"hidden-anyof"`, false},
		{"oneOf", `{"oneOf":[{"$ref":"#/components/schemas/Value"}],"examples":["hidden-oneof"]}`, `{"Value":{"type":"string","writeOnly":true}}`, `"hidden-oneof"`, false},
		{"array", `{"type":"array","items":{"$ref":"#/components/schemas/Value"},"example":["hidden-array"]}`, `{"Value":{"type":"string","format":"password","example":"hidden-array"}}`, `["hidden-array"]`, false},
		{"named-ref", `{"type":"object","properties":{"api-key":{"$ref":"#/components/schemas/Value"}},"example":{"api-key":"hidden-ref"}}`, `{"Value":{"$ref":"#/components/schemas/Next","example":"hidden-ref"},"Next":{"type":"string","default":"hidden-ref"}}`, `{"api-key":"hidden-ref"}`, false},
		{"escaped-ref", `{"type":"object","additionalProperties":{"$ref":"#/components/schemas/A~1B~0C"},"example":{"custom":"hidden-escaped"}}`, `{"A/B~C":{"allOf":[{"$ref":"#/components/schemas/Value"}]},"Value":{"type":"string","format":"password","example":"hidden-escaped"}}`, `{"custom":"hidden-escaped"}`, false},
		{"recursive", `{"$ref":"#/components/schemas/Value"}`, `{"Value":{"type":"object","properties":{"child":{"$ref":"#/components/schemas/Value"},"value":{"type":"string","writeOnly":true}},"example":{"child":{"value":"hidden-cycle"}}}}`, `{"child":{"value":"hidden-cycle"}}`, false},
		{"null-writeonly", `{"type":"string","nullable":true,"writeOnly":true,"example":null}`, `{}`, `null`, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			stage := t.TempDir()
			e := catalog.Entry{OperationID: "test", Command: []string{"test", "create"}, Method: "POST", Path: "/test", ResponseEnvelope: "none", RequestSchema: json.RawMessage(`{"$ref":"#/components/schemas/Root"}`), RequestExample: json.RawMessage(tc.example)}
			metadata, err := json.Marshal(map[string]catalog.Entry{"test": e})
			if err != nil {
				t.Fatal(err)
			}
			writeTestFile(t, filepath.Join(stage, "test", "metadata.json"), string(metadata))
			var components map[string]json.RawMessage
			if err := json.Unmarshal([]byte(tc.components), &components); err != nil {
				t.Fatal(err)
			}
			components["Root"] = json.RawMessage(tc.schema)
			raw, _ := json.Marshal(components)
			spec := fmt.Appendf(nil, `{"paths":{"/test":{"post":{"operationId":"test"}}},"components":{"schemas":%s}}`, raw)
			first, err := buildCatalog(stage, []genTag{{Pkg: "test", Cmd: "test"}}, spec)
			if err != nil {
				t.Fatal(err)
			}
			second, err := buildCatalog(stage, []genTag{{Pkg: "test", Cmd: "test"}}, spec)
			if err != nil || !bytes.Equal(first, second) {
				t.Fatalf("nondeterministic catalog: %v", err)
			}
			if bytes.Contains(first, []byte("hidden-")) {
				t.Fatal("secret sample survived publication")
			}
			var doc catalog.Document
			if err := json.Unmarshal(first, &doc); err != nil {
				t.Fatal(err)
			}
			got := doc.Entries[0].RequestExample
			if (len(got) > 0) != tc.valid {
				t.Fatalf("request example retained=%t, want %t", len(got) > 0, tc.valid)
			}
			var compact bytes.Buffer
			if tc.valid {
				if err := json.Compact(&compact, got); err != nil {
					t.Fatal(err)
				}
			}
			if tc.valid && compact.String() != tc.example {
				t.Fatalf("example changed: %s", got)
			}
			// Compare schemas with only sample annotations stripped. No field or
			// structural constraint may disappear from the published schema.
			for name, original := range components {
				before, _ := decodeSampleJSON(original)
				after, _ := decodeSampleJSON(doc.Schemas[name])
				stripSafetySamples(before)
				stripSafetySamples(after)
				b, _ := json.Marshal(before)
				a, _ := json.Marshal(after)
				if !bytes.Equal(b, a) {
					t.Fatalf("schema constraints changed for %s: %s -> %s", name, b, a)
				}
			}
		})
	}
}

func stripSafetySamples(value any) {
	switch v := value.(type) {
	case map[string]any:
		for key, child := range v {
			if key == "example" || key == "default" || key == "examples" {
				delete(v, key)
			} else {
				stripSafetySamples(child)
			}
		}
	case []any:
		for _, child := range v {
			stripSafetySamples(child)
		}
	}
}

func TestCatalogInlineSamplesAndPrecision(t *testing.T) {
	schema := json.RawMessage(`{"type":"object","required":["id"],"properties":{"id":{"type":"integer","minimum":9007199254740993,"example":9007199254740993,"default":9007199254740993,"examples":[9007199254740993]},"private_key":{"type":"string","default":"hidden-inline"}},"example":{"private_key":"hidden-inline"}}`)
	doc := catalog.Document{Schemas: map[string]json.RawMessage{}, Entries: []catalog.Entry{{OperationID: "inline", Command: []string{"test"}, Method: "POST", Path: "/test", ResponseEnvelope: "none", RequestSchema: schema, ResponseSchema: schema, Params: []catalog.Param{{Flag: "authorization", Source: "flag", In: "query", Type: "string", Schema: json.RawMessage(`{"type":"string","example":"hidden-param"}`)}}}}}
	if err := sanitizeCatalogExamples(&doc); err != nil {
		t.Fatal(err)
	}
	data, _ := json.Marshal(doc)
	if bytes.Contains(data, []byte("hidden-")) {
		t.Fatal("inline secret leaked")
	}
	if strings.Count(string(data), "9007199254740993") != 8 {
		t.Fatalf("numeric samples or constraints changed: %s", data)
	}
}

func TestOriginalExampleValidationSensitiveSchemas(t *testing.T) {
	for _, schema := range []string{
		`{"type":"object","additionalProperties":{"$ref":"#/components/schemas/Value"}}`,
		`{"allOf":[{"type":"object","additionalProperties":{"$ref":"#/components/schemas/Value"}}]}`,
		`{"type":"object","properties":{"child":{"$ref":"#/components/schemas/Root"},"custom":{"$ref":"#/components/schemas/Value"}}}`,
	} {
		doc := catalog.Document{Entries: []catalog.Entry{}, Schemas: map[string]json.RawMessage{"Root": json.RawMessage(schema), "Value": json.RawMessage(`{"type":"string","writeOnly":true}`)}}
		data, _ := json.Marshal(doc)
		index, err := catalog.Parse(data)
		if err != nil {
			t.Fatal(err)
		}
		for _, example := range []string{`{"custom":"hidden-value"}`, `{"child":{"custom":"hidden-value"}}`} {
			_, err := index.ValidExample(catalog.Entry{RequestSchema: json.RawMessage(`{"$ref":"#/components/schemas/Root"}`), RequestExample: json.RawMessage(example)})
			if err == nil {
				t.Fatalf("accepted secret/invalid example: %s", schema)
			}
			if strings.Contains(err.Error(), "hidden-value") {
				t.Fatal("secret leaked in error")
			}
		}
	}
}

func TestCatalogSampleReportAndSchemaFields(t *testing.T) {
	makeDoc := func() catalog.Document {
		return catalog.Document{Schemas: map[string]json.RawMessage{"Root": json.RawMessage(`{"type":"object","properties":{"example":{"type":"string","example":"public"},"default":{"type":"number","default":1.234567890123456789},"examples":{"type":"string"},"password":{"type":"string","example":"hidden-report"}},"example":{"password":"hidden-report"}}`)}, Entries: []catalog.Entry{{OperationID: "test", Command: []string{"test"}, Method: "POST", Path: "/test", ResponseEnvelope: "none", RequestSchema: json.RawMessage(`{"$ref":"#/components/schemas/Root"}`), RequestExample: json.RawMessage(`{"password":"hidden-report"}`)}}}
	}
	var previous []byte
	var previousReport string
	for i := 0; i < 2; i++ {
		doc := makeDoc()
		before, _ := json.MarshalIndent(doc, "", "  ")
		r, w, err := os.Pipe()
		if err != nil {
			t.Fatal(err)
		}
		stderr := os.Stderr
		os.Stderr = w
		err = sanitizeCatalogExamples(&doc)
		os.Stderr = stderr
		_ = w.Close()
		report, readErr := io.ReadAll(r)
		_ = r.Close()
		if err != nil || readErr != nil {
			t.Fatalf("sanitize=%v read=%v", err, readErr)
		}
		after, _ := json.MarshalIndent(doc, "", "  ")
		want := fmt.Sprintf("catalog samples: omitted 1 request examples, removed 2 schema samples; bytes %d -> %d\n", len(before)+1, len(after)+1)
		if string(report) != want {
			t.Fatalf("report=%q want=%q", report, want)
		}
		if i > 0 && (!bytes.Equal(previous, after) || previousReport != string(report)) {
			t.Fatal("publication or measurement changed between identical builds")
		}
		previous, previousReport = after, string(report)
		var root struct {
			Properties map[string]json.RawMessage `json:"properties"`
		}
		if err := json.Unmarshal(doc.Schemas["Root"], &root); err != nil {
			t.Fatal(err)
		}
		for _, key := range []string{"example", "default", "examples", "password"} {
			if len(root.Properties[key]) == 0 {
				t.Fatalf("removed structural property %q", key)
			}
		}
		if !bytes.Contains(after, []byte(`"example": "public"`)) || !bytes.Contains(after, []byte(`1.234567890123456789`)) {
			t.Fatal("safe examples or lossless numeric defaults removed")
		}
	}
}
