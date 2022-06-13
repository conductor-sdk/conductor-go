package examples

import (
	def2 "github.com/conductor-sdk/conductor-go/workflow/def"
)

var HttpTask = def2.NewHttpTask(
	"go_task_of_http_type", // task name
	&def2.HttpInput{ // http input
		Uri: "https://catfact.ninja/fact",
	},
)

var SimpleTask = def2.NewSimpleTask(
	"go_task_of_simple_type",
	"go_task_of_simple_type",
)
