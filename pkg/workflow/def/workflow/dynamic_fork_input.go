package workflow

import "github.com/conductor-sdk/conductor-go/pkg/http_model"

type DynamicForkInput struct {
	Tasks     []http_model.WorkflowTask
	TaskInput map[string]interface{}
}

type DynamicForkInputTask interface {
	Execute(input interface{}) *DynamicForkInput
}
