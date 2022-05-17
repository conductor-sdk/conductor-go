package workflow

import "github.com/conductor-sdk/conductor-go/pkg/http_model"

type ForkTask struct {
	Task
	forkedTasks [][]TaskInterface
}

func Fork(taskRefName string, tasks [][]TaskInterface) *ForkTask {
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

func (task *ForkTask) toWorkflowTask() []http_model.WorkflowTask {
	forkWorkflowTask := task.Task.toWorkflowTask()[0]
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

func (task *ForkTask) getJoinTask() http_model.WorkflowTask {
	join := Join(task.taskReferenceName + "_join")
	return (join.toWorkflowTask())[0]
}
