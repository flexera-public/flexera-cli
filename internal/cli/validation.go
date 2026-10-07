package cli

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"sort"
	"strconv"
	"strings"

	"github.com/flexera-public/flexera-cli/internal/catalog"
	"github.com/getkin/kin-openapi/openapi3"
)

type ValidationResult struct {
	Status string `json:"status"`
}

// ParseRequestJSON accepts exactly one JSON value and retains number precision.
func ParseRequestJSON(raw []byte) (any, error) {
	d := json.NewDecoder(bytes.NewReader(raw))
	d.UseNumber()
	var value any
	if err := d.Decode(&value); err != nil {
		return nil, Exit(2, fmt.Errorf("request body must contain valid JSON"))
	}
	if err := d.Decode(new(any)); err != io.EOF {
		return nil, Exit(2, fmt.Errorf("request body must contain exactly one JSON value"))
	}
	return value, nil
}

// DecodeResponseJSON preserves raw wire numbers even where SDK interface{}
// response decoding would round them before the printer sees the value.
func DecodeResponseJSON(raw []byte) (any, error) {
	d := json.NewDecoder(bytes.NewReader(raw))
	d.UseNumber()
	var value any
	if err := d.Decode(&value); err != nil {
		return nil, fmt.Errorf("response body is not valid JSON")
	}
	if err := d.Decode(new(any)); err != io.EOF {
		return nil, fmt.Errorf("response body must contain exactly one JSON value")
	}
	return value, nil
}

// ValidateRequestJSON checks the original body, before typed decoding. skip
// bypasses schema constraints only; syntax checking always runs.
func ValidateRequestJSON(index *catalog.Catalog, entry catalog.Entry, raw []byte, skip bool) (*ValidationResult, error) {
	value, err := ParseRequestJSON(raw)
	if err != nil {
		return nil, err
	}
	if skip {
		return &ValidationResult{Status: "skipped"}, nil
	}
	schema, err := index.RequestSchema(entry.RequestSchema)
	if err != nil {
		return nil, Exit(2, err)
	}
	if err := validateSchemaValue(schema, value); err != nil {
		details := validationDetails(err)
		return nil, Exit(2, &ValidationError{Details: details, Schema: "flexera-cli cli schema " + strings.Join(entry.Command, " ")})
	}
	return &ValidationResult{Status: "ok"}, nil
}

func validationDetails(err error) []ValidationDetail {
	details := []ValidationDetail{}
	var visit func(error)
	visit = func(err error) {
		if multiple, ok := err.(openapi3.MultiError); ok {
			for _, item := range multiple {
				visit(item)
			}
			return
		}
		var schema *openapi3.SchemaError
		if errors.As(err, &schema) {
			path := ""
			for _, segment := range schema.JSONPointer() {
				path += "/" + pointerSegment(segment)
			}
			message := schema.Reason
			if message == "" {
				message = "request value violates schema constraint " + schema.SchemaField
			}
			details = append(details, ValidationDetail{Path: path, Message: message})
			return
		}
		details = append(details, ValidationDetail{Path: "", Message: "request value does not satisfy its schema"})
	}
	visit(err)
	sort.Slice(details, func(i, j int) bool {
		if details[i].Path == details[j].Path {
			return details[i].Message < details[j].Message
		}
		return details[i].Path < details[j].Path
	})
	return details
}

func pointerSegment(key string) string { return strings.NewReplacer("~", "~0", "/", "~1").Replace(key) }

// DecodeRequestBody refuses to send a body whose supplied properties/values
// would be discarded or changed by the generated request type. The caller owns
// target and must not use it after an error. skip never disables this safeguard.
func DecodeRequestBody(index *catalog.Catalog, entry catalog.Entry, raw []byte, target any, skip bool) (*ValidationResult, error) {
	result, err := ValidateRequestJSON(index, entry, raw, skip)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(raw, target); err != nil {
		return nil, Exit(2, fmt.Errorf("request body cannot be decoded into the API request type"))
	}
	encoded, err := json.Marshal(target)
	if err != nil {
		return nil, Exit(2, fmt.Errorf("API request type cannot be encoded"))
	}
	before, err := ParseRequestJSON(raw)
	if err != nil {
		return nil, err
	}
	after, err := ParseRequestJSON(encoded)
	if err != nil {
		return nil, err
	}
	paths := changedInputs(before, after, "")
	if len(paths) > 0 {
		details := make([]ValidationDetail, 0, len(paths))
		for _, path := range paths {
			details = append(details, ValidationDetail{Path: path, Message: "supplied input would be lost or changed by the API request type; remove or correct this input, or update the client type"})
		}
		return nil, Exit(2, &ValidationError{Details: details, Schema: "flexera-cli cli schema " + strings.Join(entry.Command, " ")})
	}
	return result, nil
}

func changedInputs(before, after any, path string) []string {
	switch value := before.(type) {
	case map[string]any:
		object, ok := after.(map[string]any)
		if !ok {
			return []string{path}
		}
		keys := make([]string, 0, len(value))
		for key := range value {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		paths := []string{}
		for _, key := range keys {
			child := path + "/" + pointerSegment(key)
			item, exists := object[key]
			if !exists {
				paths = append(paths, child)
			} else {
				paths = append(paths, changedInputs(value[key], item, child)...)
			}
		}
		return paths
	case []any:
		array, ok := after.([]any)
		if !ok || len(value) != len(array) {
			return []string{path}
		}
		paths := []string{}
		for i, item := range value {
			paths = append(paths, changedInputs(item, array[i], path+"/"+strconv.Itoa(i))...)
		}
		return paths
	case json.Number:
		number, ok := after.(json.Number)
		if !ok {
			return []string{path}
		}
		left, lok := new(big.Rat).SetString(string(value))
		right, rok := new(big.Rat).SetString(string(number))
		if !lok || !rok || left.Cmp(right) != 0 {
			return []string{path}
		}
	case string:
		other, ok := after.(string)
		if !ok || value != other {
			return []string{path}
		}
	case bool:
		other, ok := after.(bool)
		if !ok || value != other {
			return []string{path}
		}
	case nil:
		if after != nil {
			return []string{path}
		}
	default:
		return []string{path}
	}
	return nil
}
