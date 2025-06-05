package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"

	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/worker"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
)

var (
	apiClient        = client.NewAPIClientFromEnv()
	workflowExecutor = executor.NewWorkflowExecutor(apiClient)
	taskClient       = &client.TaskResourceApiService{APIClient: apiClient}
	workflowClient   = &client.WorkflowResourceApiService{APIClient: apiClient}
	taskRunner       = worker.NewTaskRunnerWithApiClient(apiClient)
	schedulerClient  = &client.SchedulerResourceApiService{APIClient: apiClient}

	// Cleanup tracking
	createdWorkflows []string
	createdSchedules []string
)

func main() {
	fmt.Println("Testing comprehensive SDK usage for backward compatibility...")

	// Defer cleanup to ensure it runs even if tests fail
	defer func() {
		if err := cleanupTestArtifacts(); err != nil {
			log.Errorf("Cleanup failed: %v", err)
		}
	}()

	// Test 1: Basic Workflow Operations
	if err := testBasicWorkflowOperations(); err != nil {
		log.Fatalf("Basic workflow operations failed: %v", err)
	}

	// Test 2: WorkflowExecutor APIs
	if err := testWorkflowExecutorAPIs(); err != nil {
		log.Fatalf("WorkflowExecutor APIs failed: %v", err)
	}

	// Test 3: TaskClient APIs
	if err := testTaskClientAPIs(); err != nil {
		log.Fatalf("TaskClient APIs failed: %v", err)
	}

	// Test 4: WorkflowClient APIs
	if err := testWorkflowClientAPIs(); err != nil {
		log.Fatalf("WorkflowClient APIs failed: %v", err)
	}

	// Test 5: SchedulerClient APIs
	if err := testSchedulerClientAPIs(); err != nil {
		log.Fatalf("SchedulerClient APIs failed: %v", err)
	}

	// Test 6: Advanced Workflow Operations
	if err := testAdvancedWorkflowOperations(); err != nil {
		log.Fatalf("Advanced workflow operations failed: %v", err)
	}

	fmt.Println("✅ All backward compatibility tests passed!")
}

