package svc

import (
	"zero-demo/rpc/rbac/internal/config"
	"zero-demo/rpc/rbac/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	UserModel model.RbacUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,

		UserModel: model.NewRbacUserModel(sqlConn),
	}
}
