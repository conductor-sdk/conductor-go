package rbac

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model/rbac"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserGroup(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("Group")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var group rbac.Group
	err = json.Unmarshal([]byte(jsonTemplate), &group)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// String fields
	if group.Description != "sample_description" {
		t.Errorf("Expected Description = 'sample_description', got '%s'", group.Description)
	}
	if group.Id != "sample_id" {
		t.Errorf("Expected Id = 'sample_id', got '%s'", group.Id)
	}

	// Check DefaultAccess map[string][]Access
	if group.DefaultAccess == nil {
		t.Errorf("DefaultAccess map should not be nil")
	}
	if len(group.DefaultAccess) == 0 {
		t.Errorf("DefaultAccess map should not be empty")
	}

	// Validate map structure - should have ResourceType as key and []Access as value
	for resourceType, accessList := range group.DefaultAccess {
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
			if accessList[0] != rbac.CREATE {
				t.Errorf("Expected first access = 'CREATE', got '%s'", accessList[0])
			}
		}
	}

	// Check Roles slice
	if group.Roles == nil {
		t.Errorf("Roles slice should not be nil")
	}
	if len(group.Roles) == 0 {
		t.Errorf("Roles slice should not be empty")
	}
	if len(group.Roles) != 1 {
		t.Errorf("Expected Roles slice to have 1 element, got %d", len(group.Roles))
	}

	// Validate Role element
	role := group.Roles[0]
	if role.Name != "sample_name" {
		t.Errorf("Expected Role.Name = 'sample_name', got '%s'", role.Name)
	}

	// Check Role.Permissions slice
	if role.Permissions == nil {
		t.Errorf("Role.Permissions slice should not be nil")
	}
	if len(role.Permissions) == 0 {
		t.Errorf("Role.Permissions slice should not be empty")
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(group)
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
