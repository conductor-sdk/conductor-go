package examples

import (
	model2 "github.com/conductor-sdk/conductor-go/model"
)

func ExampleWorker(t *model2.Task) (interface{}, error) {
	taskResult := model2.GetTaskResultFromTask(t)
	taskResult.OutputData = map[string]interface{}{
		"key0": nil,
		"key1": 3,
		"key2": false,
		"foo":  "bar",
	}
	taskResult.Logs = append(
		taskResult.Logs,
		model2.TaskExecLog{
			Log: "log message",
		},
	)
	taskResult.Status = model2.COMPLETED
	return taskResult, nil
}

func SimpleWorker(t *model2.Task) (interface{}, error) {
	taskResult := model2.GetTaskResultFromTask(t)
	taskResult.OutputData = map[string]interface{}{
		"key": "value",
	}
	taskResult.Status = model2.COMPLETED
	return taskResult, nil
}

func OpenTreasureChest(t *model2.Task) (interface{}, error) {
	taskResult := model2.GetTaskResultFromTask(t)
	taskResult.OutputData = map[string]interface{}{
		"treasure": t.InputData["importantValue"],
	}
	taskResult.Status = model2.COMPLETED
	return taskResult, nil
}
