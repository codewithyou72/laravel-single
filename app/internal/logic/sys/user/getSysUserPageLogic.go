package user

import (
	"context"

	"laravel-single/app/internal/svc"
	"laravel-single/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysUserPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysUserPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysUserPageLogic {
	return &GetSysUserPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysUserPageLogic) GetSysUserPage(req *types.SysUserPageReq) (resp *types.SysUserPageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
