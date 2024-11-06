//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package executor

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/concurrency"
	"github.com/conductor-sdk/conductor-go/sdk/event/queue"
	"github.com/conductor-sdk/conductor-go/sdk/model"

	log "github.com/sirupsen/logrus"
)

type WorkflowExecutor struct {
	metadataClient *client.MetadataResourceApiService
	taskClient     *client.TaskResourceApiService
	workflowClient *client.WorkflowResourceApiService
	eventClient    *client.EventResourceApiService

	workflowMonitor *WorkflowMonitor

	startWorkflowBatchSize   int
	waitForWorkflowBatchSize int
}

const (
	startWorkflowBatchSizeEnv   = "WORKFLOW_EXECUTOR_START_BATCH_SIZE"
	waitForWorkflowBatchSizeEnv = "WORKFLOW_EXECUTOR_WAIT_BATCH_SIZE"
)

// NewWorkflowExecutor Create a new workflow executor
func NewWorkflowExecutor(apiClient *client.APIClient) *WorkflowExecutor {
	metadataClient := client.MetadataResourceApiService{
		APIClient: apiClient,
	}
	taskClient := client.TaskResourceApiService{
		APIClient: apiClient,
	}
	workflowClient := client.WorkflowResourceApiService{
		APIClient: apiClient,
	}
	eventClient := client.EventResourceApiService{
		APIClient: apiClient,
	}
	startWorkflowBatchSize, err := getEnvInt(startWorkflowBatchSizeEnv)
	if err != nil {
		startWorkflowBatchSize = 256
	}
	waitForWorkflowBatchSize, err := getEnvInt(waitForWorkflowBatchSizeEnv)
	if err != nil {
		waitForWorkflowBatchSize = 256
	}
	workflowExecutor := WorkflowExecutor{
		metadataClient:           &metadataClient,
		taskClient:               &taskClient,
		workflowClient:           &workflowClient,
		eventClient:              &eventClient,
		workflowMonitor:          NewWorkflowMonitor(&workflowClient),
		startWorkflowBatchSize:   startWorkflowBatchSize,
		waitForWorkflowBatchSize: waitForWorkflowBatchSize,
	}
	return &workflowExecutor
}

// RegisterWorkflow Registers the workflow on the server.  Overwrites if the flag is set.  If the 'overwrite' flag is not set
// and the workflow definition differs from the one on the server, the call will fail with response code 409
func (e *WorkflowExecutor) RegisterWorkflow(overwrite bool, workflow *model.WorkflowDef) error {
	_, err := e.metadataClient.RegisterWorkflowDef(
		context.Background(),
		overwrite,
		*workflow,
	)
	if err != nil {
		return err
	}

	return nil
}

// UnRegisterWorkflow Un-registers the workflow on the server.
func (e *WorkflowExecutor) UnRegisterWorkflow(name string, version int32) error {
	_, err := e.metadataClient.UnregisterWorkflowDef(
		context.Background(),
		name,
		version,
	)
	if err != nil {
		return err
	}

	return nil
}

// ExecuteWorkflow start a workflow and wait until the workflow completes or the waitUntilTask completes
// Returns the output of the workflow
func (e *WorkflowExecutor) ExecuteWorkflow(startWorkflowRequest *model.StartWorkflowRequest, waitUntilTask string) (run *model.WorkflowRun, err error) {
	requestId := ""
	version := startWorkflowRequest.Version
	workflowRun, _, error := e.workflowClient.ExecuteWorkflow(context.Background(), *startWorkflowRequest, requestId, startWorkflowRequest.Name, version, waitUntilTask)
	if error != nil {
		return nil, error
	}
	return &workflowRun, nil
}

// MonitorExecution monitors the workflow execution
// Returns the channel with the execution result of the workflow
// Note: Channels will continue to grow if the workflows do not complete and/or are not taken out
func (e *WorkflowExecutor) MonitorExecution(workflowId string) (workflowMonitor WorkflowExecutionChannel, err error) {
	return e.workflowMonitor.generateWorkflowExecutionChannel(workflowId)
}

