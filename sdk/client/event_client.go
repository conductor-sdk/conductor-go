package client

import (
	"context"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"net/http"
)

type EventClient interface {
	AddEventHandler(ctx context.Context, body model.EventHandler) (*http.Response, error)
	DeleteQueueConfig(ctx context.Context, queueType string, queueName string) (*http.Response, error)
	GetEventHandlers(ctx context.Context) ([]model.EventHandler, *http.Response, error)
	GetEventHandlersForEvent(ctx context.Context, event string, localVarOptionals *EventResourceApiGetEventHandlersForEventOpts) ([]model.EventHandler, *http.Response, error)
	GetQueueConfig(ctx context.Context, queueType string, queueName string) (map[string]interface{}, *http.Response, error)
	GetQueueNames(ctx context.Context) (map[string]string, *http.Response, error)
	PutQueueConfig(ctx context.Context, body string, queueType string, queueName string) (*http.Response, error)
	RemoveEventHandlerStatus(ctx context.Context, name string) (*http.Response, error)
	UpdateEventHandler(ctx context.Context, body model.EventHandler) (*http.Response, error)
}
