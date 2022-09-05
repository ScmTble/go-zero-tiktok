package logic

import (
	"context"
	"database/sql"
	"tiktok/video/rpc/internal/svc"
	"tiktok/video/rpc/model"
	"tiktok/video/rpc/video"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishVideoLogic {
	return &PublishVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishVideoLogic) PublishVideo(in *video.PublishVideoReq) (*video.PublishVideoResp, error) {
	v, err := l.svcCtx.VideoModel.Insert(l.ctx, &model.Videos{
		AuthorId:    in.UserId,
		PlayUrl:     in.PlayUrl,
		CoverUrl:    in.CoverUrl,
		PublishTime: time.Now(),
		Title: sql.NullString{
			String: in.Title,
			Valid:  true,
		},
	})
	if err != nil {
		return nil, err
	}
	id, _ := v.LastInsertId()
	return &video.PublishVideoResp{
		UserId:   id,
		PlayUrl:  in.PlayUrl,
		CoverUrl: in.CoverUrl,
		Title:    in.Title,
	}, nil
}
