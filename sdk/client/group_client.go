package client

import (
	"context"
	"github.com/conductor-sdk/conductor-go/sdk/model/rbac"
	"net/http"
)

type GroupClient interface {
	AddUserToGroup(ctx context.Context, groupId string, userId string) (interface{}, *http.Response, error)
	AddUsersToGroup(ctx context.Context, body []string, groupId string) (*http.Response, error)
	DeleteGroup(ctx context.Context, id string) (*http.Response, error)
	GetGrantedPermissions1(ctx context.Context, groupId string) (rbac.GrantedAccessResponse, *http.Response, error)
	GetGroup(ctx context.Context, id string) (interface{}, *http.Response, error)
	GetUsersInGroup(ctx context.Context, id string) (interface{}, *http.Response, error)
	ListGroups(ctx context.Context) ([]rbac.Group, *http.Response, error)
	RemoveUserFromGroup(ctx context.Context, groupId string, userId string) (interface{}, *http.Response, error)
	RemoveUsersFromGroup(ctx context.Context, body []string, groupId string) (*http.Response, error)
	UpsertGroup(ctx context.Context, body rbac.UpsertGroupRequest, id string) (interface{}, *http.Response, error)
}

func NewGroupClient(client *APIClient) GroupClient {
	return &GroupResourceApiService{client}
}
