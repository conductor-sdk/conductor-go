package model

import "github.com/conductor-sdk/conductor-go/pkg/model/enum/workflow_status"

type Workflow struct {
	OwnerApp                         string                         `json:"ownerApp,omitempty"`
	CreateTime                       int64                          `json:"createTime,omitempty"`
	UpdateTime                       int64                          `json:"updateTime,omitempty"`
	CreatedBy                        string                         `json:"createdBy,omitempty"`
	UpdatedBy                        string                         `json:"updatedBy,omitempty"`
	Status                           workflow_status.WorkflowStatus `json:"status,omitempty"`
	EndTime                          int64                          `json:"endTime,omitempty"`
	WorkflowId                       string                         `json:"workflowId,omitempty"`
	ParentWorkflowId                 string                         `json:"parentWorkflowId,omitempty"`
	ParentWorkflowTaskId             string                         `json:"parentWorkflowTaskId,omitempty"`
	Tasks                            []Task                         `json:"tasks,omitempty"`
	Input                            map[string]interface{}         `json:"input,omitempty"`
	Output                           map[string]interface{}         `json:"output,omitempty"`
	CorrelationId                    string                         `json:"correlationId,omitempty"`
	ReRunFromWorkflowId              string                         `json:"reRunFromWorkflowId,omitempty"`
	ReasonForIncompletion            string                         `json:"reasonForIncompletion,omitempty"`
	Event                            string                         `json:"event,omitempty"`
	TaskToDomain                     map[string]string              `json:"taskToDomain,omitempty"`
	FailedReferenceTaskNames         []string                       `json:"failedReferenceTaskNames,omitempty"`
	WorkflowDefinition               *WorkflowDef                   `json:"workflowDefinition,omitempty"`
	ExternalInputPayloadStoragePath  string                         `json:"externalInputPayloadStoragePath,omitempty"`
	ExternalOutputPayloadStoragePath string                         `json:"externalOutputPayloadStoragePath,omitempty"`
	Priority                         int32                          `json:"priority,omitempty"`
	Variables                        map[string]interface{}         `json:"variables,omitempty"`
	LastRetriedTime                  int64                          `json:"lastRetriedTime,omitempty"`
	StartTime                        int64                          `json:"startTime,omitempty"`
	WorkflowName                     string                         `json:"workflowName,omitempty"`
	WorkflowVersion                  int32                          `json:"workflowVersion,omitempty"`
}
