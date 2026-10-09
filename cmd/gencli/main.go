// Command gencli generates one cobra command package per OpenAPI tag from
// the annotated unified OpenAPI spec.
//
// Usage:
//
//	go run ./cmd/gencli -spec unified-openapi/openapi3.json \
//	    -tag Budget -pkg budget -cmd budget-gen -out internal/commands/budget/cmd_gen.go
//
// The generated file declares, in its own package:
//
//   - NewCmd() *cobra.Command — the tag command with one child per operation.
//   - one cobra command constructor per supported operation.
//
// Each leaf RunE pulls shared config/auth/printer from the command context
// (*cli.Deps) populated by the root command's PersistentPreRunE; it only
// registers operation-specific flags (path/query params + body fields, with
// a --body escape hatch). Write bodies use typed per-field flags where the
// request schema is simple, falling back to --body for complex schemas.
//
// Only operations annotated with x-flexera-action ∈ {list, get, create,
// replace, update, delete, action} are emitted. Operations with non-JSON
// bodies or array params of non-string type are
// skipped — every skip is recorded as a drop, printed to stderr, and
// causes a non-zero exit so silent spec→CLI drift fails the regen build.
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/template"
)

type operation struct {
	Method               string
	Path                 string
	OperationID          string
	Summary              string
	Description          string
	RequestSchema        map[string]interface{}
	ResponseSchema       map[string]interface{}
	RequestExample       interface{}
	RequestExampleSource string
	SuccessResponses     []successResponse
	Tag                  string
	Action               string
	Resource             string
	PathParams           []param
	QueryParams          []param
	HasHeaders           bool
	HasBody              bool
	HasRawBody           bool
	RawBodyType          string
	BodyTypeName         string // generated client's "<Method>JSONRequestBody"
	BodyFields           []bodyField
	Paginated            bool
	Success2xx           string // lowest documented numeric 2xx
	HasJSONResp          bool
	HasSchemaResp        bool
}

type param struct {
	Name        string // raw spec name (orgId, id, project_id)
	GoName      string // Go-safe lowerCamelCase (orgID, id, projectID)
	FlagName    string // kebab-case CLI flag (org-id, id, project-id)
	GoType      string // SDK type: string|int64|int|bool|[]string|time.Time|openapi_types.Date
	TimeLayout  string // non-empty for scalar date/date-time query params; CLI input is a string
	IsArray     bool
	IsInt64     bool
	IsUUID      bool
	IsEnum      bool
	EnumType    string
	In          string // path|query|header
	Required    bool
	Schema      map[string]interface{}
	Description string
	Source      string // flag|config|pagination
	ValueExpr   string // config-backed query value
	Metadata    map[string]interface{}
}

type successResponse struct {
	MediaType string
	Code      string
	HasJSON   bool
	HasSchema bool
	Schema    map[string]interface{}
}

// bodyField is a scalar (or scalar-array) top-level property of a JSON
// request body, surfaced as a typed cobra flag.
type bodyField struct {
	JSONKey  string // spec property name, used as the JSON map key
	GoVar    string // local Go variable name
	FlagName string // kebab-case CLI flag
	GoType   string // string|int|int64|float64|bool|[]string
	Required bool
	Schema   map[string]interface{}
}

type spec struct {
	Paths      map[string]json.RawMessage `json:"paths"`
	Components struct {
		Schemas map[string]json.RawMessage `json:"schemas"`
	} `json:"components"`
}

type pathItem map[string]json.RawMessage

func main() {
	var (
		specPath     = flag.String("spec", "", "path to annotated unified openapi.json")
		tag          = flag.String("tag", "", "OpenAPI tag to generate (e.g. Budget)")
		pkg          = flag.String("pkg", "", "Go package name for the generated file")
		cmdName      = flag.String("cmd", "", "root CLI command name (e.g. budget-gen)")
		outPath      = flag.String("out", "", "output Go file path")
		metadataPath = flag.String("metadata", "", "optional per-tag operation metadata JSON output (keyed by operationId; command paths are provisional)")
		service      = flag.String("service", "", "optional x-flexera-service id; only operations stamped with this service are generated")
		exclude      = flag.String("exclude", "", "optional comma-separated operationIds owned by curated commands (not generated)")
	)
	flag.Parse()
	if *specPath == "" || *tag == "" || *outPath == "" || *pkg == "" || *cmdName == "" {
		fmt.Fprintln(os.Stderr, "usage: gencli -spec <openapi.json> -tag <Tag> -pkg <pkg> -cmd <name> -out <file.go>")
		os.Exit(2)
	}

	data, err := os.ReadFile(*specPath)
	check(err)

	var s spec
	check(json.Unmarshal(data, &s))

	filter := opFilter{Service: *service, Exclude: map[string]bool{}}
	for _, id := range strings.Split(*exclude, ",") {
		if id = strings.TrimSpace(id); id != "" {
			filter.Exclude[id] = true
		}
	}
	ops, drops := collectOpsFiltered(&s, *tag, filter)
	for _, d := range drops {
		fmt.Fprintf(os.Stderr, "gencli: skip %s %s — %s\n", strings.ToUpper(d.Method), d.Path, d.Reason)
	}
	if len(ops) == 0 {
		fmt.Fprintf(os.Stderr, "no operations found for tag %q\n", *tag)
		os.Exit(1)
	}

	src, metadata, err := renderWithMetadata(*tag, *pkg, *cmdName, ops)
	check(err)

	formatted, err := format.Source(src)
	if err != nil {
		_ = os.WriteFile(*outPath+".raw", src, 0o644)
		check(fmt.Errorf("gofmt failed (raw source written to %s.raw): %w", *outPath, err))
	}

	check(os.MkdirAll(dirOf(*outPath), 0o755))
	check(os.WriteFile(*outPath, formatted, 0o644))
	if *metadataPath != "" {
		data, err := json.MarshalIndent(metadata, "", "  ")
		check(err)
		check(os.MkdirAll(dirOf(*metadataPath), 0o755))
		check(os.WriteFile(*metadataPath, append(data, '\n'), 0o644))
	}
	fmt.Fprintf(os.Stderr, "wrote %d ops to %s\n", len(ops), *outPath)

	if len(drops) > 0 {
		fmt.Fprintf(os.Stderr, "gencli: %d operation(s) for tag %q were skipped (see stderr above); failing regen build\n", len(drops), *tag)
		os.Exit(1)
	}
}

func dirOf(p string) string {
	if i := strings.LastIndexByte(p, '/'); i >= 0 {
		return p[:i]
	}
	return "."
}

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "gencli:", err)
		os.Exit(1)
	}
}

// supportedActions enumerates every x-flexera-action verb gencli emits
// a cobra leaf for. "action" covers RPC-style POST sub-action paths
// (e.g. /policy/v1/orgs/{orgId}/applied-policies/{id}/evaluate); the
// leaf's CLI verb is derived from the trailing path literal by
// assignVerbs, not the literal string "action". ("query" was a reserved
// future verb in the design's Architecture but no annotateForCLI branch
// emits it, so it is dropped here per design Q3 / S3 plan-review.)
var supportedActions = map[string]bool{
	"list": true, "get": true, "create": true,
	"replace": true, "update": true, "delete": true,
	"action": true,
}

// drop records a single operation that gencli refused to emit. main
// prints every drop to stderr and exits non-zero so silent spec→CLI
// drift fails the regen build instead of the consumer build.
type drop struct {
	Method string
	Path   string
	Action string
	Reason string
}

// opFilter narrows a tag to one service (tags such as "Project" are shared
// across services) and skips operations owned by curated commands.
type opFilter struct {
	Service string
	Exclude map[string]bool
}

func collectOps(s *spec, tag string) ([]operation, []drop) {
	return collectOpsFiltered(s, tag, opFilter{})
}

