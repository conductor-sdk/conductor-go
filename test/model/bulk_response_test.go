package model

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserBulkResponse(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("BulkResponse")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var bulkResponse model.BulkResponse
	err = json.Unmarshal([]byte(jsonTemplate), &bulkResponse)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// Check map field
	if bulkResponse.BulkErrorResults == nil {
		t.Errorf("BulkErrorResults map should not be nil")
	}
	if len(bulkResponse.BulkErrorResults) == 0 {
		t.Errorf("BulkErrorResults map should not be empty")
	}
	if bulkResponse.BulkErrorResults["sample_key"] != "sample_value" {
		t.Errorf("Expected BulkErrorResults['sample_key'] = 'sample_value', got %v", bulkResponse.BulkErrorResults["sample_key"])
	}

	// Check interface{} field (should be a slice based on template)
	if bulkResponse.BulkSuccessfulResults == nil {
		t.Errorf("BulkSuccessfulResults should not be nil")
	}
	successfulResults, ok := bulkResponse.BulkSuccessfulResults.([]interface{})
	if !ok {
		t.Errorf("BulkSuccessfulResults should be a slice, got %T", bulkResponse.BulkSuccessfulResults)
	}
	if len(successfulResults) == 0 {
		t.Errorf("BulkSuccessfulResults slice should not be empty")
	}

	// Check string field
	if bulkResponse.Message != "sample_message" {
		t.Errorf("Expected Message = 'sample_message', got '%s'", bulkResponse.Message)
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(bulkResponse)
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
