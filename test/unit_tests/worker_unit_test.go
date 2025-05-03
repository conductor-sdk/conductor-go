//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package unit_tests

import (
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
	"github.com/conductor-sdk/conductor-go/sdk/worker"
	"github.com/stretchr/testify/assert"
)

func TestSimpleTaskRunner(t *testing.T) {
	taskRunner := worker.NewTaskRunner(nil, nil)
	if taskRunner == nil {
		t.Fail()
	}
}

func TestTaskRunnerWithoutAuthenticationSettings(t *testing.T) {
	apiClient := client.NewAPIClient(
		nil,
		settings.NewHttpDefaultSettings(),
	)
	taskRunner := worker.NewTaskRunnerWithApiClient(apiClient)
	if taskRunner == nil {
		t.Fail()
	}
}

func TestTaskRunnerWithAuthenticationSettings(t *testing.T) {
	authenticationSettings := settings.NewAuthenticationSettings(
		"keyId",
		"keySecret",
	)
	apiClient := client.NewAPIClient(
		authenticationSettings,
		settings.NewHttpDefaultSettings(),
	)
	taskRunner := worker.NewTaskRunnerWithApiClient(apiClient)
	if taskRunner == nil {
		t.Fail()
	}
}
func TestPauseResume(t *testing.T) {
	authenticationSettings := settings.NewAuthenticationSettings(
		"keyId",
		"keySecret",
	)
	apiClient := client.NewAPIClient(
		authenticationSettings,
		settings.NewHttpDefaultSettings(),
	)
	taskRunner := worker.NewTaskRunnerWithApiClient(apiClient)
	taskRunner.StartWorker("test", TaskWorker, 21, time.Second)
	taskRunner.Pause("test")
	assert.Equal(t, 21, taskRunner.GetBatchSizeForTask("test"))
	taskRunner.Resume("test")
	assert.Equal(t, 21, taskRunner.GetBatchSizeForTask("test"))

}

func TestShutown(t *testing.T) {
	authenticationSettings := settings.NewAuthenticationSettings(
		"keyId",
		"keySecret",
	)
	apiClient := client.NewAPIClient(
		authenticationSettings,
		settings.NewHttpDefaultSettings(),
	)
	taskRunner := worker.NewTaskRunnerWithApiClient(apiClient)
	taskRunner.StartWorker("test_shutdown1", TaskWorker, 4, time.Second)
	taskRunner.StartWorker("test_shutdown2", TaskWorker, 4, time.Second)

	start := time.Now()
	go func() {
		time.Sleep(3 * time.Second)
		taskRunner.Shutdown("test_shutdown1")
		taskRunner.Shutdown("test_shutdown2")
	}()

	taskRunner.WaitWorkers()
	elapsed := time.Since(start)
	assert.GreaterOrEqual(t, elapsed.Seconds(), 2.9)

	assert.Equal(t, 0, taskRunner.GetBatchSizeForTask("test_shutdown1"))
	assert.Equal(t, 0, taskRunner.GetBatchSizeForTask("test_shutdown2"))

	err := taskRunner.IncreaseBatchSize("test_shutdown1", 1)
	assert.NotNil(t, err)
	assert.Equal(t, "no worker registered for taskName: test_shutdown1", err.Error())

	err = taskRunner.IncreaseBatchSize("test_shutdown2", 1)
	assert.NotNil(t, err)
	assert.Equal(t, "no worker registered for taskName: test_shutdown2", err.Error())

	pollInteval, err := taskRunner.GetPollIntervalForTask("test_shutdown1")
	assert.Equal(t, time.Duration(0), pollInteval)
	assert.Equal(t, "poll interval not registered for task: test_shutdown1", err.Error())

	pollInteval, err = taskRunner.GetPollIntervalForTask("test_shutdown2")
	assert.Equal(t, time.Duration(0), pollInteval)
	assert.Equal(t, "poll interval not registered for task: test_shutdown2", err.Error())
}

func TaskWorker(task *model.Task) (interface{}, error) {
	return map[string]interface{}{
		"zip": "10121",
	}, nil
}

func TestTaskRunnerTimeoutSettings(t *testing.T) {
	apiClient := client.NewAPIClient(
		nil,
		settings.NewHttpDefaultSettings(),
	)
	taskRunner := worker.NewTaskRunnerWithApiClient(apiClient)
	if taskRunner == nil {
		t.Fail()
	}

	// (1) default value should be negative
	defaultTimeout := -1 * time.Millisecond
	assert.Equal(t, defaultTimeout, taskRunner.GetPollTimeout())
	taskTimeout, err := taskRunner.GetPollTimeoutForTask("le_task")
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, defaultTimeout, taskTimeout)

	// (2) setting the global timeout should apply to all tasks
	timeout200 := 200 * time.Millisecond
	taskRunner.SetPollTimeout(timeout200)
	assert.Equal(t, timeout200, taskRunner.GetPollTimeout())

	taskTimeout, err = taskRunner.GetPollTimeoutForTask("le_task")
	assert.Nil(t, err)
	assert.Equal(t, timeout200, taskTimeout)

	taskTimeout, err = taskRunner.GetPollTimeoutForTask("another_task")
	assert.Nil(t, err)
	assert.Equal(t, timeout200, taskTimeout)

	// (3) changing the timeout for one task only affects that task
	timeout100 := 100 * taskTimeout
	taskRunner.SetPollTimeoutForTask("le_task", timeout100)

	assert.Equal(t, timeout200, taskRunner.GetPollTimeout())

	taskTimeout, err = taskRunner.GetPollTimeoutForTask("le_task")
	assert.Nil(t, err)
	assert.Equal(t, timeout100, taskTimeout)

	taskTimeout, err = taskRunner.GetPollTimeoutForTask("another_task")
	assert.Nil(t, err)
	assert.Equal(t, timeout200, taskTimeout)
}
