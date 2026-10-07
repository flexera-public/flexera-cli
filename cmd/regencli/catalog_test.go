package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
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
	writeTestFile(t, filepath.Join(stage, "billconnectaws", "metadata.json"), string(data))
	writeTestFile(t, filepath.Join(stage, "billconnect", "metadata.json"), `{}`)
	specData := []byte(`{"paths":{"/test":{"post":{"operationId":"test"}}},"components":{"schemas":{"A":{"type":"object","properties":{"b":{"$ref":"#/components/schemas/B"}}},"B":{"type":"object","properties":{"a":{"$ref":"#/components/schemas/A"}}},"Unused":{"type":"string"}}}}`)
	tags := []genTag{{Pkg: "billconnectaws", Cmd: "bill-connect-aws"}, {Pkg: "billconnect", Cmd: "bill-connect"}}
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
	if got := doc.Entries[0].Command; len(got) != 3 || got[0] != "bill-connect" || got[1] != "aws" || got[2] != "create" {
		t.Fatalf("bad final path: %v", got)
	}
	if _, err := buildCatalog(stage, tags, []byte(`{"paths":{"/test":{"get":{"operationId":"test"}}}}`)); err == nil {
		t.Fatal("accepted mismatched method")
	}
	if _, err := buildCatalog(stage, tags, []byte(`{"paths":{"/test":{"post":{"operationId":"test"}}},"components":{"schemas":{}}}`)); err == nil {
		t.Fatal("accepted dangling reference")
	}
	if err := os.Remove(filepath.Join(stage, "billconnectaws", "metadata.json")); err != nil {
		t.Fatal(err)
	}
	if _, err := buildCatalog(stage, tags, specData); err == nil {
		t.Fatal("accepted missing metadata")
	}
}
