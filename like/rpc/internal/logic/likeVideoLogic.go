package logic

import (
	"context"

	"tiktok/like/rpc/internal/svc"
	"tiktok/like/rpc/like"

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
	// todo: add your logic here and delete this line

	return &like.LikeVideoResp{}, nil
}