package model

import (
	"gorm.io/gorm"
	"time"
)

type UserSubscriptionModel struct {
	ID             uint           `gorm:"primaryKey;autoIncrement"`
	UserID         uint           `gorm:"not null"`
	SubscriptionID uint           `gorm:"not null"`
	StartDate      int64          `gorm:"not null"`
	EndDate        int64          `gorm:"not null"`
	CreatedAt      time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`

	Subscription     *SubscriptionModel      `gorm:"foreignKey:SubscriptionID"`
	SubscriptionDays *[]SubscriptionDayModel `gorm:"foreignKey:subscription_id;references:subscription_id"`
}

func (UserSubscriptionModel) TableName() string {
	return "user_subscriptions"
}
