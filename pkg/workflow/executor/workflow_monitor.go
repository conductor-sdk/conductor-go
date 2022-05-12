package executor

import "github.com/conductor-sdk/conductor-go/pkg/workflow/def/tasks"

func hello() {
	simpleTask := tasks.NewSimpleTask()
	print(simpleTask)
	simpleTask.DoSomething()
}
