package workflow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToWorkflowTaskValueParamEvaluator(t *testing.T) {
	caseExpr := "${workflow.input.notificationPref}"
	task := NewSwitchTask("emailorsms", caseExpr).
		SwitchCase("EMAIL", createSendEmailTask()).
		SwitchCase("SMS", createSendSMSTask())

	wfTasks := task.toWorkflowTask()

	// toWorkflowTask doesn't modify the task
	assert.Equal(t, caseExpr, task.expression)
	assert.False(t, task.useJavascript)
	assert.Len(t, task.inputParameters, 0)
	assert.Len(t, task.defaultCase, 0)
	assert.Len(t, task.DecisionCases, 2)

	// workflowTask instances
	assert.Len(t, wfTasks, 1)

	wfTask := wfTasks[0]
	assert.Equal(t, "value-param", wfTask.EvaluatorType)
	assert.Equal(t, "switchCaseValue", wfTask.Expression)
	assert.Len(t, wfTask.InputParameters, 1)
	assert.Equal(t, caseExpr, wfTask.InputParameters["switchCaseValue"])
	assert.Len(t, wfTask.DefaultCase, 0)
	assert.Len(t, wfTask.DecisionCases, 2)
	assert.NotNil(t, wfTask.DecisionCases["EMAIL"])
	assert.NotNil(t, wfTask.DecisionCases["SMS"])
}

func TestToWorkflowTaskJSEvaluator(t *testing.T) {
	caseExpr := "${workflow.input.notificationPref}"
	task := NewSwitchTask("emailorsms", caseExpr).
		SwitchCase("EMAIL", createSendEmailTask()).
		SwitchCase("SMS", createSendSMSTask()).
		UseJavascript(true)

	wfTasks := task.toWorkflowTask()

	// toWorkflowTask doesn't modify the task
	assert.Equal(t, caseExpr, task.expression)
	assert.True(t, task.useJavascript)
	assert.Len(t, task.inputParameters, 0)
	assert.Len(t, task.defaultCase, 0)
	assert.Len(t, task.DecisionCases, 2)

	// workflowTask instances
	assert.Len(t, wfTasks, 1)

	wfTask := wfTasks[0]
	assert.Equal(t, "javascript", wfTask.EvaluatorType)
	assert.Equal(t, caseExpr, wfTask.Expression)
	assert.Len(t, wfTask.InputParameters, 0)
	assert.Len(t, wfTask.DefaultCase, 0)
	assert.Len(t, wfTask.DecisionCases, 2)
	assert.NotNil(t, wfTask.DecisionCases["EMAIL"])
	assert.NotNil(t, wfTask.DecisionCases["SMS"])
}

func createSendEmailTask() TaskInterface {
	return NewSimpleTask("send_email", "send_email").
		Input("email", "${get_user_info.output.email}")
}

func createSendSMSTask() TaskInterface {
	return NewSimpleTask("send_sms", "send_sms").
		Input("phoneNumber", "${get_user_info.output.phoneNumber}")
}
