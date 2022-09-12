package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"tiktok/like/mq/internal/config"
	"tiktok/like/rpc/like"
	"tiktok/like/rpc/likeclient"
)

type ServiceContext struct {
	Config  config.Config
	LikeRpc like.LikeClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		LikeRpc: likeclient.NewLike(zrpc.MustNewClient(c.LikeRpc)),
	}
}
