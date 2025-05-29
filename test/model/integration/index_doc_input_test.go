package integration

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model/integration"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserIndexDocInput(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("IndexDocInput")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var indexDocInput integration.IndexDocInput
	err = json.Unmarshal([]byte(jsonTemplate), &indexDocInput)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// String fields
	if indexDocInput.LlmProvider != "sample_llmProvider" {
		t.Errorf("Expected LlmProvider = 'sample_llmProvider', got '%s'", indexDocInput.LlmProvider)
	}
	if indexDocInput.Model != "sample_model" {
		t.Errorf("Expected Model = 'sample_model', got '%s'", indexDocInput.Model)
	}
	if indexDocInput.EmbeddingModelProvider != "sample_embeddingModelProvider" {
		t.Errorf("Expected EmbeddingModelProvider = 'sample_embeddingModelProvider', got '%s'", indexDocInput.EmbeddingModelProvider)
	}
	if indexDocInput.EmbeddingModel != "sample_embeddingModel" {
		t.Errorf("Expected EmbeddingModel = 'sample_embeddingModel', got '%s'", indexDocInput.EmbeddingModel)
	}
	if indexDocInput.VectorDB != "sample_vectorDB" {
		t.Errorf("Expected VectorDB = 'sample_vectorDB', got '%s'", indexDocInput.VectorDB)
	}
	if indexDocInput.Text != "sample_text" {
		t.Errorf("Expected Text = 'sample_text', got '%s'", indexDocInput.Text)
	}
	if indexDocInput.DocId != "sample_docId" {
		t.Errorf("Expected DocId = 'sample_docId', got '%s'", indexDocInput.DocId)
	}
	if indexDocInput.Url != "sample_url" {
		t.Errorf("Expected Url = 'sample_url', got '%s'", indexDocInput.Url)
	}
	if indexDocInput.MediaType != "sample_mediaType" {
		t.Errorf("Expected MediaType = 'sample_mediaType', got '%s'", indexDocInput.MediaType)
	}
	if indexDocInput.Namespace != "sample_namespace" {
		t.Errorf("Expected Namespace = 'sample_namespace', got '%s'", indexDocInput.Namespace)
	}
	if indexDocInput.Index != "sample_index" {
		t.Errorf("Expected Index = 'sample_index', got '%s'", indexDocInput.Index)
	}

	// Integer fields
	if indexDocInput.ChunkSize != 123 {
		t.Errorf("Expected ChunkSize = 123, got %d", indexDocInput.ChunkSize)
	}
	if indexDocInput.ChunkOverlap != 123 {
		t.Errorf("Expected ChunkOverlap = 123, got %d", indexDocInput.ChunkOverlap)
	}

	// Check map field
	if indexDocInput.Metadata == nil {
		t.Errorf("Metadata map should not be nil")
	}
	if len(indexDocInput.Metadata) == 0 {
		t.Errorf("Metadata map should not be empty")
	}
	if indexDocInput.Metadata["sample_key"] != "sample_value" {
		t.Errorf("Expected Metadata['sample_key'] = 'sample_value', got %v", indexDocInput.Metadata["sample_key"])
	}

	// Check pointer to int field
	if indexDocInput.Dimensions == nil {
		t.Errorf("Dimensions should not be nil")
	}
	if *indexDocInput.Dimensions != 123 {
		t.Errorf("Expected Dimensions = 123, got %d", *indexDocInput.Dimensions)
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(indexDocInput)
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
