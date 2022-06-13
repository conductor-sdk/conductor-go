# Workflow SDK

APIs to create and manage the workflows.

## Workflow Creation APIs

### ConductorWorkflow
ConductorWorkflow is the SDK reprentation of a Conductor workflow.

**Create a `ConductorWorkflow` instance**
```go

//API client instance with server address and authentication details
apiClient := conductor_http_client.NewAPIClient(
    settings.NewAuthenticationSettings(
        KEY,
        SECRET,
    ),
    settings.NewHttpSettings(
        "https://play.orkes.io/api",
    ))

//Create new workflow executor
executor := executor.NewWorkflowExecutor(apiClient)

//Create a new ConductorWorkflow instance
conductorWorkflow := workflow.NewConductorWorkflow(executor).
Name("my_first_workflow").
Version(1).
OwnerEmail("developers@orkes.io")

//now, let's add a couple of simple tasks
conductorWorkflow.
	Add(workflow.NewSimpleTask("simple_task_1", "simple_task_1")).
    Add(workflow.NewSimpleTask("simple_task_2", "simple_task_2"))

//Register the workflow with server
conductorWorkflow.Register(true)        //Overwrite the existing definition with the new one
```
### Execute Workflow

#### Using Workflow Executor to start previously registered workflow
```go
//Input to the workflow
//Input can be either a map or a struct that is serializable to a JSON map
workflowInput := map[string]interface{}{}

workflowId, err := executor.StartWorkflow(&model.StartWorkflowRequest{
    Name:  conductorWorkflow.GetName(),
    Input: workflowInput,
})
```
Input to the workflow start request can be either a `map[string]interface{}` or a `struct` that is serializable as JSON map.
Here is an example of a workflow input:

```go
type WorkflowInput struct {
	Name string
	Address []string
}
```
Using this struct instance as workflow input
```go
workflowId, err := executor.StartWorkflow(&model.StartWorkflowRequest{
  Name:  conductorWorkflow.GetName(),
  Input: &WorkflowInput{
  Name: "John Doe",
  Address: []string{"street", "city", "zip"},
  },
})
```

#### Using `ConductorWorkflow` as the code to start workflows inline
**:warning:** `ConductorWorkflow` interface allows you to execute workflows defined using code without registering.
This is useful for one-off ad-hoc workflows.  For the production use cases however, you should always register
workflows with server and use the `executor` to start workflows.