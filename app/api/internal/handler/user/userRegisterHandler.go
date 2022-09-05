package user

import (
	"net/http"
	"tiktok/pkg/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"tiktok/app/api/internal/logic/user"
	"tiktok/app/api/internal/svc"
	"tiktok/app/api/internal/types"
)

func UserRegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUserRegisterLogic(r.Context(), svcCtx)
		resp, err := l.UserRegister(&req)
		result.HttpResult(r, w, resp, err)
	}
}
