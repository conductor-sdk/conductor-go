package def

import (
	"github.com/conductor-sdk/conductor-go/model"
)

type ForkTask struct {
	Task
	forkedTasks [][]TaskInterface
}

//NewForkTask creates a new fork task that executes the given tasks in parallel
/**
 * execute task specified in the forkedTasks parameter in parallel.
 *
 * <p>forkedTask is a two-dimensional list that executes the outermost list in parallel and list
 * within that is executed sequentially.
 *
 * <p>e.g. [[task1, task2],[task3, task4],[task5]] are executed as:
 *
 * <pre>
 *                    ---------------
 *                    |     fork    |
 *                    ---------------
 *                    |       |     |
 *                    |       |     |
 *                  task1  task3  task5
 *                  task2  task4    |
 *                    |      |      |
 *                 ---------------------
 *                 |       join        |
 *                 ---------------------
 * </pre>
 *
 *
 */
func NewForkTask(taskRefName string, forkedTask ...[]TaskInterface) *ForkTask {
	return &ForkTask{
		Task: Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          FORK_JOIN,
			optional:          false,
			inputParameters:   map[string]interface{}{},
		},
		forkedTasks: forkedTask,
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

// Input to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *ForkTask) Input(key string, value interface{}) *ForkTask {
	task.Task.Input(key, value)
	return task
}

// InputMap to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *ForkTask) InputMap(inputMap map[string]interface{}) *ForkTask {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}

// Optional if set to true, the task will not fail the workflow if one of the loop task fails
func (task *ForkTask) Optional(optional bool) *ForkTask {
	task.Task.Optional(optional)
	return task
}

// Description of the task
func (task *ForkTask) Description(description string) *ForkTask {
	task.Task.Description(description)
	return task
}
