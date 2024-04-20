package client

import (
	"context"
	"github.com/conductor-sdk/conductor-go/sdk/model/rbac"
	"net/http"
)

type UserClient interface {
	CheckPermissions(ctx context.Context, userId string, type_ string, id string) (interface{}, *http.Response, error)
	DeleteUser(ctx context.Context, id string) (*http.Response, error)
	GetGrantedPermissions(ctx context.Context, userId string) (interface{}, *http.Response, error)
	GetUser(ctx context.Context, id string) (interface{}, *http.Response, error)
	ListUsers(ctx context.Context, optionals *UserResourceApiListUsersOpts) ([]rbac.ConductorUser, *http.Response, error)
	UpsertUser(ctx context.Context, body rbac.UpsertUserRequest, id string) (interface{}, *http.Response, error)
}
