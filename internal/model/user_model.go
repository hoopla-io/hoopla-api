package model

import (
	"gorm.io/gorm"
	"time"
)

type UserModel struct {
	ID             uint           `gorm:"primaryKey;autoIncrement"`
	Name           string         `gorm:"column:name"`
	PhoneNumber    string         `gorm:"not null"`
	MobileProvider string         `gorm:"not null"`
	RefreshToken   string         `gorm:"not null"`
	CreatedAt      time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func (UserModel) TableName() string {
	return "users"
}
