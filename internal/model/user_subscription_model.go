package model

import "time"

// UserSubscription связывает пользователя с подпиской
type UserSubscription struct {
	ID             uint      `gorm:"primaryKey"`
	UserID         uint      `gorm:"not null"`
	SubscriptionID uint      `gorm:"not null"`
	EndDate        time.Time `gorm:"not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time

	Subscription SubscriptionModel `gorm:"foreignKey:SubscriptionID"`
}

// TableName задаёт имя таблицы в базе данных
func (UserSubscription) TableName() string {
	return "user_subscriptions"
}
