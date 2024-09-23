package logic

import (
	"context"

	"zero-demo/rpc/rbac/internal/svc"
	"zero-demo/rpc/rbac/pb/rbac"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLogic {
	return &FindUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindUserLogic) FindUser(in *rbac.FindUserReq) (*rbac.FindUserResp, error) {
	// todo: add your logic here and delete this line

	return &rbac.FindUserResp{}, nil
}
