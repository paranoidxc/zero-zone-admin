package user

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"zero-zone/app/pkg/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-zone/app/core/cmd/api/internal/logic/user"
	"zero-zone/app/core/cmd/api/internal/svc"
)

func GetLoginCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		logx.Info("GetLoginCaptchaHandler")

		l := user.NewGetLoginCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.GetLoginCaptcha()
		if err != nil {
			httpx.Error(w, err)
			return
		}

		response.Response(w, resp, err)
	}
}
