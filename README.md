# Conductor OSS Go SDK

SDK for developing Go applications that create, manage and execute workflows, and run workers.

[Conductor](https://www.conductor-oss.org/) is the leading open-source orchestration platform allowing developers to build highly scalable distributed applications.

To learn more about Conductor checkout our [developer's guide](https://docs.conductor-oss.org/devguide/concepts/index.html) and give it a ⭐ to make it famous!

[![GitHub stars](https://img.shields.io/github/stars/conductor-oss/conductor.svg?style=social&label=Star&maxAge=)](https://GitHub.com/conductor-oss/conductor/)


# Content
<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [Installation](#installation)
- [Hello World!](#hello-world)
  - [Step 1: Creating the workflow by code](#step-1-creating-the-workflow-by-code)
  - [Step 2: Creating the worker](#step-2-creating-the-worker)
  - [Step 3: Running the application](#step-3-running-the-application)
- [Further Reading](#further-reading)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Installation

1. Initialize your module. e.g.:

```shell
mkdir hello_world
cd hello_world
go mod init hello_world
```

2. Get the SDK:

```shell
go get github.com/conductor-sdk/conductor-go
```

## Hello World

In this repo you will find a basic "Hello World" under [examples/hello_world](examples/hello_world/). 

Let's analyze the app in 3 steps.


> [!note]
> You will need an up & running Conductor Server. 
>
> For details on how to run Conductor take a look at [our guide](https://conductor-oss.github.io/conductor/devguide/running/docker.html).
>
> The examples expect the server to be listening on http://localhost:8080.


### Step 1: Creating the workflow by code

The "greetings" workflow is going to be created by code and registered in Conductor. 

Check the `CreateWorkflow` function in [examples/hello_world/src/workflow.go](examples/hello_world/src/workflow.go).

```go
func CreateWorkflow(executor *executor.WorkflowExecutor) *workflow.ConductorWorkflow {
	wf := workflow.NewConductorWorkflow(executor).
		Name("greetings").
		Version(1).
		Description("Greetings workflow - Greets a user by their name").
		TimeoutPolicy(workflow.TimeOutWorkflow, 600)

	greet := workflow.NewSimpleTask("greet", "greet_ref").
		Input("person_to_be_greated", "${workflow.input.name}")

	wf.Add(greet)

	wf.OutputParameters(map[string]interface{}{
		"greetings": greet.OutputRef("hello"),
	})

	return wf
}
```

In the above code first we create a workflow by calling `workflow.NewConductorWorkflow(..)` and set its properties `Name`, `Version`, `Description` and `TimeoutPolicy`. 

Then we create a [Simple Task](https://orkes.io/content/reference-docs/worker-task) of type `"greet"` with reference name `"greet_ref"` and add it to the workflow. That task gets the workflow input `"name"` as an input with key `"person_to_be_greated"`.

> [!note]
>`"person_to_be_greated"` is too verbose! Why would you name it like that?
>
> It's just to make it clear that the workflow input is not passed automatically. 
>
> The worker will get the actual value of the workflow input because of this mapping  `Input("person_to_be_greated", "${workflow.input.name}")` in the workflow definition. 
>
>Expressions like `"${workflow.input.name}"` will be replaced by their value during execution.

Last but not least, the output of the workflow is set by calling `wf.OutputParameters(..)`. 

The value of `"greetings"` is going to be whatever `"hello"` is in the output of the executed `"greet"` task, e.g.: if the task output is:
```
{
	"hello" : "Hello, John"
}
```

The expected workflow output will be:
```
{
	"greetings": "Hello, John"
}
```

The Go code translates to this JSON defininition. You can view this in your Conductor server after registering the workflow.

```json
{
  "schemaVersion": 2,
  "name": "greetings",
  "description": "Greetings workflow - Greets a user by their name",
  "version": 1,
  "tasks": [
    {
      "name": "greet",
      "taskReferenceName": "greet_ref",
      "type": "SIMPLE",
      "inputParameters": {
        "name": "${workflow.input.name}"
      }
    }
  ],
  "outputParameters": {
    "Greetings": "${greet_ref.output.greetings}"
  },
  "timeoutPolicy": "TIME_OUT_WF",
  "timeoutSeconds": 600
}
```

> [!note]
> Workflows can also be registered using the API. Using the JSON you can make the following request:
> ```shell
> curl -X POST -H "Content-Type:application/json" \
> http://localhost:8080/api/metadata/workflow -d @greetings_workflow.json
> ```

In [Step 3](#step-3-running-the-application) you will see how to create an instance of `executor.WorkflowExecutor`.


### Step 2: Creating the worker

A worker is a function with a specific task to perform.

In this example the worker just uses the input `person_to_be_greated` to say hello, as you can see in [examples/hello_world/src/worker.go](examples/hello_world/src/worker.go).

```go
func Greet(task *model.Task) (interface{}, error) {
	return map[string]interface{}{
		"hello": "Hello, " + fmt.Sprintf("%v", task.InputData["person_to_be_greated"]),
	}, nil
}
```

To learn more about workers take a look at [Writing Workers with the Go SDK](docs/workers_sdk.md).

> [!note]
> A single workflow can have task workers written in different languages and deployed anywhere, making your workflow polyglot and distributed!

### Step 3: Running the application

The application is going to start the Greet worker (to execute tasks of type "greet") and it will register the workflow created in [step 1](#step-1-creating-the-workflow-by-code).

To begin with, let's take a look at the variable declaration in [examples/hello_world/main.go](examples/hello_world/main.go).

```go

var (
	apiClient = client.NewAPIClient(
		authSettings(),
		httpSettings(),
	)
	taskRunner       = worker.NewTaskRunnerWithApiClient(apiClient)
	workflowExecutor = executor.NewWorkflowExecutor(apiClient)
)

func authSettings() *settings.AuthenticationSettings {
	key := os.Getenv("KEY")
	secret := os.Getenv("SECRET")
	if key != "" && secret != "" {
		return settings.NewAuthenticationSettings(
			key,
			secret,
		)
	}

	return nil
}

func httpSettings() *settings.HttpSettings {
	url := os.Getenv("CONDUCTOR_SERVER_URL")
	if url == "" {
		fmt.Fprintf(os.Stderr, "Error: CONDUCTOR_SERVER_URL env variable is not set\n")
		os.Exit(1)
	}

	return settings.NewHttpSettings(url)
}
```

First we create an `APIClient` instance. This is a REST client. 

We need to pass on the proper settings to our client. For convenience to run the example you can set the following environment variables: `CONDUCTOR_SERVER_URL`, `KEY`, `SECRET`.


Now let's take a look at the `main` function:

```go
func main() {
	// Start the Greet Worker. This worker will process "greet" tasks.
	taskRunner.StartWorker("greet", hello_world.Greet, 1, time.Millisecond*100)

	// This is used to register the Workflow, it's a one-time process. You can comment from here
	wf := hello_world.CreateWorkflow(workflowExecutor)
	err := wf.Register(true)
	if err != nil {
		log.Error(err.Error())
		return
	}
	// Till Here after registering the workflow

	// Start the greetings workflow 
	id, err := workflowExecutor.StartWorkflow(
		&model.StartWorkflowRequest{
			Name:    "greetings",
			Version: 1,
			Input: map[string]string{
				"name": "Gopher",
			},
		},
	)

	if err != nil {
		log.Error(err.Error())
		return
	}

	log.Info("Started workflow with Id: ", id)

	// Get a channel to monitor the workflow execution -
	// Note: This is useful in case of short duration workflows that completes in few seconds.
	channel, _ := workflowExecutor.MonitorExecution(id)
	run := <-channel
	log.Info("Output of the workflow: ", run.Output)
}
```

The `taskRunner` uses the `apiClient` to poll for work and complete tasks. It also starts the worker and handles concurrency and polling intervals for us based on the configuration provided.

That simple line `taskRunner.StartWorker("greet", hello_world.Greet, 1, time.Millisecond*100)` is all that's needed to get our Greet worker up & running and processing tasks of type `"greet"`.

The `workflowExecutor` gives us an abstraction on top of the `apiClient` to manage workflows. It is used under the hood by `ConductorWorkflow` to register the workflow and it's also used to start and monitor the execution.

#### Running the example with a local Conductor OSS server:
```shell
export CONDUCTOR_SERVER_URL="http://localhost:8080/api"
cd examples
go run hello_world/main.go
```

#### Running the example in Orkes playground.
```shell
export CONDUCTOR_SERVER_URL="https://play.orkes.io/api"
export KEY="..."
export SECRET="..."
cd examples
go run hello_world/main.go
```

> [!note]
> Orkes Conductor requires authentication. [Get a key and secret from the server](https://orkes.io/content/how-to-videos/access-key-and-secret) to set those variables.

The above commands should give an output similar to
```shell
INFO[0000] Updated poll interval for task: greet, to: 100ms 
INFO[0000] Started 1 worker(s) for taskName greet, polling in interval of 100 ms 
INFO[0000] Started workflow with Id:14a9fcc5-3d74-11ef-83dc-acde48001122 
INFO[0000] Output of the workflow:map[Greetings:Hello, Gopher] 
```

# Further Reading

- [Writing Workers with the Go SDK](docs/workers_sdk.md)
- [Authoring Workflows with the Go SDK](docs/workflow_sdk.md)

// DUMMY change for testing purposes. DON'T MERGE.
