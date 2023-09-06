package login

import (
	"context"

	"laravel-single/app/internal/svc"
	"laravel-single/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogLoginPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLogLoginPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogLoginPageLogic {
	return &GetLogLoginPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogLoginPageLogic) GetLogLoginPage(req *types.LogLoginPageReq) (resp *types.LogLoginPageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
