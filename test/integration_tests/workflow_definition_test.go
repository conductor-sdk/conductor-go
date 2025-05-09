package integration_tests

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"github.com/stretchr/testify/assert"
)

const retryLimit = 5

func TestWorkflowCreation(t *testing.T) {
	executor := testdata.WorkflowExecutor
	workflow := testdata.NewKitchenSinkWorkflow(testdata.WorkflowExecutor)
	err := workflow.Register(true)
	if err != nil {
		t.Fatalf("Failed to register workflow: %s, reason: %s", workflow.GetName(), err.Error())
	}
	startWorkers()
	run, err := executeWorkflowWithRetries(workflow, map[string]interface{}{
		"key1": "input1",
		"key2": 101,
	})
	if err != nil {
		t.Fatalf("Failed to complete the workflow, reason: %s", err)
	}

	assert.NotEmpty(t, run, "Workflow is null", run)

	timeout := time.After(60 * time.Second)
	tick := time.Tick(1 * time.Second)
	workflowId := run.WorkflowId
	assert.NoError(t, err)

	for {
		select {
		case <-timeout:
			t.Fatalf("Timed out and workflow %s didn't complete", workflowId)
		case <-tick:
			wf, err := executor.GetWorkflow(workflowId, false)
			assert.NoError(t, err)
			if wf.Status == model.CompletedWorkflow {
				// Success! Verify the workflow details
				assert.Equal(t, model.CompletedWorkflow, wf.Status)
				assert.Equal(t, "input1", run.Input["key1"])
				return
			} else if wf.Status == model.FailedWorkflow || wf.Status == model.TerminatedWorkflow {
				t.Fatalf("Workflow failed with status: %s", wf.Status)
			}
		}
	}
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
	wf := workflow.NewConductorWorkflow(executor).
		Name("temp_wf_2_" + strconv.Itoa(time.Now().Nanosecond())).
		Version(1).
		OwnerEmail("test@orkes.io")
	wf = wf.Add(workflow.NewSetVariableTask("set_var").Input("var_value", 42))
	wf.OutputParameters(map[string]interface{}{
		"param1": "Test",
		"param2": 123,
	})
	err := wf.Register(true)

	assert.NoError(t, err, "Failed to register workflow")
	version := wf.GetVersion()
	run, err := executeWorkflowWithRetriesWithStartWorkflowRequest(
		&model.StartWorkflowRequest{
			Name:    wf.GetName(),
			Version: version,
		},
	)
	assert.NoError(t, err, "Failed to start workflow")
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
		Add(testdata.TestHttpTask)
	httpTaskWorkflow2 := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_HTTP" + correlationId2).
		OwnerEmail("test@orkes.io").
		Version(1).
		Add(testdata.TestHttpTask)
	_, err := httpTaskWorkflow1.StartWorkflow(&model.StartWorkflowRequest{CorrelationId: correlationId1})
	if err != nil {
		t.Fatal(err)
	}
	_, err = httpTaskWorkflow2.StartWorkflow(&model.StartWorkflowRequest{CorrelationId: correlationId2})
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(3 * time.Second)
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

func TestTerminateWorkflowWithFailure(t *testing.T) {

	executor := testdata.WorkflowExecutor
	wf := workflow.NewConductorWorkflow(executor).
		Name("TEST_GO_SET_VAR_USED_AS_FAILURE").
		Version(1).
		Add(workflow.NewSetVariableTask("set_var").Input("var_value", 42))
	err := testdata.ValidateWorkflowRegistration(wf)
	if err != nil {
		t.Fatal(err)
	}

	workflowWait := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_WAIT_CONDUCTOR").
		Version(1).
		Add(workflow.NewWaitTask("termination_wait")).
		FailureWorkflow(wf.GetName())
	err = testdata.ValidateWorkflowRegistration(workflowWait)
	if err != nil {
		t.Fatal(err)
	}

	id, err := workflowWait.StartWorkflow(&model.StartWorkflowRequest{})
	if err != nil {
		t.Fatal(err)
	}
	err = executor.TerminateWithFailure(id, "Terminated to trigger failure workflow", true)
	if err != nil {
		t.Fatal(err)
	}
	terminatedWfStatus, err := executor.GetWorkflow(id, false)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, terminatedWfStatus.Output["conductor.failure_workflow"])
}

