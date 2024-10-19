package model

import (
	"gin-demo/pkg/utils"
	"gorm.io/gorm"
	"time"
)

// BaseModel 是一个通用的模型，包含 ID、CreatedAt、UpdatedAt 和 DeletedAt 字段

type BaseModel struct {
	// 雪花ID
	ID int64 `gorm:"primarykey;column:id"`
	// 使用 `autoCreateTime` 标签，在创建记录时自动设置为当前时间
	CreatedAt time.Time `gorm:"type:timestamp;autoCreateTime"`
	// 使用 `autoUpdateTime` 标签，在更新记录时自动设置为当前时间
	UpdatedAt time.Time      `gorm:"type:timestamp;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp;index"`
}

// BeforeCreate 在创建记录之前自动生成雪花 ID
func (b *BaseModel) BeforeCreate(db *gorm.DB) error {
	id, err := utils.GenerateSnowflakeId()
	if err != nil {
		return err
	}
	b.ID = id
	return nil
}
