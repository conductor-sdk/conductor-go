package example

import (
	"github.com/netflix/conductor/client/go/conductor_client/model"
	"github.com/netflix/conductor/client/go/conductor_client/model/enum/task_result_status"
	log "github.com/sirupsen/logrus"
)

func TaskExecuteFunctionExample1(t *model.Task) (taskResult *model.TaskResult, err error) {
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

func TaskExecuteFunctionExample2(t *model.Task) (taskResult *model.TaskResult, err error) {
	log.Println("Executing Task_2_Execution_Function for", t.TaskType)
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
