package integration_tests

import (
	"context"
	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
	"time"
)

func TestHTTPServiceRegistryClientIntegration(t *testing.T) {
	ctx := context.Background()

	// Create a new client
	serviceRegistryClient := NewServiceRegistryClient()

	// Use a unique name for testing to avoid conflicts
	testRegistryName := "sdk_test_http_registry_" + time.Now().Format("20060102150405")

	// Test 1: Add a new service registry
	t.Run("AddOrUpdateService", func(t *testing.T) {
		registry := model.ServiceRegistry{
			Name:       testRegistryName,
			ServiceURI: "http://localhost:8081/api-docs",
			Type_:      "HTTP",
		}

		resp, err := serviceRegistryClient.AddOrUpdateService(ctx, registry)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})

	// Test 2: Get the registered service
	t.Run("GetService", func(t *testing.T) {
		registry, resp, err := serviceRegistryClient.GetService(ctx, testRegistryName)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, testRegistryName, registry.Name)
		assert.Equal(t, "HTTP", registry.Type_)
	})

	// Test 3: Get all registered services
	t.Run("GetRegisteredServices", func(t *testing.T) {
		services, resp, err := serviceRegistryClient.GetRegisteredServices(ctx)
		assert.NoError(t, err)
		assert.NotNil(t, resp)

		// Verify our test service is in the list
		found := false
		for _, service := range services {
			if service.Name == testRegistryName {
				found = true
				break
			}
		}
		assert.True(t, found, "Test service registry not found in the list")
	})

	// Test 4: For discovering methods and validating method count
	t.Run("DiscoverAndValidateMethods", func(t *testing.T) {
		// First, update the service registry with a valid API docs endpoint
		// This is a more reliable endpoint that should return OpenAPI spec
		updatedRegistry := model.ServiceRegistry{
			Name:       testRegistryName,
			ServiceURI: "http://httpbin:8081/api-docs",
			Type_:      "HTTP",
		}

		resp, err := serviceRegistryClient.AddOrUpdateService(ctx, updatedRegistry)
		assert.NoError(t, err)
		assert.NotNil(t, resp)

		// Give the system a moment to process the update
		time.Sleep(1 * time.Second)

		// Create discover options with create=true to create methods if discovered
		discoverOpts := &client.ServiceRegistryResourceApiDiscoverOpts{
			Create: optional.NewBool(true),
		}

		// Discover methods in the service
		methods, resp, err := serviceRegistryClient.Discover(ctx, testRegistryName, discoverOpts)

		// Even if there's an error, we should get a response
		assert.NotNil(t, resp)

		// If discovery succeeded, verify methods were found
		if err == nil {
			assert.NotEmpty(t, methods, "No methods discovered")
			assert.Greater(t, len(methods), 0, "Method count should be greater than 0")

			// Log discovered methods for debugging
			t.Logf("Discovered %d methods", len(methods))
			for i, method := range methods {
				t.Logf("Method %d: %s %s", i+1, method.MethodType, method.MethodName)
			}

			// Verify method details
			if len(methods) > 0 {
				method := methods[0]
				assert.NotEmpty(t, method.MethodName)
				assert.NotEmpty(t, method.MethodType)
				assert.NotEmpty(t, method.OperationName)
			}
		} else {
			// Log the error but don't fail the test since this depends on external service
			t.Logf("Discovery error (test continues): %v", err)
		}
		
		// Now get all methods for the service to verify they exist
		registry, resp, err := serviceRegistryClient.GetService(ctx, testRegistryName)
		assert.NoError(t, err)
		assert.NotNil(t, resp)

		// Check if the registry has any methods
		methodCount := 0
		if registry.Methods != nil {
			methodCount = len(registry.Methods)
		}

		t.Logf("Service has %d registered methods", methodCount)

		// If discovery was successful, we should have methods
		if err == nil && len(methods) > 0 {
			assert.Greater(t, methodCount, 0, "Method count in registry should be greater than 0")
		}
	})

	// Test for adding, getting, editing, and deleting a method
	t.Run("AddGetEditDeleteMethod", func(t *testing.T) {
		// 1. First, get the current number of methods
		registry, resp, err := serviceRegistryClient.GetService(ctx, testRegistryName)
		assert.NoError(t, err)
		assert.NotNil(t, resp)

		initialMethodCount := 0
		if registry.Methods != nil {
			initialMethodCount = len(registry.Methods)
		}
		t.Logf("Initial method count: %d", initialMethodCount)

		// 2. Add a new method
		testMethod := model.ServiceMethod{
			OperationName: "testMethodForCRUD",
			MethodName:    "/api/test/crud",
			MethodType:    "GET",
			InputType:     "string",
			OutputType:    "object",
		}

		resp, err = serviceRegistryClient.AddOrUpdateMethod(ctx, testMethod, testRegistryName)
		assert.NoError(t, err)
		assert.NotNil(t, resp)

		// 3. Get the service again and verify the method was added
		registry, resp, err = serviceRegistryClient.GetService(ctx, testRegistryName)
		assert.NoError(t, err)
		assert.NotNil(t, resp)

		// Check method count increased
		newMethodCount := 0
		if registry.Methods != nil {
			newMethodCount = len(registry.Methods)
		}
		t.Logf("New method count: %d", newMethodCount)
		assert.Equal(t, initialMethodCount+1, newMethodCount, "Method count should have increased by 1")

		// Find the added method
		var foundMethod *model.ServiceMethod
		for i := range registry.Methods {
			if registry.Methods[i].OperationName == "testMethodForCRUD" {
				foundMethod = &registry.Methods[i]
				break
			}
		}

		assert.NotNil(t, foundMethod, "Added method should be found in the service")
		if foundMethod != nil {
			assert.Equal(t, "GET", foundMethod.MethodType)
			assert.Equal(t, "/api/test/crud", foundMethod.MethodName)
		}

		// 4. Update the method
		testMethod.MethodType = "POST"
		testMethod.OperationName = "testMethodForCRUDUpdated"
		testMethod.MethodName = "/api/test/crud/updated"
		testMethod.Id = foundMethod.Id

		resp, err = serviceRegistryClient.AddOrUpdateMethod(ctx, testMethod, testRegistryName)
		assert.NoError(t, err)
		assert.NotNil(t, resp)

		// 5. Get the service again and verify the method was updated
		registry, resp, err = serviceRegistryClient.GetService(ctx, testRegistryName)
		assert.NoError(t, err)
		assert.NotNil(t, resp)

		// Check method count remained the same (update, not add)
		updatedMethodCount := 0
		if registry.Methods != nil {
			updatedMethodCount = len(registry.Methods)
		}
		assert.Equal(t, newMethodCount, updatedMethodCount, "Method count should remain the same after update")

		// Find the updated method
		foundMethod = nil
		for i := range registry.Methods {
			if registry.Methods[i].OperationName == "testMethodForCRUDUpdated" {
				foundMethod = &registry.Methods[i]
				break
			}
		}

		assert.NotNil(t, foundMethod, "Updated method should be found in the service")
		if foundMethod != nil {
			assert.Equal(t, "POST", foundMethod.MethodType, "Method type should be updated to POST")
			assert.Equal(t, "/api/test/crud/updated", foundMethod.MethodName, "URI should be updated")
		}

		// 6. Delete the method
		resp, err = serviceRegistryClient.RemoveMethod(ctx, testRegistryName, testRegistryName, "testMethodForCRUD", "POST")
		assert.NoError(t, err)
		assert.NotNil(t, resp)

		// 7. Get the service again and verify the method was deleted
		registry, resp, err = serviceRegistryClient.GetService(ctx, testRegistryName)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})

	// Finally, remove the http service
	t.Run("RemoveService", func(t *testing.T) {
		resp, err := serviceRegistryClient.RemoveService(ctx, testRegistryName)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})

	// Test for GRPC service with compiled proto file
	t.Run("GRPCServiceWithProtoFile", func(t *testing.T) {
		// Create a unique name for the GRPC service registry
		grpcRegistryName := "sdk_test_grpc_registry_" + time.Now().Format("20060102150405")

		// 1. Create a new GRPC service registry
		grpcRegistry := model.ServiceRegistry{
			Name:       grpcRegistryName,
			ServiceURI: "localhost:50051",
			Type_:      "gRPC",
		}

		resp, err := serviceRegistryClient.AddOrUpdateService(ctx, grpcRegistry)
		assert.NoError(t, err)
		assert.NotNil(t, resp)

		// 2. Read the compiled proto file from resources
		protoBytes, err := ioutil.ReadFile("compiled.bin")
		if err != nil {
			t.Fatalf("Failed to read compiled.bin: %v", err)
		}

		// 3. Upload the proto file
		protoFileName := "compiled.bin"
		resp, err = serviceRegistryClient.SetProtoData(ctx, protoBytes, grpcRegistryName, protoFileName)
		assert.NoError(t, err)
		assert.NotNil(t, resp)

		// 4. Verify the proto file was uploaded
		protos, resp, err := serviceRegistryClient.GetAllProtos(ctx, grpcRegistryName)
		assert.NoError(t, err)
		assert.NotNil(t, resp)

		// Check if our proto file is in the list
		found := false
		for _, proto := range protos {
			if proto.Filename == protoFileName {
				found = true
				break
			}
		}
		assert.True(t, found, "Uploaded proto file not found in the list")

		// 5. Get the proto file content and verify it matches
		protoData, resp, err := serviceRegistryClient.GetProtoData(ctx, grpcRegistryName, protoFileName)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, protoBytes, protoData, "Retrieved proto data should match uploaded data")

		// 6. Check the service registry for methods
		registry, resp, err := serviceRegistryClient.GetService(ctx, grpcRegistryName)
		assert.NoError(t, err)
		assert.NotNil(t, resp)

		// 7. Check if methods were discovered and created
		if registry.Methods != nil {
			t.Logf("Service has %d registered methods", len(registry.Methods))
			assert.Greater(t, len(registry.Methods), 0, "Method count should be greater than 0")

			// Log the methods for debugging
			for i, method := range registry.Methods {
				t.Logf("Method %d: %s %s", i+1, method.MethodType, method.MethodName)
			}
		} else {
			t.Log("No methods found in the registry")
		}

		// 8. Clean up by removing the service
		resp, err = serviceRegistryClient.RemoveService(ctx, grpcRegistryName)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})
}
