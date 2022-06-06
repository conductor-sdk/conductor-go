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

// Input to the task
func (task *HumanTask) Input(key string, value interface{}) *HumanTask {
	task.Task.Input(key, value)
	return task
}
func (task *HumanTask) InputMap(inputMap map[string]interface{}) *HumanTask {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}
func (task *HumanTask) Optional(optional bool) *HumanTask {
	task.Task.Optional(optional)
	return task
}
func (task *HumanTask) Description(description string) *HumanTask {
	task.Task.Description(description)
	return task
}
