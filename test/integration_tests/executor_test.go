package integration_tests

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/conductor-sdk/conductor-go/sdk/log"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
	"os"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"github.com/stretchr/testify/assert"
)

const (
	// IDs are hardcoded on purpose. It should not be found.
	notFoundWorkflowId = "2b3ea839-9aeb-11ef-9ac5-ce590b39fb93"
	notFoundTaskId     = "75c6875d-9ba8-11ef-82ba-0680bfba1f84"
)

func TestRetryNotFound(t *testing.T) {
	executor := testdata.WorkflowExecutor

	err := executor.Retry(notFoundWorkflowId, true)
	assert.Error(t, err, "Retry is expected to return an error")

	if swaggerErr, ok := err.(client.GenericSwaggerError); ok {
		assert.Equal(t, 404, swaggerErr.StatusCode())
	} else {
		assert.Fail(t, "err is not of type GenericSwaggerError")
	}
}

func TestRegisterWorkflow(t *testing.T) {
	executor := testdata.WorkflowExecutor

	wf := workflow.ConductorWorkflow{}
	wf.Name("registration_test_wf").
		Description("E2E test - Workflow Registration test").
		Version(1).
		Add(workflow.NewSimpleTask(
			"SIMPLE", "simple_ref",
		))

	// register the workflow
	err := executor.RegisterWorkflow(true, wf.ToWorkflowDef())
	assert.Nil(t, err)

	// modify the workflow and register with overwrite: false, to force a 409
	wf.Add(workflow.NewSimpleTask(
		"SIMPLE", "simple_ref_2",
	))
	err = executor.RegisterWorkflow(false, wf.ToWorkflowDef())
	assert.Error(t, err, "Registration is expected to return an error")

	if swaggerErr, ok := err.(client.GenericSwaggerError); ok {
		assert.Equal(t, 409, swaggerErr.StatusCode())
	} else {
		assert.Fail(t, "err is not of type GenericSwaggerError")
	}
}

func TestRegisterWorkflowWithTags(t *testing.T) {
	executor := testdata.WorkflowExecutor

	wf := workflow.NewConductorWorkflow(executor)

	// Create a map of tags
	tags := map[string]string{
		"environment": "production",
		"owner":       "data-team",
		"priority":    "high",
		"region":      "us-west",
		"version":     "1.2.3",
	}

	wf.Name("registration_test_wf").
		Description("E2E test - Workflow Registration test").
		Version(1).
		Add(workflow.NewSimpleTask(
			"SIMPLE", "simple_ref",
		)).
		Tags(tags)

	assert.Equal(t, tags, wf.GetTags())

	// register the workflow
	err := wf.Register(true)
	assert.Nil(t, err)

	actualTags, err := executor.GetWorkflowTags(wf.GetName())
	assert.Nil(t, err)
	assert.Equal(t, tags, actualTags)

	updateTags := map[string]string{
		"environment": "staging", // Changed from production to staging
		"owner":       "data-team",
		"priority":    "medium",  // Changed from high to medium
		"region":      "us-east", // Changed from us-west to us-east
		"version":     "1.3.0",   // Updated version
	}

	err = executor.UpdateWorkflowTags(wf.GetName(), updateTags)
	assert.Nil(t, err, "Expected no error while updating tags for workflow")

	actualTags, err = executor.GetWorkflowTags(wf.GetName())
	assert.Nil(t, err)
	assert.Equal(t, updateTags, actualTags)

	tagsToDelete := map[string]string{
		"priority": "medium",
		"version":  "1.3.0",
	}

	err = executor.DeleteWorkflowTags(wf.GetName(), tagsToDelete)
	assert.Nil(t, err)

	// After deletion, create the expected result by removing the deleted tags
	expectedRemainingTags := make(map[string]string)
	for k, v := range updateTags {
		if _, exists := tagsToDelete[k]; !exists {
			expectedRemainingTags[k] = v
		}
	}

	actualTags, err = executor.GetWorkflowTags(wf.GetName())
	assert.Nil(t, err)
	assert.Equal(t, expectedRemainingTags, actualTags, "Tags remaining after deletion should match expected")

	// remove created workflow
	executor.UnRegisterWorkflow(wf.GetName(), wf.GetVersion())
	executor.RemoveWorkflow(wf.GetName())
}