func collectOpsFiltered(s *spec, tag string, filter opFilter) ([]operation, []drop) {
	var ops []operation
	var drops []drop
	methods := []string{"get", "put", "post", "delete", "patch"}

	pathKeys := make([]string, 0, len(s.Paths))
	for k := range s.Paths {
		pathKeys = append(pathKeys, k)
	}
	sort.Strings(pathKeys)

	for _, p := range pathKeys {
		var pi pathItem
		if err := json.Unmarshal(s.Paths[p], &pi); err != nil {
			continue
		}
		var pathLevelParams []interface{}
		if rawParams, ok := pi["parameters"]; ok {
			var arr []interface{}
			_ = decodeSpecJSON(rawParams, &arr)
			pathLevelParams = arr
		}
		for _, m := range methods {
			raw, ok := pi[m]
			if !ok {
				continue
			}
			var op map[string]interface{}
			if err := decodeSpecJSON(raw, &op); err != nil {
				continue
			}
			if len(pathLevelParams) > 0 {
				opParams, _ := op["parameters"].([]interface{})
				merged := make([]interface{}, 0, len(pathLevelParams)+len(opParams))
				merged = append(merged, pathLevelParams...)
				merged = append(merged, opParams...)
				op["parameters"] = merged
			}
			tags, _ := op["tags"].([]interface{})
			matched := false
			for _, t := range tags {
				if ts, _ := t.(string); ts == tag {
					matched = true
					break
				}
			}
			if !matched {
				continue
			}
			if service, _ := op["x-flexera-service"].(string); filter.Service != "" && service != filter.Service {
				continue
			}
			if id, _ := op["operationId"].(string); filter.Exclude[id] {
				continue
			}
			action, _ := op["x-flexera-action"].(string)
			if !supportedActions[action] {
				drops = append(drops, drop{Method: m, Path: p, Action: action, Reason: fmt.Sprintf("unsupported x-flexera-action %q", action)})
				continue
			}
			opID, _ := op["operationId"].(string)
			if opID == "" {
				drops = append(drops, drop{Method: m, Path: p, Action: action, Reason: "missing operationId"})
				continue
			}
			resource, _ := op["x-flexera-resource"].(string)
			summary, _ := op["summary"].(string)
			description, _ := op["description"].(string)
			pag, _ := op["x-flexera-paginated"].(bool)

			pps, qps, hasHeaders, ok := extractParams(op)
			if !ok {
				drops = append(drops, drop{Method: m, Path: p, Action: action, Reason: "param extraction failed (unsupported param shape)"})
				continue
			}

			hasBody := false
			hasRawBody := false
			rawBodyType := ""
			hasUnsupportedBody := false
			var bodyFields []bodyField
			var requestSchema map[string]interface{}
			var requestExample interface{}
			requestExampleSource := ""
			if rb, ok := op["requestBody"].(map[string]interface{}); ok {
				if content, ok := rb["content"].(map[string]interface{}); ok {
					if jc, ok := content["application/json"].(map[string]interface{}); ok {
						hasBody = true
						bodyFields = extractBodyFields(s, jc)
						requestSchema, _ = jc["schema"].(map[string]interface{})
						requestExample = bodyExample(s, jc)
						requestExampleSource = bodyExampleSource(jc)
					} else if octets, ok := content["application/octet-stream"].(map[string]interface{}); ok {
						hasRawBody = true
						rawBodyType = "application/octet-stream"
						requestSchema, _ = octets["schema"].(map[string]interface{})
					} else if len(content) > 0 {
						hasUnsupportedBody = true
					}
				}
			}
			if hasUnsupportedBody {
				drops = append(drops, drop{Method: m, Path: p, Action: action, Reason: "non-JSON request body"})
				continue
			}

			success, hasJSON, hasSchema := successCode(op)
			if success == "" {
				drops = append(drops, drop{Method: m, Path: p, Action: action, Reason: "no 2xx success response"})
				continue
			}

			ops = append(ops, operation{
				Description:          description,
				RequestSchema:        requestSchema,
				RequestExample:       requestExample,
				RequestExampleSource: requestExampleSource,
				ResponseSchema:       successResponses(op)[0].Schema,
				SuccessResponses:     successResponses(op),
				Method:               m,
				Path:                 p,
				OperationID:          opID,
				Summary:              summary,
				Tag:                  tag,
				Action:               action,
				Resource:             resource,
				PathParams:           pps,
				QueryParams:          qps,
				HasHeaders:           hasHeaders,
				HasBody:              hasBody,
				HasRawBody:           hasRawBody,
				RawBodyType:          rawBodyType,
				BodyTypeName:         methodName(opID) + "JSONRequestBody",
				BodyFields:           bodyFields,
				Paginated:            pag,
				Success2xx:           success,
				HasJSONResp:          hasJSON,
				HasSchemaResp:        hasSchema,
			})
		}
	}
	return ops, drops
}

// extractBodyFields resolves the JSON request body schema (following a single
// $ref into components.schemas) and returns typed flags for its scalar /
// scalar-array top-level properties. It returns nil (typed mode off) when the
// schema is complex: oneOf/anyOf/allOf, a non-object, or has no scalar
// top-level properties.
// reservedFlag reports whether a flag name is reserved for CLI control
// (so a body-field flag must not shadow it).
func reservedFlag(name string) bool {
	switch name {
	case "body", "dry-run", "yes", "no-paginate", "skip-token",
		"config", "zone", "api-base-url", "login-base-url", "output", "json-style",
		"access-token", "client-id", "client-secret", "refresh-token", "org-id", "debug",
		"help", "version", "out-jq", "out-fields", "raw-output", "no-validate", "interactive":
		return true
	}
	return false
}