func TestWorkflowTest(t *testing.T) {
	httpTaskWorkflow := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_TEST").
		Version(1).
		Add(testdata.TestHttpTask)

	err := testdata.ValidateWorkflowRegistration(httpTaskWorkflow)

	if err != nil {
		t.Fatal(
			"Failed to register workflow. Reason: ", err.Error(),
		)
	}

	taskMocks := make(map[string][]model.TaskMock)
	taskMocks[testdata.TestHttpTask.ReferenceName()] = []model.TaskMock{
		{
			Status: "COMPLETED",
			Output: map[string]interface{}{
				"response": map[string]interface{}{
					"body": map[string]interface{}{
						"testKey": "testValue",
					},
					"statusCode": 200,
				},
			},
			ExecutionTime: 50, // 50 milliseconds
			QueueWaitTime: 5,  // 5 milliseconds
		},
	}
	testRequest := model.WorkflowTestRequest{
		Name:                httpTaskWorkflow.GetName(),
		Version:             httpTaskWorkflow.GetVersion(),
		Input:               map[string]interface{}{"inputParam1": "testValue1"},
		TaskRefToMockOutput: taskMocks,
		WorkflowDef:         httpTaskWorkflow.ToWorkflowDef(),
	}
	workflowResult, resp, err := testdata.WorkflowClient.TestWorkflow(context.Background(), testRequest)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotEmpty(t, workflowResult)
	assert.Equal(t, 1, len(workflowResult.Tasks))
	assert.Equal(t, model.TaskResultStatus("COMPLETED"), workflowResult.Tasks[0].Status)
}

