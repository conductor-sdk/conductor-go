package workflow

import (
	"time"
)

type WaitTask struct {
	Task
}

//NewWaitTask creates WAIT task used to wait until an external event or a timeout occurs
func NewWaitTask(taskRefName string) *WaitTask {
	return &WaitTask{
		Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          WAIT,
			optional:          false,
			inputParameters:   map[string]interface{}{},
		},
	}
}
func NewWaitForDurationTask(taskRefName string, duration time.Duration) *WaitTask {
	return &WaitTask{
		Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          WAIT,
			optional:          false,
			inputParameters: map[string]interface{}{
				"duration": duration.String(),
			},
		},
	}
}

func NewWaitUntilTask(taskRefName string, dateTime string) *WaitTask {
	return &WaitTask{
		Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          WAIT,
			optional:          false,
			inputParameters: map[string]interface{}{
				"until": dateTime,
			},
		},
	}
}

// Description of the task
func (task *WaitTask) Description(description string) *WaitTask {
	task.Task.Description(description)
	return task
}

// Optional if set to true, the task will not fail the workflow if the task fails
func (task *WaitTask) Optional(optional bool) *WaitTask {
	task.Task.Optional(optional)
	return task
}

// Input to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *WaitTask) Input(key string, value interface{}) *WaitTask {
	task.Task.Input(key, value)
	return task
}

// InputMap to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *WaitTask) InputMap(inputMap map[string]interface{}) *WaitTask {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}
