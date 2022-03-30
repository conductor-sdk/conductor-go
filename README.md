# Netflix Conductor Client SDK

To find out more about Conductor visit: [https://github.com/Netflix/conductor](https://github.com/Netflix/conductor)

`conductor-go` repository provides the client SDKs to build Task Workers in Go

## Quick Start

1. [Write TaskExecutionFunction](#Write-TaskExecutionFunction)
2. [Run workers](#Run-workers)
3. [Worker Configurations](#Worker-Configurations)



### Write Worker as a TaskExecutionFunction

```go
package example

import (
	"github.com/netflix/conductor/client/go/conductor_client/model"
	"github.com/netflix/conductor/client/go/conductor_client/model/enum/task_result_status"
	log "github.com/sirupsen/logrus"
)

func TaskExecuteFunctionExample(t *model.Task) (taskResult *model.TaskResult, err error) {
	log.Debug("Executing Task_Execution_Function_Example for", t.TaskType)
	taskResult = model.NewTaskResult(t)
	taskResult.OutputData = map[string]interface{}{
		"task": "task_1",
		"key2": "value2",
		"key3": 3,
		"key4": false,
	}
	taskResult.Logs = append(taskResult.Logs, model.LogMessage{Log: "Hello World"})
	taskResult.Status = task_result_status.COMPLETED
	err = nil
	return taskResult, err
}
```
### Run workers
Create main method that does the following:
1. Adds configurations such as metrics, authentication, thread count, Conductor server URL
2. Add your workers
3. Start the workers to poll for work

You can copy this code and put into `main.go` file

```go
package main

import (
	"os"

	"github.com/netflix/conductor/client/go/example"
	"github.com/netflix/conductor/client/go/metrics"
	"github.com/netflix/conductor/client/go/orkestrator"
	"github.com/netflix/conductor/client/go/settings"
	log "github.com/sirupsen/logrus"
)

// Example init function that shows how to configure logging
// Using json formatter and changing level to Debug
func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	//Stdout, change to a file for production use case
	log.SetOutput(os.Stdout)
	// Set to debug for demonstration.  Change to Info for production use cases.
	log.SetLevel(log.DebugLevel)
}

// Example main function to start workers
func main() {
	// MetricsSettings is optional,
	// could be nil instead and use default settings
	go metrics.ProvideMetrics(
		settings.NewMetricsSettings(
			"/metrics", // api endpoint
			2112,       // port
		),
	)
	// AuthenticationSettings and HttpSettings are optional,
	// could be nil instead and use default settings
	workerOrkestrator := orkestrator.NewWorkerOrkestrator(
		settings.NewAuthenticationSettings(
			"keyId",     // key id from your application
			"keySecret", // key secret from your application
		),
		settings.NewHttpSettings(
			"https://play.orkes.io/api", // conductor http server url
		),
	)
	workerOrkestrator.StartWorker(
		"go_task_example",                   // task definition name
		example.TaskExecuteFunctionExample1, // task execution function
		1,                                   // parallel go routines amount
		5000,                                // 5000ms
	)
	workerOrkestrator.StartWorker(
		"go_task_example",                   // task definition name
		example.TaskExecuteFunctionExample1, // task execution function
		1,                                   // parallel go routines amount
		100,                                 // 100ms
	)
	// Wait for all workers to finish, otherwise would terminate them
	workerOrkestrator.WaitWorkers()
}

```

### Running Conductor server locally in 2-minute
More details on how to run Conductor see https://netflix.github.io/conductor/server/ 

Use the script below to download and start the server locally.  The server runs in memory and no data saved upon exit.
```shell
export CONDUCTOR_VER=3.5.2
export REPO_URL=https://repo1.maven.org/maven2/com/netflix/conductor/conductor-server
curl $REPO_URL/$CONDUCTOR_VER/conductor-server-$CONDUCTOR_VER-boot.jar \
--output conductor-server-$CONDUCTOR_VER-boot.jar; java -jar conductor-server-$CONDUCTOR_VER-boot.jar 
```
### Execute workers
```shell
go ./main.go
```

### Create your first workflow
Now, let's create a new workflow and see your task worker code in execution!

Create a new Task Metadata for the worker you just created

```shell
curl -X 'POST' \
  'http://localhost:8080/api/metadata/taskdefs' \
  -H 'accept: */*' \
  -H 'Content-Type: application/json' \
  -d '[{
    "name": "go_task_example",
    "description": "Go task example",
    "retryCount": 3,
    "retryLogic": "FIXED",
    "retryDelaySeconds": 10,
    "timeoutSeconds": 300,
    "timeoutPolicy": "TIME_OUT_WF",
    "responseTimeoutSeconds": 180,
    "ownerEmail": "example@example.com"
}]'
```

Create a workflow that uses the task
```shell
curl -X 'POST' \
  'http://localhost:8080/api/metadata/workflow' \
  -H 'accept: */*' \
  -H 'Content-Type: application/json' \
  -d '{
    "name": "workflow_with_go_task_example",
    "description": "Workflow with Go Task example",
    "version": 1,
    "tasks": [
      {
        "name": "go_task_example",
        "taskReferenceName": "go_task_example_ref_1",
        "inputParameters": {},
        "type": "SIMPLE"
      }
    ],
    "inputParameters": [],
    "outputParameters": {
      "workerOutput": "${go_task_example_ref_1.output}"
    },
    "schemaVersion": 2,
    "restartable": true,
    "ownerEmail": "example@example.com",
    "timeoutPolicy": "ALERT_ONLY",
    "timeoutSeconds": 0
}'
```

Start a new workflow execution
```shell
curl -X 'POST' \
  'http://localhost:8080/api/workflow/workflow_with_go_task_example?priority=0' \
  -H 'accept: text/plain' \
  -H 'Content-Type: application/json' \
  -d '{}'
```


## Worker Configurations
Worker configuration is handled via `Configuraiton` object passed when initializing `TaskHandler`

### Server Configurations
* base_url : Conductor server address.  e.g. `http://localhost:8000` if running locally 
* debug: `true` for verbose logging `false` to display only the errors
* authentication_settings: see below
* metrics_settings: see below

### Metrics
Conductor uses [Prometheus](https://prometheus.io/) to collect metrics.

* directory: Directory where to store the metrics 
* file_name: File where the metrics are colleted. e.g. `metrics.log`
* update_interval: Time interval in seconds at which to collect metrics into the file

### Authentication
Use if your conductor server requires authentication
* key_id: Key
* key_secret: Secret for the Key 

## Unit Tests
### Simple validation

```shell
/conductor-python/src$ python3 -m unittest -v
```

### Run with code coverage

```shell
/conductor-python/src$ python3 -m coverage run --source=conductor/ -m unittest
```

Report:

```shell
/conductor-python/src$ python3 -m coverage report
```

Visual coverage results:
```shell
/conductor-python/src$ python3 -m coverage html
```

# Go client for Conductor
Go client for Conductor provides two sets of functions:

1. Workflow Management APIs (start, terminate, get workflow status etc.)
2. Worker execution framework

## Prerequisites
Go must be installed and GOPATH env variable set.

## Install

```shell
$ go get github.com/netflix/conductor/client/go
$ go get github.com/prometheus/client_golang/prometheus
$ go get github.com/prometheus/client_golang/prometheus/promauto
$ go get github.com/prometheus/client_golang/prometheus/promhttp
```
This will create a Go project under $GOPATH/src and download any dependencies.

## Run

```shell
go run $GOPATH/src/netflix-conductor/client/go/startclient/startclient.go
```

## Using Workflow Management API
Go struct ```ConductorHttpClient``` provides client API calls to the conductor server to start and manage workflows and tasks.

### Example
```go
package main

import (
    conductor "github.com/netflix/conductor/client/go"
)

func main() {
    conductorClient := conductor.NewConductorHttpClient("http://localhost:8080")
    
    // Example API that will print out workflow definition meta
    conductorClient.GetAllWorkflowDefs()
}

```

## Task Worker Execution
Task Worker execution APIs facilitates execution of a task worker using go.  The API provides necessary tools to poll for tasks at a specified interval and executing the go worker in a separate goroutine.

### Example
The following go code demonstrates workers for tasks "task_1" and "task_2".

```go
package task

import (
    "fmt"
)

// Implementation for "task_1"
func Task_1_Execution_Function(t *task.Task) (taskResult *task.TaskResult, err error) {
    log.Println("Executing Task_1_Execution_Function for", t.TaskType)

    //Do some logic
    taskResult = task.NewTaskResult(t)
    
    output := map[string]interface{}{"task":"task_1", "key2":"value2", "key3":3, "key4":false}
    taskResult.OutputData = output
    taskResult.Status = "COMPLETED"
    err = nil

    return taskResult, err
}

// Implementation for "task_2"
func Task_2_Execution_Function(t *task.Task) (taskResult *task.TaskResult, err error) {
    log.Println("Executing Task_2_Execution_Function for", t.TaskType)

    //Do some logic
    taskResult = task.NewTaskResult(t)
    
    output := map[string]interface{}{"task":"task_2", "key2":"value2", "key3":3, "key4":false}
    taskResult.OutputData = output
    taskResult.Status = "COMPLETED"
    err = nil

    return taskResult, err
}

```


Then main application to utilize these workers

```go
package main

import (
    "github.com/netflix/conductor/client/go"
    "github.com/netflix/conductor/client/go/task/sample"
)

func main() {
    c := conductor.NewConductorWorker("http://localhost:8080", 1, 10000)

    c.Start("task_1", "", sample.Task_1_Execution_Function, false)
    c.Start("task_2", "mydomain", sample.Task_2_Execution_Function, true)
}

```

Note: For the example listed above the example task implementations are in conductor/task/sample package.  Real task implementations can be placed in conductor/task directory or new subdirectory.

