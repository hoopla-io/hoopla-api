package model

import (
	"gorm.io/gorm"
	"time"
)

type PartnerModel struct {
	ID          uint           `gorm:"primaryKey;autoIncrement"`
	LogoId      uint           `gorm:"not null"`
	Name        string         `gorm:"not null"`
	Description string         `gorm:"not null"`
	CreatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (PartnerModel) TableName() string {
	return "partners"
}
