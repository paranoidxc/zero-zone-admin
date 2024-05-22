package td_firm

import (
	"context"

	"zero-zone/app/core/cmd/api/internal/svc"
	"zero-zone/app/core/cmd/api/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"zero-zone/app/core/model"
	
)

type TdFirmDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTdFirmDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TdFirmDetailLogic {
	return &TdFirmDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TdFirmDetailLogic) TdFirmDetail(req *types.TdFirmDetailReq) (resp *types.TdFirmDetailResp, err error) {
	resp = &types.TdFirmDetailResp{}
	item := &model.TdFirm{}
	item, err = l.svcCtx.FeatTdFirmModel.FindOne(l.ctx, req.FirmId)
	err = copier.Copy(resp, item)
	if err != nil {
		logx.Error("复制结果失败", err)
		return nil, err
	}
	return
}

