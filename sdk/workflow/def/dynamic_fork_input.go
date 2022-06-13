package def

import (
	"github.com/conductor-sdk/conductor-go/sdk/model"
)

//DynamicForkInput struct that represents the output of the dynamic fork preparatory task
//
//DynamicFork requires inputs that specifies the list of tasks to be executed in parallel along with the inputs
//to each of these tasks.
//This usually means having another task before the dynamic task whose output contains
//the tasks to be forked.
//
//This struct represents the output of such a task
type DynamicForkInput struct {
	Tasks     []model.WorkflowTask
	TaskInput map[string]interface{}
}
