package video

import (
	"context"
	"sync"
	"tiktok/user/rpc/user"
	"tiktok/video/rpc/video"

	"tiktok/app/api/internal/svc"
	"tiktok/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishVideoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishVideoListLogic {
	return &PublishVideoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishVideoListLogic) PublishVideoList(req *types.PublishVideoListReq) (resp *types.PublishVideoListResp, err error) {

	var group sync.WaitGroup
	group.Add(2)

	var info *user.UserInfoResponse

	var res *video.PublishVideoListResp

	go func() {
		info, err = l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserIdRequest{
			UserId: req.UserId,
		})
		defer group.Done()
		if err != nil {
			logx.Error(err)
			return
		}
	}()

	go func() {
		res, err = l.svcCtx.VideoRpc.PublishVideoList(l.ctx, &video.PublishVideoListReq{
			UserId: req.UserId,
		})
		defer group.Done()
		if err != nil {
			return
		}
	}()

	group.Wait()

	videos := make([]types.Video, len(res.Videos))

	for i, v := range res.Videos {
		videos[i].Id = v.Id
		videos[i].Author = types.UserInfo{
			UserId:        info.UserId,
			UserName:      info.UserName,
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
		}
		videos[i].PlayUrl = v.PlayUrl
		videos[i].CoverUrl = v.CoverUrl
		videos[i].FavoriteCount = 0
		videos[i].CommentCount = 0
		videos[i].IsFavorite = false
		videos[i].Title = v.Title

	}

	return &types.PublishVideoListResp{
		Videos: videos,
	}, nil
}
