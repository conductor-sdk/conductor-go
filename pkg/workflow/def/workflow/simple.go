package workflow

type SimpleTask struct {
	Task
}

func NewSimpleTask(taskType string, taskRefName string) *SimpleTask {
	return &SimpleTask{
		Task{
			name:              taskType,
			taskReferenceName: taskRefName,
			taskType:          SIMPLE,
			inputParameters:   map[string]interface{}{},
		},
	}
}

// Input to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *SimpleTask) Input(key string, value interface{}) *SimpleTask {
	task.Task.Input(key, value)
	return task
}

// InputMap to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *SimpleTask) InputMap(inputMap map[string]interface{}) *SimpleTask {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}

// Optional if set to true, the task will not fail the workflow if the task fails
func (task *SimpleTask) Optional(optional bool) *SimpleTask {
	task.Task.Optional(optional)
	return task
}

// Description of the task
func (task *SimpleTask) Description(description string) *SimpleTask {
	task.Task.Description(description)
	return task
}
