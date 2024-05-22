package td_firm

import (
	"context"

	"zero-zone/app/core/cmd/api/internal/svc"
	"zero-zone/app/core/cmd/api/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	
	errorx2 "zero-zone/app/pkg/errorx"
)

type TdFirmListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTdFirmListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TdFirmListLogic {
	return &TdFirmListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TdFirmListLogic) TdFirmList(req *types.TdFirmListReq) (resp *types.TdFirmListResp, err error) {
    where := " 1 "
    /*
    if len(req.FirmName) > 0 {
        where = where + fmt.Sprintf(" AND firm_name LIKE '%s'", "%"+req.FirmName+"%")
    }
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
    featTdFirmList, err := l.svcCtx.FeatTdFirmModel.FindAllByWhere(l.ctx, where)
	if err != nil {
		return nil, errorx2.NewSystemError(errorx2.ServerErrorCode, err.Error())
	}

	var item types.TdFirm
	TdFirmList := make([]types.TdFirm, 0)
	for _, v := range featTdFirmList {
		err := copier.Copy(&item, &v)
		if err != nil {
			return nil, errorx2.NewSystemError(errorx2.ServerErrorCode, err.Error())
		}
		TdFirmList = append(TdFirmList, item)
	}

	return &types.TdFirmListResp{
		List: TdFirmList,
	}, nil
}
