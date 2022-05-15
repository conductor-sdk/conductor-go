package workflow

import "github.com/conductor-sdk/conductor-go/pkg/http_model"

func Switch(taskRefName string, caseExpression string) *decision {
	return &decision{
		task: task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          SWITCH,
			optional:          false,
			inputParameters:   map[string]interface{}{},
		},
		decisionCases: map[string][]Task{},
		defaultCase:   []Task{},
		expression:    caseExpression,
		useJavascript: false,
		evaluatorType: "value-param",
	}
}

type decision struct {
	task
	decisionCases map[string][]Task
	defaultCase   []Task
	expression    string
	useJavascript bool
	evaluatorType string
}

func (task *decision) Description(description string) *decision {
	task.task.Description(description)
	return task
}

func (task *decision) Optional(optional bool) *decision {
	task.task.Optional(optional)
	return task
}

func (task *decision) UseJavascript(use bool) *decision {
	task.useJavascript = use
	return task
}

func (task *decision) SwitchCase(caseName string, tasks ...Task) *decision {
	task.decisionCases[caseName] = append(task.decisionCases[caseName], tasks...)
	return task
}
func (task *decision) DefaultCase(tasks ...Task) *decision {
	task.defaultCase = append(task.defaultCase, tasks...)
	return task
}

func (task *decision) toWorkflowTask() []http_model.WorkflowTask {
	if task.useJavascript {
		task.evaluatorType = "javascript"
	} else {
		task.evaluatorType = "value-param"
		task.inputParameters["switchCaseValue"] = task.expression
		task.expression = "switchCaseValue"
	}

	var decisionCases = map[string][]http_model.WorkflowTask{}
	for caseValue, tasks := range task.decisionCases {
		for _, task := range tasks {
			for _, caseTask := range task.toWorkflowTask() {
				decisionCases[caseValue] = append([]http_model.WorkflowTask{}, caseTask)
			}
		}
	}
	var defaultCase []http_model.WorkflowTask
	for _, task := range task.defaultCase {
		for _, defaultTask := range task.toWorkflowTask() {
			defaultCase = append([]http_model.WorkflowTask{}, defaultTask)
		}
	}

	workflowTasks := task.task.toWorkflowTask()
	workflowTasks[0].DecisionCases = decisionCases
	workflowTasks[0].DefaultCase = defaultCase
	workflowTasks[0].EvaluatorType = task.evaluatorType
	workflowTasks[0].Expression = task.expression
	return workflowTasks
}

// Input to the task
func (task *decision) Input(key string, value interface{}) *decision {
	task.task.Input(key, value)
	return task
}
