package security

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model/security"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserCorrelationIdsSearchRequest(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("CorrelationIdsSearchRequest")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var correlationIdsSearchRequest security.CorrelationIdsSearchRequest
	err = json.Unmarshal([]byte(jsonTemplate), &correlationIdsSearchRequest)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// Check CorrelationIds slice
	if correlationIdsSearchRequest.CorrelationIds == nil {
		t.Errorf("CorrelationIds slice should not be nil")
	}
	if len(correlationIdsSearchRequest.CorrelationIds) == 0 {
		t.Errorf("CorrelationIds slice should not be empty")
	}
	if len(correlationIdsSearchRequest.CorrelationIds) != 1 {
		t.Errorf("Expected CorrelationIds slice to have 1 element, got %d", len(correlationIdsSearchRequest.CorrelationIds))
	}
	if correlationIdsSearchRequest.CorrelationIds[0] != "sample_correlationIds" {
		t.Errorf("Expected CorrelationIds[0] = 'sample_correlationIds', got '%s'", correlationIdsSearchRequest.CorrelationIds[0])
	}

	// Check WorkflowNames slice
	if correlationIdsSearchRequest.WorkflowNames == nil {
		t.Errorf("WorkflowNames slice should not be nil")
	}
	if len(correlationIdsSearchRequest.WorkflowNames) == 0 {
		t.Errorf("WorkflowNames slice should not be empty")
	}
	if len(correlationIdsSearchRequest.WorkflowNames) != 1 {
		t.Errorf("Expected WorkflowNames slice to have 1 element, got %d", len(correlationIdsSearchRequest.WorkflowNames))
	}
	if correlationIdsSearchRequest.WorkflowNames[0] != "sample_workflowNames" {
		t.Errorf("Expected WorkflowNames[0] = 'sample_workflowNames', got '%s'", correlationIdsSearchRequest.WorkflowNames[0])
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(correlationIdsSearchRequest)
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
