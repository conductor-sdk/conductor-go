package workflow

import "github.com/conductor-sdk/conductor-go/pkg/http_model"

type DynamicForkTask struct {
	Task
	preForkTask TaskInterface
	join        JoinTask
}

func NewDynamicForkTask(taskRefName string, forkPrepareTask TaskInterface) *DynamicForkTask {
	return &DynamicForkTask{
		Task: Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          FORK_JOIN_DYNAMIC,
			optional:          false,
			inputParameters:   map[string]interface{}{},
		},
		preForkTask: forkPrepareTask,
	}
}

func NewDynamicForkWithJoinTask(taskRefName string, forkPrepareTask TaskInterface, join JoinTask) *DynamicForkTask {
	return &DynamicForkTask{
		Task: Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          FORK_JOIN_DYNAMIC,
			optional:          false,
			inputParameters:   map[string]interface{}{},
		},
		preForkTask: forkPrepareTask,
		join:        join,
	}
}

func (task *DynamicForkTask) toWorkflowTask() []http_model.WorkflowTask {
	forkWorkflowTask := task.Task.toWorkflowTask()[0]
	forkWorkflowTask.DynamicForkTasksParam = "forkedTasks"
	forkWorkflowTask.DynamicForkTasksInputParamName = "forkedTasksInputs"
	forkWorkflowTask.InputParameters["forkedTasks"] = (task.preForkTask).OutputRef("forkedTasks")
	forkWorkflowTask.InputParameters["forkedTasksInputs"] = (task.preForkTask).OutputRef("forkedTasksInputs")
	tasks := (task.preForkTask).toWorkflowTask()
	tasks = append(tasks, forkWorkflowTask, task.getJoinTask())

	return tasks
}

func (task *DynamicForkTask) getJoinTask() http_model.WorkflowTask {
	join := NewJoinTask(task.taskReferenceName + "_join")
	return (join.toWorkflowTask())[0]
}