func TestExecuteWorkflowSync(t *testing.T) {
	executor := testdata.WorkflowExecutor
	wf := workflow.NewConductorWorkflow(executor).
		Name("temp_wf_3_" + strconv.Itoa(time.Now().Nanosecond())).
		Version(1).
		OwnerEmail("test@orkes.io")
	wf = wf.Add(workflow.NewSetVariableTask("set_var").Input("var_value", 42))
	wf.OutputParameters(map[string]interface{}{
		"param1": "Test",
		"param2": 123,
	})
	err := wf.Register(true)

	assert.NoError(t, err, "Failed to register workflow")
	run, err := executeWorkflowWithRetries(wf, map[string]interface{}{
		"key1": "input1",
		"key2": 101,
	})
	if err != nil {
		t.Fatalf("Failed to complete the workflow, reason: %s", err)
	}
	assert.NotEmpty(t, run, "Workflow is null", run)
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

func startWorkers() {
	testdata.TaskRunner.StartWorker("simple_task", testdata.SimpleWorker, 10, 100*time.Millisecond)
	testdata.TaskRunner.StartWorker("dynamic_fork_prep", testdata.DynamicForkWorker, 3, 100*time.Millisecond)
}

func executeWorkflowWithRetries(wf *workflow.ConductorWorkflow, workflowInput interface{}) (*model.WorkflowRun, error) {
	for attempt := 0; attempt < retryLimit; attempt += 1 {
		workflowRun, err := wf.ExecuteWorkflowWithInput(workflowInput, "")
		if err != nil {
			time.Sleep(time.Duration(attempt+2) * time.Second)
			fmt.Println("Failed to execute workflow, reason: " + err.Error())
			continue
		}
		return workflowRun, nil
	}
	return nil, fmt.Errorf("exhausted retries for workflow execution")
}

func executeWorkflowWithRetriesWithStartWorkflowRequest(startWorkflowRequest *model.StartWorkflowRequest) (*model.WorkflowRun, error) {
	for attempt := 1; attempt <= retryLimit; attempt += 1 {
		workflowRun, err := testdata.WorkflowExecutor.ExecuteWorkflow(startWorkflowRequest, "")
		if err != nil {
			time.Sleep(time.Duration(attempt+2) * time.Second)
			fmt.Printf("Failed to execute workflow, reason: %s", err.Error())
			continue
		}
		return workflowRun, nil
	}
	return nil, fmt.Errorf("exhausted retries for workflow execution")
}

func TestUpgradeRunningWorkflowToVersion(t *testing.T) {
	input := &workflow.HttpInput{
		Method: workflow.GET,
		Uri:    "http://localhost:8081/api/hello/with-delay?name=SDK-Test&delaySeconds=1",
	}
	// Step 1: Create version 1 of a workflow
	workflowV1 := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_UPGRADE").
		Version(1).
		Add(workflow.NewHttpTask("http_1_ref_name", input))

	err := testdata.ValidateWorkflowRegistration(workflowV1)
	if err != nil {
		t.Fatal(
			"Failed to register workflow v1. Reason: ", err.Error(),
		)
	}

	// Step 2: Create version 2 of the workflow with an additional task
	workflowV2 := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_UPGRADE").
		Version(2).
		Add(workflow.NewHttpTask("http_1_ref_name", input)).Add(workflow.NewHttpTask("http_2_ref_name", input))

	err = testdata.ValidateWorkflowRegistration(workflowV2)
	if err != nil {
		t.Fatal(
			"Failed to register workflow v2. Reason: ", err.Error(),
		)
	}

	// Step 3: Start a workflow instance with version 1
	workflowInput := map[string]interface{}{
		"testKey": "testValue",
	}

	startRequest := &model.StartWorkflowRequest{
		Name:    workflowV1.GetName(),
		Version: 1,
		Input:   workflowInput,
	}

	workflowId, err := testdata.WorkflowExecutor.StartWorkflow(startRequest)
	if err != nil {
		t.Fatal(
			"Failed to start workflow. Reason: ", err.Error(),
		)
	}

	t.Logf("Started workflow with ID: %s", workflowId)

	// Step 5: Verify the workflow is running and at version 1
	workflow, err := testdata.WorkflowExecutor.GetWorkflow(workflowId, true)
	if err != nil {
		t.Fatal(
			"Failed to get workflow. Reason: ", err.Error(),
		)
	}

	if workflow.WorkflowVersion != 1 {
		t.Fatalf("Expected workflow version 1, got %d", workflow.WorkflowVersion)
	}

	// Step 6: Create upgrade request to version 2
	upgradeRequest := model.UpgradeWorkflowRequest{
		Name:    workflowV1.GetName(),
		Version: 2,
		// Optionally provide updated workflow input
		WorkflowInput: map[string]interface{}{
			"testKey": "updatedValue",
		},
	}

	// Step 7: Call upgrade API
	resp, err := testdata.WorkflowClient.UpgradeRunningWorkflowToVersion(
		context.Background(),
		upgradeRequest,
		workflowId,
	)

	// Step 8: Validate upgrade response
	if err != nil {
		t.Fatal(
			"Failed to upgrade workflow. Reason: ", err.Error(),
		)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, got %d", resp.StatusCode)
	}

	// Step 9: Wait for the upgrade to take effect
	time.Sleep(3 * time.Second)

	// Step 10: Verify the workflow has been upgraded to version 2
	upgradedWorkflow, err := testdata.WorkflowExecutor.GetWorkflow(workflowId, true)
	if err != nil {
		t.Fatal(
			"Failed to get upgraded workflow. Reason: ", err.Error(),
		)
	}

	if upgradedWorkflow.WorkflowVersion != 2 {
		t.Fatalf("Expected workflow version 2 after upgrade, got %d", upgradedWorkflow.WorkflowVersion)
	}

	// Step 11: Verify the additional task exists in the workflow
	var foundAdditionalTask bool
	for _, task := range upgradedWorkflow.Tasks {
		if task.ReferenceTaskName == "http_2_ref_name" {
			foundAdditionalTask = true
			break
		}
	}

	if !foundAdditionalTask {
		t.Fatal("Expected to find the additional task in the upgraded workflow")
	}

	// Step 12: Verify the updated workflow input
	updatedValue, ok := upgradedWorkflow.Input["testKey"].(string)
	if !ok || updatedValue != "updatedValue" {
		t.Fatalf("Expected workflow input to be updated with 'updatedValue', got %v", upgradedWorkflow.Input)
	}

	t.Log("Successfully upgraded workflow version")
}
