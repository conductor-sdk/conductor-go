package worker

import (
	"sync"
	"time"

	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/pkg/concurrency"
	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/conductor-sdk/conductor-go/pkg/model"
	log "github.com/sirupsen/logrus"
)

type WorkerManager struct {
	taskType               string
	executeFunction        model.TaskExecuteFunction
	pollInterval           time.Duration
	domain                 optional.String
	conductorClient        *conductor_http_client.TaskResourceApiService
	availableWorkerMonitor *AvailableWorkerMonitor
	runningGoRoutines      sync.WaitGroup
}

func NewWorkerManager(taskType string, executeFunction model.TaskExecuteFunction, pollInterval time.Duration, domain optional.String, conductorClient *conductor_http_client.TaskResourceApiService) *WorkerManager {
	return &WorkerManager{
		taskType:               taskType,
		executeFunction:        executeFunction,
		pollInterval:           pollInterval,
		domain:                 domain,
		conductorClient:        conductorClient,
		availableWorkerMonitor: NewAvailableWorkerMonitor(taskType),
	}
}

func (w *WorkerManager) Start() error {
	w.availableWorkerMonitor.Start()
	w.runningGoRoutines.Add(1)
	go w.manageAvailableWorkersDaemon()
	log.Debug("Started worker manager for taskType: ", w.taskType)
	return nil
}

func (w *WorkerManager) Wait() error {
	log.Debug("Waiting for worker manager of taskType: ", w.taskType)
	w.availableWorkerMonitor.Wait()
	w.runningGoRoutines.Wait()
	log.Debug("Done waiting for worker manager of taskType: ", w.taskType)
	return nil
}

func (w *WorkerManager) IncreaseWorkers(amount int) error {
	return w.availableWorkerMonitor.IncreaseAvailableWorkerAmount(amount)
}

func (w *WorkerManager) manageAvailableWorkersDaemon() {
	defer w.runningGoRoutines.Done()
	defer concurrency.HandlePanicError("manage_available_workers")
	for {
		err := w.workOnce()
		if err != nil {
			log.Warning(
				"Failed to manage workers. Reason: ", err.Error(),
				", taskType: ", w.taskType,
				", pollInterval: ", w.pollInterval.Milliseconds(), "ms",
				", domain: ", w.domain,
			)
			break
		}
	}
}

func (w *WorkerManager) executeAndUpdateTaskDaemon(task *http_model.Task) {
	defer w.runningGoRoutines.Done()
	defer concurrency.HandlePanicError("execute_and_update_task")
	defer w.availableWorkerMonitor.IncreaseAvailableWorkerAmount(1)
	taskResult, err := executeTask(task, w.executeFunction)
	if err != nil {
		log.Warning(
			"Failed to execute task. Reason: ", err.Error(),
			", taskType: ", w.taskType,
		)
		return
	}
	err = updateTaskWithRetry(w.taskType, taskResult, w.conductorClient)
	if err != nil {
		log.Warning(
			"Failed to update task with retry. Reason: ", err.Error(),
			", taskType: ", w.taskType,
		)
		return
	}
}

func (w *WorkerManager) workOnce() error {
	availableWorkerAmount, err := w.availableWorkerMonitor.GetAvailableWorkerAmount()
	if err != nil {
		return err
	}
	if availableWorkerAmount < 1 {
		log.Debug("No worker available for taskType: ", w.taskType)
		return nil
	}
	tasks, err := batchPoll(w.taskType, availableWorkerAmount, w.pollInterval, w.domain, w.conductorClient)
	if err != nil {
		return err
	}
	if len(tasks) < 1 {
		log.Debug("No tasks available for taskType: ", w.taskType)
		time.Sleep(w.pollInterval)
		return nil
	}
	for _, task := range tasks {
		err = w.availableWorkerMonitor.DecreaseAvailableWorkerAmount()
		if err != nil {
			return err
		}
		w.runningGoRoutines.Add(1)
		go w.executeAndUpdateTaskDaemon(&task)
	}
	return nil
}
