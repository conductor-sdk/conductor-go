package integration_tests

import (
	"context"
	"fmt"
	"github.com/google/uuid"
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
	//Start all the workers
	taskRunner := testdata.TaskRunner
	taskRunner.StartWorker("simple_task", testdata.SimpleWorker, 1, time.Millisecond)
	taskRunner.StartWorker("simple_task_5", testdata.SimpleWorker, 1, time.Millisecond)
	taskRunner.StartWorker("simple_task_3", testdata.SimpleWorker, 1, time.Millisecond)
	taskRunner.StartWorker("simple_task_1", testdata.SimpleWorker, 1, time.Millisecond)
	taskRunner.StartWorker("dynamic_fork_prep", testdata.DynamicForkWorker, 1, time.Millisecond)

	run, err := workflow.ExecuteWorkflowWithInput(map[string]interface{}{
		"key1": "input1",
		"key2": 101,
	}, "")
	if err != nil {
		t.Fatalf("Failed to complete the workflow, reason: %s", err)
	}
	assert.NotEmpty(t, run, "Workflow is null", run)
	assert.Equal(t, string(model.CompletedWorkflow), run.Status)
	assert.Equal(t, "input1", run.Input["key1"])
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

func TestExecuteWorkflowSync(t *testing.T) {
	executor := testdata.WorkflowExecutor
	wf := workflow.NewConductorWorkflow(executor)
	wf.Name("temp_wf_3_" + strconv.Itoa(time.Now().Nanosecond())).Version(1)
	wf = wf.Add(workflow.NewSetVariableTask("set_var").Input("var_value", 42))
	wf.OutputParameters(map[string]interface{}{
		"param1": "Test",
		"param2": 123,
	})
	err := wf.Register(true)

	assert.NoError(t, err, "Failed to register workflow")
	run, err := wf.ExecuteWorkflowWithInput(map[string]interface{}{}, "")
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

func TestExecuteWorkflowWithCorrelationIds(t *testing.T) {
	executor := testdata.WorkflowExecutor
	correlationId1 := "correlationId1-" + uuid.New().String()
	correlationId2 := "correlationId2-" + uuid.New().String()
	httpTaskWorkflow1 := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_HTTP" + correlationId1).
		OwnerEmail("test@orkes.io").
		Version(1).
		Add(httpTask)
	httpTaskWorkflow2 := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_HTTP" + correlationId2).
		OwnerEmail("test@orkes.io").
		Version(1).
		Add(httpTask)

	_, err := httpTaskWorkflow1.StartWorkflow(&model.StartWorkflowRequest{CorrelationId: correlationId1})
	if err != nil {
		t.Fatal(err)
	}
	_, err = httpTaskWorkflow2.StartWorkflow(&model.StartWorkflowRequest{CorrelationId: correlationId2})
	if err != nil {
		t.Fatal(err)
	}
	// wait a bit until indexed, no need to wait until completion
	time.Sleep(5 * time.Second)

	workflows, err := executor.GetByCorrelationIdsAndNames(true, true,
		[]string{correlationId1, correlationId2}, []string{httpTaskWorkflow1.GetName(), httpTaskWorkflow2.GetName()})
	if err != nil {
		t.Fatal(err)
	}
	assert.Contains(t, workflows, correlationId1)
	assert.Contains(t, workflows, correlationId2)
	assert.NotEmpty(t, workflows[correlationId1])
	assert.NotEmpty(t, workflows[correlationId2])
	assert.Equal(t, workflows[correlationId1][0].CorrelationId, correlationId1)
	assert.Equal(t, workflows[correlationId2][0].CorrelationId, correlationId2)
}
