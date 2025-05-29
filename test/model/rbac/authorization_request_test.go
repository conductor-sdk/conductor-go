package rbac

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model/rbac"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserAuthorizationRequest(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("AuthorizationRequest")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var authorizationRequest rbac.AuthorizationRequest
	err = json.Unmarshal([]byte(jsonTemplate), &authorizationRequest)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// Check Access slice
	if authorizationRequest.Access == nil {
		t.Errorf("Access slice should not be nil")
	}
	if len(authorizationRequest.Access) == 0 {
		t.Errorf("Access slice should not be empty")
	}
	if len(authorizationRequest.Access) != 1 {
		t.Errorf("Expected Access slice to have 1 element, got %d", len(authorizationRequest.Access))
	}
	if authorizationRequest.Access[0] != "CREATE" {
		t.Errorf("Expected Access[0] = 'CREATE', got '%s'", authorizationRequest.Access[0])
	}

	// Check Subject pointer field
	if authorizationRequest.Subject == nil {
		t.Errorf("Subject should not be nil")
	} else {
		if authorizationRequest.Subject.Id != "sample_id" {
			t.Errorf("Expected Subject.ID = 'sample_id', got '%s'", authorizationRequest.Subject.Id)
		}
		if authorizationRequest.Subject.Type_ != "USER" {
			t.Errorf("Expected Subject.Type = 'USER', got '%s'", authorizationRequest.Subject.Type_)
		}
	}

	// Check Target pointer field
	if authorizationRequest.Target == nil {
		t.Errorf("Target should not be nil")
	} else {
		if authorizationRequest.Target.Id != "sample_id" {
			t.Errorf("Expected Target.ID = 'sample_id', got '%s'", authorizationRequest.Target.Id)
		}
		if authorizationRequest.Target.Type != "WORKFLOW_DEF" {
			t.Errorf("Expected Target.Type = 'WORKFLOW_DEF', got '%s'", authorizationRequest.Target.Type)
		}
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(authorizationRequest)
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
