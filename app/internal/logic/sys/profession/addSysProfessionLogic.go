package profession

import (
	"context"

	"laravel-single/app/internal/svc"
	"laravel-single/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSysProfessionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSysProfessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSysProfessionLogic {
	return &AddSysProfessionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSysProfessionLogic) AddSysProfession(req *types.AddSysProfessionReq) error {
	// todo: add your logic here and delete this line

	return nil
}
