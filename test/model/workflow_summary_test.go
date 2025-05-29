package model_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserWorkflowSummary(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("WorkflowSummary")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var workflowSummary model.WorkflowSummary
	err = json.Unmarshal([]byte(jsonTemplate), &workflowSummary)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	if workflowSummary.WorkflowType != "sample_workflowType" {
		t.Errorf("Expected WorkflowType = 'sample_workflowType', got '%s'", workflowSummary.WorkflowType)
	}
	if workflowSummary.Version != 123 {
		t.Errorf("Expected Version = 123, got %d", workflowSummary.Version)
	}
	if workflowSummary.WorkflowId != "sample_workflowId" {
		t.Errorf("Expected WorkflowId = 'sample_workflowId', got '%s'", workflowSummary.WorkflowId)
	}
	if workflowSummary.CorrelationId != "sample_correlationId" {
		t.Errorf("Expected CorrelationId = 'sample_correlationId', got '%s'", workflowSummary.CorrelationId)
	}
	if workflowSummary.StartTime != "sample_startTime" {
		t.Errorf("Expected StartTime = 'sample_startTime', got '%s'", workflowSummary.StartTime)
	}
	if workflowSummary.UpdateTime != "sample_updateTime" {
		t.Errorf("Expected UpdateTime = 'sample_updateTime', got '%s'", workflowSummary.UpdateTime)
	}
	if workflowSummary.EndTime != "sample_endTime" {
		t.Errorf("Expected EndTime = 'sample_endTime', got '%s'", workflowSummary.EndTime)
	}
	if workflowSummary.Status != string(model.RunningWorkflow) {
		t.Errorf("Expected Status = 'RUNNING', got '%s'", workflowSummary.Status)
	}
	if workflowSummary.Input != "sample_input" {
		t.Errorf("Expected Input = 'sample_input', got '%s'", workflowSummary.Input)
	}
	if workflowSummary.Output != "sample_output" {
		t.Errorf("Expected Output = 'sample_output', got '%s'", workflowSummary.Output)
	}
	if workflowSummary.ReasonForIncompletion != "sample_reasonForIncompletion" {
		t.Errorf("Expected ReasonForIncompletion = 'sample_reasonForIncompletion', got '%s'", workflowSummary.ReasonForIncompletion)
	}
	if workflowSummary.ExecutionTime != 123 {
		t.Errorf("Expected ExecutionTime = 123, got %d", workflowSummary.ExecutionTime)
	}
	if workflowSummary.Event != "sample_event" {
		t.Errorf("Expected Event = 'sample_event', got '%s'", workflowSummary.Event)
	}
	if workflowSummary.FailedReferenceTaskNames != "sample_failedReferenceTaskNames" {
		t.Errorf("Expected FailedReferenceTaskNames = 'sample_failedReferenceTaskNames', got '%s'", workflowSummary.FailedReferenceTaskNames)
	}
	if workflowSummary.ExternalInputPayloadStoragePath != "sample_externalInputPayloadStoragePath" {
		t.Errorf("Expected ExternalInputPayloadStoragePath = 'sample_externalInputPayloadStoragePath', got '%s'", workflowSummary.ExternalInputPayloadStoragePath)
	}
	if workflowSummary.ExternalOutputPayloadStoragePath != "sample_externalOutputPayloadStoragePath" {
		t.Errorf("Expected ExternalOutputPayloadStoragePath = 'sample_externalOutputPayloadStoragePath', got '%s'", workflowSummary.ExternalOutputPayloadStoragePath)
	}
	if workflowSummary.Priority != 123 {
		t.Errorf("Expected Priority = 123, got %d", workflowSummary.Priority)
	}
	if workflowSummary.CreatedBy != "sample_createdBy" {
		t.Errorf("Expected CreatedBy = 'sample_createdBy', got '%s'", workflowSummary.CreatedBy)
	}
	if workflowSummary.OutputSize != 0 {
		t.Errorf("Expected OutputSize = 0, got %d", workflowSummary.OutputSize)
	}
	if workflowSummary.InputSize != 0 {
		t.Errorf("Expected InputSize = 0, got %d", workflowSummary.InputSize)
	}

	// Validate slice fields
	if workflowSummary.FailedTaskNames == nil {
		t.Errorf("FailedTaskNames should not be nil")
	}
	if len(workflowSummary.FailedTaskNames) == 0 {
		t.Errorf("FailedTaskNames should not be empty")
	}
	if len(workflowSummary.FailedTaskNames) != 1 {
		t.Errorf("Expected FailedTaskNames to have 1 element, got %d", len(workflowSummary.FailedTaskNames))
	}
	if workflowSummary.FailedTaskNames[0] != "sample_failedTaskNames" {
		t.Errorf("Expected FailedTaskNames[0] = 'sample_failedTaskNames', got '%s'", workflowSummary.FailedTaskNames[0])
	}

	assert.Nil(t, workflowSummary.TaskToDomain)

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(workflowSummary)
	if err != nil {
		t.Fatalf("Failed to serialize struct: %v", err)
	}

	// 5. Round-trip integrity check
	var roundTripStruct model.WorkflowSummary
	err = json.Unmarshal(serializedJSON, &roundTripStruct)
	if err != nil {
		t.Fatalf("Failed to deserialize round-trip JSON: %v", err)
	}

	// Compare key fields
	if roundTripStruct.WorkflowType != workflowSummary.WorkflowType {
		t.Errorf("Round-trip WorkflowType mismatch")
	}
}
