import (
	"context"
	"laravel-single/app/common/gormc"
	{{if .containsDbSql}}"database/sql"{{end}}
	{{if .time}}"time"{{end}}

	"gorm.io/gorm"
)
