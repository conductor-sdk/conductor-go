package worker_e2e

import (
	model2 "github.com/conductor-sdk/conductor-go/model"
	"os"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/tests/e2e/e2e_properties"
	log "github.com/sirupsen/logrus"
)

const (
	taskName = "TEST_GO_TASK_SIMPLE"

	workflowName              = "TEST_GO_WORKFLOW_SIMPLE"
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
	workerWithTaskResultOutput := func(t *model2.Task) (interface{}, error) {
		taskResult := model2.GetTaskResultFromTask(t)
		taskResult.OutputData = outputData
		taskResult.Status = model2.COMPLETED
		return taskResult, nil
	}
	workerWithGenericOutput := func(t *model2.Task) (interface{}, error) {
		return outputData, nil
	}
	workers := []model2.ExecuteTaskFunction{
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

func validateWorker(worker model2.ExecuteTaskFunction, expectedOutput map[string]interface{}) error {
	workflowIdList, err := e2e_properties.StartWorkflows(
		workflowExecutionQty,
		workflowName,
	)
	if err != nil {
		return err
	}
	err = e2e_properties.TaskRunner.StartWorker(
		taskName,
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
		taskName,
		workerQty,
	)
}
