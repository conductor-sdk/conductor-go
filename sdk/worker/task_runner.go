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
const batchPollErrorRetryInterval = 100 * time.Millisecond
const batchPollNoAvailableWorkerRetryInterval = 1 * time.Millisecond

var hostname, _ = os.Hostname()

//TaskRunner Runner for the Task Workers.  Task Runners implements the polling and execution logic for the workers
type TaskRunner struct {
	conductorTaskResourceClient *client.TaskResourceApiService

	workerWaitGroup sync.WaitGroup

	batchSizeByTaskNameMutex sync.RWMutex
	batchSizeByTaskName      map[string]int

	runningWorkersByTaskNameMutex sync.RWMutex
	runningWorkersByTaskName      map[string]int

	pollIntervalByTaskNameMutex sync.RWMutex
	pollIntervalByTaskName      map[string]time.Duration

	pausedWorkersMutex sync.RWMutex
	pausedWorkers      map[string]bool
}

func NewTaskRunner(authenticationSettings *settings.AuthenticationSettings, httpSettings *settings.HttpSettings) *TaskRunner {
	apiClient := client.NewAPIClient(
		authenticationSettings,
		httpSettings,
	)
	return NewTaskRunnerWithApiClient(apiClient)
}

func NewTaskRunnerWithApiClient(
	apiClient *client.APIClient,
) *TaskRunner {
	return &TaskRunner{
		conductorTaskResourceClient: &client.TaskResourceApiService{
			APIClient: apiClient,
		},
		batchSizeByTaskName:      make(map[string]int),
		runningWorkersByTaskName: make(map[string]int),
		pollIntervalByTaskName:   make(map[string]time.Duration),
		pausedWorkers:            make(map[string]bool),
	}
}

// StartWorkerWithDomain
//  - taskName Task name to poll and execute the work
//  - executeFunction Task execution function
//  - batchSize Amount of tasks to be polled. Each polled task will be executed and updated within its own unique goroutine.
//  - pollInterval Time to wait for between polls if there are no tasks available. Reduces excessive polling on the server when there is no work
//  - domain Task domain. Optional for polling
func (c *TaskRunner) StartWorkerWithDomain(taskName string, executeFunction model.ExecuteTaskFunction, batchSize int, pollInterval time.Duration, domain string) error {
	return c.startWorker(taskName, executeFunction, batchSize, pollInterval, domain)
}

// StartWorker
//  - taskName Task name to poll and execute the work
//  - executeFunction Task execution function
//  - batchSize Amount of tasks to be polled. Each polled task will be executed and updated within its own unique goroutine.
//  - pollInterval Time to wait for between polls if there are no tasks available. Reduces excessive polling on the server when there is no work
func (c *TaskRunner) StartWorker(taskName string, executeFunction model.ExecuteTaskFunction, batchSize int, pollInterval time.Duration) error {
	return c.startWorker(taskName, executeFunction, batchSize, pollInterval, "")
}

func (c *TaskRunner) SetBatchSize(taskName string, batchSize int) error {
	if batchSize < 0 {
		return fmt.Errorf("batchSize can not be negative")
	}
	if !c.isWorkerRegistered(taskName) {
		return fmt.Errorf("no worker registered for taskName: %s", taskName)
	}
	c.batchSizeByTaskNameMutex.Lock()
	defer c.batchSizeByTaskNameMutex.Unlock()
	previous := c.batchSizeByTaskName[taskName]
	c.batchSizeByTaskName[taskName] = batchSize
	log.Debug(
		"Set batchSize for task: ", taskName,
		", from: ", previous,
		", to: ", c.batchSizeByTaskName[taskName],
	)
	if batchSize == 0 {
		log.Info("Stopped worker for task: ", taskName)
	} else if previous == 0 && c.batchSizeByTaskName[taskName] > 0 {
		log.Info("Started worker for task: ", taskName)
	}
	return nil
}

