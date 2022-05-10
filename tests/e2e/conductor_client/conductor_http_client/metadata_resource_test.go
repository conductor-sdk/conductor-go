package conductor_http_client

import (
	"context"
	"testing"

	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/conductor-sdk/conductor-go/tests/e2e/conductor_client"
)

func TestRegisterTaskDefinition(t *testing.T) {
	apiClient := conductor_http_client.NewAPIClient(
		conductor_client.GetAuthenticationSettings(),
		conductor_client.GetHttpSettingsWithAuth(),
	)
	metadataClient := *&conductor_http_client.MetadataResourceApiService{
		APIClient: apiClient,
	}
	response, err := metadataClient.RegisterTaskDef(context.Background(), getTaskDefinition())
	if err != nil {
		t.Error("response: ", response)
	}
}

func TestRegisterWorkflowDefinition(t *testing.T) {
	apiClient := conductor_http_client.NewAPIClient(
		conductor_client.GetAuthenticationSettings(),
		conductor_client.GetHttpSettingsWithAuth(),
	)
	metadataClient := *&conductor_http_client.MetadataResourceApiService{
		APIClient: apiClient,
	}
	response, err := metadataClient.RegisterWorkflowDef(
		context.Background(),
		getWorkflowDefinition(),
	)

	if err != nil && response.Status != "409 Conflict" {
		t.Error("response:", response)
	}
}

func getTaskDefinition() []http_model.TaskDef {
	taskDef := http_model.TaskDef{
		CreateTime:                  1650595379661,
		CreatedBy:                   "",
		Name:                        "go_task_example_from_code",
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
	taskDefs := make([]http_model.TaskDef, 0)
	return append(taskDefs, taskDef)
}

func getWorkflowDefinition() http_model.WorkflowDef {
	return http_model.WorkflowDef{
		UpdateTime:  1650595431465,
		Name:        "workflow_with_go_task_example_from_code",
		Description: "Workflow with go task example from code",
		Version:     1,
		Tasks: []http_model.WorkflowTask{
			{
				Name:              "go_task_example_from_code",
				TaskReferenceName: "go_task_example_from_code_ref_0",
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
}
