package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// Command naming
//
// Each operation in a tag gets a command path relative to the tag command:
//
//	[<sub-resource>...] <action>
//
// derived only from its own HTTP path and x-flexera-action, after the parts
// every operation in the tag shares are removed:
//
//  1. Path noise is ignored: API versions (v1, v2), the config-supplied org
//     scope (orgs/{orgId}, a bare {orgId}) and a projects/{projectId} scope.
//  2. Each run of literal segments becomes one kebab-case name
//     (export/select -> export-select).
//  3. A run followed by an {id}, or a trailing run that is a resource (it has
//     an {id} form or several methods), is a sub-resource group; the leaf is
//     the CRUD action (list/get/create/update/replace/delete).
//  4. Any other trailing run names the operation itself (evaluate, report).
//  5. x-flexera-action "action" on a resource is "create"; otherwise it uses
//     its trailing literal (activate, verify).
//  6. The parent command name is not repeated in any name (risk compliance
//     compliance-cis -> risk compliance cis).
//
// Only two semantic tie-breaks exist: a collection-level operation beside its
// item-level twin becomes "<action>-all", and a project-scoped twin
// of an org-scoped operation moves under a "project" group. Anything still
// ambiguous fails generation; resolve it with commandOverrides, never with
// generated suffixes.

// commandOverrides pins command paths (relative to the tag command) for
// operations whose paths cannot be named unambiguously by the rules above.
var commandOverrides = map[string][]string{
	// POST .../client-secret rotates the secret; DELETE removes the old one.
	"Iam_Service_Account_Client_Rotate": {"client-secret", "rotate"},
	// GET /customers/tags spans all customers; .../customers/{id}/tags is one.
	"Iam_MSP_Customer_Tag_Index_Customers": {"tags", "list-all"},
	// Every device op shares devices/delete/, which would otherwise be dropped.
	"Vis_Device_estimateDelete": {"delete", "estimate"},
	"Vis_Device_delete":         {"delete", "submit"},

	// Semantic verbs for operations whose derived CRUD name misdescribes them.
	// PATCH .../bill-months reprocesses bills; PATCH .../invitations/{id}
	// resends; PATCH .../saas-data/{id} uploads (once).
	"BillAnalysis_bill_months_reprocess":                 {"reprocess"},
	"FinopsOnboarding_Processing_History_reprocess":      {"reprocess"},
	"Iam_Organization_Invitation_resend":                 {"resend"},
	"Saas_SaaS_Data_upload":                              {"upload"},
	"OptimaRecommendations_Recommendations_updateStatus": {"update-status"},
	"FinopsBilling_Shared_Cost_Rules_reorder":            {"reorder"},
	// POST .../queries runs a query; it does not list queries.
	"Vis_Query_create":         {"execute"},
	"Iam_Contracts_update":     {"update-last-seen"},
	"Iam_Contracts_activation": {"set-activation"},
	// GET /users/me is a singleton read.
	"Iam_User_Profile_show": {"get"},
	// Both return a signed URL rather than the settings object itself.
	"Iam_User_Setting_Blob_retrive_get_url": {"download-url"},
	"Iam_User_Setting_Blob_retrive_put_url": {"upload-url"},
	// GET .../{id}/validation validates one bill connect; .../validations all.
	"FinopsOnboarding_Bill_Connect_validate":                       {"validate-all"},
	"FinopsOnboarding_Bill_Connect_AWS_validate":                   {"validate"},
	"FinopsOnboarding_Bill_Connect_Azure_CSP_validate":             {"validate"},
	"FinopsOnboarding_Bill_Connect_Azure_EA_Management_validate":   {"validate"},
	"FinopsOnboarding_Bill_Connect_Azure_MCA_validate":             {"validate"},
	"FinopsOnboarding_Bill_Connect_Common_Bill_Ingestion_validate": {"validate"},
	"FinopsOnboarding_Bill_Connect_Databricks_validate":            {"validate"},
	"FinopsOnboarding_Bill_Connect_GCP_validate":                   {"validate"},
	"FinopsOnboarding_Bill_Connect_Snowflake_validate":             {"validate"},
}

