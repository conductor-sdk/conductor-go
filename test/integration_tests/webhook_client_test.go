package integration_tests

import (
	"context"
	"fmt"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestWebhooksConfigResourceApiService(t *testing.T) {
	// Setup
	webhookClient := testdata.WebhookClient // Assuming this exists in your testdata package
	ctx := context.Background()

	// Generate unique identifiers for testing
	webhookName := fmt.Sprintf("test-webhook-%d", time.Now().UnixNano())
	webhookId := fmt.Sprintf("webhook-%d", time.Now().UnixNano())

	// Test case 1: Create a new webhook
	webhookConfig := model.WebhookConfig{
		Id:             webhookId,
		Name:           webhookName,
		SourcePlatform: "GITHUB",
		HeaderKey:      "X-GitHub-Event",
		SecretKey:      "X-Hub-Signature",
		SecretValue:    "test-secret",
		Headers:        map[string]string{"Content-Type": "application/json"},
		Verifier:       "SLACK_BASED",
		WorkflowsToStart: map[string]int32{
			"TestWorkflow": 1, // Workflow name and version
		},
	}

	createdWebhook, resp, err := webhookClient.CreateWebhook(ctx, webhookConfig)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, webhookId, createdWebhook.Id)
	assert.Equal(t, webhookName, createdWebhook.Name)

	// Add cleanup to delete the webhook at the end
	defer func() {
		_, _ = webhookClient.DeleteWebhook(ctx, webhookId)
	}()

	// Test case 2: Get the webhook by ID
	retrievedWebhook, resp, err := webhookClient.GetWebhook(ctx, webhookId)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, webhookId, retrievedWebhook.Id)
	assert.Equal(t, webhookName, retrievedWebhook.Name)
	assert.Equal(t, "GITHUB", retrievedWebhook.SourcePlatform)

	// Test case 3: Get all webhooks and verify our webhook exists
	allWebhooks, resp, err := webhookClient.GetAllWebhook(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var foundWebhook bool
	for _, webhook := range allWebhooks {
		if webhook.Id == webhookId {
			foundWebhook = true
			assert.Equal(t, webhookName, webhook.Name)
			break
		}
	}
	assert.True(t, foundWebhook, "Created webhook not found in list of all webhooks")

	// Test case 4: Add tags to the webhook
	tags := []model.Tag{
		{
			Key:   "environment",
			Value: "test",
			Type_: "metadata",
		},
		{
			Key:   "owner",
			Value: "integration-test",
			Type_: "ownership",
		},
	}

	resp, err = webhookClient.PutTagForWebhook(ctx, tags, webhookId)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test case 5: Get tags for the webhook
	retrievedTags, resp, err := webhookClient.GetTagsForWebhook(ctx, webhookId)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify tags were added correctly
	assert.GreaterOrEqual(t, len(retrievedTags), 2)
	var foundEnvironmentTag, foundOwnerTag bool
	for _, tag := range retrievedTags {
		if tag.Key == "environment" && tag.Value == "test" {
			foundEnvironmentTag = true
		}
		if tag.Key == "owner" && tag.Value == "integration-test" {
			foundOwnerTag = true
		}
	}
	assert.True(t, foundEnvironmentTag, "Environment tag not found")
	assert.True(t, foundOwnerTag, "Owner tag not found")

	// Test case 6: Update the webhook
	updatedWebhook := retrievedWebhook
	updatedWebhook.SourcePlatform = "GITLAB"
	updatedWebhook.HeaderKey = "X-Gitlab-Event"

	// Add another workflow to start
	if updatedWebhook.WorkflowsToStart == nil {
		updatedWebhook.WorkflowsToStart = make(map[string]int32)
	}
	updatedWebhook.WorkflowsToStart["AnotherTestWorkflow"] = 1

	updatedWebhookResult, resp, err := webhookClient.UpdateWebhook(ctx, updatedWebhook, webhookId)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "GITLAB", updatedWebhookResult.SourcePlatform)
	assert.Equal(t, "X-Gitlab-Event", updatedWebhookResult.HeaderKey)
	assert.Contains(t, updatedWebhookResult.WorkflowsToStart, "AnotherTestWorkflow")

	// Test case 7: Verify update by getting the webhook again
	verifyWebhook, resp, err := webhookClient.GetWebhook(ctx, webhookId)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "GITLAB", verifyWebhook.SourcePlatform)
	assert.Equal(t, "X-Gitlab-Event", verifyWebhook.HeaderKey)
	assert.Contains(t, verifyWebhook.WorkflowsToStart, "AnotherTestWorkflow")

	// Test case 10: Delete the webhook
	resp, err = webhookClient.DeleteWebhook(ctx, webhookId)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test case 11: Verify the webhook was deleted
	allWebhooksAfterDelete, resp, err := webhookClient.GetAllWebhook(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	webhookFoundAfterDelete := false
	for _, webhook := range allWebhooksAfterDelete {
		if webhook.Id == webhookId {
			webhookFoundAfterDelete = true
			break
		}
	}
	assert.False(t, webhookFoundAfterDelete, "Webhook should have been deleted")
}
