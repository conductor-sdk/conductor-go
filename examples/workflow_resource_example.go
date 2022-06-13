//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package examples

import (
	"context"
	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
	"net/http"

	"github.com/antihax/optional"
)

func StartWorkflow(workflowName string, version optional.Int32, correlationId optional.String) (string, *http.Response, error) {
	workflowClient := client.WorkflowResourceApiService{
		APIClient: client.NewAPIClient(
			settings.NewAuthenticationSettings(
				"key",
				"id",
			),
			settings.NewHttpDefaultSettings(),
		),
	}
	return workflowClient.StartWorkflow(
		context.Background(),
		make(map[string]interface{}),
		workflowName,
		&client.WorkflowResourceApiStartWorkflowOpts{
			Version:       version,
			CorrelationId: correlationId,
		},
	)
}
