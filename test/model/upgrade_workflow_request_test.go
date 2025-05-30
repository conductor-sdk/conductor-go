package model

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserUpgradeWorkflowRequest(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("UpgradeWorkflowRequest")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var upgradeWorkflowRequest model.UpgradeWorkflowRequest
	err = json.Unmarshal([]byte(jsonTemplate), &upgradeWorkflowRequest)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// String field
	if upgradeWorkflowRequest.Name != "sample_name" {
		t.Errorf("Expected Name = 'sample_name', got '%s'", upgradeWorkflowRequest.Name)
	}

	// Integer field
	if upgradeWorkflowRequest.Version != 123 {
		t.Errorf("Expected Version = 123, got %d", upgradeWorkflowRequest.Version)
	}

	// Map fields
	if upgradeWorkflowRequest.WorkflowInput == nil {
		t.Errorf("WorkflowInput map should not be nil")
	}
	if len(upgradeWorkflowRequest.WorkflowInput) == 0 {
		t.Errorf("WorkflowInput map should not be empty")
	}
	if upgradeWorkflowRequest.WorkflowInput["sample_key"] != "sample_value" {
		t.Errorf("Expected WorkflowInput['sample_key'] = 'sample_value', got %v", upgradeWorkflowRequest.WorkflowInput["sample_key"])
	}

	if upgradeWorkflowRequest.TaskOutput == nil {
		t.Errorf("TaskOutput map should not be nil")
	}
	if len(upgradeWorkflowRequest.TaskOutput) == 0 {
		t.Errorf("TaskOutput map should not be empty")
	}
	if upgradeWorkflowRequest.TaskOutput["sample_key"] != "sample_value" {
		t.Errorf("Expected TaskOutput['sample_key'] = 'sample_value', got %v", upgradeWorkflowRequest.TaskOutput["sample_key"])
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(upgradeWorkflowRequest)
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
