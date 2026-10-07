package cli

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"github.com/flexera-public/flexera-cli/internal/catalog"
	flexera "github.com/flexera-public/unified-go-client"
)

func validationFixture(t *testing.T, schema string) (*catalog.Catalog, catalog.Entry) {
	t.Helper()
	entry := catalog.Entry{OperationID: "Test_create", Command: []string{"test", "create"}, Method: "POST", Path: "/test", ResponseEnvelope: "none", RequestSchema: json.RawMessage(schema)}
	doc := catalog.Document{Entries: []catalog.Entry{entry}, Schemas: map[string]json.RawMessage{}}
	data, err := json.Marshal(doc)
	if err != nil {
		t.Fatal(err)
	}
	index, err := catalog.Parse(data)
	if err != nil {
		t.Fatal(err)
	}
	return index, entry
}

func TestRawRequestValidation(t *testing.T) {
	for _, tc := range []struct {
		name, schema, body string
		valid              bool
	}{
		{"required", `{"type":"object","required":["name"],"properties":{"name":{"type":"string"}}}`, `{}`, false},
		{"enum", `{"type":"string","enum":["cost"]}`, `"usage"`, false},
		{"type", `{"type":"integer"}`, `"bad"`, false},
		{"forbidden-extra", `{"type":"object","additionalProperties":false}`, `{"extra":1}`, false},
		{"allowed-extra", `{"type":"object","additionalProperties":true}`, `{"extra":1}`, true},
		{"nullable", `{"type":"string","nullable":true}`, `null`, true},
		{"allOf", `{"allOf":[{"type":"object","required":["name"]},{"type":"object","properties":{"name":{"type":"string"}}}]}`, `{"name":"test"}`, true},
		{"oneOf", `{"oneOf":[{"type":"string"},{"type":"integer"}]}`, `12`, true},
		{"zero", `{"type":"integer"}`, `0`, true},
		{"false", `{"type":"boolean"}`, `false`, true},
		{"large", `{"type":"integer"}`, `9007199254740993`, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			index, entry := validationFixture(t, tc.schema)
			result, err := ValidateRequestJSON(index, entry, []byte(tc.body), false)
			if (err == nil) != tc.valid {
				t.Fatalf("valid=%v err=%v", tc.valid, err)
			}
			if tc.valid && result.Status != "ok" {
				t.Fatal(result)
			}
			if !tc.valid {
				var typed *ValidationError
				if !errors.As(err, &typed) || len(typed.Details) == 0 || typed.Schema != "flexera-cli cli schema test create" {
					t.Fatalf("missing details %v", err)
				}
			}
		})
	}
}

func TestValidationSkippedButSyntaxRequired(t *testing.T) {
	index, entry := validationFixture(t, `{"type":"integer"}`)
	result, err := ValidateRequestJSON(index, entry, []byte(`"not integer"`), true)
	if err != nil || result.Status != "skipped" {
		t.Fatalf("skip: %v %v", result, err)
	}
	for _, raw := range []string{``, `invalid`, `{} {}`} {
		if _, err := ValidateRequestJSON(index, entry, []byte(raw), true); err == nil {
			t.Fatalf("skip accepted syntax %q", raw)
		}
	}
}

func TestTypedDecodeRefusesDataLoss(t *testing.T) {
	index, entry := validationFixture(t, `{"type":"object","additionalProperties":true}`)
	for _, skip := range []bool{false, true} {
		var body struct {
			Name string `json:"name"`
		}
		_, err := DecodeRequestBody(index, entry, []byte(`{"name":"test","unknown":"private-secret-value"}`), &body, skip)
		var typed *ValidationError
		if !errors.As(err, &typed) || typed.Details[0].Path != "/unknown" {
			t.Fatalf("lost unknown accepted: %v", err)
		}
		encoded, _ := json.Marshal(errorProperties(err))
		if strings.Contains(string(encoded), "private-secret-value") {
			t.Fatal("secret leaked")
		}
	}
	var preserved map[string]any
	if _, err := DecodeRequestBody(index, entry, []byte(`{"extra":1}`), &preserved, false); err != nil {
		t.Fatal(err)
	}
	var rounded struct {
		ID float64 `json:"id"`
	}
	if _, err := DecodeRequestBody(index, entry, []byte(`{"id":9007199254740993}`), &rounded, true); err == nil {
		t.Fatal("rounded ID accepted")
	}
	var exact struct {
		ID int64 `json:"id"`
	}
	if _, err := DecodeRequestBody(index, entry, []byte(`{"id":9007199254740993}`), &exact, false); err != nil {
		t.Fatal(err)
	}
	var omitted struct {
		Name *string `json:"name,omitempty"`
	}
	if _, err := DecodeRequestBody(index, entry, []byte(`{"name":null}`), &omitted, true); err == nil {
		t.Fatal("explicit null loss accepted")
	}
}

func TestRealBudgetRequestValidation(t *testing.T) {
	index, err := catalog.Load()
	if err != nil {
		t.Fatal(err)
	}
	var entry catalog.Entry
	for _, e := range index.All() {
		if strings.Join(e.Command, " ") == "budget create" {
			entry = e
			break
		}
	}
	var body flexera.BudgetBudgetCreateJSONRequestBody
	result, err := DecodeRequestBody(index, entry, entry.RequestExample, &body, false)
	if err != nil || result.Status != "ok" {
		t.Fatalf("real budget example: %v %v", result, err)
	}
}

func TestNestedDataLossAndPointerEscaping(t *testing.T) {
	before, err := ParseRequestJSON([]byte(`{"rows":[{"a/b~c":1}]}`))
	if err != nil {
		t.Fatal(err)
	}
	after, err := ParseRequestJSON([]byte(`{"rows":[{}]}`))
	if err != nil {
		t.Fatal(err)
	}
	paths := changedInputs(before, after, "")
	if len(paths) != 1 || paths[0] != "/rows/0/a~1b~0c" {
		t.Fatalf("invalid JSON pointer: %v", paths)
	}
	if paths := changedInputs("string", map[string]any{}, ""); len(paths) != 1 {
		t.Fatalf("incompatible primitive accepted: %v", paths)
	}
}

func TestSchemaErrorsDoNotExposeSuppliedSecretValues(t *testing.T) {
	index, entry := validationFixture(t, `{"type":"object","properties":{"password":{"type":"string","minLength":100}}}`)
	_, err := ValidateRequestJSON(index, entry, []byte(`{"password":"private-secret-value"}`), false)
	if err == nil {
		t.Fatal("invalid value accepted")
	}
	encoded, marshalErr := json.Marshal(errorProperties(err))
	if marshalErr != nil || strings.Contains(string(encoded), "private-secret-value") {
		t.Fatalf("unsafe diagnostics: %s %v", encoded, marshalErr)
	}
}
