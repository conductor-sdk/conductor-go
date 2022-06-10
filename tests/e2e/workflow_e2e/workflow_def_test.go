package workflow_e2e

import (
	"os"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/examples"
	"github.com/conductor-sdk/conductor-go/pkg/model/enum/workflow_status"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/def/workflow"
	"github.com/conductor-sdk/conductor-go/tests/e2e/e2e_properties"
	log "github.com/sirupsen/logrus"
)

var (
	httpTask = workflow.NewHttpTask(
		"TEST_GO_TASK_HTTP",
		&workflow.HttpInput{
			Uri: "https://catfact.ninja/fact",
		},
	)

	simpleTask = workflow.NewSimpleTask(
		"TEST_GO_TASK_SIMPLE",
	)

	terminateTask = workflow.NewTerminateTask(
		"TEST_GO_TASK_TERMINATE",
		workflow_status.FAILED,
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
	httpTaskWorkflow := workflow.NewConductorWorkflow(e2e_properties.WorkflowExecutor).
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
	simpleTaskWorkflow := workflow.NewConductorWorkflow(e2e_properties.WorkflowExecutor).
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
	inlineTaskWorkflow := workflow.NewConductorWorkflow(e2e_properties.WorkflowExecutor).
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
	workflow := workflow.NewConductorWorkflow(e2e_properties.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_EVENT_SQS").
		Version(1).
		Add(sqsEventTask)
	err := e2e_properties.ValidateWorkflowRegistration(workflow)
	if err != nil {
		t.Fatal(err)
	}
}

func TestConductorEventTask(t *testing.T) {
	workflow := workflow.NewConductorWorkflow(e2e_properties.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_EVENT_CONDUCTOR").
		Version(1).
		Add(conductorEventTask)
	err := e2e_properties.ValidateWorkflowRegistration(workflow)
	if err != nil {
		t.Fatal(err)
	}
}

func TestKafkaPublishTask(t *testing.T) {
	workflow := workflow.NewConductorWorkflow(e2e_properties.WorkflowExecutor).
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
	workflow := workflow.NewConductorWorkflow(e2e_properties.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_TERMINATE").
		Version(1).
		Add(terminateTask)
	err := e2e_properties.ValidateWorkflowRegistration(workflow)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSwitchTask(t *testing.T) {
	workflow := workflow.NewConductorWorkflow(e2e_properties.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_SWITCH").
		Version(1).
		Add(switchTask)
	err := e2e_properties.ValidateWorkflowRegistration(workflow)
	if err != nil {
		t.Fatal(err)
	}
}
