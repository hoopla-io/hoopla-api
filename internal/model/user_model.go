package model

import (
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	ID             uint      `gorm:"primaryKey;autoIncrement"`
	Name           string    `gorm:"column:name"`
	PhoneNumber    string    `gorm:"not null"`
	MobileProvider string    `gorm:"not null"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (UserModel) TableName() string {
	return "users"
}
