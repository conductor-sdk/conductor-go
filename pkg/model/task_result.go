package model

type TaskResult struct {
	WorkflowInstanceId               string                 `json:"workflowInstanceId"`
	TaskId                           string                 `json:"taskId"`
	ReasonForIncompletion            string                 `json:"reasonForIncompletion,omitempty"`
	CallbackAfterSeconds             int64                  `json:"callbackAfterSeconds,omitempty"`
	WorkerId                         string                 `json:"workerId,omitempty"`
	Status                           TaskResultStatus       `json:"status,omitempty"`
	OutputData                       map[string]interface{} `json:"outputData,omitempty"`
	Logs                             []TaskExecLog          `json:"logs,omitempty"`
	ExternalOutputPayloadStoragePath string                 `json:"externalOutputPayloadStoragePath,omitempty"`
	SubWorkflowId                    string                 `json:"subWorkflowId,omitempty"`
}
