package cronjob

import (
	"context"

	"zero-zone/app/core/cmd/api/internal/svc"
	"zero-zone/app/core/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCronjobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除
func NewDeleteCronjobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCronjobLogic {
	return &DeleteCronjobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCronjobLogic) DeleteCronjob(req *types.DeleteCronjobReq) error {
	// todo: add your logic here and delete this line

	return nil
}
