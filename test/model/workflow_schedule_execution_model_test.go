package model_test

import (
	"encoding/json"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserWorkflowScheduleExecutionModel(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("WorkflowScheduleExecutionModel")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var workflowScheduleExecutionModel model.WorkflowScheduleExecutionModel
	err = json.Unmarshal([]byte(jsonTemplate), &workflowScheduleExecutionModel)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	if workflowScheduleExecutionModel.ExecutionId != "sample_executionId" {
		t.Errorf("Expected ExecutionId = 'sample_executionId', got '%s'", workflowScheduleExecutionModel.ExecutionId)
	}
	if workflowScheduleExecutionModel.ExecutionTime != 123 {
		t.Errorf("Expected ExecutionTime = 123, got %d", workflowScheduleExecutionModel.ExecutionTime)
	}
	if workflowScheduleExecutionModel.Reason != "sample_reason" {
		t.Errorf("Expected Reason = 'sample_reason', got '%s'", workflowScheduleExecutionModel.Reason)
	}
	if workflowScheduleExecutionModel.ScheduleName != "sample_scheduleName" {
		t.Errorf("Expected ScheduleName = 'sample_scheduleName', got '%s'", workflowScheduleExecutionModel.ScheduleName)
	}
	if workflowScheduleExecutionModel.ScheduledTime != 123 {
		t.Errorf("Expected ScheduledTime = 123, got %d", workflowScheduleExecutionModel.ScheduledTime)
	}
	if workflowScheduleExecutionModel.StackTrace != "sample_stackTrace" {
		t.Errorf("Expected StackTrace = 'sample_stackTrace', got '%s'", workflowScheduleExecutionModel.StackTrace)
	}
	if workflowScheduleExecutionModel.StartWorkflowRequest == nil {
		t.Errorf("StartWorkflowRequest should not be nil")
	} else {
		if workflowScheduleExecutionModel.StartWorkflowRequest.Name != "sample_name" {
			t.Errorf("Expected StartWorkflowRequest.Name = 'sample_name', got '%s'", workflowScheduleExecutionModel.StartWorkflowRequest.Name)
		}
	}
	if workflowScheduleExecutionModel.State != "POLLED" {
		t.Errorf("Expected State = 'POLLED', got '%s'", workflowScheduleExecutionModel.State)
	}
	if workflowScheduleExecutionModel.WorkflowId != "sample_workflowId" {
		t.Errorf("Expected WorkflowId = 'sample_workflowId', got '%s'", workflowScheduleExecutionModel.WorkflowId)
	}
	if workflowScheduleExecutionModel.WorkflowName != "sample_workflowName" {
		t.Errorf("Expected WorkflowName = 'sample_workflowName', got '%s'", workflowScheduleExecutionModel.WorkflowName)
	}
	if workflowScheduleExecutionModel.ZoneId != "sample_zoneId" {
		t.Errorf("Expected ZoneId = 'sample_zoneId', got '%s'", workflowScheduleExecutionModel.ZoneId)
	}
	if workflowScheduleExecutionModel.OrgId != "sample_orgId" {
		t.Errorf("Expected OrgId = 'sample_orgId', got '%s'", workflowScheduleExecutionModel.OrgId)
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(workflowScheduleExecutionModel)
	if err != nil {
		t.Fatalf("Failed to serialize struct: %v", err)
	}

	// 5. Round-trip integrity check
	var roundTripStruct model.WorkflowScheduleExecutionModel
	err = json.Unmarshal(serializedJSON, &roundTripStruct)
	if err != nil {
		t.Fatalf("Failed to deserialize round-trip JSON: %v", err)
	}

	// Compare key fields
	if roundTripStruct.ExecutionId != workflowScheduleExecutionModel.ExecutionId {
		t.Errorf("Round-trip ExecutionId mismatch: %s vs %s", workflowScheduleExecutionModel.ExecutionId, roundTripStruct.ExecutionId)
	}
	if roundTripStruct.ScheduleName != workflowScheduleExecutionModel.ScheduleName {
		t.Errorf("Round-trip ScheduleName mismatch: %s vs %s", workflowScheduleExecutionModel.ScheduleName, roundTripStruct.ScheduleName)
	}
	if roundTripStruct.WorkflowId != workflowScheduleExecutionModel.WorkflowId {
		t.Errorf("Round-trip WorkflowId mismatch: %s vs %s", workflowScheduleExecutionModel.WorkflowId, roundTripStruct.WorkflowId)
	}
}
