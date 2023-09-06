package dict

import (
	"context"

	"laravel-single/app/internal/svc"
	"laravel-single/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigDictListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetConfigDictListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigDictListLogic {
	return &GetConfigDictListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConfigDictListLogic) GetConfigDictList() (resp *types.ConfigDictListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
