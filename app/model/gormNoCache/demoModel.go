package gormNoCache

import (
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
func NewDemoModel(conn *gorm.DB) DemoModel {
	return &customDemoModel{
		defaultDemoModel: newDemoModel(conn),
	}
}

func (m *defaultDemoModel) customCacheKeys(data *Demo) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
