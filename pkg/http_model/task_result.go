package http_model

import "github.com/conductor-sdk/conductor-go/pkg/model/enum/task_result_status"

type TaskResult struct {
	WorkflowInstanceId               string                              `json:"workflowInstanceId"`
	TaskId                           string                              `json:"taskId"`
	ReasonForIncompletion            string                              `json:"reasonForIncompletion,omitempty"`
	CallbackAfterSeconds             int64                               `json:"callbackAfterSeconds,omitempty"`
	WorkerId                         string                              `json:"workerId,omitempty"`
	Status                           task_result_status.TaskResultStatus `json:"status,omitempty"`
	OutputData                       map[string]interface{}              `json:"outputData,omitempty"`
	Logs                             []TaskExecLog                       `json:"logs,omitempty"`
	ExternalOutputPayloadStoragePath string                              `json:"externalOutputPayloadStoragePath,omitempty"`
	SubWorkflowId                    string                              `json:"subWorkflowId,omitempty"`
}
