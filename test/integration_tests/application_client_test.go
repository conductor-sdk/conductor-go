package integration_tests

import (
	"context"

	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/model/rbac"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"github.com/stretchr/testify/assert"
)

func TestApplicationLifecycle(t *testing.T) {

	appClient := testdata.ApplicationClient

	// Create an application
	ctx := context.Background()
	createReq := rbac.CreateOrUpdateApplicationRequest{Name: "TestApp2"}
	createdApp, resp, err := appClient.CreateApplication(ctx, createReq)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "TestApp2", createdApp.Name)

	// Retrieve the created application
	retrievedApp, resp, err := appClient.GetApplication(ctx, createdApp.Id)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "TestApp2", retrievedApp.Name)

	// Delete the application
	_, resp, err = appClient.DeleteApplication(ctx, createdApp.Id)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify the application is deleted (this step may vary based on how your API handles deletions)
	_, resp, _ = appClient.GetApplication(ctx, createdApp.Id)
	assert.Equal(t, 404, resp.StatusCode)
}

func TestRoleManagementForApplicationUser(t *testing.T) {

	appClient := testdata.ApplicationClient

	// Create an application to use in the test
	ctx := context.Background()
	createReq := rbac.CreateOrUpdateApplicationRequest{Name: "TestAppRoleUser"}
	application, _, _ := appClient.CreateApplication(ctx, createReq)

	// Add a role to the application user
	_, resp, err := appClient.AddRoleToApplicationUser(ctx, application.Id, "ADMIN")
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Remove the role from the application user
	_, resp, err = appClient.RemoveRoleFromApplicationUser(ctx, application.Id, "ADMIN")
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Cleanup
	_, _, _ = appClient.DeleteApplication(ctx, application.Id)
}

func TestAccessKeyLifecycle(t *testing.T) {
	appClient := testdata.ApplicationClient

	// Create an application to use in the test
	ctx := context.Background()
	createReq := rbac.CreateOrUpdateApplicationRequest{Name: "TestAppAccessKey"}
	application, _, _ := appClient.CreateApplication(ctx, createReq)

	// Create an access key for the application
	accessKey, resp, err := appClient.CreateAccessKey(ctx, application.Id)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Delete the access key
	resp, err = appClient.DeleteAccessKey(ctx, application.Id, accessKey.Id)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Cleanup
	_, _, _ = appClient.DeleteApplication(ctx, application.Id)
}

func TestGetTagsForApplication(t *testing.T) {

	appClient := testdata.ApplicationClient

	// Create an application to use in the test
	ctx := context.Background()
	createReq := rbac.CreateOrUpdateApplicationRequest{Name: "TestAppTags"}
	application, _, _ := appClient.CreateApplication(ctx, createReq)

	// Assuming tags are added here or exists; this could use PutTagForApplication if needed

	// Get tags for the application
	tags, resp, err := appClient.GetTagsForApplication(ctx, application.Id)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	// Additional assertions based on expected tags
	assert.Equal(t, 0, len(tags))

	// Cleanup
	_, _, _ = appClient.DeleteApplication(ctx, application.Id)
}

func TestApplicationClientIntegration(t *testing.T) {
	// Initialize ApplicationClient
	appClient := testdata.ApplicationClient

	// Context used in all requests
	ctx := context.Background()

	// Create an application
	createReq := rbac.CreateOrUpdateApplicationRequest{Name: "IntegrationTestApp"}
	application, _, err := appClient.CreateApplication(ctx, createReq)
	assert.Nil(t, err)

	// Get the application
	gotApp, _, err := appClient.GetApplication(ctx, application.Id)
	assert.Nil(t, err)
	assert.Equal(t, application.Id, gotApp.Id)

	// List applications
	apps, _, err := appClient.ListApplications(ctx)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(apps), 1)

	// Add a tag to the application
	tags := []model.Tag{{Key: "env", Value: "development"}}
	_, err = appClient.PutTagForApplication(ctx, tags, application.Id)
	assert.Nil(t, err)

	// Get tags for the application
	retrievedTags, _, err := appClient.GetTagsForApplication(ctx, application.Id)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(retrievedTags))

	// Create an access key for the application
	accessKey, _, err := appClient.CreateAccessKey(ctx, application.Id)
	assert.Nil(t, err)

	// Toggle the access key status
	_, _, err = appClient.ToggleAccessKeyStatus(ctx, application.Id, accessKey.Id)
	assert.Nil(t, err)

	// Remove the added tag
	_, err = appClient.DeleteTagForApplication(ctx, tags, application.Id)
	assert.Nil(t, err)

	// Update the application
	updateReq := rbac.CreateOrUpdateApplicationRequest{Name: "UpdatedIntegrationTestApp"}
	updatedApp, _, err := appClient.UpdateApplication(ctx, updateReq, application.Id)
	assert.Nil(t, err)
	assert.Equal(t, "UpdatedIntegrationTestApp", updatedApp.Name)

	// Delete the access key
	_, err = appClient.DeleteAccessKey(ctx, application.Id, accessKey.Id)
	assert.Nil(t, err)

	// Finally, delete the application
	_, _, err = appClient.DeleteApplication(ctx, application.Id)
	assert.Nil(t, err)
}

