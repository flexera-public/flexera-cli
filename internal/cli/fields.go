package cli

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/itchyny/gojq"
)

// FieldPath is a sequence of literal object keys, not a jq expression.
type FieldPath []string

// ParseFields parses comma-separated, dot-separated paths. Whitespace around
// each segment is ignored. Invalid selections carry command-line exit code 2.
func ParseFields(value string) ([]FieldPath, error) {
	var paths []FieldPath
	for _, field := range strings.Split(value, ",") {
		path := FieldPath(strings.Split(field, "."))
		for i := range path {
			path[i] = strings.TrimSpace(path[i])
		}
		paths = append(paths, path)
	}
	if err := validateFieldPaths(paths); err != nil {
		return nil, Exit(2, err)
	}
	return paths, nil
}

func validateFieldPaths(paths []FieldPath) error {
	if len(paths) == 0 {
		return fmt.Errorf("fields must not be empty")
	}
	for i, path := range paths {
		if len(path) == 0 {
			return fmt.Errorf("field path must not be empty")
		}
		for _, segment := range path {
			if strings.TrimSpace(segment) == "" {
				return fmt.Errorf("field path %q contains an empty segment", strings.Join(path, "."))
			}
		}
		for _, previous := range paths[:i] {
			common := 0
			for common < len(path) && common < len(previous) && path[common] == previous[common] {
				common++
			}
			if common == len(path) || common == len(previous) {
				return fmt.Errorf("duplicate or conflicting field paths %q and %q", strings.Join(previous, "."), strings.Join(path, "."))
			}
		}
	}
	return nil
}

type fieldNode struct {
	key      string
	path     FieldPath
	children []*fieldNode
}

// CompileFields projects each array element or a single ordinary object.
// For a known envelope, only values is projected; other metadata is retained.
// Nested paths produce nested objects, with null leaves for missing/null parents.
func CompileFields(paths []FieldPath, envelope bool) (*gojq.Code, error) {
	if err := validateFieldPaths(paths); err != nil {
		return nil, Exit(2, err)
	}
	root := &fieldNode{}
	for _, path := range paths {
		node := root
		for _, key := range path {
			var child *fieldNode
			for _, candidate := range node.children {
				if candidate.key == key {
					child = candidate
					break
				}
			}
			if child == nil {
				child = &fieldNode{key: key}
				node.children = append(node.children, child)
			}
			node = child
		}
		node.path = path
	}
	query := "def project: . as $row | " + root.expression() + "; " +
		"def rows: if type == \"array\" then map(project) elif . == null then null else project end; "
	if envelope {
		query += "if type == \"object\" and has(\"values\") then .values |= rows else rows end"
	} else {
		query += "rows"
	}
	parsed, err := gojq.Parse(query)
	if err != nil {
		return nil, err
	}
	return gojq.Compile(parsed)
}

func (node *fieldNode) expression() string {
	if node.path != nil {
		path, _ := json.Marshal(node.path)
		return "($row | getpath(" + string(path) + "))"
	}
	parts := make([]string, 0, len(node.children))
	for _, child := range node.children {
		key, _ := json.Marshal(child.key)
		parts = append(parts, string(key)+": "+child.expression())
	}
	return "{" + strings.Join(parts, ", ") + "}"
}
