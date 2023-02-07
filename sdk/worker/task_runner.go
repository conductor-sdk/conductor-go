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
	"sync"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
)

// TaskRunner Runner for the Task Workers. Task Runners implements the polling and execution logic for the workers
type TaskRunner struct {
	taskWorkerByTaskName map[string]*TaskWorker
	taskWorkerMutex      sync.RWMutex

	apiClient *client.APIClient
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
		apiClient:            apiClient,
		taskWorkerByTaskName: make(map[string]*TaskWorker),
	}
}

func (c *TaskRunner) getTaskWorkerOrCreate(taskName string) *TaskWorker {
	taskWorker, err := c.getTaskWorker(taskName)
	if err != nil {
		return taskWorker
	}
	c.createTaskWorker(taskName)
	return taskWorker
}

func (c *TaskRunner) getTaskWorker(taskName string) (*TaskWorker, error) {
	c.taskWorkerMutex.RLock()
	defer c.taskWorkerMutex.RUnlock()
	taskWorker, ok := c.taskWorkerByTaskName[taskName]
	if !ok {
		return nil, fmt.Errorf("worker not found for taskName: %s", taskName)
	}
	return taskWorker, nil
}

func (c *TaskRunner) createTaskWorker(taskName string) {
	c.taskWorkerMutex.Lock()
	defer c.taskWorkerMutex.Unlock()
	c.taskWorkerByTaskName[taskName] = NewTaskWorker(
		c.apiClient,
	)
}
