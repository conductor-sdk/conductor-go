package model

import (
	"encoding/json"

	"github.com/netflix/conductor/client/go/conductor_client/model/enum/task_result_status"
)

type Task struct {
	TaskType          string                              `json:"taskType"`
	Status            task_result_status.TaskResultStatus `json:"status"`
	InputData         map[string]interface{}              `json:"inputData"`
	ReferenceTaskName string                              `json:"referenceTaskName"`
	RetryCount        int                                 `json:"retryCount"`
	Seq               int                                 `json:"seq"`
	CorrelationId     string                              `json:"correlationId"`
	PollCount         int                                 `json:"pollCount"`
	TaskDefName       string                              `json:"taskDefName"`
	// Time when the task was scheduled
	ScheduledTime int64 `json:"scheduledTime"`
	// Time when the task was first polled
	StartTime int64 `json:"startTime"`
	// Time when the task completed executing
	EndTime int64 `json:"endTime"`
	// Time when the task was last updated
	UpdateTime          int64  `json:"updateTime"`
	StartDelayInSeconds int    `json:"startDelayInSeconds"`
	RetriedTaskId       string `json:"retriedTaskId"`
	Retried             bool   `json:"retried"`
	// Default = true
	CallbackFromWorker bool `json:"callbackFromWorker"`
	// DynamicWorkflowTask
	ResponseTimeoutSeconds int                    `json:"responseTimeoutSeconds"`
	WorkflowInstanceId     string                 `json:"workflowInstanceId"`
	TaskId                 string                 `json:"taskId"`
	ReasonForIncompletion  string                 `json:"reasonForIncompletion"`
	CallbackAfterSeconds   int64                  `json:"callbackAfterSeconds"`
	WorkerId               string                 `json:"workerId"`
	OutputData             map[string]interface{} `json:"outputData"`
}

// "Constructor" to initialze non zero value defaults
func NewTask() *Task {
	task := new(Task)
	task.CallbackFromWorker = true
	task.InputData = make(map[string]interface{})
	task.OutputData = make(map[string]interface{})
	return task
}

func (t *Task) ToJSONString() (string, error) {
	var jsonString string
	b, err := json.Marshal(t)
	if err == nil {
		jsonString = string(b)
	}
	return jsonString, err
}

func ParseTask(inputJSON string) (*Task, error) {
	t := NewTask()
	err := json.Unmarshal([]byte(inputJSON), t)
	return t, err
}
