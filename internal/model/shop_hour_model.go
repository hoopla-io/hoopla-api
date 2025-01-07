package model

import (
	"gorm.io/gorm"
	"time"
)

type ShopHourModel struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	ShopID    uint           `gorm:"not null;index"`
	WeekDay   string         `gorm:"not null;index"`
	OpenAt    string         `gorm:"not null"`
	CloseAt   string         `gorm:"not null"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (ShopHourModel) TableName() string {
	return "shop_hours"
}
