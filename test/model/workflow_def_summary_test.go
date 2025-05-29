package model

import (
	"encoding/json"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"testing"

	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserWorkflowDefSummary(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("WorkflowDefSummary")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var workflowDefSummary model.WorkflowDefSummary
	err = json.Unmarshal([]byte(jsonTemplate), &workflowDefSummary)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	if workflowDefSummary.Name != "sample_name" {
		t.Errorf("Expected Name = 'sample_name', got '%s'", workflowDefSummary.Name)
	}
	if workflowDefSummary.Version != 123 {
		t.Errorf("Expected Version = 123, got %d", workflowDefSummary.Version)
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(workflowDefSummary)
	if err != nil {
		t.Fatalf("Failed to serialize struct: %v", err)
	}

	// 5. Round-trip integrity check
	var originalMap, serializedMap map[string]interface{}
	json.Unmarshal([]byte(jsonTemplate), &originalMap)
	json.Unmarshal(serializedJSON, &serializedMap)

	// Compare key fields
	if originalMap["name"] != serializedMap["name"] {
		t.Errorf("Round-trip name mismatch: %v vs %v", originalMap["name"], serializedMap["name"])
	}
	if originalMap["version"] != serializedMap["version"] {
		t.Errorf("Round-trip version mismatch: %v vs %v", originalMap["version"], serializedMap["version"])
	}
	if originalMap["createTime"] != serializedMap["createTime"] {
		t.Errorf("Round-trip createTime mismatch: %v vs %v", originalMap["createTime"], serializedMap["createTime"])
	}
}
