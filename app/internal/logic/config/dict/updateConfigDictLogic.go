package dict

import (
	"context"

	"laravel-single/app/internal/svc"
	"laravel-single/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateConfigDictLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateConfigDictLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateConfigDictLogic {
	return &UpdateConfigDictLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateConfigDictLogic) UpdateConfigDict(req *types.UpdateConfigDictReq) error {
	// todo: add your logic here and delete this line

	return nil
}
