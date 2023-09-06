package dict

import (
	"context"

	"laravel-single/app/internal/svc"
	"laravel-single/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddConfigDictLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddConfigDictLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddConfigDictLogic {
	return &AddConfigDictLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddConfigDictLogic) AddConfigDict(req *types.AddConfigDictReq) error {
	// todo: add your logic here and delete this line

	return nil
}
