package video

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"tiktok/app/api/internal/svc"
	"tiktok/app/api/internal/types"
	"tiktok/video/rpc/video"
)

type PublishVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishVideoLogic {
	return &PublishVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishVideoLogic) PublishVideo(req *types.PublishVideoReq) (resp *types.PublishVideoResp, err error) {

	uid, _ := l.ctx.Value("uid").(json.Number).Int64()

	res, err := l.svcCtx.VideoRpc.PublishVideo(l.ctx, &video.PublishVideoReq{
		UserId:   uid,
		PlayUrl:  req.PlayUrl,
		CoverUrl: req.CoverUrl,
		Title:    req.Title,
	})
	if err != nil {
		return nil, err
	}
	return &types.PublishVideoResp{
		PlayUrl:  res.PlayUrl,
		CoverUrl: res.CoverUrl,
		Title:    res.Title,
	}, nil
}
