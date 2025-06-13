package model

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserTag(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("Tag")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var tag model.Tag
	err = json.Unmarshal([]byte(jsonTemplate), &tag)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// String fields
	if tag.Key != "sample_key" {
		t.Errorf("Expected Key = 'sample_key', got '%s'", tag.Key)
	}
	if tag.Value != "sample_value" {
		t.Errorf("Expected Value = 'sample_value', got '%s'", tag.Value)
	}

	// Deprecated field - still validate if present in template
	// Note: Type_ is deprecated as of 11/21/23 but may still be in templates
	if tag.Type_ != "" {
		t.Logf("Deprecated Type_ field present with value: %s", tag.Type_)
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(tag)
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
