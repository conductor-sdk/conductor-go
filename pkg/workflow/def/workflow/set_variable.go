package workflow

func SetVariable(taskRefName string) *setVariable {
	return &setVariable{
		task: task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          SET_VARIABLE,
			optional:          false,
			inputParameters:   map[string]interface{}{},
		},
	}
}

type setVariable struct {
	task
}

func (task *setVariable) Description(description string) *setVariable {
	task.task.Description(description)
	return task
}

func (task *setVariable) Optional(optional bool) *setVariable {
	task.task.Optional(optional)
	return task
}
