package integration

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model/integration"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserChatCompletion(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("ChatCompletion")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var chatCompletion integration.ChatCompletion
	err = json.Unmarshal([]byte(jsonTemplate), &chatCompletion)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// Validate inherited LlmWorkerInput fields
	if chatCompletion.LlmProvider != "sample_llmProvider" {
		t.Errorf("Expected LlmProvider = 'sample_llmProvider', got '%s'", chatCompletion.LlmProvider)
	}
	if chatCompletion.Model != "sample_model" {
		t.Errorf("Expected Model = 'sample_model', got '%s'", chatCompletion.Model)
	}
	if chatCompletion.EmbeddingModel != "sample_embeddingModel" {
		t.Errorf("Expected EmbeddingModel = 'sample_embeddingModel', got '%s'", chatCompletion.EmbeddingModel)
	}
	if chatCompletion.EmbeddingModelProvider != "sample_embeddingModelProvider" {
		t.Errorf("Expected EmbeddingModelProvider = 'sample_embeddingModelProvider', got '%s'", chatCompletion.EmbeddingModelProvider)
	}
	if chatCompletion.Prompt != "sample_prompt" {
		t.Errorf("Expected Prompt = 'sample_prompt', got '%s'", chatCompletion.Prompt)
	}
	if chatCompletion.Temperature != 123.456 {
		t.Errorf("Expected Temperature = 123.456, got %f", chatCompletion.Temperature)
	}
	if chatCompletion.TopP != 123.456 {
		t.Errorf("Expected TopP = 123.456, got %f", chatCompletion.TopP)
	}
	if chatCompletion.MaxTokens != 123 {
		t.Errorf("Expected MaxTokens = 123, got %d", chatCompletion.MaxTokens)
	}
	if chatCompletion.MaxResults != 123 {
		t.Errorf("Expected MaxResults = 123, got %d", chatCompletion.MaxResults)
	}

	// Check slice field - StopWords
	if chatCompletion.StopWords == nil {
		t.Errorf("StopWords slice should not be nil")
	}
	if len(chatCompletion.StopWords) == 0 {
		t.Errorf("StopWords slice should not be empty")
	}
	if chatCompletion.StopWords[0] != "sample_stopWords" {
		t.Errorf("Expected StopWords[0] = 'sample_stopWords', got '%s'", chatCompletion.StopWords[0])
	}

	// Validate ChatCompletion-specific fields
	if chatCompletion.Instructions != "sample_instructions" {
		t.Errorf("Expected Instructions = 'sample_instructions', got '%s'", chatCompletion.Instructions)
	}
	if !chatCompletion.JsonOutput {
		t.Errorf("Expected JsonOutput = true, got %v", chatCompletion.JsonOutput)
	}

	// Check Messages slice
	if chatCompletion.Messages == nil {
		t.Errorf("Messages slice should not be nil")
	}
	if len(chatCompletion.Messages) == 0 {
		t.Errorf("Messages slice should not be empty")
	}
	if len(chatCompletion.Messages) != 1 {
		t.Errorf("Expected Messages slice to have 1 element, got %d", len(chatCompletion.Messages))
	}

	// Validate ChatMessage element
	message := chatCompletion.Messages[0]
	if message.Role != "sample_role" {
		t.Errorf("Expected Message.Role = 'sample_role', got '%s'", message.Role)
	}
	if message.Message != "sample_message" {
		t.Errorf("Expected Message.Message = 'sample_message', got '%s'", message.Message)
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(chatCompletion)
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

func TestSerDserChatMessage(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("ChatMessage")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var chatMessage integration.ChatMessage
	err = json.Unmarshal([]byte(jsonTemplate), &chatMessage)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	if chatMessage.Role != "sample_role" {
		t.Errorf("Expected Role = 'sample_role', got '%s'", chatMessage.Role)
	}
	if chatMessage.Message != "sample_message" {
		t.Errorf("Expected Message = 'sample_message', got '%s'", chatMessage.Message)
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(chatMessage)
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
