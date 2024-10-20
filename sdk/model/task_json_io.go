//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package model

import (
	"encoding/json"
)

type PolledTask struct {
	TaskType                         string           `json:"taskType,omitempty"`
	Status                           TaskResultStatus `json:"status,omitempty"`
	InputData                        json.RawMessage  `json:"inputData,omitempty"`
	ReferenceTaskName                string           `json:"referenceTaskName,omitempty"`
	RetryCount                       int32            `json:"retryCount,omitempty"`
	Seq                              int32            `json:"seq,omitempty"`
	CorrelationId                    string           `json:"correlationId,omitempty"`
	PollCount                        int32            `json:"pollCount,omitempty"`
	TaskDefName                      string           `json:"taskDefName,omitempty"`
	ScheduledTime                    int64            `json:"scheduledTime,omitempty"`
	StartTime                        int64            `json:"startTime,omitempty"`
	EndTime                          int64            `json:"endTime,omitempty"`
	UpdateTime                       int64            `json:"updateTime,omitempty"`
	StartDelayInSeconds              int32            `json:"startDelayInSeconds,omitempty"`
	RetriedTaskId                    string           `json:"retriedTaskId,omitempty"`
	Retried                          bool             `json:"retried,omitempty"`
	Executed                         bool             `json:"executed,omitempty"`
	CallbackFromWorker               bool             `json:"callbackFromWorker,omitempty"`
	ResponseTimeoutSeconds           int64            `json:"responseTimeoutSeconds,omitempty"`
	WorkflowInstanceId               string           `json:"workflowInstanceId,omitempty"`
	WorkflowType                     string           `json:"workflowType,omitempty"`
	TaskId                           string           `json:"taskId,omitempty"`
	ReasonForIncompletion            string           `json:"reasonForIncompletion,omitempty"`
	CallbackAfterSeconds             int64            `json:"callbackAfterSeconds,omitempty"`
	WorkerId                         string           `json:"workerId,omitempty"`
	OutputData                       json.RawMessage  `json:"outputData,omitempty"`
	WorkflowTask                     *WorkflowTask    `json:"workflowTask,omitempty"`
	Domain                           string           `json:"domain,omitempty"`
	RateLimitPerFrequency            int32            `json:"rateLimitPerFrequency,omitempty"`
	RateLimitFrequencyInSeconds      int32            `json:"rateLimitFrequencyInSeconds,omitempty"`
	ExternalInputPayloadStoragePath  string           `json:"externalInputPayloadStoragePath,omitempty"`
	ExternalOutputPayloadStoragePath string           `json:"externalOutputPayloadStoragePath,omitempty"`
	WorkflowPriority                 int32            `json:"workflowPriority,omitempty"`
	ExecutionNameSpace               string           `json:"executionNameSpace,omitempty"`
	IsolationGroupId                 string           `json:"isolationGroupId,omitempty"`
	Iteration                        int32            `json:"iteration,omitempty"`
	SubWorkflowId                    string           `json:"subWorkflowId,omitempty"`
	SubworkflowChanged               bool             `json:"subworkflowChanged,omitempty"`
	QueueWaitTime                    int64            `json:"queueWaitTime,omitempty"`
	TaskDefinition                   *TaskDef         `json:"taskDefinition,omitempty"`
	LoopOverTask                     bool             `json:"loopOverTask,omitempty"`
}

func (task *PolledTask) ToValue(newValue any) (interface{}, error) {

	bytes, err := task.InputData.MarshalJSON()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, newValue)
	if err != nil {
		return nil, err
	}
	return newValue, nil
}

func (pt *PolledTask) ToTask() (*Task, error) {
	var inputData map[string]interface{}
	if err := json.Unmarshal(pt.InputData, &inputData); err != nil {
		return nil, err
	}

	// Convert OutputData from json.RawMessage to map[string]interface{}
	var outputData map[string]interface{}
	if err := json.Unmarshal(pt.OutputData, &outputData); err != nil {
		return nil, err
	}
	return &Task{
		TaskType:                         pt.TaskType,
		Status:                           pt.Status,
		InputData:                        inputData,
		ReferenceTaskName:                pt.ReferenceTaskName,
		RetryCount:                       pt.RetryCount,
		Seq:                              pt.Seq,
		CorrelationId:                    pt.CorrelationId,
		PollCount:                        pt.PollCount,
		TaskDefName:                      pt.TaskDefName,
		ScheduledTime:                    pt.ScheduledTime,
		StartTime:                        pt.StartTime,
		EndTime:                          pt.EndTime,
		UpdateTime:                       pt.UpdateTime,
		StartDelayInSeconds:              pt.StartDelayInSeconds,
		RetriedTaskId:                    pt.RetriedTaskId,
		Retried:                          pt.Retried,
		Executed:                         pt.Executed,
		CallbackFromWorker:               pt.CallbackFromWorker,
		ResponseTimeoutSeconds:           pt.ResponseTimeoutSeconds,
		WorkflowInstanceId:               pt.WorkflowInstanceId,
		WorkflowType:                     pt.WorkflowType,
		TaskId:                           pt.TaskId,
		ReasonForIncompletion:            pt.ReasonForIncompletion,
		CallbackAfterSeconds:             pt.CallbackAfterSeconds,
		WorkerId:                         pt.WorkerId,
		OutputData:                       outputData,
		WorkflowTask:                     pt.WorkflowTask,
		Domain:                           pt.Domain,
		RateLimitPerFrequency:            pt.RateLimitPerFrequency,
		RateLimitFrequencyInSeconds:      pt.RateLimitFrequencyInSeconds,
		ExternalInputPayloadStoragePath:  pt.ExternalInputPayloadStoragePath,
		ExternalOutputPayloadStoragePath: pt.ExternalOutputPayloadStoragePath,
		WorkflowPriority:                 pt.WorkflowPriority,
		ExecutionNameSpace:               pt.ExecutionNameSpace,
		IsolationGroupId:                 pt.IsolationGroupId,
		Iteration:                        pt.Iteration,
		SubWorkflowId:                    pt.SubWorkflowId,
		SubworkflowChanged:               pt.SubworkflowChanged,
		QueueWaitTime:                    pt.QueueWaitTime,
		TaskDefinition:                   pt.TaskDefinition,
		LoopOverTask:                     pt.LoopOverTask,
	}, nil

}

func (task *PolledTask) ToTaskResult() *TaskResult {
	return &TaskResult{
		TaskId:             task.TaskId,
		WorkflowInstanceId: task.WorkflowInstanceId,
		WorkerId:           getHostname(),
	}
}

func (task *PolledTask) ToTaskResultWithError(err error) *TaskResult {
	taskResult := task.ToTaskResult()
	taskResult.ReasonForIncompletion = err.Error()
	switch err.(type) {
	case *NonRetryableError:
		taskResult.Status = FailedWithTerminalErrorTask
	default:
		taskResult.Status = FailedTask
	}
	return taskResult
}
func (task *PolledTask) ToTaskResultExecutionOutput(taskExecutionOutput interface{}) (*TaskResult, error) {
	taskResult, ok := taskExecutionOutput.(*TaskResult)
	if !ok {
		taskResult = task.ToTaskResult()
		taskResult.OutputDataRaw = taskExecutionOutput
		taskResult.Status = CompletedTask
	}
	return taskResult, nil
}
