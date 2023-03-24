package model

import (
	"time"
)

const (
	MysqlNotDel = iota
	MysqlIsDel
)

type Model struct {
	ID        uint64    `gorm:"primaryKey"`
	IsDel     int8      `gorm:"is_del"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
