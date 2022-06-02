# Netflix Conductor Client SDK

`conductor-go` repository provides the client SDKs to build Task Workers in Go

## Quick Start

1. [Setup conductor-go package](#Setup-conductor-go-package)
2. [Run workers](#Run-workers)
3. [Configuration](#Configuration)

### Setup conductor go package

Create a folder to start your project:
```shell
$ mkdir conductor_go_example/
$ cd conductor_go_example/
```

Create a new empty module named `example` ([reference](https://go.dev/ref/mod#go-mod-init)):

```shell
$ go mod init example
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
	"fmt"
	"os"
	"time"

	"github.com/conductor-sdk/conductor-go/examples"
	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/settings"
	"github.com/conductor-sdk/conductor-go/pkg/worker"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/executor"
	log "github.com/sirupsen/logrus"
)

var (
	apiClient = conductor_http_client.NewAPIClient(
		settings.NewAuthenticationSettings(
			"", // keyId
			"", // keySecret
		),
		settings.NewHttpSettings(
			"https://play.orkes.io/api", // baseUrl
		),
	)
	taskRunner       = worker.NewTaskRunnerWithApiClient(apiClient)
	workflowExecutor = executor.NewWorkflowExecutor(apiClient)
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func runHttpWorkflowExample() error {
	httpTaskWorkflow := examples.NewHttpTaskConductorWorkflow(workflowExecutor)
	_, err := httpTaskWorkflow.Register()
	if err != nil {
		return err
	}
	log.Debug("Registered workflow with http task example")
	workflowId, workflowExecutionChannel, err := httpTaskWorkflow.Start(nil)
	if err != nil {
		return err
	}
	log.Info("Started workflow http task example, workflowId: ", workflowId)
	workflow, err := executor.WaitForWorkflowCompletionUntilTimeout(
		workflowExecutionChannel,
		5*time.Second,
	)
	if err != nil {
		return err
	}
	log.Warning("Workflow with http task example is finished")
	if !examples.IsWorkflowCompleted(workflow) {
		return fmt.Errorf("failed to get completed workflow")
	}
	return nil
}

func runSimpleWorkflowExample() error {
	simpleTaskWorkflow := examples.NewSimpleTaskConductorWorkflow(workflowExecutor)
	_, err := simpleTaskWorkflow.Register()
	if err != nil {
		return err
	}
	log.Debug("Registered workflow with simple task example")
	workflowId, workflowExecutionChannel, err := simpleTaskWorkflow.Start(nil)
	if err != nil {
		return err
	}
	log.Info("Started workflow with simple task example, workflowId: ", workflowId)
	err = taskRunner.StartWorker(
		examples.SimpleTask.ReferenceName(),
		examples.SimpleWorker,
		10,
		500*time.Millisecond,
	)
	if err != nil {
		return err
	}
	workflow, err := executor.WaitForWorkflowCompletionUntilTimeout(
		workflowExecutionChannel,
		5*time.Second,
	)
	if err != nil {
		return err
	}
	taskRunner.RemoveWorker(
		examples.SimpleTask.ReferenceName(),
		10,
	)
	if !examples.IsWorkflowCompleted(workflow) {
		return fmt.Errorf("failed to get completed workflow")
	}
	return nil
}

func main() {
	runHttpWorkflowExample()
	runSimpleWorkflowExample()
}
```

Run your `main.go` file, example using [Orkes Playground Conductor Server](https://play.orkes.io/):
```shell
$ go run main.go
{"level":"debug","msg":"Refreshing authentication token","time":"2022-06-02T00:27:15-03:00"}
{"level":"debug","msg":"Registered workflow with http task example","time":"2022-06-02T00:27:16-03:00"}
{"level":"debug","msg":"Started workflow, workflowId: e6e54301-e223-11ec-a6d8-32508b865be6, name: GO_WORKFLOW_WITH_HTTP_TASK, version: 1, input: \u003cnil\u003e","time":"2022-06-02T00:27:16-03:00"}
{"level":"debug","msg":"Added workflow execution channel, workflowId: e6e54301-e223-11ec-a6d8-32508b865be6","time":"2022-06-02T00:27:16-03:00"}
{"level":"info","msg":"Started workflow http task example, workflowId: e6e54301-e223-11ec-a6d8-32508b865be6","time":"2022-06-02T00:27:16-03:00"}
{"level":"debug","msg":"Notifying finished workflowId: e6e54301-e223-11ec-a6d8-32508b865be6","time":"2022-06-02T00:27:17-03:00"}
{"level":"debug","msg":"Sent finished workflow through channel","time":"2022-06-02T00:27:17-03:00"}
{"level":"debug","msg":"Closed client workflow execution channel","time":"2022-06-02T00:27:17-03:00"}
{"level":"debug","msg":"Deleted workflow execution channel","time":"2022-06-02T00:27:17-03:00"}
{"level":"warning","msg":"Workflow with http task example is finished","time":"2022-06-02T00:27:17-03:00"}
{"level":"debug","msg":"Registered workflow with simple task example","time":"2022-06-02T00:27:17-03:00"}
{"level":"debug","msg":"Started workflow, workflowId: e79c0fd3-e223-11ec-9c27-368389356974, name: GO_WORKFLOW_WITH_SIMPLE_TASK, version: 1, input: \u003cnil\u003e","time":"2022-06-02T00:27:17-03:00"}
{"level":"debug","msg":"Added workflow execution channel, workflowId: e79c0fd3-e223-11ec-9c27-368389356974","time":"2022-06-02T00:27:17-03:00"}
{"level":"info","msg":"Started workflow with simple task example, workflowId: e79c0fd3-e223-11ec-9c27-368389356974","time":"2022-06-02T00:27:17-03:00"}
{"level":"debug","msg":"Increased max allowed workers of task: GO_TASK_OF_SIMPLE_TYPE, by: 10","time":"2022-06-02T00:27:17-03:00"}
{"level":"info","msg":"Started 10 worker(s) for taskType GO_TASK_OF_SIMPLE_TYPE, polling in interval of 500 ms","time":"2022-06-02T00:27:17-03:00"}
{"level":"debug","msg":"Polling for task: GO_TASK_OF_SIMPLE_TYPE, in batches of size: 10","time":"2022-06-02T00:27:17-03:00"}
{"level":"debug","msg":"Polled 1 tasks for taskType: GO_TASK_OF_SIMPLE_TYPE","time":"2022-06-02T00:27:18-03:00"}
{"level":"debug","msg":"Polling for task: GO_TASK_OF_SIMPLE_TYPE, in batches of size: 9","time":"2022-06-02T00:27:18-03:00"}
{"level":"debug","msg":"Updating task of type: GO_TASK_OF_SIMPLE_TYPE, taskId: e79cd324-e223-11ec-9c27-368389356974, workflowId: e79c0fd3-e223-11ec-9c27-368389356974","time":"2022-06-02T00:27:18-03:00"}
{"level":"debug","msg":"Updated task of type: GO_TASK_OF_SIMPLE_TYPE, taskId: e79cd324-e223-11ec-9c27-368389356974, workflowId: e79c0fd3-e223-11ec-9c27-368389356974","time":"2022-06-02T00:27:18-03:00"}
{"level":"debug","msg":"Notifying finished workflowId: e79c0fd3-e223-11ec-9c27-368389356974","time":"2022-06-02T00:27:18-03:00"}
{"level":"debug","msg":"Sent finished workflow through channel","time":"2022-06-02T00:27:18-03:00"}
{"level":"debug","msg":"Closed client workflow execution channel","time":"2022-06-02T00:27:18-03:00"}
{"level":"debug","msg":"Deleted workflow execution channel","time":"2022-06-02T00:27:18-03:00"}
{"level":"debug","msg":"Decreased workers for task: GO_TASK_OF_SIMPLE_TYPE, by: 10","time":"2022-06-02T00:27:18-03:00"}
```

Explanation:
// TODO

### Configuration

#### Authentication settings (optional)
Use if your conductor server requires authentication
* keyId: Key
* keySecret: Secret for the Key

```go
settings.NewAuthenticationSettings(
  "", // keyId
  "", // keySecret
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
