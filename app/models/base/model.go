package base

import (
	"gorm.io/gorm"
	"thh/helpers"
	"time"
)

// BaseModel 模型基类
type BaseModel struct {
	ID        uint64    `gorm:"column:id;primaryKey;autoIncrement;not null"`
	CreatedAt time.Time `gorm:"column:created_at;index"`
	UpdatedAt time.Time `gorm:"column:updated_at;index"`
	// 支持 gorm 软删除
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

// GetStringID 获取 ID 的字符串格式
func (itself BaseModel) GetStringID() string {
	return helpers.ToString(itself.ID)
}

type Repository interface {
	getModel()
}
