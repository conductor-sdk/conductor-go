package integration_tests

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/internal/testdata"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/stretchr/testify/assert"
)

func TestWorkflowCreation(t *testing.T) {
	workflow := testdata.NewKitchenSinkWorkflow(testdata.WorkflowExecutor)
	err := workflow.Register(true)
	if err != nil {
		t.Fatalf("Failed to register workflow: %s, reason: %s", workflow.GetName(), err.Error())
	}
	startWorkflowRequest := &model.StartWorkflowRequest{
		Name: workflow.GetName(),
	}
	id, err := workflow.StartWorkflow(startWorkflowRequest)
	if err != nil {
		t.Fatalf("Failed to start the workflow, reason: %s", err)
	}
	assert.NotEmpty(t, id, "Workflow Id is null", id)
}

func TestRemoveWorkflow(t *testing.T) {
	executor := testdata.WorkflowExecutor
	wf := workflow.NewConductorWorkflow(executor)
	wf.Name("temp_wf_" + strconv.Itoa(time.Now().Nanosecond())).Version(1)
	wf = wf.Add(workflow.NewSetVariableTask("set_var").Input("var_value", 42))
	err := wf.Register(true)

	assert.NoError(t, err, "Failed to register workflow")

	id, err := executor.StartWorkflow(&model.StartWorkflowRequest{Name: wf.GetName()})
	assert.NoError(t, err, "Failed to start workflow")
	fmt.Print("Id of the workflow, ", id)

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

func TestExecuteWorkflow(t *testing.T) {
	executor := testdata.WorkflowExecutor
	wf := workflow.NewConductorWorkflow(executor)
	wf.Name("temp_wf_2_" + strconv.Itoa(time.Now().Nanosecond())).Version(1)
	wf = wf.Add(workflow.NewSetVariableTask("set_var").Input("var_value", 42))
	wf.OutputParameters(map[string]interface{}{
		"param1": "Test",
		"param2": 123,
	})
	err := wf.Register(true)

	assert.NoError(t, err, "Failed to register workflow")
	version := wf.GetVersion()
	run, err := executor.ExecuteWorkflow(&model.StartWorkflowRequest{Name: wf.GetName(), Version: &version}, "")
	assert.NoError(t, err, "Failed to start workflow")
	fmt.Print("Id of the workflow, ", run.WorkflowId)
	assert.Equal(t, string(model.CompletedWorkflow), run.Status)

	execution, err := executor.GetWorkflow(run.WorkflowId, true)
	assert.NoError(t, err, "Failed to get workflow execution")
	assert.Equal(t, model.CompletedWorkflow, execution.Status, "Workflow is not in the completed state")

	_, err = testdata.MetadataClient.UnregisterWorkflowDef(
		context.Background(),
		wf.GetName(),
		wf.GetVersion(),
	)
	assert.NoError(t, err, "Failed to delete workflow definition ", err)
}
