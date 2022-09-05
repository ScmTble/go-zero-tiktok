package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tiktok/user/rpc/internal/config"
	"tiktok/user/rpc/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUsersModel(sqlConn, c.CacheRedis),
	}
}
