package model

import (
	"gorm.io/gorm"
	"time"
)

type ShopModel struct {
	ID           uint           `gorm:"primaryKey;autoIncrement"`
	ImageID      uint           `gorm:"not null"`
	PartnerID    uint           `gorm:"not null;index"`
	Name         string         `gorm:"not null"`
	LocationLat  float64        `gorm:"not null;index"`
	LocationLong float64        `gorm:"not null;index"`
	Distance     float64        `gorm:"index"`
	CreatedAt    time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`

	Attributes   *[]ShopAttributeModel `gorm:"foreignKey:shop_id;references:id"`
	WorkingHours *[]ShopHourModel      `gorm:"foreignKey:shop_id;references:id"`
	Pictures     *[]ShopPictureModel   `gorm:"foreignKey:shop_id;references:id"`

	PartnerDrinks     *[]PartnerDrinkModel     `gorm:"foreignKey:partner_id;references:partner_id"`
	PartnerAttributes *[]PartnerAttributeModel `gorm:"foreignKey:partner_id;references:partner_id"`

	Modules *[]ShopModuleModel `gorm:"foreignKey:shop_id;references:id"`

	Image *ImageModel `gorm:"foreignKey:id;references:image_id"`
}

func (ShopModel) TableName() string {
	return "shops"
}
