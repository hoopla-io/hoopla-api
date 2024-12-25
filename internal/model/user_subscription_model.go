package model

import "time"

type UserSubscription struct {
	ID             uint      `gorm:"primaryKey"`
	UserID         uint      `gorm:"not null"`
	SubscriptionID uint      `gorm:"not null"`
	EndDate        time.Time `gorm:"not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time

	Subscription Subscription `gorm:"foreignKey:SubscriptionID"`
}

func (UserSubscription) TableName() string {
	return "user_subscriptions"
}
