package model

import (
	"time"

	"gorm.io/gorm"
)

type CoffeeModel struct {
	gorm.Model
	ID             uint      `gorm:"primaryKey;autoIncrement"`
	ImageID        int       `gorm:"column:image_id"`
	Name           string    `gorm:"column:name"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (CoffeeModel) TableName() string {
	return "coffee"
}
