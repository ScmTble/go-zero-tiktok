package logic

import (
	"context"
	"github.com/pkg/errors"
	"tiktok/config"
	"tiktok/like/rpc/internal/svc"
	"tiktok/like/rpc/like"
	"tiktok/like/rpc/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLikeVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeVideoLogic {
	return &LikeVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LikeVideoLogic) LikeVideo(in *like.LikeVideoReq) (*like.LikeVideoResp, error) {
	likeData, err := l.svcCtx.LikeModel.FindOneByUserIdVideoId(l.ctx, in.UserId, in.VideoId)
	if errors.Is(err, model.ErrNotFound) {
		if in.StatusCode == config.UserLikeVideo {
			// 用户从未对此VideoId点赞过且点赞
			_, err := l.svcCtx.LikeModel.Insert(l.ctx, &model.Likes{
				UserId:  in.UserId,
				VideoId: in.VideoId,
				Cancel:  config.UserLikeVideo,
			})
			if err != nil {
				return nil, err
			}
		}
		// 不存在且取消点赞，状态异常
		return &like.LikeVideoResp{
			StatusCode: false,
		}, nil
	}
	// 用户已经操作过，需要更新状态并且数据库中对状态和发送来对不同
	if likeData.Cancel != int64(in.StatusCode) {
		likeData.Cancel = int64(in.StatusCode)
		if err := l.svcCtx.LikeModel.Update(l.ctx, likeData); err != nil {
			return nil, err
		}
	}
	return &like.LikeVideoResp{
		StatusCode: true,
	}, nil
}
