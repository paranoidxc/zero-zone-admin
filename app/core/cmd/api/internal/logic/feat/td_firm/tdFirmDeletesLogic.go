package td_firm

import (
	"context"

	"zero-zone/app/core/cmd/api/internal/svc"
	"zero-zone/app/core/cmd/api/internal/types"

	
	"github.com/zeromicro/go-zero/core/logx"
	
	errorx2 "zero-zone/app/pkg/errorx"
)

type TdFirmDeletesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTdFirmDeletesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TdFirmDeletesLogic {
	return &TdFirmDeletesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TdFirmDeletesLogic) TdFirmDeletes(req *types.TdFirmDeletesReq) (err error) {
	if len(req.FirmId) > 0  {
		err = l.svcCtx.FeatTdFirmModel.Deletes(l.ctx, req.FirmId)
		if err != nil {
			return  errorx2.NewSystemError(errorx2.ServerErrorCode, err.Error())
		}
	} else {
		return errorx2.NewSystemError(errorx2.ParamErrorCode, err.Error())
	}

	return
}

