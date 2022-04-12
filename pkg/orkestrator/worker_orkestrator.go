package orkestrator

import (
	"context"
	"sync"
	"time"
	"unsafe"

	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/conductor-sdk/conductor-go/pkg/metrics"
	"github.com/conductor-sdk/conductor-go/pkg/model"
	"github.com/conductor-sdk/conductor-go/pkg/model/enum/task_result_status"
	"github.com/conductor-sdk/conductor-go/pkg/settings"
	log "github.com/sirupsen/logrus"
)

type WorkerOrkestrator struct {
	conductorTaskResourceClient *conductor_http_client.TaskResourceApiService
	metricsCollector            *metrics.MetricsCollector
	waitGroup                   sync.WaitGroup
}

func NewWorkerOrkestrator(
	authenticationSettings *settings.AuthenticationSettings,
	httpSettings *settings.HttpSettings,
) *WorkerOrkestrator {
	return &WorkerOrkestrator{
		conductorTaskResourceClient: conductor_http_client.NewTaskResourceApiService(
			authenticationSettings,
			httpSettings,
		),
		metricsCollector: metrics.NewMetricsCollector(),
	}
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

func (c *WorkerOrkestrator) pollTask(taskType string) *http_model.Task {
	c.metricsCollector.IncrementTaskPoll(taskType)

	startTime := time.Now()

	task, _, err := c.conductorTaskResourceClient.Poll(
		context.Background(),
		taskType,
		nil,
	)

	spentTime := time.Since(startTime)
	c.metricsCollector.RecordTaskPollTime(
		taskType,
		spentTime.Seconds(),
	)

	if err != nil {
		log.Error("Error Parsing task:", err.Error())
		c.metricsCollector.IncrementTaskPollError(
			taskType, err,
		)
		return nil
	}

	return &task
}

func (c *WorkerOrkestrator) executeTask(t *http_model.Task, executeFunction model.TaskExecuteFunction) *http_model.TaskResult {
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

func (c *WorkerOrkestrator) updateTask(taskType string, taskResult *http_model.TaskResult) {
	_, _, err := c.conductorTaskResourceClient.UpdateTask(
		context.Background(),
		*taskResult,
	)
	if err != nil {
		log.Error("Error Updating task:", err.Error())
		c.metricsCollector.IncrementTaskUpdateError(taskType, err)
	}
}

func sleep(pollingInterval int) {
	time.Sleep(
		time.Duration(pollingInterval) * time.Millisecond,
	)
}
