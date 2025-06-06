package model

import (
	"encoding/json"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserSearchResult(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("SearchResult")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var searchResult model.SearchResult
	err = json.Unmarshal([]byte(jsonTemplate), &searchResult)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// Integer field
	if searchResult.TotalHits != 123 {
		t.Errorf("Expected TotalHits = 123, got %d", searchResult.TotalHits)
	}

	// Results slice field
	if searchResult.Results == nil {
		t.Errorf("Results slice should not be nil")
	}
	if len(searchResult.Results) == 0 {
		t.Errorf("Results slice should not be empty")
	}
	if len(searchResult.Results) != 1 {
		t.Errorf("Expected Results slice to have 1 element, got %d", len(searchResult.Results))
	}

	// Validate the result element
	result := searchResult.Results[0]
	if result == nil {
		t.Errorf("Results[0] should not be nil")
	}

	// The template resolves T to a map structure
	if resultMap, ok := result.(map[string]interface{}); ok {
		if resultMap["value"] != "sample_string_value" {
			t.Errorf("Expected Results[0]['value'] = 'sample_string_value', got %v", resultMap["value"])
		}
	} else {
		t.Errorf("Expected Results[0] to be a map[string]interface{}, got %T: %v", result, result)
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(searchResult)
	if err != nil {
		t.Fatalf("Failed to serialize struct: %v", err)
	}

	// 5. Round-trip integrity check
	var roundTripSearchResult model.SearchResult
	err = json.Unmarshal(serializedJSON, &roundTripSearchResult)
	if err != nil {
		t.Fatalf("Failed to deserialize round-trip JSON: %v", err)
	}

	// Compare key fields to ensure data integrity
	if roundTripSearchResult.TotalHits != searchResult.TotalHits {
		t.Errorf("Round-trip TotalHits mismatch: %d vs %d", searchResult.TotalHits, roundTripSearchResult.TotalHits)
	}
	if len(roundTripSearchResult.Results) != len(searchResult.Results) {
		t.Errorf("Round-trip Results length mismatch: %d vs %d", len(searchResult.Results), len(roundTripSearchResult.Results))
	}
}
