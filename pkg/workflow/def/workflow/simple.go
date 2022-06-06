package workflow

type SimpleTask struct {
	Task
}

func NewSimpleTask(taskRefName string) *SimpleTask {
	return &SimpleTask{
		Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			taskType:          SIMPLE,
			inputParameters:   map[string]interface{}{},
		},
	}
}

// Input to the task
func (task *SimpleTask) Input(key string, value interface{}) *SimpleTask {
	task.Task.Input(key, value)
	return task
}
func (task *SimpleTask) InputMap(inputMap map[string]interface{}) *SimpleTask {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}
func (task *SimpleTask) Optional(optional bool) *SimpleTask {
	task.Task.Optional(optional)
	return task
}
func (task *SimpleTask) Description(description string) *SimpleTask {
	task.Task.Description(description)
	return task
}