func TestApplicationClientErrorHandling(t *testing.T) {
	// Initialize ApplicationClient
	appClient := testdata.ApplicationClient

	// Context used in all requests
	ctx := context.Background()

	// Define an invalid application ID
	invalidAppId := "nonexistent"

	// Try to get a non-existent application
	_, resp, err := appClient.GetApplication(ctx, invalidAppId)
	assert.NotNil(t, err)
	assert.Equal(t, 404, resp.StatusCode)

	// Try to update a non-existent application
	updateReq := rbac.CreateOrUpdateApplicationRequest{Name: "NonExistentApp"}
	_, resp, err = appClient.UpdateApplication(ctx, updateReq, invalidAppId)
	assert.NotNil(t, err)
	assert.Equal(t, 404, resp.StatusCode)

	// Try to delete a non-existent application
	_, resp, err = appClient.DeleteApplication(ctx, invalidAppId)
	assert.NotNil(t, err)
	assert.Equal(t, 404, resp.StatusCode)

	// Try to add a tag to a non-existent application
	tags := []model.Tag{{Key: "env", Value: "staging"}}
	res, err := appClient.PutTagForApplication(ctx, tags, invalidAppId)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, 200, res.StatusCode) //expected
	// Assume error handling in client converts HTTP 404 to specific Go error; response might not be accessible directly

	// Try to get tags for a non-existent application
	_, _, err = appClient.GetTagsForApplication(ctx, invalidAppId)
	assert.Nil(t, err)

	// Try to delete a tag from a non-existent application
	_, err = appClient.DeleteTagForApplication(ctx, tags, invalidAppId)
	assert.Nil(t, err)

	// Try to create an access key for a non-existent application
	_, resp, err = appClient.CreateAccessKey(ctx, invalidAppId)
	assert.NotNil(t, err)
	assert.Equal(t, 404, resp.StatusCode)

	// Try to toggle access key status for a non-existent application
	_, _, err = appClient.ToggleAccessKeyStatus(ctx, invalidAppId, "fakeKey")
	assert.NotNil(t, err)
}

func TestGetAccessKeys(t *testing.T) {
	// Initialize ApplicationClient
	appClient := testdata.ApplicationClient

	// Context used in all requests
	ctx := context.Background()

	// Create an application to use in the test
	createReq := rbac.CreateOrUpdateApplicationRequest{Name: "TestAppAccessKeys"}
	application, _, err := appClient.CreateApplication(ctx, createReq)
	assert.Nil(t, err)

	// Initially check access keys when none are added
	keysBefore, resp, err := appClient.GetAccessKeys(ctx, application.Id)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Empty(t, keysBefore)

	// Create an access key
	newKey, _, err := appClient.CreateAccessKey(ctx, application.Id)
	assert.Nil(t, err)

	// Retrieve the access keys and check the list contains the new key
	keysAfter, resp, err := appClient.GetAccessKeys(ctx, application.Id)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, keysAfter)
	assert.Equal(t, newKey.Id, keysAfter[0].Id) // Assuming the client returns keys with IDs and the list has this structure

	// Cleanup - delete the access key and application
	_, err = appClient.DeleteAccessKey(ctx, application.Id, newKey.Id)
	assert.Nil(t, err)

	keysAfter, resp, err = appClient.GetAccessKeys(ctx, application.Id)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Empty(t, keysAfter)

	_, _, err = appClient.DeleteApplication(ctx, application.Id)
	assert.Nil(t, err)

	retrievedApp, resp, err := appClient.GetApplication(ctx, application.Id)
	assert.NotNil(t, err)
	assert.Equal(t, 404, resp.StatusCode)
	assert.Nil(t, retrievedApp)
}

// Implement other tests if there are more methods in ApplicationClient
