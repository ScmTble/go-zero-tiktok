package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tiktok/video/rpc/internal/config"
	"tiktok/video/rpc/model"
)

type ServiceContext struct {
	Config     config.Config
	VideoModel model.VideosModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		VideoModel: model.NewVideosModel(sqlConn, c.CacheRedis),
	}
}
