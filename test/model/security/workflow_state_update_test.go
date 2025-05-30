package security

import (
	"encoding/json"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model/security"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserWorkflowStateUpdate(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("WorkflowStateUpdate")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var workflowStateUpdate security.WorkflowStateUpdate
	err = json.Unmarshal([]byte(jsonTemplate), &workflowStateUpdate)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// String field
	if workflowStateUpdate.TaskReferenceName != "sample_taskReferenceName" {
		t.Errorf("Expected TaskReferenceName = 'sample_taskReferenceName', got '%s'", workflowStateUpdate.TaskReferenceName)
	}

	// Map field
	if workflowStateUpdate.Variables == nil {
		t.Errorf("Variables map should not be nil")
	}
	if len(workflowStateUpdate.Variables) == 0 {
		t.Errorf("Variables map should not be empty")
	}
	if workflowStateUpdate.Variables["sample_key"] != "sample_value" {
		t.Errorf("Expected Variables['sample_key'] = 'sample_value', got %v", workflowStateUpdate.Variables["sample_key"])
	}

	// Interface{} field (TaskResult)
	if workflowStateUpdate.TaskResult == nil {
		t.Errorf("TaskResult should not be nil")
	}

	// Since TaskResult is interface{}, validate based on template structure
	// The template likely resolves to a TaskResult object with various fields
	if taskResultMap, ok := workflowStateUpdate.TaskResult.(map[string]interface{}); ok {
		// Validate some expected fields in the TaskResult
		if taskResultMap["taskId"] == nil {
			t.Errorf("TaskResult should contain taskId field")
		}
		if taskResultMap["status"] == nil {
			t.Errorf("TaskResult should contain status field")
		}

		// Validate specific field values if they exist
		if taskId, exists := taskResultMap["taskId"]; exists {
			if taskIdStr, ok := taskId.(string); ok && taskIdStr != "sample_taskId" {
				t.Errorf("Expected TaskResult.taskId = 'sample_taskId', got '%s'", taskIdStr)
			}
		}
		if status, exists := taskResultMap["status"]; exists {
			if statusStr, ok := status.(string); ok && statusStr != "IN_PROGRESS" {
				t.Errorf("Expected TaskResult.status = 'IN_PROGRESS', got '%s'", statusStr)
			}
		}
	} else {
		t.Logf("TaskResult contains: %v (type: %T)", workflowStateUpdate.TaskResult, workflowStateUpdate.TaskResult)
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(workflowStateUpdate)
	if err != nil {
		t.Fatalf("Failed to serialize struct: %v", err)
	}

	// 5. Round-trip integrity check
	var roundTripUpdate security.WorkflowStateUpdate
	err = json.Unmarshal(serializedJSON, &roundTripUpdate)
	if err != nil {
		t.Fatalf("Failed to deserialize round-trip JSON: %v", err)
	}

	// Compare key fields to ensure data integrity
	if roundTripUpdate.TaskReferenceName != workflowStateUpdate.TaskReferenceName {
		t.Errorf("Round-trip TaskReferenceName mismatch: %s vs %s",
			workflowStateUpdate.TaskReferenceName, roundTripUpdate.TaskReferenceName)
	}

	if len(roundTripUpdate.Variables) != len(workflowStateUpdate.Variables) {
		t.Errorf("Round-trip Variables length mismatch: %d vs %d",
			len(workflowStateUpdate.Variables), len(roundTripUpdate.Variables))
	}

	if (workflowStateUpdate.TaskResult == nil) != (roundTripUpdate.TaskResult == nil) {
		t.Errorf("Round-trip TaskResult nil mismatch")
	}
}
