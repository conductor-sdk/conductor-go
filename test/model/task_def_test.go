package model

import (
	"encoding/json"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserTaskDef(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("TaskDef")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var taskDef model.TaskDef
	err = json.Unmarshal([]byte(jsonTemplate), &taskDef)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// Auditable fields (not inherited but present as individual fields)
	if taskDef.OwnerApp != "sample_ownerApp" {
		t.Errorf("Expected OwnerApp = 'sample_ownerApp', got '%s'", taskDef.OwnerApp)
	}
	if taskDef.CreateTime != 123 {
		t.Errorf("Expected CreateTime = 123, got %d", taskDef.CreateTime)
	}
	if taskDef.UpdateTime != 123 {
		t.Errorf("Expected UpdateTime = 123, got %d", taskDef.UpdateTime)
	}
	if taskDef.CreatedBy != "sample_createdBy" {
		t.Errorf("Expected CreatedBy = 'sample_createdBy', got '%s'", taskDef.CreatedBy)
	}
	if taskDef.UpdatedBy != "sample_updatedBy" {
		t.Errorf("Expected UpdatedBy = 'sample_updatedBy', got '%s'", taskDef.UpdatedBy)
	}

	// Core TaskDef string fields
	if taskDef.Name != "sample_name" {
		t.Errorf("Expected Name = 'sample_name', got '%s'", taskDef.Name)
	}
	if taskDef.Description != "sample_description" {
		t.Errorf("Expected Description = 'sample_description', got '%s'", taskDef.Description)
	}
	if taskDef.TimeoutPolicy != "RETRY" {
		t.Errorf("Expected TimeoutPolicy = 'RETRY', got '%s'", taskDef.TimeoutPolicy)
	}
	if taskDef.RetryLogic != "FIXED" {
		t.Errorf("Expected RetryLogic = 'FIXED', got '%s'", taskDef.RetryLogic)
	}
	if taskDef.IsolationGroupId != "sample_isolationGroupId" {
		t.Errorf("Expected IsolationGroupId = 'sample_isolationGroupId', got '%s'", taskDef.IsolationGroupId)
	}
	if taskDef.ExecutionNameSpace != "sample_executionNameSpace" {
		t.Errorf("Expected ExecutionNameSpace = 'sample_executionNameSpace', got '%s'", taskDef.ExecutionNameSpace)
	}
	if taskDef.OwnerEmail != "sample_ownerEmail" {
		t.Errorf("Expected OwnerEmail = 'sample_ownerEmail', got '%s'", taskDef.OwnerEmail)
	}
	if taskDef.BaseType != "sample_baseType" {
		t.Errorf("Expected BaseType = 'sample_baseType', got '%s'", taskDef.BaseType)
	}

	// Int32 fields
	if taskDef.RetryCount != 123 {
		t.Errorf("Expected RetryCount = 123, got %d", taskDef.RetryCount)
	}
	if taskDef.RetryDelaySeconds != 123 {
		t.Errorf("Expected RetryDelaySeconds = 123, got %d", taskDef.RetryDelaySeconds)
	}
	if taskDef.ConcurrentExecLimit != 123 {
		t.Errorf("Expected ConcurrentExecLimit = 123, got %d", taskDef.ConcurrentExecLimit)
	}
	if taskDef.RateLimitPerFrequency != 123 {
		t.Errorf("Expected RateLimitPerFrequency = 123, got %d", taskDef.RateLimitPerFrequency)
	}
	if taskDef.RateLimitFrequencyInSeconds != 123 {
		t.Errorf("Expected RateLimitFrequencyInSeconds = 123, got %d", taskDef.RateLimitFrequencyInSeconds)
	}
	if taskDef.PollTimeoutSeconds != 123 {
		t.Errorf("Expected PollTimeoutSeconds = 123, got %d", taskDef.PollTimeoutSeconds)
	}
	if taskDef.BackoffScaleFactor != 123 {
		t.Errorf("Expected BackoffScaleFactor = 123, got %d", taskDef.BackoffScaleFactor)
	}

	// Int64 fields
	if taskDef.TimeoutSeconds != 123 {
		t.Errorf("Expected TimeoutSeconds = 123, got %d", taskDef.TimeoutSeconds)
	}
	if taskDef.ResponseTimeoutSeconds != 123 {
		t.Errorf("Expected ResponseTimeoutSeconds = 123, got %d", taskDef.ResponseTimeoutSeconds)
	}
	if taskDef.TotalTimeoutSeconds != 123 {
		t.Errorf("Expected TotalTimeoutSeconds = 123, got %d", taskDef.TotalTimeoutSeconds)
	}

	// Boolean field
	if !taskDef.EnforceSchema {
		t.Errorf("Expected EnforceSchema = true, got %v", taskDef.EnforceSchema)
	}

	// String slice fields
	if taskDef.InputKeys == nil {
		t.Errorf("InputKeys slice should not be nil")
	}
	if len(taskDef.InputKeys) == 0 {
		t.Errorf("InputKeys slice should not be empty")
	}
	if taskDef.InputKeys[0] != "sample_inputKeys" {
		t.Errorf("Expected InputKeys[0] = 'sample_inputKeys', got '%s'", taskDef.InputKeys[0])
	}

	if taskDef.OutputKeys == nil {
		t.Errorf("OutputKeys slice should not be nil")
	}
	if len(taskDef.OutputKeys) == 0 {
		t.Errorf("OutputKeys slice should not be empty")
	}
	if taskDef.OutputKeys[0] != "sample_outputKeys" {
		t.Errorf("Expected OutputKeys[0] = 'sample_outputKeys', got '%s'", taskDef.OutputKeys[0])
	}

	// Map field
	if taskDef.InputTemplate == nil {
		t.Errorf("InputTemplate map should not be nil")
	}
	if len(taskDef.InputTemplate) == 0 {
		t.Errorf("InputTemplate map should not be empty")
	}
	if taskDef.InputTemplate["sample_key"] != "sample_value" {
		t.Errorf("Expected InputTemplate['sample_key'] = 'sample_value', got %v", taskDef.InputTemplate["sample_key"])
	}

	// Embedded struct fields (SchemaDef)
	if taskDef.InputSchema.Name != "sample_name" {
		t.Errorf("Expected InputSchema.Name = 'sample_name', got '%s'", taskDef.InputSchema.Name)
	}
	if taskDef.InputSchema.Type != "JSON" {
		t.Errorf("Expected InputSchema.Type = 'JSON', got '%s'", taskDef.InputSchema.Type)
	}
	if taskDef.InputSchema.Version != 1 {
		t.Errorf("Expected InputSchema.Version = 1, got %d", taskDef.InputSchema.Version)
	}

	if taskDef.OutputSchema.Name != "sample_name" {
		t.Errorf("Expected OutputSchema.Name = 'sample_name', got '%s'", taskDef.OutputSchema.Name)
	}
	if taskDef.OutputSchema.Type != "JSON" {
		t.Errorf("Expected OutputSchema.Type = 'JSON', got '%s'", taskDef.OutputSchema.Type)
	}
	if taskDef.OutputSchema.Version != 1 {
		t.Errorf("Expected OutputSchema.Version = 1, got %d", taskDef.OutputSchema.Version)
	}

	// Deprecated fields - validate if present but don't require them
	if taskDef.Tags != nil && len(taskDef.Tags) > 0 {
		t.Logf("Deprecated Tags field present with %d elements", len(taskDef.Tags))
		// Validate first tag if present
		tag := taskDef.Tags[0]
		if tag.Key != "sample_key" {
			t.Errorf("Expected Tags[0].Key = 'sample_key', got '%s'", tag.Key)
		}
		if tag.Value != "sample_value" {
			t.Errorf("Expected Tags[0].Value = 'sample_value', got '%s'", tag.Value)
		}
	}
	if taskDef.OverwriteTags {
		t.Logf("Deprecated OverwriteTags field present with value: %v", taskDef.OverwriteTags)
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(taskDef)
	if err != nil {
		t.Fatalf("Failed to serialize struct: %v", err)
	}

	// 5. Round-trip integrity check
	var roundTripTaskDef model.TaskDef
	err = json.Unmarshal(serializedJSON, &roundTripTaskDef)
	if err != nil {
		t.Fatalf("Failed to deserialize round-trip JSON: %v", err)
	}

	// Compare key fields to ensure data integrity
	if roundTripTaskDef.Name != taskDef.Name {
		t.Errorf("Round-trip Name mismatch: %s vs %s", taskDef.Name, roundTripTaskDef.Name)
	}
	if roundTripTaskDef.Description != taskDef.Description {
		t.Errorf("Round-trip Description mismatch: %s vs %s", taskDef.Description, roundTripTaskDef.Description)
	}
	if roundTripTaskDef.RetryCount != taskDef.RetryCount {
		t.Errorf("Round-trip RetryCount mismatch: %d vs %d", taskDef.RetryCount, roundTripTaskDef.RetryCount)
	}
	if roundTripTaskDef.TimeoutSeconds != taskDef.TimeoutSeconds {
		t.Errorf("Round-trip TimeoutSeconds mismatch: %d vs %d", taskDef.TimeoutSeconds, roundTripTaskDef.TimeoutSeconds)
	}
	if roundTripTaskDef.TimeoutPolicy != taskDef.TimeoutPolicy {
		t.Errorf("Round-trip TimeoutPolicy mismatch: %s vs %s", taskDef.TimeoutPolicy, roundTripTaskDef.TimeoutPolicy)
	}
	if roundTripTaskDef.RetryLogic != taskDef.RetryLogic {
		t.Errorf("Round-trip RetryLogic mismatch: %s vs %s", taskDef.RetryLogic, roundTripTaskDef.RetryLogic)
	}
	if roundTripTaskDef.EnforceSchema != taskDef.EnforceSchema {
		t.Errorf("Round-trip EnforceSchema mismatch: %v vs %v", taskDef.EnforceSchema, roundTripTaskDef.EnforceSchema)
	}

	// Compare embedded structs
	if roundTripTaskDef.InputSchema.Name != taskDef.InputSchema.Name {
		t.Errorf("Round-trip InputSchema.Name mismatch: %s vs %s",
			taskDef.InputSchema.Name, roundTripTaskDef.InputSchema.Name)
	}
	if roundTripTaskDef.OutputSchema.Type != taskDef.OutputSchema.Type {
		t.Errorf("Round-trip OutputSchema.Type mismatch: %s vs %s",
			taskDef.OutputSchema.Type, roundTripTaskDef.OutputSchema.Type)
	}
}
