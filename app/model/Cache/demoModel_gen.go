// Code generated by goctl. DO NOT EDIT.

package Cache

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	demoFieldNames          = builder.RawFieldNames(&Demo{})
	demoRows                = strings.Join(demoFieldNames, ",")
	demoRowsExpectAutoSet   = strings.Join(stringx.Remove(demoFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	demoRowsWithPlaceHolder = strings.Join(stringx.Remove(demoFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheLaravelSingleDemoIdPrefix     = "cache:laravelSingle:demo:id:"
	cacheLaravelSingleDemoUserIdPrefix = "cache:laravelSingle:demo:userId:"
)

type (
	demoModel interface {
		Insert(ctx context.Context, data *Demo) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Demo, error)
		FindOneByUserId(ctx context.Context, userId int64) (*Demo, error)
		Update(ctx context.Context, data *Demo) error
		Delete(ctx context.Context, id int64) error
	}

	defaultDemoModel struct {
		sqlc.CachedConn
		table string
	}

	Demo struct {
		Id         int64     `db:"id"`
		UserId     int64     `db:"user_id"`     // 关联user表主键id
		PictureUrl string    `db:"picture_url"` // 图片链接[image]
		VideoUrl   string    `db:"video_url"`   // 视频链接[video]
		IsDisplay  int64     `db:"is_display"`  // (用户端)是否显示[check]：1=不显示、2=显示、0=全部：默认=1
		Sort       int64     `db:"sort"`        // (降序)排序
		CreatedAt  time.Time `db:"created_at"`  // 创建时间
		UpdatedAt  time.Time `db:"updated_at"`  // 更新时间
		IsDel      int64     `db:"is_del"`      // 是否删除:1=否、2=是、0=全部：默认=1
	}
)

func newDemoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultDemoModel {
	return &defaultDemoModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`demo`",
	}
}

func (m *defaultDemoModel) withSession(session sqlx.Session) *defaultDemoModel {
	return &defaultDemoModel{
		CachedConn: m.CachedConn.WithSession(session),
		table:      "`demo`",
	}
}

func (m *defaultDemoModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	laravelSingleDemoIdKey := fmt.Sprintf("%s%v", cacheLaravelSingleDemoIdPrefix, id)
	laravelSingleDemoUserIdKey := fmt.Sprintf("%s%v", cacheLaravelSingleDemoUserIdPrefix, data.UserId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, laravelSingleDemoIdKey, laravelSingleDemoUserIdKey)
	return err
}

func (m *defaultDemoModel) FindOne(ctx context.Context, id int64) (*Demo, error) {
	laravelSingleDemoIdKey := fmt.Sprintf("%s%v", cacheLaravelSingleDemoIdPrefix, id)
	var resp Demo
	err := m.QueryRowCtx(ctx, &resp, laravelSingleDemoIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", demoRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultDemoModel) FindOneByUserId(ctx context.Context, userId int64) (*Demo, error) {
	laravelSingleDemoUserIdKey := fmt.Sprintf("%s%v", cacheLaravelSingleDemoUserIdPrefix, userId)
	var resp Demo
	err := m.QueryRowIndexCtx(ctx, &resp, laravelSingleDemoUserIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", demoRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userId); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultDemoModel) Insert(ctx context.Context, data *Demo) (sql.Result, error) {
	laravelSingleDemoIdKey := fmt.Sprintf("%s%v", cacheLaravelSingleDemoIdPrefix, data.Id)
	laravelSingleDemoUserIdKey := fmt.Sprintf("%s%v", cacheLaravelSingleDemoUserIdPrefix, data.UserId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, demoRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.PictureUrl, data.VideoUrl, data.IsDisplay, data.Sort, data.IsDel)
	}, laravelSingleDemoIdKey, laravelSingleDemoUserIdKey)
	return ret, err
}

func (m *defaultDemoModel) Update(ctx context.Context, newData *Demo) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	laravelSingleDemoIdKey := fmt.Sprintf("%s%v", cacheLaravelSingleDemoIdPrefix, data.Id)
	laravelSingleDemoUserIdKey := fmt.Sprintf("%s%v", cacheLaravelSingleDemoUserIdPrefix, data.UserId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, demoRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.UserId, newData.PictureUrl, newData.VideoUrl, newData.IsDisplay, newData.Sort, newData.IsDel, newData.Id)
	}, laravelSingleDemoIdKey, laravelSingleDemoUserIdKey)
	return err
}

func (m *defaultDemoModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheLaravelSingleDemoIdPrefix, primary)
}

func (m *defaultDemoModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", demoRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultDemoModel) tableName() string {
	return m.table
}
