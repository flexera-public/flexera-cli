package catalog

import (
	"encoding/json"
	"testing"
)

func TestRequestSchemaCachedAndReferencesResolved(t *testing.T) {
	doc := testDocument()
	doc.Entries[0].RequestSchema = json.RawMessage(`{"$ref":"#/components/schemas/Widget"}`)
	data, _ := json.Marshal(doc)
	index, err := Parse(data)
	if err != nil {
		t.Fatal(err)
	}
	first, err := index.RequestSchema(doc.Entries[0].RequestSchema)
	if err != nil {
		t.Fatal(err)
	}
	second, err := index.RequestSchema(doc.Entries[0].RequestSchema)
	if err != nil || first != second {
		t.Fatal("schema not cached")
	}
	if first.Properties["child"].Value != first {
		t.Fatal("cyclic reference not resolved")
	}
}
