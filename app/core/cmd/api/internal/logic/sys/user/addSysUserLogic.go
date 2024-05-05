package user

import (
	"context"
	"encoding/json"
	errorx2 "zero-zone/app/pkg/errorx"
	"zero-zone/app/pkg/globalkey"
	utils2 "zero-zone/app/pkg/utils"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"zero-zone/app/core/cmd/api/internal/svc"
	"zero-zone/app/core/cmd/api/internal/types"
	"zero-zone/app/core/model"
)

type AddSysUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSysUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSysUserLogic {
	return &AddSysUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSysUserLogic) AddSysUser(req *types.AddSysUserReq) error {
	_, err := l.svcCtx.SysUserModel.FindOneByAccount(l.ctx, req.Account)
	if err == model.ErrNotFound {

		_, err := l.svcCtx.SysDeptModel.FindOne(l.ctx, req.DeptId)
		if err != nil {
			return errorx2.NewDefaultError(errorx2.DeptIdErrorCode)
		}

		_, err = l.svcCtx.SysProfessionModel.FindOne(l.ctx, req.ProfessionId)
		if err != nil {
			return errorx2.NewDefaultError(errorx2.ProfessionIdErrorCode)
		}

		_, err = l.svcCtx.SysJobModel.FindOne(l.ctx, req.JobId)
		if err != nil {
			return errorx2.NewDefaultError(errorx2.JobIdErrorCode)
		}

		var sysUser = new(model.SysUser)
		err = copier.Copy(sysUser, req)
		if err != nil {
			return errorx2.NewSystemError(errorx2.ServerErrorCode, err.Error())
		}

		bytes, err := json.Marshal(req.RoleIds)
		sysUser.RoleIds = string(bytes)
		dictionary, err := l.svcCtx.SysDictionaryModel.FindOneByUniqueKey(l.ctx, "sys_pwd")
		var password string
		if dictionary.Status == globalkey.SysEnable {
			password = dictionary.Value
		} else {
			password = globalkey.SysNewUserDefaultPassword
		}

		sysUser.Password = utils2.MD5(password + l.svcCtx.Config.Salt)
		sysUser.Avatar = utils2.AvatarUrl()
		_, err = l.svcCtx.SysUserModel.Insert(l.ctx, sysUser)
		if err != nil {
			return errorx2.NewSystemError(errorx2.ServerErrorCode, err.Error())
		}

		return nil
	} else {

		return errorx2.NewDefaultError(errorx2.AddUserErrorCode)
	}
}
