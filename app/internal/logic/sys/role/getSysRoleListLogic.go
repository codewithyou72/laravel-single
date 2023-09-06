package role

import (
	"context"

	"laravel-single/app/internal/svc"
	"laravel-single/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysRoleListLogic {
	return &GetSysRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysRoleListLogic) GetSysRoleList() (resp *types.SysRoleListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
