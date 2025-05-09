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
	"fmt"
	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"net/http"
	"net/url"
)

type EventResourceApiService struct {
	*APIClient
}

/*
EventResourceApiService Add a new event handler.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
*/
func (a *EventResourceApiService) AddEventHandler(ctx context.Context, body model.EventHandler) (*http.Response, error) {
	resp, err := a.Post(ctx, "/event", body, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
EventResourceApiService Delete queue config by name
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param queueType
  - @param queueName
*/
func (a *EventResourceApiService) DeleteQueueConfig(ctx context.Context, queueType string, queueName string) (*http.Response, error) {
	path := fmt.Sprintf("/event/queue/config/%s/%s", queueType, queueName)
	resp, err := a.Delete(ctx, path, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
EventResourceApiService Get all the event handlers
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return []model.EventHandler
*/
func (a *EventResourceApiService) GetEventHandlers(ctx context.Context) ([]model.EventHandler, *http.Response, error) {
	var result []model.EventHandler
	resp, err := a.Get(ctx, "/event", nil, &result)

	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

/*
EventResourceApiService Get event handlers for a given event
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param event
 * @param optional nil or *EventResourceApiGetEventHandlersForEventOpts - Optional Parameters:
     * @param "ActiveOnly" (optional.Bool) -
@return []model.EventHandler
*/

type EventResourceApiGetEventHandlersForEventOpts struct {
	ActiveOnly optional.Bool
}

func (a *EventResourceApiService) GetEventHandlersForEvent(ctx context.Context, event string, opts *EventResourceApiGetEventHandlersForEventOpts) ([]model.EventHandler, *http.Response, error) {
	var result []model.EventHandler
	path := fmt.Sprintf("/event/%s", event)

	// Build query parameters
	queryParams := url.Values{}
	if opts != nil && opts.ActiveOnly.IsSet() {
		queryParams.Add("activeOnly", parameterToString(opts.ActiveOnly.Value(), ""))
	}

	resp, err := a.Get(ctx, path, queryParams, &result)

	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

/*
EventResourceApiService Get queue config by name
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param queueType
  - @param queueName

@return map[string]interface{}
*/
func (a *EventResourceApiService) GetQueueConfig(ctx context.Context, queueType string, queueName string) (map[string]interface{}, *http.Response, error) {
	var result map[string]interface{}
	path := fmt.Sprintf("/event/queue/config/%s/%s", queueType, queueName)
	resp, err := a.Get(ctx, path, nil, &result)

	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

/*
EventResourceApiService Get all queue configs
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return map[string]string
*/
func (a *EventResourceApiService) GetQueueNames(ctx context.Context) (map[string]string, *http.Response, error) {
	var result map[string]string
	resp, err := a.Get(ctx, "/event/queue/config", nil, &result)

	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

/*
EventResourceApiService Create or update queue config by name
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param queueType
  - @param queueName
*/
func (a *EventResourceApiService) PutQueueConfig(ctx context.Context, body string, queueType string, queueName string) (*http.Response, error) {
	path := fmt.Sprintf("/event/queue/config/%s/%s", queueType, queueName)

	resp, err := a.Put(ctx, path, body, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
EventResourceApiService Remove an event handler
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name
*/
func (a *EventResourceApiService) RemoveEventHandler(ctx context.Context, name string) (*http.Response, error) {
	path := fmt.Sprintf("/event/%s", name)
	resp, err := a.Delete(ctx, path, nil, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

/*
EventResourceApiService Update an existing event handler.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
*/
func (a *EventResourceApiService) UpdateEventHandler(ctx context.Context, body model.EventHandler) (*http.Response, error) {
	resp, err := a.Put(ctx, "/event", body, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
