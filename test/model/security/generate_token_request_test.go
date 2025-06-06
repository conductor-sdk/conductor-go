package security

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model/security"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserGenerateTokenRequest(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("GenerateTokenRequest")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var generateTokenRequest security.GenerateTokenRequest
	err = json.Unmarshal([]byte(jsonTemplate), &generateTokenRequest)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// String fields
	if generateTokenRequest.KeyId != "sample_keyId" {
		t.Errorf("Expected KeyId = 'sample_keyId', got '%s'", generateTokenRequest.KeyId)
	}
	if generateTokenRequest.KeySecret != "sample_keySecret" {
		t.Errorf("Expected KeySecret = 'sample_keySecret', got '%s'", generateTokenRequest.KeySecret)
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(generateTokenRequest)
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
