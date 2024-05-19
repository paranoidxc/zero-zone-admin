package cronjob

import (
	"context"

	"zero-zone/app/core/cmd/api/internal/svc"
	"zero-zone/app/core/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCronjobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新
func NewUpdateCronjobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCronjobLogic {
	return &UpdateCronjobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCronjobLogic) UpdateCronjob(req *types.UpdateCronjobReq) error {
	// todo: add your logic here and delete this line

	return nil
}
