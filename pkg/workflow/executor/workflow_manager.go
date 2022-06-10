package executor

import (
	"context"
	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/model"
)

//GetWorkflow Get workflow execution by workflow Id.  If includeTasks is set, also fetches all the task details.
//Returns nil if no workflow is found by the id
func (e *WorkflowExecutor) GetWorkflow(workflowId string, includeTasks bool) (*model.Workflow, error) {
	workflow, response, err := e.workflowClient.GetExecutionStatus(
		context.Background(),
		workflowId,
		&conductor_http_client.WorkflowResourceApiGetExecutionStatusOpts{
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
		&conductor_http_client.WorkflowResourceApiGetWorkflowsOpts{
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
		&conductor_http_client.WorkflowResourceApiSearchOpts{
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

func (e *WorkflowExecutor) GetAllRunningWorkflows(workflowName string, version *int32) error {
	e.workflowClient.GetRunningWorkflow(
		context.Background(),
		workflowName,
		&conductor_http_client.WorkflowResourceApiGetRunningWorkflowOpts{
			Version:   optional.Int32{},
			StartTime: optional.Int64{},
			EndTime:   optional.Int64{},
		},
	)
	return nil
}

func (e *WorkflowExecutor) Pause(workflowId string) error {
	return nil
}

func (e *WorkflowExecutor) Resume(workflowId string) error {
	return nil
}

func (e *WorkflowExecutor) Terminate(workflowId string, reason string) error {
	return nil
}

func (e *WorkflowExecutor) Restart(workflowId string, reason string) error {
	return nil
}

func (e *WorkflowExecutor) Retry(workflowId string, reason string) error {
	return nil
}

func (e *WorkflowExecutor) ReRun(workflowId string, reason string) error {
	return nil
}

func (e *WorkflowExecutor) SkipTasksFromWorkflow(workflowId string, reason string) error {
	return nil
}
