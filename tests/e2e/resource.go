package e2e

import (
	"fmt"
	"os"

	"github.com/conductor-sdk/conductor-go/examples"
	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/conductor-sdk/conductor-go/pkg/model"
	"github.com/conductor-sdk/conductor-go/pkg/settings"
	log "github.com/sirupsen/logrus"
)

type TreasureChest struct {
	ImportantValue string `json:"importantValue"`
}

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

var (
	BASE_URL = "https://play.orkes.io"

	TASK_OUTPUT = map[string]interface{}{"hello": "world"}

	AUTHENTICATION_KEY_ID     = "KEY"
	AUTHENTICATION_KEY_SECRET = "SECRET"

	WORKER_THREAD_COUNT     = 5
	WORKER_POLLING_INTERVAL = 100

	WORKFLOW_EXECUTION_AMOUNT = 5

	WORKFLOW_DEFINITIONS = []http_model.WorkflowDef{
		WORKFLOW_DEFINITION,
		TREASURE_WORKFLOW_DEFINITION,
	}
	TASK_DEFINITIONS = []http_model.TaskDef{
		TASK_DEFINITION,
		TREASURE_TASK_DEFINITION,
	}
	TASK_DEFINITION_TO_WORKER = map[string]model.TaskExecuteFunction{
		TASK_DEFINITION.Name:     examples.SimpleWorker,
		TREASURE_CHEST_TASK_NAME: examples.OpenTreasureChest,
	}
)

var API_CLIENT = getApiClientWithAuthentication()

var (
	WORKFLOW_NAME = "workflow_with_go_task_example_from_code"

	TASK_NAME           = "go_task_example_from_code"
	TASK_REFERENCE_NAME = "go_task_example_from_code_ref_0"

	WORKFLOW_DEFINITION = http_model.WorkflowDef{
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
	TASK_DEFINITION = http_model.TaskDef{
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
)

var (
	TREASURE_CHEST_WORKFLOW_NAME = "treasure_chest_workflow"
	TREASURE_CHEST_TASK_NAME     = "treasure_chest_task"

	TREASURE_WORKFLOW_DEFINITION = http_model.WorkflowDef{
		UpdateTime:  1650595431465,
		Name:        TREASURE_CHEST_WORKFLOW_NAME,
		Description: "What's inside the treasure chest?",
		Version:     1,
		Tasks: []http_model.WorkflowTask{
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

	TREASURE_TASK_DEFINITION = http_model.TaskDef{
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
