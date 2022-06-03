package workflow

type SimpleTask struct {
	Task
}

func NewSimpleTask(taskRefName string) *SimpleTask {
	return &SimpleTask{
		Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			taskType:          SIMPLE,
			inputParameters:   map[string]interface{}{},
		},
	}
}

func NewSimpleTaskWithInputParameters(taskRefName string, inputParameters map[string]interface{}) *SimpleTask {
	return &SimpleTask{
		Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			taskType:          SIMPLE,
			inputParameters:   inputParameters,
		},
	}
}
