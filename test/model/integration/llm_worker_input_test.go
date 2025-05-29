package integration

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model/integration"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserLlmWorkerInput(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("LLMWorkerInput")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var llmWorkerInput integration.LlmWorkerInput
	err = json.Unmarshal([]byte(jsonTemplate), &llmWorkerInput)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// String fields
	if llmWorkerInput.LlmProvider != "sample_llmProvider" {
		t.Errorf("Expected LlmProvider = 'sample_llmProvider', got '%s'", llmWorkerInput.LlmProvider)
	}
	if llmWorkerInput.Model != "sample_model" {
		t.Errorf("Expected Model = 'sample_model', got '%s'", llmWorkerInput.Model)
	}
	if llmWorkerInput.EmbeddingModel != "sample_embeddingModel" {
		t.Errorf("Expected EmbeddingModel = 'sample_embeddingModel', got '%s'", llmWorkerInput.EmbeddingModel)
	}
	if llmWorkerInput.EmbeddingModelProvider != "sample_embeddingModelProvider" {
		t.Errorf("Expected EmbeddingModelProvider = 'sample_embeddingModelProvider', got '%s'", llmWorkerInput.EmbeddingModelProvider)
	}
	if llmWorkerInput.Prompt != "sample_prompt" {
		t.Errorf("Expected Prompt = 'sample_prompt', got '%s'", llmWorkerInput.Prompt)
	}

	// Float64 fields
	if llmWorkerInput.Temperature != 123.456 {
		t.Errorf("Expected Temperature = 123.456, got %f", llmWorkerInput.Temperature)
	}
	if llmWorkerInput.TopP != 123.456 {
		t.Errorf("Expected TopP = 123.456, got %f", llmWorkerInput.TopP)
	}

	// Integer fields
	if llmWorkerInput.MaxTokens != 123 {
		t.Errorf("Expected MaxTokens = 123, got %d", llmWorkerInput.MaxTokens)
	}
	if llmWorkerInput.MaxResults != 123 {
		t.Errorf("Expected MaxResults = 123, got %d", llmWorkerInput.MaxResults)
	}

	// Slice field
	if llmWorkerInput.StopWords == nil {
		t.Errorf("StopWords slice should not be nil")
	}
	if len(llmWorkerInput.StopWords) == 0 {
		t.Errorf("StopWords slice should not be empty")
	}
	if len(llmWorkerInput.StopWords) != 1 {
		t.Errorf("Expected StopWords slice to have 1 element, got %d", len(llmWorkerInput.StopWords))
	}
	if llmWorkerInput.StopWords[0] != "sample_stopWords" {
		t.Errorf("Expected StopWords[0] = 'sample_stopWords', got '%s'", llmWorkerInput.StopWords[0])
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(llmWorkerInput)
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
