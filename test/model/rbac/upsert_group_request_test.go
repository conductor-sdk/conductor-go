package rbac

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model/rbac"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserUpsertGroupRequest(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("UpsertGroupRequest")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var upsertGroupRequest rbac.UpsertGroupRequest
	err = json.Unmarshal([]byte(jsonTemplate), &upsertGroupRequest)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// String field
	if upsertGroupRequest.Description != "sample_description" {
		t.Errorf("Expected Description = 'sample_description', got '%s'", upsertGroupRequest.Description)
	}

	// Check Roles slice (string slice with role names)
	if upsertGroupRequest.Roles == nil {
		t.Errorf("Roles slice should not be nil")
	}
	if len(upsertGroupRequest.Roles) == 0 {
		t.Errorf("Roles slice should not be empty")
	}
	if len(upsertGroupRequest.Roles) != 1 {
		t.Errorf("Expected Roles slice to have 1 element, got %d", len(upsertGroupRequest.Roles))
	}
	if upsertGroupRequest.Roles[0] != "USER" {
		t.Errorf("Expected Roles[0] = 'USER', got '%s'", upsertGroupRequest.Roles[0])
	}

	// Check DefaultAccess map[string][]string
	if upsertGroupRequest.DefaultAccess == nil {
		t.Errorf("DefaultAccess map should not be nil")
	}
	if len(upsertGroupRequest.DefaultAccess) == 0 {
		t.Errorf("DefaultAccess map should not be empty")
	}

	// Validate map structure - should have ResourceType as key and []Access as value
	for resourceType, accessList := range upsertGroupRequest.DefaultAccess {
		if resourceType == "" {
			t.Errorf("ResourceType key should not be empty")
		}
		if accessList == nil {
			t.Errorf("Access list should not be nil for resource type: %s", resourceType)
		}
		if len(accessList) == 0 {
			t.Errorf("Access list should not be empty for resource type: %s", resourceType)
		}
		// Validate first access in the list
		if len(accessList) > 0 {
			if accessList[0] != "CREATE" {
				t.Errorf("Expected first access = 'CREATE', got '%s'", accessList[0])
			}
		}
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(upsertGroupRequest)
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
