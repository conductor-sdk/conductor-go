package client

import (
	"context"
	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"net/http"
)

type TaskClient interface {
	All(ctx context.Context) (map[string]int64, *http.Response, error)
	AllVerbose(ctx context.Context) (map[string]map[string]map[string]int64, *http.Response, error)
	BatchPoll(ctx context.Context, tasktype string, localVarOptionals *TaskResourceApiBatchPollOpts) ([]model.Task, *http.Response, error)
	GetAllPollData(ctx context.Context) ([]model.PollData, *http.Response, error)
	GetExternalStorageLocation1(ctx context.Context, path string, operation string, payloadType string) (model.ExternalStorageLocation, *http.Response, error)
	GetPollData(ctx context.Context, taskType string) ([]model.PollData, *http.Response, error)
	GetTask(ctx context.Context, taskId string) (model.Task, *http.Response, error)
	GetTaskLogs(ctx context.Context, taskId string) ([]model.TaskExecLog, *http.Response, error)
	Log(ctx context.Context, body string, taskId string) (*http.Response, error)
	Poll(ctx context.Context, tasktype string, localVarOptionals *TaskResourceApiPollOpts) (model.Task, *http.Response, error)
	RequeuePendingTask(ctx context.Context, taskType string) (string, *http.Response, error)
	Search(ctx context.Context, localVarOptionals *TaskResourceApiSearch1Opts) (model.SearchResultTaskSummary, *http.Response, error)
	SearchV2(ctx context.Context, localVarOptionals *TaskResourceApiSearchV21Opts) (model.SearchResultTask, *http.Response, error)
	Size(ctx context.Context, localVarOptionals *TaskResourceApiSizeOpts) (map[string]int32, *http.Response, error)
	UpdateTask(ctx context.Context, taskResult *model.TaskResult) (string, *http.Response, error)
	UpdateTaskByRefName(ctx context.Context, body map[string]interface{}, workflowId string, taskRefName string, status string) (string, *http.Response, error)
	UpdateTaskByRefNameWithWorkerId(ctx context.Context, body map[string]interface{}, workflowId string, taskRefName string, status string, workerId optional.String) (string, *http.Response, error)
	updateTaskByRefName(ctx context.Context, body map[string]interface{}, workflowId string, taskRefName string, status string, workerId optional.String) (string, *http.Response, error)
	SignalAsync(ctx context.Context, body map[string]interface{}, workflowId string, status string) (*http.Response, error)
	Signal(ctx context.Context, body map[string]interface{}, workflowID string, status model.WorkflowStatus, opts ...SignalTaskOpts) (*model.SignalResponse, error)
}

func NewTaskClient(apiClient *APIClient) TaskClient {
	return &TaskResourceApiService{apiClient}
}
