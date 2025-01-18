package model

type SubscriptionModuleModel struct {
	ID             uint `gorm:"primaryKey;autoIncrement"`
	SubscriptionID uint `gorm:"not null;index"`
	ModuleID       uint `gorm:"not null;index"`
}

func (SubscriptionModuleModel) TableName() string {
	return "subscription_modules"
}
