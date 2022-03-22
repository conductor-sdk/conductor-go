package sample

import (
	"github.com/netflix/conductor/client/go/model"
	log "github.com/sirupsen/logrus"
)

func Task_1_Execution_Function(t *model.Task) (taskResult *model.TaskResult, err error) {
	log.Debug("Executing Task_1_Execution_Function for", t.TaskType)

	//Do some logic
	taskResult = model.NewTaskResult(t)

	output := map[string]interface{}{
		"task": "task_1",
		"key2": "value2",
		"key3": 3,
		"key4": false,
	}
	taskResult.OutputData = output
	taskResult.Logs = append(taskResult.Logs, model.LogMessage{Log: "Hello World"})
	taskResult.Status = "COMPLETED"
	err = nil

	return taskResult, err
}
