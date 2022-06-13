package definition

import (
	"github.com/conductor-sdk/conductor-go/sdk/model"
)

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

func (task *SwitchTask) toWorkflowTask() []model.WorkflowTask {
	if task.useJavascript {
		task.evaluatorType = "javascript"
	} else {
		task.evaluatorType = "value-param"
		task.Task.inputParameters["switchCaseValue"] = task.expression
		task.expression = "switchCaseValue"
	}
	var DecisionCases = map[string][]model.WorkflowTask{}
	for caseValue, tasks := range task.DecisionCases {
		for _, task := range tasks {
			for _, caseTask := range task.toWorkflowTask() {
				DecisionCases[caseValue] = append([]model.WorkflowTask{}, caseTask)
			}
		}
	}
	var defaultCase []model.WorkflowTask
	for _, task := range task.defaultCase {
		for _, defaultTask := range task.toWorkflowTask() {
			defaultCase = append([]model.WorkflowTask{}, defaultTask)
		}
	}
	workflowTasks := task.Task.toWorkflowTask()
	workflowTasks[0].DecisionCases = DecisionCases
	workflowTasks[0].DefaultCase = defaultCase
	workflowTasks[0].EvaluatorType = task.evaluatorType
	workflowTasks[0].Expression = task.expression
	return workflowTasks
}

// Input to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *SwitchTask) Input(key string, value interface{}) *SwitchTask {
	task.Task.Input(key, value)
	return task
}

// InputMap to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *SwitchTask) InputMap(inputMap map[string]interface{}) *SwitchTask {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}

// Description of the task
func (task *SwitchTask) Description(description string) *SwitchTask {
	task.Task.Description(description)
	return task
}

// Optional if set to true, the task will not fail the workflow if the task fails
func (task *SwitchTask) Optional(optional bool) *SwitchTask {
	task.Task.Optional(optional)
	return task
}

// UseJavascript If set to to true, the caseExpression parameter is treated as a Javascript.
//If set to false, the caseExpression follows the regular task input mapping format as described in https://conductor.netflix.com/how-tos/Tasks/task-inputs.html
func (task *SwitchTask) UseJavascript(use bool) *SwitchTask {
	task.useJavascript = use
	return task
}
