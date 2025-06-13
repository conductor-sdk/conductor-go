package integration_tests

import (
	"context"
	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

const (
	HTTPServiceName = "http-service"
	GRPCServiceName = "grpc-service"
	ProtoFilename   = "compiled.bin"
)

type ServiceRegistryClientTestSuite struct {
	client client.ServiceRegistryClient
	ctx    context.Context
}

func setupServiceRegistryTest(t *testing.T) *ServiceRegistryClientTestSuite {
	// Initialize your client here - adjust based on your client setup
	serviceClient := testdata.ServiceRegistryClient

	suite := &ServiceRegistryClientTestSuite{
		client: serviceClient,
		ctx:    context.Background(),
	}

	// Cleanup before test
	suite.cleanup(t)

	return suite
}

func (s *ServiceRegistryClientTestSuite) cleanup(t *testing.T) {
	// Remove services if they exist (ignore errors)
	s.client.RemoveService(s.ctx, HTTPServiceName)
	s.client.RemoveService(s.ctx, GRPCServiceName)
}

func TestServiceRegistryClient_HTTPService(t *testing.T) {
	suite := setupServiceRegistryTest(t)
	defer suite.cleanup(t)

	// Create HTTP service registry
	serviceRegistry := model.ServiceRegistry{
		Name:       HTTPServiceName,
		Type_:      "HTTP",
		ServiceURI: "http://httpbin:8081/api-docs",
	}

	// Add service
	resp, err := suite.client.AddOrUpdateService(suite.ctx, serviceRegistry)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)

	// Discover service methods
	discoverOpts := &client.ServiceRegistryResourceApiDiscoverOpts{
		Create: optional.NewBool(true),
	}

	methods, resp, err := suite.client.Discover(suite.ctx, HTTPServiceName, discoverOpts)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.GreaterOrEqual(t, len(methods), 15)

	// Wait for discovery to complete
	time.Sleep(1 * time.Second)

	// Get registered services
	services, resp, err := suite.client.GetRegisteredServices(suite.ctx)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)

	// Find our service
	var actualService *model.ServiceRegistry
	for _, service := range services {
		if service.Name == HTTPServiceName {
			actualService = &service
			break
		}
	}
	require.NotNil(t, actualService, "HTTP service not found")

	// Verify service properties
	assert.Equal(t, HTTPServiceName, actualService.Name)
	assert.Equal(t, "HTTP", actualService.Type_) // Adjust based on your model
	assert.Equal(t, "http://httpbin:8081/api-docs", actualService.ServiceURI)
	assert.Greater(t, len(actualService.Methods), 0)

	originalMethodCount := len(actualService.Methods)

	// Add a custom method
	method := model.ServiceMethod{
		OperationName: "TestOperation",
		MethodName:    "addBySdkTest",
		MethodType:    "GET",
		InputType:     "newHttpInputType",
		OutputType:    "newHttpOutputType",
	}

	resp, err = suite.client.AddOrUpdateMethod(suite.ctx, method, HTTPServiceName)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify method was added
	updatedService, resp, err := suite.client.GetService(suite.ctx, HTTPServiceName)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, originalMethodCount+1, len(updatedService.Methods))

	// Verify circuit breaker config (adjust field names based on your model)
	if updatedService.Config != nil && updatedService.Config.CircuitBreakerConfig != nil {
		config := updatedService.Config.CircuitBreakerConfig
		assert.Equal(t, float32(50), config.FailureRateThreshold)
		assert.Equal(t, int32(100), config.MinimumNumberOfCalls)
		assert.Equal(t, int32(100), config.PermittedNumberOfCallsInHalfOpenState)
		assert.Equal(t, int64(1000), config.WaitDurationInOpenState)
		assert.Equal(t, int32(100), config.SlidingWindowSize)
		assert.Equal(t, float32(50), config.SlowCallRateThreshold)
		assert.Equal(t, int64(1), config.MaxWaitDurationInHalfOpenState)
	}

	// Cleanup
	resp, err = suite.client.RemoveService(suite.ctx, HTTPServiceName)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestServiceRegistryClient_GRPCService(t *testing.T) {
	suite := setupServiceRegistryTest(t)
	defer suite.cleanup(t)

	// Create gRPC service registry
	serviceRegistry := model.ServiceRegistry{
		Name:       GRPCServiceName,
		Type_:      "gRPC",
		ServiceURI: "grpcbin:50051",
	}

	// Add service
	resp, err := suite.client.AddOrUpdateService(suite.ctx, serviceRegistry)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)

	// Get registered services to verify
	services, resp, err := suite.client.GetRegisteredServices(suite.ctx)
	require.NoError(t, err)

	// Find our service
	var actualService *model.ServiceRegistry
	for _, service := range services {
		if service.Name == GRPCServiceName {
			actualService = &service
			break
		}
	}
	require.NotNil(t, actualService, "gRPC service not found")

	// Verify service properties
	assert.Equal(t, GRPCServiceName, actualService.Name)
	assert.Equal(t, "gRPC", actualService.Type_)
	assert.Equal(t, "grpcbin:50051", actualService.ServiceURI)
	assert.Equal(t, 0, len(actualService.Methods))

	originalMethodCount := len(actualService.Methods)

	// Add a custom method
	method := model.ServiceMethod{
		OperationName: "TestOperation",
		MethodName:    "addBySdkTest",
		MethodType:    "GET",
		InputType:     "newHttpInputType",
		OutputType:    "newHttpOutputType",
	}

	resp, err = suite.client.AddOrUpdateMethod(suite.ctx, method, GRPCServiceName)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify method was added
	updatedService, resp, err := suite.client.GetService(suite.ctx, GRPCServiceName)
	require.NoError(t, err)
	assert.Equal(t, originalMethodCount+1, len(updatedService.Methods))

	// Read binary proto data
	binaryData, err := testdata.ReadFile("compiled.bin")
	require.NoError(t, err)
	require.Greater(t, len(binaryData), 0)

	// Set proto data
	resp, err = suite.client.SetProtoData(suite.ctx, string(binaryData), GRPCServiceName, ProtoFilename)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify proto data was set and service has methods
	finalService, resp, err := suite.client.GetService(suite.ctx, GRPCServiceName)
	require.NoError(t, err)
	assert.Greater(t, len(finalService.Methods), 11)

	// Verify circuit breaker config
	if finalService.Config != nil && finalService.Config.CircuitBreakerConfig != nil {
		config := finalService.Config.CircuitBreakerConfig
		assert.Equal(t, float32(50), config.FailureRateThreshold)
		assert.Equal(t, int32(100), config.MinimumNumberOfCalls)
		assert.Equal(t, int32(100), config.PermittedNumberOfCallsInHalfOpenState)
		assert.Equal(t, int64(1000), config.WaitDurationInOpenState)
		assert.Equal(t, int32(100), config.SlidingWindowSize)
		assert.Equal(t, float32(50), config.SlowCallRateThreshold)
		assert.Equal(t, int64(1), config.MaxWaitDurationInHalfOpenState)
	}

	// Test getting all protos
	protos, resp, err := suite.client.GetAllProtos(suite.ctx, GRPCServiceName)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Greater(t, len(protos), 0)

	// Cleanup
	resp, err = suite.client.RemoveService(suite.ctx, GRPCServiceName)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestServiceRegistryClient_CircuitBreaker(t *testing.T) {
	suite := setupServiceRegistryTest(t)
	defer suite.cleanup(t)

	// First create a service
	serviceRegistry := model.ServiceRegistry{
		Name:       HTTPServiceName,
		Type_:      "HTTP",
		ServiceURI: "http://httpbin:8081/api-docs",
	}

	resp, err := suite.client.AddOrUpdateService(suite.ctx, serviceRegistry)
	require.NoError(t, err)

	// Test circuit breaker operations
	// Get initial status
	status, resp, err := suite.client.GetCircuitBreakerStatus(suite.ctx, HTTPServiceName)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, status)

	// Open circuit breaker
	_, resp, err = suite.client.OpenCircuitBreaker(suite.ctx, HTTPServiceName)
	require.Error(t, err)
	assert.Equal(t, 404, resp.StatusCode)
	assert.NotNil(t, err.Error())

	// Close circuit breaker
	closeResp, resp, err := suite.client.CloseCircuitBreaker(suite.ctx, HTTPServiceName)
	require.Error(t, err, "No active circuit breaker for service: http-service")
	assert.Equal(t, 404, resp.StatusCode)
	assert.NotNil(t, closeResp)

	// Cleanup
	resp, err = suite.client.RemoveService(suite.ctx, HTTPServiceName)
	require.NoError(t, err)
}