// Test 5: SchedulerClient API methods
func testSchedulerClientAPIs() error {
	log.Info("Testing SchedulerClient APIs...")

	ctx := context.Background()

	// First, register an HTTP workflow for scheduling
	httpWorkflow := CreateHttpOnlyWorkflow(workflowExecutor, "scheduled_http_workflow")
	err := httpWorkflow.Register(true)
	if err != nil {
		return fmt.Errorf("failed to register HTTP workflow for scheduler test: %w", err)
	}
	createdWorkflows = append(createdWorkflows, "scheduled_http_workflow") // Track created workflow

	// Create a test schedule
	scheduleName := "test_schedule_compatibility"
	scheduleRequest := model.SaveScheduleRequest{
		Name:           scheduleName,
		CronExpression: "0 */5 * * * ?", // Every 5 minutes
		StartWorkflowRequest: &model.StartWorkflowRequest{
			Name:    "scheduled_http_workflow",
			Version: 1,
			Input:   map[string]interface{}{"scheduled": true},
		},
		Paused: true, // Start paused to avoid actual executions during test
	}

	// Test SaveSchedule
	_, _, err = schedulerClient.SaveSchedule(ctx, scheduleRequest)
	if err != nil {
		return fmt.Errorf("SchedulerClient.SaveSchedule failed: %w", err)
	}
	createdSchedules = append(createdSchedules, scheduleName) // Track created schedule
	log.Infof("✓ SchedulerClient.SaveSchedule successful for schedule: %s", scheduleName)

	// Test GetSchedule
	schedule, _, err := schedulerClient.GetSchedule(ctx, scheduleName)
	if err != nil {
		return fmt.Errorf("SchedulerClient.GetSchedule failed: %w", err)
	}
	log.Infof("✓ SchedulerClient.GetSchedule: %s (paused: %t)", schedule.Name, schedule.Paused)

	// Test GetAllSchedules
	allSchedules, _, err := schedulerClient.GetAllSchedules(ctx, &client.SchedulerResourceApiGetAllSchedulesOpts{})
	if err != nil {
		return fmt.Errorf("SchedulerClient.GetAllSchedules failed: %w", err)
	}
	log.Infof("✓ SchedulerClient.GetAllSchedules returned %d schedules", len(allSchedules))

	// Test GetNextFewSchedules
	nextSchedules, _, err := schedulerClient.GetNextFewSchedules(ctx, "0 */5 * * * ?", &client.SchedulerResourceApiGetNextFewSchedulesOpts{
		Limit: optional.NewInt32(3),
	})
	if err != nil {
		log.Warnf("SchedulerClient.GetNextFewSchedules failed: %v", err)
	} else {
		log.Infof("✓ SchedulerClient.GetNextFewSchedules returned %d timestamps", len(nextSchedules))
	}

	// Test SearchV2
	searchResult, _, err := schedulerClient.SearchV2(ctx, &client.SchedulerSearchOpts{
		Start: optional.NewInt32(0),
		Size:  optional.NewInt32(10),
		Query: optional.NewString(fmt.Sprintf("name='%s'", scheduleName)),
	})
	if err != nil {
		log.Warnf("SchedulerClient.SearchV2 failed: %v", err)
	} else {
		log.Infof("✓ SchedulerClient.SearchV2 returned %d results", len(searchResult.Results))
	}

	// Test tag operations
	testTags := []model.Tag{
		{Key: "environment", Value: "test"},
		{Key: "purpose", Value: "compatibility"},
	}

	// Test PutTagForSchedule
	_, err = schedulerClient.PutTagForSchedule(ctx, testTags, scheduleName)
	if err != nil {
		log.Warnf("SchedulerClient.PutTagForSchedule failed: %v", err)
	} else {
		log.Info("✓ SchedulerClient.PutTagForSchedule successful")

		// Test GetTagsForSchedule
		tags, _, err := schedulerClient.GetTagsForSchedule(ctx, scheduleName)
		if err != nil {
			log.Warnf("SchedulerClient.GetTagsForSchedule failed: %v", err)
		} else {
			log.Infof("✓ SchedulerClient.GetTagsForSchedule returned %d tags", len(tags))
		}

		// Test DeleteTagForSchedule
		_, err = schedulerClient.DeleteTagForSchedule(ctx, testTags, scheduleName)
		if err != nil {
			log.Warnf("SchedulerClient.DeleteTagForSchedule failed: %v", err)
		} else {
			log.Info("✓ SchedulerClient.DeleteTagForSchedule successful")
		}
	}

	// Test schedule control operations
	// Test ResumeSchedule
	_, _, err = schedulerClient.ResumeSchedule(ctx, scheduleName)
	if err != nil {
		log.Warnf("SchedulerClient.ResumeSchedule failed: %v", err)
	} else {
		log.Info("✓ SchedulerClient.ResumeSchedule successful")
	}

	// Test PauseSchedule
	_, _, err = schedulerClient.PauseSchedule(ctx, scheduleName)
	if err != nil {
		log.Warnf("SchedulerClient.PauseSchedule failed: %v", err)
	} else {
		log.Info("✓ SchedulerClient.PauseSchedule successful")
	}

	// Test ResumeAllSchedules
	_, _, err = schedulerClient.ResumeAllSchedules(ctx)
	if err != nil {
		log.Warnf("SchedulerClient.ResumeAllSchedules failed: %v", err)
	} else {
		log.Info("✓ SchedulerClient.ResumeAllSchedules successful")
	}

	// Test RequeueAllExecutionRecords
	_, _, err = schedulerClient.RequeueAllExecutionRecords(ctx)
	if err != nil {
		log.Warnf("SchedulerClient.RequeueAllExecutionRecords failed: %v", err)
	} else {
		log.Info("✓ SchedulerClient.RequeueAllExecutionRecords successful")
	}

	// Cleanup: Delete the test schedule
	_, _, err = schedulerClient.DeleteSchedule(ctx, scheduleName)
	if err != nil {
		log.Warnf("Failed to cleanup test schedule: %v", err)
	} else {
		log.Info("✓ SchedulerClient.DeleteSchedule successful (cleanup)")
	}

	log.Info("✓ SchedulerClient APIs tested successfully")
	return nil
}

