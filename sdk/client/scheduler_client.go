package client

import (
	"context"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"net/http"
)

type SchedulerClient interface {
	DeleteSchedule(ctx context.Context, name string) (interface{}, *http.Response, error)
	DeleteTagForSchedule(ctx context.Context, body []model.Tag, name string) (*http.Response, error)
	GetAllSchedules(ctx context.Context, optionals *SchedulerResourceApiGetAllSchedulesOpts) ([]model.WorkflowScheduleModel, *http.Response, error)
	GetNextFewSchedules(ctx context.Context, cronExpression string, optionals *SchedulerResourceApiGetNextFewSchedulesOpts) ([]int64, *http.Response, error)
	GetSchedule(ctx context.Context, name string) (model.WorkflowSchedule, *http.Response, error)
	GetTagsForSchedule(ctx context.Context, name string) ([]model.Tag, *http.Response, error)
	PauseAllSchedules(ctx context.Context) (map[string]interface{}, *http.Response, error)
	PauseSchedule(ctx context.Context, name string) (interface{}, *http.Response, error)
	PutTagForSchedule(ctx context.Context, body []model.Tag, name string) (*http.Response, error)
	RequeueAllExecutionRecords(ctx context.Context) (map[string]interface{}, *http.Response, error)
	ResumeAllSchedules(ctx context.Context) (map[string]interface{}, *http.Response, error)
	ResumeSchedule(ctx context.Context, name string) (interface{}, *http.Response, error)
	SaveSchedule(ctx context.Context, body model.SaveScheduleRequest) (interface{}, *http.Response, error)
	SearchV2(ctx context.Context, optionals *SchedulerSearchOpts) (model.SearchResultWorkflowSchedule, *http.Response, error)
}

func NewSchedulerClient(apiClient *APIClient) SchedulerClient {
	return &SchedulerResourceApiService{apiClient}
}
