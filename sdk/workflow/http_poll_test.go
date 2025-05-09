package workflow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpPollTaskHasNoInlineTaskDef(t *testing.T) {
	pollInput := &HttpPollInput{
		Method:          GET,
		Uri:             "http://example.com",
		PollingInterval: 30,
	}
	task := NewHttpPollTask("http_poll_task_ref", pollInput)
	wfTasks := task.toWorkflowTask()

	wfTask := wfTasks[0]
	assert.Nil(t, wfTask.TaskDefinition)
}

func TestHttpPollTaskTaskDefMapping(t *testing.T) {
	pollInput := &HttpPollInput{
		Method:              POST,
		Uri:                 "http://example.com/api",
		PollingInterval:     30,
		TerminationCriteria: "$.status == 'FAILED'",
	}
	task := NewHttpPollTask("http_poll_task_ref", pollInput)

	task.CacheConfig("cacheKey", 7200)
	task.RetryPolicy(5, "EXPONENTIAL_BACKOFF", 15, 3)
	task.RateLimitFrequency(120, 10)
	task.ConcurrentExecutionLimit(20)
	task.ExecutionTimeout(600)
	task.PollTimeout(180)
	task.ResponseTimeout(1200)
	task.TimeoutPolicy("TIMEOUT_FAILURE")

	wfTasks := task.toWorkflowTask()
	wfTask := wfTasks[0]

	// Check if TaskDefinition is set correctly
	assert.NotNil(t, wfTask.TaskDefinition)

	// Assertions for task definition mappings
	taskDef := wfTask.TaskDefinition
	assert.Equal(t, int32(5), taskDef.RetryCount)
	assert.Equal(t, "EXPONENTIAL_BACKOFF", taskDef.RetryLogic)
	assert.Equal(t, int32(15), taskDef.RetryDelaySeconds)
	assert.Equal(t, int32(3), taskDef.BackoffScaleFactor)
	assert.Equal(t, int32(120), taskDef.RateLimitFrequencyInSeconds)
	assert.Equal(t, int32(10), taskDef.RateLimitPerFrequency)
	assert.Equal(t, int32(20), taskDef.ConcurrentExecLimit)
	assert.Equal(t, int64(600), taskDef.TimeoutSeconds)
	assert.Equal(t, int32(180), taskDef.PollTimeoutSeconds)
	assert.Equal(t, int64(1200), taskDef.ResponseTimeoutSeconds)
	assert.Equal(t, "TIMEOUT_FAILURE", taskDef.TimeoutPolicy)
}

func TestHttpPollTaskInputParameters(t *testing.T) {
	pollInput := &HttpPollInput{
		Method: GET,
		Uri:    "http://example.com/status",
		Headers: map[string][]string{
			"Authorization": {"Bearer token123"},
		},
		Accept:              "application/json",
		ContentType:         "application/json",
		PollingInterval:     60,
		PollingStrategy:     "FIXED",
		TerminationCriteria: "(function(){ return $.output.response.body.status == 'FAILED';})();",
		Encode:              true,
	}
	task := NewHttpPollTask("http_poll_task_ref", pollInput)

	wfTasks := task.toWorkflowTask()
	wfTask := wfTasks[0]

	// Check task type is correctly set
	assert.Equal(t, string(HTTP_POLL), wfTask.Type_)

	// Check input parameters are correctly set
	inputParams := wfTask.InputParameters

	// Check http_request parameter
	httpRequest, ok := inputParams["http_request"].(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, "GET", httpRequest["method"])
	assert.Equal(t, "http://example.com/status", httpRequest["uri"])
	assert.Equal(t, "application/json", httpRequest["accept"])
	assert.Equal(t, "application/json", httpRequest["contentType"])
	assert.Equal(t, 60, httpRequest["pollingInterval"])
	assert.Equal(t, "FIXED", httpRequest["pollingStrategy"])
	assert.Equal(t, "(function(){ return $.output.response.body.status == 'FAILED';})();", httpRequest["terminationCondition"])
	assert.Equal(t, true, httpRequest["encode"])
}
