package integration_tests

import (
	"context"
	"fmt"
	"github.com/conductor-sdk/conductor-go/sdk/model/rbac"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGroupResourceApiService(t *testing.T) {
	// Setup
	groupClient := testdata.GroupClient // Assuming this exists in your testdata package

	// Test case 1: Create a new group
	ctx := context.Background()
	groupId := fmt.Sprintf("test-group-%d", time.Now().Unix())

	upsertRequest := rbac.UpsertGroupRequest{
		DefaultAccess: map[string][]string{
			"WORKFLOW_DEF": {"READ", "CREATE"},
			"TASK_DEF":     {"READ"},
		},
		Description: "A test group for integration testing",
		Roles:       []string{"ADMIN", "USER"},
	}

	group, resp, err := groupClient.UpsertGroup(ctx, upsertRequest, groupId)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, group)

	// Test case 2: Get the created group
	retrievedGroup, resp, err := groupClient.GetGroup(ctx, groupId)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify group properties (exact assertion depends on your response structure)
	// This is just a placeholder - adjust according to your actual response format
	groupMap, ok := retrievedGroup.(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, groupId, groupMap["id"])

	// Test case 3: List all groups and verify our group exists
	groups, resp, err := groupClient.ListGroups(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var foundGroup bool
	for _, g := range groups {
		if g.Id == groupId {
			foundGroup = true
			assert.Equal(t, groupId, g.Id)
			break
		}
	}
	assert.True(t, foundGroup, "Created group not found in list")

	user, err := testdata.CreateNewUser(ctx)
	assert.Nil(t, err)

	testUserId := user.Id
	_, resp, err = groupClient.AddUserToGroup(ctx, groupId, testUserId)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test case 5: Get users in the group and verify our user exists
	usersInGroup, resp, err := groupClient.GetUsersInGroup(ctx, groupId)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Again, exact assertion depends on your response structure
	// This is a placeholder - adjust according to your actual response format
	usersArray, ok := usersInGroup.([]interface{})
	assert.True(t, ok)

	var foundUser bool
	for _, u := range usersArray {
		user, ok := u.(map[string]interface{})
		if ok && user["id"] == testUserId {
			foundUser = true
			break
		}
	}
	assert.True(t, foundUser, "Added user not found in group")

	// Test case 6: Add multiple users to group
	user1, err := testdata.CreateNewUser(ctx)
	assert.Nil(t, err)
	user2, err := testdata.CreateNewUser(ctx)
	assert.Nil(t, err)

	additionalUsers := []string{user1.Id, user2.Id}
	resp, err = groupClient.AddUsersToGroup(ctx, additionalUsers, groupId)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test case 7: Get users again to verify multiple users were added
	usersInGroup, resp, err = groupClient.GetUsersInGroup(ctx, groupId)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test case 8: Check permissions for the group
	permissions, resp, err := groupClient.GetGrantedPermissions1(ctx, groupId)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, permissions)
	// You can add assertions here based on your expected permissions structure

	// Test case 9: Remove a single user from the group
	_, resp, err = groupClient.RemoveUserFromGroup(ctx, groupId, testUserId)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test case 10: Remove multiple users from the group
	resp, err = groupClient.RemoveUsersFromGroup(ctx, additionalUsers, groupId)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test case 11: Verify all users were removed
	usersInGroup, resp, err = groupClient.GetUsersInGroup(ctx, groupId)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// This assertion depends on how your API represents an empty group
	emptyUsersArray, ok := usersInGroup.([]interface{})
	assert.True(t, ok)
	assert.Empty(t, emptyUsersArray, "Group should have no users")

	// Test case 12: Update the group
	updatedRequest := rbac.UpsertGroupRequest{
		DefaultAccess: map[string][]string{
			"WORKFLOW_DEF": {"READ", "CREATE", "DELETE"},
			"TASK_DEF":     {"READ", "DELETE"},
		},
		Description: "A test group for integration testing-Updated",
		Roles:       []string{"ADMIN", "USER"},
	}

	updatedGroup, resp, err := groupClient.UpsertGroup(ctx, updatedRequest, groupId)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify update
	updatedGroupMap, ok := updatedGroup.(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, updatedRequest.Description, updatedGroupMap["description"])

	// Test case 13: Delete the group
	resp, err = groupClient.DeleteGroup(ctx, groupId)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test case 14: Verify the group was deleted
	groups, resp, err = groupClient.ListGroups(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	for _, g := range groups {
		assert.NotEqual(t, groupId, g.Id, "Group was not deleted successfully")
	}
}
