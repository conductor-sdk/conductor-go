package model

import (
	"encoding/json"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserWorkflowDef(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("WorkflowDef")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var workflowDef model.WorkflowDef
	err = json.Unmarshal([]byte(jsonTemplate), &workflowDef)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// Auditable fields
	if workflowDef.OwnerApp != "sample_ownerApp" {
		t.Errorf("Expected OwnerApp = 'sample_ownerApp', got '%s'", workflowDef.OwnerApp)
	}
	if workflowDef.CreateTime != 123 {
		t.Errorf("Expected CreateTime = 123, got %d", workflowDef.CreateTime)
	}
	if workflowDef.UpdateTime != 123 {
		t.Errorf("Expected UpdateTime = 123, got %d", workflowDef.UpdateTime)
	}
	if workflowDef.CreatedBy != "sample_createdBy" {
		t.Errorf("Expected CreatedBy = 'sample_createdBy', got '%s'", workflowDef.CreatedBy)
	}
	if workflowDef.UpdatedBy != "sample_updatedBy" {
		t.Errorf("Expected UpdatedBy = 'sample_updatedBy', got '%s'", workflowDef.UpdatedBy)
	}

	// Core WorkflowDef string fields
	if workflowDef.Name != "sample_name" {
		t.Errorf("Expected Name = 'sample_name', got '%s'", workflowDef.Name)
	}
	if workflowDef.Description != "sample_description" {
		t.Errorf("Expected Description = 'sample_description', got '%s'", workflowDef.Description)
	}
	if workflowDef.FailureWorkflow != "sample_failureWorkflow" {
		t.Errorf("Expected FailureWorkflow = 'sample_failureWorkflow', got '%s'", workflowDef.FailureWorkflow)
	}
	if workflowDef.OwnerEmail != "sample_ownerEmail" {
		t.Errorf("Expected OwnerEmail = 'sample_ownerEmail', got '%s'", workflowDef.OwnerEmail)
	}
	if workflowDef.TimeoutPolicy != "TIME_OUT_WF" {
		t.Errorf("Expected TimeoutPolicy = 'TIME_OUT_WF', got '%s'", workflowDef.TimeoutPolicy)
	}
	if workflowDef.WorkflowStatusListenerSink != "sample_workflowStatusListenerSink" {
		t.Errorf("Expected WorkflowStatusListenerSink = 'sample_workflowStatusListenerSink', got '%s'", workflowDef.WorkflowStatusListenerSink)
	}

	// Int32 fields
	if workflowDef.Version != 123 {
		t.Errorf("Expected Version = 123, got %d", workflowDef.Version)
	}
	if workflowDef.SchemaVersion != 123 {
		t.Errorf("Expected SchemaVersion = 123, got %d", workflowDef.SchemaVersion)
	}

	// Int64 field
	if workflowDef.TimeoutSeconds != 123 {
		t.Errorf("Expected TimeoutSeconds = 123, got %d", workflowDef.TimeoutSeconds)
	}

	// Boolean fields
	if !workflowDef.Restartable {
		t.Errorf("Expected Restartable = true, got %v", workflowDef.Restartable)
	}
	if !workflowDef.WorkflowStatusListenerEnabled {
		t.Errorf("Expected WorkflowStatusListenerEnabled = true, got %v", workflowDef.WorkflowStatusListenerEnabled)
	}
	if !workflowDef.EnforceSchema {
		t.Errorf("Expected EnforceSchema = true, got %v", workflowDef.EnforceSchema)
	}

	// String slice field
	if workflowDef.InputParameters == nil {
		t.Errorf("InputParameters slice should not be nil")
	}
	if len(workflowDef.InputParameters) == 0 {
		t.Errorf("InputParameters slice should not be empty")
	} else {
		if workflowDef.InputParameters[0] != "sample_inputParameters" {
			t.Errorf("Expected InputParameters[0] = 'sample_inputParameters', got '%s'", workflowDef.InputParameters[0])
		}
	}

	// Map fields
	if workflowDef.OutputParameters == nil {
		t.Errorf("OutputParameters map should not be nil")
	}
	if len(workflowDef.OutputParameters) == 0 {
		t.Errorf("OutputParameters map should not be empty")
	} else {
		if workflowDef.OutputParameters["sample_key"] != "sample_value" {
			t.Errorf("Expected OutputParameters['sample_key'] = 'sample_value', got %v", workflowDef.OutputParameters["sample_key"])
		}
	}

	if workflowDef.Variables == nil {
		t.Errorf("Variables map should not be nil")
	}
	if len(workflowDef.Variables) == 0 {
		t.Errorf("Variables map should not be empty")
	} else {
		if workflowDef.Variables["sample_key"] != "sample_value" {
			t.Errorf("Expected Variables['sample_key'] = 'sample_value', got %v", workflowDef.Variables["sample_key"])
		}
	}

	if workflowDef.InputTemplate == nil {
		t.Errorf("InputTemplate map should not be nil")
	}
	if len(workflowDef.InputTemplate) == 0 {
		t.Errorf("InputTemplate map should not be empty")
	} else {
		if workflowDef.InputTemplate["sample_key"] != "sample_value" {
			t.Errorf("Expected InputTemplate['sample_key'] = 'sample_value', got %v", workflowDef.InputTemplate["sample_key"])
		}
	}

	if workflowDef.Metadata == nil {
		t.Errorf("Metadata map should not be nil")
	}
	if len(workflowDef.Metadata) == 0 {
		t.Errorf("Metadata map should not be empty")
	} else {
		if workflowDef.Metadata["sample_key"] != "sample_value" {
			t.Errorf("Expected Metadata['sample_key'] = 'sample_value', got %v", workflowDef.Metadata["sample_key"])
		}
	}

	// WorkflowTask slice field
	if workflowDef.Tasks == nil {
		t.Errorf("Tasks slice should not be nil")
	}
	if len(workflowDef.Tasks) == 0 {
		t.Logf("Tasks slice is empty - this may be expected for some templates")
	} else {
		// Validate first task only if tasks exist
		task := workflowDef.Tasks[0]
		if task.Name != "sample_name" {
			t.Errorf("Expected Tasks[0].Name = 'sample_name', got '%s'", task.Name)
		}
		if task.TaskReferenceName != "sample_taskReferenceName" {
			t.Errorf("Expected Tasks[0].TaskReferenceName = 'sample_taskReferenceName', got '%s'", task.TaskReferenceName)
		}
	}

	// Pointer fields
	if workflowDef.RateLimitConfig == nil {
		t.Errorf("RateLimitConfig should not be nil")
	} else {
		if workflowDef.RateLimitConfig.RateLimitKey != "sample_rateLimitKey" {
			t.Errorf("Expected RateLimitConfig.RateLimitKey = 'sample_rateLimitKey', got '%s'", workflowDef.RateLimitConfig.RateLimitKey)
		}
		if workflowDef.RateLimitConfig.ConcurrentExecLimit != 123 {
			t.Errorf("Expected RateLimitConfig.ConcurrentExecLimit = 123, got %d", workflowDef.RateLimitConfig.ConcurrentExecLimit)
		}
	}

	if workflowDef.InputSchema == nil {
		t.Errorf("InputSchema should not be nil")
	} else {
		if workflowDef.InputSchema.Name != "sample_name" {
			t.Errorf("Expected InputSchema.Name = 'sample_name', got '%s'", workflowDef.InputSchema.Name)
		}
		if workflowDef.InputSchema.Type != "JSON" {
			t.Errorf("Expected InputSchema.Type = 'JSON', got '%s'", workflowDef.InputSchema.Type)
		}
	}

	if workflowDef.OutputSchema == nil {
		t.Errorf("OutputSchema should not be nil")
	} else {
		if workflowDef.OutputSchema.Name != "sample_name" {
			t.Errorf("Expected OutputSchema.Name = 'sample_name', got '%s'", workflowDef.OutputSchema.Name)
		}
		if workflowDef.OutputSchema.Type != "JSON" {
			t.Errorf("Expected OutputSchema.Type = 'JSON', got '%s'", workflowDef.OutputSchema.Type)
		}
	}

	// Deprecated fields - validate if present but don't require them
	if workflowDef.Tags != nil && len(workflowDef.Tags) > 0 {
		t.Logf("Deprecated Tags field present with %d elements", len(workflowDef.Tags))
		tag := workflowDef.Tags[0]
		if tag.Key != "sample_key" {
			t.Errorf("Expected Tags[0].Key = 'sample_key', got '%s'", tag.Key)
		}
		if tag.Value != "sample_value" {
			t.Errorf("Expected Tags[0].Value = 'sample_value', got '%s'", tag.Value)
		}
	}
	if workflowDef.OverwriteTags {
		t.Logf("Deprecated OverwriteTags field present with value: %v", workflowDef.OverwriteTags)
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(workflowDef)
	if err != nil {
		t.Fatalf("Failed to serialize struct: %v", err)
	}

	// 5. Round-trip integrity check
	var roundTripWorkflowDef model.WorkflowDef
	err = json.Unmarshal(serializedJSON, &roundTripWorkflowDef)
	if err != nil {
		t.Fatalf("Failed to deserialize round-trip JSON: %v", err)
	}

	// Compare key fields to ensure data integrity
	if roundTripWorkflowDef.Name != workflowDef.Name {
		t.Errorf("Round-trip Name mismatch: %s vs %s", workflowDef.Name, roundTripWorkflowDef.Name)
	}
	if roundTripWorkflowDef.Description != workflowDef.Description {
		t.Errorf("Round-trip Description mismatch: %s vs %s", workflowDef.Description, roundTripWorkflowDef.Description)
	}
	if roundTripWorkflowDef.Version != workflowDef.Version {
		t.Errorf("Round-trip Version mismatch: %d vs %d", workflowDef.Version, roundTripWorkflowDef.Version)
	}
	if roundTripWorkflowDef.TimeoutSeconds != workflowDef.TimeoutSeconds {
		t.Errorf("Round-trip TimeoutSeconds mismatch: %d vs %d", workflowDef.TimeoutSeconds, roundTripWorkflowDef.TimeoutSeconds)
	}
	if roundTripWorkflowDef.TimeoutPolicy != workflowDef.TimeoutPolicy {
		t.Errorf("Round-trip TimeoutPolicy mismatch: %s vs %s", workflowDef.TimeoutPolicy, roundTripWorkflowDef.TimeoutPolicy)
	}
	if roundTripWorkflowDef.Restartable != workflowDef.Restartable {
		t.Errorf("Round-trip Restartable mismatch: %v vs %v", workflowDef.Restartable, roundTripWorkflowDef.Restartable)
	}
	if roundTripWorkflowDef.EnforceSchema != workflowDef.EnforceSchema {
		t.Errorf("Round-trip EnforceSchema mismatch: %v vs %v", workflowDef.EnforceSchema, roundTripWorkflowDef.EnforceSchema)
	}

	// Compare pointer fields
	if (workflowDef.RateLimitConfig == nil) != (roundTripWorkflowDef.RateLimitConfig == nil) {
		t.Errorf("Round-trip RateLimitConfig nil mismatch")
	}
	if (workflowDef.InputSchema == nil) != (roundTripWorkflowDef.InputSchema == nil) {
		t.Errorf("Round-trip InputSchema nil mismatch")
	}
	if (workflowDef.OutputSchema == nil) != (roundTripWorkflowDef.OutputSchema == nil) {
		t.Errorf("Round-trip OutputSchema nil mismatch")
	}
}
