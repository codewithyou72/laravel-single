package menu

import (
	"context"

	"laravel-single/app/internal/svc"
	"laravel-single/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSysPermMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSysPermMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSysPermMenuLogic {
	return &AddSysPermMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSysPermMenuLogic) AddSysPermMenu(req *types.AddSysPermMenuReq) error {
	// todo: add your logic here and delete this line

	return nil
}
