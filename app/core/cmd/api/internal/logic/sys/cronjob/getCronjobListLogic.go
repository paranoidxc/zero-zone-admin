package cronjob

import (
	"context"

	"zero-zone/app/core/cmd/api/internal/svc"
	"zero-zone/app/core/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCronjobListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 全部不分页
func NewGetCronjobListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCronjobListLogic {
	return &GetCronjobListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCronjobListLogic) GetCronjobList(req *types.CronjobListReq) (resp *types.CronjobListResp, err error) {
	// todo: add your logic here and delete this line

	resp = &types.CronjobListResp{}

	return
}
