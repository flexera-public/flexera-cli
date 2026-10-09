package cli

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/flexera-public/flexera-cli/internal/catalog"
	"github.com/spf13/cobra"
)

func TestSchemaHelp(t *testing.T) {
	parts, err := schemaHelp(json.RawMessage(`{
		"type":"string","description":"Filter expression","enum":["a","b"],
		"format":"date","minLength":1,"maxLength":10,"pattern":"^[ab]$",
		"default":"a","example":"b"
	}`), nil, true, "filter")
	if err != nil {
		t.Fatal(err)
	}
	text := strings.Join(parts, "; ")
	for _, want := range []string{"required by API", "Filter expression", `enum: ["a","b"]`, "format: date", "minLength: 1", `pattern: "^[ab]$"`, `API default: "a"`, `illustrative example: "b"`} {
		if !strings.Contains(text, want) {
			t.Fatalf("missing %q: %s", want, text)
		}
	}
	parts, err = schemaHelp(json.RawMessage(`{"type":"integer","example":9007199254740993,"default":9007199254740993}`), nil, false, "id")
	if err != nil || !strings.Contains(strings.Join(parts, "; "), "9007199254740993") {
		t.Fatalf("numeric precision lost: %v, %v", parts, err)
	}
}

func TestSchemaHelpSensitiveAndBounded(t *testing.T) {
	for _, tc := range []struct{ name, schema string }{
		{"password", `{"type":"string","default":"hidden","example":"hidden","enum":["hidden"]}`},
		{"value", `{"type":"string","writeOnly":true,"default":"hidden","example":"hidden"}`},
		{"value", `{"type":"string","format":"password","example":"hidden"}`},
	} {
		parts, err := schemaHelp(json.RawMessage(tc.schema), json.RawMessage(`"hidden"`), true, tc.name)
		if err != nil || strings.Contains(strings.Join(parts, "; "), "hidden") {
			t.Fatalf("sensitive help: %v, %v", parts, err)
		}
	}
	if text := shortHelpValue([]byte(strings.Repeat("x", 200))); !strings.HasSuffix(text, "... (see cli schema)") {
		t.Fatalf("unbounded help: %s", text)
	}
	if _, err := schemaHelp(json.RawMessage(`invalid`), nil, false, "x"); err == nil {
		t.Fatal("invalid schema silently ignored")
	}
}

func TestNamedExampleHelp(t *testing.T) {
	value, err := parameterHelpExample(nil, json.RawMessage(`{"z":{"value":"last"},"a":{"value":9007199254740993}}`))
	if err != nil || string(value) != "9007199254740993" {
		t.Fatalf("named examples not deterministic/lossless: %s, %v", value, err)
	}

	value, err = parameterHelpExample(json.RawMessage(`"explicit"`), json.RawMessage(`{"a":{"value":"named"}}`))
	if err != nil || string(value) != `"explicit"` {
		t.Fatalf("example precedence: %s, %v", value, err)
	}
	if _, err := parameterHelpExample(nil, json.RawMessage(`invalid`)); err == nil {
		t.Fatal("invalid named examples silently ignored")
	}
	parts, err := schemaHelp(json.RawMessage(`{"type":"array","items":{"type":"string","enum":["a","b"]}}`), nil, false, "names")
	if err != nil || !strings.Contains(strings.Join(parts, "; "), `items.enum: ["a","b"]`) {
		t.Fatalf("array enum guidance missing: %v, %v", parts, err)
	}
}

func TestOperationHelpText(t *testing.T) {
	if got := operationHelpText("  Guidance \n   \nNext line\t \n"); got != "Guidance\n\nNext line" {
		t.Fatalf("unstable help whitespace: %q", got)
	}
}

func TestEnrichMetadataHelpPreservesFlagBehavior(t *testing.T) {
	index, err := catalog.Load()
	if err != nil {
		t.Fatal(err)
	}
	var entry catalog.Entry
	for _, e := range index.All() {
		for _, p := range e.Params {
			if p.Flag == "filter" && strings.Contains(p.Description, "OData-style") {
				entry = e
				break
			}
		}
		if entry.OperationID != "" {
			break
		}
	}
	if entry.OperationID == "" {
		t.Fatal("OData filter fixture missing")
	}
	root := &cobra.Command{Use: "test"}
	leaf := &cobra.Command{Use: "list", Short: entry.Summary, Annotations: map[string]string{"flexera.operationId": entry.OperationID}}
	var filter string
	leaf.Flags().StringVar(&filter, "filter", "", "filter (query)")
	root.AddCommand(leaf)
	if err := EnrichMetadataHelp(root); err != nil {
		t.Fatal(err)
	}
	flag := leaf.Flags().Lookup("filter")
	if !strings.Contains(flag.Usage, "OData-style") || leaf.Long == "" {
		t.Fatalf("guidance missing: %s, %s", flag.Usage, leaf.Long)
	}
	if filter != "" || flag.DefValue != "" || flag.Changed || len(flag.Annotations[cobra.BashCompOneRequiredFlag]) != 0 {
		t.Fatalf("help changed flag behavior: %+v", flag)
	}
	root.AddCommand(&cobra.Command{Use: "missing", Annotations: map[string]string{"flexera.operationId": "Missing"}})
	if err := EnrichMetadataHelp(root); err == nil {
		t.Fatal("missing catalog entry silently ignored")
	}
}
