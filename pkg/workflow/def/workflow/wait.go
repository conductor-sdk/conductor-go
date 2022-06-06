package workflow

type WaitTask struct {
	Task
}

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
func (task *WaitTask) Description(description string) *WaitTask {
	task.Task.Description(description)
	return task
}
func (task *WaitTask) Optional(optional bool) *WaitTask {
	task.Task.Optional(optional)
	return task
}
func (task *WaitTask) Input(key string, value interface{}) *WaitTask {
	task.Task.Input(key, value)
	return task
}
func (task *WaitTask) InputMap(inputMap map[string]interface{}) *WaitTask {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}
