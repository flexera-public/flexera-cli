package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/flexera-public/flexera-cli/internal/catalog"
)

// Only illustrative annotations are removed. Schema properties, defaults that
// are safe, and validation constraints remain intact; unsafe object samples are
// omitted whole rather than rewritten into misleading or invalid examples.
type exampleSafety struct {
	schemas    map[string]any
	secretRefs map[string]bool
	removed    int
}

func sampleRefName(ref string) string {
	return strings.NewReplacer("~1", "/", "~0", "~").Replace(strings.TrimPrefix(ref, "#/components/schemas/"))
}

// A shared component used as a credential must not expose its own annotations
// through another operation or schema expansion. Only samples are suppressed.
func (s *exampleSafety) markSecretRefs(schema any, secret bool, active map[string]bool) {
	v, ok := schema.(map[string]any)
	if !ok {
		return
	}
	secret = secret || v["writeOnly"] == true || v["format"] == "password"
	if ref, ok := v["$ref"].(string); ok {
		name := sampleRefName(ref)
		if secret {
			s.secretRefs[name] = true
		}
		// Visit once per sensitivity context, including cyclic components.
		key := fmt.Sprintf("%s:%t", ref, secret)
		if !active[key] {
			active[key] = true
			s.markSecretRefs(s.schemas[name], secret, active)
		}
	}
	if properties, ok := v["properties"].(map[string]any); ok {
		for name, child := range properties {
			s.markSecretRefs(child, secret || secretSampleName(name), active)
		}
	}
	for _, key := range []string{"items", "additionalProperties", "not", "contains", "if", "then", "else"} {
		s.markSecretRefs(v[key], secret, active)
	}
	for _, key := range []string{"allOf", "anyOf", "oneOf", "prefixItems"} {
		if variants, ok := v[key].([]any); ok {
			for _, child := range variants {
				s.markSecretRefs(child, secret, active)
			}
		}
	}
}

func decodeSampleJSON(raw json.RawMessage) (any, error) {
	var value any
	d := json.NewDecoder(bytes.NewReader(raw))
	d.UseNumber()
	err := d.Decode(&value)
	return value, err
}

func secretSampleName(name string) bool {
	name = strings.ToLower(strings.NewReplacer("_", "", "-", "", " ", "").Replace(name))
	return strings.Contains(name, "password") || strings.Contains(name, "secret") || strings.Contains(name, "token") || strings.Contains(name, "privatekey") || name == "apikey" || name == "authorization"
}

func secretSampleKeys(value any) bool {
	switch v := value.(type) {
	case map[string]any:
		for name, child := range v {
			if secretSampleName(name) || secretSampleKeys(child) {
				return true
			}
		}
	case []any:
		for _, child := range v {
			if secretSampleKeys(child) {
				return true
			}
		}
	}
	return false
}

func (s *exampleSafety) sensitive(schema, value any, active map[string]bool) bool {
	v, ok := schema.(map[string]any)
	if !ok {
		return false
	}
	if v["writeOnly"] == true || v["format"] == "password" {
		return true
	}
	if ref, ok := v["$ref"].(string); ok && !active[ref] {
		name := sampleRefName(ref)
		if s.secretRefs[name] {
			return true
		}
		active[ref] = true
		unsafe := s.sensitive(s.schemas[name], value, active)
		delete(active, ref)
		if unsafe {
			return true
		}
	}
	for _, key := range []string{"allOf", "anyOf", "oneOf"} {
		if variants, ok := v[key].([]any); ok {
			for _, variant := range variants {
				if s.sensitive(variant, value, active) {
					return true
				}
			}
		}
	}
	if object, ok := value.(map[string]any); ok {
		properties, _ := v["properties"].(map[string]any)
		for name, child := range object {
			childSchema, found := properties[name]
			if !found {
				childSchema = v["additionalProperties"]
			}
			// Consuming an instance child makes recursive schemas safe to visit
			// again; a schema-only cycle guard would miss nested credentials.
			if secretSampleName(name) || s.sensitive(childSchema, child, map[string]bool{}) {
				return true
			}
		}
	}
	if array, ok := value.([]any); ok {
		for _, child := range array {
			if s.sensitive(v["items"], child, map[string]bool{}) {
				return true
			}
		}
	}
	return false
}

