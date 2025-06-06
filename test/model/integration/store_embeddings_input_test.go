package integration

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model/integration"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserStoreEmbeddingsInput(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("StoreEmbeddingsInput")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var storeEmbeddingsInput integration.StoreEmbeddingsInput
	err = json.Unmarshal([]byte(jsonTemplate), &storeEmbeddingsInput)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// Validate inherited LlmWorkerInput fields
	if storeEmbeddingsInput.LlmProvider != "sample_llmProvider" {
		t.Errorf("Expected LlmProvider = 'sample_llmProvider', got '%s'", storeEmbeddingsInput.LlmProvider)
	}
	if storeEmbeddingsInput.Model != "sample_model" {
		t.Errorf("Expected Model = 'sample_model', got '%s'", storeEmbeddingsInput.Model)
	}
	if storeEmbeddingsInput.EmbeddingModel != "sample_embeddingModel" {
		t.Errorf("Expected EmbeddingModel = 'sample_embeddingModel', got '%s'", storeEmbeddingsInput.EmbeddingModel)
	}
	if storeEmbeddingsInput.EmbeddingModelProvider != "sample_embeddingModelProvider" {
		t.Errorf("Expected EmbeddingModelProvider = 'sample_embeddingModelProvider', got '%s'", storeEmbeddingsInput.EmbeddingModelProvider)
	}
	if storeEmbeddingsInput.Prompt != "sample_prompt" {
		t.Errorf("Expected Prompt = 'sample_prompt', got '%s'", storeEmbeddingsInput.Prompt)
	}
	if storeEmbeddingsInput.Temperature != 123.456 {
		t.Errorf("Expected Temperature = 123.456, got %f", storeEmbeddingsInput.Temperature)
	}
	if storeEmbeddingsInput.TopP != 123.456 {
		t.Errorf("Expected TopP = 123.456, got %f", storeEmbeddingsInput.TopP)
	}
	if storeEmbeddingsInput.MaxTokens != 123 {
		t.Errorf("Expected MaxTokens = 123, got %d", storeEmbeddingsInput.MaxTokens)
	}
	if storeEmbeddingsInput.MaxResults != 123 {
		t.Errorf("Expected MaxResults = 123, got %d", storeEmbeddingsInput.MaxResults)
	}

	// Check inherited StopWords slice
	if storeEmbeddingsInput.StopWords == nil {
		t.Errorf("StopWords slice should not be nil")
	}
	if len(storeEmbeddingsInput.StopWords) == 0 {
		t.Errorf("StopWords slice should not be empty")
	}
	if storeEmbeddingsInput.StopWords[0] != "sample_stopWords" {
		t.Errorf("Expected StopWords[0] = 'sample_stopWords', got '%s'", storeEmbeddingsInput.StopWords[0])
	}

	// Validate StoreEmbeddingsInput-specific fields
	// String fields
	if storeEmbeddingsInput.VectorDB != "sample_vectorDB" {
		t.Errorf("Expected VectorDB = 'sample_vectorDB', got '%s'", storeEmbeddingsInput.VectorDB)
	}
	if storeEmbeddingsInput.Index != "sample_index" {
		t.Errorf("Expected Index = 'sample_index', got '%s'", storeEmbeddingsInput.Index)
	}
	if storeEmbeddingsInput.Namespace != "sample_namespace" {
		t.Errorf("Expected Namespace = 'sample_namespace', got '%s'", storeEmbeddingsInput.Namespace)
	}
	if storeEmbeddingsInput.ID != "sample_id" {
		t.Errorf("Expected ID = 'sample_id', got '%s'", storeEmbeddingsInput.ID)
	}

	// Float32 slice field (embeddings)
	if storeEmbeddingsInput.Embeddings == nil {
		t.Errorf("Embeddings slice should not be nil")
	}
	if len(storeEmbeddingsInput.Embeddings) == 0 {
		t.Errorf("Embeddings slice should not be empty")
	}
	if len(storeEmbeddingsInput.Embeddings) != 1 {
		t.Errorf("Expected Embeddings slice to have 1 element, got %d", len(storeEmbeddingsInput.Embeddings))
	}
	if storeEmbeddingsInput.Embeddings[0] != 3.14 {
		t.Errorf("Expected Embeddings[0] = 3.14, got %f", storeEmbeddingsInput.Embeddings[0])
	}

	// Map field
	if storeEmbeddingsInput.Metadata == nil {
		t.Errorf("Metadata map should not be nil")
	}
	if len(storeEmbeddingsInput.Metadata) == 0 {
		t.Errorf("Metadata map should not be empty")
	}
	if storeEmbeddingsInput.Metadata["sample_key"] != "sample_value" {
		t.Errorf("Expected Metadata['sample_key'] = 'sample_value', got %v", storeEmbeddingsInput.Metadata["sample_key"])
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(storeEmbeddingsInput)
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
