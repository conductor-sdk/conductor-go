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

	//PutTagForWebhook Add Tag to webhook
	PutTagForWebhook(ctx context.Context, body []model.Tag, id string) (*http.Response, error)

	//GetTagsForWebhook Get all tags for webhook
	GetTagsForWebhook(ctx context.Context, id string) ([]model.Tag, *http.Response, error)

	//DeleteTagForWebhook Delete Tag from webhook
	DeleteTagForWebhook(ctx context.Context, id string, body []model.Tag) (*http.Response, error)
}

func NewWebhooksConfigClient(client *APIClient) WebhooksConfigClient {
	return &WebhooksConfigResourceApiService{client}
}
