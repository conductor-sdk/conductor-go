package workflow

type WaitTask struct {
	Task
}

func Wait(taskRefName string) *WaitTask {
	return &WaitTask{
		Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          WAIT,
			optional:          false,
			inputParameters:   map[string]interface{}{},
		},
	}
}
