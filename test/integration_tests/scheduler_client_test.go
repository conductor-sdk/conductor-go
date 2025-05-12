package integration_tests

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSchedulerResourceApiService(t *testing.T) {
	// Setup
	const WorkflowName = "TestGoSDKWorkflowForSchedulerClient"
	schedulerClient := testdata.SchedulerClient // Assuming this exists in your testdata package

	ctx := context.Background()
	_, _ = testdata.MetadataClient.UnregisterWorkflowDef(ctx, WorkflowName, 1)
	task := model.WorkflowTask{
		Name:              "simple_task",
		TaskReferenceName: "simple_task_ref",
		Description:       "Test Simple Task",
	}

	workflowTasks := []model.WorkflowTask{}
	workflowTasks = append(workflowTasks, task)

	workflowDef := model.WorkflowDef{
		Name:        WorkflowName,
		Description: "Test Workflow created by GO SDK",
		Tasks:       workflowTasks,
	}

	resp, err := testdata.MetadataClient.RegisterWorkflowDef(ctx, true, workflowDef)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)

	// Generate a unique schedule name for testing
	scheduleName := fmt.Sprintf("test_schedule_%d", time.Now().UnixNano())

	// Test case 1: Create a new schedule
	// Create the StartWorkflowRequest for the schedule
	startWorkflowRequest := model.StartWorkflowRequest{
		Name:    WorkflowName,
		Version: 1,
		Input: map[string]interface{}{
			"param1": "value1",
			"param2": 123,
		},
	}

	// Create current time and a time 7 days in the future for schedule start/end
	now := time.Now()
	sevenDaysLater := now.AddDate(0, 0, 7)

	// Create the SaveScheduleRequest
	saveRequest := model.SaveScheduleRequest{
		Name:                        scheduleName,
		CronExpression:              "0 0/5 * * * ?", // Every 5 minutes
		ScheduleStartTime:           now.UnixMilli(),
		ScheduleEndTime:             sevenDaysLater.UnixMilli(),
		StartWorkflowRequest:        &startWorkflowRequest,
		RunCatchupScheduleInstances: false,
	}

	// Save the schedule
	_, resp, err = schedulerClient.SaveSchedule(ctx, saveRequest)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Create cleanup code to delete the schedule at the end
	defer func() {
		_, _, _ = schedulerClient.DeleteSchedule(ctx, scheduleName)
	}()

	// Test case 2: Get the created schedule
	schedule, resp, err := schedulerClient.GetSchedule(ctx, scheduleName)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, scheduleName, schedule.Name)
	assert.Equal(t, "0 0/5 * * * ?", schedule.CronExpression)
	assert.Equal(t, WorkflowName, schedule.StartWorkflowRequest.Name)

	// Test case 3: Add tags to the schedule
	tags := []model.Tag{
		{
			Key:   "environment",
			Value: "test",
			Type_: "metadata",
		},
		{
			Key:   "owner",
			Value: "integration-test",
			Type_: "ownership",
		},
	}

	resp, err = schedulerClient.PutTagForSchedule(ctx, tags, scheduleName)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test case 4: Get tags for the schedule
	retrievedTags, resp, err := schedulerClient.GetTagsForSchedule(ctx, scheduleName)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify tags were added correctly
	assert.GreaterOrEqual(t, len(retrievedTags), 2)
	var foundEnvironmentTag, foundOwnerTag bool
	for _, tag := range retrievedTags {
		if tag.Key == "environment" && tag.Value == "test" {
			foundEnvironmentTag = true
		}
		if tag.Key == "owner" && tag.Value == "integration-test" {
			foundOwnerTag = true
		}
	}
	assert.True(t, foundEnvironmentTag, "Environment tag not found")
	assert.True(t, foundOwnerTag, "Owner tag not found")

	// Test case 5: Get the next few scheduled executions
	getNextOpts := &client.SchedulerResourceApiGetNextFewSchedulesOpts{
		ScheduleStartTime: optional.NewInt64(now.UnixMilli()),
		ScheduleEndTime:   optional.NewInt64(now.Add(24 * time.Hour).UnixMilli()),
		Limit:             optional.NewInt32(5),
	}

	nextSchedules, resp, err := schedulerClient.GetNextFewSchedules(ctx, "0 0/5 * * * ?", getNextOpts)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.GreaterOrEqual(t, len(nextSchedules), 1) // Should have at least one upcoming execution

	// Test case 6: Get all schedules and verify our schedule exists
	getAllOpts := &client.SchedulerResourceApiGetAllSchedulesOpts{
		WorkflowName: optional.NewString(WorkflowName),
	}

	allSchedules, resp, err := schedulerClient.GetAllSchedules(ctx, getAllOpts)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var foundSchedule bool
	for _, s := range allSchedules {
		if s.Name == scheduleName {
			foundSchedule = true
			break
		}
	}
	assert.True(t, foundSchedule, "Created schedule not found in list of all schedules")

	// Test case 7: Pause the schedule
	_, resp, err = schedulerClient.PauseSchedule(ctx, scheduleName)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test case 8: Verify the schedule is paused
	pausedSchedule, resp, err := schedulerClient.GetSchedule(ctx, scheduleName)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.True(t, pausedSchedule.Paused, "Schedule should be paused")

	// Test case 9: Resume the schedule
	_, resp, err = schedulerClient.ResumeSchedule(ctx, scheduleName)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test case 10: Verify the schedule is resumed
	resumedSchedule, resp, err := schedulerClient.GetSchedule(ctx, scheduleName)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.False(t, resumedSchedule.Paused, "Schedule should be resumed")

	// Test case 11: Search for schedule executions
	searchOpts := &client.SchedulerSearchOpts{
		Start: optional.NewInt32(0),
		Size:  optional.NewInt32(10),
		Query: optional.NewString(scheduleName),
	}

	searchResults, resp, err := schedulerClient.SearchV2(ctx, searchOpts)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, searchResults)

	// Test case 12: Delete a tag from the schedule
	tagToDelete := []model.Tag{
		{
			Key:   "environment",
			Value: "test",
			Type_: "metadata",
		},
	}

	resp, err = schedulerClient.DeleteTagForSchedule(ctx, tagToDelete, scheduleName)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test case 13: Verify tag was deleted
	updatedTags, resp, err := schedulerClient.GetTagsForSchedule(ctx, scheduleName)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	environmentTagFound := false
	for _, tag := range updatedTags {
		if tag.Key == "environment" && tag.Value == "test" {
			environmentTagFound = true
			break
		}
	}
	assert.False(t, environmentTagFound, "Environment tag should have been deleted")

	// Test case 14: Delete the schedule
	_, resp, err = schedulerClient.DeleteSchedule(ctx, scheduleName)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test case 15: Verify the schedule was deleted
	getAllOpts = &client.SchedulerResourceApiGetAllSchedulesOpts{}
	allSchedulesAfterDelete, resp, err := schedulerClient.GetAllSchedules(ctx, getAllOpts)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	scheduleFoundAfterDelete := false
	for _, s := range allSchedulesAfterDelete {
		if s.Name == scheduleName {
			scheduleFoundAfterDelete = true
			break
		}
	}
	assert.False(t, scheduleFoundAfterDelete, "Schedule should have been deleted")
}
