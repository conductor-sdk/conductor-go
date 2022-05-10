package main

import (
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/conductor-sdk/conductor-go/pkg/model"
	"github.com/conductor-sdk/conductor-go/pkg/model/enum/task_result_status"
	"github.com/conductor-sdk/conductor-go/pkg/settings"
	"github.com/conductor-sdk/conductor-go/pkg/worker"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func Worker(t *http_model.Task) (taskResult *http_model.TaskResult, err error) {
	taskResult = model.GetTaskResultFromTask(t)
	taskResult.OutputData = map[string]interface{}{
		"task": "task_1",
		"key3": 3,
		"key4": false,
	}
	taskResult.Status = task_result_status.COMPLETED
	err = nil
	return taskResult, err
}

func main() {
	taskRunner := worker.NewTaskRunner(
		settings.NewAuthenticationSettings(
			"",
			"",
		),
		settings.NewHttpSettings(
			"https://play.orkes.io",
		),
	)

	taskRunner.StartWorker(
		"go_task_example",
		Worker,
		2,
		10,
	)

	taskRunner.WaitWorkers()
}
