package workflow

import "github.com/conductor-sdk/conductor-go/pkg/http_model"

func DynamicFork(taskRefName string, forkPrepareTask Task) *dynamicFork {
	return &dynamicFork{
		task: task{
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

type dynamicFork struct {
	task
	preForkTask Task
}

func (task *dynamicFork) toWorkflowTask() []http_model.WorkflowTask {

	forkWorkflowTask := task.task.toWorkflowTask()[0]
	forkWorkflowTask.DynamicForkTasksParam = "forkedTasks"
	forkWorkflowTask.DynamicForkTasksInputParamName = "forkedTasksInputs"
	forkWorkflowTask.InputParameters["forkedTasks"] = (task.preForkTask).OutputRef("forkedTasks")
	forkWorkflowTask.InputParameters["forkedTasksInputs"] = (task.preForkTask).OutputRef("forkedTasksInputs")
	tasks := (task.preForkTask).toWorkflowTask()
	tasks = append(tasks, forkWorkflowTask, task.getJoinTask())

	return tasks
}

func (task *dynamicFork) getJoinTask() http_model.WorkflowTask {
	join := Join(task.taskReferenceName + "_join")
	return (join.toWorkflowTask())[0]
}
