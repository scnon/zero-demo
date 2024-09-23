package user

import (
	"context"

	"zero-demo/api/admin/internal/svc"
	"zero-demo/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserListLogic {
	return &QueryUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryUserListLogic) QueryUserList(req *types.QueryUserListReq) (resp *types.QueryUserListResp, err error) {
	resp = &types.QueryUserListResp{}

	userList, total, err := l.svcCtx.UserModel.FindAll(l.ctx, req.Page, req.PageSize, req.UserName, req.Status)
	if err != nil {
		return
	}

	list := []types.UserData{}
	for _, user := range *userList {
		list = append(list, types.UserData{
			Id:         user.Id,
			UserName:   user.Username,
			NickName:   user.Nickname,
			Status:     int(user.Status),
			Sort:       int(user.Sort),
			Remark:     user.Remark.String,
			CreateTime: user.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime: user.UpdateTime.Time.Format("2006-01-02 15:04:05"),
		})
	}

	resp.List = list
	resp.Total = int(total)
	return
}
