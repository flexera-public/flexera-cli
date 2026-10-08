package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/flexera-public/flexera-cli/internal/catalog"
)

func TestBuildCatalog(t *testing.T) {
	stage := t.TempDir()
	entry := catalog.Entry{OperationID: "test", Command: []string{"bill-connect-aws", "create"}, Method: "POST", Path: "/test", ResponseEnvelope: "none", RequestSchema: json.RawMessage(`{"$ref":"#/components/schemas/A"}`)}
	data, err := json.Marshal(map[string]catalog.Entry{"test": entry})
	if err != nil {
		t.Fatal(err)
	}
	writeTestFile(t, filepath.Join(stage, "finopsonboarding", "billconnectaws", "metadata.json"), string(data))
	writeTestFile(t, filepath.Join(stage, "finopsonboarding", "billconnect", "metadata.json"), `{}`)
	specData := []byte(`{"paths":{"/test":{"post":{"operationId":"test"}}},"components":{"schemas":{"A":{"type":"object","properties":{"b":{"$ref":"#/components/schemas/B"}}},"B":{"type":"object","properties":{"a":{"$ref":"#/components/schemas/A"}}},"Unused":{"type":"string"}}}}`)
	tags := []genTag{
		{Service: "finops_onboarding", Pkg: "finopsonboarding/billconnectaws", Cmd: "bill-connect-aws", Path: []string{"finops-onboarding", "bill-connect", "aws"}},
		{Service: "finops_onboarding", Pkg: "finopsonboarding/billconnect", Cmd: "bill-connect", Path: []string{"finops-onboarding", "bill-connect"}},
	}
	first, err := buildCatalog(stage, tags, specData)
	if err != nil {
		t.Fatal(err)
	}
	second, err := buildCatalog(stage, []genTag{tags[1], tags[0]}, specData)
	if err != nil || !bytes.Equal(first, second) {
		t.Fatalf("nondeterministic output: %v", err)
	}
	var doc catalog.Document
	if err := json.Unmarshal(first, &doc); err != nil {
		t.Fatal(err)
	}
	if len(doc.Schemas) != 2 || len(doc.Entries) != 1 {
		t.Fatalf("unexpected closure: %s", first)
	}
	if got := doc.Entries[0].Command; strings.Join(got, " ") != "finops-onboarding bill-connect aws create" {
		t.Fatalf("bad final path: %v", got)
	}
	if doc.Entries[0].Service != "finops_onboarding" {
		t.Fatalf("service not recorded: %q", doc.Entries[0].Service)
	}
	if _, err := buildCatalog(stage, tags, []byte(`{"paths":{"/test":{"get":{"operationId":"test"}}}}`)); err == nil {
		t.Fatal("accepted mismatched method")
	}
	if _, err := buildCatalog(stage, tags, []byte(`{"paths":{"/test":{"post":{"operationId":"test"}}},"components":{"schemas":{}}}`)); err == nil {
		t.Fatal("accepted dangling reference")
	}
	if err := os.Remove(filepath.Join(stage, "finopsonboarding", "billconnectaws", "metadata.json")); err != nil {
		t.Fatal(err)
	}
	if _, err := buildCatalog(stage, tags, specData); err == nil {
		t.Fatal("accepted missing metadata")
	}
}
