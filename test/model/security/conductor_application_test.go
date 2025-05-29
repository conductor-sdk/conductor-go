package security

import (
	"encoding/json"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model/security"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserConductorApplication(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("ConductorApplication")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var conductorApplication security.ConductorApplication
	err = json.Unmarshal([]byte(jsonTemplate), &conductorApplication)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// String fields
	if conductorApplication.ID != "sample_id" {
		t.Errorf("Expected ID = 'sample_id', got '%s'", conductorApplication.ID)
	}
	if conductorApplication.Name != "sample_name" {
		t.Errorf("Expected Name = 'sample_name', got '%s'", conductorApplication.Name)
	}
	if conductorApplication.CreatedBy != "sample_createdBy" {
		t.Errorf("Expected CreatedBy = 'sample_createdBy', got '%s'", conductorApplication.CreatedBy)
	}
	if conductorApplication.UpdatedBy != "sample_updatedBy" {
		t.Errorf("Expected UpdatedBy = 'sample_updatedBy', got '%s'", conductorApplication.UpdatedBy)
	}

	// Int64 fields
	if conductorApplication.CreateTime != 123 {
		t.Errorf("Expected CreateTime = 123, got %d", conductorApplication.CreateTime)
	}
	if conductorApplication.UpdateTime != 123 {
		t.Errorf("Expected UpdateTime = 123, got %d", conductorApplication.UpdateTime)
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(conductorApplication)
	if err != nil {
		t.Fatalf("Failed to serialize struct: %v", err)
	}

	// 5. Round-trip integrity check
	var roundTripApplication security.ConductorApplication
	err = json.Unmarshal(serializedJSON, &roundTripApplication)
	if err != nil {
		t.Fatalf("Failed to deserialize round-trip JSON: %v", err)
	}

	// Compare key fields to ensure data integrity
	if roundTripApplication.ID != conductorApplication.ID {
		t.Errorf("Round-trip ID mismatch: %s vs %s", conductorApplication.ID, roundTripApplication.ID)
	}
	if roundTripApplication.Name != conductorApplication.Name {
		t.Errorf("Round-trip Name mismatch: %s vs %s", conductorApplication.Name, roundTripApplication.Name)
	}
	if roundTripApplication.CreatedBy != conductorApplication.CreatedBy {
		t.Errorf("Round-trip CreatedBy mismatch: %s vs %s", conductorApplication.CreatedBy, roundTripApplication.CreatedBy)
	}
	if roundTripApplication.UpdatedBy != conductorApplication.UpdatedBy {
		t.Errorf("Round-trip UpdatedBy mismatch: %s vs %s", conductorApplication.UpdatedBy, roundTripApplication.UpdatedBy)
	}
	if roundTripApplication.CreateTime != conductorApplication.CreateTime {
		t.Errorf("Round-trip CreateTime mismatch: %d vs %d", conductorApplication.CreateTime, roundTripApplication.CreateTime)
	}
	if roundTripApplication.UpdateTime != conductorApplication.UpdateTime {
		t.Errorf("Round-trip UpdateTime mismatch: %d vs %d", conductorApplication.UpdateTime, roundTripApplication.UpdateTime)
	}
}
