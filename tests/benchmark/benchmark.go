package benchmark

import (
	"os"
	"time"

	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/worker"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/def/workflow"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/executor"
	"github.com/conductor-sdk/conductor-go/tests/e2e/http_client_e2e/http_client_e2e_properties"
	log "github.com/sirupsen/logrus"
)

var (
	apiClient        = conductor_http_client.NewAPIClient(nil, nil)
	taskRunner       = worker.NewTaskRunnerWithApiClient(apiClient)
	workflowExecutor = executor.NewWorkflowExecutor(apiClient)
)

var (
	conductorWorkflow = workflow.NewConductorWorkflow(workflowExecutor).
		Name(http_client_e2e_properties.WORKFLOW_NAME).
		Version(1)
)

const (
	workflowCompletionTimeout = 15 * time.Second
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	for exponent := 0; exponent < 15; exponent += 1 {
		qty := 1 << exponent
		workflowExecutionChannelList := conductorWorkflow.StartMany(qty)
		executor.WaitForCompletionOfWorkflows(
			workflowExecutionChannelList,
			workflowCompletionTimeout,
		)
	}
}
