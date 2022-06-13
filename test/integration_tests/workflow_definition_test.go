package integration_tests

import (
	"fmt"
	"github.com/conductor-sdk/conductor-go/internal/testdata"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWorkflowCreation(t *testing.T) {

	workflow := testdata.NewKitchenSinkWorkflow(testdata.WorkflowExecutor)
	startWorkflowRequest := model.StartWorkflowRequest{
		Name: workflow.GetName(),
	}
	id, err := workflow.StartWorkflow(&startWorkflowRequest)
	assert.NoError(t, err, "Failed to start the workflow", err)
	assert.NotEmpty(t, id, "Workflow Id is null", id)
	fmt.Println("Workflow Id is ", id)
}
