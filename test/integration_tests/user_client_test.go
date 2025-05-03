package integration_tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model/rbac"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"github.com/stretchr/testify/assert"
)

// TestCheckPermissions checks if permissions for a user can be retrieved correctly.
func TestCheckPermissions(t *testing.T) {
	TestUpsertUser(t)
	client := NewUserClient()
	ctx := context.Background()
	userId := "testuser"
	type_ := "WORKFLOW_DEF"
	id := "kitchen_sink"

	permissions, resp, err := client.CheckPermissions(ctx, userId, type_, id)
	if err != nil {
		t.Fatalf("CheckPermissions failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.Status)
	}
	if _, ok := permissions["CREATE"]; !ok {
		t.Errorf("Expected 'allowed' field in the response, but found %s", permissions)
	}
}

// TestDeleteUser verifies that a user can be successfully deleted.
func TestDeleteUser(t *testing.T) {
	TestUpsertUser(t)
	client := NewUserClient()
	ctx := context.Background()
	id := "testuser"

	resp, err := client.DeleteUser(ctx, id)
	if err != nil {
		t.Fatalf("DeleteUser failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.Status)
	}
}

// TestGetGrantedPermissions checks if granted permissions can be fetched for a user.
func TestGetGrantedPermissions(t *testing.T) {
	TestUpsertUser(t)
	client := NewUserClient()
	ctx := context.Background()
	userId := "testuser"

	permissions, resp, err := client.GetGrantedPermissions(ctx, userId)
	if err != nil {
		t.Fatalf("GetGrantedPermissions failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.Status)
	}
	if len(permissions.GrantedAccess) != 0 {
		t.Errorf("Expected non-empty permissions %d", len(permissions.GrantedAccess))
	}
}

// TestGetUser checks fetching a specific user's details.
func TestGetUser(t *testing.T) {
	TestUpsertUser(t)
	client := NewUserClient()
	ctx := context.Background()
	id := "testuser"

	user, resp, err := client.GetUser(ctx, id)
	if err != nil {
		t.Fatalf("GetUser failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.Status)
	}
	if user.Id != id {
		t.Errorf("Expected user ID %v, got %v", id, user.Id)
	}
}

func TestGetUserNotFound(t *testing.T) {
	TestUpsertUser(t)
	client := NewUserClient()
	ctx := context.Background()
	id := "testuserxxx_doesnot_exist"

	user, resp, _ := client.GetUser(ctx, id)

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status 404, got %v", resp.Status)
	}
	assert.Nil(t, user)

}

// TestListUsers checks listing users with optional parameters.
func TestListUsers(t *testing.T) {
	user_client := NewUserClient()
	ctx := context.Background()
	options := client.UserResourceApiListUsersOpts{Apps: optional.NewBool(true)}

	users, resp, err := user_client.ListUsers(ctx, &options)
	if err != nil {
		t.Fatalf("ListUsers failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.Status)
	}
	if len(users) == 0 {
		t.Errorf("Expected non-empty user list")
	}
}

// TestUpsertUser verifies that a user can be updated or inserted.
func TestUpsertUser(t *testing.T) {
	client := NewUserClient()
	ctx := context.Background()
	body := rbac.UpsertUserRequest{
		Name:  "testuser",
		Roles: []string{"ADMIN", "USER"},
	}
	id := "testUser"

	user, resp, err := client.UpsertUser(ctx, body, id)
	if err != nil {
		t.Fatalf("UpsertUser failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.Status)
	}
	if user.Name != body.Name {
		t.Errorf("Expected username %v, got %v", body.Name, user.Name)
	}
}

func NewUserClient() client.UserClient {
	return testdata.UserClient
}
