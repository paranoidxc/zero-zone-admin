package user

import (
	"context"
	"zero-zone/app/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
	"zero-zone/app/core/cmd/api/internal/svc"
	"zero-zone/app/core/cmd/api/internal/types"
)

type GetGenerateAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGenerateAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGenerateAvatarLogic {
	return &GetGenerateAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGenerateAvatarLogic) GetGenerateAvatar() (resp *types.GenerateAvatarResp, err error) {
	return &types.GenerateAvatarResp{
		AvatarUrl: utils.AvatarUrl(),
	}, nil
}
