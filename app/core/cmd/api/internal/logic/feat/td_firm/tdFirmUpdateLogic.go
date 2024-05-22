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

type TdFirmUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTdFirmUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TdFirmUpdateLogic {
	return &TdFirmUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TdFirmUpdateLogic) TdFirmUpdate(req *types.TdFirmUpdateReq) (err error) {
	modelParams := &model.TdFirm{}
	modelParams, err = l.svcCtx.FeatTdFirmModel.FindOne(l.ctx, req.FirmId)
	if err != nil {
		return errorx2.NewDefaultError(errorx2.UserIdErrorCode)
	}

	err = copier.Copy(modelParams, req)
	if err != nil {
		logx.Error("复制参数失败", err)
		return errorx2.NewSystemError(errorx2.ServerErrorCode, err.Error())
	}

	err = l.svcCtx.FeatTdFirmModel.Update(l.ctx, modelParams)
	if err != nil {
		return errorx2.NewSystemError(errorx2.ServerErrorCode, err.Error())
	}

	return
}