func TestGetWorkflow(t *testing.T) {
	executor := testdata.WorkflowExecutor

	wf, err := executor.GetWorkflow(notFoundWorkflowId, false)

	assert.Nil(t, wf)
	assert.Error(t, err, "GetWorkflow is expected to return an error")
	assert.Equal(t, fmt.Sprintf("no such workflow by Id %s", notFoundWorkflowId), err.Error())
}

func TestUpdateTaskByRefName(t *testing.T) {
	executor := testdata.WorkflowExecutor
	err := executor.UpdateTaskByRefName("task_ref", notFoundWorkflowId, model.CompletedTask, map[string]interface{}{})
	assert.Error(t, err, "UpdateTaskByRefName is expected to return an error")
	if swaggerErr, ok := err.(client.GenericSwaggerError); ok {
		assert.Equal(t, 404, swaggerErr.StatusCode())
	} else {
		assert.Fail(t, "err is not of type GenericSwaggerError")
	}
}

func TestUpdate(t *testing.T) {
	executor := testdata.WorkflowExecutor
	err := executor.UpdateTask(notFoundTaskId, notFoundWorkflowId, model.CompletedTask, map[string]interface{}{})
	assert.Error(t, err, "UpdateTask is expected to return an error")
	if swaggerErr, ok := err.(client.GenericSwaggerError); ok {
		assert.Equal(t, 404, swaggerErr.StatusCode())
	} else {
		assert.Fail(t, "err is not of type GenericSwaggerError")
	}
}

func TestStartWorkflowWithContext(t *testing.T) {
	executor := testdata.WorkflowExecutor

	ctx, cancel := context.WithCancel(context.Background())
	// cancel straightaway on purpose
	cancel()

	_, err := executor.StartWorkflowWithContext(ctx, &model.StartWorkflowRequest{})
	assert.Error(t, err, "StartWorkflowWithContext is expected to return an error")
	assert.Equal(t, context.Canceled, err, "Expected context canceled error")
}

func getSubWorkflow(t *testing.T) workflow.ConductorWorkflow {
	wf := workflow.ConductorWorkflow{}
	wf.Name("wait_signal_test").
		Description("wait_signal_test").
		Version(1).
		OwnerEmail("test.user@orkes.io").

		// Add the WAIT task
		Add(workflow.NewWaitTask("wait_ref")).

		// Add the JSON_JQ_TRANSFORM task
		Add(workflow.NewJQTask(
			"JSON_JQ_TRANSFORM", "json_transform_ref").
			Input("persons", []map[string]interface{}{
				{
					"name":  "some",
					"last":  "name",
					"email": "mail@mail.com",
					"id":    1,
				},
				{
					"name":  "some2",
					"last":  "name2",
					"email": "mail2@mail.com",
					"id":    2,
				},
			}).
			Input("queryExpression", ".persons | map({user:{email,id}})")).

		// Add the INLINE task
		Add(workflow.NewInlineTask(
			"INLINE", "inline_ref").
			Input("expression", "(function () {\n  return $.value1 + $.value2;\n})();").
			Input("evaluatorType", "graaljs").
			Input("value1", 1).
			Input("value2", 2))

	return wf
}

func TestRegisterWorkflowWithWaitSignal(t *testing.T) {
	executor := testdata.WorkflowExecutor
	wf := getSubWorkflow(t)

	// Register the workflow
	err := executor.RegisterWorkflow(true, wf.ToWorkflowDef())
	assert.Nil(t, err)

	// Start the workflow
	workflowId, err := executor.StartWorkflow(&model.StartWorkflowRequest{
		Name:    "wait_signal_test",
		Version: 1,
		Input:   map[string]interface{}{},
	})
	assert.Nil(t, err)
	assert.NotEmpty(t, workflowId)
	// Get the workflow to verify it's in RUNNING state
	workflow, err := executor.GetWorkflow(workflowId, true)
	assert.Nil(t, err)
	assert.Equal(t, model.RunningWorkflow, workflow.Status)

	// Signal the WAIT task to continue
	err = executor.SignalWorkflowTaskAsync(workflowId, model.CompletedTask, map[string]interface{}{
		"result": "Signal received, continuing workflow",
	})
	assert.Nil(t, err)

	// Wait a moment for processing
	time.Sleep(1 * time.Second)

	// Get the workflow again to verify it's completed
	workflow, err = executor.GetWorkflow(workflowId, true)
	assert.Nil(t, err)
	assert.Equal(t, model.CompletedWorkflow, workflow.Status)
}

