package model

import (
	"encoding/json"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserTask(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("Task")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var task model.Task
	err = json.Unmarshal([]byte(jsonTemplate), &task)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// String fields
	if task.TaskType != "sample_taskType" {
		t.Errorf("Expected TaskType = 'sample_taskType', got '%s'", task.TaskType)
	}
	if task.ReferenceTaskName != "sample_referenceTaskName" {
		t.Errorf("Expected ReferenceTaskName = 'sample_referenceTaskName', got '%s'", task.ReferenceTaskName)
	}
	if task.CorrelationId != "sample_correlationId" {
		t.Errorf("Expected CorrelationId = 'sample_correlationId', got '%s'", task.CorrelationId)
	}
	if task.TaskDefName != "sample_taskDefName" {
		t.Errorf("Expected TaskDefName = 'sample_taskDefName', got '%s'", task.TaskDefName)
	}
	if task.RetriedTaskId != "sample_retriedTaskId" {
		t.Errorf("Expected RetriedTaskId = 'sample_retriedTaskId', got '%s'", task.RetriedTaskId)
	}
	if task.WorkflowInstanceId != "sample_workflowInstanceId" {
		t.Errorf("Expected WorkflowInstanceId = 'sample_workflowInstanceId', got '%s'", task.WorkflowInstanceId)
	}
	if task.WorkflowType != "sample_workflowType" {
		t.Errorf("Expected WorkflowType = 'sample_workflowType', got '%s'", task.WorkflowType)
	}
	if task.TaskId != "sample_taskId" {
		t.Errorf("Expected TaskId = 'sample_taskId', got '%s'", task.TaskId)
	}
	if task.ReasonForIncompletion != "sample_reasonForIncompletion" {
		t.Errorf("Expected ReasonForIncompletion = 'sample_reasonForIncompletion', got '%s'", task.ReasonForIncompletion)
	}
	if task.WorkerId != "sample_workerId" {
		t.Errorf("Expected WorkerId = 'sample_workerId', got '%s'", task.WorkerId)
	}
	if task.Domain != "sample_domain" {
		t.Errorf("Expected Domain = 'sample_domain', got '%s'", task.Domain)
	}
	if task.ExecutionNameSpace != "sample_executionNameSpace" {
		t.Errorf("Expected ExecutionNameSpace = 'sample_executionNameSpace', got '%s'", task.ExecutionNameSpace)
	}
	if task.IsolationGroupId != "sample_isolationGroupId" {
		t.Errorf("Expected IsolationGroupId = 'sample_isolationGroupId', got '%s'", task.IsolationGroupId)
	}
	if task.SubWorkflowId != "sample_subWorkflowId" {
		t.Errorf("Expected SubWorkflowId = 'sample_subWorkflowId', got '%s'", task.SubWorkflowId)
	}
	if task.ParentTaskId != "sample_parentTaskId" {
		t.Errorf("Expected ParentTaskId = 'sample_parentTaskId', got '%s'", task.ParentTaskId)
	}

	// Enum field
	if task.Status != model.InProgressTask {
		t.Errorf("Expected Status = 'IN_PROGRESS', got '%s'", task.Status)
	}

	// Int32 fields
	if task.RetryCount != 123 {
		t.Errorf("Expected RetryCount = 123, got %d", task.RetryCount)
	}
	if task.Seq != 123 {
		t.Errorf("Expected Seq = 123, got %d", task.Seq)
	}
	if task.PollCount != 123 {
		t.Errorf("Expected PollCount = 123, got %d", task.PollCount)
	}
	if task.StartDelayInSeconds != 123 {
		t.Errorf("Expected StartDelayInSeconds = 123, got %d", task.StartDelayInSeconds)
	}
	if task.RateLimitPerFrequency != 123 {
		t.Errorf("Expected RateLimitPerFrequency = 123, got %d", task.RateLimitPerFrequency)
	}
	if task.RateLimitFrequencyInSeconds != 123 {
		t.Errorf("Expected RateLimitFrequencyInSeconds = 123, got %d", task.RateLimitFrequencyInSeconds)
	}
	if task.WorkflowPriority != 123 {
		t.Errorf("Expected WorkflowPriority = 123, got %d", task.WorkflowPriority)
	}
	if task.Iteration != 123 {
		t.Errorf("Expected Iteration = 123, got %d", task.Iteration)
	}

	// Int64 fields
	if task.ScheduledTime != 123 {
		t.Errorf("Expected ScheduledTime = 123, got %d", task.ScheduledTime)
	}
	if task.StartTime != 123 {
		t.Errorf("Expected StartTime = 123, got %d", task.StartTime)
	}
	if task.EndTime != 123 {
		t.Errorf("Expected EndTime = 123, got %d", task.EndTime)
	}
	if task.UpdateTime != 123 {
		t.Errorf("Expected UpdateTime = 123, got %d", task.UpdateTime)
	}
	if task.ResponseTimeoutSeconds != 123 {
		t.Errorf("Expected ResponseTimeoutSeconds = 123, got %d", task.ResponseTimeoutSeconds)
	}
	if task.CallbackAfterSeconds != 123 {
		t.Errorf("Expected CallbackAfterSeconds = 123, got %d", task.CallbackAfterSeconds)
	}
	if task.FirstStartTime != 123 {
		t.Errorf("Expected FirstStartTime = 123, got %d", task.FirstStartTime)
	}

	// Boolean fields
	if !task.Retried {
		t.Errorf("Expected Retried = true, got %v", task.Retried)
	}
	if !task.Executed {
		t.Errorf("Expected Executed = true, got %v", task.Executed)
	}
	if !task.CallbackFromWorker {
		t.Errorf("Expected CallbackFromWorker = true, got %v", task.CallbackFromWorker)
	}
	if !task.SubworkflowChanged {
		t.Errorf("Expected SubworkflowChanged = true, got %v", task.SubworkflowChanged)
	}

	// Map fields
	if task.InputData == nil {
		t.Errorf("InputData map should not be nil")
	}
	if len(task.InputData) == 0 {
		t.Errorf("InputData map should not be empty")
	}
	if task.InputData["sample_key"] != "sample_value" {
		t.Errorf("Expected InputData['sample_key'] = 'sample_value', got %v", task.InputData["sample_key"])
	}

	if task.OutputData == nil {
		t.Errorf("OutputData map should not be nil")
	}
	if len(task.OutputData) == 0 {
		t.Errorf("OutputData map should not be empty")
	}
	if task.OutputData["sample_key"] != "sample_value" {
		t.Errorf("Expected OutputData['sample_key'] = 'sample_value', got %v", task.OutputData["sample_key"])
	}

	// Pointer to struct field
	if task.WorkflowTask == nil {
		t.Errorf("WorkflowTask should not be nil")
	} else {
		// Validate nested WorkflowTask fields
		if task.WorkflowTask.Name != "sample_name" {
			t.Errorf("Expected WorkflowTask.Name = 'sample_name', got '%s'", task.WorkflowTask.Name)
		}
		if task.WorkflowTask.TaskReferenceName != "sample_taskReferenceName" {
			t.Errorf("Expected WorkflowTask.TaskReferenceName = 'sample_taskReferenceName', got '%s'", task.WorkflowTask.TaskReferenceName)
		}
	}

	// Deprecated fields - validate if present but don't require them
	if task.QueueWaitTime != 0 {
		t.Logf("Deprecated QueueWaitTime field present with value: %d", task.QueueWaitTime)
	}
	if task.TaskDefinition != nil {
		t.Logf("Deprecated TaskDefinition field present")
	}
	if task.LoopOverTask {
		t.Logf("Deprecated LoopOverTask field present with value: %v", task.LoopOverTask)
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(task)
	if err != nil {
		t.Fatalf("Failed to serialize struct: %v", err)
	}

	// 5. Round-trip integrity check
	var roundTripTask model.Task
	err = json.Unmarshal(serializedJSON, &roundTripTask)
	if err != nil {
		t.Fatalf("Failed to deserialize round-trip JSON: %v", err)
	}

	// Compare key fields to ensure data integrity
	if roundTripTask.TaskId != task.TaskId {
		t.Errorf("Round-trip TaskId mismatch: %s vs %s", task.TaskId, roundTripTask.TaskId)
	}
	if roundTripTask.TaskType != task.TaskType {
		t.Errorf("Round-trip TaskType mismatch: %s vs %s", task.TaskType, roundTripTask.TaskType)
	}
	if roundTripTask.Status != task.Status {
		t.Errorf("Round-trip Status mismatch: %s vs %s", task.Status, roundTripTask.Status)
	}
	if roundTripTask.WorkflowInstanceId != task.WorkflowInstanceId {
		t.Errorf("Round-trip WorkflowInstanceId mismatch: %s vs %s", task.WorkflowInstanceId, roundTripTask.WorkflowInstanceId)
	}
	if roundTripTask.RetryCount != task.RetryCount {
		t.Errorf("Round-trip RetryCount mismatch: %d vs %d", task.RetryCount, roundTripTask.RetryCount)
	}
	if roundTripTask.StartTime != task.StartTime {
		t.Errorf("Round-trip StartTime mismatch: %d vs %d", task.StartTime, roundTripTask.StartTime)
	}
	if roundTripTask.Executed != task.Executed {
		t.Errorf("Round-trip Executed mismatch: %v vs %v", task.Executed, roundTripTask.Executed)
	}

	// Compare nested WorkflowTask
	if (task.WorkflowTask == nil) != (roundTripTask.WorkflowTask == nil) {
		t.Errorf("Round-trip WorkflowTask nil mismatch")
	} else if task.WorkflowTask != nil {
		if roundTripTask.WorkflowTask.Name != task.WorkflowTask.Name {
			t.Errorf("Round-trip WorkflowTask.Name mismatch: %s vs %s",
				task.WorkflowTask.Name, roundTripTask.WorkflowTask.Name)
		}
	}
}
