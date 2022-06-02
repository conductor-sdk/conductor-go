package workflow

type InlineTask struct {
	Task
}

func NewInlineTask(name string, inputParameters map[string]interface{}) *InlineTask {
	return &InlineTask{
		Task{
			name:              name,
			taskReferenceName: name,
			taskType:          INLINE,
			inputParameters:   inputParameters,
		},
	}
}
