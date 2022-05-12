package tasks

func SimpleTask(name string, taskRefName string) *simpleTask {
	return &simpleTask{task{
		name:              name,
		taskReferenceName: taskRefName,
		taskType:          SIMPLE,
		optional:          false,
		inputParameters:   struct{}{},
	}}
}
func Switch(taskRefName string, caseExpression string) *decision {
	return &decision{
		task: task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          SWITCH,
			optional:          false,
			inputParameters:   struct{}{},
		},
		decisionCases:  nil,
		defaultCase:    nil,
		caseExpression: caseExpression,
		useJavascript:  false,
		evaluatorType:  "value-param",
	}
}

func example() {
	decision := Switch("shipping", "${workflow.input.shipping")
	decision.SwitchCase("Ground",
		SimpleTask("ship", "ship"),
		SimpleTask("wait_shipping", "wait_shipping"))
	decision.SwitchCase("Air", SimpleTask("air", "air"))
}
