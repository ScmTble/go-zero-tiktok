package user

import (
	"net/http"
	"tiktok/pkg/result"

	"tiktok/app/api/internal/logic/user"
	"tiktok/app/api/internal/svc"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		result.HttpResult(r, w, resp, err)
	}
}
