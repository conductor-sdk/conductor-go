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

func (task *fork) Description(description string) *fork {
	task.task.Description(description)
	return task
}

func (task *fork) Optional(optional bool) *fork {
	task.task.Optional(optional)
	return task
}
func (task *fork) JoinOn(joinOn ...string) *fork {
	task.joinTask = Join(task.taskReferenceName+"_join", joinOn...)
	return task
}

func (task *fork) toWorkflowTask() *[]http_model.WorkflowTask {
	workflowTasks := task.task.toWorkflowTask()

	for _, forkedTask := range task.forkedTasks {
		var forkedTasksInner []http_model.WorkflowTask
		for _, t := range forkedTask {
			wts := t.toWorkflowTask()
			for _, wt := range *wts {
				forkedTasksInner = append(forkedTasksInner, wt)
			}
		}
		(*workflowTasks)[0].ForkTasks = append((*workflowTasks)[0].ForkTasks, forkedTasksInner)
	}
	(*workflowTasks) = append((*workflowTasks), task.getJoinTask())
	return workflowTasks
}

func (task *fork) getJoinTask() http_model.WorkflowTask {
	if task.joinTask != nil {
		return (*task.joinTask.toWorkflowTask())[0]
	}
	join := Join(task.taskReferenceName + "_join")
	return (*join.toWorkflowTask())[0]
}

func (task *fork) Input(key string, value interface{}) *fork {
	task.task.Input(key, value)
	return task
}
