package logic

import (
	"context"
	"tiktok/video/rpc/internal/svc"
	"tiktok/video/rpc/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishVideoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishVideoListLogic {
	return &PublishVideoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishVideoListLogic) PublishVideoList(in *video.PublishVideoListReq) (*video.PublishVideoListResp, error) {

	videos, err := l.svcCtx.VideoModel.FindListByUid(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	videoresp := make([]*video.PublishVideo, len(videos))

	for i, v := range videos {
		videoresp[i] = &video.PublishVideo{
			Id:          v.Id,
			UserId:      v.AuthorId,
			PlayUrl:     v.PlayUrl,
			CoverUrl:    v.CoverUrl,
			Title:       v.Title.String,
			PublishTime: v.PublishTime.Format("2006-01-02 15:04:05"),
		}
	}

	return &video.PublishVideoListResp{
		Videos: videoresp,
	}, nil
}