// commandAliases keeps earlier leaf spellings working after a rename. An
// alias replaces only the leaf name, so it stays beside the canonical command.
var commandAliases = map[string][]string{
	"BillAnalysis_bill_months_reprocess":                           {"update"},
	"FinopsOnboarding_Processing_History_reprocess":                {"update"},
	"Iam_Organization_Invitation_resend":                           {"update"},
	"Saas_SaaS_Data_upload":                                        {"update"},
	"OptimaRecommendations_Recommendations_updateStatus":           {"updatestatus"},
	"FinopsBilling_Shared_Cost_Rules_reorder":                      {"operations-reorder"},
	"Vis_Query_create":                                             {"get-all"},
	"Iam_Contracts_update":                                         {"update"},
	"Iam_Contracts_activation":                                     {"activation"},
	"Iam_User_Profile_show":                                        {"list"},
	"Iam_User_Setting_Blob_retrive_get_url":                        {"list"},
	"Iam_User_Setting_Blob_retrive_put_url":                        {"replace"},
	"FinopsOnboarding_Bill_Connect_validate":                       {"validations"},
	"FinopsOnboarding_Bill_Connect_AWS_validate":                   {"validation"},
	"FinopsOnboarding_Bill_Connect_Azure_CSP_validate":             {"validation"},
	"FinopsOnboarding_Bill_Connect_Azure_EA_Management_validate":   {"validation"},
	"FinopsOnboarding_Bill_Connect_Azure_MCA_validate":             {"validation"},
	"FinopsOnboarding_Bill_Connect_Common_Bill_Ingestion_validate": {"validation"},
	"FinopsOnboarding_Bill_Connect_Databricks_validate":            {"validation"},
	"FinopsOnboarding_Bill_Connect_GCP_validate":                   {"validation"},
	"FinopsOnboarding_Bill_Connect_Snowflake_validate":             {"validation"},
}

var versionSegment = regexp.MustCompile(`^v[0-9]+$`)

func isPathParam(seg string) bool { return strings.HasPrefix(seg, "{") }

func isOrgParam(seg string) bool { return seg == "{orgId}" || seg == "{org_id}" }

// namingSegments returns the path segments that carry meaning for command
// names, and whether the operation is scoped to projects/{projectId}.
func namingSegments(path string) (segs []string, projectScoped bool) {
	raw := strings.Split(strings.Trim(path, "/"), "/")
	for i := 0; i < len(raw); i++ {
		seg, next := raw[i], ""
		if i+1 < len(raw) {
			next = raw[i+1]
		}
		switch {
		case seg == "" || versionSegment.MatchString(seg) || isOrgParam(seg):
		case seg == "orgs" && isOrgParam(next):
			i++
		case seg == "projects" && next == "{projectId}":
			projectScoped = true
			i++
		default:
			segs = append(segs, seg)
		}
	}
	return segs, projectScoped
}

// shapeKey identifies a path shape with parameter names erased.
func shapeKey(segs []string) string {
	shape := make([]string, len(segs))
	for i, seg := range segs {
		shape[i] = seg
		if isPathParam(seg) {
			shape[i] = "{}"
		}
	}
	return strings.Join(shape, "/")
}

func commandToken(literals []string) string {
	s := strings.NewReplacer(".", "-", "_", "-").Replace(strings.Join(literals, "-"))
	return kebab(s)
}

// sharedPrefixLen is the number of leading segments shared by every path,
// never ending on a parameter (so an {id} always stays with its operation).
func sharedPrefixLen(paths [][]string) int {
	n := 0
	for ; len(paths) > 0 && n < len(paths[0]); n++ {
		for _, p := range paths[1:] {
			if n >= len(p) || p[n] != paths[0][n] {
				goto done
			}
		}
	}
done:
	for n > 0 && isPathParam(paths[0][n-1]) {
		n--
	}
	return n
}

// commandPaths names every operation of one tag. parent is the tag command
// name. It fails, rather than inventing suffixes, if two operations would
// share a command or a command would also be a group.
func commandPaths(ops []operation, parent string) ([][]string, error) {
	segs := make([][]string, len(ops))
	project := make([]bool, len(ops))
	methodsAtShape := map[string]map[string]bool{}
	for i, o := range ops {
		segs[i], project[i] = namingSegments(o.Path)
		key := shapeKey(segs[i])
		if methodsAtShape[key] == nil {
			methodsAtShape[key] = map[string]bool{}
		}
		methodsAtShape[key][o.Method] = true
	}
	// A resource has an {id} form or several methods; org/project copies of
	// one operation share a shape but not a second method.
	isResource := func(s []string) bool {
		key := shapeKey(s)
		return len(methodsAtShape[key+"/{}"]) > 0 || len(methodsAtShape[key]) > 1
	}
	shared := sharedPrefixLen(segs)

	paths := make([][]string, len(ops))
	for i, o := range ops {
		var runs, run []string
		for _, seg := range segs[i][shared:] {
			if isPathParam(seg) {
				if len(run) > 0 {
					runs = append(runs, commandToken(run))
					run = nil
				}
				continue
			}
			run = append(run, seg)
		}
		if len(run) > 0 {
			runs = append(runs, commandToken(run))
		}

		namedOp := !endsWithParam(segs[i]) && !isResource(segs[i])
		groups, leaf := runs, o.Action
		if len(runs) > 0 && namedOp {
			groups, leaf = runs[:len(runs)-1], runs[len(runs)-1]
		}
		if leaf == "action" {
			leaf = "create"
			if namedOp {
				leaf = commandToken([]string{lastLiteral(segs[i])})
			}
		}
		paths[i] = withoutParentName(append(append([]string(nil), groups...), leaf), parent)
	}

	resolveTwins(segs, project, paths)
	for i, o := range ops {
		if override, ok := commandOverrides[o.OperationID]; ok {
			paths[i] = append([]string(nil), override...)
		}
	}
	return paths, checkCommandPaths(ops, paths)
}

