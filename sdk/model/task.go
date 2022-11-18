// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.
package model

type Task struct {
	CallbackAfterSeconds             int64                  `json:"callbackAfterSeconds,omitempty"`
	CallbackFromWorker               bool                   `json:"callbackFromWorker,omitempty"`
	CorrelationId                    string                 `json:"correlationId,omitempty"`
	Domain                           string                 `json:"domain,omitempty"`
	EndTime                          int64                  `json:"endTime,omitempty"`
	Executed                         bool                   `json:"executed,omitempty"`
	ExecutionNameSpace               string                 `json:"executionNameSpace,omitempty"`
	ExternalInputPayloadStoragePath  string                 `json:"externalInputPayloadStoragePath,omitempty"`
	ExternalOutputPayloadStoragePath string                 `json:"externalOutputPayloadStoragePath,omitempty"`
	InputData                        map[string]interface{} `json:"inputData,omitempty"`
	IsolationGroupId                 string                 `json:"isolationGroupId,omitempty"`
	Iteration                        int32                  `json:"iteration,omitempty"`
	LoopOverTask                     bool                   `json:"loopOverTask,omitempty"`
	OutputData                       map[string]interface{} `json:"outputData,omitempty"`
	PollCount                        int32                  `json:"pollCount,omitempty"`
	QueueWaitTime                    int64                  `json:"queueWaitTime,omitempty"`
	RateLimitFrequencyInSeconds      int32                  `json:"rateLimitFrequencyInSeconds,omitempty"`
	RateLimitPerFrequency            int32                  `json:"rateLimitPerFrequency,omitempty"`
	ReasonForIncompletion            string                 `json:"reasonForIncompletion,omitempty"`
	ReferenceTaskName                string                 `json:"referenceTaskName,omitempty"`
	ResponseTimeoutSeconds           int64                  `json:"responseTimeoutSeconds,omitempty"`
	Retried                          bool                   `json:"retried,omitempty"`
	RetriedTaskId                    string                 `json:"retriedTaskId,omitempty"`
	RetryCount                       int32                  `json:"retryCount,omitempty"`
	ScheduledTime                    int64                  `json:"scheduledTime,omitempty"`
	Seq                              int32                  `json:"seq,omitempty"`
	StartDelayInSeconds              int32                  `json:"startDelayInSeconds,omitempty"`
	StartTime                        int64                  `json:"startTime,omitempty"`
	Status                           string                 `json:"status,omitempty"`
	SubWorkflowId                    string                 `json:"subWorkflowId,omitempty"`
	SubworkflowChanged               bool                   `json:"subworkflowChanged,omitempty"`
	TaskDefName                      string                 `json:"taskDefName,omitempty"`
	TaskDefinition                   *TaskDef               `json:"taskDefinition,omitempty"`
	TaskId                           string                 `json:"taskId,omitempty"`
	TaskType                         string                 `json:"taskType,omitempty"`
	UpdateTime                       int64                  `json:"updateTime,omitempty"`
	WorkerId                         string                 `json:"workerId,omitempty"`
	WorkflowInstanceId               string                 `json:"workflowInstanceId,omitempty"`
	WorkflowPriority                 int32                  `json:"workflowPriority,omitempty"`
	WorkflowTask                     *WorkflowTask          `json:"workflowTask,omitempty"`
	WorkflowType                     string                 `json:"workflowType,omitempty"`
}