func extractBodyFields(s *spec, jsonContent map[string]interface{}) []bodyField {
	schema, _ := jsonContent["schema"].(map[string]interface{})
	if schema == nil {
		return nil
	}
	schema = resolveSchemaRef(s, schema, 0)
	if schema == nil {
		return nil
	}
	for _, k := range []string{"oneOf", "anyOf", "allOf"} {
		if _, ok := schema[k]; ok {
			return nil
		}
	}
	if t, _ := schema["type"].(string); t != "object" && t != "" {
		return nil
	}
	props, _ := schema["properties"].(map[string]interface{})
	if len(props) == 0 {
		return nil
	}
	keys := make([]string, 0, len(props))
	for k := range props {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var fields []bodyField
	used := map[string]bool{}
	// Reserve natural spellings first so a fallback cannot steal another property.
	for _, name := range keys {
		used[kebab(name)] = true
	}
	for _, name := range keys {
		ps, _ := props[name].(map[string]interface{})
		if ps == nil {
			continue
		}
		if readOnly, _ := ps["readOnly"].(bool); readOnly {
			continue
		}
		if _, ref := ps["$ref"]; ref {
			continue // nested object reference — leave to --body
		}
		t, _ := ps["type"].(string)
		var goType string
		switch t {
		case "string":
			goType = "string"
		case "integer":
			if f, _ := ps["format"].(string); f == "int64" {
				goType = "int64"
			} else {
				goType = "int"
			}
		case "number":
			goType = "float64"
		case "boolean":
			goType = "bool"
		case "array":
			items, _ := ps["items"].(map[string]interface{})
			it, _ := items["type"].(string)
			if it != "string" {
				continue // array of non-string / objects → --body
			}
			goType = "[]string"
		default:
			continue // object / unknown → --body
		}
		flag := kebab(name)
		if reservedFlag(flag) {
			flag = "body-" + flag
			for used[flag] || reservedFlag(flag) {
				flag = "body-" + flag
			}
		}
		used[flag] = true
		fields = append(fields, bodyField{
			JSONKey:  name,
			GoVar:    "f" + pascal(goIdent(name)),
			FlagName: flag,
			GoType:   goType,
			Required: containsString(schema["required"], name),
			Schema:   ps,
		})
	}
	return fields
}

func resolveSchemaRef(s *spec, schema map[string]interface{}, depth int) map[string]interface{} {
	if depth > 5 {
		return schema
	}
	ref, _ := schema["$ref"].(string)
	if ref == "" {
		return schema
	}
	const prefix = "#/components/schemas/"
	if !strings.HasPrefix(ref, prefix) {
		return nil
	}
	name := strings.TrimPrefix(ref, prefix)
	raw, ok := s.Components.Schemas[name]
	if !ok {
		return nil
	}
	var next map[string]interface{}
	if err := decodeSpecJSON(raw, &next); err != nil {
		return nil
	}
	return resolveSchemaRef(s, next, depth+1)
}

func extractParams(op map[string]interface{}) (path, query []param, hasHeaders bool, ok bool) {
	raw, _ := op["parameters"].([]interface{})
	for _, p := range raw {
		m, _ := p.(map[string]interface{})
		if m == nil {
			continue
		}
		name, _ := m["name"].(string)
		in, _ := m["in"].(string)
		required, _ := m["required"].(bool)
		schema, _ := m["schema"].(map[string]interface{})
		t, _ := schema["type"].(string)
		format, _ := schema["format"].(string)
		_, hasEnum := schema["enum"]
		if in == "header" {
			hasHeaders = true
			continue
		}
		var goType string
		isArray := false
		isInt64 := false
		isEnum := false
		timeLayout := ""
		switch t {
		case "integer":
			if format == "int64" {
				goType = "int64"
				isInt64 = true
			} else {
				goType = "int"
			}
			if hasEnum {
				isEnum = true
			}
		case "boolean":
			goType = "bool"
		case "string", "":
			goType = "string"
			if format == "date-time" || format == "date" {
				// Only scalar query values are supported here. Formatted enums
				// and path values require their own SDK type/call handling.
				if in != "query" || hasEnum {
					return nil, nil, false, false
				}
				goType, timeLayout = "time.Time", "2006-01-02T15:04:05.999999999Z07:00"
				if format == "date" {
					goType, timeLayout = "openapi_types.Date", "2006-01-02"
				}
			}
			if hasEnum {
				isEnum = true
			}
		case "array":
			items, _ := schema["items"].(map[string]interface{})
			it, _ := items["type"].(string)
			if it != "string" && it != "" {
				return nil, nil, false, false
			}
			if f, _ := items["format"].(string); f == "date" || f == "date-time" {
				return nil, nil, false, false
			}
			goType = "[]string"
			isArray = true
			if _, ok := items["enum"]; ok {
				isEnum = true
			}
		default:
			return nil, nil, false, false
		}
		pp := param{
			Metadata:    m,
			Schema:      schema,
			Description: stringValue(m["description"]),
			Source:      "flag",
			Name:        name,
			GoName:      goIdent(name),
			FlagName:    kebab(name),
			GoType:      goType,
			TimeLayout:  timeLayout,
			IsArray:     isArray,
			IsInt64:     isInt64,
			IsUUID:      format == "uuid",
			IsEnum:      isEnum,
			In:          in,
			Required:    required,
		}
		switch in {
		case "path":
			path = append(path, pp)
		case "query":
			query = append(query, pp)
		}
	}
	return path, query, hasHeaders, true
}

func successCode(op map[string]interface{}) (code string, hasJSON, hasSchema bool) {
	responses := successResponses(op)
	if len(responses) == 0 {
		return "", false, false
	}
	r := responses[0]
	return r.Code, r.HasJSON, r.HasSchema
}

// successResponses reads OAS 3 content.schema, not the OAS 2 response.schema.
// Numeric status codes have a stable order; wildcard/default responses cannot
// name a generated SDK response field and are not selected.
func successResponses(op map[string]interface{}) []successResponse {
	responses, _ := op["responses"].(map[string]interface{})
	var codes []string
	for c := range responses {
		n, err := strconv.Atoi(c)
		if err == nil && len(c) == 3 && n >= 200 && n < 300 {
			codes = append(codes, c)
		}
	}
	sort.Strings(codes)
	var result []successResponse
	for _, c := range codes {
		resp, ok := responses[c].(map[string]interface{})
		if !ok {
			continue
		}
		content, _ := resp["content"].(map[string]interface{})
		keys := sortedKeys(content)
		// Prefer JSON for schema metadata when multiple media types exist.
		sort.SliceStable(keys, func(i, j int) bool {
			return strings.Contains(keys[i], "json") && !strings.Contains(keys[j], "json")
		})
		r := successResponse{Code: c}
		for _, ct := range keys {
			media, _ := content[ct].(map[string]interface{})
			if r.MediaType == "" {
				r.MediaType = ct
				r.HasJSON = strings.Contains(ct, "json")
			}
			schema, _ := media["schema"].(map[string]interface{})
			if schema != nil {
				r.MediaType = ct
				r.Schema = schema
				r.HasSchema = true
				r.HasJSON = strings.Contains(ct, "json")
				break
			}
		}
		result = append(result, r)
	}
	return result
}

func sortedKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func stringValue(v interface{}) string { s, _ := v.(string); return s }

func decodeSpecJSON(raw []byte, target interface{}) error {
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.UseNumber()
	return decoder.Decode(target)
}

// Catalog examples are illustrative, not validated requests. Help deliberately
// uses @request.json instead of printing spec values that could contain secrets.
func bodyExample(s *spec, media map[string]interface{}) interface{} {
	if v, ok := media["example"]; ok {
		return v
	}

	examples, _ := media["examples"].(map[string]interface{})
	for _, key := range sortedKeys(examples) {
		example, _ := examples[key].(map[string]interface{})
		if v, ok := example["value"]; ok {
			return v
		}
	}
	schema, _ := media["schema"].(map[string]interface{})
	v, _ := synthesizeExample(s, schema, 0, map[string]bool{})
	return v
}

func bodyExampleSource(media map[string]interface{}) string {
	if _, exists := media["example"]; exists {
		return "upstream"
	}
	examples, _ := media["examples"].(map[string]interface{})
	for _, key := range sortedKeys(examples) {
		example, _ := examples[key].(map[string]interface{})
		if _, exists := example["value"]; exists {
			return "upstream"
		}
	}
	return "synthesized"
}

func synthesizeExample(s *spec, schema map[string]interface{}, depth int, seen map[string]bool) (interface{}, bool) {
	if schema == nil || depth > 4 {
		return nil, false
	}
	if ref := stringValue(schema["$ref"]); ref != "" {
		if seen[ref] || !strings.HasPrefix(ref, "#/components/schemas/") {
			return nil, false
		}
		var resolved map[string]interface{}
		if err := decodeSpecJSON(s.Components.Schemas[strings.TrimPrefix(ref, "#/components/schemas/")], &resolved); err != nil {
			return nil, false
		}
		seen[ref] = true
		defer delete(seen, ref)
		return synthesizeExample(s, resolved, depth, seen)
	}
	if schema["readOnly"] == true {
		return nil, false
	}
	for _, key := range []string{"example", "default"} {
		if v, ok := schema[key]; ok {
			return v, true
		}
	}
	if enum, _ := schema["enum"].([]interface{}); len(enum) > 0 {
		return enum[0], true
	}
	for _, key := range []string{"oneOf", "anyOf"} {
		if variants, _ := schema[key].([]interface{}); len(variants) > 0 {
			variant, _ := variants[0].(map[string]interface{})
			return synthesizeExample(s, variant, depth+1, seen)
		}
	}
	if variants, ok := schema["allOf"].([]interface{}); ok && len(variants) > 0 {
		merged := map[string]interface{}{}
		for _, variant := range variants {
			child, _ := variant.(map[string]interface{})
			value, ok := synthesizeExample(s, child, depth+1, seen)
			if !ok {
				return nil, false
			}
			object, ok := value.(map[string]interface{})
			if !ok {
				return nil, false
			}
			for key, item := range object {
				merged[key] = item
			}
		}
		return merged, true
	}
	switch stringValue(schema["type"]) {
	case "array":
		items, _ := schema["items"].(map[string]interface{})
		v, ok := synthesizeExample(s, items, depth+1, seen)
		if !ok {
			return nil, false
		}
		return []interface{}{v}, true
	case "string":
		switch stringValue(schema["format"]) {
		case "date":
			return "2000-01-01", true
		case "date-time":
			return "2000-01-01T00:00:00Z", true
		case "uuid":
			return "00000000-0000-0000-0000-000000000000", true
		case "email":
			return "user@example.invalid", true
		case "uri":
			return "https://example.invalid", true
		}
		return "REPLACE_ME", true
	case "integer", "number":
		return 0, true
	case "boolean":
		return false, true
	}
	props, _ := schema["properties"].(map[string]interface{})
	if props != nil || schema["type"] == "object" {
		result := map[string]interface{}{}
		required, _ := schema["required"].([]interface{})
		for _, key := range required {
			name, _ := key.(string)
			property, _ := props[name].(map[string]interface{})
			if property["readOnly"] == true {
				continue
			}
			if v, ok := synthesizeExample(s, property, depth+1, seen); ok {
				result[name] = v
			} else {
				return nil, false
			}
		}
		return result, true
	}
	return nil, false
}

func methodName(opID string) string {
	parts := strings.Split(opID, "_")
	var b strings.Builder
	for _, p := range parts {
		if p == "" {
			continue
		}
		b.WriteString(strings.ToUpper(p[:1]))
		b.WriteString(p[1:])
	}
	return b.String()
}

func goIdent(s string) string {
	s = strings.ReplaceAll(s, "-", "_")
	parts := strings.Split(s, "_")
	for i, p := range parts {
		if p == "" {
			continue
		}
		if i == 0 {
			parts[i] = strings.ToLower(p[:1]) + p[1:]
		} else {
			parts[i] = strings.ToUpper(p[:1]) + p[1:]
		}
	}
	out := strings.Join(parts, "")
	out = strings.ReplaceAll(out, "Id", "ID")
	out = strings.ReplaceAll(out, "Url", "URL")
	if isGoKeyword(out) {
		out = out + "_"
	}
	return out
}

func isGoKeyword(s string) bool {
	switch s {
	case "break", "case", "chan", "const", "continue", "default", "defer",
		"else", "fallthrough", "for", "func", "go", "goto", "if", "import",
		"interface", "map", "package", "range", "return", "select", "struct",
		"switch", "type", "var":
		return true
	}
	return false
}

func kebab(s string) string {
	s = strings.ReplaceAll(s, "_", "-")
	runes := []rune(s)
	var b strings.Builder
	for i, r := range runes {
		isUpper := r >= 'A' && r <= 'Z'
		if i > 0 && isUpper {
			prev := runes[i-1]
			prevUpper := prev >= 'A' && prev <= 'Z'
			nextLower := i+1 < len(runes) && runes[i+1] >= 'a' && runes[i+1] <= 'z'
			if !prevUpper || (prevUpper && nextLower) {
				b.WriteByte('-')
			}
		}
		if isUpper {
			b.WriteRune(r + 32)
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func cleanTag(t string) string {
	tmp := strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			return r
		}
		return ' '
	}, t)
	parts := strings.Fields(tmp)
	for i, p := range parts {
		if p == "" {
			continue
		}
		parts[i] = strings.ToUpper(p[:1]) + p[1:]
	}
	return strings.Join(parts, "")
}

type renderData struct {
	Tag      string
	Pkg      string
	CmdName  string
	Short    string
	TagIdent string
	Ops      []renderOp
	HasTime  bool
	HasDate  bool
}

type renderOp struct {
	OutputContract    string
	Metadata          map[string]interface{}
	Example           string
	SuccessResponses  []successResponse
	ConfigQueryParams []param
	Constructor       string
	Verb              string
	Short             string
	Method            string
	Path              string
	OperationID       string
	GenMethod         string
	PathParams        []param // non-orgId path params (flags)
	QueryParams       []param // excludes skipToken when paginated
	PathArgExprs      []string
	UUIDPathParams    []param // subset of PathParams with IsUUID, for pre-call parsing
	HasUUIDParams     bool
	HasBody           bool
	HasRawBody        bool
	RawBodyType       string
	BodyTypeName      string
	BodyFields        []bodyField
	NeedsOrgID        bool
	ParamsType        string
	CanPaginate       bool
	Success2xx        string
	HasJSONResp       bool
	HasSchemaResp     bool
	HasRawJSONResp    bool
	IsWrite           bool
	IsDestructive     bool
}

func isWriteAction(a string) bool {
	switch a {
	case "create", "update", "replace", "delete":
		return true
	}
	return false
}

// isDestructiveVerb flags verbs that destructively mutate/remove even though
// their HTTP action is not DELETE (e.g. PUT .../revoke). These require --yes.
func isDestructiveVerb(v string) bool {
	switch v {
	case "revoke", "revoke-all", "deactivate", "delete-all":
		return true
	}
	return false
}

func render(tag, pkg, cmdName string, ops []operation) ([]byte, error) {
	src, _, err := renderWithMetadata(tag, pkg, cmdName, ops)
	return src, err
}

func renderWithMetadata(tag, pkg, cmdName string, ops []operation) ([]byte, map[string]interface{}, error) {
	ops = append([]operation(nil), ops...)
	tagIdent := cleanTag(tag)
	data := renderData{
		Tag:      tag,
		Pkg:      pkg,
		CmdName:  cmdName,
		Short:    tag + " operations (generated from the unified OpenAPI spec)",
		TagIdent: tagIdent,
	}

	sort.Slice(ops, func(i, j int) bool {
		if ops[i].Action != ops[j].Action {
			return ops[i].Action < ops[j].Action
		}
		return ops[i].Path < ops[j].Path
	})

	verbs := assignVerbs(ops)

	for idx, o := range ops {
		verb := verbs[idx]

		genMethod := methodName(o.OperationID)

		var uuidPathParams []param
		hasUUIDParams := false
		needsOrgID := false
		var pathFlags []param
		var pathArgExprs []string
		for _, pp := range o.PathParams {
			if pp.Name == "orgId" || pp.Name == "org_id" {
				needsOrgID = true
				switch {
				case pp.GoType == "string":
					pathArgExprs = append(pathArgExprs, "fmt.Sprint(deps.Config.OrgID)")
				case pp.IsInt64:
					pathArgExprs = append(pathArgExprs, "int64(deps.Config.OrgID)")
				default:
					pathArgExprs = append(pathArgExprs, "deps.Config.OrgID")
				}
				continue
			}
			f := pp
			if f.IsEnum {
				f.EnumType = "flexera." + genMethod + "Params" + pascal(f.Name)
				pathArgExprs = append(pathArgExprs, f.EnumType+"("+f.GoName+")")
			} else if f.IsUUID {
				pathArgExprs = append(pathArgExprs, f.GoName+"UUID")
				uuidPathParams = append(uuidPathParams, f)
				hasUUIDParams = true
			} else {
				pathArgExprs = append(pathArgExprs, f.GoName)
			}
			pathFlags = append(pathFlags, f)
		}

		// Pagination only when there's a skipToken query param to drive.
		canPaginate := false
		var queryFlags []param
		var configQueryParams []param
		for _, qp := range o.QueryParams {
			data.HasTime = data.HasTime || qp.TimeLayout != ""
			data.HasDate = data.HasDate || qp.GoType == "openapi_types.Date"
			if qp.FlagName == "org-id" {
				needsOrgID = needsOrgID || qp.Required
				q := qp
				q.Source = "config"
				q.ValueExpr = "deps.Config.OrgID"
				if q.GoType == "string" {
					q.ValueExpr = "fmt.Sprint(deps.Config.OrgID)"
				}
				if q.GoType == "int64" {
					q.ValueExpr = "int64(deps.Config.OrgID)"
				}
				if q.IsEnum {
					q.ValueExpr = "flexera." + genMethod + "Params" + pascal(q.Name) + "(" + q.ValueExpr + ")"
				}
				configQueryParams = append(configQueryParams, q)
				continue
			}
			if o.Paginated && qp.Name == "skipToken" {
				canPaginate = true
				continue // driven by CollectPages, not a normal flag
			}
			q := qp
			if q.IsEnum {
				q.EnumType = "flexera." + genMethod + "Params" + pascal(q.Name)
			}
			queryFlags = append(queryFlags, q)
		}

		paramsType := ""
		if len(o.QueryParams) > 0 || o.HasHeaders {
			paramsType = "flexera." + genMethod + "Params"
		}

		// Path/query parameters take precedence over body flags. Keep JSON keys
		// unchanged, and repeatedly prefix a fallback until it is unused.
		used := map[string]bool{}
		for i := range pathFlags {
			f := &pathFlags[i]
			if f.FlagName == "client-id" && (cmdName == "service-account-client" || kebab(tag) == "service-account-client") {
				f.FlagName = "target-client-id"
			} else if reservedFlag(f.FlagName) {
				f.FlagName = "target-" + f.FlagName
			}
			used[f.FlagName] = true
		}
		for i := range queryFlags {
			f := &queryFlags[i]
			for reservedFlag(f.FlagName) || used[f.FlagName] {
				f.FlagName = "query-" + f.FlagName
			}
			used[f.FlagName] = true
		}
		bodyFields := append([]bodyField(nil), o.BodyFields...)
		for _, f := range bodyFields {
			used[f.FlagName] = true
		}
		for i := range bodyFields {
			f := &bodyFields[i]
			collision := reservedFlag(f.FlagName)
			for _, p := range append(append([]param(nil), pathFlags...), queryFlags...) {
				collision = collision || p.FlagName == f.FlagName
			}
			if collision {
				f.FlagName = "body-" + f.FlagName
				for reservedFlag(f.FlagName) || used[f.FlagName] {
					f.FlagName = "body-" + f.FlagName
				}
				used[f.FlagName] = true
			}
		}
		responses := o.SuccessResponses
		if len(responses) == 0 {
			responses = []successResponse{{Code: o.Success2xx, HasJSON: o.HasJSONResp, HasSchema: o.HasSchemaResp}}
		}
		r := renderOp{
			SuccessResponses:  responses,
			ConfigQueryParams: configQueryParams,
			Constructor:       "new" + tagIdent + pascal(verb) + "Cmd",
			Verb:              verb,
			Short:             shortFor(o),
			Method:            strings.ToUpper(o.Method),
			Path:              o.Path,
			OperationID:       o.OperationID,
			GenMethod:         genMethod,
			PathParams:        pathFlags,
			QueryParams:       queryFlags,
			PathArgExprs:      pathArgExprs,
			UUIDPathParams:    uuidPathParams,
			HasUUIDParams:     hasUUIDParams,
			HasBody:           o.HasBody,
			HasRawBody:        o.HasRawBody,
			RawBodyType:       o.RawBodyType,
			BodyTypeName:      o.BodyTypeName,
			BodyFields:        bodyFields,
			NeedsOrgID:        needsOrgID,
			ParamsType:        paramsType,
			CanPaginate:       canPaginate,
			Success2xx:        o.Success2xx,
			HasJSONResp:       o.HasJSONResp,
			HasSchemaResp:     o.HasSchemaResp,
			HasRawJSONResp:    o.Action == "list" && o.Success2xx == "200" && !o.HasJSONResp && o.HasSchemaResp,
			IsWrite:           isWriteAction(o.Action) || (o.Action == "action" && o.Method != "get"),
			IsDestructive:     o.Action == "delete" || isDestructiveVerb(verb),
		}
		r.OutputContract = "structured"
		for _, response := range responses {
			if !response.HasJSON && (!response.HasSchema || response.MediaType != "") {
				r.OutputContract = "text"
				break
			}
		}
		if r.OutputContract == "text" {
			for _, response := range responses {
				if stringValue(response.Schema["format"]) == "binary" {
					r.OutputContract = "binary"
				}
			}
			for _, response := range responses {
				if response.HasJSON || response.HasSchema && response.MediaType == "" {
					r.OutputContract = "mixed"
					break
				}
			}
		}
		r.Metadata = operationMetadata(cmdName, o, r)
		r.Example = helpExample(cmdName, o, r)
		data.Ops = append(data.Ops, r)
	}

	metadata := map[string]interface{}{}
	for _, r := range data.Ops {
		if _, exists := metadata[r.OperationID]; exists {
			return nil, nil, fmt.Errorf("duplicate operationId %q", r.OperationID)
		}
		metadata[r.OperationID] = r.Metadata
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, nil, err
	}
	return buf.Bytes(), metadata, nil
}

func operationMetadata(cmdName string, o operation, r renderOp) map[string]interface{} {
	params := []interface{}{}
	add := func(p param) {
		source := p.Source
		if source == "" {
			source = "flag"
		}
		var enum interface{}
		if p.Schema != nil {
			enum = p.Schema["enum"]
		}
		metadata := map[string]interface{}{
			"flag": p.FlagName, "source": source, "in": p.In,
			"type": stringValue(p.Schema["type"]), "required": p.Required,
			"enum": enum, "schema": p.Schema, "description": p.Description,
			"name": p.Name,
		}
		for _, key := range []string{"example", "examples", "style", "explode", "allowReserved", "deprecated"} {
			if value, exists := p.Metadata[key]; exists {
				metadata[key] = value
			}
		}
		params = append(params, metadata)
	}
	for _, p := range o.PathParams {
		if p.FlagName == "org-id" {
			p.Source = "config"
			add(p)
		}
	}
	for _, p := range r.PathParams {
		add(p)
	}
	for _, p := range r.ConfigQueryParams {
		add(p)
	}
	for _, p := range r.QueryParams {
		add(p)
	}
	if r.CanPaginate {
		for _, p := range o.QueryParams {
			if p.Name == "skipToken" {
				p.Source = "pagination"
				add(p)
			}
		}
	}
	bodyFlags := []string{}
	bodyFields := []interface{}{}
	for _, f := range r.BodyFields {
		bodyFlags = append(bodyFlags, f.FlagName)
		if f.Schema != nil {
			bodyFields = append(bodyFields, map[string]interface{}{
				"flag": f.FlagName, "property": f.JSONKey, "required": f.Required, "schema": f.Schema,
			})
		}
	}
	envelope := "none"
	if r.CanPaginate {
		envelope = "values"
	}
	m := map[string]interface{}{
		"operationId": r.OperationID, "command": []string{cmdName, r.Verb},
		"method": r.Method, "path": r.Path, "tag": o.Tag, "resource": o.Resource, "action": o.Action,
		"summary": o.Summary, "description": o.Description, "destructive": r.IsDestructive,
		"paginated": r.CanPaginate, "responseEnvelope": envelope, "params": params, "bodyFlags": bodyFlags,
	}
	if len(bodyFields) > 0 {
		m["bodyFields"] = bodyFields
	}
	if o.RequestSchema != nil {
		m["requestSchema"] = o.RequestSchema
	}
	if o.ResponseSchema != nil {
		m["responseSchema"] = o.ResponseSchema
	}
	if o.RequestExample != nil {
		m["requestExample"] = o.RequestExample
		if o.RequestExampleSource != "" {
			m["requestExampleSource"] = o.RequestExampleSource
		}
	}
	return m
}

func containsString(value interface{}, want string) bool {
	values, _ := value.([]interface{})
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
}

func helpExample(cmdName string, o operation, r renderOp) string {
	parts := []string{"flexera-cli", cmdName, r.Verb}
	params, _ := r.Metadata["params"].([]interface{})
	for _, value := range params {
		p := value.(map[string]interface{})
		if p["required"] == true {
			flag := p["flag"].(string)
			parts = append(parts, "--"+flag, strings.ToUpper(strings.ReplaceAll(flag, "-", "_")))
		}
	}
	if r.HasBody || r.HasRawBody {
		parts = append(parts, "--body", "@request.json")
	}
	line := "  " + strings.Join(parts, " ")
	text := "Illustrative only: replace uppercase tokens; provide your own request.json for body input.\n" + line
	if r.IsWrite {
		text += "\n" + line + " --dry-run"
	}
	if r.HasBody {
		text += "\nValidated illustrative body, when available (review before use):\n  flexera-cli cli schema " + cmdName + " " + r.Verb + " --example > request.json"
	}
	return text
}

func shortFor(o operation) string {
	if s := strings.TrimSpace(o.Summary); s != "" {
		return strings.ReplaceAll(s, `"`, `'`)
	}
	return strings.ToUpper(o.Method) + " " + o.Path
}

// assignVerbs derives a friendly CLI verb for each operation. A single op for
// an action uses the action name (list/get/create/update/replace/delete).
// When an action repeats within a tag, the ops are disambiguated by the path
// suffix after the group's common prefix: a literal trailing segment becomes
// the verb (e.g. "accept", "grant", "revoke", "report", "activation"); a
// collection-vs-item collision becomes "<action>-all" (broader) and
// "<action>" (item). Any residual collision gets a numeric suffix.
func assignVerbs(ops []operation) []string {
	verbs := make([]string, len(ops))

	// Action-verb special case: x-flexera-action == "action" represents
	// RPC-style POST sub-actions (e.g. /policy/.../{id}/evaluate) where
	// the literal verb name "action" is not a useful CLI command. Use
	// the path's trailing literal segment (e.g. "evaluate"). Indices
	// pre-assigned here are excluded from the action-keyed grouping in
	// the initial pass; collision passes 1–3 still run over the full
	// verbs slice, so a pre-assigned "evaluate" colliding across two
	// resources gets disambiguated by scope (pass 1) or pluralized by
	// numeric dedupe (pass 3).
	for i, o := range ops {
		if o.Action != "action" {
			continue
		}
		verb := lastLiteralSeg(segsOf(o.Path))
		if verb == "" {
			verb = o.Action // unreachable in practice; pass 3 dedupes
		}
		verbs[i] = verb
	}

	// Initial pass: a unique distinguishing literal tail becomes the verb
	// (accept/decline/grant/revoke/report/activation); otherwise the action.
	// Skip indices already assigned by the action-verb special case above.
	byAction := map[string][]int{}
	for i, o := range ops {
		if verbs[i] != "" {
			continue
		}
		byAction[o.Action] = append(byAction[o.Action], i)
	}
	for action, idxs := range byAction {
		if len(idxs) == 1 {
			verbs[idxs[0]] = action
			continue
		}
		grp := make([][]string, len(idxs))
		for k, i := range idxs {
			grp[k] = segsOf(ops[i].Path)
		}
		cp := commonPrefixLen(grp)
		lits := make([]string, len(idxs))
		litCount := map[string]int{}
		for k := range idxs {
			lits[k] = lastLiteralSeg(grp[k][cp:])
			if lits[k] != "" {
				litCount[lits[k]]++
			}
		}
		for k, i := range idxs {
			if lits[k] != "" && litCount[lits[k]] == 1 {
				verbs[i] = lits[k]
			} else {
				verbs[i] = action
			}
		}
	}

	// Pass 1: resolve collisions by the first distinguishing literal segment
	// (e.g. scope orgs/projects, or service root iam/policy). orgs is treated
	// as the canonical/bare variant; projects becomes "-project"; any other
	// distinguishing root (iam/policy/risk/...) is appended verbatim.
	for v, idxs := range groupVerbs(verbs) {
		if len(idxs) < 2 {
			continue
		}
		segs := make([][]string, len(idxs))
		for k, i := range idxs {
			segs[k] = segsOf(ops[i].Path)
		}
		cp := commonPrefixLen(segs)
		disc := make([]string, len(idxs))
		ok := true
		for k := range idxs {
			disc[k] = firstLiteralAtOrAfter(segs[k], cp)
			if disc[k] == "" {
				ok = false
			}
		}
		if !ok || !allDistinct(disc) {
			continue
		}
		for k, i := range idxs {
			switch disc[k] {
			case "orgs", "org":
				verbs[i] = v // canonical
			case "projects":
				verbs[i] = v + "-project"
			default:
				verbs[i] = v + "-" + pathSuffixIdent("/"+disc[k])
			}
		}
	}

	// Pass 2: remaining collisions — the deepest (most path segments) op keeps
	// the bare verb; broader ones get "-all".
	for v, idxs := range groupVerbs(verbs) {
		if len(idxs) < 2 {
			continue
		}
		sort.SliceStable(idxs, func(a, b int) bool {
			return len(segsOf(ops[idxs[a]].Path)) < len(segsOf(ops[idxs[b]].Path))
		})
		for j := 0; j < len(idxs)-1; j++ {
			verbs[idxs[j]] = v + "-all"
		}
	}

	// Pass 3: global numeric dedupe for any residual collisions.
	final := map[string]bool{}
	for i := range verbs {
		base := verbs[i]
		v := base
		for n := 2; final[v]; n++ {
			v = fmt.Sprintf("%s-%d", base, n)
		}
		final[v] = true
		verbs[i] = v
	}
	return verbs
}

// groupVerbs maps each verb to the op indices currently using it (index order).
func groupVerbs(verbs []string) map[string][]int {
	g := map[string][]int{}
	for i, v := range verbs {
		g[v] = append(g[v], i)
	}
	return g
}

// firstLiteralAtOrAfter returns the first non-parameter segment at or after
// index start, or "" if none.
func firstLiteralAtOrAfter(segs []string, start int) string {
	for i := start; i < len(segs); i++ {
		if !strings.HasPrefix(segs[i], "{") {
			return segs[i]
		}
	}
	return ""
}

// allDistinct reports whether every string in xs is unique.
func allDistinct(xs []string) bool {
	seen := make(map[string]bool, len(xs))
	for _, x := range xs {
		if seen[x] {
			return false
		}
		seen[x] = true
	}
	return true
}

// segsOf splits a path into non-empty segments.
func segsOf(p string) []string {
	var out []string
	for _, s := range strings.Split(strings.Trim(p, "/"), "/") {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

// commonPrefixLen returns the number of leading path segments shared by every
// group member.
func commonPrefixLen(groups [][]string) int {
	if len(groups) == 0 {
		return 0
	}
	minLen := len(groups[0])
	for _, g := range groups {
		if len(g) < minLen {
			minLen = len(g)
		}
	}
	n := 0
	for ; n < minLen; n++ {
		seg := groups[0][n]
		for _, g := range groups {
			if g[n] != seg {
				return n
			}
		}
	}
	return n
}

// lastLiteralSeg returns the last non-parameter segment of a suffix as a
// kebab-case verb, or "" when the suffix is empty or all parameters.
func lastLiteralSeg(suffix []string) string {
	for i := len(suffix) - 1; i >= 0; i-- {
		if strings.HasPrefix(suffix[i], "{") {
			continue
		}
		return pathSuffixIdent("/" + suffix[i])
	}
	return ""
}

func pathSuffixIdent(p string) string {
	segs := strings.Split(strings.Trim(p, "/"), "/")
	for i := len(segs) - 1; i >= 0; i-- {
		s := segs[i]
		if strings.HasPrefix(s, "{") {
			continue
		}
		s = strings.Map(func(r rune) rune {
			switch {
			case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z', r >= '0' && r <= '9':
				return r
			case r == '-' || r == '_' || r == '.':
				return '-'
			}
			return -1
		}, s)
		return kebab(s)
	}
	return "x"
}

func pascal(s string) string {
	parts := strings.FieldsFunc(s, func(r rune) bool { return r == '-' || r == '_' })
	for i, p := range parts {
		if p == "" {
			continue
		}
		parts[i] = strings.ToUpper(p[:1]) + p[1:]
	}
	return strings.Join(parts, "")
}

var tmpl = template.Must(template.New("file").Funcs(template.FuncMap{
	"pascal": pascal,
	"quote":  strconv.Quote,
	"join":   func(s []string) string { return strings.Join(s, ", ") },
}).Parse(`// Code generated by cmd/gencli. DO NOT EDIT.
// Source: unified OpenAPI spec, tag {{.Tag}}.

package {{.Pkg}}

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
{{- if .HasTime }}
	"time"
{{- end }}
{{- if .HasDate }}
	openapi_types "github.com/oapi-codegen/runtime/types"
{{- end }}
{{- range .Ops }}{{- if .HasRawBody }}
	"bytes"
{{- break}}{{- end}}{{- end}}
{{- range .Ops }}{{- if .HasUUIDParams }}
	"github.com/google/uuid"
{{- break}}{{- end}}{{- end}}

	"github.com/spf13/cobra"

	clipkg "github.com/flexera-public/flexera-cli/internal/cli"
	flexera "github.com/flexera-public/unified-go-client"
)

var (
	_ = context.Background
	_ = json.Unmarshal
	_ = fmt.Sprint
	_ = strings.TrimSpace
)

// NewCmd builds the "{{.CmdName}}" command tree.
func NewCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "{{.CmdName}}",
		Short: "{{.Short}}",
	}
	c.AddCommand(
{{- range .Ops }}
		{{.Constructor}}(),
{{- end }}
	)
	return c
}

{{ range .Ops }}
{{- $op := . }}
// {{.Constructor}} — {{.Method}} {{.Path}} (operationId: {{.OperationID}})
func {{.Constructor}}() *cobra.Command {
{{- if or .PathParams .QueryParams .CanPaginate .HasBody .HasRawBody .IsWrite }}
	var (
{{- range .PathParams }}
		{{.GoName}} {{.GoType}}
{{- end }}
{{- range .QueryParams }}
		{{.GoName}} {{if .TimeLayout}}string{{else}}{{.GoType}}{{end}}
{{- end }}
{{- if .CanPaginate }}
		noPaginate bool
		skipToken  string
{{- end }}
{{- if or .HasBody .HasRawBody }}
		bodyRaw string
{{- end }}
{{- if .HasBody }}
{{- range .BodyFields }}
		{{.GoVar}} {{.GoType}}
{{- end }}
{{- end }}
{{- if .IsWrite }}
		dryRun bool
		yes    bool
{{- if .HasBody }}
		interactive bool
{{- end }}
{{- end }}
	)
{{- end }}
	c := &cobra.Command{
		Use:   "{{.Verb}}",
		Short: {{quote .Short}},
		Example: {{quote .Example}},
		Annotations: map[string]string{"flexera.operationId": {{quote .OperationID}}, "flexera.output": {{quote .OutputContract}}, "flexera.validation": "{{if .HasBody}}body{{else}}params{{end}}"},
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
{{- if and .IsWrite .HasBody }}
			if interactive {
				if err := clipkg.GuardInteractive(cmd, bodyRaw); err != nil { return err }
				if err := clipkg.GatherInteractiveParams(cmd, "{{.OperationID}}"); err != nil { return err }
			}
{{- end }}
			// Parse formatted query flags before client creation/authentication.
{{- range .QueryParams }}{{- if .TimeLayout }}
			var {{.GoName}}Value {{.GoType}}
{{- if .Required }}
			if !cmd.Flags().Changed("{{.FlagName}}") {
				return clipkg.Exit(2, fmt.Errorf("--{{.FlagName}} is required (expected {{if eq .GoType "time.Time"}}RFC3339 date-time with timezone{{else}}YYYY-MM-DD date{{end}})"))
			}
{{- end }}
			if cmd.Flags().Changed("{{.FlagName}}") {
				v, err := time.Parse({{quote .TimeLayout}}, {{.GoName}})
				if err != nil {
					return clipkg.Exit(2, fmt.Errorf("--{{.FlagName}}: expected {{if eq .GoType "time.Time"}}RFC3339 date-time with timezone (e.g. 2021-06-28T00:00:00Z){{else}}YYYY-MM-DD date (e.g. 2021-06-28){{end}}: %w", err))
				}
				{{.GoName}}Value = {{if eq .GoType "openapi_types.Date"}}openapi_types.Date{Time: v}{{else}}v{{end}}
			}
{{- end }}{{- end }}
			deps := clipkg.DepsFrom(cmd.Context())
			effectiveParams, err := clipkg.ValidateCommandParams(cmd, "{{.OperationID}}")
			if err != nil { return err }
{{- if not .IsWrite }}
			_ = effectiveParams
{{- end }}
{{- range .PathParams }}{{ if eq .GoType "string" }}
			if strings.TrimSpace({{.GoName}}) == "" {
				return fmt.Errorf("--{{.FlagName}} is required")
			}
{{- end }}{{ end }}
{{- range .UUIDPathParams }}
			{{.GoName}}UUID, err := uuid.Parse({{.GoName}})
			if err != nil {
				return fmt.Errorf("--{{.FlagName}}: invalid UUID: %w", err)
			}
{{- end }}
{{- if .ParamsType }}
			params := {{.ParamsType}}{}
{{- range .ConfigQueryParams }}
{{- if not .Required }}
			if _, supplied := effectiveParams["{{.FlagName}}"]; supplied {
				v := {{.ValueExpr}}
				params.{{pascal .Name}} = &v
			}
{{- else }}
			params.{{pascal .Name}} = {{.ValueExpr}}
{{- end }}
{{- end }}
{{- range .QueryParams }}
{{- if .TimeLayout }}
			if cmd.Flags().Changed("{{.FlagName}}") {
				params.{{pascal .Name}} = {{if not .Required}}&{{end}}{{.GoName}}Value
			}
{{- else if .IsArray }}
			if cmd.Flags().Changed("{{.FlagName}}") {
{{- if .Required }}
{{- if .IsEnum }}
				var ev{{pascal .Name}} []{{.EnumType}}
				for _, s := range {{.GoName}} { ev{{pascal .Name}} = append(ev{{pascal .Name}}, {{.EnumType}}(s)) }
				params.{{ pascal .Name }} = ev{{pascal .Name}}
{{- else }}
				params.{{ pascal .Name }} = {{.GoName}}
{{- end }}
{{- else }}
{{- if .IsEnum }}
				var ev{{pascal .Name}} []{{.EnumType}}
				for _, s := range {{.GoName}} { ev{{pascal .Name}} = append(ev{{pascal .Name}}, {{.EnumType}}(s)) }
				params.{{ pascal .Name }} = &ev{{pascal .Name}}
{{- else }}
				v := {{.GoName}}
				params.{{ pascal .Name }} = &v
{{- end }}
{{- end }}
			}
{{- else }}
			if cmd.Flags().Changed("{{.FlagName}}") {
{{- if .Required }}
{{- if .IsEnum }}
				params.{{ pascal .Name }} = {{.EnumType}}({{.GoName}})
{{- else }}
				params.{{ pascal .Name }} = {{.GoName}}
{{- end }}
{{- else }}
{{- if .IsEnum }}
				ev := {{.EnumType}}({{.GoName}})
				params.{{ pascal .Name }} = &ev
{{- else }}
				v := {{.GoName}}
				params.{{ pascal .Name }} = &v
{{- end }}
{{- end }}
			}
{{- end }}
{{- end }}
{{- end }}
{{- if .HasBody }}
			fields := map[string]any{}
{{- range .BodyFields }}
			if cmd.Flags().Changed("{{.FlagName}}") {
				fields["{{.JSONKey}}"] = {{.GoVar}}
			}
{{- end }}
			var typed any
			if len(fields) > 0 {
				typed = fields
			}
			raw, err := clipkg.ResolveBody(bodyRaw, typed, cmd.InOrStdin())
			if err != nil {
				return err
			}
			if len(raw) == 0 {
{{- if .IsWrite }}
				if !interactive { return clipkg.Exit(2, fmt.Errorf("a request body is required: pass --body (inline JSON, @file, or @-) or the body field flags")) }
{{- else }}
				return clipkg.Exit(2, fmt.Errorf("a request body is required: pass --body (inline JSON, @file, or @-) or the body field flags"))
{{- end }}
			}
{{- if .IsWrite }}
			if interactive { raw, err = clipkg.GatherInteractiveBody(cmd, "{{.OperationID}}", raw); if err != nil { return err } }
{{- end }}
			var body flexera.{{.BodyTypeName}}
			noValidate, err := cmd.Flags().GetBool(clipkg.FlagNoValidate)
			if err != nil { return clipkg.Exit(2, err) }
			effectiveBody, validation, requestSchema, err := clipkg.PrepareRequestBody("{{.OperationID}}", raw, &body, noValidate)
			if err != nil { return err }
{{- if not .IsWrite }}
			_, _, _ = effectiveBody, validation, requestSchema
{{- end }}
{{- end }}
{{- if .HasRawBody }}
			raw, err := clipkg.ResolveRawBody(bodyRaw, cmd.InOrStdin())
			if err != nil {
				return err
			}
			if len(raw) == 0 {
				return fmt.Errorf("a request body is required: pass --body @file or --body @-")
			}
{{- end }}
{{- if .IsWrite }}
			planParams := effectiveParams
			writePlan := clipkg.Plan{Command: cmd.CommandPath(), Method: "{{.Method}}", Path: "{{.Path}}", Params: planParams, Destructive: {{.IsDestructive}}}
{{- if .NeedsOrgID }}
			writePlan.OrgID = deps.Config.OrgID
			planParams["org-id"] = deps.Config.OrgID
{{- end }}
{{- if .HasBody }}
			writePlan.Body, writePlan.Validation, writePlan.RequestSchema = effectiveBody, validation, requestSchema
{{- else if .HasRawBody }}
			writePlan.Body = raw
			writePlan.RawUpload = true
			writePlan.Validation = &clipkg.ValidationResult{Status: "unsupported"}
{{- end }}
			var writeDone bool
			var werr error
{{- if .HasBody }}
			if interactive { writeDone, werr = clipkg.ConfirmInteractive(cmd, dryRun, yes, writePlan, deps.Printer) } else { writeDone, werr = clipkg.ConfirmPlan(dryRun, yes, deps.Stdout, writePlan, deps.Printer) }
{{- else }}
			writeDone, werr = clipkg.ConfirmPlan(dryRun, yes, deps.Stdout, writePlan, deps.Printer)
{{- end }}
			if werr != nil {
				return werr
			} else if writeDone {
				return nil
			}
{{- end }}
			client, err := deps.APIClient()
			if err != nil { return err }
{{- if .CanPaginate }}
			var initialSkipToken *string
			if t := strings.TrimSpace(skipToken); t != "" {
				initialSkipToken = &t
			}
			result, err := flexera.CollectPages(cmd.Context(), noPaginate, initialSkipToken,
				func(ctx context.Context, st *string) (any, error) {
					p := params
					p.SkipToken = st
					resp, callErr := client.{{.GenMethod}}WithResponse(ctx{{ range .PathArgExprs }}, {{.}}{{ end }}, &p)
					if callErr != nil {
						return nil, callErr
					}
					switch resp.StatusCode() {
{{- range .SuccessResponses }}
{{- if and .HasJSON .HasSchema }}
					case {{.Code}}:
						if resp.JSON{{.Code}} != nil {
							if deps.Config.Output == "table" && deps.Printer.JQ == nil && len(deps.Printer.Fields) == 0 { return resp.JSON{{.Code}}, nil }
							return clipkg.DecodeResponseJSON(resp.Body)
						}
{{- else if or .HasJSON .HasSchema }}
					case {{.Code}}:
						return clipkg.DecodeResponseJSON(resp.Body)
{{- end }}
{{- end }}
					}
					return nil, flexera.ResponseError(resp.StatusCode(), resp.Body)
				})
			if err != nil {
				return err
			}
			return deps.Printer.Render(deps.Stdout, deps.Config.Output, result)
{{- else }}
			resp, err := client.{{.GenMethod}}{{if .HasRawBody}}WithBody{{end}}WithResponse(cmd.Context(){{ range .PathArgExprs }}, {{.}}{{ end }}{{ if .ParamsType }}, &params{{ end }}{{ if .HasBody }}, body{{ end }}{{ if .HasRawBody }}, "{{.RawBodyType}}", bytes.NewReader(raw){{ end }})
			if err != nil {
				return err
			}
			switch resp.StatusCode() {
{{- range .SuccessResponses }}
			case {{.Code}}:
{{- if and .HasJSON .HasSchema }}
				if resp.JSON{{.Code}} == nil { return flexera.ResponseError(resp.StatusCode(), resp.Body) }
				if deps.Config.Output == "table" && deps.Printer.JQ == nil && len(deps.Printer.Fields) == 0 { return deps.Printer.Render(deps.Stdout, deps.Config.Output, resp.JSON{{.Code}}) }
				result, err := clipkg.DecodeResponseJSON(resp.Body)
				if err != nil { return err }
				return deps.Printer.Render(deps.Stdout, deps.Config.Output, result)
{{- else if .HasJSON }}
				result, err := clipkg.DecodeResponseJSON(resp.Body)
				if err != nil { return err }
				return deps.Printer.Render(deps.Stdout, deps.Config.Output, result)
{{- else if .HasSchema }}
				{{- if .MediaType }}
				_, err := deps.Stdout.Write(resp.Body)
				return err
				{{- else }}
			result, err := clipkg.DecodeResponseJSON(resp.Body)
			if err != nil { return err }
			return deps.Printer.Render(deps.Stdout, deps.Config.Output, result)
				{{- end }}
{{- else }}
			fmt.Fprintln(deps.Stdout, "OK")
			return nil
{{- end }}
{{- end }}
			default:
				return flexera.ResponseError(resp.StatusCode(), resp.Body)
			}
{{- end }}
		},
	}
{{- range .PathParams }}
{{- if eq .GoType "string" }}
	c.Flags().StringVar(&{{.GoName}}, "{{.FlagName}}", "", "{{.Name}} (path, required)")
{{- else if eq .GoType "int64" }}
	c.Flags().Int64Var(&{{.GoName}}, "{{.FlagName}}", 0, "{{.Name}} (path, required)")
{{- else if eq .GoType "int" }}
	c.Flags().IntVar(&{{.GoName}}, "{{.FlagName}}", 0, "{{.Name}} (path, required)")
{{- else if eq .GoType "bool" }}
	c.Flags().BoolVar(&{{.GoName}}, "{{.FlagName}}", false, "{{.Name}} (path, required)")
{{- end }}
{{- end }}
{{- range .QueryParams }}
{{- if .TimeLayout }}
	c.Flags().StringVar(&{{.GoName}}, "{{.FlagName}}", "", "{{.Name}} (query, {{if eq .GoType "time.Time"}}RFC3339 date-time with timezone{{else}}YYYY-MM-DD date{{end}})")
{{- else if .IsArray }}
	c.Flags().StringSliceVar(&{{.GoName}}, "{{.FlagName}}", nil, "{{.Name}} (query)")
{{- else if eq .GoType "string" }}
	c.Flags().StringVar(&{{.GoName}}, "{{.FlagName}}", "", "{{.Name}} (query)")
{{- else if eq .GoType "int64" }}
	c.Flags().Int64Var(&{{.GoName}}, "{{.FlagName}}", 0, "{{.Name}} (query)")
{{- else if eq .GoType "int" }}
	c.Flags().IntVar(&{{.GoName}}, "{{.FlagName}}", 0, "{{.Name}} (query)")
{{- else if eq .GoType "bool" }}
	c.Flags().BoolVar(&{{.GoName}}, "{{.FlagName}}", false, "{{.Name}} (query)")
{{- end }}
{{- end }}
{{- if .CanPaginate }}
	c.Flags().BoolVar(&noPaginate, "no-paginate", false, "return only the first page (do not follow nextPage)")
	c.Flags().StringVar(&skipToken, "skip-token", "", "resume pagination from this token")
{{- end }}
{{- if or .HasBody .HasRawBody }}
{{- range .BodyFields }}
{{- if eq .GoType "string" }}
	c.Flags().StringVar(&{{.GoVar}}, "{{.FlagName}}", "", "{{.JSONKey}} (body)")
{{- else if eq .GoType "int64" }}
	c.Flags().Int64Var(&{{.GoVar}}, "{{.FlagName}}", 0, "{{.JSONKey}} (body)")
{{- else if eq .GoType "int" }}
	c.Flags().IntVar(&{{.GoVar}}, "{{.FlagName}}", 0, "{{.JSONKey}} (body)")
{{- else if eq .GoType "float64" }}
	c.Flags().Float64Var(&{{.GoVar}}, "{{.FlagName}}", 0, "{{.JSONKey}} (body)")
{{- else if eq .GoType "bool" }}
	c.Flags().BoolVar(&{{.GoVar}}, "{{.FlagName}}", false, "{{.JSONKey}} (body)")
{{- else if eq .GoType "[]string" }}
	c.Flags().StringSliceVar(&{{.GoVar}}, "{{.FlagName}}", nil, "{{.JSONKey}} (body)")
{{- end }}
{{- end }}
	c.Flags().StringVar(&bodyRaw, "body", "", "{{if .HasRawBody}}raw request body (@file | @-){{else}}raw JSON body (inline | @file | @-); overrides body field flags{{end}}")
{{- end }}
{{- if .IsWrite }}
	c.Flags().BoolVar(&dryRun, "dry-run", false, "print the planned operation as JSON and exit without calling the API")
	c.Flags().BoolVar(&yes, "yes", false, "confirm the operation (required for destructive ops)")
{{- if .HasBody }}
	c.Flags().BoolVarP(&interactive, "interactive", "i", false, "edit inputs in a terminal form, review a plan and approve with typed yes")
{{- end }}
{{- end }}
	return c
}
{{ end }}
`))