func (c *TaskRunner) IncreaseBatchSize(taskName string, batchSize int) error {
	if batchSize < 1 {
		return fmt.Errorf("batchSize value must be positive")
	}
	if !c.isWorkerRegistered(taskName) {
		return fmt.Errorf("no worker registered for taskName: %s", taskName)
	}
	c.batchSizeByTaskNameMutex.Lock()
	defer c.batchSizeByTaskNameMutex.Unlock()
	previous := c.batchSizeByTaskName[taskName]
	c.batchSizeByTaskName[taskName] += batchSize
	log.Debug(
		"Increased batchSize for task: ", taskName,
		", from: ", previous,
		", to: ", c.batchSizeByTaskName[taskName],
	)
	if previous == 0 {
		log.Info("Started worker for task: ", taskName)
	}
	return nil
}

func (c *TaskRunner) DecreaseBatchSize(taskName string, batchSize int) error {
	if batchSize < 1 {
		return fmt.Errorf("batchSize value must be positive")
	}
	if !c.isWorkerRegistered(taskName) {
		return fmt.Errorf("no worker registered for taskName: %s", taskName)
	}
	c.batchSizeByTaskNameMutex.Lock()
	defer c.batchSizeByTaskNameMutex.Unlock()
	previous := c.batchSizeByTaskName[taskName]
	c.batchSizeByTaskName[taskName] -= batchSize
	log.Debug(
		"Decreased batchSize for task: ", taskName,
		", from: ", previous,
		", to: ", c.batchSizeByTaskName[taskName],
	)
	if previous-batchSize <= 0 {
		c.batchSizeByTaskName[taskName] = 0
		log.Info("Stopped worker for task: ", taskName)
	}
	return nil
}

// Pause a running worker.  When paused worker will not poll for new task.  Worker must be resumed using Resume
func (c *TaskRunner) Pause(taskName string) {
	c.pausedWorkersMutex.Lock()
	defer c.pausedWorkersMutex.Unlock()
	c.pausedWorkers[taskName] = true
}

// Resume a running worker.  If the worker is not paused, calling this method has no impact
func (c *TaskRunner) Resume(taskName string) {
	c.pausedWorkersMutex.Lock()
	defer c.pausedWorkersMutex.Unlock()
	c.pausedWorkers[taskName] = false
}

func (c *TaskRunner) isPaused(taskName string) bool {
	c.pausedWorkersMutex.RLock()
	defer c.pausedWorkersMutex.RUnlock()
	return c.pausedWorkers[taskName]
}
func (c *TaskRunner) WaitWorkers() {
	c.workerWaitGroup.Wait()
}

