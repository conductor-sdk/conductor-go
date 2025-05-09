package unit_tests

import (
	"encoding/json"
	"fmt"

	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"

	"github.com/stretchr/testify/assert"
)

func TestRetrySettings(t *testing.T) {
	simpleTask := workflow.NewSimpleTask("worker_task", "worker_task_ref")
	simpleTask.RetryPolicy(2, workflow.FixedRetry, 10, 1)
	simpleTask.Input("url", "${workflow.input.url}")
	simpleTask.CacheConfig("${url}", 120)
	wf := workflow.NewConductorWorkflow(nil).
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
	wf := workflow.NewConductorWorkflow(nil).
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

func TestHttpPollTask(t *testing.T) {
	input := workflow.HttpPollInput{
		Method:              "GET",
		Uri:                 "http://localhost:8081/api/hello/intermittent-failures?name=Http_poll_test",
		Accept:              "application/json",
		ContentType:         "application/json",
		PollingInterval:     1,
		PollingStrategy:     "FIXED",
		TerminationCriteria: "(function(){ return $.output.body.length > 10;})();",
		Encode:              true,
	}
	httpPollTask := workflow.NewHttpPollTask("worker_task", &input)
	httpPollTask.RetryPolicy(2, workflow.FixedRetry, 10, 1)
	httpPollTask.Input("url", "${workflow.input.url}")
	httpPollTask.CacheConfig("${url}", 120)
	wf := workflow.NewConductorWorkflow(nil).
		Name("workflow_with_http_poll_task").
		Version(1).
		Add(httpPollTask)
	workflowDef := wf.ToWorkflowDef()
	assert.NotNil(t, workflowDef)
	assert.Equal(t, 1, len(workflowDef.Tasks))
	workflowTask := workflowDef.Tasks[0]
	assert.NotNil(t, workflowTask.TaskDefinition)

	// Check task def
	assert.Equal(t, int32(10), workflowTask.TaskDefinition.RetryDelaySeconds)
	assert.Equal(t, string(workflow.FixedRetry), workflowTask.TaskDefinition.RetryLogic)
	assert.Equal(t, int32(2), workflowTask.TaskDefinition.RetryCount)

	// Check task basic properties
	assert.Equal(t, "HTTP_POLL", workflowTask.Type_)
	assert.Equal(t, "worker_task", workflowTask.TaskReferenceName)
	assert.NotNil(t, workflowTask.InputParameters)

	// Check for the url parameter (added via Input method)
	assert.Equal(t, "${workflow.input.url}", workflowTask.InputParameters["url"])

	// Check http_request parameter exists
	httpRequest, ok := workflowTask.InputParameters["http_request"].(map[string]interface{})
	assert.True(t, ok, "http_request parameter should be a map")
	assert.NotNil(t, httpRequest, "http_request should not be nil")

	// Check all fields in http_request match what we provided
	assert.Equal(t, "GET", httpRequest["method"], "Method should match")
	assert.Equal(t, "http://localhost:8081/api/hello/intermittent-failures?name=Http_poll_test", httpRequest["uri"], "URI should match")
	assert.Equal(t, "application/json", httpRequest["accept"], "Accept should match")
	assert.Equal(t, "application/json", httpRequest["contentType"], "ContentType should match")
	assert.Equal(t, 1, httpRequest["pollingInterval"], "PollingInterval should match")
	assert.Equal(t, "FIXED", httpRequest["pollingStrategy"], "PollingStrategy should match")
	assert.Equal(t, "(function(){ return $.output.body.length > 10;})();", httpRequest["terminationCondition"], "TerminationCondition should match")
	assert.Equal(t, true, httpRequest["encode"], "Encode should match")

	json, _ := json.Marshal(workflowDef)
	fmt.Println(string(json))
}

func TestUpdateTaskWithTaskId(t *testing.T) {

	updateTask := workflow.NewUpdateTaskWithTaskId("update_task_ref", model.CompletedTask, "target_task_to_update")
	updateTask.MergeOutput(true)
	updateTask.TaskOutput(map[string]interface{}{"key": map[string]interface{}{"nestedKey": "nestedValue"}})

	wf := workflow.NewConductorWorkflow(nil).
		Name("workflow_with_update_task").
		Version(1).
		Add(updateTask)
	workflowDef := wf.ToWorkflowDef()

	assert.NotNil(t, workflowDef)
	assert.Equal(t, 1, len(workflowDef.Tasks))

	taskFromWorkflow := workflowDef.Tasks[0]

	assert.Equal(t, "update_task_ref", taskFromWorkflow.TaskReferenceName)
	assert.Equal(t, "target_task_to_update", taskFromWorkflow.InputParameters["taskId"])
	assert.Nil(t, taskFromWorkflow.InputParameters["workflowId"])
	assert.Nil(t, taskFromWorkflow.InputParameters["taskRefName"])
	json, _ := json.Marshal(workflowDef)
	fmt.Println(string(json))
}

func TestUpdateTaskWithWorkflowIdAndTaskRef(t *testing.T) {
	updateTask := workflow.NewUpdateTask("update_task_ref", model.CompletedTask, "target_workflow", "target_task_ref")
	updateTask.MergeOutput(true)
	integers := []int{2, 3, 5, 7, 11, 13}
	updateTask.TaskOutput(map[string]interface{}{"key": integers})
	wf := workflow.NewConductorWorkflow(nil).
		Name("workflow_with_update_task").
		Version(1).
		Add(updateTask)
	workflowDef := wf.ToWorkflowDef()

	assert.NotNil(t, workflowDef)
	assert.Equal(t, 1, len(workflowDef.Tasks))

	taskFromWorkflow := workflowDef.Tasks[0]

	assert.Equal(t, "update_task_ref", taskFromWorkflow.TaskReferenceName)
	assert.Equal(t, "target_workflow", taskFromWorkflow.InputParameters["workflowId"])
	assert.Equal(t, "target_task_ref", taskFromWorkflow.InputParameters["taskRefName"])
	assert.Nil(t, taskFromWorkflow.InputParameters["taskId"])

	json, _ := json.Marshal(workflowDef)
	fmt.Println(string(json))
}
