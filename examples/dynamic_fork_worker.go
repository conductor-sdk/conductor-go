package examples

import "github.com/conductor-sdk/conductor-go/sdk/model"

type WorkflowTask struct {
	Name              string `json:"name"`
	TaskReferenceName string `json:"taskReferenceName"`
	Type              string `json:"type,omitempty"`
}

//dynamic_fork_prep
func DynamicForkWorker(t *model.Task) (output interface{}, err error) {
	taskResult := model.NewTaskResultFromTask(t)
	tasks := []WorkflowTask{
		{
			Name:              "simple_task_1",
			TaskReferenceName: "simple_task_11",
			Type:              "SIMPLE",
		},
		{
			Name:              "simple_task_3",
			TaskReferenceName: "simple_task_12",
			Type:              "SIMPLE",
		},
		{
			Name:              "simple_task_5",
			TaskReferenceName: "simple_task_13",
			Type:              "SIMPLE",
		},
	}
	inputs := map[string]interface{}{
		"simple_task_11": map[string]interface{}{
			"key1": "value1",
			"key2": 121,
		},
		"simple_task_12": map[string]interface{}{
			"key1": "value2",
			"key2": 122,
		},
		"simple_task_13": map[string]interface{}{
			"key1": "value3",
			"key2": 123,
		},
	}

	taskResult.OutputData = map[string]interface{}{
		"forkedTasks":       tasks,
		"forkedTasksInputs": inputs,
	}
	taskResult.Status = model.CompletedTask
	err = nil
	return taskResult, err
}
