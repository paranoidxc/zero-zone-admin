package cronjob

import (
	"context"

	"zero-zone/app/core/cmd/api/internal/svc"
	"zero-zone/app/core/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCronjobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 新增
func NewAddCronjobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCronjobLogic {
	return &AddCronjobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCronjobLogic) AddCronjob(req *types.AddCronjobReq) error {
	// todo: add your logic here and delete this line

	return nil
}
