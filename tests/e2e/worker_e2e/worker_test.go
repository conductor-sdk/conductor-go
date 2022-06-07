package worker_e2e

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/pkg/model"
	"github.com/conductor-sdk/conductor-go/pkg/model/enum/task_result_status"
	"github.com/conductor-sdk/conductor-go/pkg/model/enum/workflow_status"
	"github.com/conductor-sdk/conductor-go/pkg/worker"
	"github.com/conductor-sdk/conductor-go/tests/e2e/e2e_properties"
	"github.com/conductor-sdk/conductor-go/tests/e2e/http_client_e2e"
	"github.com/conductor-sdk/conductor-go/tests/e2e/http_client_e2e/http_client_e2e_properties"
	log "github.com/sirupsen/logrus"
)

const (
	workflowCompletionTimeout = 3 * time.Second
)

var taskRunner = worker.NewTaskRunnerWithApiClient(e2e_properties.API_CLIENT)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func TestWorkers(t *testing.T) {
	outputData := map[string]interface{}{
		"key": "value",
	}
	workerWithTaskResultOutput := func(t *model.Task) (interface{}, error) {
		taskResult := model.GetTaskResultFromTask(t)
		taskResult.OutputData = outputData
		taskResult.Status = task_result_status.COMPLETED
		return taskResult, nil
	}
	workerWithGenericOutput := func(t *model.Task) (interface{}, error) {
		return outputData, nil
	}
	workers := []model.ExecuteTaskFunction{
		workerWithTaskResultOutput,
		workerWithGenericOutput,
	}
	for _, worker := range workers {
		err := validateWorker(worker, outputData)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func validateWorker(worker model.ExecuteTaskFunction, expectedOutput map[string]interface{}) error {
	workflowIdList, err := http_client_e2e.StartWorkflows(
		http_client_e2e_properties.WORKFLOW_EXECUTION_AMOUNT,
		http_client_e2e_properties.WORKFLOW_NAME,
	)
	if err != nil {
		return err
	}
	err = taskRunner.StartWorker(
		http_client_e2e_properties.TASK_NAME,
		worker,
		http_client_e2e_properties.WORKER_THREAD_COUNT,
		http_client_e2e_properties.WORKER_POLLING_INTERVAL,
	)
	if err != nil {
		return err
	}
	runningWorkflows := make([]chan error, len(workflowIdList))
	for i, workflowId := range workflowIdList {
		runningWorkflows[i] = make(chan error)
		go validateWorkflowDaemon(
			runningWorkflows[i],
			workflowId,
			expectedOutput,
		)
	}
	for _, runningWorkflowChannel := range runningWorkflows {
		err := <-runningWorkflowChannel
		if err != nil {
			return err
		}
	}
	return nil
}

func validateWorkflowDaemon(outputChannel chan error, workflowId string, expectedOutput map[string]interface{}) {
	time.Sleep(workflowCompletionTimeout)
	workflow, _, err := http_client_e2e.GetWorkflowExecutionStatus(
		workflowId,
	)
	if err != nil {
		outputChannel <- err
		return
	}
	if workflow.Status != string(workflow_status.COMPLETED) {
		outputChannel <- fmt.Errorf(
			"workflow status different than expected, workflowId: %s, workflowStatus: %s",
			workflow.WorkflowId, workflow.Status,
		)
		return
	}
	if !reflect.DeepEqual(workflow.Output, expectedOutput) {
		outputChannel <- fmt.Errorf(
			"workflow output is different than expected, workflowId: %s, output: %+v",
			workflow.WorkflowId, workflow.Output,
		)
		return
	}
	outputChannel <- nil
}
