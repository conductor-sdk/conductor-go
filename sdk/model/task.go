//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package model

type Task struct {
	TaskType                         string                 `json:"taskType,omitempty"`
	Status                           TaskResultStatus       `json:"status,omitempty"`
	InputData                        map[string]interface{} `json:"inputData,omitempty"`
	ReferenceTaskName                string                 `json:"referenceTaskName,omitempty"`
	RetryCount                       int32                  `json:"retryCount,omitempty"`
	Seq                              int32                  `json:"seq,omitempty"`
	CorrelationId                    string                 `json:"correlationId,omitempty"`
	PollCount                        int32                  `json:"pollCount,omitempty"`
	TaskDefName                      string                 `json:"taskDefName,omitempty"`
	ScheduledTime                    int64                  `json:"scheduledTime,omitempty"`
	StartTime                        int64                  `json:"startTime,omitempty"`
	EndTime                          int64                  `json:"endTime,omitempty"`
	UpdateTime                       int64                  `json:"updateTime,omitempty"`
	StartDelayInSeconds              int32                  `json:"startDelayInSeconds,omitempty"`
	RetriedTaskId                    string                 `json:"retriedTaskId,omitempty"`
	Retried                          bool                   `json:"retried,omitempty"`
	Executed                         bool                   `json:"executed,omitempty"`
	CallbackFromWorker               bool                   `json:"callbackFromWorker,omitempty"`
	ResponseTimeoutSeconds           int64                  `json:"responseTimeoutSeconds,omitempty"`
	WorkflowInstanceId               string                 `json:"workflowInstanceId,omitempty"`
	WorkflowType                     string                 `json:"workflowType,omitempty"`
	TaskId                           string                 `json:"taskId,omitempty"`
	ReasonForIncompletion            string                 `json:"reasonForIncompletion,omitempty"`
	CallbackAfterSeconds             int64                  `json:"callbackAfterSeconds,omitempty"`
	WorkerId                         string                 `json:"workerId,omitempty"`
	OutputData                       map[string]interface{} `json:"outputData,omitempty"`
	WorkflowTask                     *WorkflowTask          `json:"workflowTask,omitempty"`
	Domain                           string                 `json:"domain,omitempty"`
	RateLimitPerFrequency            int32                  `json:"rateLimitPerFrequency,omitempty"`
	RateLimitFrequencyInSeconds      int32                  `json:"rateLimitFrequencyInSeconds,omitempty"`
	ExternalInputPayloadStoragePath  string                 `json:"externalInputPayloadStoragePath,omitempty"`
	ExternalOutputPayloadStoragePath string                 `json:"externalOutputPayloadStoragePath,omitempty"`
	WorkflowPriority                 int32                  `json:"workflowPriority,omitempty"`
	ExecutionNameSpace               string                 `json:"executionNameSpace,omitempty"`
	IsolationGroupId                 string                 `json:"isolationGroupId,omitempty"`
	Iteration                        int32                  `json:"iteration,omitempty"`
	SubWorkflowId                    string                 `json:"subWorkflowId,omitempty"`
	SubworkflowChanged               bool                   `json:"subworkflowChanged,omitempty"`
	QueueWaitTime                    int64                  `json:"queueWaitTime,omitempty"`
	TaskDefinition                   *TaskDef               `json:"taskDefinition,omitempty"`
	LoopOverTask                     bool                   `json:"loopOverTask,omitempty"`
}
