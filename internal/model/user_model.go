package model

import (
	"gorm.io/gorm"
	"math"
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

	Debit  int64 `gorm:"not null"`
	Credit int64 `gorm:"not null"`
}

func (UserModel) TableName() string {
	return "users"
}

func (a UserModel) GetBalance() float64 {
	balance := a.Debit - a.Credit
	value := math.Round(float64(balance) / float64(100))
	return value
}
