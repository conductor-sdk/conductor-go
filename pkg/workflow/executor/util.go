package executor

import (
	"encoding/json"

	"github.com/conductor-sdk/conductor-go/pkg/http_model"
)

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
