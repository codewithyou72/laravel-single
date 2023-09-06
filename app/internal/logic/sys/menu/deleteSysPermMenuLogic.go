package menu

import (
	"context"

	"laravel-single/app/internal/svc"
	"laravel-single/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSysPermMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSysPermMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysPermMenuLogic {
	return &DeleteSysPermMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSysPermMenuLogic) DeleteSysPermMenu(req *types.DeleteSysPermMenuReq) error {
	// todo: add your logic here and delete this line

	return nil
}
