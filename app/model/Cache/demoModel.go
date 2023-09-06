// goctl1.5.5缓存方式生成的代码
package Cache

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DemoModel = (*customDemoModel)(nil)

type (
	// DemoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDemoModel.
	DemoModel interface {
		demoModel
	}

	customDemoModel struct {
		*defaultDemoModel
	}
)

// NewDemoModel returns a model for the database table.
func NewDemoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) DemoModel {
	return &customDemoModel{
		defaultDemoModel: newDemoModel(conn, c, opts...),
	}
}
