package model

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserDynamicForkJoinTaskList(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("DynamicForkJoinTaskList")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var dynamicForkJoinTaskList model.DynamicForkJoinTaskList
	err = json.Unmarshal([]byte(jsonTemplate), &dynamicForkJoinTaskList)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// Check slice field
	if dynamicForkJoinTaskList.DynamicTasks == nil {
		t.Errorf("DynamicTasks slice should not be nil")
	}
	if len(dynamicForkJoinTaskList.DynamicTasks) == 0 {
		t.Errorf("DynamicTasks slice should not be empty")
	}
	if len(dynamicForkJoinTaskList.DynamicTasks) != 1 {
		t.Errorf("Expected DynamicTasks slice to have 1 element, got %d", len(dynamicForkJoinTaskList.DynamicTasks))
	}

	// Validate the DynamicForkJoinTask element
	task := dynamicForkJoinTaskList.DynamicTasks[0]
	if task.TaskName != "sample_taskName" {
		t.Errorf("Expected TaskName = 'sample_taskName', got '%s'", task.TaskName)
	}
	if task.WorkflowName != "sample_workflowName" {
		t.Errorf("Expected WorkflowName = 'sample_workflowName', got '%s'", task.WorkflowName)
	}
	if task.Type != "SIMPLE" {
		t.Errorf("Expected Type = 'SIMPLE', got '%s'", task.Type)
	}
	if task.ReferenceName != "sample_referenceName" {
		t.Errorf("Expected ReferenceName = 'sample_referenceName', got '%s'", task.ReferenceName)
	}

	// Check map field in the task
	if task.Input == nil {
		t.Errorf("Input map should not be nil")
	}
	if len(task.Input) == 0 {
		t.Errorf("Input map should not be empty")
	}
	if task.Input["sample_key"] != "sample_value" {
		t.Errorf("Expected Input['sample_key'] = 'sample_value', got %v", task.Input["sample_key"])
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(dynamicForkJoinTaskList)
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
