package main

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/flexera-public/flexera-cli/internal/catalog"
)

func TestCoverageRejectsUnclassifiedAndUnexpectedOperations(t *testing.T) {
	spec := []byte(`{"paths":{"/widgets":{"get":{"operationId":"Widget_list","x-flexera-action":"list"}}}}`)
	doc := catalog.Document{Entries: []catalog.Entry{}, Schemas: map[string]json.RawMessage{}}
	data, _ := json.Marshal(doc)
	if _, err := verifyCoverage(spec, data); err == nil || !strings.Contains(err.Error(), "unclassified operation Widget_list") {
		t.Fatalf("accepted incomplete coverage: %v", err)
	}
	doc.Entries = append(doc.Entries, catalog.Entry{OperationID: "Widget_list", Command: []string{"widget", "list"}})
	data, _ = json.Marshal(doc)
	first, err := verifyCoverage(spec, data)
	if err != nil {
		t.Fatal(err)
	}
	second, err := verifyCoverage(spec, data)
	if err != nil || string(first) != string(second) {
		t.Fatal("coverage report is nondeterministic")
	}
	doc.Entries = append(doc.Entries, catalog.Entry{OperationID: "Unexpected", Command: []string{"widget", "other"}})
	data, _ = json.Marshal(doc)
	if _, err := verifyCoverage(spec, data); err == nil || !strings.Contains(err.Error(), "unexpected catalog operation") {
		t.Fatalf("accepted unexpected operation: %v", err)
	}
}
