package def

import (
	"fmt"
	"github.com/conductor-sdk/conductor-go/sdk/model"
)

const (
	loopCondition = "loop_count"
)

// DoWhileTask Do...While task
type DoWhileTask struct {
	Task
	loopCondition string
	loopOver      []TaskInterface
}

// NewDoWhileTask DoWhileTask Crate a new DoWhile task.
// terminationCondition is a Javascript expression that evaluates to True or False
func NewDoWhileTask(taskRefName string, terminationCondition string, tasks ...TaskInterface) *DoWhileTask {
	return &DoWhileTask{
		Task: Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			taskType:          DO_WHILE,
			inputParameters:   map[string]interface{}{},
		},
		loopCondition: terminationCondition,
		loopOver:      tasks,
	}
}

// NewLoopTask Loop over N times when N is specified as iterations
func NewLoopTask(taskRefName string, iterations int32, tasks ...TaskInterface) *DoWhileTask {
	return &DoWhileTask{
		Task: Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			taskType:          DO_WHILE,
			inputParameters: map[string]interface{}{
				loopCondition: iterations,
			},
		},
		loopCondition: getForLoopCondition(taskRefName, loopCondition),
		loopOver:      tasks,
	}
}

func (task *DoWhileTask) toWorkflowTask() []model.WorkflowTask {
	workflowTasks := task.Task.toWorkflowTask()
	workflowTasks[0].LoopCondition = task.loopCondition
	workflowTasks[0].LoopOver = []model.WorkflowTask{}
	for _, loopTask := range task.loopOver {
		workflowTasks[0].LoopOver = append(
			workflowTasks[0].LoopOver,
			loopTask.toWorkflowTask()...,
		)
	}
	return workflowTasks
}
func getForLoopCondition(loopValue string, taskReferenceName string) string {
	return fmt.Sprintf(
		"if ( $.%s['iteration'] < $.%s ) { true; } else { false; }",
		taskReferenceName, loopValue,
	)
}

// Optional if set to true, the task will not fail the workflow if the task fails
func (task *DoWhileTask) Optional(optional bool) *DoWhileTask {
	task.Task.Optional(optional)
	return task
}

// Input to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *DoWhileTask) Input(key string, value interface{}) *DoWhileTask {
	task.Task.Input(key, value)
	return task
}

// InputMap to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *DoWhileTask) InputMap(inputMap map[string]interface{}) *DoWhileTask {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}

// Description of the task
func (task *DoWhileTask) Description(description string) *DoWhileTask {
	task.Task.Description(description)
	return task
}
