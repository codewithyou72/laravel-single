package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"laravel-single/app/common/gormc"
	"laravel-single/app/internal/config"
	"laravel-single/app/internal/middleware"
	gormNoCacheModel "laravel-single/app/model/gormNoCache"
	"log"

	gormModel "laravel-single/app/model/gormCache"
)

type ServiceContext struct {
	Config       config.Config
	PermMenuAuth rest.Middleware
	//Gorm         *gorm.DB
	DemoModel        gormModel.DemoModel        //Gorm配置步骤4
	DemoModelNoCache gormNoCacheModel.DemoModel //Gorm不使用缓存方式
}

func NewServiceContext(c config.Config) *ServiceContext {

	gormDB, err := gormc.ConnectMysql(c.GormCon) //Gorm配置步骤3
	if err != nil {
		log.Fatal(err)
	}

	return &ServiceContext{
		Config: c,
		//Gorm:         gormCon,
		DemoModel:        gormModel.NewDemoModel(gormDB, c.RedisCache), //Gorm配置步骤5
		DemoModelNoCache: gormNoCacheModel.NewDemoModel(gormDB),        //Gorm不使用缓存方式
		PermMenuAuth:     middleware.NewPermMenuAuthMiddleware().Handle,
	}
}
