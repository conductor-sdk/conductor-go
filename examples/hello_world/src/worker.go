// Hello World Application Using Conductor
package hello_world

import (
	"fmt"

	"github.com/conductor-sdk/conductor-go/sdk/model"
)

func Greet(task *model.Task) (interface{}, error) {
	return map[string]interface{}{
		"greetings": "Hello, " + fmt.Sprintf("%v", task.InputData["name"]),
	}, nil
}
