package profession

import (
	"context"
	errorx2 "zero-zone/app/pkg/errorx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"zero-zone/app/core/cmd/api/internal/svc"
	"zero-zone/app/core/cmd/api/internal/types"
)

type GetSysProfessionPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysProfessionPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysProfessionPageLogic {
	return &GetSysProfessionPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysProfessionPageLogic) GetSysProfessionPage(req *types.SysProfessionPageReq) (resp *types.SysProfessionPageResp, err error) {
	sysProfessionList, err := l.svcCtx.SysProfessionModel.FindPage(l.ctx, req.Page, req.Limit)
	if err != nil {
		return nil, errorx2.NewSystemError(errorx2.ServerErrorCode, err.Error())
	}

	var profession types.Profession
	professionList := make([]types.Profession, 0)
	for _, v := range sysProfessionList {
		err := copier.Copy(&profession, &v)
		if err != nil {
			return nil, errorx2.NewSystemError(errorx2.ServerErrorCode, err.Error())
		}
		professionList = append(professionList, profession)
	}

	total, err := l.svcCtx.SysProfessionModel.FindCount(l.ctx)
	if err != nil {
		return nil, errorx2.NewSystemError(errorx2.ServerErrorCode, err.Error())
	}

	pagination := types.Pagination{
		Page:  req.Page,
		Limit: req.Limit,
		Total: total,
	}

	return &types.SysProfessionPageResp{
		List:       professionList,
		Pagination: pagination,
	}, nil
}
