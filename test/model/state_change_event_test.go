package model

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserStateChangeEvent(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("StateChangeEvent")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var stateChangeEvent model.StateChangeEvent
	err = json.Unmarshal([]byte(jsonTemplate), &stateChangeEvent)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// String field
	if stateChangeEvent.Type_ != "sample_type" {
		t.Errorf("Expected Type = 'sample_type', got '%s'", stateChangeEvent.Type_)
	}

	// Map field
	if stateChangeEvent.Payload == nil {
		t.Errorf("Payload map should not be nil")
	}
	if len(stateChangeEvent.Payload) == 0 {
		t.Errorf("Payload map should not be empty")
	}
	if stateChangeEvent.Payload["key"] != "sample_value" {
		t.Errorf("Expected Payload['key'] = 'sample_value', got %v", stateChangeEvent.Payload["key"])
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(stateChangeEvent)
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
