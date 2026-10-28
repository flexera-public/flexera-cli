package main

import (
	"strings"
	"testing"
)

func namesFor(t *testing.T, parent string, ops []operation) map[string]string {
	t.Helper()
	paths, err := commandPaths(ops, parent)
	if err != nil {
		t.Fatal(err)
	}
	got := map[string]string{}
	for i, o := range ops {
		got[o.OperationID] = strings.Join(paths[i], " ")
	}
	return got
}

func assertNames(t *testing.T, got, want map[string]string) {
	t.Helper()
	for id, name := range want {
		if got[id] != name {
			t.Errorf("%s = %q, want %q", id, got[id], name)
		}
	}
}

func TestCommandPathsNameSubResourcesByAction(t *testing.T) {
	ops := []operation{
		{OperationID: "License_get", Method: "get", Action: "get", Path: "/saas/v1/orgs/{orgId}/licenses/{licenseId}"},
		{OperationID: "Fees_list", Method: "get", Action: "list", Path: "/saas/v1/orgs/{orgId}/licenses/{licenseId}/fees"},
		{OperationID: "Fees_create", Method: "post", Action: "create", Path: "/saas/v1/orgs/{orgId}/licenses/{licenseId}/fees"},
		{OperationID: "Fees_get", Method: "get", Action: "get", Path: "/saas/v1/orgs/{orgId}/licenses/{licenseId}/fees/{feeId}"},
		{OperationID: "Fees_delete", Method: "delete", Action: "delete", Path: "/saas/v1/orgs/{orgId}/licenses/{licenseId}/fees/{feeId}"},
		{OperationID: "Allocations_list", Method: "get", Action: "list", Path: "/saas/v1/orgs/{orgId}/licenses/{licenseId}/terms/{termId}/purchases/{purchaseId}/allocations"},
		{OperationID: "Allocations_create", Method: "post", Action: "create", Path: "/saas/v1/orgs/{orgId}/licenses/{licenseId}/terms/{termId}/purchases/{purchaseId}/allocations"},
		{OperationID: "License_report", Method: "get", Action: "get", Path: "/saas/v1/orgs/{orgId}/licenses/{licenseId}/usage-report"},
		{OperationID: "License_activate", Method: "post", Action: "action", Path: "/saas/v1/orgs/{orgId}/licenses/{licenseId}/activate"},
	}
	want := map[string]string{
		"License_get":        "get",
		"Fees_list":          "fees list",
		"Fees_create":        "fees create",
		"Fees_get":           "fees get",
		"Fees_delete":        "fees delete",
		"Allocations_create": "terms purchases allocations create",
		"License_report":     "usage-report",
		"License_activate":   "activate",
	}
	assertNames(t, namesFor(t, "license", ops), want)

	reversed := make([]operation, len(ops))
	for i, o := range ops {
		reversed[len(ops)-1-i] = o
	}
	assertNames(t, namesFor(t, "license", reversed), want)
}

func TestCommandPathsTwinsAndStutter(t *testing.T) {
	ops := []operation{
		{OperationID: "Rules_replace_all", Method: "put", Action: "replace", Path: "/v1/orgs/{orgId}/rules"},
		{OperationID: "Rules_replace", Method: "put", Action: "replace", Path: "/v1/orgs/{orgId}/rules/{id}"},
		{OperationID: "Key_get", Method: "get", Action: "get", Path: "/v1/orgs/{orgId}/keys/{id}"},
		{OperationID: "Key_get_project", Method: "get", Action: "get", Path: "/v1/orgs/{orgId}/projects/{projectId}/keys/{id}"},
		{OperationID: "Compliance_cis", Method: "get", Action: "get", Path: "/v1/orgs/{orgId}/compliance-cis"},
		{OperationID: "Token_create", Method: "post", Action: "action", Path: "/v1/orgs/{orgId}/tokens"},
		{OperationID: "Evaluate", Method: "post", Action: "action", Path: "/v1/orgs/{orgId}/evaluate"},
		{OperationID: "Evaluate_project", Method: "post", Action: "action", Path: "/v1/orgs/{orgId}/projects/{projectId}/evaluate"},
	}
	assertNames(t, namesFor(t, "compliance", ops), map[string]string{
		"Rules_replace_all": "rules replace-all",
		"Rules_replace":     "rules replace",
		"Key_get":           "keys get",
		"Key_get_project":   "project keys get",
		"Compliance_cis":    "cis",
		"Token_create":      "tokens",
		"Evaluate":          "evaluate",
		"Evaluate_project":  "project evaluate",
	})
}

