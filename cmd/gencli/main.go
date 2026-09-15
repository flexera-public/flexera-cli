// Command gencli generates one cobra command package per OpenAPI tag from
// the annotated unified OpenAPI spec.
//
// Usage:
//
//	go run ./cmd/gencli -spec ../unified-openapi/openapi3.json \
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
// bodies, array params of non-string type, or date/date-time params are
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
	"strings"
	"text/template"
)

type operation struct {
	Method       string
	Path         string
	OperationID  string
	Summary      string
	Tag          string
	Action       string
	Resource     string
	PathParams   []param
	QueryParams  []param
	HasHeaders   bool
	HasBody      bool
	BodyTypeName string // generated client's "<Method>JSONRequestBody"
	BodyFields   []bodyField
	Paginated    bool
	Success2xx   string // "200", "201", or "204"
	HasJSONResp  bool
}

type param struct {
	Name     string // raw spec name (orgId, id, project_id)
	GoName   string // Go-safe lowerCamelCase (orgID, id, projectID)
	FlagName string // kebab-case CLI flag (org-id, id, project-id)
	GoType   string // string|int64|int|bool|[]string
	IsArray  bool
	IsInt64  bool
	IsUUID   bool
	IsEnum   bool
	EnumType string
	In       string // path|query|header
	Required bool
}

