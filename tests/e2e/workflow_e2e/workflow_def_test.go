package workflow_e2e

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/pkg/model"

	"github.com/conductor-sdk/conductor-go/examples"
	"github.com/conductor-sdk/conductor-go/pkg/model/enum/workflow_status"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/def/workflow"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/executor"
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
)

const (
	workflowValidationTimeout = 3 * time.Second
	workerQty                 = 7
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
	err := validateWorkflow(httpTaskWorkflow, workflowValidationTimeout)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSimpleTask(t *testing.T) {
	response, err := registerTask(simpleTask)
	if err != nil {
		t.Fatal("Failed to register task, response: ", response)
	}
	simpleTaskWorkflow := workflow.NewConductorWorkflow(e2e_properties.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_SIMPLE").
		Version(1).
		Add(simpleTask)
	err = e2e_properties.TaskRunner.StartWorker(
		simpleTask.ReferenceName(),
		examples.SimpleWorker,
		workerQty,
		500*time.Millisecond,
	)
	if err != nil {
		t.Fatal(err)
	}
	err = validateWorkflow(simpleTaskWorkflow, workflowValidationTimeout)
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
	err := validateWorkflow(inlineTaskWorkflow, workflowValidationTimeout)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSqsEventTask(t *testing.T) {
	sqsEventTask := workflow.NewSqsEventTask(
		"TEST_GO_TASK_EVENT_SQS",
		"QUEUE",
	)
	testEventTask(
		t,
		"TEST_GO_WORKFLOW_EVENT_SQS",
		sqsEventTask,
	)
}

func TestConductorEventTask(t *testing.T) {
	sqsEventTask := workflow.NewSqsEventTask(
		"TEST_GO_TASK_EVENT_CONDUCTOR",
		"QUEUE",
	)
	testEventTask(
		t,
		"TEST_GO_WORKFLOW_EVENT_CONDUCTOR",
		sqsEventTask,
	)
}

func TestKafkaPublishTask(t *testing.T) {
	kafkaPublishTask := workflow.NewKafkaPublishTask(
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
	kafkaPublishTaskWorkflow := workflow.NewConductorWorkflow(e2e_properties.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_KAFKA_PUBLISH").
		Version(1).
		Add(kafkaPublishTask)
	registerWorkflow(t, kafkaPublishTaskWorkflow)
}

func TestDoWhileTask(t *testing.T) {

}

func TestTerminateTask(t *testing.T) {
	terminateTaskWorkflow := workflow.NewConductorWorkflow(e2e_properties.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_TERMINATE").
		Version(1).
		Add(terminateTask)
	registerWorkflow(t, terminateTaskWorkflow)
}

func TestSwitchTask(t *testing.T) {
	switchTaskWorkflow := workflow.NewConductorWorkflow(e2e_properties.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_SWITCH").
		Version(1).
		Add(switchTask)
	registerWorkflow(t, switchTaskWorkflow)
}

func testEventTask(t *testing.T, workflowName string, event *workflow.EventTask) {
	eventTaskWorkflow := workflow.NewConductorWorkflow(e2e_properties.WorkflowExecutor).
		Name(workflowName).
		Version(1).
		Add(event)
	_, err := eventTaskWorkflow.Register(true)
	if err != nil {
		t.Fatal(err)
	}
}

func isWorkflowCompleted(workflow *model.Workflow) bool {
	return workflow.Status == string(workflow_status.COMPLETED)
}

func validateWorkflow(conductorWorkflow *workflow.ConductorWorkflow, timeout time.Duration) error {
	_, err := conductorWorkflow.Register(true)
	if err != nil {
		return err
	}
	version := conductorWorkflow.GetVersion()
	_, workflowExecutionChannel, err := conductorWorkflow.ExecuteWorkflow(
		model.NewStartWorkflowRequest(conductorWorkflow.GetName(), &version, "", map[string]interface{}{}),
	)
	if err != nil {
		return err
	}
	workflow, err := executor.WaitForWorkflowCompletionUntilTimeout(
		workflowExecutionChannel,
		timeout,
	)
	if err != nil {
		return err
	}
	if !isWorkflowCompleted(workflow) {
		return fmt.Errorf("workflow finished with status: %s", workflow.Status)
	}
	return nil
}

func registerTask(task *workflow.SimpleTask) (*http.Response, error) {
	return e2e_properties.MetadataClient.RegisterTaskDef(
		context.Background(),
		[]model.TaskDef{
			*task.ToTaskDef(),
		},
	)
}

func registerWorkflow(t *testing.T, conductorWorkflow *workflow.ConductorWorkflow) {
	_, err := conductorWorkflow.Register(true)
	if err != nil {
		t.Error(err)
	}
}
