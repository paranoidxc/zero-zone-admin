package menu

import (
	"net/http"
	"zero-zone/app/pkg/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-zone/app/core/cmd/api/internal/logic/sys/menu"
	"zero-zone/app/core/cmd/api/internal/svc"
)

func GetSysPermMenuListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := menu.NewGetSysPermMenuListLogic(r.Context(), svcCtx)
		resp, err := l.GetSysPermMenuList()
		if err != nil {
			httpx.Error(w, err)
			return
		}

		response.Response(w, resp, err)
	}
}
