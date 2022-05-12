package tasks

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
	toWorkflowTask() WorkflowTask
}

type task struct {
	Name              string   `json:"name"`
	TaskReferenceName string   `json:"taskReferenceName"`
	Type              TaskType `json:"type"`
	Optional          bool     `json:"optional"`
	InputParameters   struct{} `json:"inputParameters"`
}

type simpleTask struct {
	task
}

func (task *simpleTask) toWorkflowTask() WorkflowTask {
	return WorkflowTask{}
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
func (task *decision) toWorkflowTask() WorkflowTask {
	return WorkflowTask{}
}
