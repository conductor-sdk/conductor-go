package model

import (
	"encoding/json"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
	"reflect"
	"testing"
)

func TestSerDserEventExecution(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("EventExecution")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var eventExecution model.EventExecution
	err = json.Unmarshal([]byte(jsonTemplate), &eventExecution)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	if eventExecution.ID != "sample_id" {
		t.Errorf("Expected ID = 'sample_id', got '%s'", eventExecution.ID)
	}
	if eventExecution.MessageID != "sample_messageId" {
		t.Errorf("Expected MessageID = 'sample_messageId', got '%s'", eventExecution.MessageID)
	}
	if eventExecution.Name != "sample_name" {
		t.Errorf("Expected Name = 'sample_name', got '%s'", eventExecution.Name)
	}
	if eventExecution.Event != "sample_event" {
		t.Errorf("Expected Event = 'sample_event', got '%s'", eventExecution.Event)
	}
	if eventExecution.Created != 123 {
		t.Errorf("Expected Created = 123, got %d", eventExecution.Created)
	}

	// Check enum field (Status should be resolved from template reference)
	if eventExecution.Status != "IN_PROGRESS" {
		t.Errorf("Expected Status = 'IN_PROGRESS', got '%s'", eventExecution.Status)
	}

	// Check enum field (Action should be resolved from template reference)
	if eventExecution.Action != "start_workflow" {
		t.Errorf("Expected Action = 'start_workflow', got '%s'", eventExecution.Action)
	}

	// Check map field
	if eventExecution.Output == nil {
		t.Errorf("Output map should not be nil")
	}
	if len(eventExecution.Output) == 0 {
		t.Errorf("Output map should not be empty")
	}
	if eventExecution.Output["sample_key"] != "sample_value" {
		t.Errorf("Expected Output['sample_key'] = 'sample_value', got %v", eventExecution.Output["sample_key"])
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(eventExecution)
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
