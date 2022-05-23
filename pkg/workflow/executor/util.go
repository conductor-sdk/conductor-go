package executor

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/sirupsen/logrus"
)

func WaitForCompletionOfWorkflows(workflowExecutionChannelList []WorkflowExecutionChannel, workflowCompletionTimeout time.Duration) {
	var waitGroup sync.WaitGroup
	for _, workflowExecutionChannel := range workflowExecutionChannelList {
		go getWorkflowFromExecutionChannel(
			&waitGroup,
			workflowExecutionChannel,
			workflowCompletionTimeout,
		)
	}
	waitGroup.Wait()
}

func getWorkflowFromExecutionChannel(waitGroup *sync.WaitGroup, workflowExecutionChannel WorkflowExecutionChannel, workflowCompletionTimeout time.Duration) *http_model.Workflow {
	defer waitGroup.Done()
	select {
	case workflow := <-workflowExecutionChannel:
		return workflow
	case <-time.After(workflowCompletionTimeout):
		logrus.Warning("Timeout waiting for workflow completion")
		close(workflowExecutionChannel)
	}
	return nil
}

func getInputAsMap(input interface{}) map[string]interface{} {
	if input == nil {
		return nil
	}
	data, _ := json.Marshal(input)
	var parsedInput map[string]interface{}
	json.Unmarshal(data, &parsedInput)
	return parsedInput
}

func isWorkflowInTerminalState(workflow *http_model.Workflow) bool {
	return workflow.Status != "PAUSED" && workflow.Status != "RUNNING"
}
