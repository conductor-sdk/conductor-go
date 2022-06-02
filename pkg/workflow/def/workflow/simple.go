package workflow

type SimpleTask struct {
	Task
}

func NewSimpleTask(taskName string) *SimpleTask {
	return &SimpleTask{
		Task{
			name:              taskName,
			taskReferenceName: taskName,
			taskType:          SIMPLE,
			inputParameters:   map[string]interface{}{},
		},
	}
}

func NewSimpleTaskWithInputParameters(taskName string, inputParameters map[string]interface{}) *SimpleTask {
	return &SimpleTask{
		Task{
			name:              taskName,
			taskReferenceName: taskName,
			taskType:          SIMPLE,
			inputParameters:   inputParameters,
		},
	}
}
