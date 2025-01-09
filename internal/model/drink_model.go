package model

import (
	"gorm.io/gorm"
	"time"
)

type DrinkModel struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	ImageID   uint           `gorm:"not null"`
	Name      string         `gorm:"not null"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Image *ImageModel `gorm:"foreignKey:image_id;primaryKey:id"`
}

func (DrinkModel) TableName() string {
	return "drinks"
}
