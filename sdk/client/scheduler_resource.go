//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package client

import (
	"context"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"net/http"
)

// SchedulerClient manage workflow schedules
type SchedulerClient interface {
	// DeleteSchedule Deletes the schedule given the name
	DeleteSchedule(ctx context.Context, name string) (*http.Response, error)

	//DeleteTagForSchedule Remove the tags from the schedule
	DeleteTagForSchedule(ctx context.Context, tags []model.Tag, name string) (*http.Response, error)

	//GetAllSchedules Retrieve all the schedules
	GetAllSchedules(ctx context.Context, localVarOptionals *GetAllSchedulesOpts) ([]model.WorkflowSchedule, *http.Response, error)

	//GetNextFewSchedules given the cron expression retrieves the next few schedules.  Useful for testing
	GetNextFewSchedules(ctx context.Context, cronExpression string, localVarOptionals *NextFewSchedulesOpts) ([]int64, *http.Response, error)

	//GetSchedule Retrieve schedule definition given the name
	GetSchedule(ctx context.Context, name string) (model.WorkflowSchedule, *http.Response, error)

	//GetTagsForSchedule get tags associated with the schedule
	GetTagsForSchedule(ctx context.Context, name string) ([]model.Tag, *http.Response, error)

	//PauseAllSchedules WARNING: pauses ALL the schedules in the system.  Use with caution!
	PauseAllSchedules(ctx context.Context) (map[string]interface{}, *http.Response, error)

	//PauseSchedule pause the schedule by name
	PauseSchedule(ctx context.Context, name string) (*http.Response, error)

	//AddTagForSchedule Adds tags to the schedule
	AddTagForSchedule(ctx context.Context, tags []model.Tag, name string) (*http.Response, error)

	//ResumeAllSchedules Resume ALL the schedule.  WARNING: use with caution!
	ResumeAllSchedules(ctx context.Context) (map[string]interface{}, *http.Response, error)

	//ResumeSchedule Resumes the schedule by name
	ResumeSchedule(ctx context.Context, name string) (*http.Response, error)

	//SaveSchedule Upsert a new schedule
	SaveSchedule(ctx context.Context, body model.SaveScheduleRequest) (*http.Response, error)

	//Search Find all the executions for the given schedule
	Search(ctx context.Context, localVarOptionals *SchedulerSearchOpts) (model.SearchResultWorkflowSchedule, *http.Response, error)
}

func GetSchedulerService(client *APIClient) SchedulerClient {
	return &SchedulerResourceApiService{client}
}