func TestSubWorkflowSignal(t *testing.T) {
	executor := testdata.WorkflowExecutor

	subWf := getSubWorkflow(t)
	executor.RegisterWorkflow(true, subWf.ToWorkflowDef())

	wf := workflow.ConductorWorkflow{}
	wf.Name("signal_subworkflow").
		Description("signal_subworkflow").
		Version(1).
		OwnerEmail("test.user@orkes.io").
		// Add the SUB_WORKFLOW task
		Add(workflow.NewSubWorkflowTask("sub_workflow_ref", "wait_signal_test", 1))

	// Register the workflow
	err := executor.RegisterWorkflow(true, wf.ToWorkflowDef())
	assert.Nil(t, err)

	// Start the workflow
	parentWorkflowId, err := executor.StartWorkflow(&model.StartWorkflowRequest{
		Name:    "signal_subworkflow",
		Version: 1,
		Input:   map[string]interface{}{},
	})
	assert.Nil(t, err)
	assert.NotEmpty(t, parentWorkflowId)

	t.Logf("Started workflow with ID: %s", parentWorkflowId)

	// Wait a moment for the subworkflow to be started
	time.Sleep(1 * time.Second)

	// Get the parent workflow to find the subworkflow ID
	parentWorkflow, err := executor.GetWorkflow(parentWorkflowId, true)
	assert.Nil(t, err)
	assert.Equal(t, model.RunningWorkflow, parentWorkflow.Status, "Workflow should be in RUNNING state")

	// Ensure task list has at least one task
	assert.Greater(t, len(parentWorkflow.Tasks), 0, "Workflow should have at least one task")

	// Check that the first task is the WAIT task and it's in IN_PROGRESS state
	waitTask := parentWorkflow.Tasks[0]
	assert.Equal(t, "sub_workflow_ref", waitTask.ReferenceTaskName, "First task should be wait_ref")
	assert.Equal(t, model.InProgressTask, waitTask.Status, "WAIT task should be in IN_PROGRESS state")

	t.Logf("Verified workflow is RUNNING and WAIT task is IN_PROGRESS")

	// 3. Signal Workflow with task Completed
	err = executor.SignalWorkflowTaskAsync(parentWorkflowId, model.CompletedTask, map[string]interface{}{
		"result": "Signal received, continuing workflow",
	})
	assert.Nil(t, err)

	t.Logf("Sent COMPLETED signal to workflow")

	// Small delay to allow workflow to process
	time.Sleep(2 * time.Second)

	err = waitForWorkflowCompletion(executor, parentWorkflowId, 10*time.Second)
	assert.NoError(t, err)

	// 4. Check if WF status is completed
	workflowDetails, err := executor.GetWorkflow(parentWorkflowId, true)
	assert.Nil(t, err)
	assert.Equal(t, model.CompletedWorkflow, workflowDetails.Status, "Workflow should be in COMPLETED state")

	t.Logf("Verified workflow has COMPLETED after signaling")
}

// Helper function that polls until completion or timeout
func waitForWorkflowCompletion(executor *executor.WorkflowExecutor, workflowId string, maxWait time.Duration) error {
	deadline := time.Now().Add(maxWait)
	for time.Now().Before(deadline) {
		wf, err := executor.GetWorkflow(workflowId, false)
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}

		if wf.Status == model.CompletedWorkflow {
			return nil // Success!
		} else if wf.Status == model.FailedWorkflow || wf.Status == model.TerminatedWorkflow {
			return fmt.Errorf("workflow failed with status: %s", wf.Status)
		}

		// Exponential backoff - start with 1s, then increase
		time.Sleep(1 * time.Second)
	}
	return fmt.Errorf("timed out waiting for workflow %s to complete", workflowId)
}

