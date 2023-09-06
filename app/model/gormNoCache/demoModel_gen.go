// Code generated by goctl. DO NOT EDIT!

package gormNoCache

import (
	"context"
	"laravel-single/app/common/gormc"

	"time"

	"gorm.io/gorm"
)

type (
	demoModel interface {
		Insert(ctx context.Context, tx *gorm.DB, data *Demo) error

		FindOne(ctx context.Context, id int64) (*Demo, error)
		FindOneByUserId(ctx context.Context, userId int64) (*Demo, error)
		Update(ctx context.Context, tx *gorm.DB, data *Demo) error

		Delete(ctx context.Context, tx *gorm.DB, id int64) error
		Transaction(ctx context.Context, fn func(db *gorm.DB) error) error
	}

	defaultDemoModel struct {
		conn  *gorm.DB
		table string
	}

	Demo struct {
		Id         int64     `gorm:"column:id"`
		UserId     int64     `gorm:"column:user_id"`     // 关联user表主键id
		PictureUrl string    `gorm:"column:picture_url"` // 图片链接[image]
		VideoUrl   string    `gorm:"column:video_url"`   // 视频链接[video]
		IsDisplay  int64     `gorm:"column:is_display"`  // (用户端)是否显示[check]：1=不显示、2=显示、0=全部：默认=1
		Sort       int64     `gorm:"column:sort"`        // (降序)排序
		CreatedAt  time.Time `gorm:"column:created_at"`  // 创建时间
		UpdatedAt  time.Time `gorm:"column:updated_at"`  // 更新时间
		IsDel      int64     `gorm:"column:is_del"`      // 是否删除:1=否、2=是、0=全部：默认=1
	}
)

func (Demo) TableName() string {
	return "`demo`"
}

func newDemoModel(conn *gorm.DB) *defaultDemoModel {
	return &defaultDemoModel{
		conn:  conn,
		table: "`demo`",
	}
}

func (m *defaultDemoModel) Insert(ctx context.Context, tx *gorm.DB, data *Demo) error {
	db := m.conn
	if tx != nil {
		db = tx
	}
	err := db.WithContext(ctx).Save(&data).Error
	return err
}

func (m *defaultDemoModel) FindOne(ctx context.Context, id int64) (*Demo, error) {
	var resp Demo
	err := m.conn.WithContext(ctx).Model(&Demo{}).Where("`id` = ?", id).Take(&resp).Error
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultDemoModel) FindOneByUserId(ctx context.Context, userId int64) (*Demo, error) {
	var resp Demo
	err := m.conn.WithContext(ctx).Model(&Demo{}).Where("`user_id` = ?", userId).Take(&resp).Error
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultDemoModel) Update(ctx context.Context, tx *gorm.DB, data *Demo) error {
	db := m.conn
	if tx != nil {
		db = tx
	}
	err := db.WithContext(ctx).Save(data).Error
	return err
}

func (m *defaultDemoModel) Delete(ctx context.Context, tx *gorm.DB, id int64) error {
	db := m.conn
	if tx != nil {
		db = tx
	}
	err := db.WithContext(ctx).Delete(&Demo{}, id).Error

	return err
}

func (m *defaultDemoModel) Transaction(ctx context.Context, fn func(db *gorm.DB) error) error {
	return m.conn.WithContext(ctx).Transaction(fn)
}