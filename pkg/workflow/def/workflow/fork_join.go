package workflow

// import "github.com/conductor-sdk/conductor-go/pkg/http_model"

// func Fork(taskRefName string, tasks ...[]Task) *fork {
// 	return &fork{
// 		task: task{
// 			name:              taskRefName,
// 			taskReferenceName: taskRefName,
// 			description:       "",
// 			taskType:          FORK_JOIN,
// 			optional:          false,
// 			inputParameters:   map[string]interface{}{},
// 		},
// 		forkedTasks: tasks,
// 	}
// }

// type fork struct {
// 	task
// 	forkedTasks [][]Task
// 	joinTask    Task
// }

// func (task *fork) toWorkflowTask() []http_model.WorkflowTask {
// 	workflowTask := task.task.toWorkflowTask()[0]
// 	workflowTask.ForkTasks = task.forkedTasks

// 	for _, forkedTask := range task.forkedTasks {
// 		for _, innerForkedTask := range forkedTask {
// 			workflowTask.ForkTasks  = append(
// 				workflowTasks.forkedTask,
// 				innerForkedTask,
// 			)
// 		}
// 	}
// 	workflowTask := task.task.toWorkflowTask()[0]
// 	for _, forkedTask := range task.forkedTasks {
// 		var forkedTasksInner []http_model.WorkflowTask
// 		for _, t := range forkedTask {
// 			forkedTasksInner = append(
// 				forkedTasksInner,
// 				t.toWorkflowTask()...,
// 			)
// 		}
// 		workflowTask.ForkTasks = append(
// 			workflowTask.ForkTasks,
// 			forkedTasksInner,
// 		)
// 	}
// 	if task.joinTask != nil {
// 		workflowTasks = append(
// 			workflowTasks,
// 			task.joinTask.toWorkflowTask()...,
// 		)
// 	}
// 	return workflowTasks
// }

// func (task *fork) Input(key string, value interface{}) *fork {
// 	task.task.Input(key, value)
// 	return task
// }
