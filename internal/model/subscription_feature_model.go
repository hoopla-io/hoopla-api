package model

import "time"

type SubscriptionFeatureModel struct {
	ID             uint      `gorm:"primaryKey;autoIncrement"`
	SubscriptionID uint      `gorm:"not null;index"`
	Feature        string    `gorm:"not null"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (SubscriptionFeatureModel) TableName() string {
	return "subscription_features"
}
