package rbac

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model/rbac"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserTargetRef(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("TargetRef")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var targetRef rbac.TargetRef
	err = json.Unmarshal([]byte(jsonTemplate), &targetRef)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// String fields
	if targetRef.Id != "sample_id" {
		t.Errorf("Expected Id = 'sample_id', got '%s'", targetRef.Id)
	}
	if targetRef.Type != "WORKFLOW_DEF" {
		t.Errorf("Expected Type = 'WORKFLOW_DEF', got '%s'", targetRef.Type)
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(targetRef)
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
