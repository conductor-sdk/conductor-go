# Conductor OSS Go SDK

Go SDK for working with https://github.com/conductor-oss/conductor.

[Conductor](https://www.conductor-oss.org/) is the leading open-source orchestration platform allowing developers to build highly scalable distributed applications.

Check out the [official documentation for Conductor](https://orkes.io/content).

## ⭐ Conductor OSS

Show support for the Conductor OSS.  Please help spread the awareness by starring Conductor repo.

[![GitHub stars](https://img.shields.io/github/stars/conductor-oss/conductor.svg?style=social&label=Star&maxAge=)](https://GitHub.com/conductor-oss/conductor/)

## Content
<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [Install Conductor Go SDK](#install-conductor-go-sdk)
  - [Get Conductor Go SDK](#get-conductor-go-sdk)
- [Hello World Application Using Conductor](#hello-world-application-using-conductor)
  - [Step 1: Create Workflow](#step-1-create-workflow)
    - [Creating Workflows by Code](#creating-workflows-by-code)
    - [(Alternatively) Creating Workflows in JSON](#alternatively-creating-workflows-in-json)
  - [Step 2: Write Task Worker](#step-2-write-task-worker)
  - [Step 3: Write _Hello World_ Application](#step-3-write-_hello-world_-application)
- [Running Workflows on Conductor Standalone (Installed Locally)](#running-workflows-on-conductor-standalone-installed-locally)
  - [Setup Environment Variable](#setup-environment-variable)
  - [Start Conductor Server](#start-conductor-server)
  - [Execute Hello World Application](#execute-hello-world-application)
- [Running Workflows on Orkes Conductor](#running-workflows-on-orkes-conductor)
- [Learn More about Conductor Go SDK](#learn-more-about-conductor-go-sdk)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->



## Install Conductor Go SDK

Before installing Conductor Go SDK, it is a good practice to set up a dedicated folder for it.

```shell
mkdir quickstart/
cd quickstart/
go mod init quickstart
```

### Get Conductor Go SDK

The SDK requires Go. To install the SDK, use the following command
```shell
go get github.com/conductor-sdk/conductor-go
```
## Hello World Application Using Conductor

In this section, we will create a simple "Hello World" application that executes a "greetings" workflow managed by Conductor.

### Step 1: Create Workflow

#### Creating Workflows by Code

Create workflow/workflow.go with the following:

```go
package workflow

import (
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
)

// Name struct that represents the input to the workflow
type NameAndCity struct {
	Name string
}

func GetTaskDefinitions() []model.TaskDef {
	taskDefs := []model.TaskDef{
		{Name: "greet", TimeoutSeconds: 60},
	}
	return taskDefs
}

// Create a workflow and register it with the server
func CreateWorkflow(executor *executor.WorkflowExecutor) *workflow.ConductorWorkflow {

	wf := workflow.NewConductorWorkflow(executor).
		Name("hello5").
		Version(1).
		Description("Greetings workflow - Greets a user by their name").
		TimeoutPolicy(workflow.TimeOutWorkflow, 600)

	//Greet Task
	greet := workflow.NewSimpleTask("greet", "greet_ref").
		Input("name", "${workflow.input.Name}")

	//Add tasks to workflow
	wf.Add(greet)
	//Add the output of the workflow from the task
	wf.OutputParameters(map[string]interface{}{
		"Greetings": greet.OutputRef("greetings"),
	})
	return wf
}
```

#### (Alternatively) Creating Workflows in JSON

Create `greetings_workflow.json` with the following:

```json
{
  "name": "greetings",
  "description": "Sample greetings workflow",
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
  "timeoutPolicy": "TIME_OUT_WF",
  "timeoutSeconds": 60
}
```

Workflows must be registered to the Conductor server. Use the API to register the greetings workflow from the JSON file above:
```shell
curl -X POST -H "Content-Type:application/json" \
http://localhost:8080/api/metadata/workflow -d @greetings_workflow.json
```
> [!note]
> To use the Conductor API, the Conductor server must be up and running (see [Running over Conductor standalone (installed locally)](#running-over-conductor-standalone-installed-locally)).

### Step 2: Write Task Worker

Using Go, a worker represents a function with a specific task to perform. Create greet/greet.go

> [!note]
> A single workflow can have task workers written in different languages and deployed anywhere, making your workflow polyglot and distributed!

```go
package greet

import (
	"fmt"
	"github.com/conductor-sdk/conductor-go/sdk/model"
)

// Task worker
func Greet(task *model.Task) (interface{}, error) {
	return map[string]interface{}{
		"greetings": "Hello, " + fmt.Sprintf("%v", task.InputData["name"]),
	}, nil
}
```

Now, we are ready to write our main application, which will execute our workflow.

### Step 3: Write _Hello World_ Application

Let's add main.go with a `main` method:

```go
package main

import (
	"fmt"
	"os"
	"quickstart/greet"
	"quickstart/workflow"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/settings"

	"github.com/conductor-sdk/conductor-go/sdk/worker"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
)

var (
	apiClient = client.NewAPIClient(
		settings.NewAuthenticationSettings(
			os.Getenv("KEY"),
			os.Getenv("SECRET"),
		),
		settings.NewHttpSettings(
			os.Getenv("CONDUCTOR_SERVER_URL"),
		))
	taskRunner       = worker.NewTaskRunnerWithApiClient(apiClient)
	workflowExecutor = executor.NewWorkflowExecutor(apiClient)
	metadataClient   = client.MetadataResourceApiService{APIClient: apiClient}
)

func StartWorkers() {
	taskRunner.StartWorker("greet", greet.Greet, 1, time.Millisecond*100)
}

func main() {

	//Start the workers
	StartWorkers()
    	/* This is used to register the Workflow, it's a one-time process. Comment from here */
	wf := workflow.CreateWorkflow(workflowExecutor)
	err := wf.Register(false)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	/* Till Here after registering the workflow*/
	id, err := wf.StartWorkflowWithInput(&workflow.NameAndCity{
		Name: "Orkes",
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
    fmt.Println("Started workflow with Id: ", id)

	/*Get a channel to monitor the workflow execution -
	  Note: This is useful in case of short duration workflows that completes in few seconds.*/
	channel, _ := workflowExecutor.MonitorExecution(id)
	run := <-channel
	fmt.Println("Output of the workflow, ", run.Status)

}
```

## Running Workflows on Conductor Standalone (Installed Locally)

### Setup Environment Variable

Set the following environment variable to point the SDK to the Conductor Server API endpoint:

```shell
export CONDUCTOR_SERVER_URL=http://localhost:8080/api
```
### Start Conductor Server

To start the Conductor server in a standalone mode from a Docker image, type the command below:

```shell
docker run --init -p 8080:8080 -p 5000:5000 conductoross/conductor-standalone:3.15.0
```
To ensure the server has started successfully, open Conductor UI on http://localhost:5000.

### Execute Hello World Application

To run the application, type the following command:

```shell
go run main.go
```
Now, the workflow is executed, and its execution status can be viewed from Conductor UI (http://localhost:5000).

Navigate to the **Executions** tab to view the workflow execution.

## Running Workflows on Orkes Conductor

For running the workflow in Orkes Conductor,

- Update the Conductor server URL to your cluster name.

```shell
export CONDUCTOR_SERVER_URL=https://[cluster-name].orkesconductor.io/api
```

- If you want to run the workflow on the Orkes Conductor Playground, set the Conductor Server variable as follows:

```shell
export CONDUCTOR_SERVER_URL=https://play.orkes.io/api
```

- Orkes Conductor requires authentication. [Obtain the key and secret from the Conductor server](https://orkes.io/content/how-to-videos/access-key-and-secret) and set the following environment variables.

```shell
export CONDUCTOR_AUTH_KEY=your_key
export CONDUCTOR_AUTH_SECRET=your_key_secret
```

Run the application and view the execution status from Conductor's UI Console.

> [!NOTE]
> That's it - you just created and executed your first distributed Go app!

## Learn More about Conductor Go SDK

There are three main ways you can use Conductor when building durable, resilient, distributed applications.

1. Write service workers that implement business logic to accomplish a specific goal - such as initiating payment transfer, getting user information from the database, etc.
2. Create Conductor workflows that implement application state - A typical workflow implements the saga pattern.
3. Use Conductor SDK and APIs to manage workflows from your application.

