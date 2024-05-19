package cronjob

import (
	"context"

	"zero-zone/app/core/cmd/api/internal/svc"
	"zero-zone/app/core/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ViewCronjobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查看
func NewViewCronjobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewCronjobLogic {
	return &ViewCronjobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ViewCronjobLogic) ViewCronjob(req *types.ViewCronjobReq) (resp *types.ViewCronjobResp, err error) {
	// todo: add your logic here and delete this line

	return
}
