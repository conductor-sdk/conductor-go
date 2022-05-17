package workflow

import "github.com/conductor-sdk/conductor-go/pkg/http_model"

type TaskType string

const (
	SIMPLE            TaskType = "SIMPLE"
	DYNAMIC           TaskType = "DYNAMIC"
	FORK_JOIN         TaskType = "FORK_JOIN"
	FORK_JOIN_DYNAMIC TaskType = "FORK_JOIN_DYNAMIC"
	SWITCH            TaskType = "SWITCH"
	JOIN              TaskType = "JOIN"
	DO_WHILE          TaskType = "DO_WHILE"
	SUB_WORKFLOW      TaskType = "SUB_WORKFLOW"
	START_WORKFLOW    TaskType = "START_WORKFLOW"
	EVENT             TaskType = "EVENT"
	WAIT              TaskType = "WAIT"
	HTTP              TaskType = "HTTP"
	INLINE            TaskType = "INLINE"
	EXCLUSIVE_JOIN    TaskType = "EXCLUSIVE_JOIN"
	TERMINATE         TaskType = "TERMINATE"
	KAFKA_PUBLISH     TaskType = "KAFKA_PUBLISH"
	JSON_JQ_TRANSFORM TaskType = "JSON_JQ_TRANSFORM"
	SET_VARIABLE      TaskType = "SET_VARIABLE"
)

type TaskInterface interface {
	toWorkflowTask() []http_model.WorkflowTask
	OutputRef(path string) string
}

type Task struct {
	name              string
	taskReferenceName string
	description       string
	taskType          TaskType
	optional          bool
	inputParameters   map[string]interface{}
}

func (task *Task) toWorkflowTask() []http_model.WorkflowTask {
	return []http_model.WorkflowTask{
		{
			Name:              task.name,
			TaskReferenceName: task.taskReferenceName,
			Description:       task.description,
			InputParameters:   task.inputParameters,
			Optional:          task.optional,
			Type_:             string(task.taskType),
		},
	}
}

func (task *Task) ReferenceName() string {
	return task.taskReferenceName
}
func (task *Task) OutputRef(path string) string {
	return "${" + task.taskReferenceName + ".output." + path + "}"
}

//Note: All the below method should be implemented by the
//Implementing interface given its a fluent interface
//If not, the return type is a Task which makes it impossible to use fluent interface
//For the tasks like Switch which has other methods too - quirks with Golang!

// Input to the task
func (task *Task) Input(key string, value interface{}) *Task {
	task.inputParameters[key] = value
	return task
}

func (task *Task) Description(description string) *Task {
	task.description = description
	return task
}
func (task *Task) Optional(optional bool) *Task {
	task.optional = optional
	return task
}
