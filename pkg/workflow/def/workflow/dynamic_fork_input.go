package workflow

import (
	"github.com/conductor-sdk/conductor-go/pkg/model"
)

type DynamicForkInput struct {
	Tasks     []model.WorkflowTask
	TaskInput map[string]interface{}
}

type DynamicForkInputTask interface {
	Execute(input interface{}) *DynamicForkInput
}
