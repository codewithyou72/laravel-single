package user

import (
	"context"

	"laravel-single/app/internal/svc"
	"laravel-single/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPermMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserPermMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPermMenuLogic {
	return &GetUserPermMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserPermMenuLogic) GetUserPermMenu() (resp *types.UserPermMenuResp, err error) {
	// todo: add your logic here and delete this line

	return
}
