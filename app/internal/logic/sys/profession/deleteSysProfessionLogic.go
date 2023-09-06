package profession

import (
	"context"

	"laravel-single/app/internal/svc"
	"laravel-single/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSysProfessionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSysProfessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysProfessionLogic {
	return &DeleteSysProfessionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSysProfessionLogic) DeleteSysProfession(req *types.DeleteSysProfessionReq) error {
	// todo: add your logic here and delete this line

	return nil
}
