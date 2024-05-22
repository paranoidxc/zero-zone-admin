package td_firm

import (
	"context"

	"zero-zone/app/core/cmd/api/internal/svc"
	"zero-zone/app/core/cmd/api/internal/types"

	
	"github.com/zeromicro/go-zero/core/logx"
	
	errorx2 "zero-zone/app/pkg/errorx"
)

type TdFirmDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTdFirmDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TdFirmDeleteLogic {
	return &TdFirmDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TdFirmDeleteLogic) TdFirmDelete(req *types.TdFirmDeleteReq) (err error) {
	err = l.svcCtx.FeatTdFirmModel.Delete(l.ctx, req.FirmId)
	if err != nil {
		return errorx2.NewSystemError(errorx2.ServerErrorCode, err.Error())
	}

	return
}

