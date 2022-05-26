package benchmark

// import (
// 	"os"
// 	"testing"
// 	"time"

// 	"github.com/conductor-sdk/conductor-go/examples"
// 	"github.com/conductor-sdk/conductor-go/pkg/worker"
// 	"github.com/conductor-sdk/conductor-go/pkg/workflow/def/workflow"
// 	"github.com/conductor-sdk/conductor-go/pkg/workflow/executor"
// 	"github.com/conductor-sdk/conductor-go/tests/e2e/e2e_properties"
// 	"github.com/conductor-sdk/conductor-go/tests/e2e/http_client_e2e/http_client_e2e_properties"
// 	"github.com/sirupsen/logrus"
// )

// var (
//  apiClient        = conductor_http_client.NewAPIClient(nil, nil)
// 	apiClient        = e2e_properties.API_CLIENT
// 	taskRunner       = worker.NewTaskRunnerWithApiClient(apiClient)
// 	workflowExecutor = executor.NewWorkflowExecutor(apiClient)
// )

// var (
// 	conductorWorkflow = workflow.NewConductorWorkflow(workflowExecutor).
// 		Name(http_client_e2e_properties.WORKFLOW_NAME).
// 		Version(1)
// )

// const (
// 	workflowCompletionTimeout = 15 * time.Second
// )

// func init() {
// 	logrus.SetFormatter(&logrus.JSONFormatter{})
// 	logrus.SetOutput(os.Stdout)
// 	logrus.SetLevel(logrus.DebugLevel)
// }

// func benchmark(t *testing.T) {
// 	for exponent := 0; exponent < 3; exponent += 1 {
// 		qty := 1 << exponent
// 		taskRunner.StartWorker(
// 			http_client_e2e_properties.TASK_NAME,
// 			examples.SimpleWorker,
// 			10,
// 			1000,
// 		)
// 		channels, err := conductorWorkflow.StartManyWithTimeout(qty, workflowCompletionTimeout)
// 		if err != nil {
// 			t.Error(err)
// 		}
// 		for _, channel := range channels {
// 			workflow := <-channel
// 			if workflow == nil {
// 				t.Error("Failed to pull completed workflow before timeout")
// 			}
// 		}
// 		taskRunner.RemoveWorker(
// 			http_client_e2e_properties.TASK_NAME,
// 			10,
// 		)
// 	}
// }
