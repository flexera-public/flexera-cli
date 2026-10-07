package catalog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Expand resolves references up to depth reference hops. Cycles and references
// beyond the requested depth remain references. Property names are never refs.
func (c *Catalog) Expand(raw json.RawMessage, depth int) (json.RawMessage, error) {
	if depth < 0 {
		return nil, fmt.Errorf("schema depth must not be negative")
	}
	if len(raw) == 0 {
		return nil, nil
	}
	var value any
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.UseNumber()
	if err := decoder.Decode(&value); err != nil {
		return nil, err
	}
	var expand func(any, int, map[string]bool) (any, error)
	expand = func(value any, remaining int, active map[string]bool) (any, error) {
		switch v := value.(type) {
		case map[string]any:
			if ref, ok := v["$ref"].(string); ok && remaining > 0 && !active[ref] {
				raw, found := c.Schema(ref)
				if !found {
					return nil, fmt.Errorf("unknown schema reference %s", ref)
				}
				var target any
				d := json.NewDecoder(bytes.NewReader(raw))
				d.UseNumber()
				if err := d.Decode(&target); err != nil {
					return nil, err
				}
				active[ref] = true
				expanded, err := expand(target, remaining-1, active)
				delete(active, ref)
				return expanded, err
			}
			result := map[string]any{}
			for key, nested := range v {
				if key == "example" || key == "examples" || key == "default" || key == "enum" {
					result[key] = nested
					continue
				}
				if key == "properties" {
					if properties, ok := nested.(map[string]any); ok {
						out := map[string]any{}
						for name, property := range properties {
							expanded, err := expand(property, remaining, active)
							if err != nil {
								return nil, err
							}
							out[name] = expanded
						}
						result[key] = out
						continue
					}
				}
				expanded, err := expand(nested, remaining, active)
				if err != nil {
					return nil, err
				}
				result[key] = expanded
			}
			return result, nil
		case []any:
			result := make([]any, len(v))
			for i, item := range v {
				expanded, err := expand(item, remaining, active)
				if err != nil {
					return nil, err
				}
				result[i] = expanded
			}
			return result, nil
		default:
			return value, nil
		}
	}
	result, err := expand(value, depth, map[string]bool{})
	if err != nil {
		return nil, err
	}
	return json.Marshal(result)
}

// ValidExample refuses absent, invalid or secret-bearing illustrative samples.
// Validation is local: external reference resolution is deliberately disabled.
func (c *Catalog) ValidExample(e Entry) (json.RawMessage, error) {
	if len(e.RequestExample) == 0 || len(e.RequestSchema) == 0 {
		return nil, fmt.Errorf("no request example is available; use cli schema to construct your own body")
	}
	var value any
	d := json.NewDecoder(bytes.NewReader(e.RequestExample))
	d.UseNumber()
	if err := d.Decode(&value); err != nil {
		return nil, fmt.Errorf("request example is invalid JSON")
	}
	if sensitiveExample(value) {
		return nil, fmt.Errorf("request example contains sensitive fields; supply your own private request file")
	}
	schema, err := c.RequestSchema(e.RequestSchema)
	if err != nil {
		return nil, fmt.Errorf("request example schema could not be resolved")
	}
	if schemaSensitiveValue(schema, value, map[*openapi3.Schema]bool{}) {
		return nil, fmt.Errorf("request example contains schema-marked sensitive values; supply your own private request file")
	}
	if err := schema.VisitJSON(value, openapi3.MultiErrors(), openapi3.VisitAsRequest()); err != nil {
		return nil, fmt.Errorf("illustrative request example does not satisfy its request schema; supply your own body")
	}
	return append(json.RawMessage(nil), e.RequestExample...), nil
}

func sensitiveExample(value any) bool {
	switch v := value.(type) {
	case map[string]any:
		for key, item := range v {
			normalized := strings.ToLower(strings.NewReplacer("_", "", "-", "", " ", "").Replace(key))
			if strings.Contains(normalized, "password") || strings.Contains(normalized, "secret") || strings.Contains(normalized, "token") || strings.Contains(normalized, "privatekey") || normalized == "apikey" || normalized == "authorization" {
				return true
			}
			if sensitiveExample(item) {
				return true
			}
		}
	case []any:
		for _, item := range v {
			if sensitiveExample(item) {
				return true
			}
		}
	}
	return false
}

func schemaSensitiveValue(schema *openapi3.Schema, value any, active map[*openapi3.Schema]bool) bool {
	if schema == nil || active[schema] {
		return false
	}
	if schema.WriteOnly || schema.Format == "password" {
		return true
	}
	active[schema] = true
	defer delete(active, schema)
	for _, variants := range []openapi3.SchemaRefs{schema.AllOf, schema.OneOf, schema.AnyOf} {
		for _, ref := range variants {
			if schemaSensitiveValue(ref.Value, value, active) {
				return true
			}
		}
	}
	if object, ok := value.(map[string]any); ok {
		for key, item := range object {
			ref := schema.Properties[key]
			if ref == nil {
				ref = schema.AdditionalProperties.Schema
			}
			if ref != nil && schemaSensitiveValue(ref.Value, item, map[*openapi3.Schema]bool{}) {
				return true
			}
		}
	}
	if array, ok := value.([]any); ok && schema.Items != nil {
		for _, item := range array {
			if schemaSensitiveValue(schema.Items.Value, item, map[*openapi3.Schema]bool{}) {
				return true
			}
		}
	}
	return false
}
