package model

import (
	"time"
)

type ShopModel struct {
	ID             uint      `gorm:"primaryKey;autoIncrement"`
	CompanyID      int       `gorm:"column:company_id"`
	ImageID        int       `gorm:"column:image_id"`
	Name           string    `gorm:"column:name"`
	Location       string    `gorm:"column:location"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (ShopModel) TableName() string {
	return "shops"
}