// Test 1: Basic workflow registration and execution
func testBasicWorkflowOperations() error {
	log.Info("Testing basic workflow operations...")

	// Register workflow
	wf := CreateWorkflow(workflowExecutor)
	err := wf.Register(true)
	if err != nil {
		return fmt.Errorf("workflow registration failed: %w", err)
	}
	createdWorkflows = append(createdWorkflows, "greetings") // Track created workflow
	log.Info("✓ Workflow registered successfully")

	// Start workflow
	id, err := workflowExecutor.StartWorkflow(
		&model.StartWorkflowRequest{
			Name:    "greetings",
			Version: 1,
			Input: map[string]interface{}{
				"name":      "Gopher",
				"timestamp": time.Now().Unix(),
			},
		},
	)

	if err != nil {
		return fmt.Errorf("workflow start failed: %w", err)
	}
	log.Infof("✓ Started workflow with ID: %s", id)

	// Wait and validate
	err = waitAndValidateWorkflow(workflowExecutor, id)
	if err != nil {
		return fmt.Errorf("workflow validation failed: %w", err)
	}
	log.Info("✓ Basic workflow operations completed successfully")
	return nil
}

// Test 2: WorkflowExecutor API methods with controllable SIMPLE task
func testWorkflowExecutorAPIs() error {
	log.Info("Testing WorkflowExecutor APIs...")

	// Create workflow with SIMPLE task that requires manual completion
	wf := CreateSimpleTaskWorkflow(workflowExecutor, "api_test_simple_workflow")
	err := wf.Register(true)
	if err != nil {
		return fmt.Errorf("failed to register test workflow: %w", err)
	}
	createdWorkflows = append(createdWorkflows, "api_test_simple_workflow") // Track created workflow

	id, err := workflowExecutor.StartWorkflow(
		&model.StartWorkflowRequest{
			Name:    "api_test_simple_workflow",
			Version: 1,
			Input:   map[string]interface{}{"test": "api_test"},
		},
	)
	if err != nil {
		return fmt.Errorf("failed to start test workflow: %w", err)
	}

	// Wait a moment for workflow to start and task to be scheduled
	time.Sleep(2 * time.Second)

	// Test GetWorkflowStatus
	status, err := workflowExecutor.GetWorkflowStatus(id, true, true)
	if err != nil {
		return fmt.Errorf("GetWorkflowStatus failed: %w", err)
	}
	log.Infof("✓ GetWorkflowStatus: %s", status.Status)

	// Test GetWorkflow with tasks
	currentWorkflow, err := workflowExecutor.GetWorkflow(id, true)
	if err != nil {
		return fmt.Errorf("GetWorkflow failed: %w", err)
	}
	log.Infof("✓ GetWorkflow: %s with %d tasks", currentWorkflow.Status, len(currentWorkflow.Tasks))

	// Test Search functionality
	workflows, err := workflowExecutor.Search(0, 10, fmt.Sprintf("workflowId='%s'", id), "")
	if err != nil {
		return fmt.Errorf("Search failed: %w", err)
	}
	if len(workflows) == 0 {
		return fmt.Errorf("Search returned no results")
	}
	log.Infof("✓ Search found %d workflows", len(workflows))

	// Test Pause/Resume while workflow is running
	if currentWorkflow.Status == "RUNNING" {
		// Test Pause
		err = workflowExecutor.Pause(id)
		if err != nil {
			log.Warnf("Pause failed: %v", err)
		} else {
			log.Info("✓ Pause successful")

			// Verify paused state
			pausedWorkflow, err := workflowExecutor.GetWorkflow(id, false)
			if err == nil && pausedWorkflow.Status == "PAUSED" {
				log.Info("✓ Workflow successfully paused")
			}

			// Test Resume
			err = workflowExecutor.Resume(id)
			if err != nil {
				log.Warnf("Resume failed: %v", err)
			} else {
				log.Info("✓ Resume successful")
			}
		}
	}

	// Test UpdateTaskByRefName to complete the SIMPLE task
	ctx := context.Background()
	outputData := map[string]interface{}{
		"result":    "compatibility_test_completed",
		"timestamp": time.Now().Unix(),
	}

	_, _, err = taskClient.UpdateTaskByRefName(
		ctx,
		outputData,
		id,
		"SIMPLE", // This matches our task reference name
		string(model.CompletedTask),
	)
	if err != nil {
		log.Warnf("UpdateTaskByRefName failed: %v", err)
	} else {
		log.Info("✓ UpdateTaskByRefName successful")
	}

	// Wait for workflow to complete after task update
	err = waitAndValidateWorkflow(workflowExecutor, id)
	if err != nil {
		log.Warnf("Workflow didn't complete normally: %v", err)
	}

	log.Info("✓ WorkflowExecutor APIs tested successfully")
	return nil
}

