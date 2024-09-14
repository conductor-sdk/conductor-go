package workflow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpTaskHasNoInlineTaskDef(t *testing.T) {
	input := &HttpInput{
		Method: GET,
		Uri:    "http://example.com",
	}
	task := NewHttpTask("http_task_ref", input)
	wfTasks := task.toWorkflowTask()

	wfTask := wfTasks[0]
	assert.Nil(t, wfTask.TaskDefinition)
}

func TestHttpTaskTaskDefMapping(t *testing.T) {
	task := NewHttpTask("http_task_ref", &HttpInput{
		Method: POST,
		Uri:    "http://example.com/api",
	})

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
