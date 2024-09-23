package logic

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"zero-demo/rpc/rbac/internal/model"
	"zero-demo/rpc/rbac/internal/svc"
	"zero-demo/rpc/rbac/pb/rbac"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserLogic) AddUser(in *rbac.User) (*rbac.User, error) {
	if strings.TrimSpace(in.Account) == "" || strings.TrimSpace(in.Password) == "" {
		return nil, errors.New("账号或密码不能为空")
	}

	account, err := l.svcCtx.UserModel.FindOneByAccount(l.ctx, in.Account)
	if err != nil {
		return nil, err
	}

	if account != nil {
		return nil, errors.New("用户已存在")
	}

	if len(in.Password) < 8 || len(in.Password) > 18 {
		return nil, errors.New("密码长度必须在8-18位之间")
	}

	user := &model.RbacUser{
		Account:  in.Account,
		Password: in.Password,
		UserName: in.UserName,
		Roles:    sql.NullString{String: in.Roles},
	}

	_, err = l.svcCtx.UserModel.Insert(l.ctx, user)
	if err != nil {
		return nil, err
	}
	return &rbac.User{
		Account:  user.Account,
		Password: user.Password,
		UserName: user.UserName,
		Roles:    user.Roles.String,
	}, nil
}
