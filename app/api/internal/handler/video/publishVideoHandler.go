package video

import (
	"net/http"
	"tiktok/pkg/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"tiktok/app/api/internal/logic/video"
	"tiktok/app/api/internal/svc"
	"tiktok/app/api/internal/types"
)

func PublishVideoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublishVideoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := video.NewPublishVideoLogic(r.Context(), svcCtx)
		resp, err := l.PublishVideo(&req)
		result.HttpResult(r, w, resp, err)
	}
}
