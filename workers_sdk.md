# Writing Workers

Workers are the implementation of Tasks in the workflow.

Let's create a simple worker implementation as `main.go`

```go
package main

import (
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/conductor-sdk/conductor-go/pkg/model"
	"github.com/conductor-sdk/conductor-go/pkg/model/enum/task_result_status"
	"github.com/conductor-sdk/conductor-go/pkg/settings"
	"github.com/conductor-sdk/conductor-go/pkg/worker"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

//init set the logging for the workers
func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

//TaskOutput struct that represents the output of the task execution
type TaskOutput struct {
	Keys    []string
	Message string
	Value   float64
}

//SimpleWorker function accepts Task as input and returns TaskOutput as result
//If there is a failure, error can be returned and the task will be marked as FAILED
func SimpleWorker(t *model.Task) (interface{}, error) {
	taskResult := &TaskOutput{
		Keys:    []string{"Key1", "Key2"},
		Message: "Hello World",
		Value:   rand.ExpFloat64(),
	}
	return taskResult, nil
}

func main() {
	taskRunner := worker.NewTaskRunner(
		settings.NewAuthenticationSettings(
			__KEY__,
			__SECRET__,
		),
		settings.NewHttpSettings(
			"https://play.orkes.io",
		),
	)
	//Start worker execution with a batch size of 1 and sleep interval of 1 second if there are no tasks to be executed
	taskRunner.StartWorker("simple_worker", SimpleWorker, 1, time.Second * 1)
    
	//Block
	taskRunner.WaitWorkers()
}

```
**Note:**
Replace `KEY` and `SECRET` by obtaining a new key and secret from Orkes Playground as described [Generating Access Keys for Programmatic Access](https://orkes.io/content/docs/getting-started/concepts/access-control#access-keys)

Also - replace `simple_worker` with the name of your task.

### Run workers
Start the workers by running `go run`
```shell
go run main.go
```

### Long-running tasks
For the long-running tasks you might want to spawn another process/routine and update the status of the task at a later point and complete the
execution function without actually marking the task as `COMPLETED`.  Use `TaskResult` struct that allows you to specify more fined grained control.

Here is an example of a task execution function that returns with `IN_PROGERSS` status asking server to push the task again in 60 seconds.
```go
func LongRunningTaskWorker(t *model.Task) (interface{}, error) {
	taskResult := model.NewTaskResult(t)
	taskResult.OutputData = map[string]interface{}{}
    
	//Keep the status as IN_PROGRESS
	taskResult.Status = task_result_status.IN_PROGRESS
	//Time after which the task should be sent back to worker
	taskResult.CallbackAfterSeconds = 60
	return taskResult, nil
}
```

## Task Management APIs

### Get Task Details
```go
task, err := executor.GetTask(taskId)
```

### Updating the Task result outside the worker implementation
#### Update task by id
```go
output :=  &TaskOutput{
Keys:    []string{"Key1", "Key2"},
Message: "Hello World",
Value:   rand.ExpFloat64(),
}
executor.UpdateTask(taskId, workflowInstanceId, task_result_status.COMPLETED, ouptut)
```

#### Update task by Reference Name
```go
output :=  &TaskOutput{
Keys:    []string{"Key1", "Key2"},
Message: "Hello World",
Value:   rand.ExpFloat64(),
}
executor.UpdateTaskByRefName("task_ref_name", workflowInstanceId, task_result_status.COMPLETED, ouptut)
```