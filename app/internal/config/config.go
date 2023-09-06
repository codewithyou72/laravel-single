package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"laravel-single/app/common/gormc"
)

type Config struct {
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}

	GormCon    gormc.Mysql //Gorm配置步骤2
	RedisCache cache.CacheConf

	PublicDir string //index.go设置对外公开的public文件夹名称、你可以根据自己需要修改这个文件夹
}
