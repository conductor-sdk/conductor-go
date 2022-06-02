package workflow_e2e

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/examples"
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
		http_client_e2e_properties.TASK_NAME,
		http_client_e2e_properties.TASK_NAME,
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

func TestHttpTask(t *testing.T) {
	_, err := httpTaskWorkflow.Register()
	if err != nil {
		t.Fatal(err)
	}
	_, workflowExecutionChannel, err := httpTaskWorkflow.Start(nil)
	if err != nil {
		t.Fatal(err)
	}
	workflow, err := executor.WaitForWorkflowCompletionUntilTimeout(
		workflowExecutionChannel,
		5*time.Second,
	)
	if err != nil {
		t.Fatal(err)
	}
	if !isWorkflowCompleted(workflow) {
		t.Fatal("Workflow finished with incomplete status, workflow: ", workflow.Status)
	}
}

func TestSimpleTask(t *testing.T) {
	_, err := simpleTaskWorkflow.Register()
	if err != nil {
		t.Fatal(err)
	}
	_, workflowExecutionChannel, err := simpleTaskWorkflow.Start(nil)
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
	workflow, err := executor.WaitForWorkflowCompletionUntilTimeout(
		workflowExecutionChannel,
		5*time.Second,
	)
	if err != nil {
		t.Fatal(err)
	}
	taskRunner.RemoveWorker(
		simpleTask.ReferenceName(),
		http_client_e2e_properties.WORKER_THREAD_COUNT,
	)
	if !isWorkflowCompleted(workflow) {
		t.Fatal("Workflow finished with incomplete status, workflow: ", workflow.Status)
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
	err := validateWorkflow(inlineTaskWorkflow, 5*time.Second)
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
