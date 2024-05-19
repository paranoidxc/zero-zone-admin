package cronjob

import (
	"net/http"
	"zero-zone/app/pkg/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-zone/app/core/cmd/api/internal/logic/sys/cronjob"
	"zero-zone/app/core/cmd/api/internal/svc"
	"zero-zone/app/core/cmd/api/internal/types"
)

// 更新
func UpdateCronjobHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateCronjobReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := cronjob.NewUpdateCronjobLogic(r.Context(), svcCtx)
		err := l.UpdateCronjob(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}

		response.Response(w, nil, err)
	}
}
