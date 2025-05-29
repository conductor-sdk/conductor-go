package integration

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model/integration"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserCategory(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("Category")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var category integration.Category
	err = json.Unmarshal([]byte(jsonTemplate), &category)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized enum value
	if category != integration.API {
		t.Errorf("Expected Category = 'API', got '%s'", category)
	}

	// Validate it's one of the valid enum values
	validCategories := []integration.Category{
		integration.API,
		integration.AI_MODEL,
		integration.VECTOR_DB,
		integration.RELATIONAL_DB,
		integration.MESSAGE_BROKER,
		integration.GIT,
		integration.EMAIL,
	}

	isValid := false
	for _, validCategory := range validCategories {
		if category == validCategory {
			isValid = true
			break
		}
	}
	if !isValid {
		t.Errorf("Category '%s' is not a valid enum value", category)
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(category)
	if err != nil {
		t.Fatalf("Failed to serialize struct: %v", err)
	}

	// 5. Round-trip integrity check
	var originalValue, serializedValue string
	json.Unmarshal([]byte(jsonTemplate), &originalValue)
	json.Unmarshal(serializedJSON, &serializedValue)

	if !reflect.DeepEqual(originalValue, serializedValue) {
		t.Errorf("Round-trip integrity failed: expected '%s', got '%s'", originalValue, serializedValue)
	}
}
