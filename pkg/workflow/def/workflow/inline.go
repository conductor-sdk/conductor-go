package workflow

type InlineTask struct {
	Task
}

const (
	JavascriptEvaluator = "javascript"
)

func NewInlineTask(name string, script string) *InlineTask {
	return &InlineTask{
		Task{
			name:              name,
			taskReferenceName: name,
			taskType:          INLINE,
			inputParameters: map[string]interface{}{
				"evaluatorType": JavascriptEvaluator,
				"expression":    script,
			},
		},
	}
}

// Input to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *InlineTask) Input(key string, value interface{}) *InlineTask {
	task.Task.Input(key, value)
	return task
}

// InputMap to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *InlineTask) InputMap(inputMap map[string]interface{}) *InlineTask {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}

// Optional if set to true, the task will not fail the workflow if the task fails
func (task *InlineTask) Optional(optional bool) *InlineTask {
	task.Task.Optional(optional)
	return task
}

// Description of the task
func (task *InlineTask) Description(description string) *InlineTask {
	task.Task.Description(description)
	return task
}
