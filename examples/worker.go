package examples

import (
	"github.com/conductor-sdk/conductor-go/pkg/model"
)

func ExampleWorker(t *model.Task) (interface{}, error) {
	taskResult := model.GetTaskResultFromTask(t)
	taskResult.OutputData = map[string]interface{}{
		"key0": nil,
		"key1": 3,
		"key2": false,
		"foo":  "bar",
	}
	taskResult.Logs = append(
		taskResult.Logs,
		model.TaskExecLog{
			Log: "log message",
		},
	)
	taskResult.Status = model.COMPLETED
	return taskResult, nil
}

func SimpleWorker(t *model.Task) (interface{}, error) {
	taskResult := model.GetTaskResultFromTask(t)
	taskResult.OutputData = map[string]interface{}{
		"key": "value",
	}
	taskResult.Status = model.COMPLETED
	return taskResult, nil
}

func OpenTreasureChest(t *model.Task) (interface{}, error) {
	taskResult := model.GetTaskResultFromTask(t)
	taskResult.OutputData = map[string]interface{}{
		"treasure": t.InputData["importantValue"],
	}
	taskResult.Status = model.COMPLETED
	return taskResult, nil
}
