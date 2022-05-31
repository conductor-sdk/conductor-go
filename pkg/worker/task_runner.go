package worker

import (
	"sync"
	"time"

	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/model"
	"github.com/conductor-sdk/conductor-go/pkg/settings"
	log "github.com/sirupsen/logrus"
)

type TaskRunner struct {
	conductorTaskResourceClient conductor_http_client.TaskResourceApiService
	workerManagerByTaskType     map[string]AvailableWorkerChannel
	runningWorkers              sync.WaitGroup
}

func NewTaskRunner(authenticationSettings *settings.AuthenticationSettings, httpSettings *settings.HttpSettings) *TaskRunner {
	apiClient := conductor_http_client.NewAPIClient(
		authenticationSettings,
		httpSettings,
	)
	return NewTaskRunnerWithApiClient(apiClient)
}

func NewTaskRunnerWithApiClient(
	apiClient *conductor_http_client.APIClient,
) *TaskRunner {
	return &TaskRunner{
		conductorTaskResourceClient: conductor_http_client.TaskResourceApiService{
			APIClient: apiClient,
		},
		workerManagerByTaskType: make(map[string]AvailableWorkerChannel),
	}
}

// StartWorkerWithDomain
//  - taskType Task Type to poll and execute the work
//  - executeFunction Task execution function
//  - batchSize No. of tasks to poll for.  Each polled task is executed in a goroutine.  Batching improves the throughput
//  - pollInterval Time to wait for between polls if there are no tasks available.  Reduces excessive polling on the server when there is no work
//  - domain Task domain. Optional for polling
func (c *TaskRunner) StartWorkerWithDomain(taskType string, executeFunction model.TaskExecuteFunction, threadCount int, pollInterval time.Duration, domain string) error {
	return c.startWorker(taskType, executeFunction, threadCount, pollInterval, optional.NewString(domain))
}

// StartWorker
//  - taskType Task Type to poll and execute the work
//  - executeFunction Task execution function
//  - batchSize No. of tasks to poll for.  Each polled task is executed in a goroutine.  Batching improves the throughput
//  - pollInterval Time to wait for between polls if there are no tasks available.  Reduces excessive polling on the server when there is no work
func (c *TaskRunner) StartWorker(taskType string, executeFunction model.TaskExecuteFunction, batchSize int, pollInterval time.Duration) error {
	return c.startWorker(taskType, executeFunction, batchSize, pollInterval, optional.EmptyString())
}

func (c *TaskRunner) WaitWorkers() {
	c.runningWorkers.Wait()
}

func (c *TaskRunner) startWorker(taskType string, executeFunction model.TaskExecuteFunction, threadCount int, pollInterval time.Duration, taskDomain optional.String) error {
	availableWorkerChannel, ok := c.workerManagerByTaskType[taskType]
	if !ok {
		c.runningWorkers.Add(1)
		availableWorkerChannel, err := startWorkerManager(
			taskType,
			executeFunction,
			pollInterval,
			taskDomain,
			c.conductorTaskResourceClient,
		)
		if err != nil {
			return err
		}
		c.workerManagerByTaskType[taskType] = availableWorkerChannel
	}
	availableWorkerChannel <- threadCount
	log.Debug(
		"Started worker for task: ", taskType,
		", polling in batches of: ", threadCount,
		", with poll interval of: ", pollInterval.Milliseconds(), "ms",
	)
	return nil
}
