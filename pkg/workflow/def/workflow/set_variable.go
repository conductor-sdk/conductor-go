package workflow

func SetVariable(taskRefName string) *setVariable {
	return &setVariable{
		task: Task{
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
	task Task
}

func (task *setVariable) Description(description string) *setVariable {
	task.task.Description(description)
	return task
}

func (task *setVariable) Optional(optional bool) *setVariable {
	task.task.Optional(optional)
	return task
}
func (task *setVariable) Input(key string, value interface{}) *setVariable {
	task.task.Input(key, value)
	return task
}
