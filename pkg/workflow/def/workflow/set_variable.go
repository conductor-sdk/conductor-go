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
