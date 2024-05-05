package dept

import (
	"context"
	"fmt"
	errorx2 "zero-zone/app/pkg/errorx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"zero-zone/app/core/cmd/api/internal/svc"
	"zero-zone/app/core/cmd/api/internal/types"
)

type GetSysDeptListByWhereLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysDeptListByWhereLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysDeptListByWhereLogic {
	return &GetSysDeptListByWhereLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysDeptListByWhereLogic) GetSysDeptListByWhere(req *types.SearchDeptReq) (resp *types.SysDeptPageResp, err error) {
	fmt.Println("req", req)
	fmt.Printf("%+v\n", req)

	where := " 1 "
	if len(req.Name) > 0 {
		where = where + fmt.Sprintf(" AND name LIKE '%s'", "%"+req.Name+"%")
	}
	if len(req.FullName) > 0 {
		where = where + fmt.Sprintf(" AND full_name LIKE '%s'", "%"+req.FullName+"%")
	}
	if req.Status >= 0 {
		// todo 默认零时很有问题
	}
	sysDeptList, err := l.svcCtx.SysDeptModel.FindPagination(l.ctx, where, req.Page, req.Limit)
	if err != nil {
		return nil, errorx2.NewSystemError(errorx2.ServerErrorCode, err.Error())
	}

	var dept types.Dept
	deptList := make([]types.Dept, 0)
	for _, v := range sysDeptList {
		err := copier.Copy(&dept, &v)
		if err != nil {
			return nil, errorx2.NewSystemError(errorx2.ServerErrorCode, err.Error())
		}
		deptList = append(deptList, dept)
	}

	total, err := l.svcCtx.SysDeptModel.FindPaginationCount(l.ctx, where)
	if err != nil {
		return nil, errorx2.NewSystemError(errorx2.ServerErrorCode, err.Error())
	}

	pagination := types.Pagination{
		Page:  req.Page,
		Limit: req.Limit,
		Total: total,
	}

	return &types.SysDeptPageResp{
		List:       deptList,
		Pagination: pagination,
	}, nil
}
