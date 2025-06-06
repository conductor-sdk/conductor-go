package integration

import (
	"encoding/json"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model/integration"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserIntegrationDef(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("IntegrationDef")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var integrationDef integration.IntegrationDef
	err = json.Unmarshal([]byte(jsonTemplate), &integrationDef)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// String fields
	if integrationDef.Type != "sample_type" {
		t.Errorf("Expected Type = 'sample_type', got '%s'", integrationDef.Type)
	}
	if integrationDef.Name != "sample_name" {
		t.Errorf("Expected Name = 'sample_name', got '%s'", integrationDef.Name)
	}
	if integrationDef.Description != "sample_description" {
		t.Errorf("Expected Description = 'sample_description', got '%s'", integrationDef.Description)
	}
	if integrationDef.CategoryLabel != "sample_categoryLabel" {
		t.Errorf("Expected CategoryLabel = 'sample_categoryLabel', got '%s'", integrationDef.CategoryLabel)
	}
	if integrationDef.IconName != "sample_iconName" {
		t.Errorf("Expected IconName = 'sample_iconName', got '%s'", integrationDef.IconName)
	}

	// Enum field
	if integrationDef.Category != integration.API {
		t.Errorf("Expected Category = 'API', got '%s'", integrationDef.Category)
	}

	// Boolean field
	if !integrationDef.Enabled {
		t.Errorf("Expected Enabled = true, got %v", integrationDef.Enabled)
	}

	// Check Configuration slice
	if integrationDef.Configuration == nil {
		t.Errorf("Configuration slice should not be nil")
	}
	if len(integrationDef.Configuration) == 0 {
		t.Errorf("Configuration slice should not be empty")
	}
	if len(integrationDef.Configuration) != 1 {
		t.Errorf("Expected Configuration slice to have 1 element, got %d", len(integrationDef.Configuration))
	}

	// Check Tags slice (string slice)
	if integrationDef.Tags == nil {
		t.Errorf("Tags slice should not be nil")
	}
	if len(integrationDef.Tags) == 0 {
		t.Errorf("Tags slice should not be empty")
	}
	if integrationDef.Tags[0] != "sample_tags" {
		t.Errorf("Expected Tags[0] = 'sample_tags', got '%s'", integrationDef.Tags[0])
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(integrationDef)
	if err != nil {
		t.Fatalf("Failed to serialize struct: %v", err)
	}

	// 5. Round-trip integrity check
	// Serialize and deserialize again to compare structs instead of JSON maps
	var roundTripIntegrationDef integration.IntegrationDef
	err = json.Unmarshal(serializedJSON, &roundTripIntegrationDef)
	if err != nil {
		t.Fatalf("Failed to deserialize round-trip JSON: %v", err)
	}

	// Compare the original deserialized struct with the round-trip struct
	// Use a more flexible comparison that handles nil vs empty slice differences
	if !deepEqualWithNilSliceHandling(integrationDef, roundTripIntegrationDef) {
		t.Errorf("Round-trip integrity failed - structs don't match")
		t.Logf("Original JSON: %s", jsonTemplate)
		t.Logf("Serialized JSON: %s", string(serializedJSON))

		// Detailed field comparison for debugging
		if integrationDef.Type != roundTripIntegrationDef.Type {
			t.Logf("Type differs: %s vs %s", integrationDef.Type, roundTripIntegrationDef.Type)
		}
		if len(integrationDef.Configuration) != len(roundTripIntegrationDef.Configuration) {
			t.Logf("Configuration length differs: %d vs %d", len(integrationDef.Configuration), len(roundTripIntegrationDef.Configuration))
		}
	}
}

// Helper function to handle nil vs empty slice differences in deep comparison
func deepEqualWithNilSliceHandling(a, b interface{}) bool {
	// For this specific test, we'll use a simpler approach:
	// Serialize both structs and compare the JSON strings
	aJSON, err1 := json.Marshal(a)
	bJSON, err2 := json.Marshal(b)

	if err1 != nil || err2 != nil {
		return false
	}

	return string(aJSON) == string(bJSON)
}
