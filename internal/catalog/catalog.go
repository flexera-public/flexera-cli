// Package catalog exposes generated operation metadata without importing the
// command tree. The embedded artifact is published with the generated command
// tree after staged operation-coverage and live-tree verification.
package catalog

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"sync"
)

type Param struct {
	Flag        string          `json:"flag"`
	Source      string          `json:"source"`
	In          string          `json:"in"`
	Type        string          `json:"type"`
	Required    bool            `json:"required"`
	Enum        json.RawMessage `json:"enum"`
	Schema      json.RawMessage `json:"schema"`
	Description string          `json:"description,omitempty"`
}

type Entry struct {
	OperationID      string          `json:"operationId"`
	Command          []string        `json:"command"`
	Method           string          `json:"method"`
	Path             string          `json:"path"`
	Tag              string          `json:"tag"`
	Service          string          `json:"service,omitempty"`
	Resource         string          `json:"resource"`
	Action           string          `json:"action"`
	Summary          string          `json:"summary"`
	Description      string          `json:"description"`
	Destructive      bool            `json:"destructive"`
	Paginated        bool            `json:"paginated"`
	ResponseEnvelope string          `json:"responseEnvelope"`
	Params           []Param         `json:"params"`
	BodyFlags        []string        `json:"bodyFlags"`
	RequestSchema    json.RawMessage `json:"requestSchema,omitempty"`
	ResponseSchema   json.RawMessage `json:"responseSchema,omitempty"`
	RequestExample   json.RawMessage `json:"requestExample,omitempty"`
}

// Document is the final publication shape. Per-tag generator metadata is a
// separate operationId-keyed map whose provisional command paths are reconciled
// by regencli before constructing this document.
type Document struct {
	Entries []Entry                    `json:"entries"`
	Schemas map[string]json.RawMessage `json:"schemas"`
}

type Catalog struct {
	document Document
	byID     map[string]int
}

//go:embed catalog_gen.json
var artifact []byte

var once sync.Once
var loaded *Catalog
var loadErr error

// Load parses and validates the embedded catalog once, including transitive
// schema references. It does not require config, authentication or API access.
func Load() (*Catalog, error) {
	once.Do(func() {
		loaded, loadErr = Parse(artifact)
		if loadErr == nil && len(loaded.document.Entries) == 0 {
			loaded = nil
			loadErr = fmt.Errorf("catalog: operation artifact has not been generated")
		}
	})
	return loaded, loadErr
}

// Parse validates a publication document; useful to staging callers and tests.
func Parse(data []byte) (*Catalog, error) {
	var doc Document
	if err := json.Unmarshal(data, &doc); err != nil {
		return nil, fmt.Errorf("catalog: decoding: %w", err)
	}
	if doc.Entries == nil || doc.Schemas == nil {
		return nil, fmt.Errorf("catalog: entries and schemas must be present")
	}
	sort.Slice(doc.Entries, func(i, j int) bool { return doc.Entries[i].OperationID < doc.Entries[j].OperationID })
	c := &Catalog{document: doc, byID: map[string]int{}}
	commands := map[string]bool{}
	for i, e := range doc.Entries {
		if strings.TrimSpace(e.OperationID) == "" || len(e.Command) == 0 || !strings.HasPrefix(e.Path, "/") {
			return nil, fmt.Errorf("catalog: entry %d has missing identity or command/path", i)
		}
		for _, token := range e.Command {
			if token == "" || strings.ContainsAny(token, " \t\r\n") {
				return nil, fmt.Errorf("catalog: invalid command token %q", token)
			}
		}
		switch e.Method {
		case "GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS":
		default:
			return nil, fmt.Errorf("catalog: %s: invalid method %q", e.OperationID, e.Method)
		}
		if e.ResponseEnvelope != "none" && e.ResponseEnvelope != "values" {
			return nil, fmt.Errorf("catalog: %s: invalid response envelope", e.OperationID)
		}
		if (e.ResponseEnvelope == "values") != e.Paginated {
			return nil, fmt.Errorf("catalog: %s: inconsistent pagination envelope", e.OperationID)
		}
		if _, exists := c.byID[e.OperationID]; exists {
			return nil, fmt.Errorf("catalog: duplicate operationId %q", e.OperationID)
		}
		path := strings.Join(e.Command, " ")
		if commands[path] {
			return nil, fmt.Errorf("catalog: duplicate command %q", path)
		}
		commands[path] = true
		c.byID[e.OperationID] = i
		flags := map[string]bool{}
		for _, p := range e.Params {
			if p.Flag == "" || (p.Source != "flag" && p.Source != "config" && p.Source != "pagination") || (p.In != "path" && p.In != "query") {
				return nil, fmt.Errorf("catalog: %s: invalid parameter", e.OperationID)
			}
			// A shared config org-id can describe both a path and query param.
			if flags[p.Flag] && p.Source != "config" {
				return nil, fmt.Errorf("catalog: %s: duplicate flag %q", e.OperationID, p.Flag)
			}
			flags[p.Flag] = true
		}
		for _, flag := range e.BodyFlags {
			if flag == "" || flags[flag] {
				return nil, fmt.Errorf("catalog: %s: duplicate/empty body flag %q", e.OperationID, flag)
			}
			flags[flag] = true
		}
	}
	// Walking the whole document validates nested refs and cycles without
	// expansion. Every referenced schema must exist, including unused schemas.
	var value interface{}
	if err := json.Unmarshal(data, &value); err != nil {
		return nil, err
	}
	if err := validateRefs(value, doc.Schemas); err != nil {
		return nil, err
	}
	return c, nil
}

