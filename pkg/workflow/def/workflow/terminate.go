package workflow

import (
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	workflow_status "github.com/conductor-sdk/conductor-go/pkg/model/enum"
)

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

type terminate struct {
	task
}

func (task *terminate) Description(description string) *terminate {
	task.task.Description(description)
	return task
}

func (task *terminate) Optional(optional bool) *terminate {
	task.task.Optional(optional)
	return task
}

func (task *terminate) toWorkflowTask() *[]http_model.WorkflowTask {
	return task.task.toWorkflowTask()
}
