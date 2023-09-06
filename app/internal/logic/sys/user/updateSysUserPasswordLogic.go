package user

import (
	"context"

	"laravel-single/app/internal/svc"
	"laravel-single/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysUserPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSysUserPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysUserPasswordLogic {
	return &UpdateSysUserPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSysUserPasswordLogic) UpdateSysUserPassword(req *types.UpdateSysUserPasswordReq) error {
	// todo: add your logic here and delete this line

	return nil
}
