package integration_tests

import (
	"context"
	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"net/http"
	"testing"
)

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

	// Create a task mock for the HTTP task to simulate its output
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
			QueueWaitTime: 5, // 5 milliseconds
		},
	}

	// Prepare the workflow test request
	testRequest := model.WorkflowTestRequest{
		Name:                httpTaskWorkflow.GetName(),
		Version:             httpTaskWorkflow.GetVersion(),
		Input:               map[string]interface{}{"inputParam1": "testValue1"},
		TaskRefToMockOutput: taskMocks,
		// Optionally include the workflow definition directly
		WorkflowDef: httpTaskWorkflow.ToWorkflowDef(),
	}

	// Call the test workflow API
	workflowResult, resp, err := testdata.WorkflowClient.TestWorkflow(context.Background(), testRequest)

	// Validate response
	if err != nil {
		t.Fatalf("Failed to test workflow. API Error: %s", err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, got %d", resp.StatusCode)
	}

	// Validate workflow result
	if workflowResult.Status != "COMPLETED" {
		t.Fatalf("Expected workflow status COMPLETED, got %s", workflowResult.Status)
	}

	// Validate tasks in workflow result
	if len(workflowResult.Tasks) != 1 {
		t.Fatalf("Expected 1 task in workflow result, got %d", len(workflowResult.Tasks))
	}

	// Validate http task output
	httpTask := workflowResult.Tasks[0]
	if httpTask.Status != "COMPLETED" {
		t.Fatalf("Expected HTTP task status COMPLETED, got %s", httpTask.Status)
	}

	// Check execution metrics
	if httpTask.QueueWaitTime != 5 {
		t.Fatalf("Expected HTTP task queue wait time to be 5ms, got %d", httpTask.QueueWaitTime)
	}
}

func TestUpgradeRunningWorkflowToVersion(t *testing.T) {
	// Create an HTTP task with a longer delay to ensure workflow stays in RUNNING state
	httpInput := &workflow.HttpInput{
		Method: "GET",
		Uri:    "http://httpbin:8081/api/hello/with-delay?name=Sdktest&delaySeconds=2",
	}

	// Step 1: Create version 1 of a workflow
	workflowV1 := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_UPGRADE").
		Version(1).
		Add(workflow.NewHttpTask("additional_task_1", httpInput))

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
		Add(workflow.NewHttpTask("additional_task_1", &workflow.HttpInput{
			Method: "GET",
			Uri:    "http://httpbin:8081/api/hello/with-delay?name=Sdktest",
		})).
		Add(workflow.NewHttpTask("additional_task_2", &workflow.HttpInput{
			Method: "GET",
			Uri:    "http://httpbin:8081/api/hello/with-delay?name=Sdktest",
		}))

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

	// Step 5: Verify the workflow is running (not completed) and at version 1
	workflow, err := testdata.WorkflowExecutor.GetWorkflow(workflowId, true)
	if err != nil {
		t.Fatal(
			"Failed to get workflow. Reason: ", err.Error(),
		)
	}

	// Check that workflow is still in RUNNING state
	if workflow.Status != "RUNNING" {
		t.Fatal(
			"Workflow is not in RUNNING state, cannot upgrade. Current state: ", workflow.Status,
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
		// Optionally provide task output for tasks that will be added
		TaskOutput: map[string]interface{}{
			"additional_task_2": map[string]interface{}{
				"outputKey": "outputValue",
			},
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

	// Step 10: Verify the workflow has been upgraded to version 2
	upgradedWorkflow, err := testdata.WorkflowExecutor.GetWorkflow(workflowId, true)
	if err != nil {
		t.Fatal(
			"Failed to get upgraded workflow. Reason: ", err.Error(),
		)
	}

	// Step 11: Verify the additional task exists in the workflow
	var foundAdditionalTask bool
	for _, task := range upgradedWorkflow.Tasks {
		if task.ReferenceTaskName == "additional_task_2" {
			foundAdditionalTask = true

			// If task already has output data, check for our expected value
			if task.Status == "COMPLETED" {
				outputValue, ok := task.OutputData["outputKey"].(string)
				if !ok || outputValue != "outputValue" {
					t.Fatalf("Expected additional task output to contain outputKey with value 'outputValue', got %v", task.OutputData)
				}
			}
			break
		}
	}

	if !foundAdditionalTask {
		t.Fatal("Expected to find the additional task in the upgraded workflow")
	}
}

func TestJumpToTask(t *testing.T) {
	workflowDef := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_JUMP").
		Version(1).
		Add(testdata.TestSimpleTask).
		Add(testdata.TestHttpTask)

	err := testdata.ValidateWorkflowRegistration(workflowDef)
	if err != nil {
		t.Fatal(
			"Failed to register workflow. Reason: ", err.Error(),
		)
	}

	// Start a workflow instance
	workflowInput := map[string]interface{}{
		"testKey": "testValue",
	}

	startRequest := &model.StartWorkflowRequest{
		Name:    workflowDef.GetName(),
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

	// Verify workflow is running
	workflow, err := testdata.WorkflowExecutor.GetWorkflow(workflowId, true)
	if err != nil {
		t.Fatal(
			"Failed to get workflow. Reason: ", err.Error(),
		)
	}

	if workflow.Status != "RUNNING" {
		t.Fatalf("Expected workflow status RUNNING, got %s", workflow.Status)
	}

	// Define custom input for the task we're jumping to
	jumpTaskInput := map[string]interface{}{
		"uri":    "http://httpbin:8081/api/hello?name=Test1",
		"method": "GET",
	}

	// Jump to the third task, skipping the second
	opts := &client.WorkflowResourceApiJumpToTaskOpts{
		TaskReferenceName: optional.NewString("TEST_GO_TASK_HTTP"),
	}

	resp, err := testdata.WorkflowClient.JumpToTask(
		context.Background(),
		jumpTaskInput,
		workflowId,
		opts,
	)

	if err != nil {
		t.Fatal(
			"Failed to jump to task. Reason: ", err.Error(),
		)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, got %d", resp.StatusCode)
	}

	// Get the updated workflow
	updatedWorkflow, err := testdata.WorkflowExecutor.GetWorkflow(workflowId, true)
	if err != nil {
		t.Fatal(
			"Failed to get updated workflow. Reason: ", err.Error(),
		)
	}

	// Verify that the second task was skipped and the third task is being executed or completed
	var firstTaskSkipped bool
	var secondTaskActive bool

	for _, task := range updatedWorkflow.Tasks {
		if task.ReferenceTaskName == "TEST_GO_TASK_SIMPLE" {
			if task.Status == "SKIPPED" {
				firstTaskSkipped = true
			}
		}

		if task.ReferenceTaskName == "TEST_GO_TASK_HTTP" {
			if task.Status == "IN_PROGRESS" || task.Status == "COMPLETED" {
				secondTaskActive = true
			}
		}
	}

	if !firstTaskSkipped {
		t.Fatal("Expected the second task to be skipped")
	}

	if !secondTaskActive {
		t.Fatal("Expected the third task to be active after jumping")
	}

	t.Log("Successfully tested jump to task functionality")
}
