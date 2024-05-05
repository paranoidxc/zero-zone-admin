package user

import (
	"context"
	errorx2 "zero-zone/app/pkg/errorx"
	"zero-zone/app/pkg/globalkey"

	"github.com/zeromicro/go-zero/core/logx"
	"zero-zone/app/core/cmd/api/internal/svc"
	"zero-zone/app/core/cmd/api/internal/types"
)

type DeleteSysUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSysUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysUserLogic {
	return &DeleteSysUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSysUserLogic) DeleteSysUser(req *types.DeleteSysUserReq) error {
	if req.Id == globalkey.SysSuperUserId {
		return errorx2.NewDefaultError(errorx2.ForbiddenErrorCode)
	}

	err := l.svcCtx.SysUserModel.Delete(l.ctx, req.Id)
	if err != nil {
		return errorx2.NewSystemError(errorx2.ServerErrorCode, err.Error())
	}

	return nil
}
