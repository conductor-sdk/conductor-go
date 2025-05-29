package model

import (
	"encoding/json"
	_ "reflect"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserStartWorkflowRequest(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("StartWorkflowRequest")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var startWorkflowRequest model.StartWorkflowRequest
	err = json.Unmarshal([]byte(jsonTemplate), &startWorkflowRequest)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// String fields
	if startWorkflowRequest.Name != "sample_name" {
		t.Errorf("Expected Name = 'sample_name', got '%s'", startWorkflowRequest.Name)
	}
	if startWorkflowRequest.CorrelationId != "sample_correlationId" {
		t.Errorf("Expected CorrelationId = 'sample_correlationId', got '%s'", startWorkflowRequest.CorrelationId)
	}
	if startWorkflowRequest.ExternalInputPayloadStoragePath != "sample_externalInputPayloadStoragePath" {
		t.Errorf("Expected ExternalInputPayloadStoragePath = 'sample_externalInputPayloadStoragePath', got '%s'", startWorkflowRequest.ExternalInputPayloadStoragePath)
	}
	if startWorkflowRequest.CreatedBy != "sample_createdBy" {
		t.Errorf("Expected CreatedBy = 'sample_createdBy', got '%s'", startWorkflowRequest.CreatedBy)
	}
	if startWorkflowRequest.IdempotencyKey != "sample_idempotencyKey" {
		t.Errorf("Expected IdempotencyKey = 'sample_idempotencyKey', got '%s'", startWorkflowRequest.IdempotencyKey)
	}

	// Int32 fields
	if startWorkflowRequest.Version != 123 {
		t.Errorf("Expected Version = 123, got %d", startWorkflowRequest.Version)
	}
	if startWorkflowRequest.Priority != 123 {
		t.Errorf("Expected Priority = 123, got %d", startWorkflowRequest.Priority)
	}

	// Enum field
	if startWorkflowRequest.IdempotencyStrategy != model.FailOnConflict {
		t.Errorf("Expected IdempotencyStrategy = 'FAIL', got '%s'", startWorkflowRequest.IdempotencyStrategy)
	}

	// Map fields
	if startWorkflowRequest.Input == nil {
		t.Errorf("Input map should not be nil")
	}
	if len(startWorkflowRequest.Input) == 0 {
		t.Errorf("Input map should not be empty")
	}
	if startWorkflowRequest.Input["sample_key"] != "sample_value" {
		t.Errorf("Expected Input['sample_key'] = 'sample_value', got %v", startWorkflowRequest.Input["sample_key"])
	}

	if startWorkflowRequest.TaskToDomain == nil {
		t.Errorf("TaskToDomain map should not be nil")
	}
	if len(startWorkflowRequest.TaskToDomain) == 0 {
		t.Errorf("TaskToDomain map should not be empty")
	}
	if startWorkflowRequest.TaskToDomain["sample_key"] != "sample_value" {
		t.Errorf("Expected TaskToDomain['sample_key'] = 'sample_value', got %v", startWorkflowRequest.TaskToDomain["sample_key"])
	}

	// Pointer to struct field
	if startWorkflowRequest.WorkflowDef == nil {
		t.Errorf("WorkflowDef should not be nil")
	} else {
		// Validate nested WorkflowDef fields
		if startWorkflowRequest.WorkflowDef.Name != "sample_name" {
			t.Errorf("Expected WorkflowDef.Name = 'sample_name', got '%s'", startWorkflowRequest.WorkflowDef.Name)
		}
		if startWorkflowRequest.WorkflowDef.Version != 123 {
			t.Errorf("Expected WorkflowDef.Version = 123, got %d", startWorkflowRequest.WorkflowDef.Version)
		}
		if startWorkflowRequest.WorkflowDef.Description != "sample_description" {
			t.Errorf("Expected WorkflowDef.Description = 'sample_description', got '%s'", startWorkflowRequest.WorkflowDef.Description)
		}

		// Check WorkflowDef maps and slices
		if startWorkflowRequest.WorkflowDef.Tasks == nil || len(startWorkflowRequest.WorkflowDef.Tasks) == 0 {
			t.Errorf("WorkflowDef.Tasks should not be nil or empty")
		}
		if startWorkflowRequest.WorkflowDef.OutputParameters == nil || len(startWorkflowRequest.WorkflowDef.OutputParameters) == 0 {
			t.Errorf("WorkflowDef.OutputParameters should not be nil or empty")
		}
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(startWorkflowRequest)
	if err != nil {
		t.Fatalf("Failed to serialize struct: %v", err)
	}

	// 5. Round-trip integrity check
	var roundTripRequest model.StartWorkflowRequest
	err = json.Unmarshal(serializedJSON, &roundTripRequest)
	if err != nil {
		t.Fatalf("Failed to deserialize round-trip JSON: %v", err)
	}

	// Compare key fields to ensure data integrity
	if roundTripRequest.Name != startWorkflowRequest.Name {
		t.Errorf("Round-trip Name mismatch: %s vs %s", startWorkflowRequest.Name, roundTripRequest.Name)
	}
	if roundTripRequest.CorrelationId != startWorkflowRequest.CorrelationId {
		t.Errorf("Round-trip CorrelationId mismatch: %s vs %s", startWorkflowRequest.CorrelationId, roundTripRequest.CorrelationId)
	}
	if roundTripRequest.Version != startWorkflowRequest.Version {
		t.Errorf("Round-trip Version mismatch: %d vs %d", startWorkflowRequest.Version, roundTripRequest.Version)
	}
	if roundTripRequest.Priority != startWorkflowRequest.Priority {
		t.Errorf("Round-trip Priority mismatch: %d vs %d", startWorkflowRequest.Priority, roundTripRequest.Priority)
	}
	if roundTripRequest.IdempotencyStrategy != startWorkflowRequest.IdempotencyStrategy {
		t.Errorf("Round-trip IdempotencyStrategy mismatch: %s vs %s", startWorkflowRequest.IdempotencyStrategy, roundTripRequest.IdempotencyStrategy)
	}

	// Compare nested WorkflowDef
	if (startWorkflowRequest.WorkflowDef == nil) != (roundTripRequest.WorkflowDef == nil) {
		t.Errorf("Round-trip WorkflowDef nil mismatch")
	} else if startWorkflowRequest.WorkflowDef != nil {
		if roundTripRequest.WorkflowDef.Name != startWorkflowRequest.WorkflowDef.Name {
			t.Errorf("Round-trip WorkflowDef.Name mismatch: %s vs %s",
				startWorkflowRequest.WorkflowDef.Name, roundTripRequest.WorkflowDef.Name)
		}
	}
}
