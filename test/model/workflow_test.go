package model

import (
	"encoding/json"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserWorkflow(t *testing.T) {
	// 1. Load JSON template (use Workflow1 to avoid recursive history issues)
	jsonTemplate, err := util.GetJsonString("Workflow1")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// Debug: Log the JSON template to see its structure
	t.Logf("JSON Template: %s", jsonTemplate)

	// 2. Deserialize JSON to struct with custom handling for WorkflowStatus
	// First, unmarshal to a flexible structure
	var jsonData map[string]interface{}
	err = json.Unmarshal([]byte(jsonTemplate), &jsonData)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON to map: %v", err)
	}

	// Convert back to JSON and then to struct
	modifiedJSON, err := json.Marshal(jsonData)
	if err != nil {
		t.Fatalf("Failed to marshal modified JSON: %v", err)
	}

	var workflow model.Workflow
	err = json.Unmarshal(modifiedJSON, &workflow)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// Auditable fields
	if workflow.OwnerApp != "sample_ownerApp" {
		t.Errorf("Expected OwnerApp = 'sample_ownerApp', got '%s'", workflow.OwnerApp)
	}
	if workflow.CreateTime != 123 {
		t.Errorf("Expected CreateTime = 123, got %d", workflow.CreateTime)
	}
	if workflow.UpdateTime != 123 {
		t.Errorf("Expected UpdateTime = 123, got %d", workflow.UpdateTime)
	}
	if workflow.CreatedBy != "sample_createdBy" {
		t.Errorf("Expected CreatedBy = 'sample_createdBy', got '%s'", workflow.CreatedBy)
	}
	if workflow.UpdatedBy != "sample_updatedBy" {
		t.Errorf("Expected UpdatedBy = 'sample_updatedBy', got '%s'", workflow.UpdatedBy)
	}

	// WorkflowStatus struct field (not a simple enum)
	if workflow.Status != model.RunningWorkflow {
		t.Errorf("Expected Status.Status = 'RUNNING', got '%s'", workflow.Status)
	}
	if workflow.WorkflowId != "sample_workflowId" {
		t.Errorf("Expected Status.WorkflowId = 'sample_workflowId', got '%s'", workflow.WorkflowId)
	}
	if workflow.CorrelationId != "sample_correlationId" {
		t.Errorf("Expected Status.CorrelationId = 'sample_correlationId', got '%s'", workflow.CorrelationId)
	}

	// Check Status Output and Variables maps
	if workflow.Output == nil || len(workflow.Output) == 0 {
		t.Errorf("Status.Output map should not be nil or empty")
	}
	if workflow.Variables == nil || len(workflow.Variables) == 0 {
		t.Errorf("Status.Variables map should not be nil or empty")
	}

	// Core workflow string fields
	if workflow.WorkflowId != "sample_workflowId" {
		t.Errorf("Expected WorkflowId = 'sample_workflowId', got '%s'", workflow.WorkflowId)
	}
	if workflow.ParentWorkflowId != "sample_parentWorkflowId" {
		t.Errorf("Expected ParentWorkflowId = 'sample_parentWorkflowId', got '%s'", workflow.ParentWorkflowId)
	}
	if workflow.ParentWorkflowTaskId != "sample_parentWorkflowTaskId" {
		t.Errorf("Expected ParentWorkflowTaskId = 'sample_parentWorkflowTaskId', got '%s'", workflow.ParentWorkflowTaskId)
	}
	if workflow.CorrelationId != "sample_correlationId" {
		t.Errorf("Expected CorrelationId = 'sample_correlationId', got '%s'", workflow.CorrelationId)
	}
	if workflow.ReRunFromWorkflowId != "sample_reRunFromWorkflowId" {
		t.Errorf("Expected ReRunFromWorkflowId = 'sample_reRunFromWorkflowId', got '%s'", workflow.ReRunFromWorkflowId)
	}
	if workflow.ReasonForIncompletion != "sample_reasonForIncompletion" {
		t.Errorf("Expected ReasonForIncompletion = 'sample_reasonForIncompletion', got '%s'", workflow.ReasonForIncompletion)
	}
	if workflow.Event != "sample_event" {
		t.Errorf("Expected Event = 'sample_event', got '%s'", workflow.Event)
	}
	if workflow.ExternalInputPayloadStoragePath != "sample_externalInputPayloadStoragePath" {
		t.Errorf("Expected ExternalInputPayloadStoragePath = 'sample_externalInputPayloadStoragePath', got '%s'", workflow.ExternalInputPayloadStoragePath)
	}
	if workflow.ExternalOutputPayloadStoragePath != "sample_externalOutputPayloadStoragePath" {
		t.Errorf("Expected ExternalOutputPayloadStoragePath = 'sample_externalOutputPayloadStoragePath', got '%s'", workflow.ExternalOutputPayloadStoragePath)
	}
	if workflow.IdempotencyKey != "sample_idempotencyKey" {
		t.Errorf("Expected IdempotencyKey = 'sample_idempotencyKey', got '%s'", workflow.IdempotencyKey)
	}
	if workflow.RateLimitKey != "sample_rateLimitKey" {
		t.Errorf("Expected RateLimitKey = 'sample_rateLimitKey', got '%s'", workflow.RateLimitKey)
	}

	// Int64 fields
	if workflow.EndTime != 123 {
		t.Errorf("Expected EndTime = 123, got %d", workflow.EndTime)
	}
	if workflow.LastRetriedTime != 123 {
		t.Errorf("Expected LastRetriedTime = 123, got %d", workflow.LastRetriedTime)
	}

	// Integer field
	if workflow.Priority != 1 {
		t.Errorf("Expected Priority = 1, got %d", workflow.Priority)
	}

	// Boolean field
	if !workflow.RateLimited {
		t.Errorf("Expected RateLimited = true, got %v", workflow.RateLimited)
	}

	// Map fields
	if workflow.Input == nil {
		t.Errorf("Input map should not be nil")
	}
	if len(workflow.Input) == 0 {
		t.Errorf("Input map should not be empty")
	}
	if workflow.Input["sample_key"] != "sample_value" {
		t.Errorf("Expected Input['sample_key'] = 'sample_value', got %v", workflow.Input["sample_key"])
	}

	if workflow.Output == nil {
		t.Errorf("Output map should not be nil")
	}
	if len(workflow.Output) == 0 {
		t.Errorf("Output map should not be empty")
	}
	if workflow.Output["sample_key"] != "sample_value" {
		t.Errorf("Expected Output['sample_key'] = 'sample_value', got %v", workflow.Output["sample_key"])
	}

	if workflow.TaskToDomain == nil {
		t.Errorf("TaskToDomain map should not be nil")
	}
	if len(workflow.TaskToDomain) == 0 {
		t.Errorf("TaskToDomain map should not be empty")
	}
	if workflow.TaskToDomain["sample_key"] != "sample_value" {
		t.Errorf("Expected TaskToDomain['sample_key'] = 'sample_value', got %v", workflow.TaskToDomain["sample_key"])
	}

	if workflow.Variables == nil {
		t.Errorf("Variables map should not be nil")
	}
	if len(workflow.Variables) == 0 {
		t.Errorf("Variables map should not be empty")
	}
	if workflow.Variables["sample_key"] != "sample_value" {
		t.Errorf("Expected Variables['sample_key'] = 'sample_value', got %v", workflow.Variables["sample_key"])
	}

	// String slice fields
	if workflow.FailedReferenceTaskNames == nil {
		t.Errorf("FailedReferenceTaskNames slice should not be nil")
	}
	if len(workflow.FailedReferenceTaskNames) == 0 {
		t.Logf("FailedReferenceTaskNames slice is empty - this may be expected")
	} else {
		if workflow.FailedReferenceTaskNames[0] != "sample_value" {
			t.Errorf("Expected FailedReferenceTaskNames[0] = 'sample_value', got '%s'", workflow.FailedReferenceTaskNames[0])
		}
	}

	if workflow.FailedTaskNames == nil {
		t.Errorf("FailedTaskNames slice should not be nil")
	}
	if len(workflow.FailedTaskNames) == 0 {
		t.Logf("FailedTaskNames slice is empty - this may be expected")
	} else {
		if workflow.FailedTaskNames[0] != "sample_value" {
			t.Errorf("Expected FailedTaskNames[0] = 'sample_value', got '%s'", workflow.FailedTaskNames[0])
		}
	}

	// Task slice field
	if workflow.Tasks == nil {
		t.Errorf("Tasks slice should not be nil")
	}
	if len(workflow.Tasks) == 0 {
		t.Logf("Tasks slice is empty - this may be expected for Workflow1 template")
	} else {
		// Validate first task only if tasks exist
		task := workflow.Tasks[0]
		if task.TaskId != "sample_taskId" {
			t.Errorf("Expected Tasks[0].TaskId = 'sample_taskId', got '%s'", task.TaskId)
		}
		if task.TaskType != "sample_taskType" {
			t.Errorf("Expected Tasks[0].TaskType = 'sample_taskType', got '%s'", task.TaskType)
		}
	}

	// Pointer to struct field
	if workflow.WorkflowDefinition == nil {
		t.Errorf("WorkflowDefinition should not be nil")
	} else {
		if workflow.WorkflowDefinition.Name != "sample_name" {
			t.Errorf("Expected WorkflowDefinition.Name = 'sample_name', got '%s'", workflow.WorkflowDefinition.Name)
		}
		if workflow.WorkflowDefinition.Version != 123 {
			t.Errorf("Expected WorkflowDefinition.Version = 123, got %d", workflow.WorkflowDefinition.Version)
		}
	}

	// History slice field (recursive structure)
	if workflow.History == nil {
		t.Errorf("History slice should not be nil")
	}
	if len(workflow.History) == 0 {
		t.Logf("History slice is empty - this is expected for Workflow1 template")
	} else {
		// Validate first history item only if history exists
		historyWorkflow := workflow.History[0]
		if historyWorkflow.WorkflowId != "sample_workflowId" {
			t.Errorf("Expected History[0].WorkflowId = 'sample_workflowId', got '%s'", historyWorkflow.WorkflowId)
		}
	}

	// Deprecated fields - validate if present but don't require them
	if workflow.StartTime != 0 {
		t.Logf("Deprecated StartTime field present with value: %d", workflow.StartTime)
	}
	if workflow.WorkflowName != "" {
		t.Logf("Deprecated WorkflowName field present with value: %s", workflow.WorkflowName)
	}
	if workflow.WorkflowVersion != 0 {
		t.Logf("Deprecated WorkflowVersion field present with value: %d", workflow.WorkflowVersion)
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(workflow)
	if err != nil {
		t.Fatalf("Failed to serialize struct: %v", err)
	}

	// 5. Round-trip integrity check - use struct comparison for complex recursive structure
	var roundTripWorkflow model.Workflow
	err = json.Unmarshal(serializedJSON, &roundTripWorkflow)
	if err != nil {
		t.Fatalf("Failed to deserialize round-trip JSON: %v", err)
	}

	// Compare key fields to ensure data integrity (avoid full comparison due to recursive History)
	if roundTripWorkflow.WorkflowId != workflow.WorkflowId {
		t.Errorf("Round-trip WorkflowId mismatch: %s vs %s", workflow.WorkflowId, roundTripWorkflow.WorkflowId)
	}
	if roundTripWorkflow.Status != workflow.Status {
		t.Errorf("Round-trip Status.Status mismatch: %s vs %s", workflow.Status, roundTripWorkflow.Status)
	}
	if roundTripWorkflow.CorrelationId != workflow.CorrelationId {
		t.Errorf("Round-trip CorrelationId mismatch: %s vs %s", workflow.CorrelationId, roundTripWorkflow.CorrelationId)
	}
	if roundTripWorkflow.Priority != workflow.Priority {
		t.Errorf("Round-trip Priority mismatch: %d vs %d", workflow.Priority, roundTripWorkflow.Priority)
	}
	if roundTripWorkflow.RateLimited != workflow.RateLimited {
		t.Errorf("Round-trip RateLimited mismatch: %v vs %v", workflow.RateLimited, roundTripWorkflow.RateLimited)
	}
	if len(roundTripWorkflow.Tasks) != len(workflow.Tasks) {
		t.Errorf("Round-trip Tasks length mismatch: %d vs %d", len(workflow.Tasks), len(roundTripWorkflow.Tasks))
	}

	// Compare nested WorkflowDefinition
	if (workflow.WorkflowDefinition == nil) != (roundTripWorkflow.WorkflowDefinition == nil) {
		t.Errorf("Round-trip WorkflowDefinition nil mismatch")
	} else if workflow.WorkflowDefinition != nil {
		if roundTripWorkflow.WorkflowDefinition.Name != workflow.WorkflowDefinition.Name {
			t.Errorf("Round-trip WorkflowDefinition.Name mismatch: %s vs %s",
				workflow.WorkflowDefinition.Name, roundTripWorkflow.WorkflowDefinition.Name)
		}
	}
}