// bodyField is a scalar (or scalar-array) top-level property of a JSON
// request body, surfaced as a typed cobra flag.
type bodyField struct {
	JSONKey  string // spec property name, used as the JSON map key
	GoVar    string // local Go variable name
	FlagName string // kebab-case CLI flag
	GoType   string // string|int|int64|float64|bool|[]string
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
		specPath = flag.String("spec", "", "path to annotated unified openapi.json")
		tag      = flag.String("tag", "", "OpenAPI tag to generate (e.g. Budget)")
		pkg      = flag.String("pkg", "", "Go package name for the generated file")
		cmdName  = flag.String("cmd", "", "root CLI command name (e.g. budget-gen)")
		outPath  = flag.String("out", "", "output Go file path")
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

	ops, drops := collectOps(&s, *tag)
	for _, d := range drops {
		fmt.Fprintf(os.Stderr, "gencli: skip %s %s — %s\n", strings.ToUpper(d.Method), d.Path, d.Reason)
	}
	if len(ops) == 0 {
		fmt.Fprintf(os.Stderr, "no operations found for tag %q\n", *tag)
		os.Exit(1)
	}

	src, err := render(*tag, *pkg, *cmdName, ops)
	check(err)

	formatted, err := format.Source(src)
	if err != nil {
		_ = os.WriteFile(*outPath+".raw", src, 0o644)
		check(fmt.Errorf("gofmt failed (raw source written to %s.raw): %w", *outPath, err))
	}

	check(os.MkdirAll(dirOf(*outPath), 0o755))
	check(os.WriteFile(*outPath, formatted, 0o644))
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

func collectOps(s *spec, tag string) ([]operation, []drop) {
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
			_ = json.Unmarshal(rawParams, &arr)
			pathLevelParams = arr
		}
		for _, m := range methods {
			raw, ok := pi[m]
			if !ok {
				continue
			}
			var op map[string]interface{}
			if err := json.Unmarshal(raw, &op); err != nil {
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
			pag, _ := op["x-flexera-paginated"].(bool)

			pps, qps, hasHeaders, ok := extractParams(op)
			if !ok {
				drops = append(drops, drop{Method: m, Path: p, Action: action, Reason: "param extraction failed (unsupported param shape)"})
				continue
			}

			hasBody := false
			hasNonJSONBody := false
			var bodyFields []bodyField
			if rb, ok := op["requestBody"].(map[string]interface{}); ok {
				if content, ok := rb["content"].(map[string]interface{}); ok {
					if jc, ok := content["application/json"].(map[string]interface{}); ok {
						hasBody = true
						bodyFields = extractBodyFields(s, jc)
					} else if len(content) > 0 {
						hasNonJSONBody = true
					}
				}
			}
			if hasNonJSONBody {
				drops = append(drops, drop{Method: m, Path: p, Action: action, Reason: "non-JSON request body"})
				continue
			}

			success, hasJSON := successCode(op)
			if success == "" {
				drops = append(drops, drop{Method: m, Path: p, Action: action, Reason: "no 2xx success response"})
				continue
			}

			ops = append(ops, operation{
				Method:       m,
				Path:         p,
				OperationID:  opID,
				Summary:      summary,
				Tag:          tag,
				Action:       action,
				Resource:     resource,
				PathParams:   pps,
				QueryParams:  qps,
				HasHeaders:   hasHeaders,
				HasBody:      hasBody,
				BodyTypeName: methodName(opID) + "JSONRequestBody",
				BodyFields:   bodyFields,
				Paginated:    pag,
				Success2xx:   success,
				HasJSONResp:  hasJSON,
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
	case "body", "dry-run", "yes", "no-paginate", "skip-token":
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
	for _, name := range keys {
		ps, _ := props[name].(map[string]interface{})
		if ps == nil {
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
			// Collides with a CLI control flag (--body/--dry-run/--yes/etc.);
			// leave it to the --body escape hatch rather than shadowing.
			continue
		}
		fields = append(fields, bodyField{
			JSONKey:  name,
			GoVar:    "f" + pascal(goIdent(name)),
			FlagName: flag,
			GoType:   goType,
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
	if err := json.Unmarshal(raw, &next); err != nil {
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
			if format == "date-time" || format == "date" {
				return nil, nil, false, false
			}
			goType = "string"
			if hasEnum {
				isEnum = true
			}
		case "array":
			items, _ := schema["items"].(map[string]interface{})
			it, _ := items["type"].(string)
			if it != "string" && it != "" {
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
			Name:     name,
			GoName:   goIdent(name),
			FlagName: kebab(name),
			GoType:   goType,
			IsArray:  isArray,
			IsInt64:  isInt64,
			IsUUID:   format == "uuid",
			IsEnum:   isEnum,
			In:       in,
			Required: required,
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

func successCode(op map[string]interface{}) (code string, hasJSON bool) {
	responses, _ := op["responses"].(map[string]interface{})
	for _, c := range []string{"200", "201", "204"} {
		resp, ok := responses[c].(map[string]interface{})
		if !ok {
			continue
		}
		content, _ := resp["content"].(map[string]interface{})
		for ct := range content {
			if strings.Contains(ct, "json") {
				return c, true
			}
		}
		return c, false
	}
	return "", false
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
}

type renderOp struct {
	Constructor   string
	Verb          string
	Short         string
	Method        string
	Path          string
	OperationID   string
	GenMethod     string
	PathParams    []param // non-orgId path params (flags)
	QueryParams   []param // excludes skipToken when paginated
	PathArgExprs  []string
	UUIDPathParams []param // subset of PathParams with IsUUID, for pre-call parsing
	HasUUIDParams bool
	HasBody       bool
	BodyTypeName  string
	BodyFields    []bodyField
	NeedsOrgID    bool
	ParamsType    string
	CanPaginate   bool
	Success2xx    string
	HasJSONResp   bool
	IsWrite       bool
	IsDestructive bool
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
		for _, qp := range o.QueryParams {
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

		data.Ops = append(data.Ops, renderOp{
			Constructor:    "new" + tagIdent + pascal(verb) + "Cmd",
			Verb:           verb,
			Short:          shortFor(o),
			Method:         strings.ToUpper(o.Method),
			Path:           o.Path,
			OperationID:    o.OperationID,
			GenMethod:      genMethod,
			PathParams:     pathFlags,
			QueryParams:    queryFlags,
			PathArgExprs:   pathArgExprs,
			UUIDPathParams: uuidPathParams,
			HasUUIDParams:  hasUUIDParams,
			HasBody:        o.HasBody,
			BodyTypeName:   o.BodyTypeName,
			BodyFields:     o.BodyFields,
			NeedsOrgID:     needsOrgID,
			ParamsType:     paramsType,
			CanPaginate:    canPaginate,
			Success2xx:     o.Success2xx,
			HasJSONResp:    o.HasJSONResp,
			IsWrite:        isWriteAction(o.Action),
			IsDestructive:  o.Action == "delete" || isDestructiveVerb(verb),
		})
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
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
	"join":   func(s []string) string { return strings.Join(s, ", ") },
}).Parse(`// Code generated by cmd/gencli. DO NOT EDIT.
// Source: unified OpenAPI spec, tag {{.Tag}}.

package {{.Pkg}}

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
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
{{- if or .PathParams .QueryParams .HasBody .IsWrite }}
	var (
{{- range .PathParams }}
		{{.GoName}} {{.GoType}}
{{- end }}
{{- range .QueryParams }}
		{{.GoName}} {{.GoType}}
{{- end }}
{{- if .CanPaginate }}
		noPaginate bool
		skipToken  string
{{- end }}
{{- if .HasBody }}
		bodyRaw string
{{- range .BodyFields }}
		{{.GoVar}} {{.GoType}}
{{- end }}
{{- end }}
{{- if .IsWrite }}
		dryRun bool
		yes    bool
{{- end }}
	)
{{- end }}
	c := &cobra.Command{
		Use:   "{{.Verb}}",
		Short: "{{.Short}}",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			deps := clipkg.DepsFrom(cmd.Context())
{{- if .NeedsOrgID }}
			if err := deps.Config.RequireOrgID(); err != nil {
				return err
			}
{{- end }}
			client, err := deps.APIClient()
			if err != nil {
				return err
			}
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
{{- range .QueryParams }}
{{- if .IsArray }}
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
				return fmt.Errorf("a request body is required: pass --body (inline JSON, @file, or @-) or the body field flags")
			}
			var body flexera.{{.BodyTypeName}}
			if err := json.Unmarshal(raw, &body); err != nil {
				return fmt.Errorf("decoding request body: %w", err)
			}
{{- end }}
{{- if .IsWrite }}
			writePlan := map[string]any{"method": "{{.Method}} {{.Path}}"}
{{- if .NeedsOrgID }}
			writePlan["orgId"] = deps.Config.OrgID
{{- end }}
{{- if .HasBody }}
			writePlan["body"] = json.RawMessage(raw)
{{- end }}
			if writeDone, werr := clipkg.ConfirmWrite(dryRun, yes, {{.IsDestructive}}, deps.Stdout, writePlan); werr != nil {
				return werr
			} else if writeDone {
				return nil
			}
{{- end }}
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
					if resp.JSON200 == nil {
						return nil, flexera.ResponseError(resp.StatusCode(), resp.Body)
					}
					return resp.JSON200, nil
				})
			if err != nil {
				return err
			}
			return deps.Printer.Render(deps.Stdout, deps.Config.Output, result)
{{- else }}
			resp, err := client.{{.GenMethod}}WithResponse(cmd.Context(){{ range .PathArgExprs }}, {{.}}{{ end }}{{ if .ParamsType }}, &params{{ end }}{{ if .HasBody }}, body{{ end }})
			if err != nil {
				return err
			}
{{- if eq .Success2xx "204" }}
			if resp.StatusCode() != 204 {
				return flexera.ResponseError(resp.StatusCode(), resp.Body)
			}
			fmt.Fprintln(deps.Stdout, "OK")
			return nil
{{- else if and (eq .Success2xx "200") .HasJSONResp }}
			if resp.JSON200 == nil {
				return flexera.ResponseError(resp.StatusCode(), resp.Body)
			}
			return deps.Printer.Render(deps.Stdout, deps.Config.Output, resp.JSON200)
{{- else if and (eq .Success2xx "201") .HasJSONResp }}
			if resp.JSON201 == nil {
				return flexera.ResponseError(resp.StatusCode(), resp.Body)
			}
			return deps.Printer.Render(deps.Stdout, deps.Config.Output, resp.JSON201)
{{- else }}
			if resp.StatusCode() >= 300 {
				return flexera.ResponseError(resp.StatusCode(), resp.Body)
			}
			fmt.Fprintln(deps.Stdout, "OK")
			return nil
{{- end }}
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
{{- if .IsArray }}
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
{{- if .HasBody }}
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
	c.Flags().StringVar(&bodyRaw, "body", "", "raw JSON body (inline | @file | @-); overrides body field flags")
{{- end }}
{{- if .IsWrite }}
	c.Flags().BoolVar(&dryRun, "dry-run", false, "print the planned operation as JSON and exit without calling the API")
	c.Flags().BoolVar(&yes, "yes", false, "confirm the operation (required for destructive ops)")
{{- end }}
	return c
}
{{ end }}
`))
