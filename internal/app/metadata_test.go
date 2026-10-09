package app_test

import (
	"encoding/json"
	"io"
	"strings"
	"testing"

	"github.com/flexera-public/flexera-cli/internal/app"
	"github.com/flexera-public/flexera-cli/internal/catalog"
)

func TestPublishedInputMetadataMatchesFlags(t *testing.T) {
	root, _ := app.NewRootCmd(io.Discard, io.Discard, func(string) string { return "" }, offlineDoer{t}, "test")
	entries, err := catalog.All()
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range entries {
		leaf, _, err := root.Find(entry.Command)
		if err != nil {
			t.Fatal(err)
		}
		for _, param := range entry.Params {
			if param.Name == "" {
				t.Errorf("%s --%s has no API parameter name", entry.OperationID, param.Flag)
			}
		}
		for _, field := range entry.BodyFields {
			flag := leaf.Flags().Lookup(field.Flag)
			if flag == nil || field.Property == "" || !json.Valid(field.Schema) {
				t.Errorf("%s invalid body field: %+v", entry.OperationID, field)
				continue
			}
			if field.Required && !strings.Contains(flag.Usage, "required by API") {
				t.Errorf("%s --%s missing requiredness guidance", entry.OperationID, field.Flag)
			}
		}
	}
}

func TestSchemaPublishesInputGuidanceOffline(t *testing.T) {
	code, out, stderr := schemaRun(t, "cli", "schema", "finops-billing", "invoice-templates", "list")
	if code != 0 {
		t.Fatalf("%d: %s", code, stderr)
	}
	var entry catalog.Entry
	if err := json.Unmarshal([]byte(out), &entry); err != nil {
		t.Fatal(err)
	}
	for _, param := range entry.Params {
		if param.Flag == "filter" {
			if param.Name != "filter" || !strings.Contains(param.Description, "OData-style") || !strings.Contains(string(param.Examples), "templateName co") {
				t.Fatalf("filter guidance missing: %+v", param)
			}
			return
		}
	}
	t.Fatal("filter parameter missing")
}
