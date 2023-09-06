package user

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"laravel-single/app/internal/svc"
	"laravel-single/app/internal/types"
)

type GetLoginCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLoginCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLoginCaptchaLogic {
	return &GetLoginCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLoginCaptchaLogic) GetLoginCaptcha() (resp *types.LoginCaptchaResp, err error) {

	one, err := l.svcCtx.DemoModel.FindOne(l.ctx, 1)
	if err != nil {
		return nil, err
	}
	fmt.Println(one)
	return
}

//原生方式使用gorm，对接之前gin代码适用这种方式
/*func (l *GetLoginCaptchaLogic) GetLoginCaptcha2() (resp *types.LoginCaptchaResp, err error) {

	//var demo2 []mysql.Demo

	sql2 := l.svcCtx.Gorm.ToSQL(func(tx *gorm.DB) *gorm.DB { //这种方法有效
		return tx.Model(&mysql.Demo{}).Where("id = ?", 100).Limit(10).Order("age desc").Find(&[]mysql.Demo{})
	})

	fmt.Println("SQL2语句=", sql2)

	var demo []mysql.Demo

	l.svcCtx.Gorm.Statement.Limit(10).Find(&demo)
	//l.svcCtx.Gorm.Statement.Where("picture_url LIKE ?", "%jinzhu%").Limit(10).Find(&demo)

	sql := l.svcCtx.Gorm.Statement.SQL.String()

	fmt.Println("SQL语句=", sql)
	fmt.Println("demo=", demo)

	//l.svcCtx.Gorm.
	return
}
*/
