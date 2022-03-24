package conductor

import (
	"os"
	"time"

	"github.com/netflix/conductor/client/go/metrics"
	"github.com/netflix/conductor/client/go/model"
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

type ConductorWorker struct {
	conductorHttpClient *ConductorHttpClient
	metricsCollector    *metrics.MetricsCollector
	pollingInterval     int
	threadCount         int
}

func NewConductorWorker(
	authenticationSettings *settings.AuthenticationSettings,
	httpSettings *settings.HttpSettings,
	metricsCollector *metrics.MetricsCollector,
	threadCount int,
	pollingInterval int,
) *ConductorWorker {
	conductorWorker := new(ConductorWorker)
	conductorHttpClient := NewConductorHttpClient(
		authenticationSettings,
		httpSettings,
	)
	conductorWorker.metricsCollector = metricsCollector
	conductorWorker.pollingInterval = pollingInterval
	conductorWorker.threadCount = threadCount
	conductorWorker.conductorHttpClient = conductorHttpClient
	return conductorWorker
}

func (c *ConductorWorker) Execute(t *model.Task, executeFunction func(t *model.Task) (*model.TaskResult, error)) {
	taskResult, err := executeFunction(t)
	if taskResult == nil {
		log.Error("TaskResult cannot be nil: ", t.TaskId)
		return
	}
	if err != nil {
		log.Error("Error Executing task:", err.Error())
		taskResult.Status = model.FAILED
		taskResult.ReasonForIncompletion = err.Error()
	}

	taskResultJsonString, err := taskResult.ToJSONString()
	if err != nil {
		log.Error("Error Forming TaskResult JSON body", err)
		return
	}
	_, _ = c.conductorHttpClient.UpdateTask(taskResultJsonString)
}

func (c *ConductorWorker) PollAndExecute(taskType string, domain string, executeFunction func(t *model.Task) (*model.TaskResult, error)) {
	for {
		c.metricsCollector.IncrementCounter()

		time.Sleep(time.Duration(c.pollingInterval) * time.Millisecond)

		// Poll for Task taskType
		polled, err := c.conductorHttpClient.PollForTask(taskType, hostname, domain)
		if err != nil {
			log.Error("Error Polling task:", err.Error())
			continue
		}
		if polled == "" {
			log.Debug("No task found for:", taskType)
			continue
		}

		// Parse Http response into Task
		parsedTask, err := model.ParseTask(polled)
		if err != nil {
			log.Error("Error Parsing task:", err.Error())
			continue
		}

		// Execute given function
		c.Execute(parsedTask, executeFunction)
	}
}

func (c *ConductorWorker) Start(taskType string, domain string, executeFunction func(t *model.Task) (*model.TaskResult, error), wait bool) {
	log.Println("Polling for task:", taskType, "with a:", c.pollingInterval, "(ms) polling interval with", c.threadCount, "goroutines for task execution, with workerId as", hostname)
	for i := 1; i <= c.threadCount; i++ {
		go c.PollAndExecute(taskType, domain, executeFunction)
	}

	// wait infinitely while the go routines are running
	if wait {
		select {}
	}
}
