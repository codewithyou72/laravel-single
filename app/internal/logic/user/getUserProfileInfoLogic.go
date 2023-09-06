package user

import (
	"context"

	"laravel-single/app/internal/svc"
	"laravel-single/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserProfileInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserProfileInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserProfileInfoLogic {
	return &GetUserProfileInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserProfileInfoLogic) GetUserProfileInfo() (resp *types.UserProfileInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
