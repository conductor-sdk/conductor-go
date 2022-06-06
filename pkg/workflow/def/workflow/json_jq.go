package workflow

type JQTask struct {
	Task
}

func NewJQTask(name string, script string) *JQTask {
	return &JQTask{
		Task{
			name:              name,
			taskReferenceName: name,
			taskType:          JSON_JQ_TRANSFORM,
			inputParameters: map[string]interface{}{
				"queryExpression": script,
			},
		},
	}
}

// Input to the task
func (task *JQTask) Input(key string, value interface{}) *JQTask {
	task.Task.Input(key, value)
	return task
}
func (task *JQTask) Optional(optional bool) *JQTask {
	task.Task.Optional(optional)
	return task
}
func (task *JQTask) Description(description string) *JQTask {
	task.Task.Description(description)
	return task
}
