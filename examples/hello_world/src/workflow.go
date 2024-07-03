// Hello World Application Using Conductor
package hello_world

import (
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
)

// Creates the "greetings" workflow definition.
// This single-task workflow expects an input "Name" which is used by
// the "greet" Task to say hello!
func CreateWorkflow(executor *executor.WorkflowExecutor) *workflow.ConductorWorkflow {
	wf := workflow.NewConductorWorkflow(executor).
		Name("greetings").
		Version(1).
		Description("Greetings workflow - Greets a user by their name").
		TimeoutPolicy(workflow.TimeOutWorkflow, 600)

	//New Simple Task - "greet" Task
	greet := workflow.NewSimpleTask("greet", "greet_ref").
		Input("name", "${workflow.input.Name}")

	//Add task to workflow
	wf.Add(greet)

	//Add the output of the workflow from the task
	wf.OutputParameters(map[string]interface{}{
		"Greetings": greet.OutputRef("greetings"),
	})

	return wf
}

func GetTaskDefinitions() []model.TaskDef {
	taskDefs := []model.TaskDef{
		{Name: "greet", TimeoutSeconds: 60},
	}
	return taskDefs
}
