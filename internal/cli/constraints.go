package cli

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/mail"
	"net/url"
	"strconv"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/google/uuid"
)

func validateSchemaValue(schema *openapi3.Schema, value any) error {
	if err := strictConstraints(schema, value); err != nil {
		return err
	}
	if err := structuralSchema(schema, map[*openapi3.Schema]*openapi3.Schema{}).VisitJSON(value, openapi3.MultiErrors(), openapi3.VisitAsRequest()); err != nil {
		return err
	}
	return nil
}

func structuralSchema(schema *openapi3.Schema, seen map[*openapi3.Schema]*openapi3.Schema) *openapi3.Schema {
	if schema == nil {
		return nil
	}
	if copy, ok := seen[schema]; ok {
		return copy
	}
	copy := *schema
	seen[schema] = &copy
	if schema.Type != nil && (schema.Type.Is("integer") || schema.Type.Is("number")) {
		copy.Type = openapi3.NewFloat64Schema().Type
		copy.Min = nil
		copy.Max = nil
		copy.MultipleOf = nil
		copy.ExclusiveMin = openapi3.ExclusiveBound{}
		copy.ExclusiveMax = openapi3.ExclusiveBound{}
		copy.Enum = nil
	}
	ref := func(value *openapi3.SchemaRef) *openapi3.SchemaRef {
		if value == nil {
			return nil
		}
		r := *value
		r.Value = structuralSchema(value.Value, seen)
		return &r
	}
	copy.Properties = openapi3.Schemas{}
	for key, value := range schema.Properties {
		copy.Properties[key] = ref(value)
	}
	copy.Items = ref(schema.Items)
	copy.AdditionalProperties.Schema = ref(schema.AdditionalProperties.Schema)
	copy.Not = ref(schema.Not)
	variants := func(values openapi3.SchemaRefs) openapi3.SchemaRefs {
		out := make(openapi3.SchemaRefs, len(values))
		for i, value := range values {
			out[i] = ref(value)
		}
		return out
	}
	copy.AllOf = variants(schema.AllOf)
	copy.OneOf = variants(schema.OneOf)
	copy.AnyOf = variants(schema.AnyOf)
	return &copy
}
func strictConstraints(s *openapi3.Schema, value any) error {
	if s == nil || value == nil {
		return nil
	}
	fail := func(field, reason string) error {
		return &openapi3.SchemaError{Schema: s, SchemaField: field, Reason: reason}
	}
	for _, ref := range s.AllOf {
		if err := strictConstraints(ref.Value, value); err != nil {
			return err
		}
	}
	for _, variants := range []openapi3.SchemaRefs{s.OneOf, s.AnyOf} {
		if len(variants) > 0 {
			matches := 0
			for _, ref := range variants {
				if ref.Value != nil && validateSchemaValue(ref.Value, value) == nil {
					matches++
				}
			}
			if matches == 0 {
				return fail("composition", "value does not satisfy a schema variant")
			}
		}
	}
	if n, ok := value.(json.Number); ok {
		r, ok := new(big.Rat).SetString(string(n))
		if !ok {
			return fail("type", "invalid numeric input")
		}
		if s.Type != nil && s.Type.Is("integer") && !r.IsInt() {
			return fail("type", "value must be an exact integer")
		}
		bound := func(value float64) *big.Rat {
			r, _ := new(big.Rat).SetString(strconv.FormatFloat(value, 'g', -1, 64))
			return r
		}
		if s.Min != nil {
			cmp := r.Cmp(bound(*s.Min))
			if cmp < 0 || cmp == 0 && s.ExclusiveMin.IsTrue() {
				return fail("minimum", "value is below its numeric minimum")
			}
		}
		if s.Max != nil {
			cmp := r.Cmp(bound(*s.Max))
			if cmp > 0 || cmp == 0 && s.ExclusiveMax.IsTrue() {
				return fail("maximum", "value exceeds its numeric maximum")
			}
		}
		if s.MultipleOf != nil {
			multiple := bound(*s.MultipleOf)
			if multiple.Sign() != 0 && !new(big.Rat).Quo(r, multiple).IsInt() {
				return fail("multipleOf", "value is not an exact permitted multiple")
			}
		}
		if len(s.Enum) > 0 {
			matches := false
			for _, v := range s.Enum {
				raw, err := json.Marshal(v)
				if err != nil {
					continue
				}
				other, ok := new(big.Rat).SetString(string(raw))
				matches = matches || ok && r.Cmp(other) == 0
			}
			if !matches {
				return fail("enum", "value is not an allowed numeric enum")
			}
		}
	}
	if text, ok := value.(string); ok {
		valid := true
		switch s.Format {
		case "uuid":
			_, err := uuid.Parse(text)
			valid = err == nil
		case "date":
			_, err := time.Parse("2006-01-02", text)
			valid = err == nil
		case "date-time":
			_, err := time.Parse(time.RFC3339Nano, text)
			valid = err == nil
		case "email":
			address, err := mail.ParseAddress(text)
			valid = err == nil && address.Address == text
		case "uri":
			u, err := url.Parse(text)
			valid = err == nil && u.IsAbs()
		}
		if !valid {
			return fail("format", fmt.Sprintf("value must have valid %s format", s.Format))
		}
	}
	if object, ok := value.(map[string]any); ok {
		for key, item := range object {
			ref := s.Properties[key]
			if ref == nil {
				ref = s.AdditionalProperties.Schema
			}
			if ref != nil {
				if err := strictConstraints(ref.Value, item); err != nil {
					return err
				}
			}
		}
	}
	if array, ok := value.([]any); ok && s.Items != nil {
		for _, item := range array {
			if err := strictConstraints(s.Items.Value, item); err != nil {
				return err
			}
		}
	}
	return nil
}
