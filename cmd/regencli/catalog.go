package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/flexera-public/flexera-cli/internal/catalog"
)

// buildCatalog consumes staged gencli metadata without publishing or requiring
// complete spec coverage. The returned byte length is the measured artifact size;
// no size cap is imposed before a real generated catalog has been measured.
func buildCatalog(stage string, tags []genTag, specData []byte, curated ...catalog.Entry) ([]byte, error) {
	var s spec
	if err := json.Unmarshal(specData, &s); err != nil {
		return nil, fmt.Errorf("catalog spec: %w", err)
	}
	type location struct{ method, path string }
	operations := map[string]location{}
	paths := make([]string, 0, len(s.Paths))
	for path := range s.Paths {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		for _, method := range []string{"get", "post", "put", "patch", "delete", "head", "options", "trace"} {
			raw, ok := s.Paths[path][method]
			if !ok {
				continue
			}
			var op struct {
				ID string `json:"operationId"`
			}
			if err := json.Unmarshal(raw, &op); err != nil {
				return nil, fmt.Errorf("catalog spec %s %s: %w", method, path, err)
			}
			if op.ID == "" {
				continue
			}
			if _, exists := operations[op.ID]; exists {
				return nil, fmt.Errorf("catalog spec: duplicate operationId %q", op.ID)
			}
			operations[op.ID] = location{strings.ToUpper(method), path}
		}
	}

	// Copy before sorting: callers may reuse their generation order.
	tags = append([]genTag(nil), tags...)
	sort.Slice(tags, func(i, j int) bool { return tags[i].Pkg < tags[j].Pkg })
	parentCommand := ""
	for _, tag := range tags {
		if tag.Pkg == billConnectParent {
			parentCommand = tag.Cmd
		}
	}
	doc := catalog.Document{Entries: append([]catalog.Entry{}, curated...), Schemas: map[string]json.RawMessage{}}
	seen := map[string]bool{}
	for _, e := range curated {
		if seen[e.OperationID] {
			return nil, fmt.Errorf("catalog: duplicate curated operation %s", e.OperationID)
		}
		seen[e.OperationID] = true
	}
	for _, tag := range tags {
		path := filepath.Join(stage, tag.Pkg, "metadata.json")
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("catalog metadata %s: %w", path, err)
		}
		var entries map[string]catalog.Entry
		if err := json.Unmarshal(data, &entries); err != nil {
			return nil, fmt.Errorf("catalog metadata %s: %w", path, err)
		}
		if entries == nil {
			return nil, fmt.Errorf("catalog metadata %s: expected operation map", path)
		}
		ids := make([]string, 0, len(entries))
		for id := range entries {
			ids = append(ids, id)
		}
		sort.Strings(ids)
		for _, id := range ids {
			e := entries[id]
			if id == "" || e.OperationID != id {
				return nil, fmt.Errorf("catalog metadata %s: operationId %q does not match key %q", path, e.OperationID, id)
			}
			if seen[id] {
				return nil, fmt.Errorf("catalog metadata: duplicate operationId %q across tags", id)
			}
			seen[id] = true
			op, exists := operations[id]
			if !exists {
				return nil, fmt.Errorf("catalog metadata: operationId %q missing from spec", id)
			}
			if e.Method != op.method || e.Path != op.path {
				return nil, fmt.Errorf("catalog metadata %s: method/path %s %s does not match spec %s %s", id, e.Method, e.Path, op.method, op.path)
			}
			if len(e.Command) == 0 || e.Command[0] != tag.Cmd {
				return nil, fmt.Errorf("catalog metadata %s: missing or mismatched provisional command", id)
			}
			if parentCommand != "" {
				for _, child := range billConnectChildren {
					if tag.Pkg == child.Pkg {
						e.Command = append([]string{parentCommand, child.Use}, e.Command[1:]...)
						break
					}
				}
			}
			doc.Entries = append(doc.Entries, e)
		}
	}
	sort.Slice(doc.Entries, func(i, j int) bool { return doc.Entries[i].OperationID < doc.Entries[j].OperationID })

	// Walk metadata as well as schemas, including Param.Schema and every
	// request/response composition. Mark components before walking for cycles.
	var walk func(interface{}) error
	walk = func(value interface{}) error {
		switch v := value.(type) {
		case map[string]interface{}:
			if raw, exists := v["$ref"]; exists {
				ref, ok := raw.(string)
				const prefix = "#/components/schemas/"
				name := strings.TrimPrefix(ref, prefix)
				if !ok || !strings.HasPrefix(ref, prefix) || name == "" || strings.Contains(name, "/") {
					return fmt.Errorf("catalog: unsupported schema ref %v", raw)
				}
				name = strings.NewReplacer("~1", "/", "~0", "~").Replace(name)
				if _, loaded := doc.Schemas[name]; !loaded {
					schema, exists := s.Components.Schemas[name]
					if !exists {
						return fmt.Errorf("catalog: dangling schema ref %q", ref)
					}
					doc.Schemas[name] = schema
					var nested interface{}
					if err := json.Unmarshal(schema, &nested); err != nil {
						return fmt.Errorf("catalog schema %q: %w", name, err)
					}
					if err := walk(nested); err != nil {
						return err
					}
				}
			}
			keys := make([]string, 0, len(v))
			for key := range v {
				keys = append(keys, key)
			}
			sort.Strings(keys)
			for _, key := range keys {
				// Match catalog.Parse: examples are instance data, not refs.
				if key == "example" || key == "examples" || key == "requestExample" {
					continue
				}
				if key == "properties" {
					if properties, ok := v[key].(map[string]interface{}); ok {
						for _, property := range properties {
							if err := walk(property); err != nil {
								return err
							}
						}
						continue
					}
				}
				if err := walk(v[key]); err != nil {
					return err
				}
			}
		case []interface{}:
			for _, nested := range v {
				if err := walk(nested); err != nil {
					return err
				}
			}
		}
		return nil
	}
	metadata, err := json.Marshal(doc.Entries)
	if err != nil {
		return nil, fmt.Errorf("catalog entries: %w", err)
	}
	var value interface{}
	if err := json.Unmarshal(metadata, &value); err != nil {
		return nil, err
	}
	if err := walk(value); err != nil {
		return nil, err
	}
	if err := sanitizeCatalogExamples(&doc); err != nil {
		return nil, fmt.Errorf("catalog sample safety: %w", err)
	}
	data, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("catalog encoding: %w", err)
	}
	data = append(data, '\n')
	if _, err := catalog.Parse(data); err != nil {
		return nil, fmt.Errorf("validate staged catalog: %w", err)
	}
	return data, nil
}