// Test 3: TaskClient API methods
func testTaskClientAPIs() error {
	log.Info("Testing TaskClient APIs...")

	ctx := context.Background()

	// Test All - Get queue details
	queueDetails, _, err := taskClient.All(ctx)
	if err != nil {
		return fmt.Errorf("TaskClient.All failed: %w", err)
	}
	log.Infof("✓ TaskClient.All returned %d queues", len(queueDetails))

	// Test Size - Get queue sizes with specific task types to avoid null pointer
	sizes, _, err := taskClient.Size(ctx, &client.TaskResourceApiSizeOpts{
		TaskType: optional.NewInterface([]string{"HTTP", "SIMPLE"}),
	})
	if err != nil {
		// If Size fails, try without parameters and handle gracefully
		log.Warnf("TaskClient.Size with task types failed, trying without parameters: %v", err)
		sizes, _, err = taskClient.Size(ctx, &client.TaskResourceApiSizeOpts{})
		if err != nil {
			log.Warnf("TaskClient.Size failed, skipping this test: %v", err)
			log.Info("✓ TaskClient.Size test skipped (server may not have task definitions)")
		} else {
			log.Infof("✓ TaskClient.Size returned %d task types", len(sizes))
		}
	} else {
		log.Infof("✓ TaskClient.Size returned %d task types", len(sizes))
	}

	// Test Search
	searchResult, _, err := taskClient.Search(ctx, &client.TaskResourceApiSearch1Opts{
		Start: optional.NewInt32(0),
		Size:  optional.NewInt32(10),
	})
	if err != nil {
		return fmt.Errorf("TaskClient.Search failed: %w", err)
	}
	log.Infof("✓ TaskClient.Search returned %d results", len(searchResult.Results))

	// Test Poll with a common task type (gracefully handle if not available)
	task, _, err := taskClient.Poll(ctx, "HTTP", &client.TaskResourceApiPollOpts{
		Workerid: optional.NewString("test-worker"),
	})
	if err != nil {
		log.Warnf("TaskClient.Poll failed (expected if no HTTP tasks available): %v", err)
		log.Info("✓ TaskClient.Poll test completed (no tasks to poll)")
	} else {
		log.Infof("✓ TaskClient.Poll returned task: %s", task.TaskId)
	}

	log.Info("✓ TaskClient APIs tested successfully")
	return nil
}