func TestSubWorkflowSignalWithSyncConsistency(t *testing.T) {
	executor := testdata.WorkflowExecutor

	// First, make sure the wait_signal_test workflow is registered
	// (We can reuse the code from the previous test or assume it's already registered)

	// Create the workflow start request
	startRequest := &model.StartWorkflowRequest{
		Name:    "signal_subworkflow",
		Version: 1,
		Input:   map[string]interface{}{},
	}

	// Execute the workflow with BLOCKING_WORKFLOW return strategy
	workflowRun, err := executor.ExecuteAndGetBlockingWorkflow(
		startRequest,
		"",            // No waitUntilTask
		10,            // waitForSeconds
		"SYNCHRONOUS", // consistency
	)
	assert.NotNil(t, workflowRun)

	parentWorkflowId := workflowRun.WorkflowId

	t.Logf("Started workflow with ID: %s", parentWorkflowId)

	// Wait a moment for the subworkflow to be started
	time.Sleep(1 * time.Second)

	// Get the parent workflow to find the subworkflow ID
	parentWorkflow, err := executor.GetWorkflow(parentWorkflowId, true)
	assert.Nil(t, err)
	assert.Equal(t, model.RunningWorkflow, parentWorkflow.Status, "Workflow should be in RUNNING state")

	// Ensure task list has at least one task
	assert.Greater(t, len(parentWorkflow.Tasks), 0, "Workflow should have at least one task")

	// Check that the first task is the WAIT task and it's in IN_PROGRESS state
	waitTask := parentWorkflow.Tasks[0]
	assert.Equal(t, "wait_ref", waitTask.ReferenceTaskName, "First task should be wait_ref")
	assert.Equal(t, model.InProgressTask, waitTask.Status, "WAIT task should be in IN_PROGRESS state")

	t.Logf("Verified workflow is RUNNING and WAIT task is IN_PROGRESS")

	// 3. Signal Workflow with task Completed
	err = executor.SignalWorkflowTaskAsync(parentWorkflowId, model.CompletedTask, map[string]interface{}{
		"result": "Signal received, continuing workflow",
	})
	assert.Nil(t, err)

	t.Logf("Sent COMPLETED signal to workflow")

	// Small delay to allow workflow to process
	time.Sleep(2 * time.Second)

	// 4. Check if WF status is completed
	workflowDetails, err := executor.GetWorkflow(parentWorkflowId, true)
	assert.Nil(t, err)
	assert.Equal(t, model.CompletedWorkflow, workflowDetails.Status, "Workflow should be in COMPLETED state")

	t.Logf("Verified workflow has COMPLETED after signaling")
}

func TestSubWorkflowSignalWithDurableConsistency(t *testing.T) {
	executor := testdata.WorkflowExecutor

	// Create the workflow start request
	startRequest := &model.StartWorkflowRequest{
		Name:    "signal_subworkflow",
		Version: 1,
		Input:   map[string]interface{}{},
	}

	// Execute the workflow with BLOCKING_WORKFLOW return strategy
	workflowRun, err := executor.ExecuteAndGetBlockingWorkflow(
		startRequest,
		"",            // No waitUntilTask
		10,            // waitForSeconds
		"SYNCHRONOUS", // consistency
	)
	assert.NotNil(t, workflowRun)

	parentWorkflowId := workflowRun.WorkflowId

	t.Logf("Started workflow with ID: %s", parentWorkflowId)

	// Wait a moment for the subworkflow to be started
	time.Sleep(1 * time.Second)

	// Get the parent workflow to find the subworkflow ID
	parentWorkflow, err := executor.GetWorkflow(parentWorkflowId, true)
	assert.Nil(t, err)
	assert.Equal(t, model.RunningWorkflow, parentWorkflow.Status, "Workflow should be in RUNNING state")

	// Ensure task list has at least one task
	assert.Greater(t, len(parentWorkflow.Tasks), 0, "Workflow should have at least one task")

	// Check that the first task is the WAIT task and it's in IN_PROGRESS state
	waitTask := parentWorkflow.Tasks[0]
	assert.Equal(t, "wait_ref", waitTask.ReferenceTaskName, "First task should be wait_ref")
	assert.Equal(t, model.InProgressTask, waitTask.Status, "WAIT task should be in IN_PROGRESS state")

	t.Logf("Verified workflow is RUNNING and WAIT task is IN_PROGRESS")

	// 3. Signal Workflow with task Completed
	err = executor.SignalWorkflowTaskAsync(parentWorkflowId, model.CompletedTask, map[string]interface{}{
		"result": "Signal received, continuing workflow",
	})
	assert.Nil(t, err)

	t.Logf("Sent COMPLETED signal to workflow")

	// Small delay to allow workflow to process
	time.Sleep(2 * time.Second)

	// 4. Check if WF status is completed
	workflowDetails, err := executor.GetWorkflow(parentWorkflowId, true)
	assert.Nil(t, err)
	assert.Equal(t, model.CompletedWorkflow, workflowDetails.Status, "Workflow should be in COMPLETED state")

	t.Logf("Verified workflow has COMPLETED after signaling")
}

