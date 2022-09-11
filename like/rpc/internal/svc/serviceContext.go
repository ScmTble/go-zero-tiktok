package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tiktok/like/rpc/internal/config"
	"tiktok/like/rpc/model"
)

type ServiceContext struct {
	Config    config.Config
	LikeModel model.LikesModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		LikeModel: model.NewLikesModel(sqlx.NewMysql(c.Mysql.DataSource), c.CacheRedis),
	}
}
