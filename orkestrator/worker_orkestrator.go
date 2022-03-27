package orkestrator

import (
	"os"
	"time"
	"unsafe"

	"github.com/netflix/conductor/client/go/conductor_client/http"
	"github.com/netflix/conductor/client/go/conductor_client/model"
	"github.com/netflix/conductor/client/go/conductor_client/model/enum/task_result_status"
	"github.com/netflix/conductor/client/go/metrics"
	"github.com/netflix/conductor/client/go/settings"
	log "github.com/sirupsen/logrus"
)

var (
	hostname, hostnameError = os.Hostname()
)

func init() {
	if hostnameError != nil {
		log.Fatal("Could not get hostname")
	}
}

type WorkerOrkestrator struct {
	conductorHttpClient *http.ConductorHttpClient
	metricsCollector    *metrics.MetricsCollector
	pollingInterval     int
	threadCount         int
}

func NewWorkerOrkestrator(
	authenticationSettings *settings.AuthenticationSettings,
	httpSettings *settings.HttpSettings,
	metricsCollector *metrics.MetricsCollector,
	threadCount int,
	pollingInterval int,
) *WorkerOrkestrator {
	workerOrkestrator := new(WorkerOrkestrator)
	conductorHttpClient := http.NewConductorHttpClient(
		authenticationSettings,
		httpSettings,
	)
	workerOrkestrator.metricsCollector = metricsCollector
	workerOrkestrator.pollingInterval = pollingInterval
	workerOrkestrator.threadCount = threadCount
	workerOrkestrator.conductorHttpClient = conductorHttpClient
	return workerOrkestrator
}

func (c *WorkerOrkestrator) StartWorker(taskType string, domain string, executeFunction model.TaskExecuteFunction, wait bool) {
	log.Println(
		"Polling for task:", taskType,
		"with a:", c.pollingInterval,
		"(ms) polling interval with", c.threadCount,
		"goroutines for task execution, with workerId as", hostname,
	)
	for goRoutines := 1; goRoutines <= c.threadCount; goRoutines++ {
		go c.run(taskType, domain, executeFunction)
	}
	// wait infinitely while the go routines are running
	if wait {
		select {}
	}
}

func (c *WorkerOrkestrator) run(taskType string, domain string, executeFunction model.TaskExecuteFunction) {
	for {
		c.runOnce(taskType, domain, executeFunction)
		c.sleep()
	}
}

func (c *WorkerOrkestrator) runOnce(taskType string, domain string, executeFunction model.TaskExecuteFunction) {
	task := c.pollTask(taskType, domain)
	if task == nil {
		return
	}
	taskResult := c.executeTask(task, executeFunction)
	c.updateTask(taskType, taskResult)
}

func (c *WorkerOrkestrator) sleep() {
	time.Sleep(
		time.Duration(c.pollingInterval) * time.Millisecond,
	)
}

func (c *WorkerOrkestrator) pollTask(taskType string, domain string) *model.Task {
	c.metricsCollector.IncrementTaskPoll(taskType)

	startTime := time.Now()
	polled, err := c.conductorHttpClient.PollForTask(
		taskType,
		hostname,
		domain,
	)
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
