package workflow

import "github.com/conductor-sdk/conductor-go/pkg/http_model"

func Fork(taskRefName string, tasks ...[]Task) *fork {
	return &fork{
		task: task{
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

type fork struct {
	task
	forkedTasks [][]Task
	joinTask    Task
}

func (task *fork) toWorkflowTask() []http_model.WorkflowTask {
	forkWorkflowTask := task.task.toWorkflowTask()[0]
	forkWorkflowTask.ForkTasks = make([][]http_model.WorkflowTask, len(task.forkedTasks))
	for i, forkedTask := range task.forkedTasks {
		forkWorkflowTask.ForkTasks[i] = make([]http_model.WorkflowTask, len(forkedTask))
		for j, innerForkedTask := range forkedTask {
			forkWorkflowTask.ForkTasks[i][j] = innerForkedTask.toWorkflowTask()[0]
		}
	}
	return []http_model.WorkflowTask{
		forkWorkflowTask,
		task.getJoinTask(),
	}
}

func (task *fork) getJoinTask() http_model.WorkflowTask {
	if task.joinTask != nil {
		return (task.joinTask.toWorkflowTask())[0]
	}
	join := Join(task.taskReferenceName + "_join")
	return (join.toWorkflowTask())[0]
}
