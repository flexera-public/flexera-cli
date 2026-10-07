package catalog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/getkin/kin-openapi/openapi3"
)

type schemaCacheKey struct {
	catalog *Catalog
	raw     string
}
type schemaCacheValue struct {
	once   sync.Once
	schema *openapi3.Schema
	err    error
}

var requestSchemas sync.Map

type catalogSchemaSet struct {
	once  sync.Once
	roots map[string]*openapi3.Schema
	err   error
}

var catalogSchemas sync.Map

func schemaKey(raw json.RawMessage) string {
	var compact bytes.Buffer
	if err := json.Compact(&compact, raw); err != nil {
		return string(raw)
	}
	return compact.String()
}

func (c *Catalog) resolvedSchemas() (*catalogSchemaSet, error) {
	value, _ := catalogSchemas.LoadOrStore(c, &catalogSchemaSet{})
	set := value.(*catalogSchemaSet)
	set.once.Do(func() {
		schemas := map[string]json.RawMessage{}
		for name, raw := range c.document.Schemas {
			schemas[name] = raw
		}
		names := map[string]string{}
		add := func(raw json.RawMessage) {
			if len(raw) == 0 {
				return
			}
			key := schemaKey(raw)
			if names[key] != "" {
				return
			}
			name := fmt.Sprintf("__cli_root_%d", len(names))
			for schemas[name] != nil {
				name += "_"
			}
			names[key] = name
			schemas[name] = raw
		}
		for _, entry := range c.document.Entries {
			add(entry.RequestSchema)
			for _, param := range entry.Params {
				add(param.Schema)
			}
		}
		document := map[string]any{"openapi": "3.0.3", "info": map[string]string{"title": "CLI validation", "version": "1"}, "paths": map[string]any{}, "components": map[string]any{"schemas": schemas}}
		data, err := json.Marshal(document)
		if err != nil {
			set.err = err
			return
		}
		loader := openapi3.NewLoader()
		loader.IsExternalRefsAllowed = false
		doc, err := loader.LoadFromData(data)
		if err != nil {
			set.err = fmt.Errorf("catalog schema references could not be resolved")
			return
		}
		set.roots = map[string]*openapi3.Schema{}
		for key, name := range names {
			set.roots[key] = doc.Components.Schemas[name].Value
		}
	})
	return set, set.err
}

// RequestSchema resolves local OAS 3.0 references once per catalog/schema.
// The returned schema is shared and must not be modified by callers.
func (c *Catalog) RequestSchema(raw json.RawMessage) (*openapi3.Schema, error) {
	if len(raw) == 0 {
		return nil, fmt.Errorf("operation has no request schema")
	}
	set, err := c.resolvedSchemas()
	if err != nil {
		return nil, err
	}
	if schema := set.roots[schemaKey(raw)]; schema != nil {
		return schema, nil
	}
	value, _ := requestSchemas.LoadOrStore(schemaCacheKey{c, schemaKey(raw)}, &schemaCacheValue{})
	cached := value.(*schemaCacheValue)
	cached.once.Do(func() {
		schemas := map[string]json.RawMessage{}
		for name, body := range c.document.Schemas {
			schemas[name] = body
		}
		name := "__flexera_cli_request__"
		for schemas[name] != nil {
			name += "_"
		}
		schemas[name] = raw
		document := map[string]any{"openapi": "3.0.3", "info": map[string]string{"title": "CLI validation", "version": "1"}, "paths": map[string]any{}, "components": map[string]any{"schemas": schemas}}
		data, err := json.Marshal(document)
		if err != nil {
			cached.err = err
			return
		}
		loader := openapi3.NewLoader()
		loader.IsExternalRefsAllowed = false
		doc, err := loader.LoadFromData(data)
		if err != nil {
			cached.err = fmt.Errorf("request schema references could not be resolved")
			return
		}
		cached.schema = doc.Components.Schemas[name].Value
	})
	return cached.schema, cached.err
}
