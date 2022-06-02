package examples

import "github.com/conductor-sdk/conductor-go/pkg/workflow/def/workflow"

var HttpTask = workflow.NewHttpTask(
	"GO_TASK_OF_HTTP_TYPE", // task name
	&workflow.HttpInput{ // http input
		Uri: "https://catfact.ninja/fact",
	},
)

var SimpleTask = workflow.NewSimpleTask(
	"GO_TASK_OF_SIMPLE_TYPE", // task name
)
