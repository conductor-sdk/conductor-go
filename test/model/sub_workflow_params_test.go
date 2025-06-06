package model

import (
	"encoding/json"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserSubWorkflowParams(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("SubWorkflowParams")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var subWorkflowParams model.SubWorkflowParams
	err = json.Unmarshal([]byte(jsonTemplate), &subWorkflowParams)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// String fields
	if subWorkflowParams.Name != "sample_name" {
		t.Errorf("Expected Name = 'sample_name', got '%s'", subWorkflowParams.Name)
	}
	if subWorkflowParams.IdempotencyKey != "sample_idempotencyKey" {
		t.Errorf("Expected IdempotencyKey = 'sample_idempotencyKey', got '%s'", subWorkflowParams.IdempotencyKey)
	}
	if subWorkflowParams.IdempotencyStrategy != "FAIL" {
		t.Errorf("Expected IdempotencyStrategy = 'FAIL', got '%s'", subWorkflowParams.IdempotencyStrategy)
	}

	// Int32 field
	if subWorkflowParams.Version != 123 {
		t.Errorf("Expected Version = 123, got %d", subWorkflowParams.Version)
	}

	// Map field
	if subWorkflowParams.TaskToDomain == nil {
		t.Errorf("TaskToDomain map should not be nil")
	}
	if len(subWorkflowParams.TaskToDomain) == 0 {
		t.Errorf("TaskToDomain map should not be empty")
	}
	if subWorkflowParams.TaskToDomain["sample_key"] != "sample_value" {
		t.Errorf("Expected TaskToDomain['sample_key'] = 'sample_value', got %v", subWorkflowParams.TaskToDomain["sample_key"])
	}

	// Interface{} field - WorkflowDefinition
	if subWorkflowParams.WorkflowDefinition == nil {
		t.Errorf("WorkflowDefinition should not be nil")
	}

	// Validate WorkflowDefinition structure (should resolve to WorkflowDef object)
	if subWorkflowParams.WorkflowDefinition == nil {
		t.Errorf("WorkflowDefinition should not be nil")
	} else {
		// Validate WorkflowDef fields
		if subWorkflowParams.WorkflowDefinition.Name != "sample_name" {
			t.Errorf("Expected WorkflowDefinition.Name = 'sample_name', got '%s'", subWorkflowParams.WorkflowDefinition.Name)
		}
		if subWorkflowParams.WorkflowDefinition.Version != 123 {
			t.Errorf("Expected WorkflowDefinition.Version = 123, got %d", subWorkflowParams.WorkflowDefinition.Version)
		}
	}

	// Interface{} field - Priority
	if subWorkflowParams.Priority == nil {
		t.Errorf("Priority should not be nil")
	}

	// Validate Priority (could be string or number based on template)
	if priorityStr, ok := subWorkflowParams.Priority.(string); ok {
		if priorityStr != "sample_object_priority" {
			t.Errorf("Expected Priority = 'sample_object_priority', got '%s'", priorityStr)
		}
	} else if priorityFloat, ok := subWorkflowParams.Priority.(float64); ok {
		t.Logf("Priority is numeric: %v", priorityFloat)
	} else {
		t.Logf("Priority contains: %v (type: %T)", subWorkflowParams.Priority, subWorkflowParams.Priority)
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(subWorkflowParams)
	if err != nil {
		t.Fatalf("Failed to serialize struct: %v", err)
	}

	// 5. Round-trip integrity check
	var roundTripParams model.SubWorkflowParams
	err = json.Unmarshal(serializedJSON, &roundTripParams)
	if err != nil {
		t.Fatalf("Failed to deserialize round-trip JSON: %v", err)
	}

	// Compare key fields to ensure data integrity
	if roundTripParams.Name != subWorkflowParams.Name {
		t.Errorf("Round-trip Name mismatch: %s vs %s", subWorkflowParams.Name, roundTripParams.Name)
	}
	if roundTripParams.Version != subWorkflowParams.Version {
		t.Errorf("Round-trip Version mismatch: %d vs %d", subWorkflowParams.Version, roundTripParams.Version)
	}
	if roundTripParams.IdempotencyKey != subWorkflowParams.IdempotencyKey {
		t.Errorf("Round-trip IdempotencyKey mismatch: %s vs %s", subWorkflowParams.IdempotencyKey, roundTripParams.IdempotencyKey)
	}
	if roundTripParams.IdempotencyStrategy != subWorkflowParams.IdempotencyStrategy {
		t.Errorf("Round-trip IdempotencyStrategy mismatch: %s vs %s", subWorkflowParams.IdempotencyStrategy, roundTripParams.IdempotencyStrategy)
	}

	// Compare interface{} fields
	if (subWorkflowParams.WorkflowDefinition == nil) != (roundTripParams.WorkflowDefinition == nil) {
		t.Errorf("Round-trip WorkflowDefinition nil mismatch")
	}
	if (subWorkflowParams.Priority == nil) != (roundTripParams.Priority == nil) {
		t.Errorf("Round-trip Priority nil mismatch")
	}
}
