package integration

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model/integration"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserEmbeddingRequest(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("EmbeddingRequest")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var embeddingRequest integration.EmbeddingRequest
	err = json.Unmarshal([]byte(jsonTemplate), &embeddingRequest)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	if embeddingRequest.LlmProvider != "sample_llmProvider" {
		t.Errorf("Expected LlmProvider = 'sample_llmProvider', got '%s'", embeddingRequest.LlmProvider)
	}
	if embeddingRequest.Model != "sample_model" {
		t.Errorf("Expected Model = 'sample_model', got '%s'", embeddingRequest.Model)
	}
	if embeddingRequest.Text != "sample_text" {
		t.Errorf("Expected Text = 'sample_text', got '%s'", embeddingRequest.Text)
	}
	if embeddingRequest.Dimensions != 123 {
		t.Errorf("Expected Dimensions = 123, got %d", embeddingRequest.Dimensions)
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(embeddingRequest)
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
