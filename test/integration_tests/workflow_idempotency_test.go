package integration_tests

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"github.com/stretchr/testify/assert"
)

func TestIdempotencyCombinations(t *testing.T) {
	executor := testdata.WorkflowExecutor
	wf := workflow.NewConductorWorkflow(executor)
	wf.Name("temp_wf_" + strconv.Itoa(time.Now().Nanosecond())).Version(1)
	wf = wf.Add(workflow.NewSetVariableTask("set_var").Input("var_value", 42))
	err := wf.Register(true)

	assert.NoError(t, err, "Failed to register workflow")

	id, olderr := executor.StartWorkflow(&model.StartWorkflowRequest{Name: wf.GetName(), IdempotencyKey: "test", IdempotencyStrategy: model.FailOnConflict})
	assert.NoError(t, olderr, "Failed to start workflow")

	id2, err := executor.StartWorkflow(&model.StartWorkflowRequest{Name: wf.GetName(), IdempotencyKey: "test", IdempotencyStrategy: model.ReturnExisting})
	assert.NoError(t, err, "Failed to start workflow")
	assert.Equal(t, id, id2) //should return an existing workflow

	_, err = executor.StartWorkflow(&model.StartWorkflowRequest{Name: wf.GetName(), IdempotencyKey: "test", IdempotencyStrategy: model.FailOnConflict})
	assert.Error(t, err, "Failed to start workflow")

	execution, err := executor.GetWorkflow(id, true)
	assert.NoError(t, err, "Failed to get workflow execution")
	assert.Equal(t, model.CompletedWorkflow, execution.Status, "Workflow is not in the completed state")

	err = executor.RemoveWorkflow(id)
	assert.NoError(t, err, "Failed to remove workflow execution")

	_, err = executor.GetWorkflow(id, true)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no such workflow by Id")

	_, err = testdata.MetadataClient.UnregisterWorkflowDef(
		context.Background(),
		wf.GetName(),
		wf.GetVersion(),
	)
	assert.NoError(t, err, "Failed to delete workflow definition ", err)
}

func TestIdempotencyFailOnRunning(t *testing.T) {
	executor := testdata.WorkflowExecutor
	wf := workflow.NewConductorWorkflow(executor)
	wf.Name("temp_wf_" + strconv.Itoa(time.Now().Nanosecond())).Version(1)
	wf = wf.Add(workflow.NewSimpleTask("simple_task_1", "simple_task_1"))
	err := wf.Register(true)
	assert.NoError(t, err, "Failed to register workflow")

	// (1) workflow should start
	id, err := executor.StartWorkflow(&model.StartWorkflowRequest{Name: wf.GetName(), IdempotencyKey: "test", IdempotencyStrategy: model.FailOnRunning})
	assert.NoError(t, err, "Failed to start workflow")

	// (2) workflow start should fail because (1) is running
	_, err = executor.StartWorkflow(&model.StartWorkflowRequest{Name: wf.GetName(), IdempotencyKey: "test", IdempotencyStrategy: model.FailOnRunning})
	assert.Error(t, err, "Workflow should have failed but there was no error")

	// complete task so that workflow is completed
	err = executor.UpdateTaskByRefName("simple_task_1", id, model.CompletedTask, map[string]interface{}{})
	assert.NoError(t, err, "Failed to update task")

	checkWorkflowIsCompleted(t, executor, id)

	//  workflow should start
	id2, err := executor.StartWorkflow(&model.StartWorkflowRequest{Name: wf.GetName(), IdempotencyKey: "test", IdempotencyStrategy: model.FailOnRunning})
	assert.NoError(t, err, "Failed to start workflow")
	assert.NotEqual(t, id, id2)
}

func checkWorkflowIsCompleted(t *testing.T, executor *executor.WorkflowExecutor, id string) {
	timeout := time.After(5 * time.Second)
	tick := time.Tick(1 * time.Second)

	for {
		select {
		case <-timeout:
			t.Fatalf("Timed out and workflow %s didn't complete", id)
		case <-tick:
			wf, err := executor.GetWorkflow(id, false)
			assert.NoError(t, err)
			assert.Equal(t, model.CompletedWorkflow, wf.Status)
			return
		}
	}
}
