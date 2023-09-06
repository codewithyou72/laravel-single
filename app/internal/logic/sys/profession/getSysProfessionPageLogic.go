package profession

import (
	"context"

	"laravel-single/app/internal/svc"
	"laravel-single/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysProfessionPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysProfessionPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysProfessionPageLogic {
	return &GetSysProfessionPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysProfessionPageLogic) GetSysProfessionPage(req *types.SysProfessionPageReq) (resp *types.SysProfessionPageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
