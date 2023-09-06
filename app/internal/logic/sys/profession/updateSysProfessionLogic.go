package profession

import (
	"context"

	"laravel-single/app/internal/svc"
	"laravel-single/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysProfessionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSysProfessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysProfessionLogic {
	return &UpdateSysProfessionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSysProfessionLogic) UpdateSysProfession(req *types.UpdateSysProfessionReq) error {
	// todo: add your logic here and delete this line

	return nil
}
