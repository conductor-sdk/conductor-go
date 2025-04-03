package integration_tests

import (
	"context"
	"fmt"
	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
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
	assert.Error(t, err, "Retry is expected to return an error")

	if swaggerErr, ok := err.(client.GenericSwaggerError); ok {
		assert.Equal(t, 409, swaggerErr.StatusCode())
	} else {
		assert.Fail(t, "err is not of type GenericSwaggerError")
	}
}

func TestRegisterWorkflowWithTags(t *testing.T) {
	executor := testdata.WorkflowExecutor

	wf := workflow.NewConductorWorkflow(executor)

	// remove already created workflow before starting test
	executor.UnRegisterWorkflow(wf.GetName(), 1)

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

	// register the workflow
	err := executor.RegisterWorkflow(true, wf.ToWorkflowDef())
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
