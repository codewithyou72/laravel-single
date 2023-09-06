package user

import (
	"context"

	"laravel-single/app/internal/svc"
	"laravel-single/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysUserRdpjInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysUserRdpjInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysUserRdpjInfoLogic {
	return &GetSysUserRdpjInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysUserRdpjInfoLogic) GetSysUserRdpjInfo(req *types.GetSysUserRdpjInfoReq) (resp *types.GetSysUserRdpjInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
