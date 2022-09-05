package user

import (
	"context"
	"tiktok/pkg/jwtx"
	"tiktok/user/rpc/user"
	"time"

	"tiktok/app/api/internal/svc"
	"tiktok/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	loginResp, err := l.svcCtx.UserRpc.UserLogin(l.ctx, &user.LoginRequest{
		UserName: req.UserName,
		PassWord: req.PassWord,
	})

	if err != nil {
		return nil, err
	}
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessSecret := l.svcCtx.Config.JwtAuth.AccessSecret
	token, _ := jwtx.GetToken(accessSecret, now, accessExpire, loginResp.UserId)
	return &types.LoginResponse{
		UserId: loginResp.UserId,
		Token:  token,
	}, nil
}
