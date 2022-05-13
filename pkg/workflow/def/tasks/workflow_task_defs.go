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
}

type task struct {
	name              string
	taskReferenceName string
	description       string
	taskType          TaskType
	optional          bool
	inputParameters   map[string]interface{}
}

type simpleTask struct {
	task
}

func (t *simpleTask) Description(description string) *simpleTask {
	t.description = description
	return t
}
func (t *simpleTask) Optional(optional bool) *simpleTask {
	t.optional = optional
	return t
}

func (task *simpleTask) ToWorkflowTask() *http_model.WorkflowTask {
	return &http_model.WorkflowTask{
		Name:              task.name,
		TaskReferenceName: task.taskReferenceName,
		Description:       task.description,
		InputParameters:   task.inputParameters,
		Type_:             string(SIMPLE),
		Optional:          false,
		TaskDefinition:    nil,
	}
}

type wait struct {
	task
}

type decision struct {
	task
	decisionCases  map[string][]Task
	defaultCase    []Task
	caseExpression string
	useJavascript  bool
	evaluatorType  string
}

func (t *decision) Description(description string) *decision {
	t.description = description
	return t
}

func (task *decision) UseJavascript(use bool) {
	task.useJavascript = use
}

func (task *decision) SwitchCase(caseName string, tasks ...Task) *decision {
	task.decisionCases[caseName] = append(task.decisionCases[caseName], tasks...)
	return task
}
func (task *decision) DefaultCase(tasks ...Task) *decision {
	task.defaultCase = append(task.defaultCase, tasks...)
	return task
}
func (task *decision) ToWorkflowTask() *http_model.WorkflowTask {
	return &http_model.WorkflowTask{
		Name:              task.name,
		TaskReferenceName: task.taskReferenceName,
		Description:       task.description,
		InputParameters:   map[string]interface{}{},
		Type_:             string(task.taskType),
		CaseExpression:    task.caseExpression,
		DecisionCases:     map[string][]http_model.WorkflowTask{},
		DefaultCase:       nil,
		Optional:          false,
		EvaluatorType:     task.evaluatorType,
		Expression:        "",
	}
}
