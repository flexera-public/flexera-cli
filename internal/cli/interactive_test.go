package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/spf13/cobra"
)

type scriptPrompter struct {
	asks      []PromptField
	answer    string
	approvals int
	values    map[string]any
	options   []string
}

func (p *scriptPrompter) Ask(_ context.Context, f PromptField) (any, error) {
	p.asks = append(p.asks, f)
	if value, ok := p.values[f.Title]; ok {
		return value, nil
	}
	if f.Value != nil {
		return f.Value, nil
	}
	if f.Schema.Type != nil && f.Schema.Type.Is("boolean") {
		return false, nil
	}
	return "sample", nil
}
func (p *scriptPrompter) SelectFields(_ context.Context, title string, options, selected []string) ([]string, error) {
	return p.options, nil
}
func (p *scriptPrompter) Approve(context.Context) (string, error) {
	p.approvals++
	return p.answer, nil
}

func TestInteractiveTerminalGuards(t *testing.T) {
	for _, tc := range []struct {
		name     string
		in, err  bool
		body     string
		fragment string
	}{{"stdin", false, true, "", "terminal input and stderr"}, {"stderr", true, false, "", "terminal input and stderr"}, {"stdin-body", true, true, "@-", "cannot use --body @-"}} {
		t.Run(tc.name, func(t *testing.T) {
			input := strings.NewReader("")
			var stderr bytes.Buffer
			cmd := &cobra.Command{Use: "write"}
			cmd.SetIn(input)
			cmd.SetErr(&stderr)
			deps := &Deps{IsTerminal: func(value any) bool {
				if value == input {
					return tc.in
				}
				if value == &stderr {
					return tc.err
				}
				return false
			}}
			cmd.SetContext(WithDeps(context.Background(), deps))
			err := GuardInteractive(cmd, tc.body)
			var exit *ExitError
			if !errors.As(err, &exit) || exit.Code != 2 || !strings.Contains(err.Error(), tc.fragment) {
				t.Fatalf("guard: %v", err)
			}
		})
	}
}

func TestInteractivePrefillAndOptionalFields(t *testing.T) {
	var schema openapi3.Schema
	if err := json.Unmarshal([]byte(`{"type":"object","required":["name"],"properties":{"name":{"type":"string"},"optional":{"type":"boolean"}}}`), &schema); err != nil {
		t.Fatal(err)
	}
	p := &scriptPrompter{values: map[string]any{"Body.name": "edited", "Body.optional": false}, options: []string{"optional"}}
	cmd := &cobra.Command{Use: "write"}
	cmd.SetContext(context.Background())
	value, err := promptSchema(cmd, p, "Body", &schema, map[string]any{"name": "prefilled"}, 0, false)
	if err != nil {
		t.Fatal(err)
	}
	body := value.(map[string]any)
	if body["name"] != "edited" || body["optional"] != false || p.asks[0].Value != "prefilled" {
		t.Fatalf("prefill not editable: %v %+v", body, p.asks)
	}
}

func TestInteractiveArrayHonorsMinimumItems(t *testing.T) {
	for _, tc := range []struct {
		name     string
		minItems uint64
		wantLen  int
	}{
		{name: "empty allowed", minItems: 0, wantLen: 0},
		{name: "minimum required", minItems: 1, wantLen: 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			schema := openapi3.NewArraySchema().WithItems(openapi3.NewStringSchema())
			schema.MinItems = tc.minItems
			p := &scriptPrompter{}
			cmd := &cobra.Command{Use: "write"}
			cmd.SetContext(context.Background())
			value, err := promptSchema(cmd, p, "Items", schema, []any{}, 0, false)
			if err != nil {
				t.Fatal(err)
			}
			items, ok := value.([]any)
			if !ok || len(items) != tc.wantLen {
				t.Fatalf("items = %#v, want length %d", value, tc.wantLen)
			}
			if tc.minItems == 0 && (len(p.asks) != 1 || p.asks[0].Title != "Items: add another?") {
				t.Fatalf("empty array should not prompt for an item: %+v", p.asks)
			}
			if tc.minItems > 0 && (len(p.asks) != 2 || p.asks[0].Title != "Items[0]") {
				t.Fatalf("minimum item prompt missing: %+v", p.asks)
			}
		})
	}
}

