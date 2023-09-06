package user

import (
	"context"

	"laravel-single/app/internal/svc"
	"laravel-single/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
	// todo: add your logic here and delete this line

	return
}
