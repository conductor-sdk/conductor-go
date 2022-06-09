package workflow_e2e

import (
	"context"
	"fmt"
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
	workflowValidationTimeout = 10 * time.Second
	workflowBulkQty           = 22

	workerQty          = 7
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
	err := validateWorkflow(httpTaskWorkflow, workflowValidationTimeout)
	if err != nil {
		t.Fatal(err)
	}
	err = validateWorkflowBulk(httpTaskWorkflow, workflowValidationTimeout, workflowBulkQty)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSimpleTask(t *testing.T) {
	err := validateTaskRegistration(simpleTask)
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
	err = validateWorkflow(simpleTaskWorkflow, workflowValidationTimeout)
	if err != nil {
		t.Fatal(err)
	}
	err = validateWorkflowBulk(simpleTaskWorkflow, workflowValidationTimeout, workflowBulkQty)
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
	err = validateWorkflowBulk(inlineTaskWorkflow, workflowValidationTimeout, workflowBulkQty)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSqsEventTask(t *testing.T) {
	workflow := workflow.NewConductorWorkflow(e2e_properties.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_EVENT_SQS").
		Version(1).
		Add(sqsEventTask)
	err := validateWorkflowRegistration(workflow)
	if err != nil {
		t.Fatal(err)
	}
}

func TestConductorEventTask(t *testing.T) {
	workflow := workflow.NewConductorWorkflow(e2e_properties.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_EVENT_CONDUCTOR").
		Version(1).
		Add(conductorEventTask)
	err := validateWorkflowRegistration(workflow)
	if err != nil {
		t.Fatal(err)
	}
}

func TestKafkaPublishTask(t *testing.T) {
	workflow := workflow.NewConductorWorkflow(e2e_properties.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_KAFKA_PUBLISH").
		Version(1).
		Add(kafkaPublishTask)
	err := validateWorkflowRegistration(workflow)
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
	err := validateWorkflowRegistration(workflow)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSwitchTask(t *testing.T) {
	workflow := workflow.NewConductorWorkflow(e2e_properties.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_SWITCH").
		Version(1).
		Add(switchTask)
	err := validateWorkflowRegistration(workflow)
	if err != nil {
		t.Fatal(err)
	}
}

func isWorkflowCompleted(workflow *model.Workflow) bool {
	return workflow.Status == string(workflow_status.COMPLETED)
}

func validateWorkflow(conductorWorkflow *workflow.ConductorWorkflow, timeout time.Duration) error {
	err := validateWorkflowRegistration(conductorWorkflow)
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
		return fmt.Errorf("workflow finished with unexpected status: %s", workflow.Status)
	}
	return nil
}

func validateWorkflowBulk(conductorWorkflow *workflow.ConductorWorkflow, timeout time.Duration, amount int) error {
	err := validateWorkflowRegistration(conductorWorkflow)
	if err != nil {
		return err
	}
	version := conductorWorkflow.GetVersion()
	startWorkflowRequests := make([]model.StartWorkflowRequest, amount)
	for i := 0; i < amount; i += 1 {
		startWorkflowRequests[i] = *model.NewStartWorkflowRequest(
			conductorWorkflow.GetName(), &version, "", map[string]interface{}{},
		)
	}
	workflowExecutionChannels, err := conductorWorkflow.ExecuteWorkflowBulk(
		startWorkflowRequests...,
	)
	if err != nil {
		return err
	}
	for _, workflowExecutionChannel := range workflowExecutionChannels {
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
	}
	return nil
}

func validateTaskRegistration(task *workflow.SimpleTask) error {
	response, err := e2e_properties.MetadataClient.RegisterTaskDef(
		context.Background(),
		[]model.TaskDef{
			*task.ToTaskDef(),
		},
	)
	if err != nil {
		log.Debug(
			"Failed to validate task registration. Reason: ", err.Error(),
			", response: ", *response,
		)
		return err
	}
	return nil
}

func validateWorkflowRegistration(workflow *workflow.ConductorWorkflow) error {
	response, err := workflow.Register(true)
	if err != nil {
		log.Debug(
			"Failed to validate workflow registration. Reason: ", err.Error(),
			", response: ", *response,
		)
		return err
	}
	return nil
}
