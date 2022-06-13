//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package workflow_e2e

import (
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/definition"
	"os"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/examples"
	"github.com/conductor-sdk/conductor-go/tests/e2e/e2e_properties"
	log "github.com/sirupsen/logrus"
)

var (
	httpTask = definition.NewHttpTask(
		"TEST_GO_TASK_HTTP",
		&definition.HttpInput{
			Uri: "https://catfact.ninja/fact",
		},
	)

	simpleTask = definition.NewSimpleTask(
		"TEST_GO_TASK_SIMPLE", "TEST_GO_TASK_SIMPLE",
	)

	terminateTask = definition.NewTerminateTask(
		"TEST_GO_TASK_TERMINATE",
		model.FailedWorkflow,
		"Task used to mark workflow as failed",
	)

	switchTask = definition.NewSwitchTask(
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

	inlineTask = definition.NewInlineTask(
		"TEST_GO_TASK_INLINE",
		"function e() { if ($.value == 1){return {\"result\": true}} else { return {\"result\": false}}} e();",
	)

	kafkaPublishTask = definition.NewKafkaPublishTask(
		"TEST_GO_TASK_KAFKA_PUBLISH",
		&definition.KafkaPublishTaskInput{
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

	sqsEventTask = definition.NewSqsEventTask(
		"TEST_GO_TASK_EVENT_SQS",
		"QUEUE",
	)

	conductorEventTask = definition.NewConductorEventTask(
		"TEST_GO_TASK_EVENT_CONDUCTOR",
		"EVENT_NAME",
	)
)

const (
	workflowValidationTimeout = 5 * time.Second
	workflowBulkQty           = 10

	workerQty          = 3
	workerPollInterval = 500 * time.Millisecond
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func TestHttpTask(t *testing.T) {
	httpTaskWorkflow := definition.NewConductorWorkflow(e2e_properties.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_HTTP").
		Version(1).
		Add(httpTask)
	err := e2e_properties.ValidateWorkflow(httpTaskWorkflow, workflowValidationTimeout)
	if err != nil {
		t.Fatal(err)
	}
	err = e2e_properties.ValidateWorkflowBulk(httpTaskWorkflow, workflowValidationTimeout, workflowBulkQty)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSimpleTask(t *testing.T) {
	err := e2e_properties.ValidateTaskRegistration(*simpleTask.ToTaskDef())
	if err != nil {
		t.Fatal(err)
	}
	simpleTaskWorkflow := definition.NewConductorWorkflow(e2e_properties.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_SIMPLE").
		Version(1).
		Add(simpleTask)
	err = e2e_properties.TaskRunner.StartWorker(
		simpleTask.ReferenceName(),
		examples.SimpleWorker,
		workerQty,
		workerPollInterval,
	)
	if err != nil {
		t.Fatal(err)
	}
	err = e2e_properties.ValidateWorkflow(simpleTaskWorkflow, workflowValidationTimeout)
	if err != nil {
		t.Fatal(err)
	}
	err = e2e_properties.ValidateWorkflowBulk(simpleTaskWorkflow, workflowValidationTimeout, workflowBulkQty)
	if err != nil {
		t.Fatal(err)
	}
	err = e2e_properties.TaskRunner.RemoveWorker(
		simpleTask.ReferenceName(),
		workerQty,
	)
	if err != nil {
		t.Fatal(err)
	}
}

func TestInlineTask(t *testing.T) {
	inlineTaskWorkflow := definition.NewConductorWorkflow(e2e_properties.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_INLINE_TASK").
		Version(1).
		Add(inlineTask)
	err := e2e_properties.ValidateWorkflow(inlineTaskWorkflow, workflowValidationTimeout)
	if err != nil {
		t.Fatal(err)
	}
	err = e2e_properties.ValidateWorkflowBulk(inlineTaskWorkflow, workflowValidationTimeout, workflowBulkQty)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSqsEventTask(t *testing.T) {
	workflow := definition.NewConductorWorkflow(e2e_properties.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_EVENT_SQS").
		Version(1).
		Add(sqsEventTask)
	err := e2e_properties.ValidateWorkflowRegistration(workflow)
	if err != nil {
		t.Fatal(err)
	}
}

func TestConductorEventTask(t *testing.T) {
	workflow := definition.NewConductorWorkflow(e2e_properties.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_EVENT_CONDUCTOR").
		Version(1).
		Add(conductorEventTask)
	err := e2e_properties.ValidateWorkflowRegistration(workflow)
	if err != nil {
		t.Fatal(err)
	}
}

func TestKafkaPublishTask(t *testing.T) {
	workflow := definition.NewConductorWorkflow(e2e_properties.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_KAFKA_PUBLISH").
		Version(1).
		Add(kafkaPublishTask)
	err := e2e_properties.ValidateWorkflowRegistration(workflow)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDoWhileTask(t *testing.T) {

}

func TestTerminateTask(t *testing.T) {
	workflow := definition.NewConductorWorkflow(e2e_properties.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_TERMINATE").
		Version(1).
		Add(terminateTask)
	err := e2e_properties.ValidateWorkflowRegistration(workflow)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSwitchTask(t *testing.T) {
	workflow := definition.NewConductorWorkflow(e2e_properties.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_SWITCH").
		Version(1).
		Add(switchTask)
	err := e2e_properties.ValidateWorkflowRegistration(workflow)
	if err != nil {
		t.Fatal(err)
	}
}
