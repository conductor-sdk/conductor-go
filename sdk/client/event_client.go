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

	"github.com/conductor-sdk/conductor-go/sdk/model"
)

type EventHandlerClient interface {
	AddEventHandler(ctx context.Context, body model.EventHandler) (*http.Response, error)
	GetEventHandlers(ctx context.Context) ([]model.EventHandler, *http.Response, error)
	GetEventHandlersForEvent(ctx context.Context, event string, localVarOptionals *EventResourceApiGetEventHandlersForEventOpts) ([]model.EventHandler, *http.Response, error)
	RemoveEventHandler(ctx context.Context, name string) (*http.Response, error)
	UpdateEventHandler(ctx context.Context, body model.EventHandler) (*http.Response, error)
}

func NewEventHandlerClient(client *APIClient) EventHandlerClient {
	return &EventResourceApiService{client}
}
