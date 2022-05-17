package workflow

import (
	"fmt"

	"github.com/conductor-sdk/conductor-go/pkg/http_model"
)

var (
	LOOP_CONDITION = "VALUE"
)

type DoWhileTask struct {
	task          Task
	loopCondition string
	loopOver      []Task
}

func DoWhile(taskRefName string, terminationCondition string, tasks []Task) *DoWhileTask {
	return &DoWhileTask{
		task: Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          DO_WHILE,
			optional:          false,
			inputParameters:   map[string]interface{}{},
		},
		loopCondition: terminationCondition,
		loopOver:      tasks,
	}
}

// Loop over N times when N is specified as iterations
// can be  static number e.g. 5 or a parameter expression like ${task_ref.output.some_value} that is a number
func Loop(taskRefName string, iterations string, tasks []Task) *DoWhileTask {
	return &DoWhileTask{
		task: Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          DO_WHILE,
			optional:          false,
			inputParameters: map[string]interface{}{
				"value": iterations,
			},
		},
		loopCondition: getForLoopCondition(LOOP_CONDITION, iterations),
		loopOver:      tasks,
	}
}

// Input to the task
func (task *DoWhileTask) Input(key string, value interface{}) *DoWhileTask {
	task.task.Input(key, value)
	return task
}

func (task *DoWhileTask) toWorkflowTask() []http_model.WorkflowTask {
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
func (task *DoWhileTask) Optional(optional bool) *DoWhileTask {
	task.task.Optional(optional)
	return task
}

func getForLoopCondition(taskReferencename string, iterations string) string {
	return fmt.Sprintf(
		"if ( $.%s['iteration'] < $.%s ) { true; } else { false; }",
		taskReferencename, iterations,
	)
}
