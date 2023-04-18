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
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/internal/testdata"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

const (
	workflowValidationTimeout = 7 * time.Second
	workflowBulkQty           = 10
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

var (
	httpTask = workflow.NewHttpTask(
		"TEST_GO_TASK_HTTP",
		&workflow.HttpInput{
			Uri: "https://orkes-api-tester.orkesconductor.com/get",
		},
	)

	simpleTask = workflow.NewSimpleTask(
		"TEST_GO_TASK_SIMPLE", "TEST_GO_TASK_SIMPLE",
	)

	terminateTask = workflow.NewTerminateTask(
		"TEST_GO_TASK_TERMINATE",
		model.FailedWorkflow,
		"Task used to mark workflow as failed",
	)

	switchTask = workflow.NewSwitchTask(
		"TEST_GO_TASK_SWITCH",
		"switchCaseValue",
	).
		Input("switchCaseValue", "${workflow.input.service}").
		UseJavascript(true).
		SwitchCase(
			"REQUEST",
			httpTask,
		).
		SwitchCase(
			"STOP",
			terminateTask,
		)

	inlineTask = workflow.NewInlineTask(
		"TEST_GO_TASK_INLINE",
		"function e() { if ($.value == 1){return {\"result\": true}} else { return {\"result\": false}}} e();",
	)

	kafkaPublishTask = workflow.NewKafkaPublishTask(
		"TEST_GO_TASK_KAFKA_PUBLISH",
		&workflow.KafkaPublishTaskInput{
			Topic:            "userTopic",
			Value:            "Message to publish",
			BootStrapServers: "localhost:9092",
			Headers: map[string]interface{}{
				"x-Auth": "Auth-key",
			},
			Key:           "123",
			KeySerializer: "org.apache.kafka.common.serialization.IntegerSerializer",
		},
	)

	sqsEventTask = workflow.NewSqsEventTask(
		"TEST_GO_TASK_EVENT_SQS",
		"QUEUE",
	)

	conductorEventTask = workflow.NewConductorEventTask(
		"TEST_GO_TASK_EVENT_CONDUCTOR",
		"EVENT_NAME",
	)
)

func TestHttpTask(t *testing.T) {
	httpTaskWorkflow := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_HTTP").
		OwnerEmail("test@orkes.io").
		Version(1).
		Add(httpTask)
	err := testdata.ValidateWorkflow(httpTaskWorkflow, workflowValidationTimeout, model.CompletedWorkflow)
	if err != nil {
		t.Fatal(err)
	}
	err = testdata.ValidateWorkflowBulk(httpTaskWorkflow, workflowValidationTimeout, workflowBulkQty)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSimpleTask(t *testing.T) {
	err := testdata.ValidateTaskRegistration(*simpleTask.ToTaskDef())
	if err != nil {
		t.Fatal(err)
	}
	simpleTaskWorkflow := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_SIMPLE").
		Version(1).
		Add(simpleTask)
	err = testdata.TaskRunner.StartWorker(
		simpleTask.ReferenceName(),
		testdata.SimpleWorker,
		testdata.WorkerQty,
		testdata.WorkerPollInterval,
	)
	if err != nil {
		t.Fatal(err)
	}
	err = testdata.ValidateWorkflow(simpleTaskWorkflow, workflowValidationTimeout, model.CompletedWorkflow)
	if err != nil {
		t.Fatal(err)
	}
	err = testdata.ValidateWorkflowBulk(simpleTaskWorkflow, workflowValidationTimeout, workflowBulkQty)
	if err != nil {
		t.Fatal(err)
	}
	err = testdata.TaskRunner.DecreaseBatchSize(
		simpleTask.ReferenceName(),
		testdata.WorkerQty,
	)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSimpleTaskWithoutRetryCount(t *testing.T) {
	taskToRegister := simpleTask.ToTaskDef()
	taskToRegister.RetryCount = 0
	err := testdata.ValidateTaskRegistration(*taskToRegister)
	if err != nil {
		t.Fatal(err)
	}
	simpleTaskWorkflow := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_SIMPLE").
		Version(1).
		Add(simpleTask)
	err = testdata.TaskRunner.StartWorker(
		simpleTask.ReferenceName(),
		testdata.SimpleWorker,
		testdata.WorkerQty,
		testdata.WorkerPollInterval,
	)
	if err != nil {
		t.Fatal(err)
	}
	err = testdata.ValidateWorkflow(simpleTaskWorkflow, workflowValidationTimeout, model.CompletedWorkflow)
	if err != nil {
		t.Fatal(err)
	}
	err = testdata.ValidateWorkflowBulk(simpleTaskWorkflow, workflowValidationTimeout, workflowBulkQty)
	if err != nil {
		t.Fatal(err)
	}
	err = testdata.TaskRunner.DecreaseBatchSize(
		simpleTask.ReferenceName(),
		testdata.WorkerQty,
	)
	if err != nil {
		t.Fatal(err)
	}
}

func TestInlineTask(t *testing.T) {
	inlineTaskWorkflow := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_INLINE_TASK").
		Version(1).
		Add(inlineTask)
	err := testdata.ValidateWorkflow(inlineTaskWorkflow, workflowValidationTimeout, model.CompletedWorkflow)
	if err != nil {
		t.Fatal(err)
	}
	err = testdata.ValidateWorkflowBulk(inlineTaskWorkflow, workflowValidationTimeout, workflowBulkQty)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSqsEventTask(t *testing.T) {
	workflow := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_EVENT_SQS").
		Version(1).
		Add(sqsEventTask)
	err := testdata.ValidateWorkflowRegistration(workflow)
	if err != nil {
		t.Fatal(err)
	}
}

func TestConductorEventTask(t *testing.T) {
	workflow := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_EVENT_CONDUCTOR").
		Version(1).
		Add(conductorEventTask)
	err := testdata.ValidateWorkflowRegistration(workflow)
	if err != nil {
		t.Fatal(err)
	}
}

func TestKafkaPublishTask(t *testing.T) {
	workflow := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_KAFKA_PUBLISH").
		Version(1).
		Add(kafkaPublishTask)
	err := testdata.ValidateWorkflowRegistration(workflow)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDoWhileTask(t *testing.T) {

}

func TestTerminateTask(t *testing.T) {
	workflow := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_TERMINATE").
		Version(1).
		Add(terminateTask)
	err := testdata.ValidateWorkflowRegistration(workflow)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSwitchTask(t *testing.T) {
	workflow := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_SWITCH").
		Version(1).
		Add(switchTask)
	err := testdata.ValidateWorkflowRegistration(workflow)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDynamicForkWorkflow(t *testing.T) {
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("dynamic_workflow_array_sub_workflow").
		Version(1).
		Add(createDynamicForkTask())
	err := wf.Register(true)
	if err != nil {
		t.Fatal()
	}
}

func createDynamicForkTask() *workflow.DynamicForkTask {
	return workflow.NewDynamicForkTaskWithoutPrepareTask(
		"dynamic_workflow_array_sub_workflow",
	).Input(
		"forkTaskWorkflow", "extract_user",
	).Input(
		"forkTaskInputs", []map[string]interface{}{
			{
				"input": "value1",
			},
			{
				"sub_workflow_2_inputs": map[string]interface{}{
					"key":  "value",
					"key2": 23,
				},
			},
		},
	)
}

func TestComplexSwitchWorkflow(t *testing.T) {
	wf := testdata.GetWorkflowWithComplexSwitchTask()
	err := testdata.ValidateWorkflowRegistration(wf)
	if err != nil {
		t.Fatal(err)
	}
	receivedWf, _, err := testdata.MetadataClient.Get(context.Background(), wf.GetName(), nil)
	if err != nil {
		t.Fatal(err)
	}
	counter := countMultipleSwitchInnerTasks(receivedWf.Tasks...)
	assert.Equal(t, 7, counter)
}

func TestWorkflowParameterIO(t *testing.T) {
	amount := 10
	taskName := "workflow_parameter_io_task_go_sdk"

	tasks := make([]*workflow.SimpleTask, amount)
	tasks[0] = workflow.NewSimpleTask(taskName, fmt.Sprintf("%s_%d", taskName, 0)).
		Input("X", "${workflow.input.A}").
		Input("Y", "${workflow.input.B}")
	tasks[1] = workflow.NewSimpleTask(taskName, fmt.Sprintf("%s_%d", taskName, 1)).
		Input("X", "${workflow.input.B}").
		Input("Y", fmt.Sprintf("${%s.output.result}", tasks[0].ReferenceName()))
	for idx := 2; idx < amount; idx += 1 {
		tasks[idx] = workflow.NewSimpleTask(taskName, fmt.Sprintf("%s_%d", taskName, idx)).
			Input("X", fmt.Sprintf("${%s.output.result}", tasks[idx-2].ReferenceName())).
			Input("Y", fmt.Sprintf("${%s.output.result}", tasks[idx-1].ReferenceName()))
	}

	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("GoSdkWorkflowParameterIO").
		OwnerEmail("test@orkes.io").
		Version(1).
		InputParameters("A", "B")

	for _, task := range tasks {
		testdata.ValidateTaskRegistration(*task.ToTaskDef())
		wf.Add(task)
	}
	wf.Register(true)

	testdata.TaskRunner.StartWorker(taskName, testdata.FibonacciWorker, 10, 100*time.Millisecond)

	err := testdata.ValidateWorkflowWithInput(
		wf,
		map[string]interface{}{
			"A": 0,
			"B": 1,
		},
		10*time.Second,
		model.CompletedWorkflow,
	)
	if err != nil {
		t.Fatal(err)
	}
}

func countMultipleSwitchInnerTasks(tasks ...model.WorkflowTask) int {
	counter := 0
	for _, task := range tasks {
		counter += countSwitchInnerTasks(task)
	}
	return counter
}

func countSwitchInnerTasks(task model.WorkflowTask) int {
	fmt.Println(task.Type_)
	counter := 1
	if task.Type_ != "SWITCH" {
		return counter
	}
	for _, value := range task.DecisionCases {
		counter += countMultipleSwitchInnerTasks(value...)
	}
	return counter
}
