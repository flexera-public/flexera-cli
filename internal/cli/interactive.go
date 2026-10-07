package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sort"
	"strings"

	"github.com/flexera-public/flexera-cli/internal/catalog"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

type PromptField struct {
	Title  string
	Schema *openapi3.Schema
	Value  any
	Secret bool
}
type Prompter interface {
	Ask(context.Context, PromptField) (any, error)
	SelectFields(context.Context, string, []string, []string) ([]string, error)
	Approve(context.Context) (string, error)
}

func terminalStream(stream any) bool {
	file, ok := stream.(*os.File)
	return ok && term.IsTerminal(int(file.Fd()))
}

// GuardInteractive checks actual prompt streams, never stdout. Tests inject a
// terminal predicate through Deps; production never bypasses these checks.
func GuardInteractive(cmd *cobra.Command, body string) error {
	if strings.TrimSpace(body) == "@-" {
		return Exit(2, fmt.Errorf("interactive mode cannot use --body @-; use inline JSON or @file"))
	}
	isTerminal := terminalStream
	if deps := DepsFrom(cmd.Context()); deps != nil && deps.IsTerminal != nil {
		isTerminal = deps.IsTerminal
	}
	if !isTerminal(cmd.InOrStdin()) || !isTerminal(cmd.ErrOrStderr()) {
		return Exit(2, fmt.Errorf("interactive mode requires terminal input and stderr"))
	}
	return nil
}

func interactivePrompter(cmd *cobra.Command) Prompter {
	if deps := DepsFrom(cmd.Context()); deps != nil && deps.Prompter != nil {
		return deps.Prompter
	}
	return &HuhPrompter{Input: cmd.InOrStdin(), Output: cmd.ErrOrStderr(), Accessible: os.Getenv("FLEXERA_CLI_ACCESSIBLE") == "1"}
}

// GatherInteractiveParams updates the real flags/config so the same final
// values are subsequently validated, displayed, and passed to the SDK.
func GatherInteractiveParams(cmd *cobra.Command, operationID string) error {
	index, err := catalog.Load()
	if err != nil {
		return err
	}
	entry, ok := index.Lookup(operationID)
	if !ok {
		return Exit(2, fmt.Errorf("unknown interactive operation %s", operationID))
	}
	deps := DepsFrom(cmd.Context())
	prompter := interactivePrompter(cmd)
	for _, p := range entry.Params {
		present := false
		if p.Source == "config" {
			present = deps.OrgIDPresent || deps.Config.OrgID != 0
		} else if flag := cmd.Flags().Lookup(p.Flag); flag != nil {
			present = flag.Changed
		}
		if !p.Required || present {
			continue
		}
		schema, err := index.RequestSchema(p.Schema)
		if err != nil {
			return err
		}
		value, err := prompter.Ask(cmd.Context(), PromptField{Title: "Required parameter: " + p.Flag, Schema: schema, Secret: sensitiveName(p.Flag)})
		if err != nil {
			return err
		}
		encoded, err := json.Marshal(value)
		if err != nil {
			return err
		}
		text := string(encoded)
		if s, ok := value.(string); ok {
			text = s
		}
		if err := cmd.Flags().Set(p.Flag, text); err != nil {
			return Exit(2, fmt.Errorf("cannot set interactive --%s", p.Flag))
		}
		if p.Source == "config" {
			org, err := cmd.Flags().GetInt(p.Flag)
			if err != nil {
				return err
			}
			deps.Config.OrgID = org
			deps.OrgIDPresent = true
		}
	}
	return nil
}

func GatherInteractiveBody(cmd *cobra.Command, operationID string, raw []byte) (json.RawMessage, error) {
	index, err := catalog.Load()
	if err != nil {
		return nil, err
	}
	entry, ok := index.Lookup(operationID)
	if !ok {
		return nil, Exit(2, fmt.Errorf("unknown interactive operation %s", operationID))
	}
	schema, err := index.RequestSchema(entry.RequestSchema)
	if err != nil {
		return nil, err
	}
	var value any = map[string]any{}
	if len(raw) > 0 {
		value, err = ParseRequestJSON(raw)
		if err != nil {
			return nil, err
		}
	}
	value, err = promptSchema(cmd, interactivePrompter(cmd), "Request body", schema, value, 0, false)
	if err != nil {
		return nil, err
	}
	return json.Marshal(value)
}