func (c *TaskRunner) startWorker(taskName string, executeFunction model.ExecuteTaskFunction, batchSize int, pollInterval time.Duration, taskDomain string) error {
	c.SetPollIntervalForTask(taskName, pollInterval)
	c.Resume(taskName)
	previousMaxAllowedWorkers, err := c.getMaxAllowedWorkers(taskName)
	if err != nil {
		return err
	}
	err = c.increaseMaxAllowedWorkers(taskName, batchSize)
	if err != nil {
		return err
	}
	if previousMaxAllowedWorkers < 1 {
		c.workerWaitGroup.Add(1)
		go c.pollAndExecute(taskName, executeFunction, taskDomain)
	}
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

func (c *TaskRunner) pollAndExecute(taskName string, executeFunction model.ExecuteTaskFunction, domain string) {
	defer c.workerWaitGroup.Done()
	defer concurrency.HandlePanicError("poll_and_execute")
	for c.isWorkerRegistered(taskName) {
		err := c.runBatch(taskName, executeFunction, domain)
		if err != nil {
			log.Debug(
				"transient error when poll and execute",
				", reason: ", err.Error(),
				", taskName: ", taskName,
				", domain: ", domain,
			)
		}
	}
}

func (c *TaskRunner) runBatch(taskName string, executeFunction model.ExecuteTaskFunction, domain string) error {
	batchSize, err := c.getAvailableWorkerAmount(taskName)
	if err != nil {
		return err
	}

	if batchSize < 1 || c.isPaused(taskName) {
		time.Sleep(batchPollNoAvailableWorkerRetryInterval)
		return nil
	}
	tasks, err := c.batchPoll(taskName, batchSize, domain)
	if err != nil {
		time.Sleep(batchPollErrorRetryInterval)
		return err
	}
	if len(tasks) < 1 {
		log.Debug("No tasks available for: ", taskName)
		pollInterval, err := c.GetPollIntervalForTask(taskName)
		if err != nil {
			return err
		}
		time.Sleep(pollInterval)
		return nil
	}
	c.increaseRunningWorkers(taskName, len(tasks))
	for _, task := range tasks {
		go c.executeAndUpdateTask(taskName, task, executeFunction)
	}
	return nil
}

func (c *TaskRunner) executeAndUpdateTask(taskName string, task model.Task, executeFunction model.ExecuteTaskFunction) error {
	defer c.runningWorkerDone(taskName)
	defer concurrency.HandlePanicError("execute_and_update_task")
	taskResult, err := c.executeTask(&task, executeFunction)
	if err != nil {
		metrics.IncrementTaskExecuteError(
			taskName, err,
		)
		return err
	}
	err = c.updateTaskWithRetry(taskName, taskResult)
	return err
}

func (c *TaskRunner) batchPoll(taskName string, count int, domain string) ([]model.Task, error) {
	timeout, err := c.GetPollIntervalForTask(taskName)
	if err != nil {
		return nil, fmt.Errorf("failed to get poll interval for task %s, reason: %s", taskName, err.Error())
	}
	var domainOptional optional.String
	if domain != "" {
		domainOptional = optional.NewString(domain)
	}
	log.Debug(
		"Polling for task: ", taskName,
		", in batches of size: ", count,
	)
	metrics.IncrementTaskPoll(taskName)
	startTime := time.Now()
	tasks, response, err := c.conductorTaskResourceClient.BatchPoll(
		context.Background(),
		taskName,
		&client.TaskResourceApiBatchPollOpts{
			Domain:   domainOptional,
			Workerid: optional.NewString(hostname),
			Count:    optional.NewInt32(int32(count)),
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

func (c *TaskRunner) executeTask(t *model.Task, executeFunction model.ExecuteTaskFunction) (*model.TaskResult, error) {
	log.Trace(
		"Executing task of type: ", t.TaskDefName,
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
		return model.NewTaskResultFromTaskWithError(t, err), nil
	}

	taskResult, err := model.GetTaskResultFromTaskExecutionOutput(t, taskExecutionOutput)
	if err != nil {
		return model.NewTaskResultFromTaskWithError(t, err), nil
	}
	log.Trace(
		"Executed task of type: ", t.TaskDefName,
		", taskId: ", t.TaskId,
		", workflowId: ", t.WorkflowInstanceId,
	)
	return taskResult, nil
}

func (c *TaskRunner) updateTaskWithRetry(taskName string, taskResult *model.TaskResult) error {
	log.Debug(
		"Updating task of type: ", taskName,
		", taskId: ", taskResult.TaskId,
		", workflowId: ", taskResult.WorkflowInstanceId,
	)
	for attempt := 0; attempt <= taskUpdateRetryAttemptsLimit; attempt += 1 {
		if attempt > 0 {
			// Wait for [10s, 20s, 30s] before next attempt
			amount := attempt * 10
			time.Sleep(time.Duration(amount) * time.Second)
		}
		response, err := c.updateTask(taskName, taskResult)
		if err == nil {
			log.Debug(
				"Updated task of type: ", taskName,
				", taskId: ", taskResult.TaskId,
				", workflowId: ", taskResult.WorkflowInstanceId,
			)
			return nil
		}
		metrics.IncrementTaskUpdateError(taskName, err)
		log.Debug(
			"Failed to update task",
			", reason: ", err.Error(),
			", task type: ", taskName,
			", taskId: ", taskResult.TaskId,
			", workflowId: ", taskResult.WorkflowInstanceId,
			", response: ", *response,
		)
	}
	return fmt.Errorf("failed to update task %s after %d attempts", taskName, taskUpdateRetryAttemptsLimit)
}

func (c *TaskRunner) updateTask(taskName string, taskResult *model.TaskResult) (*http.Response, error) {
	startTime := time.Now()
	_, response, err := c.conductorTaskResourceClient.UpdateTask(context.Background(), taskResult)
	spentTime := time.Since(startTime).Milliseconds()
	metrics.RecordTaskUpdateTime(taskName, float64(spentTime))
	return response, err
}

func (c *TaskRunner) getAvailableWorkerAmount(taskName string) (int, error) {
	allowed, err := c.getMaxAllowedWorkers(taskName)
	if err != nil {
		return -1, err
	}
	running, err := c.getRunningWorkers(taskName)
	if err != nil {
		return -1, err
	}
	return allowed - running, nil
}

func (c *TaskRunner) getMaxAllowedWorkers(taskName string) (int, error) {
	c.batchSizeByTaskNameMutex.RLock()
	defer c.batchSizeByTaskNameMutex.RUnlock()
	amount, ok := c.batchSizeByTaskName[taskName]
	if !ok {
		return 0, nil
	}
	return amount, nil
}

func (c *TaskRunner) getRunningWorkers(taskName string) (int, error) {
	c.runningWorkersByTaskNameMutex.RLock()
	defer c.runningWorkersByTaskNameMutex.RUnlock()
	amount, ok := c.runningWorkersByTaskName[taskName]
	if !ok {
		return 0, nil
	}
	return amount, nil
}

func (c *TaskRunner) isWorkerRegistered(taskName string) bool {
	c.batchSizeByTaskNameMutex.RLock()
	defer c.batchSizeByTaskNameMutex.RUnlock()
	_, ok := c.batchSizeByTaskName[taskName]
	return ok
}

func (c *TaskRunner) increaseRunningWorkers(taskName string, amount int) error {
	c.runningWorkersByTaskNameMutex.Lock()
	defer c.runningWorkersByTaskNameMutex.Unlock()
	c.runningWorkersByTaskName[taskName] += amount
	log.Trace("Increased running workers for task: ", taskName, ", by: ", amount)
	return nil
}

func (c *TaskRunner) runningWorkerDone(taskName string) error {
	c.runningWorkersByTaskNameMutex.Lock()
	defer c.runningWorkersByTaskNameMutex.Unlock()
	c.runningWorkersByTaskName[taskName] -= 1
	log.Trace("Running worker done for task: ", taskName)
	return nil
}

func (c *TaskRunner) increaseMaxAllowedWorkers(taskName string, batchSize int) error {
	c.batchSizeByTaskNameMutex.Lock()
	defer c.batchSizeByTaskNameMutex.Unlock()
	c.batchSizeByTaskName[taskName] += batchSize
	log.Debug("Increased max allowed workers of task: ", taskName, ", by: ", batchSize)
	return nil
}

func (c *TaskRunner) SetPollIntervalForTask(taskName string, pollInterval time.Duration) error {
	c.pollIntervalByTaskNameMutex.Lock()
	defer c.pollIntervalByTaskNameMutex.Unlock()
	c.pollIntervalByTaskName[taskName] = pollInterval
	log.Debug("Updated poll interval for task: ", taskName, ", to: ", pollInterval.Milliseconds(), "ms")
	return nil
}

func (c *TaskRunner) GetPollIntervalForTask(taskName string) (pollInterval time.Duration, err error) {
	c.pollIntervalByTaskNameMutex.RLock()
	defer c.pollIntervalByTaskNameMutex.RUnlock()
	pollInterval, ok := c.pollIntervalByTaskName[taskName]
	if !ok {
		return pollInterval, fmt.Errorf("poll interval not registered for task: %s", taskName)
	}
	return pollInterval, nil
}

func (c *TaskRunner) GetBatchSizeForAll() (batchSizeByTaskName map[string]int) {
	c.batchSizeByTaskNameMutex.RLock()
	defer c.batchSizeByTaskNameMutex.RUnlock()
	batchSizeByTaskName = make(map[string]int)
	for taskName, batchSize := range c.batchSizeByTaskName {
		batchSizeByTaskName[taskName] = batchSize
	}
	return batchSizeByTaskName
}

func (c *TaskRunner) GetBatchSizeForTask(taskName string) (batchSize int) {
	c.batchSizeByTaskNameMutex.RLock()
	defer c.batchSizeByTaskNameMutex.RUnlock()
	batchSize, ok := c.batchSizeByTaskName[taskName]
	if !ok {
		return 0
	}
	return batchSize
}
