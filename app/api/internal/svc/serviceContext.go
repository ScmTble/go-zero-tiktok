package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"tiktok/app/api/internal/config"
	"tiktok/app/api/internal/middleware"
	"tiktok/user/rpc/user"
	"tiktok/user/rpc/userclient"
	"tiktok/video/rpc/video"
	"tiktok/video/rpc/videoclient"
)

type ServiceContext struct {
	Config     config.Config
	UserRpc    user.UserClient
	VideoRpc   video.VideoClient
	UploadFile rest.Middleware
	Redis      *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		UserRpc:    userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		VideoRpc:   videoclient.NewVideo(zrpc.MustNewClient(c.VideoRpc)),
		UploadFile: middleware.NewUploadFileMiddleware().Handle,
		Redis:      redis.New(c.Redis.Host),
	}
}
