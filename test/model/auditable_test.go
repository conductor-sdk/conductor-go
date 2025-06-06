package model

import (
	"encoding/json"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"reflect"
	"testing"

	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserAuditable(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("Auditable")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var auditable model.Auditable
	err = json.Unmarshal([]byte(jsonTemplate), &auditable)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	expected := model.Auditable{
		OwnerApp:   "sample_ownerApp",
		CreateTime: 123,
		UpdateTime: 123,
		CreatedBy:  "sample_createdBy",
		UpdatedBy:  "sample_updatedBy",
	}

	if !reflect.DeepEqual(auditable, expected) {
		t.Errorf("Deserialized struct mismatch:\nExpected: %+v\nActual: %+v", expected, auditable)
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(auditable)
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
