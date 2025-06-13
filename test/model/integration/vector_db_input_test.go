package integration

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model/integration"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserVectorDbInput(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("VectorDBInput")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var vectorDbInput integration.VectorDbInput
	err = json.Unmarshal([]byte(jsonTemplate), &vectorDbInput)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// Validate inherited LLMWorkerInput fields
	if vectorDbInput.LlmProvider != "sample_llmProvider" {
		t.Errorf("Expected LlmProvider = 'sample_llmProvider', got '%s'", vectorDbInput.LlmProvider)
	}
	if vectorDbInput.Model != "sample_model" {
		t.Errorf("Expected Model = 'sample_model', got '%s'", vectorDbInput.Model)
	}
	if vectorDbInput.EmbeddingModel != "sample_embeddingModel" {
		t.Errorf("Expected EmbeddingModel = 'sample_embeddingModel', got '%s'", vectorDbInput.EmbeddingModel)
	}
	if vectorDbInput.EmbeddingModelProvider != "sample_embeddingModelProvider" {
		t.Errorf("Expected EmbeddingModelProvider = 'sample_embeddingModelProvider', got '%s'", vectorDbInput.EmbeddingModelProvider)
	}
	if vectorDbInput.Prompt != "sample_prompt" {
		t.Errorf("Expected Prompt = 'sample_prompt', got '%s'", vectorDbInput.Prompt)
	}
	if vectorDbInput.Temperature != 123.456 {
		t.Errorf("Expected Temperature = 123.456, got %f", vectorDbInput.Temperature)
	}
	if vectorDbInput.TopP != 123.456 {
		t.Errorf("Expected TopP = 123.456, got %f", vectorDbInput.TopP)
	}
	if vectorDbInput.MaxTokens != 123 {
		t.Errorf("Expected MaxTokens = 123, got %d", vectorDbInput.MaxTokens)
	}
	if vectorDbInput.MaxResults != 123 {
		t.Errorf("Expected MaxResults = 123, got %d", vectorDbInput.MaxResults)
	}

	// Check inherited StopWords slice
	if vectorDbInput.StopWords == nil {
		t.Errorf("StopWords slice should not be nil")
	}
	if len(vectorDbInput.StopWords) == 0 {
		t.Errorf("StopWords slice should not be empty")
	}
	if vectorDbInput.StopWords[0] != "sample_stopWords" {
		t.Errorf("Expected StopWords[0] = 'sample_stopWords', got '%s'", vectorDbInput.StopWords[0])
	}

	// Validate VectorDbInput-specific fields
	// String fields
	if vectorDbInput.VectorDB != "sample_vectorDB" {
		t.Errorf("Expected VectorDB = 'sample_vectorDB', got '%s'", vectorDbInput.VectorDB)
	}
	if vectorDbInput.Index != "sample_index" {
		t.Errorf("Expected Index = 'sample_index', got '%s'", vectorDbInput.Index)
	}
	if vectorDbInput.Namespace != "sample_namespace" {
		t.Errorf("Expected Namespace = 'sample_namespace', got '%s'", vectorDbInput.Namespace)
	}
	if vectorDbInput.Query != "sample_query" {
		t.Errorf("Expected Query = 'sample_query', got '%s'", vectorDbInput.Query)
	}

	// Integer field
	if vectorDbInput.Dimensions != 123 {
		t.Errorf("Expected Dimensions = 123, got %d", vectorDbInput.Dimensions)
	}

	// Float32 slice field (embeddings)
	if vectorDbInput.Embeddings == nil {
		t.Errorf("Embeddings slice should not be nil")
	}
	if len(vectorDbInput.Embeddings) == 0 {
		t.Errorf("Embeddings slice should not be empty")
	}
	if len(vectorDbInput.Embeddings) != 1 {
		t.Errorf("Expected Embeddings slice to have 1 element, got %d", len(vectorDbInput.Embeddings))
	}
	if vectorDbInput.Embeddings[0] != 3.14 {
		t.Errorf("Expected Embeddings[0] = 3.14, got %f", vectorDbInput.Embeddings[0])
	}

	// Map field
	if vectorDbInput.Metadata == nil {
		t.Errorf("Metadata map should not be nil")
	}
	if len(vectorDbInput.Metadata) == 0 {
		t.Errorf("Metadata map should not be empty")
	}
	if vectorDbInput.Metadata["sample_key"] != "sample_value" {
		t.Errorf("Expected Metadata['sample_key'] = 'sample_value', got %v", vectorDbInput.Metadata["sample_key"])
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(vectorDbInput)
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
