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
	"fmt"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/concurrency"
	"github.com/conductor-sdk/conductor-go/sdk/metrics"
	"github.com/conductor-sdk/conductor-go/sdk/model"

	log "github.com/sirupsen/logrus"
)

const batchPollErrorRetryInterval = 100 * time.Millisecond

type TaskWorker struct {
	taskClient       *client.TaskResourceApiService
	workerProperties *WorkerProperties
}

func NewTaskWorker(apiClient *client.APIClient, workerProperties *WorkerProperties) *TaskWorker {
	return &TaskWorker{
		taskClient: &client.TaskResourceApiService{
			APIClient: apiClient,
		},
		workerProperties: workerProperties,
	}
}

func (tw *TaskWorker) StartWorker() {
	go tw.work4ever()
}

func (p *TaskWorker) work4ever() {
	defer concurrency.HandlePanicError("work4ever")
	for !p.workerProperties.IsPaused() {
		err := p.workOnce()
		if err != nil {
			log.Debug(
				"transient error on poll and execute",
				", reason: ", err.Error(),
				", taskName: ", p.workerProperties.TaskName,
				", domain: ", p.workerProperties.GetTaskDomain(),
			)
		}
	}
}

func (p *TaskWorker) workOnce() error {
	if p.workerProperties.IsPaused() {
		time.Sleep(batchPollErrorRetryInterval)
		return nil
	}
	batchSize := p.workerProperties.GetAvailableWorkers()
	if batchSize < 1 {
		return fmt.Errorf("no available worker, response value: %d", batchSize)
	}
	tasks, err := BatchPoll(
		p.workerProperties.TaskName,
		batchSize,
		p.workerProperties.GetTaskDomain(),
		p.workerProperties.GetPollInterval(),
		p.taskClient,
	)
	if err != nil {
		log.Debug(
			"Failed to poll tasks for taskName: ", p.workerProperties.TaskName,
			", reason: ", err.Error(),
		)
		time.Sleep(batchPollErrorRetryInterval)
		return err
	}
	if len(tasks) < 1 {
		log.Debug("No tasks available for: ", p.workerProperties.TaskName)
		time.Sleep(p.workerProperties.GetPollInterval())
		return nil
	}
	for _, task := range tasks {
		go p.executeAndUpdateTask(task)
	}
	return nil
}

func (p *TaskWorker) executeAndUpdateTask(task model.Task) error {
	p.workerProperties.IncrementRunningWorker()
	defer p.runningWorkerDone(taskName)
	defer concurrency.HandlePanicError("execute_and_update_task " + string(task.TaskId) + ": " + string(task.Status))
	taskResult, err := ExecuteTask(&task, executeFunction)
	if err != nil {
		metrics.IncrementTaskExecuteError(
			taskName, err,
		)
		return err
	}
	err = UpdateTaskWithRetry(taskName, taskResult, p.taskClient)
	return err
}

func (p *TaskWorker) getAvailableWorkerAmount(taskName string) (int, error) {
	allowed, err := p.getMaxAllowedWorkers(taskName)
	if err != nil {
		return -1, err
	}
	running, err := p.getRunningWorkers(taskName)
	if err != nil {
		return -1, err
	}
	return allowed - running, nil
}