func promptSchema(cmd *cobra.Command, prompter Prompter, title string, schema *openapi3.Schema, value any, depth int, secret bool) (any, error) {
	if schema == nil {
		return nil, Exit(2, fmt.Errorf("cannot represent form schema safely; use --body @file"))
	}
	secret = secret || schema.WriteOnly || schema.Format == "password"
	complex := len(schema.Properties) > 0 || len(schema.AllOf) > 0 || len(schema.OneOf) > 0 || len(schema.AnyOf) > 0 || schema.Type != nil && (schema.Type.Is("object") || schema.Type.Is("array"))
	if depth > 3 && complex {
		return editSchemaValue(cmd, schema, value, secret)
	}
	variants := schema.OneOf
	if len(variants) == 0 {
		variants = schema.AnyOf
	}
	if len(variants) > 0 {
		labels := make([]string, len(variants))
		for i, ref := range variants {
			labels[i] = fmt.Sprintf("Variant %d", i+1)
			if ref.Value != nil && ref.Value.Title != "" {
				labels[i] = ref.Value.Title
			}
		}
		selected, err := prompter.SelectFields(cmd.Context(), title+": choose one variant", labels, nil)
		if err != nil {
			return nil, err
		}
		if len(selected) != 1 {
			return nil, Exit(2, fmt.Errorf("choose exactly one schema variant"))
		}
		for i, label := range labels {
			if label == selected[0] {
				return promptSchema(cmd, prompter, title, variants[i].Value, value, depth+1, secret)
			}
		}
		return nil, Exit(2, fmt.Errorf("unknown schema variant"))
	}
	if len(schema.AllOf) > 0 {
		return editSchemaValue(cmd, schema, value, secret)
	}
	if len(schema.Enum) > 0 {
		return prompter.Ask(cmd.Context(), PromptField{Title: title, Schema: schema, Value: value, Secret: secret})
	}
	if schema.Type != nil && schema.Type.Is("object") || len(schema.Properties) > 0 {
		if secret {
			return editSchemaValue(cmd, schema, value, true)
		}
		object := map[string]any{}
		if existing, ok := value.(map[string]any); ok {
			for key, item := range existing {
				object[key] = item
			}
		}
		required := map[string]bool{}
		for _, key := range schema.Required {
			required[key] = true
		}
		keys := make([]string, 0, len(schema.Properties))
		for key := range schema.Properties {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		optional := []string{}
		for _, key := range keys {
			property := schema.Properties[key].Value
			if property == nil || property.ReadOnly {
				continue
			}
			existing, present := object[key]
			if !required[key] && !present {
				optional = append(optional, key)
				continue
			}
			edited, err := promptSchema(cmd, prompter, title+"."+key, property, existing, depth+1, sensitiveName(key))
			if err != nil {
				return nil, err
			}
			object[key] = edited
		}
		if len(optional) > 0 {
			selected, err := prompter.SelectFields(cmd.Context(), title+": Configure optional fields?", optional, nil)
			if err != nil {
				return nil, err
			}
			for _, key := range selected {
				ref := schema.Properties[key]
				if ref == nil {
					return nil, Exit(2, fmt.Errorf("unknown optional field"))
				}
				edited, err := promptSchema(cmd, prompter, title+"."+key, ref.Value, nil, depth+1, sensitiveName(key))
				if err != nil {
					return nil, err
				}
				object[key] = edited
			}
		}
		return object, nil
	}
	if schema.Type != nil && schema.Type.Is("array") {
		if schema.Items == nil {
			return editSchemaValue(cmd, schema, value, secret)
		}
		array := []any{}
		if existing, ok := value.([]any); ok {
			array = append(array, existing...)
		}
		for i, item := range array {
			edited, err := promptSchema(cmd, prompter, fmt.Sprintf("%s[%d]", title, i), schema.Items.Value, item, depth+1, secret)
			if err != nil {
				return nil, err
			}
			array[i] = edited
		}
		for uint64(len(array)) < schema.MinItems {
			item, err := promptSchema(cmd, prompter, fmt.Sprintf("%s[%d]", title, len(array)), schema.Items.Value, nil, depth+1, secret)
			if err != nil {
				return nil, err
			}
			array = append(array, item)
		}
		for {
			boolean := openapi3.NewBoolSchema()
			more, err := prompter.Ask(cmd.Context(), PromptField{Title: title + ": add another?", Schema: boolean, Value: false})
			if err != nil {
				return nil, err
			}
			add, ok := more.(bool)
			if !ok {
				return nil, Exit(2, fmt.Errorf("invalid array confirmation"))
			}
			if !add {
				break
			}
			item, err := promptSchema(cmd, prompter, fmt.Sprintf("%s[%d]", title, len(array)), schema.Items.Value, nil, depth+1, secret)
			if err != nil {
				return nil, err
			}
			array = append(array, item)
		}
		return array, nil
	}
	if schema.Type != nil && (schema.Type.Is("string") || schema.Type.Is("integer") || schema.Type.Is("number") || schema.Type.Is("boolean")) {
		return prompter.Ask(cmd.Context(), PromptField{Title: title, Schema: schema, Value: value, Secret: secret})
	}
	return editSchemaValue(cmd, schema, value, secret)
}

func editSchemaValue(cmd *cobra.Command, schema *openapi3.Schema, value any, secret bool) (any, error) {
	if secret {
		return nil, Exit(2, fmt.Errorf("form schema contains sensitive fields that cannot be edited safely; use --body @file"))
	}
	_, redacted := redactPreview(value, []*openapi3.Schema{schema})
	if redacted || hasSensitiveSchema(schema, map[*openapi3.Schema]bool{}) {
		return nil, Exit(2, fmt.Errorf("sensitive schema cannot be opened in an editor; use --body @file"))
	}
	editor := os.Getenv("EDITOR")
	if editor == "" {
		return nil, Exit(2, fmt.Errorf("unsupported/deep form schema requires $EDITOR; use --body @file"))
	}
	// A single executable path avoids shell interpretation of editor arguments.
	file, err := os.CreateTemp("", "flexera-cli-request-*.json")
	if err != nil {
		return nil, err
	}
	defer os.Remove(file.Name())
	if value == nil {
		value = editorExample(schema, 0, map[*openapi3.Schema]bool{})
	}
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		file.Close()
		return nil, err
	}
	if _, err := file.Write(data); err != nil {
		file.Close()
		return nil, err
	}
	if err := file.Close(); err != nil {
		return nil, err
	}
	process := exec.CommandContext(cmd.Context(), editor, file.Name())
	process.Stdin = cmd.InOrStdin()
	process.Stdout = cmd.ErrOrStderr()
	process.Stderr = cmd.ErrOrStderr()
	if err := process.Run(); err != nil {
		return nil, Exit(2, fmt.Errorf("editor failed; use --body @file"))
	}
	data, err = os.ReadFile(file.Name())
	if err != nil {
		return nil, err
	}
	return ParseRequestJSON(data)
}

