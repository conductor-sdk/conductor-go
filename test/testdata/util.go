//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package testdata

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/authentication"
	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/log"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
	"github.com/conductor-sdk/conductor-go/sdk/worker"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
)

const (
	AUTHENTICATION_KEY_ID     = "KEY"
	AUTHENTICATION_KEY_SECRET = "SECRET"
	BASE_URL                  = "CONDUCTOR_SERVER_URL"
)

var (
	apiClient = getApiClientWithAuthentication()
)

var (
	MetadataClient = client.MetadataResourceApiService{
		APIClient: apiClient,
	}
	TaskClient = client.TaskResourceApiService{
		APIClient: apiClient,
	}
	WorkflowClient = client.WorkflowResourceApiService{
		APIClient: apiClient,
	}
	EventClient = client.EventResourceApiService{
		APIClient: apiClient,
	}
	TagsClient = client.TagsApiService{
		APIClient: apiClient,
	}
	ApplicationClient = client.NewApplicationClient(apiClient)
	EnvironmentClient = client.NewEnvironmentClient(apiClient)
	IntegrationClient = client.NewIntegrationClient(apiClient)
	PromptClient      = client.NewPromptClient(apiClient)
	UserClient        = client.NewUserClient(apiClient)
)

var TaskRunner = worker.NewTaskRunnerWithApiClient(apiClient)

var WorkflowExecutor = executor.NewWorkflowExecutor(apiClient)

func ValidateWorkflowDaemon(waitTime time.Duration, outputChannel chan error, workflowId string, expectedOutput map[string]interface{}, expectedStatus model.WorkflowStatus) {
	time.Sleep(waitTime)
	workflow, _, err := WorkflowClient.GetExecutionStatus(
		context.Background(),
		workflowId,
		nil,
	)
	if err != nil {
		outputChannel <- err
		return
	}
	if !isWorkflowCompleted(&workflow, expectedStatus) {
		outputChannel <- fmt.Errorf(
			"workflow status different than expected, workflowId: %s, workflowStatus: %s",
			workflow.WorkflowId, workflow.Status,
		)
		return
	}
	if !reflect.DeepEqual(workflow.Output, expectedOutput) {
		outputChannel <- fmt.Errorf(
			"workflow output is different than expected, workflowId: %s, output: %+v",
			workflow.WorkflowId, workflow.Output,
		)
		return
	}
	outputChannel <- nil
}

func getApiClientWithAuthentication() *client.APIClient {
	return client.NewAPIClientWithTokenExpiration(
		getAuthenticationSettings(),
		getHttpSettingsWithAuth(),
		authentication.NewTokenExpiration(
			3*time.Second,
			30*time.Second,
		),
	)
}

func getAuthenticationSettings() *settings.AuthenticationSettings {
	return settings.NewAuthenticationSettings(
		os.Getenv(AUTHENTICATION_KEY_ID),
		os.Getenv(AUTHENTICATION_KEY_SECRET),
	)
}

func getHttpSettingsWithAuth() *settings.HttpSettings {
	return settings.NewHttpSettings(
		os.Getenv(BASE_URL),
	)
}

func StartWorkflows(workflowQty int, workflowName string) ([]string, error) {
	workflowIdList := make([]string, workflowQty)
	for i := 0; i < workflowQty; i += 1 {
		workflowId, _, err := WorkflowClient.StartWorkflow(
			context.Background(),
			make(map[string]interface{}),
			workflowName,
			nil,
		)
		if err != nil {
			return nil, err
		}
		log.Debug(
			"Started workflow",
			", workflowName: ", workflowName,
			", workflowId: ", workflowId,
		)
		workflowIdList[i] = workflowId
	}
	return workflowIdList, nil
}

func ValidateWorkflow(conductorWorkflow *workflow.ConductorWorkflow, timeout time.Duration, expectedStatus model.WorkflowStatus) error {
	err := ValidateWorkflowRegistration(conductorWorkflow)
	if err != nil {
		return err
	}
	workflowId, err := conductorWorkflow.StartWorkflowWithInput(make(map[string]interface{}))
	if err != nil {
		return err
	}
	log.Debug("Started workflowId: ", workflowId)
	workflowExecutionChannel, err := WorkflowExecutor.MonitorExecution(workflowId)
	if err != nil {
		return err
	}
	log.Debug("Generated workflowExecutionChannel for workflowId: ", workflowId)
	workflow, err := executor.WaitForWorkflowCompletionUntilTimeout(
		workflowExecutionChannel,
		timeout,
	)
	if err != nil {
		return err
	}
	log.Debug("Workflow completed, workflowId: ", workflowId)
	if !isWorkflowCompleted(workflow, expectedStatus) {
		return fmt.Errorf("workflow finished with unexpected status: %s", workflow.Status)
	}
	return nil
}

