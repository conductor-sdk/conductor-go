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
	"time"

	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/metrics"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	log "github.com/sirupsen/logrus"
)

const taskUpdateRetryAttemptsLimit = 3

var hostname, _ = os.Hostname()

func BatchPoll(taskName string, count int, domain string, timeout time.Duration, taskClient *client.TaskResourceApiService) ([]model.Task, error) {
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
	tasks, response, err := taskClient.BatchPoll(
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

func UpdateTaskWithRetry(taskName string, taskResult *model.TaskResult, taskClient *client.TaskResourceApiService) error {
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
		_, err := UpdateTask(taskName, taskResult, taskClient)
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
			", response: ", err,
		)
	}
	return fmt.Errorf("failed to update task %s after %d attempts", taskName, taskUpdateRetryAttemptsLimit)
}

func UpdateTask(taskName string, taskResult *model.TaskResult, taskClient *client.TaskResourceApiService) (*http.Response, error) {
	startTime := time.Now()
	_, response, err := taskClient.UpdateTask(context.Background(), taskResult)
	spentTime := time.Since(startTime).Milliseconds()
	metrics.RecordTaskUpdateTime(taskName, float64(spentTime))
	return response, err
}
