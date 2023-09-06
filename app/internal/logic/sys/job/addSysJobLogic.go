package job

import (
	"context"

	"laravel-single/app/internal/svc"
	"laravel-single/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSysJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSysJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSysJobLogic {
	return &AddSysJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSysJobLogic) AddSysJob(req *types.AddSysJobReq) error {
	// todo: add your logic here and delete this line

	return nil
}
