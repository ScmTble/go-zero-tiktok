package logic

import (
	"context"
	"github.com/pkg/errors"
	"tiktok/pkg/tool"

	"tiktok/user/rpc/internal/svc"
	"tiktok/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserLoginLogic) UserLogin(in *user.LoginRequest) (*user.LoginResponse, error) {
	u, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.UserName)
	if err != nil {
		return nil, err
	}
	password := tool.EnCoder(in.PassWord)
	if u.Password != password {
		return nil, errors.New("用户名或密码错误")
	}
	return &user.LoginResponse{
		UserId:   u.Id,
		UserName: u.Name,
	}, nil
}
