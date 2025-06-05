package src

import (
	"fmt"
	"github.com/conductor-sdk/conductor-go/sdk/model"
)

func SimpleTask(task *model.Task) (interface{}, error) {
	return map[string]interface{}{
		"greetings": "Hello, " + fmt.Sprintf("%v", task.InputData["name"]),
	}, nil
}