// Helper method to register all the complex workflows
func registerComplexWorkflows() {
	executor := testdata.WorkflowExecutor

	// Subworkflow-2
	wfDef, err := getWorkflowDef("complex_wf_signal_test_subworkflow_2.json")
	if err != nil {
		log.Fatalf("Failed to get workflow definition: %v", err)
	}
	err = executor.RegisterWorkflow(true, wfDef)
	if err != nil {
		log.Fatalf("Failed to register workflow: %v", err)
	}

	// Subworkflow-1
	wfDef, err = getWorkflowDef("complex_wf_signal_test_subworkflow_1.json")
	if err != nil {
		log.Fatalf("Failed to get workflow definition: %v", err)
	}
	err = executor.RegisterWorkflow(true, wfDef)
	if err != nil {
		log.Fatalf("Failed to register workflow: %v", err)
	}

	// Main WF
	wfDef, err = getWorkflowDef("complex_wf_signal_test.json")
	if err != nil {
		log.Fatalf("Failed to get workflow definition: %v", err)
	}
	err = executor.RegisterWorkflow(true, wfDef)
	if err != nil {
		log.Fatalf("Failed to register workflow: %v", err)
	}
}

func getWorkflowDef(filename string) (*model.WorkflowDef, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	defer file.Close()

	var def model.WorkflowDef
	if err := json.NewDecoder(file).Decode(&def); err != nil {
		return nil, fmt.Errorf("failed to decode JSON from %s: %w", filename, err)
	}

	return &def, nil
}

