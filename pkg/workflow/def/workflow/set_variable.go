package workflow

type SetVariableTask struct {
	Task
}

func NewSetVariableTask(taskRefName string) *SetVariableTask {
	return &SetVariableTask{
		Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          SET_VARIABLE,
			optional:          false,
			inputParameters:   map[string]interface{}{},
		},
	}
}

// Input to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *SetVariableTask) Input(key string, value interface{}) *SetVariableTask {
	task.Task.Input(key, value)
	return task
}

// InputMap to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *SetVariableTask) InputMap(inputMap map[string]interface{}) *SetVariableTask {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}

// Optional if set to true, the task will not fail the workflow if the task fails
func (task *SetVariableTask) Optional(optional bool) *SetVariableTask {
	task.Task.Optional(optional)
	return task
}

// Description of the task
func (task *SetVariableTask) Description(description string) *SetVariableTask {
	task.Task.Description(description)
	return task
}
