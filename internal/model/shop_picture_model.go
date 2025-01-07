package model

import (
	"gorm.io/gorm"
	"time"
)

type ShopPictureModel struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	ShopID    uint           `gorm:"not null;index"`
	ImageID   uint           `gorm:"not null;index"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Image *ImageModel `gorm:"foreignKey:image_id;primaryKey:id"`
}

func (ShopPictureModel) TableName() string {
	return "shop_pictures"
}
