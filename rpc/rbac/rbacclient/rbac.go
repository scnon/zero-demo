// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: rbac.proto

package rbacclient

import (
	"context"

	"zero-demo/rpc/rbac/pb/rbac"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AuthUserReq    = rbac.AuthUserReq
	AuthUserResp   = rbac.AuthUserResp
	DeleteUserReq  = rbac.DeleteUserReq
	DeleteUserResp = rbac.DeleteUserResp
	FindUserReq    = rbac.FindUserReq
	FindUserResp   = rbac.FindUserResp
	Request        = rbac.Request
	Response       = rbac.Response
	User           = rbac.User

	Rbac interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
		AddUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
		AuthUser(ctx context.Context, in *AuthUserReq, opts ...grpc.CallOption) (*AuthUserResp, error)
		DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*DeleteUserResp, error)
		FindUser(ctx context.Context, in *FindUserReq, opts ...grpc.CallOption) (*FindUserResp, error)
	}

	defaultRbac struct {
		cli zrpc.Client
	}
)

func NewRbac(cli zrpc.Client) Rbac {
	return &defaultRbac{
		cli: cli,
	}
}

func (m *defaultRbac) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := rbac.NewRbacClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}

func (m *defaultRbac) AddUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	client := rbac.NewRbacClient(m.cli.Conn())
	return client.AddUser(ctx, in, opts...)
}

func (m *defaultRbac) AuthUser(ctx context.Context, in *AuthUserReq, opts ...grpc.CallOption) (*AuthUserResp, error) {
	client := rbac.NewRbacClient(m.cli.Conn())
	return client.AuthUser(ctx, in, opts...)
}

func (m *defaultRbac) DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*DeleteUserResp, error) {
	client := rbac.NewRbacClient(m.cli.Conn())
	return client.DeleteUser(ctx, in, opts...)
}

func (m *defaultRbac) FindUser(ctx context.Context, in *FindUserReq, opts ...grpc.CallOption) (*FindUserResp, error) {
	client := rbac.NewRbacClient(m.cli.Conn())
	return client.FindUser(ctx, in, opts...)
}
