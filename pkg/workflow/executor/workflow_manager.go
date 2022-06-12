package executor

import (
	"context"
	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/pkg/client"
	"github.com/conductor-sdk/conductor-go/pkg/model"
)

//GetWorkflow Get workflow execution by workflow Id.  If includeTasks is set, also fetches all the task details.
//Returns nil if no workflow is found by the id
func (e *WorkflowExecutor) GetWorkflow(workflowId string, includeTasks bool) (*model.Workflow, error) {
	workflow, response, err := e.workflowClient.GetExecutionStatus(
		context.Background(),
		workflowId,
		&client.WorkflowResourceApiGetExecutionStatusOpts{
			IncludeTasks: optional.NewBool(includeTasks)},
	)
	if response.StatusCode == 404 {
		return nil, nil
	}
	return &workflow, err
}

//GetWorkflowStatus Get the status of the workflow execution.
//This is a lightweight method that returns only overall state of the workflow
func (e *WorkflowExecutor) GetWorkflowStatus(workflowId string, includeOutput bool, includeVariables bool) (*model.WorkflowState, error) {
	state, response, err := e.workflowClient.GetWorkflowState(context.Background(), workflowId, includeOutput, includeVariables)
	if response.StatusCode == 404 {
		return nil, nil
	}
	return &state, err
}

//GetByCorrelationIds Given the list of correlation ids, find and return workflows
//Returns a map with key as correlationId and value as a list of Workflows
//When IncludeClosed is set to true, the return value also includes workflows that are completed otherwise only running workflows are returned
func (e *WorkflowExecutor) GetByCorrelationIds(workflowName string, includeClosed bool, includeTasks bool, correlationIds ...string) (map[string][]model.Workflow, error) {
	workflows, _, err := e.workflowClient.GetWorkflows(
		context.Background(),
		correlationIds,
		workflowName,
		&client.WorkflowResourceApiGetWorkflowsOpts{
			IncludeClosed: optional.NewBool(includeClosed),
			IncludeTasks:  optional.NewBool(includeTasks),
		})
	if err != nil {
		return nil, err
	}
	return workflows, nil
}

//Search searches for workflows
//
// - Start: Start index - used for pagination
//
// - Size:  Number of results to return
//
// - Query: Query expression.  In the format FIELD = 'VALUE' or FIELD IN (value1, value2)
//   		Only AND operations are supported.  e.g. workflowId IN ('a', 'b', 'c') ADN workflowType ='test_workflow'
//			AND startTime BETWEEN 1000 and 2000
//			Supported fields for Query are:workflowId,workflowType,status,startTime
// - FreeText: Full text search.  All the workflow input, output and task outputs upto certain limit (check with your admins to find the size limit)
//			are full text indexed and can be used to search
func (e *WorkflowExecutor) Search(start int32, size int32, query string, freeText string) ([]model.WorkflowSummary, error) {
	workflows, _, err := e.workflowClient.Search(
		context.Background(),
		&client.WorkflowResourceApiSearchOpts{
			Start:    optional.NewInt32(start),
			Size:     optional.NewInt32(size),
			FreeText: optional.NewString(freeText),
			Query:    optional.NewString(query),
		},
	)
	if err != nil {
		return nil, err
	}
	return workflows.Results, nil
}

//Pause the execution of a running workflow.
//Any tasks that are currently running will finish but no new tasks are scheduled until the workflow is resumed
func (e *WorkflowExecutor) Pause(workflowId string) error {
	_, err := e.workflowClient.PauseWorkflow(context.Background(), workflowId)
	if err != nil {
		return err
	}
	return err
}

//Resume the execution of a workflow that is paused.  If the workflow is not paused, this method has no effect
func (e *WorkflowExecutor) Resume(workflowId string) error {
	_, err := e.workflowClient.ResumeWorkflow(context.Background(), workflowId)
	if err != nil {
		return err
	}
	return err
}

//Terminate a running workflow.  Reason must be provided that is captured as the termination resaon for the workflow
func (e *WorkflowExecutor) Terminate(workflowId string, reason string) error {
	_, err := e.workflowClient.Terminate(context.Background(), workflowId,
		&client.WorkflowResourceApiTerminateOpts{Reason: optional.NewString(reason)},
	)
	if err != nil {
		return err
	}
	return err
}

//Restart a workflow execution from the beginning with the same input.
//When called on a workflow that is not in a terminal status, this operation has no effect
//If useLatestDefinition is set, the restarted workflow fetches the latest definition from the metadata store
func (e *WorkflowExecutor) Restart(workflowId string, useLatestDefinition bool) error {
	_, err := e.workflowClient.Restart(
		context.Background(),
		workflowId,
		&client.WorkflowResourceApiRestartOpts{
			UseLatestDefinitions: optional.NewBool(useLatestDefinition),
		})
	if err != nil {
		return err
	}
	return err
}

//Retry a failed workflow from the last task that failed.  When called the task in the failed state is scheduled again
//and workflow moves to RUNNING status.  If resumeSubworkflowTasks is set and the last failed task was a sub-workflow
//the server restarts the subworkflow from the failed task.  If set to false, the sub-workflow is re-executed.
func (e *WorkflowExecutor) Retry(workflowId string, resumeSubworkflowTasks bool) error {
	_, err := e.workflowClient.Retry(
		context.Background(),
		workflowId,
		&client.WorkflowResourceApiRetryOpts{
			ResumeSubworkflowTasks: optional.NewBool(resumeSubworkflowTasks),
		},
	)
	if err != nil {
		return nil
	}
	return err
}

// ReRun a completed workflow from a specific task (ReRunFromTaskId) and optionally change the input
// Also update the completed tasks with new input (ReRunFromTaskId) if required
func (e *WorkflowExecutor) ReRun(workflowId string, reRunRequest model.RerunWorkflowRequest) (id string, error error) {
	id, _, err := e.workflowClient.Rerun(
		context.Background(),
		reRunRequest,
		workflowId,
	)
	if err != nil {
		return "", err
	}
	return id, err
}

//SkipTasksFromWorkflow Skips a given task execution from a current running workflow.
//When skipped the task's input and outputs are updated  from skipTaskRequest parameter.
func (e *WorkflowExecutor) SkipTasksFromWorkflow(workflowId string, taskReferenceName string, skipTaskRequest model.SkipTaskRequest) error {
	_, err := e.workflowClient.SkipTaskFromWorkflow(
		context.Background(),
		workflowId,
		taskReferenceName,
		skipTaskRequest,
	)
	if err != nil {
		return err
	}
	return nil
}
