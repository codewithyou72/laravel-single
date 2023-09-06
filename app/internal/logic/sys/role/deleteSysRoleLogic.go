package role

import (
	"context"

	"laravel-single/app/internal/svc"
	"laravel-single/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSysRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSysRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysRoleLogic {
	return &DeleteSysRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSysRoleLogic) DeleteSysRole(req *types.DeleteSysRoleReq) error {
	// todo: add your logic here and delete this line

	return nil
}
