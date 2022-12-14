// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	like "tiktok/app/api/internal/handler/like"
	user "tiktok/app/api/internal/handler/user"
	video "tiktok/app/api/internal/handler/video"
	"tiktok/app/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: user.UserRegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.UserLoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/douyin/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/",
				Handler: user.UserInfoHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/douyin/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.UploadFile},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/action",
					Handler: video.PublishVideoHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/list/:userId",
					Handler: video.PublishVideoListHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/douyin/publish"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/action",
				Handler: like.LikeVideoHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/douyin/favourite"),
	)
}
