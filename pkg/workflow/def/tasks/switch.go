package tasks

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
func (task *decision) ToWorkflowTask() *http_model.WorkflowTask {

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
			decisionCases[caseValue] = append([]http_model.WorkflowTask{}, *task.ToWorkflowTask())
		}
	}
	var defaultCase []http_model.WorkflowTask
	for _, task := range task.defaultCase {
		defaultCase = append([]http_model.WorkflowTask{}, *task.ToWorkflowTask())
	}

	workflowTask := task.task.ToWorkflowTask()
	workflowTask.DecisionCases = decisionCases
	workflowTask.DefaultCase = defaultCase
	workflowTask.EvaluatorType = task.evaluatorType
	workflowTask.Expression = task.expression

	return workflowTask
}

// Input input to the task
func (task *decision) Input(key string, value interface{}) *decision {
	task.task.Input(key, value)
	return task
}
