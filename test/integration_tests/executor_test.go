package integration_tests

import (
	"context"
	"fmt"
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
	err = executor.SignalTask(workflowId, model.CompletedTask, map[string]interface{}{
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

	// Now create the signal_subworkflow using the fluent API
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
	err = executor.SignalTask(parentWorkflowId, model.CompletedTask, map[string]interface{}{
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
	workflowRun, err := executor.ExecuteWorkflowWithBlockingWorkflow(
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
	err = executor.SignalTask(parentWorkflowId, model.CompletedTask, map[string]interface{}{
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

	// First, make sure the wait_signal_test workflow is registered
	// (We can reuse the code from the previous test or assume it's already registered)

	// Create the workflow start request
	startRequest := &model.StartWorkflowRequest{
		Name:    "signal_subworkflow",
		Version: 1,
		Input:   map[string]interface{}{},
	}

	// Execute the workflow with BLOCKING_WORKFLOW return strategy
	workflowRun, err := executor.ExecuteWorkflowWithBlockingWorkflow(
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
	err = executor.SignalTask(parentWorkflowId, model.CompletedTask, map[string]interface{}{
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
func registerComplexWorkflows(t *testing.T) {
	executor := testdata.WorkflowExecutor

	// Register complex_wf_signal_test_subworkflow_2 (innermost workflow)
	subWorkflow2 := workflow.ConductorWorkflow{}
	subWorkflow2.Name("complex_wf_signal_test_subworkflow_2").
		Description("complex_wf_signal_test_subworkflow_2").
		Version(1).
		OwnerEmail("shailesh.padave@orkes.io").
		// Add HTTP task
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
		// Add YIELD task
		Add(workflow.NewSimpleTask("yield", "simple_task_ref_1"))

	err := executor.RegisterWorkflow(true, subWorkflow2.ToWorkflowDef())
	assert.Nil(t, err)

	// Register complex_wf_signal_test_subworkflow_1 (middle workflow)
	subWorkflow1 := workflow.ConductorWorkflow{}
	subWorkflow1.Name("complex_wf_signal_test_subworkflow_1").
		Description("complex_wf_signal_test_subworkflow_1").
		Version(1).
		OwnerEmail("shailesh.padave@orkes.io").
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
		// Add YIELD task
		Add(workflow.NewSimpleTask("yield", "simple_task_ref_2")).
		// Add SUB_WORKFLOW task
		Add(workflow.NewSubWorkflowTask("sub_workflow_ref", "complex_wf_signal_test_subworkflow_2", 1))

	err = executor.RegisterWorkflow(true, subWorkflow1.ToWorkflowDef())
	assert.Nil(t, err)

	// Register complex_wf_signal_test (main workflow)
	mainWorkflow := workflow.ConductorWorkflow{}
	mainWorkflow.Name("complex_wf_signal_test").
		Description("http_yield_signal_test").
		Version(1).
		OwnerEmail("shailesh.padave@orkes.io").
		// Add HTTP task
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
		// Add SUB_WORKFLOW task
		Add(workflow.NewSubWorkflowTask("sub_workflow_ref", "complex_wf_signal_test_subworkflow_1", 1))

	err = executor.RegisterWorkflow(true, mainWorkflow.ToWorkflowDef())
	assert.Nil(t, err)
}

// Test complex workflow signaling with BLOCKING_WORKFLOW strategy
func TestComplexWorkflowSignalWithBlockingWorkflow(t *testing.T) {
	executor := testdata.WorkflowExecutor

	// Register all the complex workflows
	registerComplexWorkflows(t)

	startRequest := &model.StartWorkflowRequest{
		Name:    "complex_wf_signal_test",
		Version: 1,
		Input:   map[string]interface{}{},
	}

	// 1. Start workflow
	workflowRun, err := executor.ExecuteWorkflowWithBlockingWorkflow(
		startRequest,
		"",        // No waitUntilTask
		10,        // waitForSeconds
		"DURABLE", // consistency
	)
	assert.Nil(t, err)
	assert.NotEmpty(t, workflowRun)
	workflowId := workflowRun.WorkflowId

	t.Logf("Started complex workflow with ID: %s", workflowId)

	// Wait for workflow to execute the HTTP task and start the subworkflow
	time.Sleep(20 * time.Millisecond)

	// 2. Get workflow and check its status
	workflow, err := executor.GetWorkflow(workflowId, true)
	assert.Nil(t, err)
	assert.Equal(t, model.RunningWorkflow, workflow.Status, "Workflow should be RUNNING")

	// 3. Check that first task (HTTP) is completed and second task (SUBWORKFLOW) is in progress
	assert.GreaterOrEqual(t, len(workflow.Tasks), 2, "Workflow should have at least 2 tasks")

	httpTask := workflow.Tasks[0]
	assert.Equal(t, "JSON_JQ_TRANSFORM", httpTask.ReferenceTaskName, "First task should be http_ref")
	assert.Equal(t, model.CompletedTask, httpTask.Status, "HTTP task should be COMPLETED")

	subWorkflowTask := workflow.Tasks[1]
	assert.Equal(t, "sub_workflow_ref", subWorkflowTask.ReferenceTaskName, "Second task should be sub_workflow_ref")
	assert.Equal(t, model.InProgressTask, subWorkflowTask.Status, "SUBWORKFLOW task should be IN_PROGRESS")

	// 4. Signal with BLOCKING_WORKFLOW return strategy
	response1, err := executor.SignalTaskAndReturnBlockingWorkflow(
		workflow.WorkflowId,
		model.CompletedTask,
		map[string]interface{}{"result": "Signal received for first subworkflow"},
	)
	assert.Nil(t, err)

	// 5. Check the response
	assert.NotNil(t, response1, "Response should not be null")

	// Wait for the second subworkflow to start and reach the YIELD task
	time.Sleep(10 * time.Millisecond)

	// 6. Signal second subworkflow with BLOCKING_WORKFLOW return strategy
	response2, err := executor.SignalTaskAndReturnBlockingWorkflow(
		workflow.WorkflowId,
		model.CompletedTask,
		map[string]interface{}{"result": "Signal received for second subworkflow"},
	)
	assert.Nil(t, err)

	// 7. Check the response
	assert.NotNil(t, response2, "Response should not be null")
}

// Test complex workflow signaling with TARGET_WORKFLOW strategy
func TestComplexWorkflowSignalWithTargetWorkflow(t *testing.T) {
	executor := testdata.WorkflowExecutor

	// Register all the complex workflows
	registerComplexWorkflows(t)

	startRequest := &model.StartWorkflowRequest{
		Name:    "complex_wf_signal_test",
		Version: 1,
		Input:   map[string]interface{}{},
	}

	// 1. Start workflow
	workflowRun, err := executor.ExecuteWorkflowWithTargetWorkflow(
		startRequest,
		"",        // No waitUntilTask
		10,        // waitForSeconds
		"DURABLE", // consistency
	)
	assert.Nil(t, err)
	assert.NotEmpty(t, workflowRun)
	workflowId := workflowRun.WorkflowId

	t.Logf("Started complex workflow with ID: %s", workflowId)

	// Wait for workflow to execute the HTTP task and start the subworkflow
	time.Sleep(20 * time.Millisecond)

	// 2. Get workflow and check its status
	workflow, err := executor.GetWorkflow(workflowId, true)
	assert.Nil(t, err)
	assert.Equal(t, model.RunningWorkflow, workflow.Status, "Workflow should be RUNNING")

	// 3. Check that first task (HTTP) is completed and second task (SUBWORKFLOW) is in progress
	assert.GreaterOrEqual(t, len(workflow.Tasks), 2, "Workflow should have at least 2 tasks")

	httpTask := workflow.Tasks[0]
	assert.Equal(t, "JSON_JQ_TRANSFORM", httpTask.ReferenceTaskName, "First task should be http_ref")
	assert.Equal(t, model.CompletedTask, httpTask.Status, "HTTP task should be COMPLETED")

	subWorkflowTask := workflow.Tasks[1]
	assert.Equal(t, "sub_workflow_ref", subWorkflowTask.ReferenceTaskName, "Second task should be sub_workflow_ref")
	assert.Equal(t, model.InProgressTask, subWorkflowTask.Status, "SUBWORKFLOW task should be IN_PROGRESS")

	// 4. Signal with TARGET_WORKFLOW return strategy
	response1, err := executor.SignalTaskAndReturnTargetWorkflow(
		workflow.WorkflowId,
		model.CompletedTask,
		map[string]interface{}{"result": "Signal received for first subworkflow"},
	)
	assert.Nil(t, err)

	// 5. Check the response
	assert.NotNil(t, response1, "Response should not be null")

	// Wait for the second subworkflow to start and reach the YIELD task
	time.Sleep(10 * time.Millisecond)

	// 6. Signal second subworkflow with BLOCKING_WORKFLOW return strategy
	response2, err := executor.SignalTaskAndReturnTargetWorkflow(
		workflow.WorkflowId,
		model.CompletedTask,
		map[string]interface{}{"result": "Signal received for second subworkflow"},
	)
	assert.Nil(t, err)

	// 7. Check the response
	assert.NotNil(t, response2, "Response should not be null")
}

// Test complex workflow signaling with BLOCKING_TASK strategy
func TestComplexWorkflowSignalWithBlockingTask(t *testing.T) {
	executor := testdata.WorkflowExecutor

	// Register all the complex workflows
	registerComplexWorkflows(t)

	startRequest := &model.StartWorkflowRequest{
		Name:    "complex_wf_signal_test",
		Version: 1,
		Input:   map[string]interface{}{},
	}

	// 1. Start workflow
	workflowRun, err := executor.ExecuteWorkflowWithBlockingTask(
		startRequest,
		"",            // No waitUntilTask
		10,            // waitForSeconds
		"SYNCHRONOUS", // consistency
	)
	assert.Nil(t, err)
	assert.NotEmpty(t, workflowRun)
	workflowId := workflowRun.WorkflowId

	t.Logf("Started complex workflow with ID: %s", workflowId)

	// Wait for workflow to execute the HTTP task and start the subworkflow
	time.Sleep(20 * time.Millisecond)

	// 2. Get workflow and check its status
	workflow, err := executor.GetWorkflow(workflowId, true)
	assert.Nil(t, err)
	assert.Equal(t, model.RunningWorkflow, workflow.Status, "Workflow should be RUNNING")

	// 3. Check that first task (HTTP) is completed and second task (SUBWORKFLOW) is in progress
	assert.GreaterOrEqual(t, len(workflow.Tasks), 2, "Workflow should have at least 2 tasks")

	httpTask := workflow.Tasks[0]
	assert.Equal(t, "JSON_JQ_TRANSFORM", httpTask.ReferenceTaskName, "First task should be http_ref")
	assert.Equal(t, model.CompletedTask, httpTask.Status, "HTTP task should be COMPLETED")

	subWorkflowTask := workflow.Tasks[1]
	assert.Equal(t, "sub_workflow_ref", subWorkflowTask.ReferenceTaskName, "Second task should be sub_workflow_ref")
	assert.Equal(t, model.InProgressTask, subWorkflowTask.Status, "SUBWORKFLOW task should be IN_PROGRESS")

	// 4. Signal with TARGET_WORKFLOW return strategy
	response1, err := executor.SignalTaskAndReturnBlockingTask(
		workflow.WorkflowId,
		model.CompletedTask,
		map[string]interface{}{"result": "Signal received for first subworkflow"},
	)
	assert.Nil(t, err)

	// 5. Check the response
	assert.NotNil(t, response1, "Response should not be null")

	// Wait for the second subworkflow to start and reach the YIELD task
	time.Sleep(10 * time.Millisecond)

	// 6. Signal second subworkflow with BLOCKING_WORKFLOW return strategy
	response2, err := executor.SignalTaskAndReturnBlockingTask(
		workflow.WorkflowId,
		model.CompletedTask,
		map[string]interface{}{"result": "Signal received for second subworkflow"},
	)
	assert.Nil(t, err)

	// 7. Check the response
	assert.NotNil(t, response2, "Response should not be null")
}

// Test complex workflow signaling with BLOCKING_TASK_INPUT strategy
func TestComplexWorkflowSignalWithBlockingTaskInput(t *testing.T) {
	executor := testdata.WorkflowExecutor

	// Register all the complex workflows
	registerComplexWorkflows(t)

	startRequest := &model.StartWorkflowRequest{
		Name:    "complex_wf_signal_test",
		Version: 1,
		Input:   map[string]interface{}{},
	}

	// 1. Start workflow
	workflowRun, err := executor.ExecuteWorkflowWithBlockingTaskInput(
		startRequest,
		"",            // No waitUntilTask
		10,            // waitForSeconds
		"SYNCHRONOUS", // consistency
	)
	assert.Nil(t, err)
	assert.NotEmpty(t, workflowRun)
	workflowId := workflowRun.WorkflowId

	t.Logf("Started complex workflow with ID: %s", workflowId)

	// Wait for workflow to execute the HTTP task and start the subworkflow
	time.Sleep(20 * time.Millisecond)

	// 2. Get workflow and check its status
	workflow, err := executor.GetWorkflow(workflowId, true)
	assert.Nil(t, err)
	assert.Equal(t, model.RunningWorkflow, workflow.Status, "Workflow should be RUNNING")

	// 3. Check that first task (HTTP) is completed and second task (SUBWORKFLOW) is in progress
	assert.GreaterOrEqual(t, len(workflow.Tasks), 2, "Workflow should have at least 2 tasks")

	httpTask := workflow.Tasks[0]
	assert.Equal(t, "JSON_JQ_TRANSFORM", httpTask.ReferenceTaskName, "First task should be http_ref")
	assert.Equal(t, model.CompletedTask, httpTask.Status, "HTTP task should be COMPLETED")

	subWorkflowTask := workflow.Tasks[1]
	assert.Equal(t, "sub_workflow_ref", subWorkflowTask.ReferenceTaskName, "Second task should be sub_workflow_ref")
	assert.Equal(t, model.InProgressTask, subWorkflowTask.Status, "SUBWORKFLOW task should be IN_PROGRESS")

	// 4. Signal with TARGET_WORKFLOW return strategy
	response1, err := executor.SignalTaskAndReturnBlockingTaskInput(
		workflow.WorkflowId,
		model.CompletedTask,
		map[string]interface{}{"result": "Signal received for first subworkflow"},
	)
	assert.Nil(t, err)

	// 5. Check the response
	assert.NotNil(t, response1, "Response should not be null")

	// Wait for the second subworkflow to start and reach the YIELD task
	time.Sleep(10 * time.Millisecond)

	// 6. Signal second subworkflow with BLOCKING_TASK_INPUT return strategy
	response2, err := executor.SignalTaskAndReturnBlockingTaskInput(
		workflow.WorkflowId,
		model.CompletedTask,
		map[string]interface{}{"result": "Signal received for second subworkflow"},
	)
	assert.Nil(t, err)

	// 7. Check the response
	assert.NotNil(t, response2, "Response should not be null")
}
