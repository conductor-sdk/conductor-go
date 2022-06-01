package examples

import (
	"fmt"
	"os"
	"time"

	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/conductor-sdk/conductor-go/pkg/model"
	"github.com/conductor-sdk/conductor-go/pkg/model/enum/task_result_status"
	"github.com/conductor-sdk/conductor-go/pkg/settings"
	"github.com/conductor-sdk/conductor-go/pkg/worker"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/def/workflow"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/executor"
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
		"https://play.orkes.io/api", // baseUrl
	)
)

var (
	apiClient = conductor_http_client.NewAPIClient(
		authenticationSettings,
		httpSettings,
	)

	taskRunner       = worker.NewTaskRunnerWithApiClient(apiClient)
	workflowExecutor = executor.NewWorkflowExecutor(apiClient)
)

var (
	httpTask = workflow.NewHttpTask(
		"GO_TASK_OF_HTTP_TYPE",
		&workflow.HttpInput{
			Uri: "https://catfact.ninja/fact",
		},
	)
	httpTaskWorkflow = workflow.NewConductorWorkflow(workflowExecutor).
				Name("GO_WORKFLOW_WITH_HTTP_TASK").
				Version(1).
				Add(httpTask)
)

var (
	simpleTask = workflow.NewSimpleTask(
		"GO_TASK_OF_SIMPLE_TYPE", // taskName
		"GO_TASK_OF_SIMPLE_TYPE", // taskReferenceName
	)
	simpleTaskWorkflow = workflow.NewConductorWorkflow(workflowExecutor).
				Name("GO_WORKFLOW_WITH_SIMPLE_TASK").
				Version(1).
				Add(simpleTask)
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func Worker(t *http_model.Task) (taskResult *http_model.TaskResult, err error) {
	taskResult = model.GetTaskResultFromTask(t)
	taskResult.OutputData = map[string]interface{}{
		"task": "task_1",
		"key3": 3,
		"key4": false,
	}
	taskResult.Status = task_result_status.COMPLETED
	return taskResult, nil
}

func runHttpWorkflowExample() error {
	_, err := httpTaskWorkflow.Register()
	if err != nil {
		return err
	}
	workflowId, workflowExecutionChannel, err := httpTaskWorkflow.Start(nil)
	if err != nil {
		return err
	}
	workflow, err := executor.WaitForWorkflowCompletionUntilTimeout(
		workflowId,
		workflowExecutionChannel,
		5*time.Second,
	)
	if err != nil {
		return err
	}
	if !executor.IsWorkflowCompleted(workflow) {
		return fmt.Errorf("failed to get completed workflow")
	}
	return nil
}

func runSimpleWorkflowExample() error {
	_, err := simpleTaskWorkflow.Register()
	if err != nil {
		return err
	}
	workflowId, workflowExecutionChannel, err := simpleTaskWorkflow.Start(nil)
	if err != nil {
		return err
	}
	err = taskRunner.StartWorker(
		simpleTask.ReferenceName(),
		Worker,
		2,
		500*time.Millisecond,
	)
	if err != nil {
		return err
	}
	workflow, err := executor.WaitForWorkflowCompletionUntilTimeout(
		workflowId,
		workflowExecutionChannel,
		5*time.Second,
	)
	if err != nil {
		return err
	}
	taskRunner.RemoveWorker(
		simpleTask.ReferenceName(),
		2,
	)
	if !executor.IsWorkflowCompleted(workflow) {
		return fmt.Errorf("failed to get completed workflow")
	}
	return nil
}

func main() {
	runHttpWorkflowExample()
	runSimpleWorkflowExample()
}
