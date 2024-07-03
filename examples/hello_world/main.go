package main

import (
	"fmt"
	"os"
	"time"

	hello_world "github.com/conductor-sdk/conductor-go/examples/hello_world/src"
	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/settings"

	"github.com/conductor-sdk/conductor-go/sdk/worker"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
)

var (
	apiClient = client.NewAPIClient(
		authSettings(),
		httpSettings(),
	)
	taskRunner       = worker.NewTaskRunnerWithApiClient(apiClient)
	workflowExecutor = executor.NewWorkflowExecutor(apiClient)
)

func main() {
	//Start the Greet Worker. This worked will process "greet" tasks.
	taskRunner.StartWorker("greet", hello_world.Greet, 1, time.Millisecond*100)

	/* This is used to register the Workflow, it's a one-time process. You can comment from here */
	wf := hello_world.CreateWorkflow(workflowExecutor)
	err := wf.Register(false)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	/* Till Here after registering the workflow*/

	// The workflow input is a Map. You could pass a struct here as well.
	id, err := wf.StartWorkflowWithInput(map[string]string{
		"Name": "Orkes",
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
