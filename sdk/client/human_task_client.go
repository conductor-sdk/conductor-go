// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package client

import (
	"context"
	"net/http"

	"github.com/conductor-sdk/conductor-go/sdk/model/human"
)

type HumanTaskClient interface {
	AssignAndClaim(ctx context.Context, taskId string, userId string, optionals *HumanTaskApiAssignAndClaimOpts) (human.HumanTaskEntry, *http.Response, error)
	BackPopulateFullTextIndex(ctx context.Context, var100 int32) (map[string]interface{}, *http.Response, error)
	ClaimTask(ctx context.Context, taskId string, optionals *HumanTaskApiClaimTaskOpts) (human.HumanTaskEntry, *http.Response, error)
	DeleteTaskFromHumanTaskRecords(ctx context.Context, body []string) (*http.Response, error)
	DeleteTaskFromHumanTaskRecords1(ctx context.Context, taskId string) (*http.Response, error)
	DeleteTemplateByName(ctx context.Context, name string) (*http.Response, error)
	DeleteTemplatesByNameAndVersion(ctx context.Context, name string, version int32) (*http.Response, error)
	GetAllTemplates(ctx context.Context, optionals *HumanTaskApiGetAllTemplatesOpts) ([]human.HumanTaskSearch, *http.Response, error)
	GetTask1(ctx context.Context, taskId string, optionals *HumanTaskApiGetTask1Opts) (human.HumanTaskEntry, *http.Response, error)
	GetTaskDisplayNames(ctx context.Context, searchType string) ([]string, *http.Response, error)
	GetTemplateByNameAndVersion(ctx context.Context, name string, version int32) (human.HumanTaskSearch, *http.Response, error)
	GetTemplateByTaskId(ctx context.Context, humanTaskId string) (human.HumanTaskSearch, *http.Response, error)
	ReassignTask(ctx context.Context, body []human.HumanTaskAssignment, taskId string) (*http.Response, error)
	ReleaseTask(ctx context.Context, taskId string) (*http.Response, error)
	SaveTemplate(ctx context.Context, body human.HumanTaskSearch, optionals *HumanTaskApiSaveTemplateOpts) (human.HumanTaskSearch, *http.Response, error)
	SaveTemplates(ctx context.Context, body []human.HumanTaskSearch, optionals *HumanTaskApiSaveTemplatesOpts) ([]human.HumanTaskSearch, *http.Response, error)
	Search(ctx context.Context, body human.HumanTaskSearch) (human.HumanTaskSearchResult, *http.Response, error)
	SkipTask(ctx context.Context, taskId string, optionals *HumanTaskApiSkipTaskOpts) (*http.Response, error)
	UpdateTaskOutput(ctx context.Context, body map[string]interface{}, taskId string, optionals *HumanTaskApiUpdateTaskOutputOpts) (*http.Response, error)
	UpdateTaskOutputByRef(ctx context.Context, body map[string]interface{}, workflowId string, taskRefName string, optionals *HumanTaskApiUpdateTaskOutputByRefOpts) (*http.Response, error)
}

func NewHumanTaskClient(client *APIClient) HumanTaskClient {
	return &HumanTaskApiService{client}
}
