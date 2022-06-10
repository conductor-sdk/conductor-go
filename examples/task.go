package examples

import "github.com/conductor-sdk/conductor-go/pkg/workflow/def/workflow"

var HttpTask = workflow.NewHttpTask(
	"go_task_of_http_type", // task name
	&workflow.HttpInput{ // http input
		Uri: "https://catfact.ninja/fact",
	},
)

var SimpleTask = workflow.NewSimpleTask(
	"go_task_of_simple_type",
	"go_task_of_simple_type",
)
