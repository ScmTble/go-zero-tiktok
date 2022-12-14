package svc

import (
	"github.com/streadway/amqp"
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

type RabbitMQContext struct {
	Conn        *amqp.Connection
	Channel     *amqp.Channel
	ContentType string
}

type ServiceContext struct {
	Config     config.Config
	UserRpc    user.UserClient
	VideoRpc   video.VideoClient
	UploadFile rest.Middleware
	Redis      *redis.Redis
	RabbitMQ   *RabbitMQContext
}

func NewServiceContext(c config.Config) *ServiceContext {

	coon, err := amqp.Dial(c.RabbitMQ.GetUrl())
	if err != nil {
		panic(err)
	}
	channel, err := coon.Channel()
	if err != nil {
		panic(err)
	}

	rabbitMQ := &RabbitMQContext{
		Conn:        coon,
		Channel:     channel,
		ContentType: "text/plain",
	}

	return &ServiceContext{
		Config:     c,
		UserRpc:    userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		VideoRpc:   videoclient.NewVideo(zrpc.MustNewClient(c.VideoRpc)),
		UploadFile: middleware.NewUploadFileMiddleware().Handle,
		Redis:      redis.New(c.Redis.Host),
		RabbitMQ:   rabbitMQ,
	}
}
