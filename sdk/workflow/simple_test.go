package workflow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleTaskHasNoInlineTaskDef(t *testing.T) {
	task := NewSimpleTask("simple", "simple_ref")
	wfTasks := task.toWorkflowTask()

	wfTask := wfTasks[0]
	assert.Nil(t, wfTask.TaskDefinition)
}

func TestSimpleTaskTaskDefMapping(t *testing.T) {
	task := NewSimpleTask("simple", "simple_ref")

	task.CacheConfig("cacheKey", 3600)
	task.RetryPolicy(3, "FIXED", 10, 2)
	task.RateLimitFrequency(60, 5)
	task.ConcurrentExecutionLimit(10)
	task.ExecutionTimeout(300)
	task.PollTimeout(120)
	task.ResponseTimeout(600)
	task.TimeoutPolicy("RETRY")

	wfTasks := task.toWorkflowTask()
	wfTask := wfTasks[0]

	assert.NotNil(t, wfTask.TaskDefinition)

	// Assertions for task definition mappings
	taskDef := wfTask.TaskDefinition
	assert.Equal(t, int32(3), taskDef.RetryCount)
	assert.Equal(t, "FIXED", taskDef.RetryLogic)
	assert.Equal(t, int32(10), taskDef.RetryDelaySeconds)
	assert.Equal(t, int32(2), taskDef.BackoffScaleFactor)
	assert.Equal(t, int32(60), taskDef.RateLimitFrequencyInSeconds)
	assert.Equal(t, int32(5), taskDef.RateLimitPerFrequency)
	assert.Equal(t, int32(10), taskDef.ConcurrentExecLimit)
	assert.Equal(t, int64(300), taskDef.TimeoutSeconds)
	assert.Equal(t, int32(120), taskDef.PollTimeoutSeconds)
	assert.Equal(t, int64(600), taskDef.ResponseTimeoutSeconds)
	assert.Equal(t, "RETRY", taskDef.TimeoutPolicy)
}
