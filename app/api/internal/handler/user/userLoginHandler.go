package user

import (
	"net/http"
	"tiktok/pkg/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"tiktok/app/api/internal/logic/user"
	"tiktok/app/api/internal/svc"
	"tiktok/app/api/internal/types"
)

func UserLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUserLoginLogic(r.Context(), svcCtx)
		resp, err := l.UserLogin(&req)
		result.HttpResult(r, w, resp, err)
	}
}
