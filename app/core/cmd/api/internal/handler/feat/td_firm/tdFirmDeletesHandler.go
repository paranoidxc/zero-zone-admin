package td_firm

import (
	"net/http"
	"zero-zone/app/pkg/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-zone/app/core/cmd/api/internal/logic/feat/td_firm"
	"zero-zone/app/core/cmd/api/internal/svc"
	"zero-zone/app/core/cmd/api/internal/types"
)

func TdFirmDeletesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TdFirmDeletesReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, err)
			return
		}

		l := td_firm.NewTdFirmDeletesLogic(r.Context(), svcCtx)
		err := l.TdFirmDeletes(&req)
		if err != nil {
			response.Response(w, nil, err)
			return
		}

		response.Response(w, nil, nil)
	}
}
