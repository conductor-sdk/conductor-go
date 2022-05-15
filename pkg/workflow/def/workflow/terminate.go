package workflow

import (
	workflow_status "github.com/conductor-sdk/conductor-go/pkg/model/enum"
)

type terminate struct {
	task
}

func Terminate(taskRefName string, terminationReason string, status workflow_status.WorkflowStatus) *terminate {
	terminate := &terminate{
		task: task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          TERMINATE,
			optional:          false,
			inputParameters:   map[string]interface{}{},
		},
	}
	terminate.task.Input("terminationStatus", status)
	terminate.task.Input("terminationReason", terminationReason)
	return terminate
}
