## Workflow monitoring

The SDK not only provides translation between a programming language and the Conductor server, but also help you to keep track of your running workflows.

Thinking about the Go way of doing that, channels might be the most suitable way to get notified of your workflow status. Take a look of a quick example here:

```go
workflowExecutor := executor.NewWorkflowExecutor(apiClient)

httpTask = workflow.NewHttpTask(
    "TEST_GO_TASK_HTTP",
    &workflow.HttpInput{
        Uri: "https://catfact.ninja/fact",
    },
)

httpTaskWorkflow := workflow.NewConductorWorkflow(e2e_properties.WorkflowExecutor).
    Name("TEST_GO_WORKFLOW_HTTP").
    Version(1).
    Add(httpTask)

workflowId, workflowExecutionChannel, err := httpTaskWorkflow.ExecuteWorkflowWithInput(nil)

workflow := <-workflowExecutionChannel

if workflow.Status != string(workflow_status.COMPLETED) {
    return fmt.Errorf("workflow finished with unexpected status: %s", workflow.Status)
}

return nil
```

This example do:
* Create Workflow Executor
* Create HTTP task
  * Don't forget to register each workflow task before registering the workflow
* Create Workflow with HTTP task 
* Start a workflow with null input and get it's execution channel
* Receive workflow response from channel
* Validate workflow status

Well, this should work until you face some faulty workflow execution that got stuck. Instead of manually waiting for the channel response, you can use a function in the SDK that waits up to a timeout:

```go
workflow, err := executor.WaitForWorkflowCompletionUntilTimeout(
    workflowExecutionChannel,
    5 * time.Second,
)
```

This example waits for the event that happens first:
* Workflow is in terminal state
* Timeout
