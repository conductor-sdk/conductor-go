package workflow

type HumanTask struct {
	Task
}

func NewHumanTask(taskRefName string) *HumanTask {
	return &HumanTask{
		Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			taskType:          HUMAN,
			inputParameters:   map[string]interface{}{},
		},
	}
}

// Input to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *HumanTask) Input(key string, value interface{}) *HumanTask {
	task.Task.Input(key, value)
	return task
}

// InputMap to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *HumanTask) InputMap(inputMap map[string]interface{}) *HumanTask {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}

// Optional if set to true, the task will not fail the workflow if the task fails
func (task *HumanTask) Optional(optional bool) *HumanTask {
	task.Task.Optional(optional)
	return task
}

// Description of the task
func (task *HumanTask) Description(description string) *HumanTask {
	task.Task.Description(description)
	return task
}
