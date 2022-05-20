package examples

import (
	"context"
	"os"

	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/conductor-sdk/conductor-go/pkg/model"
	"github.com/conductor-sdk/conductor-go/pkg/model/enum/task_result_status"
	"github.com/conductor-sdk/conductor-go/pkg/settings"
	"github.com/conductor-sdk/conductor-go/pkg/worker"
	log "github.com/sirupsen/logrus"
)

var (
	// To obtain a key / secret for your server, see
	// https://orkes.io/content/docs/getting-started/concepts/access-control#access-keys
	// If you are testing against a server that does not require authentication, pass nil
	authenticationSettings = settings.NewAuthenticationSettings(
		"", // keyId
		"", // keySecret
	)

	httpSettings = settings.NewHttpSettings(
		"https://play.orkes.io", // baseUrl
	)
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	startWorkflows()
	startWorkers()
}

func Worker(t *http_model.Task) (taskResult *http_model.TaskResult, err error) {
	taskResult = model.GetTaskResultFromTask(t)
	taskResult.OutputData = map[string]interface{}{
		"task": "task_1",
		"key3": 3,
		"key4": false,
	}
	taskResult.Status = task_result_status.COMPLETED
	err = nil
	return taskResult, err
}

func startWorkflows() {
	apiClient := conductor_http_client.NewAPIClient(
		authenticationSettings,
		httpSettings,
	)
	workflowClient := conductor_http_client.WorkflowResourceApiService{
		APIClient: apiClient,
	}
	workflowClient.StartWorkflow(
		context.Background(),
		nil,
		"workflow_name",
		nil,
	)
}

func startWorkers() {
	taskRunner := worker.NewTaskRunner(
		authenticationSettings,
		httpSettings,
	)
	taskRunner.StartWorker(
		"go_task_example",
		Worker,
		2,
		10,
	)
	taskRunner.WaitWorkers()
}
