package model

import (
	"gorm.io/gorm"
	"time"
)

type SubscriptionModel struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	Name      string         `gorm:"type:varchar(255);not null"`
	Days      uint           `gorm:"not null"`
	Price     float64        `gorm:"not null"`
	Currency  string         `gorm:"type:varchar(50);not null"`
	CupsDay   uint           `gorm:"column:cups_day"`
	Priority  int            `gorm:"default:0"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Features []SubscriptionFeatureModel `gorm:"foreignKey:subscription_id;references:id"`
	WeekDays []SubscriptionDayModel     `gorm:"foreignKey:subscription_id;references:id"`
}

func (SubscriptionModel) TableName() string {
	return "subscriptions"
}
