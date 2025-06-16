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

func TestSecretResourceApiService(t *testing.T) {
	// Setup
	secretClient := testdata.SecretClient // Assuming this exists in your testdata package
	ctx := context.Background()

	// Generate a unique secret key for testing
	secretKey := fmt.Sprintf("test-secret-%d", time.Now().UnixNano())
	secretValue := "this-is-a-test-secret-value"

	// Test case 1: Put a new secret
	_, resp, err := secretClient.PutSecret(ctx, secretValue, secretKey)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test case 2: Check if secret exists
	existsResult, resp, err := secretClient.SecretExists(ctx, secretKey)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, existsResult)

	// Test case 3: Get the secret value
	secretRetrieved, resp, err := secretClient.GetSecret(ctx, secretKey)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, secretValue, secretRetrieved)

	// Test case 4: Add tags to the secret
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

	resp, err = secretClient.PutTagForSecret(ctx, tags, secretKey)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test case 5: Get tags for the secret
	retrievedTags, resp, err := secretClient.GetTags(ctx, secretKey)
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

	// Test case 6: List all secret names
	allSecretNames, resp, err := secretClient.ListAllSecretNames(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Check if our secret is in the list
	var foundSecret bool
	for _, name := range allSecretNames {
		if name == secretKey {
			foundSecret = true
			break
		}
	}
	assert.True(t, foundSecret, "Created secret not found in list of all secrets")

	// Test case 7: List secrets that user can grant access to
	accessibleSecrets, resp, err := secretClient.ListSecretsThatUserCanGrantAccessTo(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, accessibleSecrets)
	assert.Equal(t, len(allSecretNames), len(accessibleSecrets))

	// Test case 8: List secrets with tags that user can grant access to
	secretsWithTags, resp, err := secretClient.ListSecretsWithTagsThatUserCanGrantAccessTo(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Check if our secret is in the list with correct tags
	foundSecretWithTags := false
	for _, secret := range secretsWithTags {
		if secret.Name == secretKey {
			foundSecretWithTags = true

			// Verify tags are present
			foundEnvironmentTag = false
			foundOwnerTag = false
			for _, tag := range secret.Tags {
				if tag.Key == "environment" && tag.Value == "test" {
					foundEnvironmentTag = true
				}
				if tag.Key == "owner" && tag.Value == "integration-test" {
					foundOwnerTag = true
				}
			}
			assert.True(t, foundEnvironmentTag, "Environment tag not found in secret with tags")
			assert.True(t, foundOwnerTag, "Owner tag not found in secret with tags")

			break
		}
	}
	assert.True(t, foundSecretWithTags, "Created secret not found in list of secrets with tags")

	// Test case 9: Delete a tag from the secret
	tagToDelete := []model.Tag{
		{
			Key:   "environment",
			Value: "test",
			Type_: "metadata",
		},
	}

	resp, err = secretClient.DeleteTagForSecret(ctx, tagToDelete, secretKey)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test case 10: Verify tag was deleted
	updatedTags, resp, err := secretClient.GetTags(ctx, secretKey)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	environmentTagFound := false
	for _, tag := range updatedTags {
		if tag.Key == "environment" && tag.Value == "test" {
			environmentTagFound = true
			break
		}
	}
	assert.False(t, environmentTagFound, "Environment tag should have been deleted")

	// Test case 11: Test cache clearing operations (these usually don't return meaningful results in tests)
	_, resp, err = secretClient.ClearLocalCache(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	_, resp, err = secretClient.ClearRedisCache(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test case 12: Delete the secret
	_, resp, err = secretClient.DeleteSecret(ctx, secretKey)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test case 13: Verify the secret was deleted
	_, resp, err = secretClient.GetSecret(ctx, secretKey)
	assert.NotNil(t, err)
	assert.Equal(t, 404, resp.StatusCode) // Assuming 404 is returned for non-existent secret
}
