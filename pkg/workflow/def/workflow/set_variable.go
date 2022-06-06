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

// Input to the task
func (task *SetVariableTask) Input(key string, value interface{}) *SetVariableTask {
	task.Task.Input(key, value)
	return task
}
func (task *SetVariableTask) InputMap(inputMap map[string]interface{}) *SetVariableTask {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}
func (task *SetVariableTask) Optional(optional bool) *SetVariableTask {
	task.Task.Optional(optional)
	return task
}
func (task *SetVariableTask) Description(description string) *SetVariableTask {
	task.Task.Description(description)
	return task
}
