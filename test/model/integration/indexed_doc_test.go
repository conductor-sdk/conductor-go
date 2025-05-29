package integration

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model/integration"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserIndexedDoc(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("IndexedDoc")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var indexedDoc integration.IndexedDoc
	err = json.Unmarshal([]byte(jsonTemplate), &indexedDoc)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// String fields
	if indexedDoc.DocId != "sample_docId" {
		t.Errorf("Expected DocId = 'sample_docId', got '%s'", indexedDoc.DocId)
	}
	if indexedDoc.ParentDocId != "sample_parentDocId" {
		t.Errorf("Expected ParentDocId = 'sample_parentDocId', got '%s'", indexedDoc.ParentDocId)
	}
	if indexedDoc.Text != "sample_text" {
		t.Errorf("Expected Text = 'sample_text', got '%s'", indexedDoc.Text)
	}

	// Float field
	if indexedDoc.Score != 123.456 {
		t.Errorf("Expected Score = 123.456, got %f", indexedDoc.Score)
	}

	// Check map field
	if indexedDoc.Metadata == nil {
		t.Errorf("Metadata map should not be nil")
	}
	if len(indexedDoc.Metadata) == 0 {
		t.Errorf("Metadata map should not be empty")
	}
	if indexedDoc.Metadata["sample_key"] != "sample_value" {
		t.Errorf("Expected Metadata['sample_key'] = 'sample_value', got %v", indexedDoc.Metadata["sample_key"])
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(indexedDoc)
	if err != nil {
		t.Fatalf("Failed to serialize struct: %v", err)
	}

	// 5. Round-trip integrity check
	var originalMap, serializedMap map[string]interface{}
	json.Unmarshal([]byte(jsonTemplate), &originalMap)
	json.Unmarshal(serializedJSON, &serializedMap)

	if !reflect.DeepEqual(originalMap, serializedMap) {
		t.Errorf("Round-trip integrity failed")
	}
}
