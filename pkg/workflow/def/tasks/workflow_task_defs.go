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
	name              string `json:"name"`
	taskReferenceName string `json:"taskReferenceName"`
	description       string
	taskType          TaskType `json:"type"`
	optional          bool     `json:"optional"`
	inputParameters   struct{} `json:"inputParameters"`
}

type simpleTask struct {
	task
}

func (t *simpleTask) Description(description string) *simpleTask {
	t.description = description
	return t
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
func (task *decision) toWorkflowTask() WorkflowTask {
	return WorkflowTask{
		Name:                           task.name,
		TaskReferenceName:              task.taskReferenceName,
		Description:                    task.description,
		InputParameters:                struct{}{},
		Type:                           string(task.taskType),
		CaseExpression:                 nil,
		ScriptExpression:               nil,
		DecisionCases:                  struct{}{},
		DynamicForkJoinTasksParam:      nil,
		DynamicForkTasksParam:          nil,
		DynamicForkTasksInputParamName: nil,
		DefaultCase:                    nil,
		StartDelay:                     0,
		Optional:                       false,
		EvaluatorType:                  task.evaluatorType,
		Expression:                     nil,
	}
}
