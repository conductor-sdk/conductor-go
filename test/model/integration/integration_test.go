package integration

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/model/integration"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserIntegration(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("Integration")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var integration integration.Integration
	err = json.Unmarshal([]byte(jsonTemplate), &integration)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// String fields
	if integration.Category != "API" {
		t.Errorf("Expected Category = 'API', got '%s'", integration.Category)
	}
	if integration.Description != "sample_description" {
		t.Errorf("Expected Description = 'sample_description', got '%s'", integration.Description)
	}
	if integration.Name != "sample_name" {
		t.Errorf("Expected Name = 'sample_name', got '%s'", integration.Name)
	}
	if integration.Type_ != "sample_type" {
		t.Errorf("Expected Type_ = 'sample_type', got '%s'", integration.Type_)
	}

	// Integer fields (only validate fields that exist in the struct)
	if integration.ModelsCount != 123 {
		t.Errorf("Expected ModelsCount = 123, got %d", integration.ModelsCount)
	}

	// Boolean field
	if !integration.Enabled {
		t.Errorf("Expected Enabled = true, got %v", integration.Enabled)
	}

	// Check Configuration map
	if integration.Configuration == nil {
		t.Errorf("Configuration map should not be nil")
	}
	if len(integration.Configuration) == 0 {
		t.Errorf("Configuration map should not be empty")
	}
	// Configuration uses ConfigKey as key, so check for resolved key
	// Based on template, this should be resolved from ConfigKey reference

	// Check Tags slice
	if integration.Tags == nil {
		t.Errorf("Tags slice should not be nil")
	}
	if len(integration.Tags) == 0 {
		t.Errorf("Tags slice should not be empty")
	}
	if len(integration.Tags) != 1 {
		t.Errorf("Expected Tags slice to have 1 element, got %d", len(integration.Tags))
	}

	// Validate TagObject element
	tag := integration.Tags[0]
	if tag.Key != "sample_key" {
		t.Errorf("Expected Tag.Key = 'sample_key', got '%s'", tag.Key)
	}
	if tag.Value != "sample_value" {
		t.Errorf("Expected Tag.Value = 'sample_value', got '%s'", tag.Value)
	}

	// Check Apis slice
	if integration.Apis == nil {
		t.Errorf("Apis slice should not be nil")
	}
	if len(integration.Apis) == 0 {
		t.Errorf("Apis slice should not be empty")
	}
	if len(integration.Apis) != 1 {
		t.Errorf("Expected Apis slice to have 1 element, got %d", len(integration.Apis))
	}

	// Validate IntegrationApi element
	api := integration.Apis[0]
	if api.IntegrationName != "sample_integrationName" {
		t.Errorf("Expected Api.IntegrationName = 'sample_integrationName', got '%s'", api.IntegrationName)
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(integration)
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

func TestSerDserIntegrationApi(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("IntegrationApi")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var integrationApi integration.IntegrationApi
	err = json.Unmarshal([]byte(jsonTemplate), &integrationApi)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// String fields
	if integrationApi.Api != "sample_api" {
		t.Errorf("Expected Api = 'sample_api', got '%s'", integrationApi.Api)
	}
	if integrationApi.Description != "sample_description" {
		t.Errorf("Expected Description = 'sample_description', got '%s'", integrationApi.Description)
	}
	if integrationApi.IntegrationName != "sample_integrationName" {
		t.Errorf("Expected IntegrationName = 'sample_integrationName', got '%s'", integrationApi.IntegrationName)
	}

	// Boolean field
	if !integrationApi.Enabled {
		t.Errorf("Expected Enabled = true, got %v", integrationApi.Enabled)
	}

	// Check Configuration map with ConfigKey
	if integrationApi.Configuration == nil {
		t.Errorf("Configuration map should not be nil")
	}
	if len(integrationApi.Configuration) == 0 {
		t.Errorf("Configuration map should not be empty")
	}

	// Check Tags slice
	if integrationApi.Tags == nil {
		t.Errorf("Tags slice should not be nil")
	}
	if len(integrationApi.Tags) == 0 {
		t.Errorf("Tags slice should not be empty")
	}
	if len(integrationApi.Tags) != 1 {
		t.Errorf("Expected Tags slice to have 1 element, got %d", len(integrationApi.Tags))
	}

	// Validate TagObject element
	tag := integrationApi.Tags[0]
	if tag.Key != "sample_key" {
		t.Errorf("Expected Tag.Key = 'sample_key', got '%s'", tag.Key)
	}
	if tag.Value != "sample_value" {
		t.Errorf("Expected Tag.Value = 'sample_value', got '%s'", tag.Value)
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(integrationApi)
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

func TestSerDserTagObject(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("Tag")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var tagObject model.TagObject
	err = json.Unmarshal([]byte(jsonTemplate), &tagObject)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	if tagObject.Key != "sample_key" {
		t.Errorf("Expected Key = 'sample_key', got '%s'", tagObject.Key)
	}
	if tagObject.Value != "sample_value" {
		t.Errorf("Expected Value = 'sample_value', got '%s'", tagObject.Value)
	}
	// Type_ field validation would depend on template content

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(tagObject)
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
