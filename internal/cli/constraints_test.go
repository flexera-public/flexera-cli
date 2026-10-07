package cli

import (
	"encoding/json"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestStrictFormatsAndLargeNumericConstraints(t *testing.T) {
	for _, tc := range []struct {
		schema, value string
		valid         bool
	}{
		{`{"type":"integer","maximum":9007199254740992}`, `9007199254740993`, false},
		{`{"type":"integer"}`, `9007199254740992.1`, false},
		{`{"type":"integer","minimum":9007199254740992,"exclusiveMinimum":true}`, `9007199254740993`, true},
		{`{"type":"string","format":"uuid"}`, `"bad"`, false},
		{`{"type":"string","format":"date"}`, `"2025-02-30"`, false},
		{`{"type":"string","format":"email"}`, `"bad"`, false},
		{`{"type":"string","format":"uri"}`, `"relative/path"`, false},
	} {
		var schema openapi3.Schema
		if err := json.Unmarshal([]byte(tc.schema), &schema); err != nil {
			t.Fatal(err)
		}
		value, err := ParseRequestJSON([]byte(tc.value))
		if err != nil {
			t.Fatal(err)
		}
		if err := validateSchemaValue(&schema, value); (err == nil) != tc.valid {
			t.Errorf("schema=%s value=%s valid=%v err=%v", tc.schema, tc.value, tc.valid, err)
		}
	}
}