func TestInteractiveApprovalAndDryRun(t *testing.T) {
	for _, tc := range []struct {
		name, answer    string
		dry, yes        bool
		code, approvals int
	}{{"yes", "yes", false, false, 0, 1}, {"cancel", "Yes", false, false, 1, 1}, {"auto", "", false, true, 0, 0}, {"dry", "", true, false, 0, 0}} {
		t.Run(tc.name, func(t *testing.T) {
			var out, stderr bytes.Buffer
			p := &scriptPrompter{answer: tc.answer}
			cmd := &cobra.Command{Use: "write"}
			cmd.SetOut(&out)
			cmd.SetErr(&stderr)
			cmd.SetContext(WithDeps(context.Background(), &Deps{Prompter: p}))
			plan := Plan{Command: "write", Method: "POST", Path: "/test", Body: json.RawMessage(`{"name":"sample"}`), Validation: &ValidationResult{Status: "ok"}}
			done, err := ConfirmInteractive(cmd, tc.dry, tc.yes, plan, Printer{})
			var exit *ExitError
			code := 0
			if errors.As(err, &exit) {
				code = exit.Code
			} else if err != nil {
				t.Fatal(err)
			}
			if code != tc.code || p.approvals != tc.approvals || done != tc.dry {
				t.Fatalf("approval mismatch done=%v err=%v prompts=%d", done, err, p.approvals)
			}
			if !strings.Contains(stderr.String(), "will perform") {
				t.Fatal("missing human plan")
			}
			if tc.dry && !json.Valid(out.Bytes()) {
				t.Fatal("missing JSON plan")
			}
			if !tc.dry && out.Len() != 0 {
				t.Fatal("approval polluted stdout")
			}
		})
	}
}

func TestUnsupportedSensitiveEditorNeverExports(t *testing.T) {
	var schema openapi3.Schema
	if err := json.Unmarshal([]byte(`{"type":"object","properties":{"password":{"type":"string"}}}`), &schema); err != nil {
		t.Fatal(err)
	}
	cmd := &cobra.Command{Use: "test"}
	cmd.SetContext(context.Background())
	cmd.SetOut(io.Discard)
	if _, err := editSchemaValue(cmd, &schema, map[string]any{}, false); err == nil || !strings.Contains(err.Error(), "use --body @file") {
		t.Fatalf("unsafe editor: %v", err)
	}
}

func TestPrivateInteractiveBodyExport(t *testing.T) {
	path := filepath.Join(t.TempDir(), "request.json")
	p := &scriptPrompter{values: map[string]any{"New private body file path": path}, options: []string{"Save body"}}
	var out, stderr bytes.Buffer
	cmd := &cobra.Command{Use: "write"}
	cmd.SetOut(&out)
	cmd.SetErr(&stderr)
	cmd.SetContext(WithDeps(context.Background(), &Deps{Prompter: p}))
	body := json.RawMessage(`{"password":"private-secret"}`)
	_, err := ConfirmInteractive(cmd, true, false, Plan{Method: "POST", Path: "/test", Body: body}, Printer{})
	if err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(path)
	if err != nil || string(data) != string(body) {
		t.Fatalf("body export changed %s %v", data, err)
	}
	info, err := os.Stat(path)
	if err != nil || info.Mode().Perm() != 0600 {
		t.Fatalf("unsafe private export mode %v %v", info, err)
	}
	if strings.Contains(out.String()+stderr.String(), "private-secret") {
		t.Fatal("private body leaked to preview")
	}
	if p.approvals != 0 {
		t.Fatal("dryrun asked for apply approval")
	}
}
