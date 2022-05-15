package tasks

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
//loopValue can be  static number e.g. 5 or a parameter experession like ${task_ref.output.some_value} that is a number
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

func (task *doWhile) Description(description string) *doWhile {
	task.task.Description(description)
	return task
}

func (task *doWhile) Optional(optional bool) *doWhile {
	task.task.Optional(optional)
	return task
}

// Input input to the task
func (task *doWhile) Input(key string, value interface{}) *doWhile {
	task.task.Input(key, value)
	return task
}

func (task *doWhile) ToWorkflowTask() *http_model.WorkflowTask {
	workflowTask := task.task.ToWorkflowTask()
	workflowTask.LoopCondition = task.loopCondition
	workflowTask.LoopOver = []http_model.WorkflowTask{}
	for _, loopTask := range task.loopOver {
		workflowTask.LoopOver = append(workflowTask.LoopOver, *loopTask.ToWorkflowTask())
	}
	return workflowTask
}
