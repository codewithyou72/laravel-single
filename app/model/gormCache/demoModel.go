// goctl1.5.5非缓存方式生成的代码
package gormCache

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ DemoModel = (*customDemoModel)(nil)

type (
	// DemoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDemoModel.
	DemoModel interface {
		demoModel
		customDemoLogicModel
	}

	customDemoModel struct {
		*defaultDemoModel
	}

	customDemoLogicModel interface {
	}
)

// NewDemoModel returns a model for the database table.
func NewDemoModel(conn *gorm.DB, c cache.CacheConf) DemoModel {
	return &customDemoModel{
		defaultDemoModel: newDemoModel(conn, c),
	}
}

func (m *defaultDemoModel) customCacheKeys(data *Demo) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