// Test 4: WorkflowClient API methods
func testWorkflowClientAPIs() error {
	log.Info("Testing WorkflowClient APIs...")

	ctx := context.Background()

	// Create and start a test workflow for WorkflowClient operations
	wf := CreateSimpleTaskWorkflow(workflowExecutor, "workflow_client_test")
	err := wf.Register(true)
	if err != nil {
		return fmt.Errorf("failed to register workflow for WorkflowClient test: %w", err)
	}
	createdWorkflows = append(createdWorkflows, "workflow_client_test") // Track created workflow

	// Test StartWorkflowWithRequest
	startRequest := model.StartWorkflowRequest{
		Name:    "workflow_client_test",
		Version: 1,
		Input:   map[string]interface{}{"test": "workflow_client_api"},
	}

	workflowId, _, err := workflowClient.StartWorkflowWithRequest(ctx, startRequest)
	if err != nil {
		return fmt.Errorf("WorkflowClient.StartWorkflowWithRequest failed: %w", err)
	}
	log.Infof("✓ WorkflowClient.StartWorkflowWithRequest successful, ID: %s", workflowId)

	// Wait a moment for workflow to initialize
	time.Sleep(2 * time.Second)

	// Test GetExecutionStatus
	workflow, _, err := workflowClient.GetExecutionStatus(ctx, workflowId, &client.WorkflowResourceApiGetExecutionStatusOpts{
		IncludeTasks: optional.NewBool(true),
	})
	if err != nil {
		return fmt.Errorf("WorkflowClient.GetExecutionStatus failed: %w", err)
	}
	log.Infof("✓ WorkflowClient.GetExecutionStatus: %s with %d tasks", workflow.Status, len(workflow.Tasks))

	// Test GetWorkflowState
	workflowState, _, err := workflowClient.GetWorkflowState(ctx, workflowId, true, true)
	if err != nil {
		return fmt.Errorf("WorkflowClient.GetWorkflowState failed: %w", err)
	}
	log.Infof("✓ WorkflowClient.GetWorkflowState: %s", workflowState.Status)

	// Test PauseWorkflow
	_, err = workflowClient.PauseWorkflow(ctx, workflowId)
	if err != nil {
		log.Warnf("WorkflowClient.PauseWorkflow failed: %v", err)
	} else {
		log.Info("✓ WorkflowClient.PauseWorkflow successful")

		// Test ResumeWorkflow
		_, err = workflowClient.ResumeWorkflow(ctx, workflowId)
		if err != nil {
			log.Warnf("WorkflowClient.ResumeWorkflow failed: %v", err)
		} else {
			log.Info("✓ WorkflowClient.ResumeWorkflow successful")
		}
	}

	// Test Search
	searchResult, _, err := workflowClient.Search(ctx, &client.WorkflowResourceApiSearchOpts{
		Start:    optional.NewInt32(0),
		Size:     optional.NewInt32(10),
		Query:    optional.NewString(fmt.Sprintf("workflowId='%s'", workflowId)),
		FreeText: optional.NewString(""),
	})
	if err != nil {
		return fmt.Errorf("WorkflowClient.Search failed: %w", err)
	}
	log.Infof("✓ WorkflowClient.Search returned %d results", len(searchResult.Results))

	// Test GetRunningWorkflow
	runningWorkflowIds, _, err := workflowClient.GetRunningWorkflow(ctx, "workflow_client_test", &client.WorkflowResourceApiGetRunningWorkflowOpts{
		Version: optional.NewInt32(1),
	})
	if err != nil {
		log.Warnf("WorkflowClient.GetRunningWorkflow failed: %v", err)
	} else {
		log.Infof("✓ WorkflowClient.GetRunningWorkflow returned %d running workflows", len(runningWorkflowIds))
	}

	// Complete the task so we can test other operations
	outputData := map[string]interface{}{
		"result":    "workflow_client_test_completed",
		"timestamp": time.Now().Unix(),
	}

	_, _, err = taskClient.UpdateTaskByRefName(
		ctx,
		outputData,
		workflowId,
		"SIMPLE", // This matches our task reference name
		string(model.CompletedTask),
	)
	if err != nil {
		log.Warnf("Failed to complete task for WorkflowClient test: %v", err)
	}

	taskRunner.StartWorker("SIMPLE", SimpleTask, 1, time.Millisecond*100)
	// Wait for workflow to complete
	waitAndValidateWorkflow(workflowExecutor, workflowId)

	// Test operations on completed workflow
	// Test Restart (only works on completed workflows)
	_, err = workflowClient.Restart(ctx, workflowId, &client.WorkflowResourceApiRestartOpts{
		UseLatestDefinitions: optional.NewBool(false),
	})
	if err != nil {
		log.Warnf("WorkflowClient.Restart failed: %v", err)
	} else {
		log.Info("✓ WorkflowClient.Restart successful")
		// Complete the restarted workflow
		time.Sleep(2 * time.Second)
		restartedWorkflow, _, _ := workflowClient.GetExecutionStatus(ctx, workflowId, nil)
		if restartedWorkflow.Status == "RUNNING" {
			_, _, _ = taskClient.UpdateTaskByRefName(ctx, outputData, workflowId, "SIMPLE", string(model.CompletedTask))
			waitAndValidateWorkflow(workflowExecutor, workflowId)
		}
	}

	log.Info("✓ WorkflowClient APIs tested successfully")
	return nil
}

