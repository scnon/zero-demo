// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: rbac.proto

package server

import (
	"context"

	"zero-demo/rpc/rbac/internal/logic"
	"zero-demo/rpc/rbac/internal/svc"
	"zero-demo/rpc/rbac/pb/rbac"
)

type RbacServer struct {
	svcCtx *svc.ServiceContext
	rbac.UnimplementedRbacServer
}

func NewRbacServer(svcCtx *svc.ServiceContext) *RbacServer {
	return &RbacServer{
		svcCtx: svcCtx,
	}
}

func (s *RbacServer) Ping(ctx context.Context, in *rbac.Request) (*rbac.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}

func (s *RbacServer) AddUser(ctx context.Context, in *rbac.User) (*rbac.User, error) {
	l := logic.NewAddUserLogic(ctx, s.svcCtx)
	return l.AddUser(in)
}

func (s *RbacServer) AuthUser(ctx context.Context, in *rbac.AuthUserReq) (*rbac.AuthUserResp, error) {
	l := logic.NewAuthUserLogic(ctx, s.svcCtx)
	return l.AuthUser(in)
}

func (s *RbacServer) DeleteUser(ctx context.Context, in *rbac.DeleteUserReq) (*rbac.DeleteUserResp, error) {
	l := logic.NewDeleteUserLogic(ctx, s.svcCtx)
	return l.DeleteUser(in)
}

func (s *RbacServer) FindUser(ctx context.Context, in *rbac.FindUserReq) (*rbac.FindUserResp, error) {
	l := logic.NewFindUserLogic(ctx, s.svcCtx)
	return l.FindUser(in)
}
