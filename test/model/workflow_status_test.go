package model_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserWorkflowStatus(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("WorkflowStatus")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var workflowStatus model.WorkflowStatus
	err = json.Unmarshal([]byte(jsonTemplate), &workflowStatus)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	if workflowStatus.WorkflowId != "sample_workflowId" {
		t.Errorf("Expected WorkflowId = 'sample_workflowId', got '%s'", workflowStatus.WorkflowId)
	}
	if workflowStatus.CorrelationId != "sample_correlationId" {
		t.Errorf("Expected CorrelationId = 'sample_correlationId', got '%s'", workflowStatus.CorrelationId)
	}
	if workflowStatus.Output == nil {
		t.Errorf("Output should not be nil")
	}
	if len(workflowStatus.Output) == 0 {
		t.Errorf("Output should not be empty")
	}
	if workflowStatus.Output != nil && workflowStatus.Output["key"] != "sample_value" {
		t.Errorf("Expected Output['sample_key'] = 'sample_value', got %v", workflowStatus.Output["sample_key"])
	}
	if workflowStatus.Variables == nil {
		t.Errorf("Variables should not be nil")
	}
	if len(workflowStatus.Variables) == 0 {
		t.Errorf("Variables should not be empty")
	}
	if workflowStatus.Variables != nil && workflowStatus.Variables["key"] != "sample_value" {
		t.Errorf("Expected Variables['sample_key'] = 'sample_value', got %v", workflowStatus.Variables["sample_key"])
	}
	if workflowStatus.Status != model.RunningWorkflow {
		t.Errorf("Expected Status = 'RUNNING', got '%s'", workflowStatus.Status)
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(workflowStatus)
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
