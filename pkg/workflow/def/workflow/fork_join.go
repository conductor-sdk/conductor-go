package workflow

import (
	"github.com/conductor-sdk/conductor-go/pkg/model"
)

type ForkTask struct {
	Task
	forkedTasks [][]TaskInterface
}

func NewForkTask(taskRefName string, tasks ...[]TaskInterface) *ForkTask {
	return &ForkTask{
		Task: Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          FORK_JOIN,
			optional:          false,
			inputParameters:   map[string]interface{}{},
		},
		forkedTasks: tasks,
	}
}

func (task *ForkTask) toWorkflowTask() []model.WorkflowTask {
	forkWorkflowTask := task.Task.toWorkflowTask()[0]
	forkWorkflowTask.ForkTasks = make([][]model.WorkflowTask, len(task.forkedTasks))
	for i, forkedTask := range task.forkedTasks {
		forkWorkflowTask.ForkTasks[i] = make([]model.WorkflowTask, len(forkedTask))
		for j, innerForkedTask := range forkedTask {
			forkWorkflowTask.ForkTasks[i][j] = innerForkedTask.toWorkflowTask()[0]
		}
	}
	return []model.WorkflowTask{
		forkWorkflowTask,
		task.getJoinTask(),
	}
}

func (task *ForkTask) getJoinTask() model.WorkflowTask {
	join := NewJoinTask(task.taskReferenceName + "_join")
	return (join.toWorkflowTask())[0]
}

// Input to the task
func (task *ForkTask) Input(key string, value interface{}) *ForkTask {
	task.Task.Input(key, value)
	return task
}
func (task *ForkTask) InputMap(inputMap map[string]interface{}) *ForkTask {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}
func (task *ForkTask) Optional(optional bool) *ForkTask {
	task.Task.Optional(optional)
	return task
}
func (task *ForkTask) Description(description string) *ForkTask {
	task.Task.Description(description)
	return task
}
