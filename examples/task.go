package examples

import (
	"github.com/conductor-sdk/conductor-go/sdk/workflow/definition"
)

var HttpTask = definition.NewHttpTask(
	"go_task_of_http_type", // task name
	&definition.HttpInput{ // http input
		Uri: "https://catfact.ninja/fact",
	},
)

var SimpleTask = definition.NewSimpleTask(
	"go_task_of_simple_type",
	"go_task_of_simple_type",
)