func TestSignal_AllStrategies_Comprehensive(t *testing.T) {
	registerComplexWorkflows()
	testCases := []struct {
		name                         string
		returnStrategy               model.ReturnStrategy
		consistency                  model.WorkflowConsistency
		expectedIsTarget             bool
		expectedIsBlocking           bool
		expectedIsTask               bool
		expectedIsTaskInput          bool
		shouldHaveWorkflow           bool
		shouldHaveTask               bool
		shouldHaveTaskInput          bool
		shouldValidateWorkflowFields bool
		shouldValidateTaskFields     bool
	}{
		{
			name:                         "TARGET_WORKFLOW",
			returnStrategy:               model.ReturnTargetWorkflow,
			consistency:                  model.SynchronousConsistency,
			expectedIsTarget:             true,
			expectedIsBlocking:           false,
			expectedIsTask:               false,
			expectedIsTaskInput:          false,
			shouldHaveWorkflow:           true,
			shouldHaveTask:               false,
			shouldHaveTaskInput:          false,
			shouldValidateWorkflowFields: true,
			shouldValidateTaskFields:     false,
		},
		{
			name:                         "BLOCKING_WORKFLOW",
			returnStrategy:               model.ReturnBlockingWorkflow,
			consistency:                  model.SynchronousConsistency,
			expectedIsTarget:             false,
			expectedIsBlocking:           true,
			expectedIsTask:               false,
			expectedIsTaskInput:          false,
			shouldHaveWorkflow:           true,
			shouldHaveTask:               false,
			shouldHaveTaskInput:          false,
			shouldValidateWorkflowFields: true,
			shouldValidateTaskFields:     false,
		},
		{
			name:                         "BLOCKING_TASK",
			returnStrategy:               model.ReturnBlockingTask,
			consistency:                  model.SynchronousConsistency,
			expectedIsTarget:             false,
			expectedIsBlocking:           false,
			expectedIsTask:               true,
			expectedIsTaskInput:          false,
			shouldHaveWorkflow:           false,
			shouldHaveTask:               true,
			shouldHaveTaskInput:          false,
			shouldValidateWorkflowFields: false,
			shouldValidateTaskFields:     true,
		},
		{
			name:                         "BLOCKING_TASK_INPUT",
			returnStrategy:               model.ReturnBlockingTaskInput,
			consistency:                  model.SynchronousConsistency,
			expectedIsTarget:             false,
			expectedIsBlocking:           false,
			expectedIsTask:               false,
			expectedIsTaskInput:          true,
			shouldHaveWorkflow:           false,
			shouldHaveTask:               true,
			shouldHaveTaskInput:          true,
			shouldValidateWorkflowFields: false,
			shouldValidateTaskFields:     true,
		},
		{
			name:                         "TARGET_WORKFLOW",
			returnStrategy:               model.ReturnTargetWorkflow,
			consistency:                  model.DurableConsistency,
			expectedIsTarget:             true,
			expectedIsBlocking:           false,
			expectedIsTask:               false,
			expectedIsTaskInput:          false,
			shouldHaveWorkflow:           true,
			shouldHaveTask:               false,
			shouldHaveTaskInput:          false,
			shouldValidateWorkflowFields: true,
			shouldValidateTaskFields:     false,
		},
		{
			name:                         "BLOCKING_WORKFLOW",
			returnStrategy:               model.ReturnBlockingWorkflow,
			consistency:                  model.DurableConsistency,
			expectedIsTarget:             false,
			expectedIsBlocking:           true,
			expectedIsTask:               false,
			expectedIsTaskInput:          false,
			shouldHaveWorkflow:           true,
			shouldHaveTask:               false,
			shouldHaveTaskInput:          false,
			shouldValidateWorkflowFields: true,
			shouldValidateTaskFields:     false,
		},
		{
			name:                         "BLOCKING_TASK",
			returnStrategy:               model.ReturnBlockingTask,
			consistency:                  model.DurableConsistency,
			expectedIsTarget:             false,
			expectedIsBlocking:           false,
			expectedIsTask:               true,
			expectedIsTaskInput:          false,
			shouldHaveWorkflow:           false,
			shouldHaveTask:               true,
			shouldHaveTaskInput:          false,
			shouldValidateWorkflowFields: false,
			shouldValidateTaskFields:     true,
		},
		{
			name:                         "BLOCKING_TASK_INPUT",
			returnStrategy:               model.ReturnBlockingTaskInput,
			consistency:                  model.DurableConsistency,
			expectedIsTarget:             false,
			expectedIsBlocking:           false,
			expectedIsTask:               false,
			expectedIsTaskInput:          true,
			shouldHaveWorkflow:           false,
			shouldHaveTask:               true,
			shouldHaveTaskInput:          true,
			shouldValidateWorkflowFields: false,
			shouldValidateTaskFields:     true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup workflow (same as your default test)
			executor := testdata.WorkflowExecutor

			startRequest := &model.StartWorkflowRequest{
				Name:    "complex_wf_signal_test",
				Version: 1,
				Input:   map[string]interface{}{},
			}

			// 1. Start workflow
			workflowRun, err := executor.ExecuteAndGetTarget(
				startRequest,
				"", // No waitUntilTask
				10, // waitForSeconds
				tc.consistency.String(),
			)
			assert.Nil(t, err)
			assert.NotEmpty(t, workflowRun)
			workflowId := workflowRun.WorkflowId

			t.Logf("Started complex workflow with ID: %s for strategy: %s, Consistency: %s", workflowId, tc.name, tc.consistency.String())

			// Wait for workflow to execute the HTTP task and start the subworkflow
			time.Sleep(20 * time.Millisecond)

			// 2. Get workflow and check its status
			workflow, err := executor.GetWorkflow(workflowId, true)
			assert.Nil(t, err)
			assert.Equal(t, model.RunningWorkflow, workflow.Status, "Workflow should be RUNNING")

			// 3. Check that first task (HTTP) is completed and second task (SUBWORKFLOW) is in progress
			assert.GreaterOrEqual(t, len(workflow.Tasks), 2, "Workflow should have at least 2 tasks")

			httpTask := workflow.Tasks[0]
			assert.Equal(t, "http_ref", httpTask.ReferenceTaskName, "First task should be http_ref")
			assert.Equal(t, model.CompletedTask, httpTask.Status, "HTTP task should be COMPLETED")

			subWorkflowTask := workflow.Tasks[1]
			assert.Equal(t, "sub_workflow_ref", subWorkflowTask.ReferenceTaskName, "Second task should be sub_workflow_ref")
			assert.Equal(t, model.InProgressTask, subWorkflowTask.Status, "SUBWORKFLOW task should be IN_PROGRESS")

			// 4. Signal with the test strategy
			resp, err := executor.Signal(workflowId, model.CompletedWorkflow,
				map[string]interface{}{"result": fmt.Sprintf("Signal received for %s", tc.name)},
				client.SignalTaskOpts{
					ReturnStrategy: tc.returnStrategy,
				})
			// Validate response type
			assert.Equal(t, tc.returnStrategy, resp.ResponseType, fmt.Sprintf("Response type should be %s", tc.returnStrategy))

			// ========== BASIC VALIDATIONS (same for all strategies) ==========
			assert.NoError(t, err, "Signal should not return an error")
			assert.NotNil(t, resp, "Response should not be nil")

			// Type check validations
			assert.Equal(t, tc.expectedIsTarget, resp.IsTargetWorkflow(), "IsTargetWorkflow check")
			assert.Equal(t, tc.expectedIsBlocking, resp.IsBlockingWorkflow(), "IsBlockingWorkflow check")
			assert.Equal(t, tc.expectedIsTask, resp.IsBlockingTask(), "IsBlockingTask check")
			assert.Equal(t, tc.expectedIsTaskInput, resp.IsBlockingTaskInput(), "IsBlockingTaskInput check")

			// Validate workflow identifiers (common for all)
			assert.NotEmpty(t, resp.WorkflowId, "WorkflowId should not be empty")
			assert.NotEmpty(t, resp.TargetWorkflowId, "TargetWorkflowId should not be empty")
			assert.NotEmpty(t, resp.WorkflowId, "WorkflowId should match the input workflowId")

			// Validate workflow status (common for all)
			assert.NotEmpty(t, resp.TargetWorkflowStatus, "TargetWorkflowStatus should not be empty")

			// Validate input/output data (common for all)
			assert.NotNil(t, resp.Input, "Input should not be nil")
			assert.NotNil(t, resp.Output, "Output should not be nil")

			// ========== WORKFLOW-SPECIFIC VALIDATIONS ==========
			if tc.shouldValidateWorkflowFields {
				// Validate workflow status and timestamps
				assert.NotEmpty(t, resp.Status, "Status should not be empty")
				assert.Greater(t, resp.CreateTime, int64(0), "CreateTime should be greater than 0")
				assert.Greater(t, resp.UpdateTime, int64(0), "UpdateTime should be greater than 0")
				assert.GreaterOrEqual(t, resp.UpdateTime, resp.CreateTime, "UpdateTime should be >= CreateTime")

				// Validate tasks array for workflow responses
				assert.NotNil(t, resp.Tasks, "Tasks should not be nil")
				if len(resp.Tasks) > 0 {
					assert.Greater(t, len(resp.Tasks), 0, "Should have at least one task")

					// Validate first task
					firstTask := resp.Tasks[0]
					assert.NotEmpty(t, firstTask.TaskId, "Task ID should not be empty")
					assert.NotEmpty(t, firstTask.TaskType, "Task type should not be empty")
					assert.NotEmpty(t, firstTask.ReferenceTaskName, "Reference task name should not be empty")
					assert.NotEmpty(t, firstTask.TaskDefName, "Task definition name should not be empty")
					assert.NotNil(t, firstTask.WorkflowInstanceId, "Task workflow instance ID should not be nil")
				}
			}

			// ========== TASK-SPECIFIC VALIDATIONS ==========
			if tc.shouldValidateTaskFields {
				// Task-specific field validations
				assert.NotEmpty(t, resp.TaskType, "TaskType should not be empty")
				assert.NotEmpty(t, resp.TaskId, "TaskId should not be empty")
				assert.NotEmpty(t, resp.ReferenceTaskName, "ReferenceTaskName should not be empty")
				assert.NotEmpty(t, resp.TaskDefName, "TaskDefName should not be empty")
				assert.NotEmpty(t, resp.WorkflowType, "WorkflowType should not be empty")
				assert.NotEmpty(t, resp.Status, "Status should not be empty")
			}

			// ========== HELPER METHOD VALIDATIONS ==========
			if tc.shouldHaveWorkflow {
				// Test GetWorkflow helper method
				workflowFromResp, err := resp.GetWorkflow()
				assert.NoError(t, err, "GetWorkflow should not return an error")
				assert.NotNil(t, workflowFromResp, "Workflow should not be nil")

				// Validate workflow data matches response
				assert.Equal(t, resp.WorkflowId, workflowFromResp.WorkflowId, "Workflow ID should match")
				assert.Equal(t, resp.Status, workflowFromResp.Status, "Workflow status should match")
				assert.Equal(t, resp.CreateTime, workflowFromResp.CreateTime, "Create time should match")
				assert.Equal(t, resp.UpdateTime, workflowFromResp.UpdateTime, "Update time should match")
				assert.Equal(t, resp.CreatedBy, workflowFromResp.CreatedBy, "Created by should match")
				assert.Equal(t, len(resp.Tasks), len(workflowFromResp.Tasks), "Tasks count should match")
			} else {
				_, err := resp.GetWorkflow()
				assert.Error(t, err, "GetWorkflow should return error for non-workflow responses")
			}

			if tc.shouldHaveTask {
				// Test GetBlockingTask helper method
				task, err := resp.GetBlockingTask()
				assert.NoError(t, err, "GetBlockingTask should not return an error")
				assert.NotNil(t, task, "Task should not be nil")

				// Validate task data matches response
				assert.Equal(t, resp.TaskId, task.TaskId, "Task ID should match")
				assert.Equal(t, resp.TaskType, task.TaskType, "Task type should match")
				assert.Equal(t, resp.ReferenceTaskName, task.ReferenceTaskName, "Reference task name should match")
				assert.Equal(t, resp.TaskDefName, task.TaskDefName, "Task definition name should match")
				assert.Equal(t, resp.WorkflowType, task.WorkflowType, "Workflow type should match")
			} else {
				_, err := resp.GetBlockingTask()
				assert.Error(t, err, "GetBlockingTask should return error for non-task responses")
			}

			if tc.shouldHaveTaskInput {
				// Test GetTaskInput helper method
				taskInput, err := resp.GetTaskInput()
				assert.NoError(t, err, "GetTaskInput should not return an error")
				assert.NotNil(t, taskInput, "Task input should not be nil")
				assert.Equal(t, resp.Input, taskInput, "Task input should match response input")
			} else {
				_, err := resp.GetTaskInput()
				assert.Error(t, err, "GetTaskInput should return error for non-task-input responses")
			}

			resp, err = executor.Signal(workflowId, model.CompletedWorkflow,
				map[string]interface{}{"result": fmt.Sprintf("Signal received for %s", tc.name)},
				client.SignalTaskOpts{
					ReturnStrategy: tc.returnStrategy,
				})
			assert.NoError(t, err, "Signal should not return an error")
			if tc.returnStrategy == model.ReturnBlockingTask {
				assert.Empty(t, resp, "Signal response should not be nil")
			}
		})
	}
}

