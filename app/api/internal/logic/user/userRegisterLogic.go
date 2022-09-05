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

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	// todo: add your logic here and delete this line
	response, err := l.svcCtx.UserRpc.UserRegister(l.ctx, &user.RegisterRequest{
		UserName: req.UserName,
		PassWord: req.PassWord,
	})
	if err != nil {
		return
	}
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessSecret := l.svcCtx.Config.JwtAuth.AccessSecret
	token, _ := jwtx.GetToken(accessSecret, now, accessExpire, response.UserId)
	return &types.RegisterResponse{
		UserId: response.UserId,
		Token:  token,
	}, nil
}
