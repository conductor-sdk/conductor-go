package integration_tests

import (
	"context"
	"fmt"
	"github.com/conductor-sdk/conductor-go/internal/testdata"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestWorkflowCreation(t *testing.T) {

	workflow := testdata.NewKitchenSinkWorkflow(testdata.WorkflowExecutor)
	startWorkflowRequest := model.StartWorkflowRequest{
		Name: workflow.GetName(),
	}
	id, err := workflow.StartWorkflow(&startWorkflowRequest)
	assert.NoError(t, err, "Failed to start the workflow", err)
	assert.NotEmpty(t, id, "Workflow Id is null", id)
	fmt.Println("Workflow Id is ", id)
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

	execution, err = executor.GetWorkflow(id, true)
	assert.Error(t, err, "Workflow found even after removing")

	_, err = testdata.MetadataClient.UnregisterWorkflowDef(context.Background(), wf.GetName(), wf.GetVersion())
	assert.NoError(t, err, "Failed to delete workflow definition ", err)

}
