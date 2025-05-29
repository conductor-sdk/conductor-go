package model

import (
	"encoding/json"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserSchemaDef(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("SchemaDef")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var schemaDef model.SchemaDef
	err = json.Unmarshal([]byte(jsonTemplate), &schemaDef)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// Validate inherited Auditable fields
	if schemaDef.OwnerApp != "sample_ownerApp" {
		t.Errorf("Expected OwnerApp = 'sample_ownerApp', got '%s'", schemaDef.OwnerApp)
	}
	if schemaDef.CreateTime != 123 {
		t.Errorf("Expected CreateTime = 123, got %d", schemaDef.CreateTime)
	}
	if schemaDef.UpdateTime != 123 {
		t.Errorf("Expected UpdateTime = 123, got %d", schemaDef.UpdateTime)
	}
	if schemaDef.CreatedBy != "sample_createdBy" {
		t.Errorf("Expected CreatedBy = 'sample_createdBy', got '%s'", schemaDef.CreatedBy)
	}
	if schemaDef.UpdatedBy != "sample_updatedBy" {
		t.Errorf("Expected UpdatedBy = 'sample_updatedBy', got '%s'", schemaDef.UpdatedBy)
	}

	// Validate SchemaDef-specific fields
	// String fields
	if schemaDef.Name != "sample_name" {
		t.Errorf("Expected Name = 'sample_name', got '%s'", schemaDef.Name)
	}
	if schemaDef.Type != "JSON" {
		t.Errorf("Expected Type = 'JSON', got '%s'", schemaDef.Type)
	}
	if schemaDef.ExternalRef != "sample_externalRef" {
		t.Errorf("Expected ExternalRef = 'sample_externalRef', got '%s'", schemaDef.ExternalRef)
	}

	// Integer field
	if schemaDef.Version != 1 {
		t.Errorf("Expected Version = 1, got %d", schemaDef.Version)
	}

	// Map field
	if schemaDef.Data == nil {
		t.Errorf("Data map should not be nil")
	}
	if len(schemaDef.Data) == 0 {
		t.Errorf("Data map should not be empty")
	}
	if schemaDef.Data["sample_key"] != "sample_value" {
		t.Errorf("Expected Data['sample_key'] = 'sample_value', got %v", schemaDef.Data["sample_key"])
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(schemaDef)
	if err != nil {
		t.Fatalf("Failed to serialize struct: %v", err)
	}

	// 5. Round-trip integrity check
	var roundTripSchemaDef model.SchemaDef
	err = json.Unmarshal(serializedJSON, &roundTripSchemaDef)
	if err != nil {
		t.Fatalf("Failed to deserialize round-trip JSON: %v", err)
	}

	// Compare key fields to ensure data integrity
	if roundTripSchemaDef.Name != schemaDef.Name {
		t.Errorf("Round-trip Name mismatch: %s vs %s", schemaDef.Name, roundTripSchemaDef.Name)
	}
	if roundTripSchemaDef.Type != schemaDef.Type {
		t.Errorf("Round-trip Type mismatch: %s vs %s", schemaDef.Type, roundTripSchemaDef.Type)
	}
	if roundTripSchemaDef.Version != schemaDef.Version {
		t.Errorf("Round-trip Version mismatch: %d vs %d", schemaDef.Version, roundTripSchemaDef.Version)
	}
	if roundTripSchemaDef.OwnerApp != schemaDef.OwnerApp {
		t.Errorf("Round-trip OwnerApp mismatch: %s vs %s", schemaDef.OwnerApp, roundTripSchemaDef.OwnerApp)
	}
	if roundTripSchemaDef.CreateTime != schemaDef.CreateTime {
		t.Errorf("Round-trip CreateTime mismatch: %d vs %d", schemaDef.CreateTime, roundTripSchemaDef.CreateTime)
	}
}
