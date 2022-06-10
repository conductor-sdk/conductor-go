package e2e_properties

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/model"
	"github.com/conductor-sdk/conductor-go/pkg/model/enum/workflow_status"
	"github.com/conductor-sdk/conductor-go/pkg/settings"
	"github.com/conductor-sdk/conductor-go/pkg/worker"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/def/workflow"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/executor"

	log "github.com/sirupsen/logrus"
)

const (
	AUTHENTICATION_KEY_ID     = "KEY"
	AUTHENTICATION_KEY_SECRET = "SECRET"
	BASE_URL                  = "https://play.orkes.io/api"
)

var (
	apiClient = getApiClientWithAuthentication()
)

var (
	MetadataClient = conductor_http_client.MetadataResourceApiService{
		APIClient: apiClient,
	}
	TaskClient = conductor_http_client.TaskResourceApiService{
		APIClient: apiClient,
	}
	WorkflowClient = conductor_http_client.WorkflowResourceApiService{
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
	if workflow.Status != string(workflow_status.COMPLETED) {
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

func getApiClientWithAuthentication() *conductor_http_client.APIClient {
	return conductor_http_client.NewAPIClient(
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
	WORKFLOW_DEFINITIONS = []model.WorkflowDef{
		WORKFLOW_DEFINITION,
		TREASURE_WORKFLOW_DEFINITION,
	}
	TASK_DEFINITIONS = []model.TaskDef{
		TASK_DEFINITION,
		TREASURE_TASK_DEFINITION,
	}
)

var (
	WORKFLOW_NAME = "workflow_with_go_task_example_from_code"

	TASK_NAME = "GO_TASK_OF_SIMPLE_TYPE"

	WORKFLOW_DEFINITION = model.WorkflowDef{
		UpdateTime:  1650595431465,
		Name:        WORKFLOW_NAME,
		Description: "Workflow with go task example from code",
		Version:     1,
		Tasks: []model.WorkflowTask{
			{
				Name:              TASK_NAME,
				TaskReferenceName: TASK_NAME,
				Type_:             "SIMPLE",
				StartDelay:        0,
				Optional:          false,
				AsyncComplete:     false,
			},
		},
		OutputParameters: map[string]interface{}{
			"workerOutput": "${go_task_example_from_code_ref_0.output}",
		},
		SchemaVersion:                 2,
		Restartable:                   true,
		WorkflowStatusListenerEnabled: false,
		OwnerEmail:                    "gustavo.gardusi@orkes.io",
		TimeoutPolicy:                 "ALERT_ONLY",
		TimeoutSeconds:                0,
	}

	TASK_DEFINITION = model.TaskDef{
		Name:                        TASK_NAME,
		Description:                 "Go task example from code",
		RetryCount:                  3,
		TimeoutSeconds:              300,
		InputKeys:                   make([]string, 0),
		OutputKeys:                  make([]string, 0),
		TimeoutPolicy:               "TIME_OUT_WF",
		RetryLogic:                  "FIXED",
		RetryDelaySeconds:           10,
		ResponseTimeoutSeconds:      180,
		InputTemplate:               make(map[string]interface{}),
		RateLimitPerFrequency:       0,
		RateLimitFrequencyInSeconds: 1,
		OwnerEmail:                  "gustavo.gardusi@orkes.io",
		BackoffScaleFactor:          1,
	}
)

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

func ValidateWorkflow(conductorWorkflow *workflow.ConductorWorkflow, timeout time.Duration) error {
	err := ValidateWorkflowRegistration(conductorWorkflow)
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

func ValidateWorkflowBulk(conductorWorkflow *workflow.ConductorWorkflow, timeout time.Duration, amount int) error {
	err := ValidateWorkflowRegistration(conductorWorkflow)
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

func ValidateTaskRegistration(task *workflow.SimpleTask) error {
	response, err := MetadataClient.RegisterTaskDef(
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

func ValidateWorkflowRegistration(workflow *workflow.ConductorWorkflow) error {
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
	return workflow.Status == string(workflow_status.COMPLETED)
}