func editorExample(schema *openapi3.Schema, depth int, seen map[*openapi3.Schema]bool) any {
	if schema == nil || depth > 4 || seen[schema] || schema.ReadOnly {
		return nil
	}
	seen[schema] = true
	defer delete(seen, schema)
	if schema.Example != nil {
		return schema.Example
	}
	if schema.Default != nil {
		return schema.Default
	}
	if len(schema.Enum) > 0 {
		return schema.Enum[0]
	}
	for _, variants := range []openapi3.SchemaRefs{schema.OneOf, schema.AnyOf} {
		if len(variants) > 0 {
			return editorExample(variants[0].Value, depth+1, seen)
		}
	}
	if schema.Type != nil {
		switch {
		case schema.Type.Is("string"):
			return "REPLACE_ME"
		case schema.Type.Is("integer") || schema.Type.Is("number"):
			return 0
		case schema.Type.Is("boolean"):
			return false
		case schema.Type.Is("array"):
			if schema.Items != nil {
				return []any{editorExample(schema.Items.Value, depth+1, seen)}
			}
		}
	}
	value := map[string]any{}
	for _, ref := range schema.AllOf {
		if object, ok := editorExample(ref.Value, depth+1, seen).(map[string]any); ok {
			for key, item := range object {
				value[key] = item
			}
		}
	}
	for _, key := range schema.Required {
		if ref := schema.Properties[key]; ref != nil && !ref.Value.ReadOnly {
			value[key] = editorExample(ref.Value, depth+1, seen)
		}
	}
	return value
}

