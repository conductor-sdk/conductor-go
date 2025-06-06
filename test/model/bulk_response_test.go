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

	// Test backward compatibility - existing code should still work
	if bulkResponse.BulkSuccessfulResults == nil {
		t.Errorf("BulkSuccessfulResults should not be nil for backward compatibility")
	}

	if len(bulkResponse.BulkSuccessfulResults) == 0 {
		t.Errorf("BulkSuccessfulResults slice should not be empty")
	}

	// Test new functionality
	successfulResults := bulkResponse.GetSuccessfulResults()
	if successfulResults == nil {
		t.Errorf("GetSuccessfulResults should not return nil")
	}

	// Test that both approaches give consistent results
	stringResults := bulkResponse.GetSuccessfulResultsAsStrings()
	if len(stringResults) != len(bulkResponse.BulkSuccessfulResults) {
		t.Errorf("Inconsistent results between old and new methods")
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
	var originalStruct, roundTripStruct model.BulkResponse

	err = json.Unmarshal([]byte(jsonTemplate), &originalStruct)
	if err != nil {
		t.Fatalf("Failed to deserialize original JSON: %v", err)
	}

	err = json.Unmarshal(serializedJSON, &roundTripStruct)
	if err != nil {
		t.Fatalf("Failed to deserialize round-trip JSON: %v", err)
	}

	// Compare the important fields instead of raw JSON
	if !reflect.DeepEqual(originalStruct.BulkErrorResults, roundTripStruct.BulkErrorResults) {
		t.Errorf("BulkErrorResults mismatch after round-trip")
	}

	if originalStruct.Message != roundTripStruct.Message {
		t.Errorf("Message mismatch after round-trip")
	}

	// Compare successful results semantically
	originalResults := originalStruct.GetSuccessfulResultsAsStrings()
	roundTripResults := roundTripStruct.GetSuccessfulResultsAsStrings()
	if !reflect.DeepEqual(originalResults, roundTripResults) {
		t.Errorf("BulkSuccessfulResults content mismatch after round-trip")
	}
}
