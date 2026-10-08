package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/flexera-public/flexera-cli/internal/catalog"
)

// Only exact wrappers belong here; workflow commands intentionally do not.
var curatedOperations = map[string]string{
	"Auth_Token_token":                 "auth token create",
	"Iam_User_Memberships_index":       "iam user-memberships orgs",
	"Policy_Action_Status_index":       "policy action-status list",
	"Policy_Action_Status_show":        "policy action-status get",
	"Policy_Applied_Policy_index":      "policy applied-policy list",
	"Policy_Applied_Policy_show":       "policy applied-policy get",
	"Policy_Applied_Policy_create":     "policy applied-policy create",
	"Policy_Applied_Policy_update":     "policy applied-policy update",
	"Policy_Applied_Policy_delete":     "policy applied-policy delete",
	"Policy_Applied_Policy_evaluate":   "policy applied-policy evaluate",
	"Policy_Applied_Policy_showLog":    "policy applied-policy log",
	"Policy_Applied_Policy_showStatus": "policy applied-policy status",
	"Policy_Archived_Incident_index":   "policy archived-incident list",
	"Policy_Archived_Incident_show":    "policy archived-incident get",
	"Policy_Policy_Template_index":     "policy policy-template list",
	"Policy_Policy_Template_show":      "policy policy-template get",
	"Policy_Policy_Template_create":    "policy policy-template create",
	"Policy_Policy_Template_validate":  "policy policy-template validate",
	"Policy_Policy_Template_update":    "policy policy-template update",
	"Policy_Policy_Template_delete":    "policy policy-template delete",
	"Policy_Policy_Template_evaluate":  "policy policy-template evaluate",
}

func curatedOperationIDs() []string {
	ids := make([]string, 0, len(curatedOperations))
	for id := range curatedOperations {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	return ids
}

func curatedEntries(specData []byte) ([]catalog.Entry, error) {
	var s spec
	if err := json.Unmarshal(specData, &s); err != nil {
		return nil, err
	}
	var entries []catalog.Entry
	for path, pi := range s.Paths {
		for method, raw := range pi {
			var op map[string]json.RawMessage
			if json.Unmarshal(raw, &op) != nil {
				continue
			}
			var id string
			_ = json.Unmarshal(op["operationId"], &id)
			command, owned := curatedOperations[id]
			if !owned {
				continue
			}
			e := catalog.Entry{OperationID: id, Command: strings.Fields(command), Method: strings.ToUpper(method), Path: path, ResponseEnvelope: "none", Params: []catalog.Param{}, BodyFlags: []string{}}
			_ = json.Unmarshal(op["summary"], &e.Summary)
			_ = json.Unmarshal(op["description"], &e.Description)
			_ = json.Unmarshal(op["x-flexera-resource"], &e.Resource)
			_ = json.Unmarshal(op["x-flexera-action"], &e.Action)
			_ = json.Unmarshal(op["x-flexera-service"], &e.Service)
			var tags []string
			_ = json.Unmarshal(op["tags"], &tags)
			if len(tags) > 0 {
				e.Tag = tags[0]
			}
			// These exact curated wrappers do not expose --yes. Destructive is
			// the existing confirmation contract, not an HTTP-write indicator.
			e.Destructive = false
			if e.Action == "list" {
				var paginated bool
				_ = json.Unmarshal(op["x-flexera-paginated"], &paginated)
				e.Paginated = paginated
				if paginated {
					e.ResponseEnvelope = "values"
				}
			}
			var rb struct {
				Content map[string]struct {
					Schema json.RawMessage `json:"schema"`
				} `json:"content"`
			}
			_ = json.Unmarshal(op["requestBody"], &rb)
			for _, media := range []string{"application/json", "application/x-www-form-urlencoded"} {
				if c, ok := rb.Content[media]; ok {
					e.RequestSchema = c.Schema
					break
				}
			}
			var responses map[string]struct {
				Content map[string]struct {
					Schema json.RawMessage `json:"schema"`
				} `json:"content"`
			}
			_ = json.Unmarshal(op["responses"], &responses)
			codes := []string{}
			for code := range responses {
				if len(code) == 3 && code >= "200" && code <= "299" {
					codes = append(codes, code)
				}
			}
			sort.Strings(codes)
			if len(codes) > 0 {
				e.ResponseSchema = responses[codes[0]].Content["application/json"].Schema
			}
			// Curated wrappers have adapters distinct from generated flags. Their
			// precise API parameters remain available as schema metadata.
			var params []struct {
				Name, In, Description string
				Required              bool
				Schema                json.RawMessage
			}
			_ = json.Unmarshal(op["parameters"], &params)
			for _, p := range params {
				if p.In != "path" && p.In != "query" {
					continue
				}
				var schema struct {
					Type string
					Enum json.RawMessage
				}
				_ = json.Unmarshal(p.Schema, &schema)
				flag := kebab(p.Name)
				source := "flag"
				if flag == "org-id" {
					source = "config"
				}
				if p.Name == "appliedPolicyId" || p.Name == "actionStatusId" || p.Name == "archivedIncidentId" || p.Name == "policyTemplateId" {
					flag = "id"
				}
				e.Params = append(e.Params, catalog.Param{Flag: flag, Source: source, In: p.In, Type: schema.Type, Required: p.Required, Schema: p.Schema, Enum: schema.Enum, Description: p.Description})
			}
			entries = append(entries, e)
		}
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].OperationID < entries[j].OperationID })
	return entries, nil
}

func verifyCoverage(specData, catalogData []byte) ([]byte, error) {
	var s spec
	if err := json.Unmarshal(specData, &s); err != nil {
		return nil, err
	}
	var doc catalog.Document
	if err := json.Unmarshal(catalogData, &doc); err != nil {
		return nil, err
	}
	seen := map[string]catalog.Entry{}
	for _, e := range doc.Entries {
		seen[e.OperationID] = e
	}
	type disposition struct {
		OperationID string   `json:"operationId"`
		Command     []string `json:"command"`
		Ownership   string   `json:"ownership"`
	}
	report := []disposition{}
	expected := map[string]bool{}
	for _, pi := range s.Paths {
		for _, raw := range pi {
			var op map[string]json.RawMessage
			if json.Unmarshal(raw, &op) != nil {
				continue
			}
			if _, ok := op["x-flexera-action"]; !ok {
				continue
			}
			var id string
			_ = json.Unmarshal(op["operationId"], &id)
			if id == "" || expected[id] {
				return nil, fmt.Errorf("coverage: missing or duplicate operationId %q", id)
			}
			expected[id] = true
			e, ok := seen[id]
			if !ok {
				return nil, fmt.Errorf("coverage: unclassified operation %s", id)
			}
			ownership := "generated"
			if _, ok := curatedOperations[id]; ok {
				ownership = "exact-curated"
			}
			report = append(report, disposition{id, e.Command, ownership})
		}
	}
	for id := range seen {
		if !expected[id] {
			return nil, fmt.Errorf("coverage: unexpected catalog operation %s", id)
		}
	}
	sort.Slice(report, func(i, j int) bool { return report[i].OperationID < report[j].OperationID })
	data, err := json.MarshalIndent(report, "", "  ")
	return append(data, '\n'), err
}
