package integration_tests

import (
	"context"
	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/model/integration"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestIntegrationClient(t *testing.T) {
	ctx := context.Background()

	// Instantiate the IntegrationClient
	// Note: Ensure that NewIntegrationClient is correctly implemented to connect to a test environment
	integrationClient := NewIntegrationClient() // Adjust the function name and parameters as necessary
	promptClient := NewPromptClient()

	// Insert a few integration providers before retrieving

	integrationEntries := []integration.IntegrationUpdate{
		{
			Category:      "AI_MODEL",
			Configuration: map[integration.ConfigKey]interface{}{integration.APIKey: "value1", integration.Password: "password1"},
			Description:   "First example AI model provider.",
			Enabled:       true,
			Type_:         "azure_openai",
		},
		{
			Category:      "AI_MODEL",
			Configuration: map[integration.ConfigKey]interface{}{integration.APIKey: "value2", integration.Password: "password2"},
			Description:   "Second example AI model provider.",
			Enabled:       true,
			Type_:         "cohere",
		},
	}
	names := []string{"test_integration_1", "test_integration_2"}

	for i, entry := range integrationEntries {
		resp, err := integrationClient.SaveIntegrationProvider(ctx, entry, names[i])
		require.NoError(t, err, "error saving integration provider")
		require.NotNil(t, resp, "response should not be nil for SaveIntegrationProvider")
	}

	// Testing GetIntegrationProviders with some filters if applicable
	providers, resp, err := integrationClient.GetIntegrationProviders(ctx, nil)
	require.NoError(t, err, "error fetching integration providers")
	require.NotNil(t, resp, "response should not be nil for GetIntegrationProviders")
	require.GreaterOrEqual(t, len(providers), len(integrationEntries), "the number of providers fetched should match the entries inserted")

	// Testing GetIntegrationProvider for each inserted entry
	for i, entry := range integrationEntries {
		provider, resp, err := integrationClient.GetIntegrationProvider(ctx, names[i])
		require.NoError(t, err, "error fetching a specific integration provider")
		require.NotNil(t, resp, "response should not be nil for GetIntegrationProvider")
		require.Equal(t, entry.Category, provider.Category, "category should match the inserted provider")
	}

	apiUpdate := integration.IntegrationApiUpdate{
		Description: "A test API for integration",
		Enabled:     true,
	}

	providerName := names[0]
	apiModel := "DefaultModel"
	promptName := "TestPrompt"
	description := "greetings"
	opts := client.PromptResourceApiSaveMessageTemplateOpts{Models: []string{providerName + ":" + apiModel}}
	promptClient.SaveMessageTemplate(ctx, "Say hello to ${name}", description, promptName, &opts)

	promptTemplate, resp, err := promptClient.GetMessageTemplate(ctx, promptName)
	require.NoError(t, err)
	require.NotNil(t, resp, "response should not be nil for GetMessageTemplate")
	require.Equal(t, resp.StatusCode, 200)
	require.NotNil(t, promptTemplate)
	require.Equal(t, promptName, promptTemplate.Name)
	require.Equal(t, description, promptTemplate.Description)

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

	// Add Tags on prompt template
	resp, err = promptClient.PutTagForPromptTemplate(ctx, tags, promptName)
	require.NoError(t, err)
	require.NotNil(t, resp, "response should not be nil for PutTagForPromptTemplate")
	require.Equal(t, resp.StatusCode, 200)

	// Get Tags on prompt template
	tags, resp, err = promptClient.GetTagsForPromptTemplate(ctx, promptName)
	require.NoError(t, err)
	require.NotNil(t, resp, "response should not be nil for GetTagsForPromptTemplate")
	require.Equal(t, resp.StatusCode, 200)
	require.Equal(t, len(tags), 2)

	// Delete Tag on prompt template
	resp, err = promptClient.DeleteTagForPromptTemplate(ctx, tags, promptName)
	require.NoError(t, err)
	require.NotNil(t, resp, "response should not be nil for DeleteTagForPromptTemplate")
	require.Equal(t, resp.StatusCode, 200)

	// Get Tags on prompt template
	tags, resp, err = promptClient.GetTagsForPromptTemplate(ctx, promptName)
	require.NoError(t, err)
	require.NotNil(t, resp, "response should not be nil for GetTagsForPromptTemplate")
	require.Equal(t, resp.StatusCode, 200)
	require.Equal(t, len(tags), 0)

	// Create an Integration API
	_, err = integrationClient.SaveIntegrationApi(ctx, apiUpdate, providerName, apiModel)
	require.NoError(t, err, "Failed to save integration API")

	// Retrieve the Integration API
	api, resp, err := integrationClient.GetIntegrationApi(ctx, providerName, apiModel)
	require.NoError(t, err, "Failed to retrieve integration API")
	require.NotNil(t, resp, "Response should not be nil")
	require.Equal(t, apiModel, api.Api, "API name should match the saved API")

	// Retrieve all Integration APIs for the provider
	apis, resp, err := integrationClient.GetIntegrationApis(ctx, providerName, optional.NewBool(true))
	require.NoError(t, err, "Failed to retrieve integration APIs")
	require.NotNil(t, resp, "Response should not be nil")
	require.Contains(t, apis, api, "Retrieved APIs should contain the saved API")

	// Test Prompts with Integration
	// Associate a prompt with the integration
	_, err = integrationClient.AssociatePromptWithIntegration(ctx, providerName, apiModel, promptName)
	require.NoError(t, err, "Failed to associate prompt with integration")

	// Retrieve prompts associated with the integration
	prompts, resp, err := integrationClient.GetPromptsWithIntegration(ctx, providerName, apiModel)
	require.NoError(t, err, "Failed to get prompts with integration")
	require.NotNil(t, resp, "Response should not be nil")
	require.NotEmpty(t, prompts, "Expected non-empty list of prompts")

	integrations, resp, err := testdata.IntegrationClient.GetAllIntegrations(
		context.Background(), &client.IntegrationResourceApiGetAllIntegrationsOpts{
			ActiveOnly: optional.NewBool(true),
		})

	if err != nil {
		t.Fatalf("Failed to get integrations. Reason: %s", err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, got %d", resp.StatusCode)
	}
	for i, integration := range integrations {
		if !integration.Enabled {
			t.Fatalf("Integration #%d (%s) is not active, but should be", i, integration.Name)
		}
	}
	require.Equal(t, 3, len(integrations))

	for _, integration := range integrations {
		require.NotNil(t, integration)
		require.True(t, integration.Enabled)
		require.Equal(t, "AI_MODEL", integration.Category)
	}

	integrationDefs, resp, err := testdata.IntegrationClient.GetIntegrationProviderDefs(ctx)
	require.NoError(t, err, "Failed to retrieve integration providers")
	require.NotNil(t, resp, "Response should not be nil")
	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Equal(t, 29, len(integrationDefs))

	promptClient.DeleteMessageTemplate(ctx, promptName)
	template, res, err := promptClient.GetMessageTemplate(ctx, promptName)
	require.NotNil(t, err)
	require.Nil(t, template)
	require.Equal(t, 404, res.StatusCode)

	// Delete the Integration API
	_, err = integrationClient.DeleteIntegrationApi(ctx, providerName, apiModel)
	require.NoError(t, err, "Failed to delete integration API")

	// Cleanup: Deleting providers to clean the test environment
	for i := range integrationEntries {
		resp, err = integrationClient.DeleteIntegrationProvider(ctx, names[i])
		require.NoError(t, err, "error deleting integration provider")
		require.NotNil(t, resp, "response should not be nil for DeleteIntegrationProvider")
	}
}

func NewIntegrationClient() client.IntegrationClient {
	return testdata.IntegrationClient
}
func NewPromptClient() client.PromptClient {
	return testdata.PromptClient
}
