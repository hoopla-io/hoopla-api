package model

import (
	"time"

	"gorm.io/gorm"
)

type ShopWorkTimeModel struct {
	gorm.Model
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	ShopID      int       `gorm:"column:shop_id"`
	DayRange    string    `gorm:"column:day_range"`
	OpeningTime string    `gorm:"column:opening_time"`
	ClosingTime string    `gorm:"column:closing_time"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (ShopWorkTimeModel) TableName() string {
	return "shop_worktimes"
}
