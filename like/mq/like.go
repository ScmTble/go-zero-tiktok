package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"tiktok/like/mq/internal/config"
	"tiktok/like/mq/internal/listen"
	"tiktok/like/mq/internal/svc"
)

var configFile = flag.String("f", "etc/like.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	if err := c.SetUp(); err != nil {
		fmt.Println(err)
		panic(err)
	}

	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	svcCtx := svc.NewServiceContext(c)

	// Mq ：消息队列.
	listen.AddMqService(serviceGroup, context.Background(), svcCtx)
	// asynq：延迟队列、定时任务
	//services = append(services, AsynqMqs(c, ctx, svcContext)...)
	// other mq ....

	serviceGroup.Start()
}
