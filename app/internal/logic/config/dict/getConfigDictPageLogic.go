package dict

import (
	"context"

	"laravel-single/app/internal/svc"
	"laravel-single/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigDictPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetConfigDictPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigDictPageLogic {
	return &GetConfigDictPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConfigDictPageLogic) GetConfigDictPage(req *types.ConfigDictPageReq) (resp *types.ConfigDictPageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
