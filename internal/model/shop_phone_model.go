package model

import (
	"time"

	"gorm.io/gorm"
)

type ShopPhoneModel struct {
	gorm.Model
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	ShopID      int       `gorm:"column:shop_id"`
	PhoneNumber string    `gorm:"column:phone_number"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (ShopPhoneModel) TableName() string {
	return "shop_phones"
}
