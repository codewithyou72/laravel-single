package dept

import (
	"context"

	"laravel-single/app/internal/svc"
	"laravel-single/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysDeptListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysDeptListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysDeptListLogic {
	return &GetSysDeptListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysDeptListLogic) GetSysDeptList() (resp *types.SysDeptListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
