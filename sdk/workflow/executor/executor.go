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
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/concurrency"
	"github.com/conductor-sdk/conductor-go/sdk/event/queue"
	"github.com/conductor-sdk/conductor-go/sdk/model"

	log "github.com/sirupsen/logrus"
)

type WorkflowExecutor struct {
	metadataClient *client.MetadataResourceApiService
	taskClient     *client.TaskResourceApiService
	tagsClient     *client.TagsApiService
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
	tagsClient := client.TagsApiService{
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
		tagsClient:               &tagsClient,
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
	return e.RegisterWorkflowWithContext(context.Background(), overwrite, workflow)
}

// UnRegisterWorkflow Un-registers the workflow on the server.
func (e *WorkflowExecutor) UnRegisterWorkflow(name string, version int32) error {
	return e.UnRegisterWorkflowWithContext(context.Background(), name, version)
}

// ExecuteWorkflow start a workflow and wait until the workflow completes or the waitUntilTask completes
// Returns the output of the workflow
func (e *WorkflowExecutor) ExecuteWorkflow(startWorkflowRequest *model.StartWorkflowRequest, waitUntilTask string) (run *model.WorkflowRun, err error) {
	return e.ExecuteWorkflowWithContext(context.Background(), startWorkflowRequest, waitUntilTask)
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
	return e.StartWorkflowWithContext(context.Background(), startWorkflowRequest)
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

// WaitForRunningWorkflowsUntilTimeout Helper method to wait for running workflows until the timeout for the workflow execution to complete
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
	return e.GetWorkflowWithContext(context.Background(), workflowId, includeTasks)
}

// GetWorkflowStatus Get the status of the workflow execution.
// This is a lightweight method that returns only overall state of the workflow
func (e *WorkflowExecutor) GetWorkflowStatus(workflowId string, includeOutput bool, includeVariables bool) (*model.WorkflowState, error) {
	return e.GetWorkflowStatusWithContext(context.Background(), workflowId, includeOutput, includeVariables)
}

// GetByCorrelationIds Given the list of correlation ids, find and return workflows
// Returns a map with key as correlationId and value as a list of Workflows
// When IncludeClosed is set to true, the return value also includes workflows that are completed otherwise only running workflows are returned
func (e *WorkflowExecutor) GetByCorrelationIds(workflowName string, includeClosed bool, includeTasks bool, correlationIds ...string) (map[string][]model.Workflow, error) {
	return e.GetByCorrelationIdsWithContext(context.Background(), workflowName, includeClosed, includeTasks, correlationIds...)
}

// GetByCorrelationIdsAndNames Given the list of correlation ids and list of workflow names, find and return workflows
// Returns a map with key as correlationId and value as a list of Workflows
// When IncludeClosed is set to true, the return value also includes workflows that are completed otherwise only running workflows are returned
func (e *WorkflowExecutor) GetByCorrelationIdsAndNames(includeClosed bool, includeTasks bool, correlationIds []string, workflowNames []string) (map[string][]model.Workflow, error) {
	return e.GetByCorrelationIdsAndNamesWithContext(context.Background(), includeClosed, includeTasks, correlationIds, workflowNames)
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
	return e.SearchWithContext(context.Background(), start, size, query, freeText)
}

// Pause the execution of a running workflow.
// Any tasks that are currently running will finish but no new tasks are scheduled until the workflow is resumed
func (e *WorkflowExecutor) Pause(workflowId string) error {
	return e.PauseWithContext(context.Background(), workflowId)
}

// Resume the execution of a workflow that is paused.  If the workflow is not paused, this method has no effect
func (e *WorkflowExecutor) Resume(workflowId string) error {
	return e.ResumeWithContext(context.Background(), workflowId)
}

// Terminate Terminates a running workflow. Reason must be provided that is captured as the termination reason for the workflow.
func (e *WorkflowExecutor) Terminate(workflowId string, reason string) error {
	return e.TerminateWithContext(context.Background(), workflowId, reason)
}

// TerminateWithFailure Terminates a running workflow.
func (e *WorkflowExecutor) TerminateWithFailure(workflowId string, reason string, triggerFailureWorkflow bool) error {
	return e.TerminateWithFailureWithContext(context.Background(), workflowId, reason, triggerFailureWorkflow)
}

// Restart a workflow execution from the beginning with the same input.
// When called on a workflow that is not in a terminal status, this operation has no effect
// If useLatestDefinition is set, the restarted workflow fetches the latest definition from the metadata store
func (e *WorkflowExecutor) Restart(workflowId string, useLatestDefinition bool) error {
	return e.RestartWithContext(context.Background(), workflowId, useLatestDefinition)
}

// Retry a failed workflow from the last task that failed.  When called the task in the failed state is scheduled again
// and workflow moves to RUNNING status.  If resumeSubworkflowTasks is set and the last failed task was a sub-workflow
// the server restarts the subworkflow from the failed task.  If set to false, the sub-workflow is re-executed.
func (e *WorkflowExecutor) Retry(workflowId string, resumeSubworkflowTasks bool) error {
	return e.RetryWithContext(context.Background(), workflowId, resumeSubworkflowTasks)
}

// ReRun a completed workflow from a specific task (ReRunFromTaskId) and optionally change the input
// Also update the completed tasks with new input (ReRunFromTaskId) if required
func (e *WorkflowExecutor) ReRun(workflowId string, reRunRequest model.RerunWorkflowRequest) (id string, error error) {
	return e.ReRunWithContext(context.Background(), workflowId, reRunRequest)
}

// SkipTasksFromWorkflow Skips a given task execution from a current running workflow.
// When skipped the task's input and outputs are updated  from skipTaskRequest parameter.
func (e *WorkflowExecutor) SkipTasksFromWorkflow(workflowId string, taskReferenceName string, skipTaskRequest model.SkipTaskRequest) error {
	return e.SkipTasksFromWorkflowWithContext(context.Background(), workflowId, taskReferenceName, skipTaskRequest)
}

// UpdateTask update the task with output and status.
func (e *WorkflowExecutor) UpdateTask(taskId string, workflowInstanceId string, status model.TaskResultStatus, output interface{}) error {
	return e.UpdateTaskWithContext(context.Background(), taskId, workflowInstanceId, status, output)
}

// UpdateTaskByRefName Update the execution status and output of the task and status
func (e *WorkflowExecutor) UpdateTaskByRefName(taskRefName string, workflowInstanceId string, status model.TaskResultStatus, output interface{}) error {
	return e.UpdateTaskByRefNameWithContext(context.Background(), taskRefName, workflowInstanceId, status, output)
}

// GetTask by task Id returns nil if no such task is found by the id
func (e *WorkflowExecutor) GetTask(taskId string) (task *model.Task, err error) {
	return e.GetTaskWithContext(context.Background(), taskId)
}

// RemoveWorkflow Remove workflow execution permanently from the system
// Returns nil if no workflow is found by the id
func (e *WorkflowExecutor) RemoveWorkflow(workflowId string) error {
	return e.RemoveWorkflowWithContext(context.Background(), workflowId)
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

// PutQueueConfiguration Create or update a queue configuration
// Returns nil if no error occurred
func (e *WorkflowExecutor) PutQueueConfiguration(queueConfiguration queue.QueueConfiguration) (*http.Response, error) {
	return e.PutQueueConfigurationWithContext(context.Background(), queueConfiguration)
}

// ExecuteWorkflow Executes a workflow
// Returns workflow Id for the newly started workflow
func (e *WorkflowExecutor) executeWorkflow(workflow *model.WorkflowDef, request *model.StartWorkflowRequest) (workflowId string, err error) {
	return e.executeWorkflowWithContext(context.Background(), workflow, request)
}

func (e *WorkflowExecutor) AddWorkflowTags(workflowName string, tags map[string]string) error {
	return e.addWorkflowTagsWithContext(context.Background(), workflowName, tags)
}

func (e *WorkflowExecutor) GetWorkflowTags(workflowName string) (map[string]string, error) {
	return e.getWorkflowTagsWithContext(context.Background(), workflowName)
}

func (e *WorkflowExecutor) UpdateWorkflowTags(workflowName string, tags map[string]string) error {
	return e.updateWorkflowTagWithContext(context.Background(), workflowName, tags)
}

func (e *WorkflowExecutor) DeleteWorkflowTags(workflowName string, tags map[string]string) error {
	return e.deleteWorkflowTagWithContext(context.Background(), workflowName, tags)
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
