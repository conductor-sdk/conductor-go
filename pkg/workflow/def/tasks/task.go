package tasks

import "github.com/conductor-sdk/conductor-go/pkg/http_model"

type TaskType string

const (
	SIMPLE            TaskType = "SIMPLE"
	DYNAMIC           TaskType = "DYNAMIC"
	FORK_JOIN         TaskType = "FORK_JOIN"
	FORK_JOIN_DYNAMIC TaskType = "FORK_JOIN_DYNAMIC"
	DECISION          TaskType = "DECISION"
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

type Task interface {
	ToWorkflowTask() *http_model.WorkflowTask
	//Future methods
	//validate() error_string each task should implement this to check for correctness
}

type task struct {
	name              string
	taskReferenceName string
	description       string
	taskType          TaskType
	optional          bool
	inputParameters   map[string]interface{}
}

func (task *task) ToWorkflowTask() *http_model.WorkflowTask {

	return &http_model.WorkflowTask{
		Name:              task.name,
		TaskReferenceName: task.taskReferenceName,
		Description:       task.description,
		InputParameters:   task.inputParameters,
		Optional:          task.optional,
		TaskDefinition:    nil,
		Type_:             string(task.taskType),
	}
}
func (task *task) Description(description string) *task {
	task.description = description
	return task
}
func (task *task) Optional(optional bool) *task {
	task.optional = optional
	return task
}

// Input input to the task
func (task *task) Input(key string, value interface{}) *task {
	task.inputParameters[key] = value
	return task
}
