package rbac

import (
	"encoding/json"
	"github.com/conductor-sdk/conductor-go/sdk/model/rbac"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
	"reflect"
	"testing"
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

	// Check DefaultAccess map[string][]string (backward compatibility)
	if group.DefaultAccess == nil {
		t.Errorf("DefaultAccess map should not be nil")
	}
	if len(group.DefaultAccess) == 0 {
		t.Errorf("DefaultAccess map should not be empty")
	}

	// Validate map structure - should have ResourceType as key and []string as value
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
		// Validate first access in the list (based on your JSON output showing "CREATE")
		if len(accessList) > 0 {
			if accessList[0] != "CREATE" {
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
		t.Logf("Original: %+v", originalMap)
		t.Logf("Serialized: %+v", serializedMap)
	}

	// 6. Test new type-safe helper methods (after round-trip check)
	// Get the first resource type for testing
	var firstResourceType string
	for resourceType := range group.DefaultAccess {
		firstResourceType = resourceType
		break
	}

	if firstResourceType != "" {
		// Test GetDefaultAccessTyped method
		typedAccesses := group.GetDefaultAccessTyped(firstResourceType)
		if typedAccesses == nil {
			t.Errorf("GetDefaultAccessTyped should not return nil")
		}

		if len(typedAccesses) != len(group.DefaultAccess[firstResourceType]) {
			t.Errorf("Typed access length (%d) should match string access length (%d)",
				len(typedAccesses), len(group.DefaultAccess[firstResourceType]))
		}

		// Verify conversion between string and typed access
		for i, typedAccess := range typedAccesses {
			if string(typedAccess) != group.DefaultAccess[firstResourceType][i] {
				t.Errorf("Typed access[%d] (%s) should match string access (%s)",
					i, typedAccess, group.DefaultAccess[firstResourceType][i])
			}
		}

		// Test SetDefaultAccessTyped method
		testAccesses := []rbac.Access{rbac.CREATE, rbac.READ, rbac.UPDATE}
		group.SetDefaultAccessTyped("test_resource", testAccesses)

		// Verify it was set correctly as strings
		if group.DefaultAccess["test_resource"] == nil {
			t.Errorf("SetDefaultAccessTyped should set the access list")
		}

		expectedStrings := []string{"CREATE", "READ", "UPDATE"}
		if !reflect.DeepEqual(group.DefaultAccess["test_resource"], expectedStrings) {
			t.Errorf("SetDefaultAccessTyped should set correct string values, got %v",
				group.DefaultAccess["test_resource"])
		}

		// Test AddDefaultAccess method
		originalLength := len(group.DefaultAccess[firstResourceType])
		group.AddDefaultAccess(firstResourceType, rbac.DELETE)

		if len(group.DefaultAccess[firstResourceType]) != originalLength+1 {
			t.Errorf("AddDefaultAccess should increase access list length")
		}

		lastAccess := group.DefaultAccess[firstResourceType][len(group.DefaultAccess[firstResourceType])-1]
		if lastAccess != "DELETE" {
			t.Errorf("AddDefaultAccess should add DELETE access, got %s", lastAccess)
		}

		// Test HasDefaultAccess method
		if !group.HasDefaultAccess(firstResourceType, rbac.DELETE) {
			t.Errorf("HasDefaultAccess should return true for DELETE access")
		}

		if group.HasDefaultAccess(firstResourceType, rbac.Access("NONEXISTENT")) {
			t.Errorf("HasDefaultAccess should return false for non-existent access")
		}
	}
}
