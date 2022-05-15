package workflow

import (
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
)

func DoWhile(taskRefName string, terminationCondition string, tasks ...Task) *doWhile {
	loop := &doWhile{
		task: task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          DO_WHILE,
			optional:          false,
			inputParameters:   map[string]interface{}{},
		},
	}
	loop.loopCondition = terminationCondition
	loop.loopOver = tasks
	return loop
}

//Loop N times when N is specified as loopValue
// can be  static number e.g. 5 or a parameter experession like ${task_ref.output.some_value} that is a number
func Loop(taskRefName string, loopValue interface{}, tasks ...Task) *doWhile {
	loop := &doWhile{
		task: task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          DO_WHILE,
			optional:          false,
			inputParameters:   map[string]interface{}{},
		},
	}
	loop.task.Input("value", loopValue)
	loop.loopCondition = getForLoopCondition("value", taskRefName)
	loop.loopOver = tasks
	return loop
}

func getForLoopCondition(loopValue string, taskReferencename string) string {
	return "if ( $." + taskReferencename + "['iteration'] < $." + loopValue + ") { true; } else { false; }"
}

type doWhile struct {
	task
	loopCondition string
	loopOver      []Task
}

// Input to the task
func (task *doWhile) Input(key string, value interface{}) *doWhile {
	task.task.Input(key, value)
	return task
}

func (task *doWhile) toWorkflowTask() []http_model.WorkflowTask {
	workflowTasks := task.task.toWorkflowTask()
	workflowTasks[0].LoopCondition = task.loopCondition
	workflowTasks[0].LoopOver = []http_model.WorkflowTask{}
	for _, loopTask := range task.loopOver {
		workflowTasks[0].LoopOver = append(
			workflowTasks[0].LoopOver,
			loopTask.toWorkflowTask()...,
		)
	}
	return workflowTasks
}
