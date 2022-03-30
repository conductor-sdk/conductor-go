package task_execute_function

import (
	"github.com/conductor-sdk/conductor-go/conductor_client/model"
	"github.com/conductor-sdk/conductor-go/conductor_client/model/enum/task_result_status"
	log "github.com/sirupsen/logrus"
)

func Example1(t *model.Task) (taskResult *model.TaskResult, err error) {
	log.Debug("Executing Task_1_Execution_Function for", t.TaskType)
	taskResult = model.NewTaskResult(t)
	taskResult.OutputData = map[string]interface{}{
		"task": "task_1",
		"key2": "value2",
		"key3": 3,
		"key4": false,
	}
	taskResult.Logs = append(taskResult.Logs, model.LogMessage{Log: "Hello World"})
	taskResult.Status = task_result_status.COMPLETED
	err = nil
	return taskResult, err
}

func Example2(t *model.Task) (taskResult *model.TaskResult, err error) {
	log.Debug("Executing Task_2_Execution_Function for", t.TaskType)
	taskResult = model.NewTaskResult(t)
	taskResult.OutputData = map[string]interface{}{
		"task": "task_2",
		"key2": "value2",
		"key3": 3,
		"key4": false,
	}
	taskResult.Status = task_result_status.COMPLETED
	return taskResult, nil
}