// StartWorkflow Start workflows
// Returns the id of the newly created workflow
func (e *WorkflowExecutor) StartWorkflow(startWorkflowRequest *model.StartWorkflowRequest) (workflowId string, err error) {
	id, _, err := e.workflowClient.StartWorkflowWithRequest(
		context.Background(),
		*startWorkflowRequest,
	)
	if err != nil {
		return "", err
	}
	return id, nil
}

// StartWorkflows Start workflows in bulk
// Returns RunningWorkflow struct that contains the workflowId, Err (if failed to start) and an execution channel
// which can be used to monitor the completion of the workflow execution.  The channel is available if monitorExecution is set
func (e *WorkflowExecutor) StartWorkflows(monitorExecution bool, startWorkflowRequests ...*model.StartWorkflowRequest) []*RunningWorkflow {
	amount := len(startWorkflowRequests)
	log.Debug(fmt.Sprintf("Starting %d workflows", amount))
	startingWorkflowChannel := make([]chan *RunningWorkflow, amount)
	for idx := 0; idx < len(startWorkflowRequests); {
		var waitGroup sync.WaitGroup
		for batchIdx := 0; idx < len(startWorkflowRequests) && batchIdx < e.startWorkflowBatchSize; batchIdx, idx = batchIdx+1, idx+1 {
			waitGroup.Add(1)
			startingWorkflowChannel[idx] = make(chan *RunningWorkflow)
			go e.startWorkflowDaemon(monitorExecution, startWorkflowRequests[idx], startingWorkflowChannel[idx], &waitGroup)
		}
		waitGroup.Wait()
	}
	startedWorkflows := make([]*RunningWorkflow, amount)
	for i := 0; i < amount; i += 1 {
		startedWorkflows[i] = <-startingWorkflowChannel[i]
	}
	log.Debug(fmt.Sprintf("Started %d workflows", amount))
	return startedWorkflows
}

// WaitForWorkflowCompletionUntilTimeout Helper method to wait on the channel until the timeout for the workflow execution to complete
func WaitForWorkflowCompletionUntilTimeout(executionChannel WorkflowExecutionChannel, timeout time.Duration) (workflow *model.Workflow, err error) {
	select {
	case workflow, ok := <-executionChannel:
		if !ok {
			return nil, fmt.Errorf("channel closed")
		}
		return workflow, nil
	case <-time.After(timeout):
		return nil, fmt.Errorf("timeout")
	}
}

// WaitForRunningWorkflowUntilTimeout Helper method to wait for running workflows until the timeout for the workflow execution to complete
func (e *WorkflowExecutor) WaitForRunningWorkflowsUntilTimeout(timeout time.Duration, runningWorkflows ...*RunningWorkflow) {
	for idx := 0; idx < len(runningWorkflows); {
		var waitGroup sync.WaitGroup
		for batchIdx := 0; idx < len(runningWorkflows) && batchIdx < e.waitForWorkflowBatchSize; batchIdx, idx = batchIdx+1, idx+1 {
			waitGroup.Add(1)
			go waitForRunningWorkflowUntilTimeoutDaemon(timeout, runningWorkflows[idx], &waitGroup)
		}
		waitGroup.Wait()
	}
}

func waitForRunningWorkflowUntilTimeoutDaemon(timeout time.Duration, runningWorkflow *RunningWorkflow, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	runningWorkflow.WaitForCompletionUntilTimeout(timeout)
}

// GetWorkflow Get workflow execution by workflow Id.  If includeTasks is set, also fetches all the task details.
// Returns nil if no workflow is found by the id
func (e *WorkflowExecutor) GetWorkflow(workflowId string, includeTasks bool) (*model.Workflow, error) {
	return e.getWorkflow(4, workflowId, includeTasks)
}

