package integration_tests

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/internal/testdata"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
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