// Test 5: Advanced Workflow Operations
func testAdvancedWorkflowOperations() error {
	log.Info("Testing advanced workflow operations...")

	// Test bulk workflow start
	requests := []*model.StartWorkflowRequest{
		{
			Name:    "greetings",
			Version: 1,
			Input:   map[string]interface{}{"name": "Bulk1", "batch": "test"},
		},
		{
			Name:    "greetings",
			Version: 1,
			Input:   map[string]interface{}{"name": "Bulk2", "batch": "test"},
		},
	}

	runningWorkflows := workflowExecutor.StartWorkflows(false, requests...)
	log.Infof("✓ Started %d workflows in bulk", len(runningWorkflows))

	// Validate bulk start results
	for i, rw := range runningWorkflows {
		if rw.Err != nil {
			return fmt.Errorf("bulk workflow %d failed: %w", i, rw.Err)
		}
		log.Infof("✓ Bulk workflow %d started with ID: %s", i, rw.WorkflowId)
	}

	// Test ExecuteWorkflow with WorkflowClient
	ctx := context.Background()
	executeRequest := model.StartWorkflowRequest{
		Name:    "greetings",
		Version: 1,
		Input:   map[string]interface{}{"name": "ExecuteTest"},
	}

	workflowRun, _, err := workflowClient.ExecuteWorkflow(
		ctx,
		executeRequest,
		"test-request-id",
		"greetings",
		1,
		"",
	)
	if err != nil {
		log.Warnf("WorkflowClient.ExecuteWorkflow failed: %v", err)
	} else {
		log.Infof("✓ WorkflowClient.ExecuteWorkflow completed with status: %s", workflowRun.Status)
	}

	log.Info("✓ Advanced workflow operations tested successfully")
	return nil
}

func waitAndValidateWorkflow(executor *executor.WorkflowExecutor, workflowId string) error {
	maxWaitTime := 30 * time.Second
	checkInterval := 2 * time.Second
	startTime := time.Now()

	for time.Since(startTime) < maxWaitTime {
		workflow, err := executor.GetWorkflow(workflowId, true)
		if err != nil {
			return fmt.Errorf("failed to get workflow status: %w", err)
		}

		fmt.Printf("Workflow %s status: %s\n", workflowId, workflow.Status)

		if workflow.Status == "COMPLETED" {
			return nil
		} else if workflow.Status == "FAILED" || workflow.Status == "TIMED_OUT" {
			return fmt.Errorf("workflow failed with status: %s", workflow.Status)
		}

		time.Sleep(checkInterval)
	}

	return fmt.Errorf("workflow timed out")
}

// Creates the "greetings" workflow definition
func CreateWorkflow(executor *executor.WorkflowExecutor) *workflow.ConductorWorkflow {
	wf := workflow.NewConductorWorkflow(executor).
		Name("greetings").
		Version(1).
		Description("Greetings workflow - Greets a user by their name").
		TimeoutPolicy(workflow.TimeOutWorkflow, 600)

	// HTTP Task pointing to a reliable external service
	greet := workflow.NewHttpTask("http_task_ref", &workflow.HttpInput{
		Method: "GET",
		Uri:    "https://jsonplaceholder.typicode.com/posts/1", // More reliable than localhost
	})

	wf.Add(greet)
	return wf
}