func TestCommandPathsRejectAmbiguity(t *testing.T) {
	collision := []operation{
		{OperationID: "A", Method: "get", Action: "list", Path: "/v1/widgets"},
		{OperationID: "B", Method: "get", Action: "list", Path: "/v2/widgets"},
		{OperationID: "C", Method: "get", Action: "list", Path: "/v3/widgets"},
	}
	if _, err := commandPaths(collision, "widget"); err == nil || !strings.Contains(err.Error(), "both map to command") {
		t.Fatalf("collision error = %v", err)
	}
	leafAndGroup := []operation{
		{OperationID: "List", Method: "get", Action: "list", Path: "/v1/things"},
		{OperationID: "Report", Method: "get", Action: "get", Path: "/v1/things/{id}/report"},
		{OperationID: "Report_rows", Method: "get", Action: "list", Path: "/v1/things/{id}/report/{reportId}/rows"},
	}
	if _, err := commandPaths(leafAndGroup, "thing"); err == nil || !strings.Contains(err.Error(), "also a command group") {
		t.Fatalf("leaf/group error = %v", err)
	}
}

func TestCommandPathsRejectHiddenDestructiveOps(t *testing.T) {
	ops := []operation{
		{OperationID: "Estimate", Method: "post", Action: "get", Path: "/v1/devices/delete/estimate"},
		{OperationID: "Submit", Method: "post", Action: "action", Path: "/v1/devices/delete/submit"},
	}
	if _, err := commandPaths(ops, "device"); err == nil || !strings.Contains(err.Error(), "does not say so") {
		t.Fatalf("hidden destructive error = %v", err)
	}
}

func TestIsDestructiveUsesMethodAndPath(t *testing.T) {
	for _, tc := range []struct {
		o    operation
		want bool
	}{
		{operation{Method: "delete", Action: "delete", Path: "/things/{id}"}, true},
		{operation{Method: "put", Action: "action", Path: "/tokens/{id}/revoke"}, true},
		{operation{Method: "post", Action: "create", Path: "/customers/bulk-delete"}, true},
		{operation{Method: "post", Action: "action", Path: "/devices/delete/submit"}, true},
		{operation{Method: "put", Action: "replace", Path: "/tokens/revoke-all"}, true},
		{operation{Method: "post", Action: "action", Path: "/tokens/delete-all"}, true},
		{operation{Method: "post", Action: "action", Path: "/things/delete-all/{id}"}, true},
		{operation{Method: "post", Action: "action", Path: "/schedules/{id}/deactivate"}, true},
		{operation{Method: "get", Action: "list", Path: "/tokens/revoke-all"}, false},
		{operation{Method: "get", Action: "get", Path: "/devices/delete/estimate"}, false},
		{operation{Method: "post", Action: "create", Path: "/things"}, false},
	} {
		if got := isDestructive(tc.o); got != tc.want {
			t.Errorf("isDestructive(%s %s) = %v, want %v", tc.o.Method, tc.o.Path, got, tc.want)
		}
	}
}

func TestCommandAliasesRejectCollisions(t *testing.T) {
	ops := []operation{
		{OperationID: "Thing_list", Method: "get", Action: "list", Path: "/v1/things"},
		{OperationID: "Thing_resend", Method: "patch", Action: "update", Path: "/v1/things/{id}"},
	}
	defer func(old map[string][]string) { commandAliases = old }(commandAliases)
	commandAliases = map[string][]string{"Thing_resend": {"old-resend"}}
	if _, err := commandPaths(ops, "thing"); err != nil {
		t.Fatalf("distinct alias rejected: %v", err)
	}
	commandAliases = map[string][]string{"Thing_resend": {"list"}}
	if _, err := commandPaths(ops, "thing"); err == nil || !strings.Contains(err.Error(), "collides") {
		t.Fatalf("alias collision error = %v", err)
	}
}