// withoutParentName drops groups named exactly like the parent command and
// trims a "<parent>-" prefix from any name (risk compliance compliance-cis ->
// risk compliance cis).
func withoutParentName(command []string, parent string) []string {
	out := command[:0]
	for i, name := range command {
		if name == parent && i < len(command)-1 {
			continue
		}
		out = append(out, strings.TrimPrefix(name, parent+"-"))
	}
	return out
}

func endsWithParam(segs []string) bool {
	return len(segs) > 0 && isPathParam(segs[len(segs)-1])
}

func lastLiteral(segs []string) string {
	for i := len(segs) - 1; i >= 0; i-- {
		if !isPathParam(segs[i]) {
			return segs[i]
		}
	}
	return ""
}

// resolveTwins applies the two semantic tie-breaks to pairs of operations that
// would otherwise share a command: a collection-level twin of an item-level
// operation gets "<action>-all", and a project-scoped twin of an org-scoped
// operation moves under a "project" group.
func resolveTwins(segs [][]string, project []bool, paths [][]string) {
	byCommand := map[string][]int{}
	for i := range paths {
		key := strings.Join(paths[i], " ")
		byCommand[key] = append(byCommand[key], i)
	}
	for _, idxs := range byCommand {
		if len(idxs) != 2 {
			continue
		}
		a, b := idxs[0], idxs[1]
		aItem, bItem := endsWithParam(segs[a]), endsWithParam(segs[b])
		switch {
		case project[a] == project[b] && aItem != bItem:
			collection := a
			if aItem {
				collection = b
			}
			paths[collection][len(paths[collection])-1] += "-all"
		case project[a] != project[b]:
			scoped := a
			if project[b] {
				scoped = b
			}
			paths[scoped] = append([]string{"project"}, paths[scoped]...)
		}
	}
}

// destructiveWords are the names a destructive command must contain, so a
// command name never hides that it removes something.
var destructiveWords = []string{"delete", "revoke", "deactivate", "remove", "purge"}

func checkCommandPaths(ops []operation, paths [][]string) error {
	owner := map[string]int{}
	groups := map[string]bool{}
	for i, p := range paths {
		key := strings.Join(p, " ")
		if isDestructive(ops[i]) && !containsAny(key, destructiveWords) {
			return fmt.Errorf("destructive operation %s (%s %s) maps to command %q, which does not say so; add a commandOverrides entry",
				ops[i].OperationID, strings.ToUpper(ops[i].Method), ops[i].Path, key)
		}
		if j, dup := owner[key]; dup {
			return fmt.Errorf("operations %s (%s %s) and %s (%s %s) both map to command %q; add a commandOverrides entry",
				ops[j].OperationID, strings.ToUpper(ops[j].Method), ops[j].Path,
				ops[i].OperationID, strings.ToUpper(ops[i].Method), ops[i].Path, key)
		}
		owner[key] = i
		for n := 1; n < len(p); n++ {
			groups[strings.Join(p[:n], " ")] = true
		}
	}
	names := make([]string, 0, len(owner))
	for key := range owner {
		names = append(names, key)
	}
	sort.Strings(names)
	for _, key := range names {
		if groups[key] {
			o := ops[owner[key]]
			return fmt.Errorf("command %q (%s) is also a command group; add a commandOverrides entry", key, o.OperationID)
		}
	}
	aliasOwner := map[string]string{}
	for i, p := range paths {
		for _, alias := range commandAliases[ops[i].OperationID] {
			key := strings.Join(append(append([]string(nil), p[:len(p)-1]...), alias), " ")
			_, isCommand := owner[key]
			other, isAlias := aliasOwner[key]
			if isCommand || groups[key] || isAlias {
				return fmt.Errorf("alias %q of %s collides with an existing command, group or alias %s", key, ops[i].OperationID, other)
			}
			aliasOwner[key] = ops[i].OperationID
		}
	}
	return nil
}

func containsAny(s string, words []string) bool {
	for _, w := range words {
		if strings.Contains(s, w) {
			return true
		}
	}
	return false
}
