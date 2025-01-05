package model

import (
	"time"
)

type UserModel struct {
	ID             uint      `gorm:"primaryKey;autoIncrement"`
	Name           string    `gorm:"column:name"`
	PhoneNumber    string    `gorm:"not null"`
	MobileProvider string    `gorm:"not null"`
	RefreshToken   string    `gorm:"not null"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (UserModel) TableName() string {
	return "users"
}
