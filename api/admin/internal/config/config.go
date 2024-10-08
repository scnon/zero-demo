package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	RbacClientConf zrpc.RpcClientConf

	Mysql struct {
		DataSource string
	}

	Auth struct {
		AccessSecret  string
		AccessExpire  int64
		RefreshSecret string
	}

	Redis struct {
		Address string
		Pass    string
	}
}