// Add a separate test for default strategy to ensure it behaves like TARGET_WORKFLOW
func TestSignal_DefaultStrategy_IsTargetWorkflow(t *testing.T) {
	// Setup workflow (same setup code as in your original test)
	executor := testdata.WorkflowExecutor
	registerComplexWorkflows()

	startRequest := &model.StartWorkflowRequest{
		Name:    "complex_wf_signal_test",
		Version: 1,
		Input:   map[string]interface{}{},
	}

	workflowRun, err := executor.ExecuteAndGetTarget(startRequest, "", 10, "DURABLE")
	assert.Nil(t, err)
	workflowId := workflowRun.WorkflowId

	time.Sleep(20 * time.Millisecond)

	// Signal with NO ReturnStrategy specified (should default to TARGET_WORKFLOW)
	resp, err := executor.Signal(workflowId, model.CompletedWorkflow,
		map[string]interface{}{"result": "Signal received for default strategy"},
		client.SignalTaskOpts{}) // Empty options - should use default

	// Should behave exactly like TARGET_WORKFLOW
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, model.ReturnTargetWorkflow, resp.ResponseType, "Default strategy should be TARGET_WORKFLOW")
	assert.True(t, resp.IsTargetWorkflow(), "Default should behave like TARGET_WORKFLOW")

	// All the same validations as TARGET_WORKFLOW...
	assert.NotEmpty(t, resp.WorkflowId)
	assert.Greater(t, resp.CreateTime, int64(0))
	assert.Greater(t, resp.UpdateTime, int64(0))

	workflow, err := resp.GetWorkflow()
	assert.NoError(t, err)
	assert.NotNil(t, workflow)

	// Signal with NO ReturnStrategy specified (should default to TARGET_WORKFLOW)
	resp, err = executor.Signal(workflowId, model.CompletedWorkflow,
		map[string]interface{}{"result": "Signal received for default strategy"},
		client.SignalTaskOpts{}) // Empty options - should use default
	assert.NoError(t, err)
}
