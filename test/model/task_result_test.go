package model

import (
	"encoding/json"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserTaskResult(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("TaskResult")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var taskResult model.TaskResult
	err = json.Unmarshal([]byte(jsonTemplate), &taskResult)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// String fields
	if taskResult.WorkflowInstanceId != "sample_workflowInstanceId" {
		t.Errorf("Expected WorkflowInstanceId = 'sample_workflowInstanceId', got '%s'", taskResult.WorkflowInstanceId)
	}
	if taskResult.TaskId != "sample_taskId" {
		t.Errorf("Expected TaskId = 'sample_taskId', got '%s'", taskResult.TaskId)
	}
	if taskResult.ReasonForIncompletion != "sample_reasonForIncompletion" {
		t.Errorf("Expected ReasonForIncompletion = 'sample_reasonForIncompletion', got '%s'", taskResult.ReasonForIncompletion)
	}
	if taskResult.WorkerId != "sample_workerId" {
		t.Errorf("Expected WorkerId = 'sample_workerId', got '%s'", taskResult.WorkerId)
	}
	if taskResult.ExternalOutputPayloadStoragePath != "sample_externalOutputPayloadStoragePath" {
		t.Errorf("Expected ExternalOutputPayloadStoragePath = 'sample_externalOutputPayloadStoragePath', got '%s'", taskResult.ExternalOutputPayloadStoragePath)
	}
	if taskResult.SubWorkflowId != "sample_subWorkflowId" {
		t.Errorf("Expected SubWorkflowId = 'sample_subWorkflowId', got '%s'", taskResult.SubWorkflowId)
	}

	// Enum field
	if taskResult.Status != model.InProgressTask {
		t.Errorf("Expected Status = 'IN_PROGRESS', got '%s'", taskResult.Status)
	}

	// Integer fields (assuming these exist based on typical TaskResult structure)
	if taskResult.CallbackAfterSeconds != 123 {
		t.Errorf("Expected CallbackAfterSeconds = 123, got %d", taskResult.CallbackAfterSeconds)
	}

	// Boolean field
	if !taskResult.ExtendLease {
		t.Errorf("Expected ExtendLease = true, got %v", taskResult.ExtendLease)
	}

	// Map field
	if taskResult.OutputData == nil {
		t.Errorf("OutputData map should not be nil")
	}
	if len(taskResult.OutputData) == 0 {
		t.Errorf("OutputData map should not be empty")
	}
	if taskResult.OutputData["sample_key"] != "sample_value" {
		t.Errorf("Expected OutputData['sample_key'] = 'sample_value', got %v", taskResult.OutputData["sample_key"])
	}

	// Slice field - TaskExecLog slice
	if taskResult.Logs == nil {
		t.Errorf("Logs slice should not be nil")
	}
	if len(taskResult.Logs) == 0 {
		t.Errorf("Logs slice should not be empty")
	}
	if len(taskResult.Logs) != 1 {
		t.Errorf("Expected Logs slice to have 1 element, got %d", len(taskResult.Logs))
	}

	// Validate TaskExecLog element
	log := taskResult.Logs[0]
	if log.TaskId != "sample_taskId" {
		t.Errorf("Expected Log.TaskId = 'sample_taskId', got '%s'", log.TaskId)
	}
	if log.Log != "sample_log" {
		t.Errorf("Expected Log.Log = 'sample_log', got '%s'", log.Log)
	}
	if log.CreatedTime != 123 {
		t.Errorf("Expected Log.CreatedTime = 123, got %d", log.CreatedTime)
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(taskResult)
	if err != nil {
		t.Fatalf("Failed to serialize struct: %v", err)
	}

	// 5. Round-trip integrity check
	var roundTripTaskResult model.TaskResult
	err = json.Unmarshal(serializedJSON, &roundTripTaskResult)
	if err != nil {
		t.Fatalf("Failed to deserialize round-trip JSON: %v", err)
	}

	// Compare key fields to ensure data integrity
	if roundTripTaskResult.WorkflowInstanceId != taskResult.WorkflowInstanceId {
		t.Errorf("Round-trip WorkflowInstanceId mismatch: %s vs %s",
			taskResult.WorkflowInstanceId, roundTripTaskResult.WorkflowInstanceId)
	}
	if roundTripTaskResult.TaskId != taskResult.TaskId {
		t.Errorf("Round-trip TaskId mismatch: %s vs %s",
			taskResult.TaskId, roundTripTaskResult.TaskId)
	}
	if roundTripTaskResult.Status != taskResult.Status {
		t.Errorf("Round-trip Status mismatch: %s vs %s",
			taskResult.Status, roundTripTaskResult.Status)
	}
	if roundTripTaskResult.WorkerId != taskResult.WorkerId {
		t.Errorf("Round-trip WorkerId mismatch: %s vs %s",
			taskResult.WorkerId, roundTripTaskResult.WorkerId)
	}
	if roundTripTaskResult.ExtendLease != taskResult.ExtendLease {
		t.Errorf("Round-trip ExtendLease mismatch: %v vs %v",
			taskResult.ExtendLease, roundTripTaskResult.ExtendLease)
	}
	if len(roundTripTaskResult.Logs) != len(taskResult.Logs) {
		t.Errorf("Round-trip Logs length mismatch: %d vs %d",
			len(taskResult.Logs), len(roundTripTaskResult.Logs))
	}
}