func (e *WorkflowExecutor) getWorkflow(retry int, workflowId string, includeTasks bool) (*model.Workflow, error) {
	workflow, response, err := e.workflowClient.GetExecutionStatus(
		context.Background(),
		workflowId,
		&client.WorkflowResourceApiGetExecutionStatusOpts{
			IncludeTasks: optional.NewBool(includeTasks)},
	)

	// 404s in GetWorkflow are a bit inconsistent with other errors since
	// it's using fmt.Errorf("..."). Keeping it this way to avoid breaking changes
	if response.StatusCode == 404 {
		return nil, fmt.Errorf("no such workflow by Id %s", workflowId)
	}

	if response.StatusCode > 399 && response.StatusCode < 500 && response.StatusCode != 429 {
		return nil, err
	}

	if err != nil {
		if retry < 0 {
			return nil, err
		} else {
			time.Sleep(time.Duration(4-retry) * 10 * time.Second)
			retry = retry - 1
			return e.getWorkflow(retry, workflowId, includeTasks)
		}
	}

	return &workflow, nil
}

// GetWorkflowStatus Get the status of the workflow execution.
// This is a lightweight method that returns only overall state of the workflow
func (e *WorkflowExecutor) GetWorkflowStatus(workflowId string, includeOutput bool, includeVariables bool) (*model.WorkflowState, error) {
	state, response, err := e.workflowClient.GetWorkflowState(context.Background(), workflowId, includeOutput, includeVariables)
	if response.StatusCode == 404 {
		return nil, nil
	}
	return &state, err
}

// GetByCorrelationIds Given the list of correlation ids, find and return workflows
// Returns a map with key as correlationId and value as a list of Workflows
// When IncludeClosed is set to true, the return value also includes workflows that are completed otherwise only running workflows are returned
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

// GetByCorrelationIdsAndNames Given the list of correlation ids and list of workflow names, find and return workflows
// Returns a map with key as correlationId and value as a list of Workflows
// When IncludeClosed is set to true, the return value also includes workflows that are completed otherwise only running workflows are returned
func (e *WorkflowExecutor) GetByCorrelationIdsAndNames(includeClosed bool, includeTasks bool, correlationIds []string, workflowNames []string) (map[string][]model.Workflow, error) {
	workflows, _, err := e.workflowClient.GetWorkflowsBatch(
		context.Background(),
		map[string][]string{
			"workflowNames":  workflowNames,
			"correlationIds": correlationIds,
		},
		&client.WorkflowResourceApiGetWorkflowsOpts{
			IncludeClosed: optional.NewBool(includeClosed),
			IncludeTasks:  optional.NewBool(includeTasks),
		})
	if err != nil {
		return nil, err
	}
	return workflows, nil
}

// Search searches for workflows
//
// - Start: Start index - used for pagination
//
// - Size:  Number of results to return
//
//   - Query: Query expression.  In the format FIELD = 'VALUE' or FIELD IN (value1, value2)
//     Only AND operations are supported.  e.g. workflowId IN ('a', 'b', 'c') ADN workflowType ='test_workflow'
//     AND startTime BETWEEN 1000 and 2000
//     Supported fields for Query are:workflowId,workflowType,status,startTime
//   - FreeText: Full text search.  All the workflow input, output and task outputs upto certain limit (check with your admins to find the size limit)
//     are full text indexed and can be used to search
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

// Pause the execution of a running workflow.
// Any tasks that are currently running will finish but no new tasks are scheduled until the workflow is resumed
func (e *WorkflowExecutor) Pause(workflowId string) error {
	_, err := e.workflowClient.PauseWorkflow(context.Background(), workflowId)
	if err != nil {
		return err
	}
	return nil
}

// Resume the execution of a workflow that is paused.  If the workflow is not paused, this method has no effect
func (e *WorkflowExecutor) Resume(workflowId string) error {
	_, err := e.workflowClient.ResumeWorkflow(context.Background(), workflowId)
	if err != nil {
		return err
	}
	return nil
}

// Terminate a running workflow.  Reason must be provided that is captured as the termination resaon for the workflow
func (e *WorkflowExecutor) Terminate(workflowId string, reason string) error {
	if strings.TrimSpace(workflowId) == "" {
		err := errors.New("workflow id cannot be empty when calling terminate workflow API")
		log.Error("Failed to terminate workflow: ", err.Error())
		return err
	}
	_, err := e.workflowClient.Terminate(context.Background(), workflowId,
		&client.WorkflowResourceApiTerminateOpts{Reason: optional.NewString(reason), TriggerFailureWorkflow: optional.NewBool(false)},
	)
	if err != nil {
		return err
	}
	return nil
}

