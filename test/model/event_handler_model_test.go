package model

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserEventHandler(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("EventHandler")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var eventHandler model.EventHandler
	err = json.Unmarshal([]byte(jsonTemplate), &eventHandler)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	if eventHandler.Name != "sample_name" {
		t.Errorf("Expected Name = 'sample_name', got '%s'", eventHandler.Name)
	}
	if eventHandler.Event != "sample_event" {
		t.Errorf("Expected Event = 'sample_event', got '%s'", eventHandler.Event)
	}
	if eventHandler.Condition != "sample_condition" {
		t.Errorf("Expected Condition = 'sample_condition', got '%s'", eventHandler.Condition)
	}
	if eventHandler.EvaluatorType != "sample_evaluatorType" {
		t.Errorf("Expected EvaluatorType = 'sample_evaluatorType', got '%s'", eventHandler.EvaluatorType)
	}
	if !eventHandler.Active {
		t.Errorf("Expected Active = true, got %v", eventHandler.Active)
	}

	// Check slice field - Actions should be populated
	if eventHandler.Actions == nil {
		t.Errorf("Actions slice should not be nil")
	}
	if len(eventHandler.Actions) == 0 {
		t.Errorf("Actions slice should not be empty")
	}
	if len(eventHandler.Actions) != 1 {
		t.Errorf("Expected Actions slice to have 1 element, got %d", len(eventHandler.Actions))
	}

	// Validate the Action element
	action := eventHandler.Actions[0]
	if action.Action != "start_workflow" {
		t.Errorf("Expected Action.Action = 'start_workflow', got '%s'", action.Action)
	}
	if !action.ExpandInlineJSON {
		t.Errorf("Expected Action.ExpandInlineJSON = true, got %v", action.ExpandInlineJSON)
	}

	// Check nested struct pointers - based on template, these should be populated
	if action.StartWorkflow == nil {
		t.Errorf("Action.StartWorkflow should not be nil")
	}
	if action.CompleteTask == nil {
		t.Errorf("Action.CompleteTask should not be nil")
	}
	if action.FailTask == nil {
		t.Errorf("Action.FailTask should not be nil")
	}
	if action.TerminateWorkflow == nil {
		t.Errorf("Action.TerminateWorkflow should not be nil")
	}
	if action.UpdateWorkflowVariables == nil {
		t.Errorf("Action.UpdateWorkflowVariables should not be nil")
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(eventHandler)
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

func TestSerDserTerminateWorkflow(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("EventHandler.TerminateWorkflow")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var terminateWorkflow model.TerminateWorkflow
	err = json.Unmarshal([]byte(jsonTemplate), &terminateWorkflow)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	if terminateWorkflow.WorkflowId != "sample_workflowId" {
		t.Errorf("Expected WorkflowId = 'sample_workflowId', got '%s'", terminateWorkflow.WorkflowId)
	}
	if terminateWorkflow.TerminationReason != "sample_terminationReason" {
		t.Errorf("Expected TerminationReason = 'sample_terminationReason', got '%s'", terminateWorkflow.TerminationReason)
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(terminateWorkflow)
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

func TestSerDserUpdateWorkflowVariables(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("EventHandler.UpdateWorkflowVariables")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var updateWorkflowVariables model.UpdateWorkflowVariables
	err = json.Unmarshal([]byte(jsonTemplate), &updateWorkflowVariables)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	if updateWorkflowVariables.WorkflowId != "sample_workflowId" {
		t.Errorf("Expected WorkflowId = 'sample_workflowId', got '%s'", updateWorkflowVariables.WorkflowId)
	}

	// Check map field
	if updateWorkflowVariables.Variables == nil {
		t.Errorf("Variables map should not be nil")
	}
	if len(updateWorkflowVariables.Variables) == 0 {
		t.Errorf("Variables map should not be empty")
	}
	if updateWorkflowVariables.Variables["key"] != "sample_value" {
		t.Errorf("Expected Variables['key'] = 'sample_value', got %v", updateWorkflowVariables.Variables["key"])
	}

	// Check pointer to bool field
	if updateWorkflowVariables.AppendArray == nil {
		t.Errorf("AppendArray should not be nil")
	}
	if !*updateWorkflowVariables.AppendArray {
		t.Errorf("Expected AppendArray = true, got %v", *updateWorkflowVariables.AppendArray)
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(updateWorkflowVariables)
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

func TestSerDserAction(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("EventHandler.Action")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var action model.Action
	err = json.Unmarshal([]byte(jsonTemplate), &action)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	if action.Action != "start_workflow" {
		t.Errorf("Expected Action = 'start_workflow', got '%s'", action.Action)
	}
	if !action.ExpandInlineJSON {
		t.Errorf("Expected ExpandInlineJSON = true, got %v", action.ExpandInlineJSON)
	}

	// Check nested struct pointers - these should be populated from template references
	if action.StartWorkflow == nil {
		t.Errorf("StartWorkflow should not be nil")
	}
	if action.CompleteTask == nil {
		t.Errorf("CompleteTask should not be nil")
	}
	if action.FailTask == nil {
		t.Errorf("FailTask should not be nil")
	}
	if action.TerminateWorkflow == nil {
		t.Errorf("TerminateWorkflow should not be nil")
	}
	if action.UpdateWorkflowVariables == nil {
		t.Errorf("UpdateWorkflowVariables should not be nil")
	}

	// Validate nested StartWorkflow struct fields
	if action.StartWorkflow.Name != "sample_name" {
		t.Errorf("Expected StartWorkflow.Name = 'sample_name', got '%s'", action.StartWorkflow.Name)
	}
	if action.StartWorkflow.Version != 123 {
		t.Errorf("Expected StartWorkflow.Version = 123, got %d", action.StartWorkflow.Version)
	}
	if action.StartWorkflow.CorrelationId != "sample_correlationId" {
		t.Errorf("Expected StartWorkflow.CorrelationId = 'sample_correlationId', got '%s'", action.StartWorkflow.CorrelationId)
	}
	// Check maps in StartWorkflow
	if action.StartWorkflow.Input == nil || len(action.StartWorkflow.Input) == 0 {
		t.Errorf("StartWorkflow.Input map should not be nil or empty")
	}
	if action.StartWorkflow.TaskToDomain == nil || len(action.StartWorkflow.TaskToDomain) == 0 {
		t.Errorf("StartWorkflow.TaskToDomain map should not be nil or empty")
	}

	// Validate nested TaskDetails struct fields (CompleteTask)
	if action.CompleteTask.WorkflowId != "sample_workflowId" {
		t.Errorf("Expected CompleteTask.WorkflowId = 'sample_workflowId', got '%s'", action.CompleteTask.WorkflowId)
	}
	if action.CompleteTask.TaskRefName != "sample_taskRefName" {
		t.Errorf("Expected CompleteTask.TaskRefName = 'sample_taskRefName', got '%s'", action.CompleteTask.TaskRefName)
	}
	if action.CompleteTask.TaskId != "sample_taskId" {
		t.Errorf("Expected CompleteTask.TaskId = 'sample_taskId', got '%s'", action.CompleteTask.TaskId)
	}
	// Check Output map in TaskDetails
	if action.CompleteTask.Output == nil || len(action.CompleteTask.Output) == 0 {
		t.Errorf("CompleteTask.Output map should not be nil or empty")
	}

	// Validate nested TerminateWorkflow struct fields
	if action.TerminateWorkflow.WorkflowId != "sample_workflowId" {
		t.Errorf("Expected TerminateWorkflow.WorkflowId = 'sample_workflowId', got '%s'", action.TerminateWorkflow.WorkflowId)
	}
	if action.TerminateWorkflow.TerminationReason != "sample_terminationReason" {
		t.Errorf("Expected TerminateWorkflow.TerminationReason = 'sample_terminationReason', got '%s'", action.TerminateWorkflow.TerminationReason)
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(action)
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

func TestSerDserStartWorkflow(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("EventHandler.StartWorkflow")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var startWorkflow model.StartWorkflow
	err = json.Unmarshal([]byte(jsonTemplate), &startWorkflow)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	if startWorkflow.Name != "sample_name" {
		t.Errorf("Expected Name = 'sample_name', got '%s'", startWorkflow.Name)
	}
	if startWorkflow.Version != 123 {
		t.Errorf("Expected Version = 123, got %d", startWorkflow.Version)
	}
	if startWorkflow.CorrelationId != "sample_correlationId" {
		t.Errorf("Expected CorrelationId = 'sample_correlationId', got '%s'", startWorkflow.CorrelationId)
	}

	// Check map fields
	if startWorkflow.Input == nil {
		t.Errorf("Input map should not be nil")
	}
	if len(startWorkflow.Input) == 0 {
		t.Errorf("Input map should not be empty")
	}
	if startWorkflow.Input["key"] != "sample_value" {
		t.Errorf("Expected Input['key'] = 'sample_value', got %v", startWorkflow.Input["key"])
	}

	if startWorkflow.TaskToDomain == nil {
		t.Errorf("TaskToDomain map should not be nil")
	}
	if len(startWorkflow.TaskToDomain) == 0 {
		t.Errorf("TaskToDomain map should not be empty")
	}
	if startWorkflow.TaskToDomain["key"] != "sample_value" {
		t.Errorf("Expected TaskToDomain['key'] = 'sample_value', got %v", startWorkflow.TaskToDomain["key"])
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(startWorkflow)
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

func TestSerDserTaskDetails(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("EventHandler.TaskDetails")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var taskDetails model.TaskDetails
	err = json.Unmarshal([]byte(jsonTemplate), &taskDetails)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	if taskDetails.WorkflowId != "sample_workflowId" {
		t.Errorf("Expected WorkflowId = 'sample_workflowId', got '%s'", taskDetails.WorkflowId)
	}
	if taskDetails.TaskRefName != "sample_taskRefName" {
		t.Errorf("Expected TaskRefName = 'sample_taskRefName', got '%s'", taskDetails.TaskRefName)
	}
	if taskDetails.TaskId != "sample_taskId" {
		t.Errorf("Expected TaskId = 'sample_taskId', got '%s'", taskDetails.TaskId)
	}

	// Check map field
	if taskDetails.Output == nil {
		t.Errorf("Output map should not be nil")
	}
	if len(taskDetails.Output) == 0 {
		t.Errorf("Output map should not be empty")
	}
	if taskDetails.Output["key"] != "sample_value" {
		t.Errorf("Expected Output['key'] = 'sample_value', got %v", taskDetails.Output["key"])
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(taskDetails)
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
