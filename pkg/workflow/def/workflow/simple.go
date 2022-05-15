package workflow

type simpleTask struct {
	task
}

func Simple(name string, taskRefName string) *simpleTask {
	return &simpleTask{
		task{
			name:              name,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          SIMPLE,
			optional:          false,
			inputParameters:   map[string]interface{}{},
		},
	}
}
