package cli

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/flexera-public/flexera-cli/internal/catalog"
	"github.com/spf13/cobra"
)

func TestParameterPresenceAndConstraints(t *testing.T) {
	for _, tc := range []struct {
		name, kind, schema string
		input              ParamInput
		skip, valid        bool
	}{
		{"missing-integer", "integer", `{"type":"integer"}`, ParamInput{}, false, false},
		{"zero", "integer", `{"type":"integer"}`, ParamInput{Value: 0, Present: true}, false, true},
		{"false", "boolean", `{"type":"boolean"}`, ParamInput{Value: false, Present: true}, false, true},
		{"missing-bool", "boolean", `{"type":"boolean"}`, ParamInput{}, true, false},
		{"minimum", "integer", `{"type":"integer","minimum":1}`, ParamInput{Value: 0, Present: true}, false, false},
		{"skip-min", "integer", `{"type":"integer","minimum":1}`, ParamInput{Value: 0, Present: true}, true, true},
		{"enum", "string", `{"type":"string","enum":["cost"]}`, ParamInput{Value: "usage", Present: true}, false, false},
		{"blank", "string", `{"type":"string"}`, ParamInput{Value: " ", Present: true}, true, false},
		{"pattern", "string", `{"type":"string","pattern":"^[a-z]+$"}`, ParamInput{Value: "123", Present: true}, false, false},
		{"array", "array", `{"type":"array","minItems":1,"items":{"type":"string"}}`, ParamInput{Value: []string{"value"}, Present: true}, false, true},
		{"empty-array", "array", `{"type":"array","minItems":1,"items":{"type":"string"}}`, ParamInput{Value: []string{}, Present: true}, false, false},
		{"date-time", "string", `{"type":"string","format":"date-time"}`, ParamInput{Value: "not-a-date", Present: true}, false, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			index, entry := validationFixture(t, `{"type":"object"}`)
			entry.Params = []catalog.Param{{Flag: "value", In: "query", Source: "flag", Required: true, Type: tc.kind, Schema: json.RawMessage(tc.schema)}}
			effective, err := ValidateParams(index, entry, map[string]ParamInput{"value": tc.input}, tc.skip)
			if (err == nil) != tc.valid {
				t.Fatalf("valid=%v err=%v", tc.valid, err)
			}
			if tc.valid && effective["value"] == nil {
				t.Fatal("supplied value omitted")
			}
			if err != nil {
				var exit *ExitError
				if !errors.As(err, &exit) || exit.Code != 2 {
					t.Fatalf("not usage error: %v", err)
				}
			}
		})
	}
}

func TestCommandParamsConfigOrgAndRequiredNumeric(t *testing.T) {
	root := &cobra.Command{Use: "root"}
	RegisterPersistentFlags(root.PersistentFlags())
	cmd := &cobra.Command{Use: "budget"}
	root.AddCommand(cmd)
	deps := &Deps{}
	deps.Config.OrgID = 123
	cmd.SetContext(WithDeps(context.Background(), deps))
	values, err := ValidateCommandParams(cmd, "Budget_Budget_create")
	if err != nil || values["org-id"] != 123 {
		t.Fatalf("config org missing %v %v", values, err)
	}
	deps.Config.OrgID = 0
	if _, err := ValidateCommandParams(cmd, "Budget_Budget_create"); err == nil {
		t.Fatal("missing config org accepted")
	}
	deps.OrgIDPresent = true
	if _, err := ValidateCommandParams(cmd, "Budget_Budget_create"); err == nil {
		t.Fatal("schema minimum for org zero ignored")
	}
	if err := root.PersistentFlags().Set(FlagNoValidate, "true"); err != nil {
		t.Fatal(err)
	}
	values, err = ValidateCommandParams(cmd, "Budget_Budget_create")
	if err != nil || values["org-id"] != 0 {
		t.Fatalf("configured zero presence lost under skip: %v %v", values, err)
	}
}