func (e *WorkflowExecutor) TerminateWithFailure(workflowId string, reason string, triggerFailureWorkflow bool) error {
	if strings.TrimSpace(workflowId) == "" {
		err := errors.New("workflow id cannot be empty when calling terminate workflow API")
		log.Error("Failed to terminate workflow: ", err.Error())
		return err
	}
	_, err := e.workflowClient.Terminate(context.Background(), workflowId,
		&client.WorkflowResourceApiTerminateOpts{Reason: optional.NewString(reason), TriggerFailureWorkflow: optional.NewBool(triggerFailureWorkflow)},
	)
	if err != nil {
		return err
	}
	return nil
}

// Restart a workflow execution from the beginning with the same input.
// When called on a workflow that is not in a terminal status, this operation has no effect
// If useLatestDefinition is set, the restarted workflow fetches the latest definition from the metadata store
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
	return nil
}

// Retry a failed workflow from the last task that failed.  When called the task in the failed state is scheduled again
// and workflow moves to RUNNING status.  If resumeSubworkflowTasks is set and the last failed task was a sub-workflow
// the server restarts the subworkflow from the failed task.  If set to false, the sub-workflow is re-executed.
func (e *WorkflowExecutor) Retry(workflowId string, resumeSubworkflowTasks bool) error {
	_, err := e.workflowClient.Retry(
		context.Background(),
		workflowId,
		&client.WorkflowResourceApiRetryOpts{
			ResumeSubworkflowTasks: optional.NewBool(resumeSubworkflowTasks),
		},
	)
	if err != nil {
		return err
	}
	return nil
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
	return id, nil
}

// SkipTasksFromWorkflow Skips a given task execution from a current running workflow.
// When skipped the task's input and outputs are updated  from skipTaskRequest parameter.
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

// UpdateTask update the task with output and status.
func (e *WorkflowExecutor) UpdateTask(taskId string, workflowInstanceId string, status model.TaskResultStatus, output interface{}) error {
	taskResult, err := getTaskResultFromOutput(taskId, workflowInstanceId, output)
	if err != nil {
		return err
	}
	taskResult.Status = status
	_, _, err = e.taskClient.UpdateTask(context.Background(), taskResult)
	if err != nil {
		return err
	}
	return nil
}

// UpdateTaskByRefName Update the execution status and output of the task and status
func (e *WorkflowExecutor) UpdateTaskByRefName(taskRefName string, workflowInstanceId string, status model.TaskResultStatus, output interface{}) error {
	outputData, err := model.ConvertToMap(output)
	if err != nil {
		return err
	}

	_, _, err = e.taskClient.UpdateTaskByRefName(context.Background(), outputData, workflowInstanceId, taskRefName, string(status))
	if err != nil {
		return err
	}

	return nil
}

