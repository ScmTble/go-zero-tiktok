package video

import (
	"net/http"
	"tiktok/pkg/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"tiktok/app/api/internal/logic/video"
	"tiktok/app/api/internal/svc"
	"tiktok/app/api/internal/types"
)

func PublishVideoListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublishVideoListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := video.NewPublishVideoListLogic(r.Context(), svcCtx)
		resp, err := l.PublishVideoList(&req)
		result.HttpResult(r, w, resp, err)
	}
}
