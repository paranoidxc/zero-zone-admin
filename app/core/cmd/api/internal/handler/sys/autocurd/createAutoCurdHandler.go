package autocurd

import (
	"net/http"
	"zero-zone/app/pkg/response"

	"zero-zone/app/core/cmd/api/internal/logic/sys/autocurd"
	"zero-zone/app/core/cmd/api/internal/svc"
)

// 新增
func CreateAutoCurdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := autocurd.NewCreateAutoCurdLogic(r.Context(), svcCtx)
		err := l.CreateAutoCurd()
		if err != nil {
			response.Response(w, nil, err)
			return
		}

		response.Response(w, nil, nil)
	}
}
