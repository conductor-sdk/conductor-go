package rbac

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model/rbac"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserUpsertUserRequest(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("UpsertUserRequest")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var upsertUserRequest rbac.UpsertUserRequest
	err = json.Unmarshal([]byte(jsonTemplate), &upsertUserRequest)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// String field
	if upsertUserRequest.Name != "sample_name" {
		t.Errorf("Expected Name = 'sample_name', got '%s'", upsertUserRequest.Name)
	}

	// Check Roles slice
	if upsertUserRequest.Roles == nil {
		t.Errorf("Roles slice should not be nil")
	}
	if len(upsertUserRequest.Roles) == 0 {
		t.Errorf("Roles slice should not be empty")
	}
	if len(upsertUserRequest.Roles) != 1 {
		t.Errorf("Expected Roles slice to have 1 element, got %d", len(upsertUserRequest.Roles))
	}
	if upsertUserRequest.Roles[0] != "USER" {
		t.Errorf("Expected Roles[0] = 'USER', got '%s'", upsertUserRequest.Roles[0])
	}

	// Check Groups slice
	if upsertUserRequest.Groups == nil {
		t.Errorf("Groups slice should not be nil")
	}
	if len(upsertUserRequest.Groups) == 0 {
		t.Errorf("Groups slice should not be empty")
	}
	if len(upsertUserRequest.Groups) != 1 {
		t.Errorf("Expected Groups slice to have 1 element, got %d", len(upsertUserRequest.Groups))
	}
	if upsertUserRequest.Groups[0] != "sample_groups" {
		t.Errorf("Expected Groups[0] = 'sample_groups', got '%s'", upsertUserRequest.Groups[0])
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(upsertUserRequest)
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
