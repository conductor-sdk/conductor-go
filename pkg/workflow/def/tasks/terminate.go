package tasks

import (
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	workflow_status "github.com/conductor-sdk/conductor-go/pkg/model/enum"
)

func Terminate(taskRefName string, terminationReason string, status workflow_status.WorkflowStatus) *decision {
	decision := &decision{
		task: task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          TERMINATE,
			optional:          false,
			inputParameters:   map[string]interface{}{},
		},
	}
	decision.task.Input("terminationStatus", status)
	decision.task.Input("terminationReason", terminationReason)
	return decision
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

func (task *terminate) ToWorkflowTask() *http_model.WorkflowTask {
	return task.task.ToWorkflowTask()
}
