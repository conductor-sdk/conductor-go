package worker_e2e

import (
	"os"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/pkg/model"
	"github.com/conductor-sdk/conductor-go/pkg/model/enum/task_result_status"
	"github.com/conductor-sdk/conductor-go/tests/e2e/e2e_properties"
	log "github.com/sirupsen/logrus"
)

const (
	workflowCompletionTimeout = 5 * time.Second
	workflowExecutionQty      = 15

	workerQty          = 7
	workerPollInterval = 250 * time.Millisecond
)

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
	workflowIdList, err := e2e_properties.StartWorkflows(
		workflowExecutionQty,
		e2e_properties.WORKFLOW_NAME,
	)
	if err != nil {
		return err
	}
	err = e2e_properties.TaskRunner.StartWorker(
		e2e_properties.TASK_NAME,
		worker,
		workerQty,
		workerPollInterval,
	)
	if err != nil {
		return err
	}
	runningWorkflows := make([]chan error, len(workflowIdList))
	for i, workflowId := range workflowIdList {
		runningWorkflows[i] = make(chan error)
		go e2e_properties.ValidateWorkflowDaemon(
			workflowCompletionTimeout,
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
	return e2e_properties.TaskRunner.RemoveWorker(
		e2e_properties.TASK_NAME,
		workerQty,
	)
}
