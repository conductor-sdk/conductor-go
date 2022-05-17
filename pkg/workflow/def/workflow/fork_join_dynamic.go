package workflow

import "github.com/conductor-sdk/conductor-go/pkg/http_model"

type DynamicForkTask struct {
	task        Task
	preForkTask *Task
	join        *JoinTask
}

func DynamicFork(taskRefName string, forkPrepareTask *Task) *DynamicForkTask {
	return &DynamicForkTask{
		task: Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          FORK_JOIN_DYNAMIC,
			optional:          false,
			inputParameters:   nil,
		},
		preForkTask: forkPrepareTask,
		// join:        NewJoin(taskRefName + "_join"),
	}
}

func NewDynamicForkWithJoin(taskRefName string, forkPrepareTask *Task, join *JoinTask) *DynamicForkTask {
	return &DynamicForkTask{
		task: Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          FORK_JOIN_DYNAMIC,
			optional:          false,
			inputParameters:   nil,
		},
		preForkTask: forkPrepareTask,
		join:        join,
	}
}

func (task *DynamicForkTask) toWorkflowTask() []http_model.WorkflowTask {
	forkWorkflowTask := task.task.toWorkflowTask()[0]
	forkWorkflowTask.DynamicForkTasksParam = "forkedTasks"
	forkWorkflowTask.DynamicForkTasksInputParamName = "forkedTasksInputs"
	forkWorkflowTask.InputParameters["forkedTasks"] = task.preForkTask.OutputRef("forkedTasks")
	forkWorkflowTask.InputParameters["forkedTasksInputs"] = (task.preForkTask).OutputRef("forkedTasksInputs")
	tasks := (task.preForkTask).toWorkflowTask()
	// tasks = append(tasks, forkWorkflowTask, task.getJoinTask())
	return tasks
}
