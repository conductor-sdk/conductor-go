package workflow_e2e

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/examples"
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/conductor-sdk/conductor-go/pkg/worker"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/def/workflow"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/executor"
	"github.com/conductor-sdk/conductor-go/tests/e2e/e2e_properties"
	"github.com/conductor-sdk/conductor-go/tests/e2e/http_client_e2e/http_client_e2e_properties"
	"github.com/sirupsen/logrus"
)

var taskRunner = worker.NewTaskRunnerWithApiClient(e2e_properties.API_CLIENT)
var workflowExecutor = executor.NewWorkflowExecutor(e2e_properties.API_CLIENT)

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
		http_client_e2e_properties.TASK_NAME,
		http_client_e2e_properties.TASK_NAME,
	)

	simpleTaskWorkflow = workflow.NewConductorWorkflow(workflowExecutor).
				Name(http_client_e2e_properties.WORKFLOW_NAME).
				Version(1).
				Add(simpleTask)
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
}

func TestHttpTask(t *testing.T) {
	_, err := httpTaskWorkflow.Register()
	if err != nil {
		t.Fatal(err)
	}
	workflowExecutionChannel, err := httpTaskWorkflow.Start(nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = waitUntilTimeout(workflowExecutionChannel, 5*time.Second)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSimpleTask(t *testing.T) {
	_, err := simpleTaskWorkflow.Register()
	if err != nil {
		t.Fatal(err)
	}
	workflowExecutionChannel, err := simpleTaskWorkflow.Start(nil)
	if err != nil {
		t.Fatal(err)
	}
	err = taskRunner.StartWorker(
		simpleTask.ReferenceName(),
		examples.SimpleWorker,
		http_client_e2e_properties.WORKER_THREAD_COUNT,
		http_client_e2e_properties.WORKER_POLLING_INTERVAL,
	)
	if err != nil {
		t.Fatal(err)
	}
	_, err = waitUntilTimeout(workflowExecutionChannel, 5*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	taskRunner.RemoveWorker(
		simpleTask.ReferenceName(),
		http_client_e2e_properties.WORKER_THREAD_COUNT,
	)
}

func waitUntilTimeout(channel executor.WorkflowExecutionChannel, timeout time.Duration) (*http_model.Workflow, error) {
	select {
	case value := <-channel:
		return value, nil
	case <-time.After(timeout):
		return nil, fmt.Errorf("timeout waiting for channel")
	}
}
