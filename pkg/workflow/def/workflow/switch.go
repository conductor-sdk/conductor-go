package workflow

import "github.com/conductor-sdk/conductor-go/pkg/http_model"

type SwitchTask struct {
	Task
	DecisionCases map[string][]TaskInterface
	defaultCase   []TaskInterface
	expression    string
	useJavascript bool
	evaluatorType string
}

func NewSwitchTask(taskRefName string, caseExpression string) *SwitchTask {
	return &SwitchTask{
		Task: Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          SWITCH,
			optional:          false,
			inputParameters:   map[string]interface{}{},
		},
		DecisionCases: make(map[string][]TaskInterface),
		defaultCase:   make([]TaskInterface, 0),
		expression:    caseExpression,
		useJavascript: false,
		evaluatorType: "value-param",
	}
}

func (task *SwitchTask) SwitchCase(caseName string, tasks ...TaskInterface) *SwitchTask {
	task.DecisionCases[caseName] = tasks
	return task
}
func (task *SwitchTask) DefaultCase(tasks ...TaskInterface) *SwitchTask {
	task.defaultCase = tasks
	return task
}

func (task *SwitchTask) toWorkflowTask() []http_model.WorkflowTask {
	if task.useJavascript {
		task.evaluatorType = "javascript"
	} else {
		task.evaluatorType = "value-param"
		task.Task.inputParameters["switchCaseValue"] = task.expression
		task.expression = "switchCaseValue"
	}
	var DecisionCases = map[string][]http_model.WorkflowTask{}
	for caseValue, tasks := range task.DecisionCases {
		for _, task := range tasks {
			for _, caseTask := range task.toWorkflowTask() {
				DecisionCases[caseValue] = append([]http_model.WorkflowTask{}, caseTask)
			}
		}
	}
	var defaultCase []http_model.WorkflowTask
	for _, task := range task.defaultCase {
		for _, defaultTask := range task.toWorkflowTask() {
			defaultCase = append([]http_model.WorkflowTask{}, defaultTask)
		}
	}
	workflowTasks := task.Task.toWorkflowTask()
	workflowTasks[0].DecisionCases = DecisionCases
	workflowTasks[0].DefaultCase = defaultCase
	workflowTasks[0].EvaluatorType = task.evaluatorType
	workflowTasks[0].Expression = task.expression
	return workflowTasks
}

// Input to the task
func (task *SwitchTask) Input(key string, value interface{}) *SwitchTask {
	task.Task.Input(key, value)
	return task
}
func (task *SwitchTask) InputMap(inputMap map[string]interface{}) *SwitchTask {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}
func (task *SwitchTask) Description(description string) *SwitchTask {
	task.Task.Description(description)
	return task
}

func (task *SwitchTask) Optional(optional bool) *SwitchTask {
	task.Task.Optional(optional)
	return task
}

func (task *SwitchTask) UseJavascript(use bool) *SwitchTask {
	task.useJavascript = use
	return task
}
