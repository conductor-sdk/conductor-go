//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package integration_tests

import (
	"github.com/conductor-sdk/conductor-go/sdk/worker"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/conductor-sdk/conductor-go/test/testdata"
)

// Test data structures
type TestData struct {
	InputKey TestInputKey `json:"inputKey"`
}

type TestInputKey struct {
	StatusCode   int    `json:"statusCode"`
	ReasonPhrase string `json:"reasonPhrase"`
	HostName     string `json:"hostName"`
}

func TestWorkerBatchSize(t *testing.T) {
	simpleTaskWorkflow := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_SIMPLE").
		Version(1).
		Add(testdata.TestSimpleTask)
	err := testdata.TaskRunner.StartWorker(
		testdata.TestSimpleTask.ReferenceName(),
		testdata.SimpleWorker,
		5,
		testdata.WorkerPollInterval,
	)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(1 * time.Second)
	if testdata.TaskRunner.GetBatchSizeForTask(testdata.TestSimpleTask.ReferenceName()) != 5 {
		t.Fatal("unexpected batch size")
	}
	err = testdata.ValidateWorkflowBulk(simpleTaskWorkflow, testdata.WorkflowValidationTimeout, testdata.WorkflowBulkQty)
	if err != nil {
		t.Fatal(err)
	}
	err = testdata.TaskRunner.SetBatchSize(
		testdata.TestSimpleTask.ReferenceName(),
		0,
	)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(1 * time.Second)
	if testdata.TaskRunner.GetBatchSizeForTask(testdata.TestSimpleTask.ReferenceName()) != 0 {
		t.Fatal("unexpected batch size")
	}
	err = testdata.TaskRunner.SetBatchSize(
		testdata.TestSimpleTask.ReferenceName(),
		8,
	)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(1 * time.Second)
	if testdata.TaskRunner.GetBatchSizeForTask(testdata.TestSimpleTask.ReferenceName()) != 8 {
		t.Fatal("unexpected batch size")
	}
	err = testdata.ValidateWorkflowBulk(simpleTaskWorkflow, testdata.WorkflowValidationTimeout, testdata.WorkflowBulkQty)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFaultyWorker(t *testing.T) {

	taskName := "TEST_GO_FAULTY_TASK"
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_FAULTY_WORKFLOW").
		Version(1).
		Add(workflow.NewSimpleTask(taskName, taskName))
	err := wf.Register(true)
	if err != nil {
		t.Fatal(err)
	}
	err = testdata.TaskRunner.StartWorker(
		taskName,
		testdata.FaultyWorker,
		5,
		testdata.WorkerPollInterval,
	)
	if err != nil {
		t.Fatal(err)
	}
	err = testdata.ValidateWorkflow(wf, 5*time.Second, model.FailedWorkflow)
	if err != nil {
		t.Fatal(err)
	}
}

func TestWorkerWithNonRetryableError(t *testing.T) {

	taskName := "TEST_GO_NON_RETRYABLE_ERROR_TASK"
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_NON_RETRYABLE_ERROR_WF").
		Version(1).
		Add(workflow.NewSimpleTask(taskName, taskName))
	err := wf.Register(true)
	if err != nil {
		t.Fatal(err)
	}
	err = testdata.TaskRunner.StartWorker(
		taskName,
		testdata.FaultyWorker,
		5,
		testdata.WorkerPollInterval,
	)
	if err != nil {
		t.Fatal(err)
	}
	err = testdata.ValidateWorkflow(wf, 5*time.Second, model.FailedWorkflow)
	if err != nil {
		t.Fatal(err)
	}
}

func LegacyTestWorker(task *model.Task) (interface{}, error) {
	return map[string]interface{}{
		"message":   "processed by legacy worker",
		"taskId":    task.TaskId,
		"timestamp": time.Now().Unix(),
	}, nil
}

func TypedTestWorker(ctx worker.TaskContext, data *TestData) (TestData, error) {
	if data == nil {
		data = &TestData{}
	}

	return TestData{
		InputKey: TestInputKey{
			StatusCode:   data.InputKey.StatusCode + 100,
			ReasonPhrase: "processed by typed worker",
			HostName:     "test-host-processed",
		},
	}, nil
}

func SimpleTestWorker(ctx worker.TaskContext, input map[string]interface{}) (map[string]interface{}, error) {
	name, exists := input["name"]
	if !exists {
		name = "DefaultName"
	}

	return map[string]interface{}{
		"greeting":    "Hello, " + name.(string) + "!",
		"processedBy": "simple typed worker",
		"taskId":      ctx.GetTaskId(),
		"retryCount":  ctx.GetRetryCount(),
	}, nil
}

// Test: Legacy Worker Integration
func TestLegacyWorkerIntegration(t *testing.T) {
	taskName := "TEST_GO_LEGACY_TASK"
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_LEGACY_WORKFLOW").
		Version(1).
		Add(workflow.NewSimpleTask(taskName, taskName).
			Input("testData", "legacy test input"))

	err := wf.Register(true)
	if err != nil {
		t.Fatal(err)
	}

	err = testdata.TaskRunner.StartWorker(
		taskName,
		LegacyTestWorker,
		2,
		testdata.WorkerPollInterval,
	)
	if err != nil {
		t.Fatal(err)
	}

	err = testdata.ValidateWorkflow(wf, 10*time.Second, model.CompletedWorkflow)
	if err != nil {
		t.Fatal(err)
	}
}

// Test: New Typed Worker Integration
func TestTypedWorkerIntegration(t *testing.T) {
	taskName := "TEST_GO_TYPED_TASK"
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_TYPED_WORKFLOW").
		Version(1).
		Add(workflow.NewSimpleTask(taskName, taskName).
			Input("inputKey", map[string]interface{}{
				"statusCode":   200,
				"reasonPhrase": "initial",
				"hostName":     "original-host",
			}))

	err := wf.Register(true)
	if err != nil {
		t.Fatal(err)
	}

	// Start new typed worker
	typedWorker := worker.NewWorkerWithCtx(taskName, TypedTestWorker)
	typedWorker.Options = &worker.TaskWorkerOptions{
		BatchSize:    3,
		PollInterval: testdata.WorkerPollInterval,
	}

	err = typedWorker.Start(testdata.TaskRunner)
	if err != nil {
		t.Fatal(err)
	}

	err = testdata.ValidateWorkflow(wf, 10*time.Second, model.CompletedWorkflow)
	if err != nil {
		t.Fatal(err)
	}
}

// Test: Simple Typed Worker Integration
func TestSimpleTypedWorkerIntegration(t *testing.T) {
	taskName := "TEST_GO_SIMPLE_TYPED_TASK"
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_SIMPLE_TYPED_WORKFLOW").
		Version(1).
		Add(workflow.NewSimpleTask(taskName, taskName).
			Input("name", "TestUser"))

	err := wf.Register(true)
	if err != nil {
		t.Fatal(err)
	}

	// Start simple typed worker
	simpleWorker := worker.NewWorkerWithCtx(taskName, SimpleTestWorker)
	simpleWorker.Options = &worker.TaskWorkerOptions{
		BatchSize:    2,
		PollInterval: testdata.WorkerPollInterval,
	}

	err = simpleWorker.Start(testdata.TaskRunner)
	if err != nil {
		t.Fatal(err)
	}

	err = testdata.ValidateWorkflow(wf, 10*time.Second, model.CompletedWorkflow)
	if err != nil {
		t.Fatal(err)
	}
}

// Test: Multi-Task Workflow with Mixed Worker Types
func TestMultiTaskWorkflowIntegration(t *testing.T) {
	legacyTaskName := "TEST_GO_MULTI_LEGACY_TASK"
	typedTaskName := "TEST_GO_MULTI_TYPED_TASK"
	simpleTaskName := "TEST_GO_MULTI_SIMPLE_TASK"

	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_MULTI_TASK_WORKFLOW").
		Version(1).
		Add(workflow.NewSimpleTask(simpleTaskName, "simple_ref").
			Input("name", "MultiTaskTest")).
		Add(workflow.NewSimpleTask(legacyTaskName, "legacy_ref").
			Input("data", "${simple_ref.output.greeting}")).
		Add(workflow.NewSimpleTask(typedTaskName, "typed_ref").
			Input("inputKey", map[string]interface{}{
				"statusCode": 500,
				"hostName":   "multi-task-host",
			}))

	err := wf.Register(true)
	if err != nil {
		t.Fatal(err)
	}

	// Start legacy worker
	err = testdata.TaskRunner.StartWorker(
		legacyTaskName,
		LegacyTestWorker,
		1,
		testdata.WorkerPollInterval,
	)
	if err != nil {
		t.Fatal(err)
	}

	// Start typed worker
	typedWorker := worker.NewWorkerWithCtx(typedTaskName, TypedTestWorker)
	typedWorker.Options = &worker.TaskWorkerOptions{
		BatchSize:    1,
		PollInterval: testdata.WorkerPollInterval,
	}
	err = typedWorker.Start(testdata.TaskRunner)
	if err != nil {
		t.Fatal(err)
	}

	// Start simple worker
	simpleWorker := worker.NewWorkerWithCtx(simpleTaskName, SimpleTestWorker)
	simpleWorker.Options = &worker.TaskWorkerOptions{
		BatchSize:    1,
		PollInterval: testdata.WorkerPollInterval,
	}
	err = simpleWorker.Start(testdata.TaskRunner)
	if err != nil {
		t.Fatal(err)
	}

	err = testdata.ValidateWorkflow(wf, 15*time.Second, model.CompletedWorkflow)
	if err != nil {
		t.Fatal(err)
	}
}

// Test: Parallel Execution with Fork-Join
func TestParallelExecutionIntegration(t *testing.T) {
	taskName := "TEST_GO_PARALLEL_SIMPLE_TASK"
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_PARALLEL_EXECUTION_WORKFLOW").
		Version(1).
		Add(workflow.NewForkTask("parallel_fork",
			[]workflow.TaskInterface{
				workflow.NewSimpleTask(taskName, "parallel_1").Input("name", "Task1"),
				workflow.NewSimpleTask(taskName, "parallel_2").Input("name", "Task2"),
			},
			[]workflow.TaskInterface{
				workflow.NewSimpleTask(taskName, "parallel_3").Input("name", "Task3"),
			},
		)).
		Add(workflow.NewJoinTask("join_ref", "parallel_1", "parallel_2", "parallel_3"))

	err := wf.Register(true)
	if err != nil {
		t.Fatal(err)
	}

	// Start simple worker for parallel tasks
	simpleWorker := worker.NewWorkerWithCtx(taskName, SimpleTestWorker)
	simpleWorker.Options = &worker.TaskWorkerOptions{
		BatchSize:    5, // Higher batch size for parallel execution
		PollInterval: testdata.WorkerPollInterval,
	}

	err = simpleWorker.Start(testdata.TaskRunner)
	if err != nil {
		t.Fatal(err)
	}

	err = testdata.ValidateWorkflow(wf, 15*time.Second, model.CompletedWorkflow)
	if err != nil {
		t.Fatal(err)
	}
}

// Test: Batch Processing
func TestBatchProcessing(t *testing.T) {
	taskName := "TEST_GO_BATCH_TASK"
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_BATCH_WORKFLOW").
		Version(1).
		Add(workflow.NewSimpleTask(taskName, "batch_1").Input("id", 1)).
		Add(workflow.NewSimpleTask(taskName, "batch_2").Input("id", 2)).
		Add(workflow.NewSimpleTask(taskName, "batch_3").Input("id", 3)).
		Add(workflow.NewSimpleTask(taskName, "batch_4").Input("id", 4)).
		Add(workflow.NewSimpleTask(taskName, "batch_5").Input("id", 5))

	err := wf.Register(true)
	if err != nil {
		t.Fatal(err)
	}

	// Start worker with higher batch size
	err = testdata.TaskRunner.StartWorker(
		taskName,
		LegacyTestWorker,
		10,
		testdata.WorkerPollInterval,
	)
	if err != nil {
		t.Fatal(err)
	}

	err = testdata.ValidateWorkflow(wf, 15*time.Second, model.CompletedWorkflow)
	if err != nil {
		t.Fatal(err)
	}
}
