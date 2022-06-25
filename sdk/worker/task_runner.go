//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package worker

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/concurrency"
	"github.com/conductor-sdk/conductor-go/sdk/metrics"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/settings"

	"github.com/antihax/optional"
	log "github.com/sirupsen/logrus"
)

const taskUpdateRetryAttemptsLimit = 3

var hostname, _ = os.Hostname()

//TaskRunner Runner for the Task Workers.  Task Runners implements the polling and execution logic for the workers
type TaskRunner struct {
	conductorTaskResourceClient *client.TaskResourceApiService
	workerWaitGroup             sync.WaitGroup
	workerOrkestratorByTaskName map[string]*WorkerOrkestrator
	workerOrkestratorMutex      sync.RWMutex
}

func NewTaskRunner(authenticationSettings *settings.AuthenticationSettings, httpSettings *settings.HttpSettings) *TaskRunner {
	apiClient := client.NewAPIClient(
		authenticationSettings,
		httpSettings,
	)
	return NewTaskRunnerWithApiClient(apiClient)
}

func NewTaskRunnerWithApiClient(apiClient *client.APIClient) *TaskRunner {
	return &TaskRunner{
		conductorTaskResourceClient: &client.TaskResourceApiService{
			APIClient: apiClient,
		},
		workerOrkestratorByTaskName: make(map[string]*WorkerOrkestrator),
	}
}

// StartWorkerWithDomain
//  - taskName Task name to poll and execute the work
//  - executeFunction Task execution function
//  - batchSize Amount of tasks to be polled. Each polled task will be executed and updated within its own unique goroutine.
//  - pollInterval Time to wait for between polls if there are no tasks available. Reduces excessive polling on the server when there is no work
//  - domain Task domain. Optional for polling
func (tr *TaskRunner) StartWorkerWithDomain(taskName string, executeFunction model.ExecuteTaskFunction, batchSize int, pollInterval time.Duration, domain string) error {
	return tr.startWorker(taskName, executeFunction, batchSize, pollInterval, domain)
}

// StartWorker
//  - taskName Task name to poll and execute the work
//  - executeFunction Task execution function
//  - batchSize Amount of tasks to be polled. Each polled task will be executed and updated within its own unique goroutine.
//  - pollInterval Time to wait for between polls if there are no tasks available. Reduces excessive polling on the server when there is no work
func (tr *TaskRunner) StartWorker(taskName string, executeFunction model.ExecuteTaskFunction, batchSize int, pollInterval time.Duration) error {
	return tr.startWorker(taskName, executeFunction, batchSize, pollInterval, "")
}

func (tr *TaskRunner) WaitWorkers() {
	tr.workerWaitGroup.Wait()
}

func (tr *TaskRunner) startWorker(taskName string, executeFunction model.ExecuteTaskFunction, batchSize int, pollInterval time.Duration, taskDomain string) error {
	_, isWorkerRegistered := tr.getWorkerOrkestratorForTask(taskName)
	if isWorkerRegistered {
		return fmt.Errorf("worker already registered for taskName: %s", taskName)
	}
	workerOrkestrator := NewWorkerOrkestrator(
		taskName,
		batchSize,
		pollInterval,
		executeFunction,
		optional.NewString(taskDomain),
	)
	tr.registerWorkerOrkestratorForTask(taskName, workerOrkestrator)
	tr.workerWaitGroup.Add(1)
	go tr.pollAndExecuteDaemon(taskName)
	log.Info(
		fmt.Sprintf(
			"Started %d worker(s) for taskName %s, polling in interval of %d ms",
			batchSize,
			taskName,
			pollInterval.Milliseconds(),
		),
	)
	return nil
}

func (tr *TaskRunner) pollAndExecuteDaemon(taskName string) {
	defer tr.workerWaitGroup.Done()
	defer concurrency.HandlePanicError("poll_and_execute")
	for {
		workerOrkestrator, isWorkerRegistered := tr.getWorkerOrkestratorForTask(taskName)
		if !isWorkerRegistered {
			log.Warning("Stop polling for taskName: ", taskName, ", reason: no worker registered.")
			break
		}
		isTaskQueueEmpty, err := tr.pollAndExecute(workerOrkestrator)
		if err != nil {
			log.Warning("Failed to poll for task, reason: ", err.Error(), ", taskName: ", taskName)
			break
		}
		if isTaskQueueEmpty {
			log.Debug("No tasks available for: ", taskName)
			time.Sleep(workerOrkestrator.GetPollInterval())
		}
	}
}

func (tr *TaskRunner) pollAndExecute(workerOrkestrator *WorkerOrkestrator) (isTaskQueueEmpty bool, err error) {
	tasks, err := tr.batchPoll(
		workerOrkestrator.taskName,
		workerOrkestrator.GetAvailableWorkers(),
		workerOrkestrator.GetPollInterval(),
		workerOrkestrator.GetDomain(),
	)
	if err != nil {
		return false, err
	}
	if len(tasks) < 1 {
		return true, nil
	}
	workerOrkestrator.IncreaseRunningWorkers(len(tasks))
	for _, task := range tasks {
		go tr.executeAndUpdateTaskDaemon(workerOrkestrator, task)
	}
	return false, nil
}

