package dept

import (
	"net/http"
	"zero-zone/app/core/cmd/api/internal/types"
	errorx2 "zero-zone/app/pkg/errorx"
	"zero-zone/app/pkg/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-zone/app/core/cmd/api/internal/logic/sys/dept"
	"zero-zone/app/core/cmd/api/internal/svc"
)

func GetSysDeptListByNameHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchDeptReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorx2.NewHandlerError(errorx2.ParamErrorCode, err.Error()))
			return
		}

		l := dept.NewGetSysDeptListByWhereLogic(r.Context(), svcCtx)
		resp, err := l.GetSysDeptListByWhere(&req)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		response.Response(w, resp, err)
	}
}
