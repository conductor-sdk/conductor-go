package workflow

type SetVariableTask struct {
	Task
}

func SetVariable(taskRefName string) *SetVariableTask {
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
