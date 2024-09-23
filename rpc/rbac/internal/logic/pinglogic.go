package logic

import (
	"context"

	"zero-demo/rpc/rbac/internal/svc"
	"zero-demo/rpc/rbac/pb/rbac"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *rbac.Request) (*rbac.Response, error) {
	// todo: add your logic here and delete this line

	return &rbac.Response{}, nil
}
