package orkestrator

import (
	"sync"
	"time"
	"unsafe"

	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/http"
	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/model"
	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/model/enum/task_result_status"
	"github.com/conductor-sdk/conductor-go/pkg/metrics"
	"github.com/conductor-sdk/conductor-go/pkg/settings"
	log "github.com/sirupsen/logrus"
)

type WorkerOrkestrator struct {
	conductorHttpClient *http.ConductorHttpClient
	metricsCollector    *metrics.MetricsCollector
	waitGroup           sync.WaitGroup
}

func NewWorkerOrkestrator(
	authenticationSettings *settings.AuthenticationSettings,
	httpSettings *settings.HttpSettings,
) *WorkerOrkestrator {
	workerOrkestrator := new(WorkerOrkestrator)
	workerOrkestrator.metricsCollector = metrics.NewMetricsCollector()
	workerOrkestrator.conductorHttpClient = http.NewConductorHttpClient(
		authenticationSettings,
		httpSettings,
	)
	return workerOrkestrator
}

func (c *WorkerOrkestrator) StartWorker(taskType string, executeFunction model.TaskExecuteFunction, parallelGoRoutinesAmount int, pollingInterval int) {
	for goRoutines := 1; goRoutines <= parallelGoRoutinesAmount; goRoutines++ {
		c.waitGroup.Add(1)
		go c.run(taskType, executeFunction, pollingInterval)
	}
	log.Debug(
		"Started worker for task:", taskType,
		", go routines amount:", parallelGoRoutinesAmount,
		", polling interval:", pollingInterval, "(ms)",
	)
}

func (c *WorkerOrkestrator) WaitWorkers() {
	c.waitGroup.Wait()
}

func (c *WorkerOrkestrator) run(taskType string, executeFunction model.TaskExecuteFunction, pollingInterval int) {
	for {
		c.runOnce(taskType, executeFunction)
		sleep(pollingInterval)
	}
	c.waitGroup.Done()
}

func (c *WorkerOrkestrator) runOnce(taskType string, executeFunction model.TaskExecuteFunction) {
	task := c.pollTask(taskType)
	if task == nil {
		return
	}
	taskResult := c.executeTask(task, executeFunction)
	c.updateTask(taskType, taskResult)
}

func (c *WorkerOrkestrator) pollTask(taskType string) *model.Task {
	c.metricsCollector.IncrementTaskPoll(taskType)

	startTime := time.Now()
	polled, err := c.conductorHttpClient.PollForTask(taskType)
	spentTime := time.Since(startTime)
	c.metricsCollector.RecordTaskPollTime(
		taskType,
		spentTime.Seconds(),
	)

	if err != nil {
		log.Error("Error Polling task:", err.Error())
		c.metricsCollector.IncrementTaskPollError(
			taskType, err,
		)
		return nil
	}
	if polled == "" {
		log.Debug("No task found for:", taskType)
		return nil
	}

	parsedTask, err := model.ParseTask(polled)
	if err != nil {
		log.Error("Error Parsing task:", err.Error())
		c.metricsCollector.IncrementTaskPollError(
			taskType, err,
		)
		return nil
	}

	return parsedTask
}

func (c *WorkerOrkestrator) executeTask(t *model.Task, executeFunction model.TaskExecuteFunction) *model.TaskResult {
	startTime := time.Now()
	taskResult, err := executeFunction(t)
	spentTime := time.Since(startTime)
	c.metricsCollector.RecordTaskExecuteTime(
		t.TaskDefName, spentTime.Seconds(),
	)
	if taskResult == nil {
		log.Error("TaskResult cannot be nil: ", t.TaskId)
		return nil
	}
	if err != nil {
		log.Error("Error Executing task:", err.Error())
		taskResult.Status = task_result_status.FAILED
		taskResult.ReasonForIncompletion = err.Error()
		c.metricsCollector.IncrementTaskExecuteError(
			t.TaskDefName, err,
		)
	}
	size := unsafe.Sizeof(taskResult)
	c.metricsCollector.RecordTaskResultPayloadSize(
		t.TaskDefName, float64(size),
	)
	return taskResult
}

func (c *WorkerOrkestrator) updateTask(taskType string, taskResult *model.TaskResult) {
	taskResultJsonString, err := taskResult.ToJSONString()
	if err != nil {
		log.Error("Error Forming TaskResult JSON body", err)
		c.metricsCollector.IncrementTaskUpdateError(
			taskType, err,
		)
		return
	}
	_, _ = c.conductorHttpClient.UpdateTask(taskResultJsonString)
}

func sleep(pollingInterval int) {
	time.Sleep(
		time.Duration(pollingInterval) * time.Millisecond,
	)
}