// GetTask by task Id returns nil if no such task is found by the id
func (e *WorkflowExecutor) GetTask(taskId string) (task *model.Task, err error) {
	t, _, err := e.taskClient.GetTask(context.Background(), taskId)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

// RemoveWorkflow Remove workflow execution permanently from the system
// Returns nil if no workflow is found by the id
func (e *WorkflowExecutor) RemoveWorkflow(workflowId string) error {
	_, err := e.workflowClient.Delete(context.Background(), workflowId, &client.WorkflowResourceApiDeleteOpts{ArchiveWorkflow: optional.NewBool(false)})
	if err != nil {
		return err
	}

	return nil
}

// DeleteQueueConfiguration Delete queue configuration permanently from the system
// Returns nil if no error occurred
func (e *WorkflowExecutor) DeleteQueueConfiguration(queueConfiguration queue.QueueConfiguration) (*http.Response, error) {
	return e.eventClient.DeleteQueueConfig(context.Background(), queueConfiguration.QueueType, queueConfiguration.QueueName)
}

// GetQueueConfiguration Get queue configuration if present
// Returns queue configuration if present
func (e *WorkflowExecutor) GetQueueConfiguration(queueConfiguration queue.QueueConfiguration) (map[string]interface{}, *http.Response, error) {
	return e.eventClient.GetQueueConfig(context.Background(), queueConfiguration.QueueType, queueConfiguration.QueueName)
}

// GetQueueConfiguration Create or update a queue configuration
// Returns nil if no error occurred
func (e *WorkflowExecutor) PutQueueConfiguration(queueConfiguration queue.QueueConfiguration) (*http.Response, error) {
	body, err := queueConfiguration.GetConfiguration()
	if err != nil {
		return nil, err
	}
	return e.eventClient.PutQueueConfig(context.Background(), body, queueConfiguration.QueueType, queueConfiguration.QueueName)
}

func getTaskResultFromOutput(taskId string, workflowInstanceId string, taskExecutionOutput interface{}) (*model.TaskResult, error) {
	taskResult, ok := taskExecutionOutput.(*model.TaskResult)
	if !ok {
		taskResult = model.NewTaskResult(taskId, workflowInstanceId)
		outputData, err := model.ConvertToMap(taskExecutionOutput)
		if err != nil {
			return nil, err
		}
		taskResult.OutputData = outputData
		taskResult.Status = model.CompletedTask
	}
	return taskResult, nil
}

// ExecuteWorkflow Executes a workflow
// Returns workflow Id for the newly started workflow
func (e *WorkflowExecutor) executeWorkflow(workflow *model.WorkflowDef, request *model.StartWorkflowRequest) (workflowId string, err error) {
	startWorkflowRequest := model.StartWorkflowRequest{
		Name:                            request.Name,
		Version:                         request.Version,
		CorrelationId:                   request.CorrelationId,
		Input:                           request.Input,
		TaskToDomain:                    request.TaskToDomain,
		ExternalInputPayloadStoragePath: request.ExternalInputPayloadStoragePath,
		Priority:                        request.Priority,
	}
	if workflow != nil {
		startWorkflowRequest.WorkflowDef = workflow
	}
	workflowId, response, err := e.workflowClient.StartWorkflowWithRequest(
		context.Background(),
		startWorkflowRequest,
	)
	if err != nil {
		log.Debug(
			"Failed to start workflow",
			", reason: ", err.Error(),
			", name: ", request.Name,
			", version: ", request.Version,
			", input: ", request.Input,
			", workflowId: ", workflowId,
			", response: ", response,
		)
		return "", err
	}
	log.Debug(
		"Started workflow",
		", workflowId: ", workflowId,
		", name: ", request.Name,
		", version: ", request.Version,
		", input: ", request.Input,
	)
	return workflowId, err
}

func (e *WorkflowExecutor) startWorkflowDaemon(monitorExecution bool, request *model.StartWorkflowRequest, runningWorkflowChannel chan *RunningWorkflow, waitGroup *sync.WaitGroup) {
	defer concurrency.HandlePanicError("start_workflow")
	workflowId, err := e.executeWorkflow(nil, request)
	waitGroup.Done()
	if err != nil {
		runningWorkflowChannel <- NewRunningWorkflow("", nil, err)
		return
	}
	if !monitorExecution {
		runningWorkflowChannel <- NewRunningWorkflow(workflowId, nil, nil)
		return
	}
	executionChannel, err := e.workflowMonitor.generateWorkflowExecutionChannel(workflowId)
	if err != nil {
		runningWorkflowChannel <- NewRunningWorkflow(workflowId, nil, err)
		return
	}
	runningWorkflowChannel <- NewRunningWorkflow(workflowId, executionChannel, nil)
}

func getEnvStr(key string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return v, fmt.Errorf("env not set: %s", key)
	}
	return v, nil
}

func getEnvInt(key string) (int, error) {
	s, err := getEnvStr(key)
	if err != nil {
		return 0, err
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return v, nil
}
