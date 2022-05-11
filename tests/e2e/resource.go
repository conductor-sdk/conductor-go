package e2e

import (
	"os"

	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/conductor-sdk/conductor-go/pkg/settings"
)

var (
	BASE_URL = "https://play.orkes.io"

	WORKFLOW_NAME       = "workflow_with_go_task_example_from_code"
	TASK_NAME           = "go_task_example_from_code"
	TASK_REFERENCE_NAME = "go_task_example_from_code_ref_0"

	TASK_OUTPUT = map[string]interface{}{"hello": "world"}

	AUTHENTICATION_KEY_ID     = "KEY"
	AUTHENTICATION_KEY_SECRET = "SECRET"
)

var WORKFLOW_DEFINITION = http_model.WorkflowDef{
	UpdateTime:  1650595431465,
	Name:        WORKFLOW_NAME,
	Description: "Workflow with go task example from code",
	Version:     1,
	Tasks: []http_model.WorkflowTask{
		{
			Name:              TASK_NAME,
			TaskReferenceName: TASK_REFERENCE_NAME,
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

var TASK_DEFINITION = http_model.TaskDef{
	CreateTime:                  1650595379661,
	CreatedBy:                   "",
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

var apiClient = getApiClientWithAuthentication()

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