func ValidateWorkflowBulk(conductorWorkflow *workflow.ConductorWorkflow, timeout time.Duration, amount int) error {
	err := ValidateWorkflowRegistration(conductorWorkflow)
	if err != nil {
		return err
	}
	version := conductorWorkflow.GetVersion()
	startWorkflowRequests := make([]*model.StartWorkflowRequest, amount)
	for i := 0; i < amount; i += 1 {
		startWorkflowRequests[i] = model.NewStartWorkflowRequest(
			conductorWorkflow.GetName(),
			version,
			"",
			make(map[string]interface{}),
		)
	}
	runningWorkflows := WorkflowExecutor.StartWorkflows(true, startWorkflowRequests...)
	WorkflowExecutor.WaitForRunningWorkflowsUntilTimeout(timeout, runningWorkflows...)
	for _, runningWorkflow := range runningWorkflows {
		if runningWorkflow.Err != nil {
			return err
		}
		if runningWorkflow.CompletedWorkflow == nil {
			return fmt.Errorf("invalid completed workflows")
		}
		if !isWorkflowCompleted(runningWorkflow.CompletedWorkflow, model.CompletedWorkflow) {
			return fmt.Errorf("workflow finished with status: %s", runningWorkflow.CompletedWorkflow.Status)
		}
	}
	return nil
}

func ValidateTaskRegistration(taskDefs ...model.TaskDef) error {
	response, err := MetadataClient.RegisterTaskDef(
		context.Background(),
		taskDefs,
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

func ValidateWorkflowRegistration(workflow *workflow.ConductorWorkflow) error {
	for attempt := 0; attempt < 5; attempt += 1 {
		err := workflow.Register(true)
		if err != nil {
			time.Sleep(time.Duration(attempt+2) * time.Second)
			fmt.Println("Failed to validate workflow registration, reason: " + err.Error())
			continue
		}
		return nil
	}
	return fmt.Errorf("exhausted retries")
}

func ValidateWorkflowDeletion(workflow *workflow.ConductorWorkflow) error {
	for attempt := 0; attempt < 5; attempt += 1 {
		err := workflow.UnRegister()
		if err != nil {
			time.Sleep(time.Duration(attempt+2) * time.Second)
			fmt.Println("Failed to validate workflow deletion, reason: " + err.Error())
			continue
		}
		return nil
	}
	return fmt.Errorf("exhausted retries")
}

func isWorkflowCompleted(workflow *model.Workflow, expectedStatus model.WorkflowStatus) bool {
	return workflow.Status == expectedStatus
}

// Common Test Tasks
const (
	WorkflowValidationTimeout = 7 * time.Second
	WorkflowBulkQty           = 10
)

var (
	TestHttpTask = workflow.NewHttpTask(
		"TEST_GO_TASK_HTTP",
		&workflow.HttpInput{
			Uri: "https://orkes-api-tester.orkesconductor.com/get",
		},
	)

	TestSimpleTask = workflow.NewSimpleTask(
		"TEST_GO_TASK_SIMPLE", "TEST_GO_TASK_SIMPLE",
	)

	TestTerminateTask = workflow.NewTerminateTask(
		"TEST_GO_TASK_TERMINATE",
		model.FailedWorkflow,
		"Task used to mark workflow as failed",
	)

	TestSwitchTask = workflow.NewSwitchTask(
		"TEST_GO_TASK_SWITCH",
		"switchCaseValue",
	).
		Input("switchCaseValue", "${workflow.input.service}").
		UseJavascript(true).
		SwitchCase(
			"REQUEST",
			TestHttpTask,
		).
		SwitchCase(
			"STOP",
			TestTerminateTask,
		)

	TestInlineTask = workflow.NewInlineTask(
		"TEST_GO_TASK_INLINE",
		"function e() { if ($.value == 1){return {\"result\": true}} else { return {\"result\": false}}} e();",
	)

	TestKafkaPublishTask = workflow.NewKafkaPublishTask(
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

	TestSqsEventTask = workflow.NewSqsEventTask(
		"TEST_GO_TASK_EVENT_SQS",
		"QUEUE",
	)

	TestConductorEventTask = workflow.NewConductorEventTask(
		"TEST_GO_TASK_EVENT_CONDUCTOR",
		"EVENT_NAME",
	)
)
