package def

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

// Input to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *JQTask) Input(key string, value interface{}) *JQTask {
	task.Task.Input(key, value)
	return task
}

// InputMap to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *JQTask) InputMap(inputMap map[string]interface{}) *JQTask {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}

// Optional if set to true, the task will not fail the workflow if the task fails
func (task *JQTask) Optional(optional bool) *JQTask {
	task.Task.Optional(optional)
	return task
}

// Description of the task
func (task *JQTask) Description(description string) *JQTask {
	task.Task.Description(description)
	return task
}