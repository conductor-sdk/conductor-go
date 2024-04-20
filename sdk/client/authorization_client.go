package client

import (
	"context"
	"github.com/conductor-sdk/conductor-go/sdk/model/rbac"
	"net/http"
)

type AuthorizationClient interface {
	GetPermissions(ctx context.Context, type_ string, id string) (interface{}, *http.Response, error)
	GrantPermissions(ctx context.Context, body rbac.AuthorizationRequest) (*http.Response, error)
	RemovePermissions(ctx context.Context, body rbac.AuthorizationRequest) (*http.Response, error)
}
