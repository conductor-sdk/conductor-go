package workflow_e2e

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/examples"
	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/conductor-sdk/conductor-go/pkg/model/enum/workflow_status"
	"github.com/conductor-sdk/conductor-go/pkg/worker"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/def/workflow"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/executor"
	"github.com/conductor-sdk/conductor-go/tests/e2e/e2e_properties"
	"github.com/conductor-sdk/conductor-go/tests/e2e/http_client_e2e/http_client_e2e_properties"
	log "github.com/sirupsen/logrus"
)

var (
	taskRunner       = worker.NewTaskRunnerWithApiClient(e2e_properties.API_CLIENT)
	workflowExecutor = executor.NewWorkflowExecutor(e2e_properties.API_CLIENT)
	metadataClient   = conductor_http_client.MetadataResourceApiService{
		APIClient: e2e_properties.API_CLIENT,
	}
)

const workflowValidationTimeout = 5 * time.Second

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func TestHttpTask(t *testing.T) {
	httpTask := workflow.NewHttpTask(
		"TEST_GO_TASK_HTTP",
		&workflow.HttpInput{
			Uri: "https://catfact.ninja/fact",
		},
	)
	httpTaskWorkflow := workflow.NewConductorWorkflow(workflowExecutor).
		Name("TEST_GO_WORKFLOW_HTTP").
		Version(1).
		Add(httpTask)
	err := validateWorkflow(httpTaskWorkflow, workflowValidationTimeout)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSimpleTask(t *testing.T) {
	simpleTask := workflow.NewSimpleTask(
		"TEST_GO_TASK_SIMPLE",
	)
	response, err := registerTask(simpleTask)
	if err != nil {
		t.Fatal("Failed to register task, response: ", response)
	}
	simpleTaskWorkflow := workflow.NewConductorWorkflow(workflowExecutor).
		Name("TEST_GO_WORKFLOW_SIMPLE").
		Version(1).
		Add(simpleTask)
	err = taskRunner.StartWorker(
		simpleTask.ReferenceName(),
		examples.SimpleWorker,
		5,
		500*time.Millisecond,
	)
	if err != nil {
		t.Fatal(err)
	}
	err = validateWorkflow(simpleTaskWorkflow, workflowValidationTimeout)
	if err != nil {
		t.Fatal(err)
	}
	err = taskRunner.RemoveWorker(
		simpleTask.ReferenceName(),
		http_client_e2e_properties.WORKER_THREAD_COUNT,
	)
	if err != nil {
		t.Fatal(err)
	}
}

func TestInlineTask(t *testing.T) {
	jsCode := "function e() { if ($.value == 1){return {\"result\": true}} else { return {\"result\": false}}} e();"
	inlineTask := workflow.NewInlineTask(
		"TEST_GO_TASK_INLINE",
		map[string]interface{}{
			"value":         "${workflow.input.value}",
			"evaluatorType": "javascript",
			"expression":    jsCode,
		},
	)
	inlineTaskWorkflow := workflow.NewConductorWorkflow(workflowExecutor).
		Name("TEST_GO_WORKFLOW_INLINE_TASK").
		Version(1).
		Add(inlineTask)
	err := validateWorkflow(inlineTaskWorkflow, workflowValidationTimeout)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSqsEventTask(t *testing.T) {
	sqsEventTask := workflow.NewSqsEventTask(
		"TEST_GO_TASK_EVENT_SQS",
		"QUEUE",
	)
	testEventTask(
		t,
		"TEST_GO_WORKFLOW_EVENT_SQS",
		sqsEventTask,
	)
}

func TestConductorEventTask(t *testing.T) {
	sqsEventTask := workflow.NewSqsEventTask(
		"TEST_GO_TASK_EVENT_CONDUCTOR",
		"QUEUE",
	)
	testEventTask(
		t,
		"TEST_GO_WORKFLOW_EVENT_CONDUCTOR",
		sqsEventTask,
	)
}

func testEventTask(t *testing.T, workflowName string, event *workflow.EventTask) {
	eventTaskWorkflow := workflow.NewConductorWorkflow(workflowExecutor).
		Name(workflowName).
		Version(1).
		Add(event)
	_, err := eventTaskWorkflow.Register()
	if err != nil {
		t.Fatal(err)
	}
}

func isWorkflowCompleted(workflow *http_model.Workflow) bool {
	return workflow.Status == string(workflow_status.COMPLETED)
}

func validateWorkflow(conductorWorkflow *workflow.ConductorWorkflow, timeout time.Duration) error {
	_, err := conductorWorkflow.Register()
	if err != nil {
		return err
	}
	_, workflowExecutionChannel, err := conductorWorkflow.Start(nil)
	if err != nil {
		return err
	}
	workflow, err := executor.WaitForWorkflowCompletionUntilTimeout(
		workflowExecutionChannel,
		timeout,
	)
	if err != nil {
		return err
	}
	if !isWorkflowCompleted(workflow) {
		return fmt.Errorf("workflow finished with status: %s", workflow.Status)
	}
	return nil
}

func registerTask(task *workflow.SimpleTask) (*http.Response, error) {
	return metadataClient.RegisterTaskDef(
		context.Background(),
		[]http_model.TaskDef{
			*task.ToTaskDef(),
		},
	)
}
