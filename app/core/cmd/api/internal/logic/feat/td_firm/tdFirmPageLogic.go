package td_firm

import (
	"context"
	"fmt"
	"strings"

	"zero-zone/app/core/cmd/api/internal/svc"
	"zero-zone/app/core/cmd/api/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"

	errorx2 "zero-zone/app/pkg/errorx"
)

type TdFirmPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTdFirmPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TdFirmPageLogic {
	return &TdFirmPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TdFirmPageLogic) TdFirmPage(req *types.TdFirmPageReq) (resp *types.TdFirmPageResp, err error) {
	where := " 1 "
	if len(strings.TrimSpace(req.FirmName)) > 0 {
		where = where + fmt.Sprintf(" AND firm_name LIKE '%s'", "%"+strings.TrimSpace(req.FirmName)+"%")
	}
	/*
	   if len(req.FirmAlias) > 0 {
	       where = where + fmt.Sprintf(" AND firm_alias LIKE '%s'", "%"+req.FirmAlias+"%")
	   }
	   if len(req.FirmCode) > 0 {
	       where = where + fmt.Sprintf(" AND firm_code LIKE '%s'", "%"+req.FirmCode+"%")
	   }
	   if len(req.FirmDesc) > 0 {
	       where = where + fmt.Sprintf(" AND firm_desc LIKE '%s'", "%"+req.FirmDesc+"%")
	   }
	*/

	featTdFirmPage, err := l.svcCtx.FeatTdFirmModel.FindPageByWhere(l.ctx, where, req.Page, req.Limit)
	if err != nil {
		return nil, errorx2.NewSystemError(errorx2.ServerErrorCode, err.Error())
	}

	var item types.TdFirm
	TdFirmPage := make([]types.TdFirm, 0)
	for _, v := range featTdFirmPage {
		err := copier.Copy(&item, &v)
		if err != nil {
			return nil, errorx2.NewSystemError(errorx2.ServerErrorCode, err.Error())
		}
		TdFirmPage = append(TdFirmPage, item)
	}

	total, err := l.svcCtx.FeatTdFirmModel.FindPageByWhereCount(l.ctx, where)
	if err != nil {
		return nil, errorx2.NewSystemError(errorx2.ServerErrorCode, err.Error())
	}

	pagination := types.Pagination{
		Page:  req.Page,
		Limit: req.Limit,
		Total: total,
	}

	return &types.TdFirmPageResp{
		List:       TdFirmPage,
		Pagination: pagination,
	}, nil
}
