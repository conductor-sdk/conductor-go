package examples

import (
	"github.com/conductor-sdk/conductor-go/sdk/workflow/def"
)

var HttpTask = def.NewHttpTask(
	"go_task_of_http_type", // task name
	&def.HttpInput{ // http input
		Uri: "https://catfact.ninja/fact",
	},
)

var SimpleTask = def.NewSimpleTask(
	"go_task_of_simple_type",
	"go_task_of_simple_type",
)
