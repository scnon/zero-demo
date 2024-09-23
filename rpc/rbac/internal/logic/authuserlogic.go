package logic

import (
	"context"

	"zero-demo/rpc/rbac/internal/svc"
	"zero-demo/rpc/rbac/pb/rbac"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuthUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthUserLogic {
	return &AuthUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AuthUserLogic) AuthUser(in *rbac.AuthUserReq) (*rbac.AuthUserResp, error) {
	// todo: add your logic here and delete this line

	return &rbac.AuthUserResp{}, nil
}
