package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

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
