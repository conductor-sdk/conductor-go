package client

import (
	"context"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"net/http"
)

type WebhooksConfigClient interface {
	//CreateWebhook create and register a new webhook
	CreateWebhook(ctx context.Context, body model.WebhookConfig) (model.WebhookConfig, *http.Response, error)

	//DeleteWebhook delete the webhook
	DeleteWebhook(ctx context.Context, id string) (*http.Response, error)

	//GetAllWebhook return all the webhooks
	GetAllWebhook(ctx context.Context) ([]model.WebhookConfig, *http.Response, error)

	//GetWebhook Get a specific webhook
	GetWebhook(ctx context.Context, id string) (model.WebhookConfig, *http.Response, error)

	//UpdateWebhook Update webhook
	UpdateWebhook(ctx context.Context, body model.WebhookConfig, id string) (model.WebhookConfig, *http.Response, error)
}

func GetWebhooksConfigService(client *APIClient) WebhooksConfigClient {
	return &WebhooksConfigResourceApiService{client}
}
