package model

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserExternalStorageLocation(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("ExternalStorageLocation")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var externalStorageLocation model.ExternalStorageLocation
	err = json.Unmarshal([]byte(jsonTemplate), &externalStorageLocation)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	if externalStorageLocation.Uri != "sample_uri" {
		t.Errorf("Expected Uri = 'sample_uri', got '%s'", externalStorageLocation.Uri)
	}
	if externalStorageLocation.Path != "sample_path" {
		t.Errorf("Expected Path = 'sample_path', got '%s'", externalStorageLocation.Path)
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(externalStorageLocation)
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
