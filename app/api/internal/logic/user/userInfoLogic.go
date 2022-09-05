package user

import (
	"context"
	"encoding/json"
	"tiktok/user/rpc/user"

	"tiktok/app/api/internal/svc"
	"tiktok/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line
	value := l.ctx.Value("uid")
	uid, _ := value.(json.Number).Int64()
	info, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserIdRequest{UserId: uid})
	if err != nil {
		return nil, err
	}
	return &types.UserInfoResponse{
		UserId:         info.UserId,
		UserName:       info.UserName,
		FollowCount:    0,
		FollowerCount:  0,
		IsFollow:       false,
		TotalFavorited: 0,
		FavoriteCount:  0,
	}, nil
}
