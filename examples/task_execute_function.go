package examples

import (
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/conductor-sdk/conductor-go/pkg/model"
	"github.com/conductor-sdk/conductor-go/pkg/model/enum/task_result_status"
)

func ExampleWorker(t *http_model.Task) (taskResult *http_model.TaskResult, err error) {
	taskResult = model.GetTaskResultFromTask(t)
	taskResult.OutputData = map[string]interface{}{
		"task": "task_1",
		"key2": "value2",
		"key3": 3,
		"key4": false,
	}
	taskResult.Logs = append(
		taskResult.Logs,
		http_model.TaskExecLog{
			Log: "Hello World",
		},
	)
	taskResult.Status = task_result_status.COMPLETED
	err = nil
	return taskResult, err
}

func SimpleWorker(t *http_model.Task) (taskResult *http_model.TaskResult, err error) {
	taskResult = model.GetTaskResultFromTask(t)
	taskResult.OutputData = map[string]interface{}{
		"key": "value",
	}
	taskResult.Status = task_result_status.COMPLETED
	return taskResult, nil
}

func OpenTreasureChest(t *http_model.Task) (taskResult *http_model.TaskResult, err error) {
	taskResult = model.GetTaskResultFromTask(t)
	taskResult.OutputData = map[string]interface{}{
		"treasure": t.InputData["importantValue"],
	}
	taskResult.Status = task_result_status.COMPLETED
	return taskResult, nil
}
