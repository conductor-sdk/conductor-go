package integration

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model/integration"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserLlmResponse(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("LLMResponse")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var llmResponse integration.LlmResponse
	err = json.Unmarshal([]byte(jsonTemplate), &llmResponse)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// String field
	if llmResponse.FinishReason != "sample_finishReason" {
		t.Errorf("Expected FinishReason = 'sample_finishReason', got '%s'", llmResponse.FinishReason)
	}

	// Integer field
	if llmResponse.TokenUsed != 123 {
		t.Errorf("Expected TokenUsed = 123, got %d", llmResponse.TokenUsed)
	}

	// Interface{} field - validate it's not nil and has expected content
	if llmResponse.Result == nil {
		t.Errorf("Result should not be nil")
	}

	// Since Result is interface{}, check if it contains expected sample value
	if resultStr, ok := llmResponse.Result.(string); ok {
		if resultStr != "sample_result" {
			t.Errorf("Expected Result = 'sample_result', got '%s'", resultStr)
		}
	} else {
		// If not a string, just validate it's populated
		t.Logf("Result field contains: %v (type: %T)", llmResponse.Result, llmResponse.Result)
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(llmResponse)
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
