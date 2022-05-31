package worker

import (
	"fmt"
	"time"

	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/pkg/concurrency"
	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/conductor-sdk/conductor-go/pkg/model"
	log "github.com/sirupsen/logrus"
)

type AvailableWorkerChannel chan int

const (
	availableWorkerChannelTimeout = 30 * time.Second
)

func startWorkerManager(taskType string, executeFunction model.TaskExecuteFunction, pollInterval time.Duration, domain optional.String, conductorClient conductor_http_client.TaskResourceApiService) (AvailableWorkerChannel, error) {
	availableWorkerChannel := make(AvailableWorkerChannel)
	go manageWorkers(availableWorkerChannel, taskType, executeFunction, pollInterval, domain, conductorClient)
	return availableWorkerChannel, nil
}

func manageWorkers(availableWorkerChannel AvailableWorkerChannel, taskType string, executeFunction model.TaskExecuteFunction, pollInterval time.Duration, domain optional.String, conductorClient conductor_http_client.TaskResourceApiService) {
	defer concurrency.HandlePanicError("poll_and_execute")
	logMessage := fmt.Sprint(
		", taskType: ", taskType,
		", pollInterval: ", pollInterval.Milliseconds(),
		"ms, domain: ", domain,
	)
	for {
		availableWorkers, err := getAvailableWorkers(availableWorkerChannel)
		if err != nil {
			log.Warning("Failed to get available workers, reason: ", err.Error(), logMessage)
			break
		}
		err = manageAvailableWorkers(availableWorkerChannel, taskType, availableWorkers, executeFunction, pollInterval, domain, conductorClient)
		if err != nil {
			log.Warning("Failed to poll and execute, reason: ", err.Error(), logMessage)
			break
		}
	}
}

func getAvailableWorkers(availableWorkerChannel AvailableWorkerChannel) (int, error) {
	select {
	case availableWorkers, ok := <-availableWorkerChannel:
		if !ok {
			return 0, fmt.Errorf("available worker channel closed")
		}
		return availableWorkers, nil
	case <-time.After(availableWorkerChannelTimeout):
		return 0, fmt.Errorf("timeout waiting for available worker")
	}
}

func manageAvailableWorkers(availableWorkerChannel AvailableWorkerChannel, taskType string, batchSize int, executeFunction model.TaskExecuteFunction, pollInterval time.Duration, domain optional.String, conductorClient conductor_http_client.TaskResourceApiService) error {
	tasks, err := batchPoll(taskType, batchSize, pollInterval, domain, conductorClient)
	if err != nil {
		return err
	}
	if len(tasks) < 1 {
		log.Debug("No tasks available for: ", taskType)
		time.Sleep(pollInterval)
		return nil
	}
	for _, task := range tasks {
		go executeAndUpdateTask(taskType, task, executeFunction, conductorClient, availableWorkerChannel)
	}
	return nil
}

func executeAndUpdateTask(taskType string, task http_model.Task, executeFunction model.TaskExecuteFunction, conductorClient conductor_http_client.TaskResourceApiService, availableWorkerChannel AvailableWorkerChannel) {
	defer concurrency.HandlePanicError("execute_and_update_task")
	taskResult, err := executeTask(&task, executeFunction)
	if err != nil {
		log.Warning(
			"Failed to execute task: ",
			", taskType: ", taskType,
			", reason: ", err.Error(),
		)
		return
	}
	err = updateTaskWithRetry(taskType, taskResult, conductorClient)
	if err != nil {
		log.Warning(
			"Failed to update task: ",
			", taskType: ", taskType,
			", reason: ", err.Error(),
		)
		return
	}
	availableWorkerChannel <- 1
}
