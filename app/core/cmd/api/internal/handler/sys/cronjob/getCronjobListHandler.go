package cronjob

import (
	"net/http"
	"zero-zone/app/pkg/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-zone/app/core/cmd/api/internal/logic/sys/cronjob"
	"zero-zone/app/core/cmd/api/internal/svc"
	"zero-zone/app/core/cmd/api/internal/types"
)

// 全部不分页
func GetCronjobListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CronjobListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := cronjob.NewGetCronjobListLogic(r.Context(), svcCtx)
		resp, err := l.GetCronjobList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}

		response.Response(w, resp, err)
	}
}
