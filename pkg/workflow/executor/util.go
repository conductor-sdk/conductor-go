package executor

import (
	"encoding/json"

	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/sirupsen/logrus"
)

func getInputAsMap(input interface{}) (map[string]interface{}, error) {
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

func isWorkflowInTerminalState(workflow *http_model.Workflow) bool {
	return workflow.Status != "PAUSED" && workflow.Status != "RUNNING"
}