func hasSensitiveSchema(schema *openapi3.Schema, seen map[*openapi3.Schema]bool) bool {
	if schema == nil || seen[schema] {
		return false
	}
	seen[schema] = true
	if schema.WriteOnly || schema.Format == "password" {
		return true
	}
	for key, ref := range schema.Properties {
		if sensitiveName(key) || hasSensitiveSchema(ref.Value, seen) {
			return true
		}
	}
	if schema.Items != nil && hasSensitiveSchema(schema.Items.Value, seen) {
		return true
	}
	if schema.AdditionalProperties.Schema != nil && hasSensitiveSchema(schema.AdditionalProperties.Schema.Value, seen) {
		return true
	}
	for _, variants := range []openapi3.SchemaRefs{schema.AllOf, schema.OneOf, schema.AnyOf} {
		for _, ref := range variants {
			if hasSensitiveSchema(ref.Value, seen) {
				return true
			}
		}
	}
	return false
}

func ConfirmInteractive(cmd *cobra.Command, dryRun, yes bool, plan Plan, printer Printer) (bool, error) {
	if err := plan.RenderHuman(cmd.ErrOrStderr()); err != nil {
		return false, err
	}
	if len(plan.Body) > 0 && !plan.RawUpload {
		body, err := ParseRequestJSON(plan.Body)
		if err != nil {
			return false, err
		}
		_, sensitive := redactPreview(body, []*openapi3.Schema{plan.RequestSchema})
		if sensitive || len(plan.Body) > 1024 {
			choices, err := interactivePrompter(cmd).SelectFields(cmd.Context(), "Save the exact body to a private file? (may contain secrets)", []string{"Save body"}, nil)
			if err != nil {
				return false, err
			}
			if len(choices) > 0 {
				pathValue, err := interactivePrompter(cmd).Ask(cmd.Context(), PromptField{Title: "New private body file path", Schema: openapi3.NewStringSchema()})
				if err != nil {
					return false, err
				}
				path, ok := pathValue.(string)
				if !ok || strings.TrimSpace(path) == "" {
					return false, Exit(2, fmt.Errorf("provide a new private file path"))
				}
				file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
				if err != nil {
					return false, Exit(2, fmt.Errorf("cannot create private body file; choose a new path"))
				}
				if _, err := file.Write(plan.Body); err != nil {
					file.Close()
					os.Remove(path)
					return false, err
				}
				if err := file.Close(); err != nil {
					os.Remove(path)
					return false, err
				}
				fmt.Fprintf(cmd.ErrOrStderr(), "Exact body saved with mode 0600; use --body %s with the reviewed parameters.\n", shellQuote("@"+path))
			}
		}
	}
	if dryRun {
		return true, plan.RenderJSON(cmd.OutOrStdout(), printer)
	}
	if yes {
		return false, nil
	}
	answer, err := interactivePrompter(cmd).Approve(cmd.Context())
	if err != nil {
		return false, err
	}
	if answer != "yes" {
		_, _ = io.WriteString(cmd.ErrOrStderr(), "Apply cancelled.\n")
		return false, Exit(1, fmt.Errorf("apply cancelled"))
	}
	return false, nil
}