func (tr *TaskRunner) executeAndUpdateTaskDaemon(workerOrkestrator *WorkerOrkestrator, task model.Task) {
	defer workerOrkestrator.DecreaseRunningWorker()
	defer concurrency.HandlePanicError("execute_and_update_task")
	taskResult, err := tr.executeTask(&task, workerOrkestrator.GetExecuteTaskFunction())
	if err != nil {
		metrics.IncrementTaskExecuteError(
			workerOrkestrator.taskName, err,
		)
		log.Warning(
			"Failed to execute task, reason: ", err.Error(),
			", taskName: ", workerOrkestrator.taskName,
			", taskId: ", task.TaskId,
			", workflowId: ", task.WorkflowInstanceId,
		)
		return
	}
	err = tr.updateTaskWithRetry(workerOrkestrator.taskName, taskResult)
	if err != nil {
		log.Warning(
			"Failed to update task with retry, reason: ", err.Error(),
			", taskName: ", workerOrkestrator.taskName,
			", taskId: ", task.TaskId,
			", workflowId: ", task.WorkflowInstanceId,
		)
		return
	}
}

func (tr *TaskRunner) batchPoll(taskName string, batchSize int, timeout time.Duration, domain optional.String) ([]model.Task, error) {
	if batchSize < 1 {
		// TODO wait until there is available workers
		time.Sleep(1 * time.Millisecond)
		return nil, nil
	}
	log.Debug(
		"Polling for task: ", taskName,
		", in batches of size: ", batchSize,
	)
	metrics.IncrementTaskPoll(taskName)
	startTime := time.Now()
	tasks, response, err := tr.conductorTaskResourceClient.BatchPoll(
		context.Background(),
		taskName,
		&client.TaskResourceApiBatchPollOpts{
			Domain:   domain,
			Workerid: optional.NewString(hostname),
			Count:    optional.NewInt32(int32(batchSize)),
			Timeout:  optional.NewInt32(int32(timeout.Milliseconds())),
		},
	)
	spentTime := time.Since(startTime)
	metrics.RecordTaskPollTime(
		taskName,
		spentTime.Seconds(),
	)
	if err != nil {
		metrics.IncrementTaskPollError(
			taskName, err,
		)
		return nil, err
	}
	if response.StatusCode == 204 {
		return nil, nil
	}
	log.Debug(fmt.Sprintf("Polled %d tasks for taskName: %s", len(tasks), taskName))
	return tasks, nil
}

func (tr *TaskRunner) executeTask(t *model.Task, executeFunction model.ExecuteTaskFunction) (*model.TaskResult, error) {
	log.Trace(
		"Executing task, taskName: ", t.TaskDefName,
		", taskId: ", t.TaskId,
		", workflowId: ", t.WorkflowInstanceId,
	)
	startTime := time.Now()
	taskExecutionOutput, err := executeFunction(t)
	spentTime := time.Since(startTime)
	metrics.RecordTaskExecuteTime(
		t.TaskDefName, float64(spentTime.Milliseconds()),
	)
	if err != nil {
		return nil, err
	}
	taskResult, err := model.GetTaskResultFromTaskExecutionOutput(t, taskExecutionOutput)
	if err != nil {
		return nil, err
	}
	log.Trace(
		"Executed task, taskName: ", t.TaskDefName,
		", taskId: ", t.TaskId,
		", workflowId: ", t.WorkflowInstanceId,
	)
	return taskResult, nil
}

func (tr *TaskRunner) updateTaskWithRetry(taskName string, taskResult *model.TaskResult) error {
	log.Debug(
		"Updating task, taskName: ", taskName,
		", taskId: ", taskResult.TaskId,
		", workflowId: ", taskResult.WorkflowInstanceId,
	)
	for attempt := 0; attempt < taskUpdateRetryAttemptsLimit; attempt += 1 {
		response, err := tr.updateTask(taskName, taskResult)
		if err == nil {
			log.Debug(
				"Updated task, taskName: ", taskName,
				", taskId: ", taskResult.TaskId,
				", workflowId: ", taskResult.WorkflowInstanceId,
			)
			return nil
		}
		metrics.IncrementTaskUpdateError(taskName, err)
		log.Debug(
			"Failed to update task",
			", reason: ", err.Error(),
			", taskName: ", taskName,
			", taskId: ", taskResult.TaskId,
			", workflowId: ", taskResult.WorkflowInstanceId,
			", response: ", *response,
		)
		amount := (1 << attempt)
		time.Sleep(time.Duration(amount) * time.Second)
	}
	return fmt.Errorf("failed to update task %s after %d attempts", taskName, taskUpdateRetryAttemptsLimit)
}

func (tr *TaskRunner) updateTask(taskName string, taskResult *model.TaskResult) (*http.Response, error) {
	startTime := time.Now()
	_, response, err := tr.conductorTaskResourceClient.UpdateTask(context.Background(), taskResult)
	spentTime := time.Since(startTime).Milliseconds()
	metrics.RecordTaskUpdateTime(taskName, float64(spentTime))
	return response, err
}

func (tr *TaskRunner) getWorkerOrkestratorForTask(taskName string) (workerOrkestrator *WorkerOrkestrator, isWorkerRegistered bool) {
	tr.workerOrkestratorMutex.RLock()
	defer tr.workerOrkestratorMutex.RUnlock()
	workerOrkestrator, ok := tr.workerOrkestratorByTaskName[taskName]
	return workerOrkestrator, ok
}

func (tr *TaskRunner) registerWorkerOrkestratorForTask(taskName string, workerOrkestrator *WorkerOrkestrator) {
	tr.workerOrkestratorMutex.Lock()
	defer tr.workerOrkestratorMutex.Unlock()
	tr.workerOrkestratorByTaskName[taskName] = workerOrkestrator
}