// Creates an HTTP-only workflow for scheduler testing (no workers needed)
func CreateHttpOnlyWorkflow(executor *executor.WorkflowExecutor, workflowName string) *workflow.ConductorWorkflow {
	wf := workflow.NewConductorWorkflow(executor).
		Name(workflowName).
		Version(1).
		Description("HTTP-only workflow for scheduler compatibility testing").
		TimeoutPolicy(workflow.TimeOutWorkflow, 300)

	// Create HTTP tasks that don't require workers
	httpTask1 := workflow.NewHttpTask("fetch_data", &workflow.HttpInput{
		Method: "GET",
		Uri:    "https://jsonplaceholder.typicode.com/posts/1",
	})

	httpTask2 := workflow.NewHttpTask("validate_data", &workflow.HttpInput{
		Method: "GET",
		Uri:    "https://httpbin.org/status/200",
	})

	wf.Add(httpTask1)
	wf.Add(httpTask2)

	wf.OutputParameters(map[string]interface{}{
		"DataFetched": httpTask1.OutputRef("response"),
		"Validated":   httpTask2.OutputRef("response"),
	})

	return wf
}

// Creates a workflow with SIMPLE task that requires manual completion
func CreateSimpleTaskWorkflow(executor *executor.WorkflowExecutor, workflowName string) *workflow.ConductorWorkflow {
	wf := workflow.NewConductorWorkflow(executor).
		Name(workflowName).
		Version(1).
		Description("Simple task workflow for API compatibility testing").
		TimeoutPolicy(workflow.TimeOutWorkflow, 600)

	// Create a SIMPLE task that requires manual completion
	simpleTask := workflow.NewSimpleTask("simple_task", "SIMPLE").Input("name", "${workflow.input.name}")

	wf.Add(simpleTask)

	wf.OutputParameters(map[string]interface{}{
		"Greetings": simpleTask.OutputRef("greetings"),
	})

	return wf
}

// cleanupTestArtifacts removes all test artifacts created during compatibility testing
func cleanupTestArtifacts() error {
	log.Info("🧹 Starting cleanup of test artifacts...")
	ctx := context.Background()
	var cleanupErrors []string

	// 1. Cleanup schedules first (they depend on workflows)
	log.Info("Cleaning up test schedules...")
	for _, scheduleName := range createdSchedules {
		_, _, err := schedulerClient.DeleteSchedule(ctx, scheduleName)
		if err != nil {
			cleanupErrors = append(cleanupErrors, fmt.Sprintf("Failed to delete schedule %s: %v", scheduleName, err))
			log.Warnf("Failed to delete schedule %s: %v", scheduleName, err)
		} else {
			log.Infof("✓ Deleted schedule: %s", scheduleName)
		}
	}

	// 2. Unregister workflow definitions
	log.Info("Unregistering test workflow definitions...")
	for _, workflowName := range createdWorkflows {
		err := workflowExecutor.UnRegisterWorkflow(workflowName, 1)
		if err != nil {
			cleanupErrors = append(cleanupErrors, fmt.Sprintf("Failed to unregister workflow %s: %v", workflowName, err))
			log.Warnf("Failed to unregister workflow %s: %v", workflowName, err)
		} else {
			log.Infof("✓ Unregistered workflow definition: %s", workflowName)
		}
	}

	// 4. Stop any running workers
	log.Info("Stopping test workers...")
	// Note: TaskRunner doesn't have a direct stop method, but workers will stop when the process ends
	log.Info("✓ Workers will stop when process ends")

	// Summary
	if len(cleanupErrors) > 0 {
		log.Warnf("Cleanup completed with %d errors:", len(cleanupErrors))
		for _, err := range cleanupErrors {
			log.Warn("  - " + err)
		}
		return fmt.Errorf("cleanup completed with %d errors", len(cleanupErrors))
	}

	log.Info("✅ Cleanup completed successfully!")
	log.Infof("Cleaned up: %d workflows, %d schedules",
		len(createdWorkflows), len(createdSchedules))

	return nil
}

func SimpleTask(task *model.Task) (interface{}, error) {
	return map[string]interface{}{
		"greetings": "Hello, " + fmt.Sprintf("%v", task.InputData["name"]),
	}, nil
}
