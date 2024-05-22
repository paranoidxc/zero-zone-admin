package td_firm

import (
	"context"

	"zero-zone/app/core/cmd/api/internal/svc"
	"zero-zone/app/core/cmd/api/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"zero-zone/app/core/model"
	errorx2 "zero-zone/app/pkg/errorx"
)

type TdFirmCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTdFirmCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TdFirmCreateLogic {
	return &TdFirmCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TdFirmCreateLogic) TdFirmCreate(req *types.TdFirmCreateReq) (err error) {
	var modelParams = new(model.TdFirm)
	err = copier.Copy(modelParams, req)
	if err != nil {
		return errorx2.NewSystemError(errorx2.ServerErrorCode, err.Error())
	}
	_, err = l.svcCtx.FeatTdFirmModel.Insert(l.ctx, modelParams)
	if err != nil {
		return errorx2.NewSystemError(errorx2.ServerErrorCode, err.Error())
	}

	return
}