func (s *exampleSafety) sanitize(schema any, secret bool) {
	v, ok := schema.(map[string]any)
	if !ok {
		return
	}
	secret = secret || v["writeOnly"] == true || v["format"] == "password"
	if ref, ok := v["$ref"].(string); ok {
		secret = secret || s.secretRefs[sampleRefName(ref)]
	}
	for _, key := range []string{"example", "default", "examples"} {
		value, found := v[key]
		if !found {
			continue
		}
		unsafe := secret || secretSampleKeys(value) || s.sensitive(v, value, map[string]bool{})
		if key == "examples" {
			if examples, ok := value.([]any); ok {
				for _, example := range examples {
					unsafe = unsafe || s.sensitive(v, example, map[string]bool{})
				}
			}
		}
		if unsafe {
			delete(v, key)
			s.removed++
		}
	}
	if properties, ok := v["properties"].(map[string]any); ok {
		for name, child := range properties {
			s.sanitize(child, secret || secretSampleName(name))
		}
	}
	for _, key := range []string{"items", "additionalProperties", "not", "contains", "if", "then", "else"} {
		s.sanitize(v[key], secret)
	}
	for _, key := range []string{"allOf", "anyOf", "oneOf", "prefixItems"} {
		if variants, ok := v[key].([]any); ok {
			for _, child := range variants {
				s.sanitize(child, secret)
			}
		}
	}
}

func sanitizeCatalogExamples(doc *catalog.Document) error {
	// Validate against the complete original schemas, before removing samples.
	original, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		return err
	}
	index, err := catalog.Parse(original)
	if err != nil {
		return fmt.Errorf("validate original catalog: %w", err)
	}
	s := exampleSafety{schemas: map[string]any{}, secretRefs: map[string]bool{}}
	for name, raw := range doc.Schemas {
		value, err := decodeSampleJSON(raw)
		if err != nil {
			return fmt.Errorf("catalog schema %q: %w", name, err)
		}
		s.schemas[name] = value
	}
	active := map[string]bool{}
	for _, value := range s.schemas {
		s.markSecretRefs(value, false, active)
	}
	for _, e := range doc.Entries {
		for _, raw := range []json.RawMessage{e.RequestSchema, e.ResponseSchema} {
			if len(raw) != 0 {
				value, err := decodeSampleJSON(raw)
				if err != nil {
					return err
				}
				s.markSecretRefs(value, false, active)
			}
		}
		for _, p := range e.Params {
			if len(p.Schema) != 0 {
				value, err := decodeSampleJSON(p.Schema)
				if err != nil {
					return err
				}
				s.markSecretRefs(value, secretSampleName(p.Flag), active)
			}
		}
	}
	// Keep this index immutable while sanitizing, so map iteration order cannot
	// change transitive sensitivity decisions.
	sanitizeRaw := func(raw json.RawMessage, secret bool) (json.RawMessage, error) {
		if len(raw) == 0 {
			return raw, nil
		}
		value, err := decodeSampleJSON(raw)
		if err != nil {
			return nil, err
		}
		s.sanitize(value, secret)
		return json.Marshal(value)
	}
	omitted := 0
	for i := range doc.Entries {
		e := &doc.Entries[i]
		if len(e.RequestExample) != 0 {
			_, validationErr := index.ValidExample(*e)
			schema, schemaErr := decodeSampleJSON(e.RequestSchema)
			value, valueErr := decodeSampleJSON(e.RequestExample)
			if validationErr != nil || schemaErr != nil || valueErr != nil || s.sensitive(schema, value, map[string]bool{}) {
				e.RequestExample = nil
				omitted++
			}
		}
		for _, raw := range []*json.RawMessage{&e.RequestSchema, &e.ResponseSchema} {
			*raw, err = sanitizeRaw(*raw, false)
			if err != nil {
				return err
			}
		}
		e.Params = append([]catalog.Param(nil), e.Params...)
		for j := range e.Params {
			p := &e.Params[j]
			p.Schema, err = sanitizeRaw(p.Schema, secretSampleName(p.Flag))
			if err != nil {
				return err
			}
		}
	}
	for name, raw := range doc.Schemas {
		doc.Schemas[name], err = sanitizeRaw(raw, s.secretRefs[name])
		if err != nil {
			return err
		}
	}
	after, err := json.MarshalIndent(doc, "", "  ")
	if err == nil {
		fmt.Fprintf(os.Stderr, "catalog samples: omitted %d request examples, removed %d schema samples; bytes %d -> %d\n", omitted, s.removed, len(original)+1, len(after)+1)
	}
	return err
}