func validateRefs(value interface{}, schemas map[string]json.RawMessage) error {
	switch v := value.(type) {
	case map[string]interface{}:
		if raw, ok := v["$ref"]; ok {
			ref, ok := raw.(string)
			name, valid := schemaName(ref)
			if !ok || !valid {
				return fmt.Errorf("catalog: unsupported schema ref %q", raw)
			}
			if _, exists := schemas[name]; !exists {
				return fmt.Errorf("catalog: dangling schema ref %q", ref)
			}
		}
		for key, nested := range v {
			// Examples are instance data, not schema definitions.
			if key == "example" || key == "examples" || key == "requestExample" {
				continue
			}
			if key == "properties" {
				if properties, ok := nested.(map[string]interface{}); ok {
					for _, property := range properties {
						if err := validateRefs(property, schemas); err != nil {
							return err
						}
					}
					continue
				}
			}
			if err := validateRefs(nested, schemas); err != nil {
				return err
			}
		}
	case []interface{}:
		for _, nested := range v {
			if err := validateRefs(nested, schemas); err != nil {
				return err
			}
		}
	}
	return nil
}

func schemaName(ref string) (string, bool) {
	const prefix = "#/components/schemas/"
	if !strings.HasPrefix(ref, prefix) {
		return "", false
	}
	name := strings.TrimPrefix(ref, prefix)
	if name == "" || strings.Contains(name, "/") {
		return "", false
	}
	return strings.NewReplacer("~1", "/", "~0", "~").Replace(name), true
}

// Accessors return copies so callers cannot mutate the shared lazy catalog.
func (c *Catalog) Lookup(operationID string) (Entry, bool) {
	i, ok := c.byID[operationID]
	if !ok {
		return Entry{}, false
	}
	data, _ := json.Marshal(c.document.Entries[i])
	var result Entry
	_ = json.Unmarshal(data, &result)
	return result, true
}

func (c *Catalog) All() []Entry {
	result := make([]Entry, 0, len(c.document.Entries))
	for _, e := range c.document.Entries {
		copy, _ := c.Lookup(e.OperationID)
		result = append(result, copy)
	}
	return result
}

// Schema accepts an OAS component schema ref and preserves its original JSON.
func (c *Catalog) Schema(ref string) (json.RawMessage, bool) {
	name, valid := schemaName(ref)
	if !valid {
		return nil, false
	}
	raw, ok := c.document.Schemas[name]
	return append(json.RawMessage(nil), raw...), ok
}

func Lookup(operationID string) (Entry, bool, error) {
	c, err := Load()
	if err != nil {
		return Entry{}, false, err
	}
	e, ok := c.Lookup(operationID)
	return e, ok, nil
}

func All() ([]Entry, error) {
	c, err := Load()
	if err != nil {
		return nil, err
	}
	return c.All(), nil
}

func Schema(ref string) (json.RawMessage, bool, error) {
	c, err := Load()
	if err != nil {
		return nil, false, err
	}
	s, ok := c.Schema(ref)
	return s, ok, nil
}
