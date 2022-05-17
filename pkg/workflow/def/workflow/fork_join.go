package workflow

import "github.com/conductor-sdk/conductor-go/pkg/http_model"

type ForkTask struct {
	task        Task
	forkedTasks [][]Task
	// joinTask    *Task
}

func Fork(taskRefName string, tasks [][]Task, inputParameters map[string]interface{}) *ForkTask {
	return &ForkTask{
		task: Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          FORK_JOIN,
			optional:          false,
			inputParameters:   inputParameters,
		},
		forkedTasks: tasks,
	}
}

func (task *ForkTask) toWorkflowTask() []http_model.WorkflowTask {
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
		// task.getJoinTask(),
	}
}

// func (task *ForkTask) getJoinTask() http_model.WorkflowTask {
// if task.joinTask != nil {
// 	return (task.joinTask.toWorkflowTask())[0]
// }
// join := Join(task.task.taskReferenceName + "_join")
// return (join.toWorkflowTask())[0]
// }
