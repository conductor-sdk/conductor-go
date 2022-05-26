package executor

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/conductor-sdk/conductor-go/pkg/model/enum/workflow_status"
	"github.com/sirupsen/logrus"
)

func GetInputAsMap(input interface{}) (map[string]interface{}, error) {
	if input == nil {
		return nil, nil
	}
	data, err := json.Marshal(input)
	if err != nil {
		logrus.Debug(
			"Failed to parse input",
			", reason: ", err.Error(),
		)
		return nil, err
	}
	var parsedInput map[string]interface{}
	json.Unmarshal(data, &parsedInput)
	return parsedInput, nil
}

func IsWorkflowInTerminalState(workflow *http_model.Workflow) bool {
	for _, terminalState := range workflow_status.WorkflowTerminalStates {
		if workflow.Status == string(terminalState) {
			return true
		}
	}
	return false
}

func IsWorkflowCompleted(workflow *http_model.Workflow) bool {
	return workflow.Status == string(workflow_status.COMPLETED)
}

func WaitForWorkflowCompletionUntilTimeout(workflowId string, executionChannel WorkflowExecutionChannel, timeout time.Duration) (*http_model.Workflow, error) {
	select {
	case workflow, ok := <-executionChannel:
		if !ok {
			return nil, fmt.Errorf(
				"failed to wait for workflow completion, reason: channel closed, workflowId: %s",
				workflowId,
			)
		}
		return workflow, nil
	case <-time.After(timeout):
		return nil, fmt.Errorf(
			"timeout waiting for workflow completion, workflowId: %s",
			workflowId,
		)
	}
}
