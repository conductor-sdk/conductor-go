package unit_tests

import (
	"encoding/json"
	"fmt"

	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"github.com/stretchr/testify/assert"
)

func TestRetrySettings(t *testing.T) {
	simpleTask := workflow.NewSimpleTask("worker_task", "worker_task_ref")
	simpleTask.RetryPolicy(2, workflow.FixedRetry, 10, 1)
	simpleTask.Input("url", "${workflow.input.url}")
	simpleTask.CacheConfig("${url}", 120)
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("workflow_with_task_retries").
		Version(1).
		Add(simpleTask)
	workflowDef := wf.ToWorkflowDef()
	assert.NotNil(t, workflowDef)
	assert.Equal(t, 1, len(workflowDef.Tasks))
	workflowTask := workflowDef.Tasks[0]
	assert.NotNil(t, workflowTask.TaskDefinition)
	assert.Equal(t, int32(10), workflowTask.TaskDefinition.RetryDelaySeconds)
	assert.Equal(t, string(workflow.FixedRetry), workflowTask.TaskDefinition.RetryLogic)
	json, _ := json.Marshal(workflowDef)
	fmt.Println(string(json))
}

func TestHttpTask(t *testing.T) {
	input := workflow.HttpInput{
		Method: "GET",
		Uri:    "https://orkes-api-tester.orkesconductor.coma/api",
	}
	httpTask := workflow.NewHttpTask("worker_task", &input)
	httpTask.RetryPolicy(2, workflow.FixedRetry, 10, 1)
	httpTask.Input("url", "${workflow.input.url}")
	httpTask.CacheConfig("${url}", 120)
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("workflow_with_http_task_retries").
		Version(1).
		Add(httpTask)
	workflowDef := wf.ToWorkflowDef()
	assert.NotNil(t, workflowDef)
	assert.Equal(t, 1, len(workflowDef.Tasks))
	workflowTask := workflowDef.Tasks[0]
	assert.NotNil(t, workflowTask.TaskDefinition)
	assert.Equal(t, int32(10), workflowTask.TaskDefinition.RetryDelaySeconds)
	assert.Equal(t, string(workflow.FixedRetry), workflowTask.TaskDefinition.RetryLogic)
	json, _ := json.Marshal(workflowDef)
	fmt.Println(string(json))
}

func TestUpdateTaskWithTaskId(t *testing.T) {

	updateTask :=  workflow.NewUpdateTaskWithTaskId("update_task_ref", model.CompletedTask, "target_task_to_update")
	updateTask.MergeOutput(true)
	updateTask.TaskOutput(map[string]interface{}{"key": map[string]interface{}{"nestedKey": "nestedValue"}})

	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("workflow_with_update_task").
		Version(1).
		Add(updateTask)
	workflowDef := wf.ToWorkflowDef()
	assert.NotNil(t, workflowDef)
	assert.Equal(t, 1, len(workflowDef.Tasks))
	json, _ := json.Marshal(workflowDef)
	fmt.Println(string(json))
}

func TestUpdateTaskWithWorkflowIdAndTaskRef(t *testing.T) {

	updateTask := workflow.NewUpdateTask("update_task_ref", model.CompletedTask, "target_workflow", "target_task_ref")
	updateTask.MergeOutput(true)
	integers := []int{2, 3, 5, 7, 11, 13}
	updateTask.TaskOutput(map[string]interface{}{"key": integers})
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("workflow_with_update_task").
		Version(1).
		Add(updateTask)
	workflowDef := wf.ToWorkflowDef()
	assert.NotNil(t, workflowDef)
	assert.Equal(t, 1, len(workflowDef.Tasks))
	json, _ := json.Marshal(workflowDef)
	fmt.Println(string(json))
}
