//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package e2e_properties

import (
	"context"
	"fmt"
	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
	"github.com/conductor-sdk/conductor-go/sdk/worker"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/definition"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
	"os"
	"reflect"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	AUTHENTICATION_KEY_ID     = "KEY"
	AUTHENTICATION_KEY_SECRET = "SECRET"
	BASE_URL                  = "https://pg-staging.orkesconductor.com/api"
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
)

var TaskRunner = worker.NewTaskRunnerWithApiClient(apiClient)

var WorkflowExecutor = executor.NewWorkflowExecutor(apiClient)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

type TreasureChest struct {
	ImportantValue string `json:"importantValue"`
}

func ValidateWorkflowDaemon(waitTime time.Duration, outputChannel chan error, workflowId string, expectedOutput map[string]interface{}) {
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
	if workflow.Status != model.CompletedWorkflow {
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
	return client.NewAPIClient(
		getAuthenticationSettings(),
		getHttpSettingsWithAuth(),
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
		BASE_URL,
	)
}

var (
	TREASURE_CHEST_WORKFLOW_NAME = "treasure_chest_workflow"
	TREASURE_CHEST_TASK_NAME     = "treasure_chest_task"

	TREASURE_WORKFLOW_DEFINITION = model.WorkflowDef{
		UpdateTime:  1650595431465,
		Name:        TREASURE_CHEST_WORKFLOW_NAME,
		Description: "What's inside the treasure chest?",
		Version:     1,
		Tasks: []model.WorkflowTask{
			{
				Name:              TREASURE_CHEST_TASK_NAME,
				TaskReferenceName: TREASURE_CHEST_TASK_NAME,
				Type_:             "SIMPLE",
				StartDelay:        0,
				Optional:          false,
				AsyncComplete:     false,
				InputParameters: map[string]interface{}{
					"importantValue": "${workflow.input.importantValue}",
				},
			},
		},
		InputParameters: []string{"importantValue"},
		OutputParameters: map[string]interface{}{
			"workerOutput": fmt.Sprintf("${%s.output}", TREASURE_CHEST_TASK_NAME),
		},
		SchemaVersion:                 2,
		Restartable:                   true,
		WorkflowStatusListenerEnabled: false,
		OwnerEmail:                    "gustavo.gardusi@orkes.io",
		TimeoutPolicy:                 "ALERT_ONLY",
		TimeoutSeconds:                0,
	}

	TREASURE_TASK_DEFINITION = model.TaskDef{
		Name:                        TREASURE_CHEST_TASK_NAME,
		Description:                 "Go task example from code",
		RetryCount:                  3,
		TimeoutSeconds:              300,
		InputKeys:                   []string{"importantValue"},
		OutputKeys:                  make([]string, 0),
		TimeoutPolicy:               "TIME_OUT_WF",
		RetryLogic:                  "FIXED",
		RetryDelaySeconds:           10,
		ResponseTimeoutSeconds:      180,
		RateLimitPerFrequency:       0,
		RateLimitFrequencyInSeconds: 1,
		OwnerEmail:                  "gustavo.gardusi@orkes.io",
		BackoffScaleFactor:          1,
	}

	IMPORTANT_VALUE = "Go is really nice :)"

	TREASURE_WORKFLOW_INPUT = &TreasureChest{
		ImportantValue: IMPORTANT_VALUE,
	}
)

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

func ValidateWorkflow(conductorWorkflow *definition.ConductorWorkflow, timeout time.Duration) error {
	err := ValidateWorkflowRegistration(conductorWorkflow)
	if err != nil {
		return err
	}
	workflowId, err := conductorWorkflow.StartWorkflowWithInput(make(map[string]interface{}))
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	workflowExecutionChannel, err := WorkflowExecutor.MonitorExecution(workflowId)
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

func ValidateWorkflowBulk(conductorWorkflow *definition.ConductorWorkflow, timeout time.Duration, amount int) error {
	err := ValidateWorkflowRegistration(conductorWorkflow)
	if err != nil {
		return err
	}
	version := conductorWorkflow.GetVersion()
	startWorkflowRequests := make([]*model.StartWorkflowRequest, amount)
	for i := 0; i < amount; i += 1 {
		startWorkflowRequests[i] = model.NewStartWorkflowRequest(
			conductorWorkflow.GetName(),
			&version,
			"",
			make(map[string]interface{}),
		)
	}
	runningWorkflows := WorkflowExecutor.StartWorkflows(true, startWorkflowRequests...)
	for _, runningWorkflow := range runningWorkflows {
		if runningWorkflow.Err != nil {
			return err
		}
		workflow, err := executor.WaitForWorkflowCompletionUntilTimeout(
			runningWorkflow.WorkflowExecutionChannel,
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

func ValidateWorkflowRegistration(workflow *definition.ConductorWorkflow) error {
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

func isWorkflowCompleted(workflow *model.Workflow) bool {
	return workflow.Status == model.CompletedWorkflow
}
