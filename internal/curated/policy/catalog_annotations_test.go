package policy

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestCatalogOperationAnnotations(t *testing.T) {
	want := map[string]string{
		"policy applied-policy list":      "Policy_Applied_Policy_index",
		"policy applied-policy get":       "Policy_Applied_Policy_show",
		"policy applied-policy create":    "Policy_Applied_Policy_create",
		"policy applied-policy update":    "Policy_Applied_Policy_update",
		"policy applied-policy delete":    "Policy_Applied_Policy_delete",
		"policy applied-policy evaluate":  "Policy_Applied_Policy_evaluate",
		"policy applied-policy log":       "Policy_Applied_Policy_showLog",
		"policy applied-policy status":    "Policy_Applied_Policy_showStatus",
		"policy action-status list":       "Policy_Action_Status_index",
		"policy action-status get":        "Policy_Action_Status_show",
		"policy archived-incident list":   "Policy_Archived_Incident_index",
		"policy archived-incident get":    "Policy_Archived_Incident_show",
		"policy policy-template list":     "Policy_Policy_Template_index",
		"policy policy-template get":      "Policy_Policy_Template_show",
		"policy policy-template create":   "Policy_Policy_Template_create",
		"policy policy-template validate": "Policy_Policy_Template_validate",
		"policy policy-template update":   "Policy_Policy_Template_update",
		"policy policy-template delete":   "Policy_Policy_Template_delete",
		"policy policy-template evaluate": "Policy_Policy_Template_evaluate",
	}
	const annotation = "flexera.operationId"
	seen := make(map[string]bool)
	count := 0
	var walk func(*cobra.Command)
	walk = func(cmd *cobra.Command) {
		path := cmd.CommandPath()
		got, annotated := cmd.Annotations[annotation]
		if expected, ok := want[path]; ok {
			seen[path] = true
			if !annotated || got != expected {
				t.Errorf("%s: operationId = %q, want %q", path, got, expected)
			}
			if cmd.HasSubCommands() {
				t.Errorf("%s: expected a direct API leaf", path)
			}
		} else if annotated {
			t.Errorf("%s: unexpected operationId annotation %q (group/meta commands must be unannotated)", path, got)
		}
		if annotated {
			count++
		}
		for _, child := range cmd.Commands() {
			walk(child)
		}
	}
	walk(NewCmd())
	if count != 19 {
		t.Errorf("annotated command count = %d, want 19", count)
	}
	for path := range want {
		if !seen[path] {
			t.Errorf("missing direct API leaf %s", path)
		}
	}
}
