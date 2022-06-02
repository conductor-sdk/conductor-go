package workflow

type SimpleTask struct {
	Task
}

func NewSimpleTask(name string, taskRefName string) *SimpleTask {
	return &SimpleTask{
		Task{
			name:              name,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          SIMPLE,
			optional:          false,
			inputParameters:   map[string]interface{}{},
		},
	}
}
