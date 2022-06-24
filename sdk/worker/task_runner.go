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
	maxAllowedWorkersByTaskName map[string]int
	runningWorkersByTaskName    map[string]int
	mutex                       sync.Mutex
	workerWaitGroup             sync.WaitGroup
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
		maxAllowedWorkersByTaskName: make(map[string]int),
		runningWorkersByTaskName:    make(map[string]int),
	}
}

// StartWorkerWithDomain
//  - taskName Task name to poll and execute the work
//  - executeFunction Task execution function
//  - batchSize Amount of tasks to be polled. Each polled task will be executed and updated within its own unique goroutine.
//  - pollInterval Time to wait for between polls if there are no tasks available. Reduces excessive polling on the server when there is no work
//  - domain Task domain. Optional for polling
func (c *TaskRunner) StartWorkerWithDomain(taskName string, executeFunction model.ExecuteTaskFunction, threadCount int, pollInterval time.Duration, domain string) error {
	return c.startWorker(taskName, executeFunction, threadCount, pollInterval, domain)
}

// StartWorker
//  - taskName Task name to poll and execute the work
//  - executeFunction Task execution function
//  - batchSize Amount of tasks to be polled. Each polled task will be executed and updated within its own unique goroutine.
//  - pollInterval Time to wait for between polls if there are no tasks available. Reduces excessive polling on the server when there is no work
func (c *TaskRunner) StartWorker(taskName string, executeFunction model.ExecuteTaskFunction, batchSize int, pollInterval time.Duration) error {
	return c.startWorker(taskName, executeFunction, batchSize, pollInterval, "")
}

func (c *TaskRunner) RemoveWorker(taskName string, threadCount int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if threadCount >= c.maxAllowedWorkersByTaskName[taskName] {
		c.maxAllowedWorkersByTaskName[taskName] = 0
	} else {
		c.maxAllowedWorkersByTaskName[taskName] -= threadCount
	}
	log.Debug(
		"Decreased workers for task: ", taskName,
		", by: ", threadCount,
	)
	return nil
}

func (c *TaskRunner) WaitWorkers() {
	c.workerWaitGroup.Wait()
}

func (c *TaskRunner) startWorker(taskName string, executeFunction model.ExecuteTaskFunction, threadCount int, pollInterval time.Duration, taskDomain string) error {
	var domain optional.String
	if taskDomain != "" {
		domain = optional.NewString(taskDomain)
	}
	previousMaxAllowedWorkers, err := c.getMaxAllowedWorkers(taskName)
	if err != nil {
		return err
	}
	err = c.increaseMaxAllowedWorkers(taskName, threadCount)
	if err != nil {
		return err
	}
	if previousMaxAllowedWorkers == 0 {
		c.workerWaitGroup.Add(1)
		go c.pollAndExecute(taskName, executeFunction, pollInterval, domain)
	}
	log.Info(
		fmt.Sprintf(
			"Started %d worker(s) for taskName %s, polling in interval of %d ms",
			threadCount,
			taskName,
			pollInterval.Milliseconds(),
		),
	)
	return nil
}

func (c *TaskRunner) pollAndExecute(taskName string, executeFunction model.ExecuteTaskFunction, pollInterval time.Duration, domain optional.String) error {
	defer concurrency.HandlePanicError("poll_and_execute")
	for c.isWorkerAlive(taskName) {
		isTaskQueueEmpty, err := c.runBatch(taskName, executeFunction, pollInterval, domain)
		if err != nil {
			log.Warning(
				"Failed to poll and execute",
				", reason: ", err.Error(),
				", taskName: ", taskName,
				", pollInterval: ", pollInterval.Milliseconds(), " ms",
				", domain: ", domain,
			)
		} else if isTaskQueueEmpty {
			log.Debug("No tasks available for: ", taskName)
			time.Sleep(pollInterval)
		}
	}
	c.workerWaitGroup.Done()
	return nil
}

func (c *TaskRunner) runBatch(taskName string, executeFunction model.ExecuteTaskFunction, pollInterval time.Duration, domain optional.String) (bool, error) {
	batchSize, err := c.getAvailableWorkerAmount(taskName)
	if err != nil {
		return false, err
	}
	if batchSize < 1 {
		// TODO wait until there is available workers
		time.Sleep(1 * time.Millisecond)
		return false, nil
	}
	tasks, err := c.batchPoll(taskName, batchSize, pollInterval, domain)
	if err != nil {
		return false, err
	}
	if len(tasks) < 1 {
		return true, nil
	}
	c.increaseRunningWorkers(taskName, len(tasks))
	for _, task := range tasks {
		go c.executeAndUpdateTask(taskName, task, executeFunction)
	}
	return false, nil
}

func (c *TaskRunner) executeAndUpdateTask(taskName string, task model.Task, executeFunction model.ExecuteTaskFunction) error {
	defer concurrency.HandlePanicError("execute_and_update_task")
	taskResult, err := c.executeTask(&task, executeFunction)
	if err != nil {
		metrics.IncrementTaskExecuteError(
			taskName, err,
		)
		return err
	}
	err = c.updateTaskWithRetry(taskName, taskResult)
	if err != nil {
		return err
	}
	return c.runningWorkerDone(taskName)
}

func (c *TaskRunner) batchPoll(taskName string, count int, timeout time.Duration, domain optional.String) ([]model.Task, error) {
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
			Domain:   domain,
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
		return nil, err
	}
	taskResult, err := model.GetTaskResultFromTaskExecutionOutput(t, taskExecutionOutput)
	if err != nil {
		return nil, err
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
	for attempt := 0; attempt < taskUpdateRetryAttemptsLimit; attempt += 1 {
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
	c.mutex.Lock()
	defer c.mutex.Unlock()
	amount, ok := c.maxAllowedWorkersByTaskName[taskName]
	if !ok {
		return 0, nil
	}
	return amount, nil
}

func (c *TaskRunner) getRunningWorkers(taskName string) (int, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	amount, ok := c.runningWorkersByTaskName[taskName]
	if !ok {
		return 0, nil
	}
	return amount, nil
}

func (c *TaskRunner) isWorkerAlive(taskName string) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	allowed, ok := c.maxAllowedWorkersByTaskName[taskName]
	return ok && allowed > 0
}

func (c *TaskRunner) increaseRunningWorkers(taskName string, amount int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.runningWorkersByTaskName[taskName] += amount
	log.Trace("Increased running workers for task: ", taskName, ", by: ", amount)
	return nil
}

func (c *TaskRunner) runningWorkerDone(taskName string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.runningWorkersByTaskName[taskName] -= 1
	log.Trace("Running worker done for task: ", taskName)
	return nil
}

func (c *TaskRunner) increaseMaxAllowedWorkers(taskName string, threadCount int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.maxAllowedWorkersByTaskName[taskName] += threadCount
	log.Debug("Increased max allowed workers of task: ", taskName, ", by: ", threadCount)
	return nil
}
