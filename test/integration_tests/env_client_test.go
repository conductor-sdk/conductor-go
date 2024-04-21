package integration_tests

import (
	"context"
	"github.com/conductor-sdk/conductor-go/sdk/authentication"
	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestCreateOrUpdateEnvVariable(t *testing.T) {
	ctx := context.Background()
	envClient := NewEnvironmentClient()

	resp, err := envClient.CreateOrUpdateEnvVariable(ctx, "test value", "testKey")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	resp, err = envClient.CreateOrUpdateEnvVariable(ctx, "", "") // Edge case with empty values
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestDeleteEnvVariable(t *testing.T) {
	TestCreateOrUpdateEnvVariable(t)
	ctx := context.Background()
	envClient := NewEnvironmentClient()

	message, resp, err := envClient.DeleteEnvVariable(ctx, "testKey")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "", message)

}

func TestDeleteTagForEnvVar(t *testing.T) {
	TestCreateOrUpdateEnvVariable(t)
	ctx := context.Background()
	envClient := NewEnvironmentClient()
	tags := []model.Tag{{Key: "tag1", Value: "value1"}}

	resp, err := envClient.DeleteTagForEnvVar(ctx, tags, "envVarName")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

}

func TestGetEnvVariable(t *testing.T) {
	ctx := context.Background()
	envClient := NewEnvironmentClient()

	value, resp, err := envClient.Get(ctx, "testKey")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, value)

}

func TestGetAllEnvVariables(t *testing.T) {
	TestCreateOrUpdateEnvVariable(t)
	ctx := context.Background()
	envClient := NewEnvironmentClient()

	variables, resp, err := envClient.GetAll(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Greater(t, len(variables), 0) // Expecting at least one variable
}

func TestGetTagsForEnvVar(t *testing.T) {
	TestUpsertUser(t)
	ctx := context.Background()
	envClient := NewEnvironmentClient()
	tags := []model.Tag{{Key: "tag1", Value: "value1"}}
	envClient.PutTagForEnvVar(ctx, tags, "envVarName")
	tags, resp, err := envClient.GetTagsForEnvVar(ctx, "envVarName")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Greater(t, len(tags), 0)

}

func TestPutTagForEnvVar(t *testing.T) {
	ctx := context.Background()
	envClient := NewEnvironmentClient()
	tags := []model.Tag{{Key: "tag1", Value: "value1"}}

	resp, err := envClient.PutTagForEnvVar(ctx, tags, "envVarName")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

}

func NewEnvironmentClient() client.EnvironmentClient {
	// Your code to initialize and return a client that implements UserClient
	// This is typically setup for integration testing with real API or a mock server
	apiClient := client.NewAPIClientWithTokenExpiration(
		settings.NewAuthenticationSettings("api_key_user_03", "api_key_user_03"),
		settings.NewHttpSettings("http://localhost:8080/api"),
		authentication.NewTokenExpiration(
			3*time.Second,
			30*time.Second,
		),
	)
	return client.NewEnvironmentClient(apiClient)
}