func TestServiceRegistryClient_MethodOperations(t *testing.T) {
	suite := setupServiceRegistryTest(t)
	defer suite.cleanup(t)

	// Create a service first
	serviceRegistry := model.ServiceRegistry{
		Name:       HTTPServiceName,
		Type_:      "HTTP",
		ServiceURI: "http://httpbin:8081/api-docs",
	}

	resp, err := suite.client.AddOrUpdateService(suite.ctx, serviceRegistry)
	require.NoError(t, err)

	// Add a method
	method := model.ServiceMethod{
		OperationName: "TestOperation",
		MethodName:    "testMethod",
		MethodType:    "POST",
		InputType:     "TestInput",
		OutputType:    "TestOutput",
	}

	resp, err = suite.client.AddOrUpdateMethod(suite.ctx, method, HTTPServiceName)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify method exists
	service, resp, err := suite.client.GetService(suite.ctx, HTTPServiceName)
	require.NoError(t, err)

	var foundMethod bool
	for _, m := range service.Methods {
		if m.MethodName == "testMethod" {
			foundMethod = true
			break
		}
	}
	assert.True(t, foundMethod, "Method should be found in service")

	// Remove the method
	resp, err = suite.client.RemoveMethod(suite.ctx, HTTPServiceName, "TestOperation", "testMethod", "POST")
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify method is removed
	service, resp, err = suite.client.GetService(suite.ctx, HTTPServiceName)
	require.NoError(t, err)

	foundMethod = false
	for _, m := range service.Methods {
		if m.MethodName == "testMethod" {
			foundMethod = true
			break
		}
	}
	assert.False(t, foundMethod, "Method should be removed from service")

	// Cleanup
	resp, err = suite.client.RemoveService(suite.ctx, HTTPServiceName)
	require.NoError(t, err)
}
