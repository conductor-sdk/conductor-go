package model

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserRateLimitConfig(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("RateLimitConfig")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var rateLimitConfig model.RateLimitConfig
	err = json.Unmarshal([]byte(jsonTemplate), &rateLimitConfig)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// String field
	if rateLimitConfig.RateLimitKey != "sample_rateLimitKey" {
		t.Errorf("Expected RateLimitKey = 'sample_rateLimitKey', got '%s'", rateLimitConfig.RateLimitKey)
	}

	// Integer field
	if rateLimitConfig.ConcurrentExecLimit != 123 {
		t.Errorf("Expected ConcurrentExecLimit = 123, got %d", rateLimitConfig.ConcurrentExecLimit)
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(rateLimitConfig)
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
