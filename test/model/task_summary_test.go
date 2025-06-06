package model

import (
	"encoding/json"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserTaskSummary(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("TaskSummary")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var taskSummary model.TaskSummary
	err = json.Unmarshal([]byte(jsonTemplate), &taskSummary)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// String fields - Workflow context
	if taskSummary.WorkflowId != "sample_workflowId" {
		t.Errorf("Expected WorkflowId = 'sample_workflowId', got '%s'", taskSummary.WorkflowId)
	}
	if taskSummary.WorkflowType != "sample_workflowType" {
		t.Errorf("Expected WorkflowType = 'sample_workflowType', got '%s'", taskSummary.WorkflowType)
	}
	if taskSummary.CorrelationId != "sample_correlationId" {
		t.Errorf("Expected CorrelationId = 'sample_correlationId', got '%s'", taskSummary.CorrelationId)
	}

	// String fields - Time stamps (stored as strings)
	if taskSummary.ScheduledTime != "sample_scheduledTime" {
		t.Errorf("Expected ScheduledTime = 'sample_scheduledTime', got '%s'", taskSummary.ScheduledTime)
	}
	if taskSummary.StartTime != "sample_startTime" {
		t.Errorf("Expected StartTime = 'sample_startTime', got '%s'", taskSummary.StartTime)
	}
	if taskSummary.UpdateTime != "sample_updateTime" {
		t.Errorf("Expected UpdateTime = 'sample_updateTime', got '%s'", taskSummary.UpdateTime)
	}
	if taskSummary.EndTime != "sample_endTime" {
		t.Errorf("Expected EndTime = 'sample_endTime', got '%s'", taskSummary.EndTime)
	}

	// String fields - Task context
	if taskSummary.Status != "IN_PROGRESS" {
		t.Errorf("Expected Status = 'IN_PROGRESS', got '%s'", taskSummary.Status)
	}
	if taskSummary.ReasonForIncompletion != "sample_reasonForIncompletion" {
		t.Errorf("Expected ReasonForIncompletion = 'sample_reasonForIncompletion', got '%s'", taskSummary.ReasonForIncompletion)
	}
	if taskSummary.TaskDefName != "sample_taskDefName" {
		t.Errorf("Expected TaskDefName = 'sample_taskDefName', got '%s'", taskSummary.TaskDefName)
	}
	if taskSummary.TaskType != "sample_taskType" {
		t.Errorf("Expected TaskType = 'sample_taskType', got '%s'", taskSummary.TaskType)
	}
	if taskSummary.TaskId != "sample_taskId" {
		t.Errorf("Expected TaskId = 'sample_taskId', got '%s'", taskSummary.TaskId)
	}

	// String fields - Data and storage
	if taskSummary.Input != "sample_input" {
		t.Errorf("Expected Input = 'sample_input', got '%s'", taskSummary.Input)
	}
	if taskSummary.Output != "sample_output" {
		t.Errorf("Expected Output = 'sample_output', got '%s'", taskSummary.Output)
	}
	if taskSummary.ExternalInputPayloadStoragePath != "sample_externalInputPayloadStoragePath" {
		t.Errorf("Expected ExternalInputPayloadStoragePath = 'sample_externalInputPayloadStoragePath', got '%s'", taskSummary.ExternalInputPayloadStoragePath)
	}
	if taskSummary.ExternalOutputPayloadStoragePath != "sample_externalOutputPayloadStoragePath" {
		t.Errorf("Expected ExternalOutputPayloadStoragePath = 'sample_externalOutputPayloadStoragePath', got '%s'", taskSummary.ExternalOutputPayloadStoragePath)
	}
	if taskSummary.Domain != "sample_domain" {
		t.Errorf("Expected Domain = 'sample_domain', got '%s'", taskSummary.Domain)
	}

	// Int64 fields - Timing
	if taskSummary.ExecutionTime != 123 {
		t.Errorf("Expected ExecutionTime = 123, got %d", taskSummary.ExecutionTime)
	}
	if taskSummary.QueueWaitTime != 123 {
		t.Errorf("Expected QueueWaitTime = 123, got %d", taskSummary.QueueWaitTime)
	}

	// Int field - Priority
	if taskSummary.WorkflowPriority != 123 {
		t.Errorf("Expected WorkflowPriority = 123, got %d", taskSummary.WorkflowPriority)
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(taskSummary)
	if err != nil {
		t.Fatalf("Failed to serialize struct: %v", err)
	}

	// 5. Round-trip integrity check
	var roundTripTaskSummary model.TaskSummary
	err = json.Unmarshal(serializedJSON, &roundTripTaskSummary)
	if err != nil {
		t.Fatalf("Failed to deserialize round-trip JSON: %v", err)
	}

	// Compare key fields to ensure data integrity
	if roundTripTaskSummary.WorkflowId != taskSummary.WorkflowId {
		t.Errorf("Round-trip WorkflowId mismatch: %s vs %s",
			taskSummary.WorkflowId, roundTripTaskSummary.WorkflowId)
	}
	if roundTripTaskSummary.TaskId != taskSummary.TaskId {
		t.Errorf("Round-trip TaskId mismatch: %s vs %s",
			taskSummary.TaskId, roundTripTaskSummary.TaskId)
	}
	if roundTripTaskSummary.Status != taskSummary.Status {
		t.Errorf("Round-trip Status mismatch: %s vs %s",
			taskSummary.Status, roundTripTaskSummary.Status)
	}
	if roundTripTaskSummary.TaskType != taskSummary.TaskType {
		t.Errorf("Round-trip TaskType mismatch: %s vs %s",
			taskSummary.TaskType, roundTripTaskSummary.TaskType)
	}
	if roundTripTaskSummary.ExecutionTime != taskSummary.ExecutionTime {
		t.Errorf("Round-trip ExecutionTime mismatch: %d vs %d",
			taskSummary.ExecutionTime, roundTripTaskSummary.ExecutionTime)
	}
	if roundTripTaskSummary.QueueWaitTime != taskSummary.QueueWaitTime {
		t.Errorf("Round-trip QueueWaitTime mismatch: %d vs %d",
			taskSummary.QueueWaitTime, roundTripTaskSummary.QueueWaitTime)
	}
	if roundTripTaskSummary.WorkflowPriority != taskSummary.WorkflowPriority {
		t.Errorf("Round-trip WorkflowPriority mismatch: %d vs %d",
			taskSummary.WorkflowPriority, roundTripTaskSummary.WorkflowPriority)
	}
}
