package client

import (
	"context"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"net/http"
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
