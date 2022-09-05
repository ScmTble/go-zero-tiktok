package logic

import (
	"context"
	"tiktok/pkg/tool"
	"tiktok/user/rpc/model"

	"tiktok/user/rpc/internal/svc"
	"tiktok/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserRegisterLogic) UserRegister(in *user.RegisterRequest) (*user.RegisterResponse, error) {

	result, err := l.svcCtx.UserModel.Insert(l.ctx, &model.Users{
		Name:     in.UserName,
		Password: tool.EnCoder(in.PassWord),
	})
	if err != nil {
		return nil, err
	}
	userId, _ := result.LastInsertId()
	return &user.RegisterResponse{
		UserId:   userId,
		UserName: in.UserName,
	}, nil
}
