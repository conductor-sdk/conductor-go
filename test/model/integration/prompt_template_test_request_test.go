package integration

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model/integration"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserPromptTemplateTestRequest(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("PromptTemplateTestRequest")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var promptTemplateTestRequest integration.PromptTemplateTestRequest
	err = json.Unmarshal([]byte(jsonTemplate), &promptTemplateTestRequest)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// String fields
	if promptTemplateTestRequest.LlmProvider != "sample_llmProvider" {
		t.Errorf("Expected LlmProvider = 'sample_llmProvider', got '%s'", promptTemplateTestRequest.LlmProvider)
	}
	if promptTemplateTestRequest.Model != "sample_model" {
		t.Errorf("Expected Model = 'sample_model', got '%s'", promptTemplateTestRequest.Model)
	}
	if promptTemplateTestRequest.Prompt != "sample_prompt" {
		t.Errorf("Expected Prompt = 'sample_prompt', got '%s'", promptTemplateTestRequest.Prompt)
	}

	// Float64 fields
	if promptTemplateTestRequest.Temperature != 123.456 {
		t.Errorf("Expected Temperature = 123.456, got %f", promptTemplateTestRequest.Temperature)
	}
	if promptTemplateTestRequest.TopP != 123.456 {
		t.Errorf("Expected TopP = 123.456, got %f", promptTemplateTestRequest.TopP)
	}

	// Map field
	if promptTemplateTestRequest.PromptVariables == nil {
		t.Errorf("PromptVariables map should not be nil")
	}
	if len(promptTemplateTestRequest.PromptVariables) == 0 {
		t.Errorf("PromptVariables map should not be empty")
	}
	if promptTemplateTestRequest.PromptVariables["key"] != "sample_value" {
		t.Errorf("Expected PromptVariables['key'] = 'sample_value', got %v", promptTemplateTestRequest.PromptVariables["key"])
	}

	// Slice field
	if promptTemplateTestRequest.StopWords == nil {
		t.Errorf("StopWords slice should not be nil")
	}
	if len(promptTemplateTestRequest.StopWords) == 0 {
		t.Errorf("StopWords slice should not be empty")
	}
	if len(promptTemplateTestRequest.StopWords) != 1 {
		t.Errorf("Expected StopWords slice to have 1 element, got %d", len(promptTemplateTestRequest.StopWords))
	}
	if promptTemplateTestRequest.StopWords[0] != "sample_stopWords" {
		t.Errorf("Expected StopWords[0] = 'sample_stopWords', got '%s'", promptTemplateTestRequest.StopWords[0])
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(promptTemplateTestRequest)
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
