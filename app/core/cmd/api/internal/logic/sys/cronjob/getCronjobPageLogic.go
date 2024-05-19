package cronjob

import (
	"context"

	"zero-zone/app/core/cmd/api/internal/svc"
	"zero-zone/app/core/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCronjobPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页
func NewGetCronjobPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCronjobPageLogic {
	return &GetCronjobPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCronjobPageLogic) GetCronjobPage(req *types.CronjobPageReq) (resp *types.CronjobPageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
