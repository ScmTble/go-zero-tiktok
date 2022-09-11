package like

import (
	"net/http"
	"tiktok/pkg/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"tiktok/app/api/internal/logic/like"
	"tiktok/app/api/internal/svc"
	"tiktok/app/api/internal/types"
)

func LikeVideoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LikeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := like.NewLikeVideoLogic(r.Context(), svcCtx)
		resp, err := l.LikeVideo(&req)
		result.HttpResult(r, w, resp, err)
	}
}
