package integration_tests

import (
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"github.com/stretchr/testify/assert"
)

func TestRetryNotFound(t *testing.T) {
	executor := testdata.WorkflowExecutor
	// Workflow id is hardcoded on purpose. It should not be found.
	err := executor.Retry("2b3ea839-9aeb-11ef-9ac5-ce590b39fb93", true)
	assert.Error(t, err, "Retry is expected to return an error")

	if swaggerErr, ok := err.(client.GenericSwaggerError); ok {
		// hmm... this should be a 404 or 400, but it's a 500 right now.
		assert.Error(t, err, "GetWorkflow was expected to return a 500 error")
		assert.Equal(t, 500, swaggerErr.StatusCode())
	} else {
		assert.Fail(t, "err is not of type GenericSwaggerError ")
	}

}
