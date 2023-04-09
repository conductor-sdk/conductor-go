package client

import (
	"context"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"net/http"
)

type SchedulerService interface {
	DeleteSchedule(ctx context.Context, name string) (*http.Response, error)

	DeleteTagForSchedule(ctx context.Context, body []model.Tag, name string) (*http.Response, error)

	GetAllSchedules(ctx context.Context, localVarOptionals *SchedulerResourceApiGetAllSchedulesOpts) ([]model.WorkflowSchedule, *http.Response, error)

	GetNextFewSchedules(ctx context.Context, cronExpression string, localVarOptionals *SchedulerResourceApiGetNextFewSchedulesOpts) ([]int64, *http.Response, error)

	GetSchedule(ctx context.Context, name string) (model.WorkflowSchedule, *http.Response, error)

	GetTagsForSchedule(ctx context.Context, name string) ([]model.Tag, *http.Response, error)

	PauseAllSchedules(ctx context.Context) (map[string]interface{}, *http.Response, error)

	PauseSchedule(ctx context.Context, name string) (*http.Response, error)

	AddTagForSchedule(ctx context.Context, body []model.Tag, name string) (*http.Response, error)

	RequeueAllExecutionRecords(ctx context.Context) (map[string]interface{}, *http.Response, error)

	ResumeAllSchedules(ctx context.Context) (map[string]interface{}, *http.Response, error)

	ResumeSchedule(ctx context.Context, name string) (*http.Response, error)

	SaveSchedule(ctx context.Context, body model.SaveScheduleRequest) (*http.Response, error)

	Search(ctx context.Context, localVarOptionals *SchedulerSearchOpts) (model.SearchResultWorkflowScheduleExecutionModel, *http.Response, error)
}
