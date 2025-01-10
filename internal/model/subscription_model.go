package model

import (
	"gorm.io/gorm"
	"time"
)

type SubscriptionModel struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	Name      string         `gorm:"unique"`
	Days      uint           `gorm:"not null"`
	Price     float64        `gorm:"not null"`
	Currency  string         `gorm:"not null"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (SubscriptionModel) TableName() string {
	return "subscriptions"
}
