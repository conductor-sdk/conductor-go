package workflow_e2e

import (
	"os"
	"testing"

	"github.com/conductor-sdk/conductor-go/examples"
	"github.com/conductor-sdk/conductor-go/pkg/model/enum/task_result_status"
	"github.com/conductor-sdk/conductor-go/pkg/worker"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/def/workflow"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/executor"
	"github.com/conductor-sdk/conductor-go/tests/e2e/e2e_properties"
	"github.com/conductor-sdk/conductor-go/tests/e2e/http_client_e2e/http_client_e2e_properties"
	log "github.com/sirupsen/logrus"
)

var taskRunner = worker.NewTaskRunnerWithApiClient(e2e_properties.API_CLIENT)
var workflowExecutor = executor.NewWorkflowExecutor(e2e_properties.API_CLIENT)

var (
	workflows = []*workflow.ConductorWorkflow{
		HTTP_WORKFLOW,
		SIMPLE_WORKFLOW,
	}
)

var (
	HTTP_TASK_WORKFLOW_NAME = "GO_WORKFLOW_WITH_HTTP_TASK"
	HTTP_TASK_NAME          = "GO_TASK_OF_HTTP_TYPE"

	HTTP_TASK = workflow.NewHttpTask(
		HTTP_TASK_NAME,
		&workflow.HttpInput{
			Uri: "https://catfact.ninja/fact",
		},
	)

	HTTP_WORKFLOW = workflow.NewConductorWorkflow(workflowExecutor).
			Name(HTTP_TASK_WORKFLOW_NAME).
			Version(1).
			Add(HTTP_TASK)
)

var (
	SIMPLE_TASK_WORKFLOW_NAME = "GO_WORKFLOW_WITH_SIMPLE_TASK"
	SIMPLE_TASK_NAME          = "GO_TASK_OF_SIMPLE_TYPE"

	SIMPLE_TASK = workflow.NewSimpleTask(
		SIMPLE_TASK_NAME,
		SIMPLE_TASK_NAME,
	)

	SIMPLE_WORKFLOW = workflow.NewConductorWorkflow(workflowExecutor).
			Name(SIMPLE_TASK_WORKFLOW_NAME).
			Version(1).
			Add(SIMPLE_TASK)
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func TestValidateWorkflowDefinitions(t *testing.T) {
	for _, conductorWorkflow := range workflows {
		response, err := conductorWorkflow.Register()
		if err != nil {
			t.Error("Response: ", response, ", error: ", err)
		}
	}
}

func TestWorkflowDefExecutionWithSingleStart(t *testing.T) {
	workflowExecutionChannelList := make(
		[][]executor.WorkflowExecutionChannel,
		len(workflows),
	)
	for i, conductorWorkflow := range workflows {
		qty := 5
		workflowExecutionChannelList[i] = make([]executor.WorkflowExecutionChannel, qty)
		for j := 0; j < qty; j += 1 {
			workflowExecutionChannel, err := conductorWorkflow.Start(nil)
			if err != nil {
				t.Error(err)
			}
			workflowExecutionChannelList[i][j] = workflowExecutionChannel
		}
	}

	taskRunner.StartWorker(
		SIMPLE_TASK_NAME,
		examples.SimpleWorker,
		http_client_e2e_properties.WORKER_THREAD_COUNT,
		http_client_e2e_properties.WORKER_POLLING_INTERVAL,
	)

	for _, channels := range workflowExecutionChannelList {
		for _, channel := range channels {
			workflow := <-channel
			if workflow == nil || workflow.Status != string(task_result_status.COMPLETED) {
				t.Error()
			}
		}
	}

	taskRunner.RemoveWorker(
		SIMPLE_WORKFLOW.GetName(),
		http_client_e2e_properties.WORKER_THREAD_COUNT,
	)
}
