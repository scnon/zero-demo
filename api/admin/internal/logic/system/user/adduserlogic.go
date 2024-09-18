package user

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"zero-demo/api/admin/internal/model"
	"zero-demo/api/admin/internal/svc"
	"zero-demo/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req *types.AddUserReq) (err error) {
	account, _ := l.svcCtx.UserModel.FindOneByAccount(l.ctx, req.Account)

	if account != nil {
		return errors.New("用户已存在")
	}

	user := model.SysUser{
		Account:    req.Account,
		Password:   req.Password,
		UserName:   req.UserName,
		StatusId:   int64(req.Status),
		Sort:       int64(req.Sort),
		CreateTime: time.Time{},
		UpdateTime: sql.NullTime{},
	}
	_, err = l.svcCtx.UserModel.Insert(l.ctx, &user)

	if err != nil {
		return err
	}

	return
}
