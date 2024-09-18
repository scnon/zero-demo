package svc

import (
	"zero-demo/api/admin/internal/config"
	"zero-demo/api/admin/internal/middleware"
	"zero-demo/api/admin/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config config.Config

	UserModel model.SysUserModel
	CheckUrl  rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config: c,

		UserModel: model.NewSysUserModel(sqlConn),
		CheckUrl:  middleware.NewCheckUrlMiddleware().Handle,
	}
}
