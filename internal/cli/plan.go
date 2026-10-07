package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/flexera-public/flexera-cli/internal/catalog"
	"github.com/getkin/kin-openapi/openapi3"
)

type Plan struct {
	Command, Method, Path string
	OrgID                 int
	Params                map[string]any
	Body                  json.RawMessage
	Destructive           bool
	Validation            *ValidationResult
	RawUpload             bool
	RequestSchema         *openapi3.Schema
}

func sensitiveName(name string) bool {
	name = strings.ToLower(strings.NewReplacer("-", "", "_", "").Replace(name))
	return strings.Contains(name, "secret") || strings.Contains(name, "password") || strings.Contains(name, "token") || strings.Contains(name, "privatekey") || name == "apikey" || name == "authorization"
}

func redactPreview(value any, schemas []*openapi3.Schema) (any, bool) {
	for _, schema := range schemas {
		if schema != nil && (schema.WriteOnly || schema.Format == "password") {
			return "[REDACTED]", true
		}
	}
	// Composition is merged for sensitivity only, never to fabricate a diff.
	expanded := append([]*openapi3.Schema{}, schemas...)
	seen := map[*openapi3.Schema]bool{}
	for i := 0; i < len(expanded); i++ {
		schema := expanded[i]
		if schema == nil || seen[schema] {
			continue
		}
		if schema.WriteOnly || schema.Format == "password" {
			return "[REDACTED]", true
		}
		seen[schema] = true
		for _, variants := range []openapi3.SchemaRefs{schema.AllOf, schema.OneOf, schema.AnyOf} {
			for _, ref := range variants {
				expanded = append(expanded, ref.Value)
			}
		}
	}
	switch v := value.(type) {
	case map[string]any:
		out := map[string]any{}
		redacted := false
		for key, item := range v {
			if sensitiveName(key) {
				out[key] = "[REDACTED]"
				redacted = true
				continue
			}
			children := []*openapi3.Schema{}
			for _, schema := range expanded {
				if schema != nil {
					if ref := schema.Properties[key]; ref != nil {
						children = append(children, ref.Value)
					} else if ref := schema.AdditionalProperties.Schema; ref != nil {
						children = append(children, ref.Value)
					}
				}
			}
			safe, changed := redactPreview(item, children)
			out[key] = safe
			redacted = redacted || changed
		}
		return out, redacted
	case []any:
		children := []*openapi3.Schema{}
		for _, schema := range expanded {
			if schema != nil && schema.Items != nil {
				children = append(children, schema.Items.Value)
			}
		}
		out := make([]any, len(v))
		redacted := false
		for i, item := range v {
			safe, changed := redactPreview(item, children)
			out[i] = safe
			redacted = redacted || changed
		}
		return out, redacted
	default:
		return value, false
	}
}

// RenderJSON is a safety output boundary; response shaping never applies.
func (p Plan) RenderJSON(w io.Writer, printer Printer) error {
	plan := map[string]any{"command": p.Command, "method": p.Method + " " + p.Path, "params": p.Params}
	if p.OrgID != 0 {
		plan["orgId"] = p.OrgID
	}
	redacted := false
	if len(p.Body) > 0 {
		if p.RawUpload {
			plan["body"] = map[string]any{"bytes": len(p.Body), "content": "[RAW UPLOAD OMITTED]"}
			redacted = true
		} else {
			body, err := ParseRequestJSON(p.Body)
			if err != nil {
				return err
			}
			safe, changed := redactPreview(body, []*openapi3.Schema{p.RequestSchema})
			plan["body"] = safe
			redacted = changed
		}
	}
	params, changed := redactPreview(p.Params, nil)
	plan["params"] = params
	redacted = redacted || changed
	output := map[string]any{"dryRun": true, "destructive": p.Destructive, "plan": plan}
	if p.Validation != nil {
		output["validation"] = p.Validation
	}
	if redacted {
		output["redacted"] = true
	}
	return WriteJSON(w, output, printer.Style, printer.IsTTY)
}

func ConfirmPlan(dryRun, yes bool, w io.Writer, plan Plan, printer Printer) (bool, error) {
	if dryRun {
		return true, plan.RenderJSON(w, printer)
	}
	if plan.Destructive && !yes {
		return false, Exit(2, fmt.Errorf("destructive operation requires --yes (or use --dry-run to preview)"))
	}
	return false, nil
}

// PrepareRequestBody validates original input and retains the effective encoded
// body for a preview identical to what the typed SDK will transmit.
func PrepareRequestBody(operationID string, raw []byte, target any, skip bool) (json.RawMessage, *ValidationResult, *openapi3.Schema, error) {
	index, err := catalog.Load()
	if err != nil {
		return nil, nil, nil, Exit(2, err)
	}
	entry, found := index.Lookup(operationID)
	if !found {
		return nil, nil, nil, Exit(2, fmt.Errorf("unknown request operation %s", operationID))
	}
	result, err := DecodeRequestBody(index, entry, raw, target, skip)
	if err != nil {
		return nil, nil, nil, err
	}
	schema, err := index.RequestSchema(entry.RequestSchema)
	if err != nil {
		return nil, nil, nil, Exit(2, err)
	}
	effective, err := json.Marshal(target)
	if err != nil {
		return nil, nil, nil, Exit(2, fmt.Errorf("API request type cannot be encoded"))
	}
	return effective, result, schema, nil
}
