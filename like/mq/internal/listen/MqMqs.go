package listen

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"tiktok/like/mq/internal/svc"
	"tiktok/like/rpc/like"
)

func AddMqService(group *service.ServiceGroup, ctx context.Context, svcCtx *svc.ServiceContext) {
	listener := MustNewListener(ctx, svcCtx, handler)
	group.Add(listener)
}

func handler(ctx context.Context, svcCtx *svc.ServiceContext, data []byte) error {
	logx.Info("receive message success")

	msg := make(map[string]int64, 3)

	if err := json.Unmarshal(data, &msg); err != nil {
		return err
	}
	_, err := svcCtx.LikeRpc.LikeVideo(ctx, &like.LikeVideoReq{
		UserId:     msg["UserId"],
		VideoId:    msg["VideoId"],
		StatusCode: uint32(msg["StatusCode"]),
	})
	if err != nil {
		return err
	}
	return nil
}
