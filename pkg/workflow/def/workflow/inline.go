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

// Input to the task
func (task *InlineTask) Input(key string, value interface{}) *InlineTask {
	task.Task.Input(key, value)
	return task
}
func (task *InlineTask) Optional(optional bool) *InlineTask {
	task.Task.Optional(optional)
	return task
}
func (task *InlineTask) Description(description string) *InlineTask {
	task.Task.Description(description)
	return task
}
