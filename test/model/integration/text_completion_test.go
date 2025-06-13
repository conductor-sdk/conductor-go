package integration

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model/integration"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserTextCompletion(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("TextCompletion")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var textCompletion integration.TextCompletion
	err = json.Unmarshal([]byte(jsonTemplate), &textCompletion)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// Validate all inherited LLMWorkerInput fields
	if textCompletion.LlmProvider != "sample_llmProvider" {
		t.Errorf("Expected LlmProvider = 'sample_llmProvider', got '%s'", textCompletion.LlmProvider)
	}
	if textCompletion.Model != "sample_model" {
		t.Errorf("Expected Model = 'sample_model', got '%s'", textCompletion.Model)
	}
	if textCompletion.EmbeddingModel != "sample_embeddingModel" {
		t.Errorf("Expected EmbeddingModel = 'sample_embeddingModel', got '%s'", textCompletion.EmbeddingModel)
	}
	if textCompletion.EmbeddingModelProvider != "sample_embeddingModelProvider" {
		t.Errorf("Expected EmbeddingModelProvider = 'sample_embeddingModelProvider', got '%s'", textCompletion.EmbeddingModelProvider)
	}
	if textCompletion.Prompt != "sample_prompt" {
		t.Errorf("Expected Prompt = 'sample_prompt', got '%s'", textCompletion.Prompt)
	}

	// Float64 fields
	if textCompletion.Temperature != 123.456 {
		t.Errorf("Expected Temperature = 123.456, got %f", textCompletion.Temperature)
	}
	if textCompletion.TopP != 123.456 {
		t.Errorf("Expected TopP = 123.456, got %f", textCompletion.TopP)
	}

	// Integer fields
	if textCompletion.MaxTokens != 123 {
		t.Errorf("Expected MaxTokens = 123, got %d", textCompletion.MaxTokens)
	}
	if textCompletion.MaxResults != 123 {
		t.Errorf("Expected MaxResults = 123, got %d", textCompletion.MaxResults)
	}

	// Slice field
	if textCompletion.StopWords == nil {
		t.Errorf("StopWords slice should not be nil")
	}
	if len(textCompletion.StopWords) == 0 {
		t.Errorf("StopWords slice should not be empty")
	}
	if len(textCompletion.StopWords) != 1 {
		t.Errorf("Expected StopWords slice to have 1 element, got %d", len(textCompletion.StopWords))
	}
	if textCompletion.StopWords[0] != "sample_stopWords" {
		t.Errorf("Expected StopWords[0] = 'sample_stopWords', got '%s'", textCompletion.StopWords[0])
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(textCompletion)
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
