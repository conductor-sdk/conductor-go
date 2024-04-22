package unit_tests

import (
	"encoding/json"
	"fmt"
	"github.com/conductor-sdk/conductor-go/internal/testdata"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRetrySettings(t *testing.T) {
	simpleTask := workflow.NewSimpleTask("worker_task", "worker_task_ref")
	simpleTask.RetryPolicy(2, workflow.FixedRetry, 10, 1)
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("workflow_with_task_retries").
		Version(1).
		Add(simpleTask)
	workflowDef := wf.ToWorkflowDef()
	assert.NotNil(t, workflowDef)
	assert.Equal(t, 1, len(workflowDef.Tasks))
	workflowTask := workflowDef.Tasks[0]
	assert.NotNil(t, workflowTask.TaskDefinition)
	assert.Equal(t, int32(10), workflowTask.TaskDefinition.RetryDelaySeconds)
	assert.Equal(t, string(workflow.FixedRetry), workflowTask.TaskDefinition.RetryLogic)
	json, _ := json.Marshal(workflowDef)
	fmt.Println(string(json))
}
