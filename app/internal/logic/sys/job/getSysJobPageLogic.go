package job

import (
	"context"

	"laravel-single/app/internal/svc"
	"laravel-single/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysJobPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysJobPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysJobPageLogic {
	return &GetSysJobPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysJobPageLogic) GetSysJobPage(req *types.SysJobPageReq) (resp *types.SysJobPageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
