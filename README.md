# Netflix Conductor Client SDK

`conductor-go` repository provides the client SDKs to build Task Workers in Go

## Quick Start

1. [Setup conductor-go package](#Setup-conductor-go-package)
2. [Run workers](#Run-workers)
3. [Configuration](#Configuration)

### Setup conductor go package

Create a folder to start your project:
```shell
mkdir conductor_go_example/
cd conductor_go_example/
```

Create a new empty module named `example` ([reference](https://go.dev/ref/mod#go-mod-init)):

```shell
$ go init example
go: creating new go.mod: module example
```

Install `conductor-go` package:
```shell
$ go get github.com/conductor-sdk/conductor-go
go: downloading github.com/conductor-sdk/conductor-go v1.1.2
go get: added github.com/conductor-sdk/conductor-go v1.1.2
```

(*Optional*) Using a specific version ([reference](https://go.dev/ref/mod#go-get)). Example with `code_review` branch:
```shell
$ go get github.com/conductor-sdk/conductor-go@code_review
go: downloading github.com/conductor-sdk/conductor-go v1.1.3-0.20220601175614-e039dcf37361
go get: upgraded github.com/conductor-sdk/conductor-go v1.1.2 => v1.1.3-0.20220601175614-e039dcf37361
```

### Run workers

Create a `main.go` file and paste this code there:
```go
package main

import (
	"context"
	"os"

	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/conductor-sdk/conductor-go/pkg/model"
	"github.com/conductor-sdk/conductor-go/pkg/model/enum/task_result_status"
	"github.com/conductor-sdk/conductor-go/pkg/settings"
	"github.com/conductor-sdk/conductor-go/pkg/worker"
	log "github.com/sirupsen/logrus"
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
	return taskResult, nil
}

func main() {
	taskRunner := worker.NewTaskRunner(
		settings.NewAuthenticationSettings(
			__KEY__,    // KeyId
			__SECRET__, // KeySecret
		),
		settings.NewHttpSettings(
			"https://play.orkes.io/api", // BaseUrl
		),
	)
	taskRunner.StartWorker(
		"go_task_example", // TaskType
		Worker,            // TaskExecutionFunction
		2,                 // WorkerAmount
		10,                // PollingInterval
	)
	taskRunner.WaitWorkers()
}
```

Run your `main.go` file
```shell
$ go run main.go
// TODO
```

Explanation:


### Configuration

#### Authentication settings (optional)
Use if your conductor server requires authentication
* keyId: Key
* keySecret: Secret for the Key

```go
authenticationSettings := settings.NewAuthenticationSettings(
  "keyId",
  "keySecret",
)
```

### Worker Settings

You can create a new worker by calling `taskRunner.StartWorker` with:
* taskType : Task definition name (e.g `"go_task_example"`)
* executeFunction : Task Execution Function (e.g. `example.TaskExecuteFunctionExample1` from `example` folder)
* threadCount : Amount of Go routines to be executed in parallel for new worker (e.g. `1`, single thread)
* pollIntervalInMillis : Amount of ms to wait between polling for task

```go
taskRunner.StartWorker(
	"go_task_example",              // task definition name
	Worker, // task execution function
	1,                              // thread count
	1000,                           // polling interval in milli-seconds
)
```
### Start a workflow using APIs
```go

apiClient := conductor_http_client.NewAPIClient(
    settings.NewAuthenticationSettings(
        KEY,
        SECRET,
    ),
    settings.NewHttpSettings(
        "https://play.orkes.io",
    ),
)

workflowClient := *&conductor_http_client.WorkflowResourceApiService{
    APIClient: apiClient,
}
workflowId, _, _ := workflowClient.StartWorkflow(
    context.Background(),
    map[string]interface{}{},
    "PopulationMinMax",
    &conductor_http_client.WorkflowResourceApiStartWorkflowOpts{},
)
log.Info("Workflow Id is ", workflowId)
	
```
