package user

import (
	"context"

	"zero-demo/api/admin/internal/svc"
	"zero-demo/api/admin/internal/types"
	"zero-demo/rpc/rbac/pb/rbac"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.DeleteUserReq) error {
	_, err := l.svcCtx.RbacRpc.DeleteUser(l.ctx, &rbac.DeleteUserReq{
		Ids: req.Ids,
	})
	return err
}
