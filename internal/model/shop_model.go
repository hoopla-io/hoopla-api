package model

import (
	"gorm.io/gorm"
	"time"
)

type ShopModel struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	PartnerID uint           `gorm:"not null;index"`
	Name      string         `gorm:"not null"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Attributes *[]ShopAttributeModel `gorm:"foreignKey:shop_id;references:id"`
}

func (ShopModel) TableName() string {
	return "shops"
}
