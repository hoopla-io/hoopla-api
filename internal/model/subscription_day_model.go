package model

import (
	"gorm.io/gorm"
	"time"
)

var WeekDays = map[int16]map[string]string{
	1: {"en": "Monday", "ru": "Понедельник", "uz": "Dushanba"},
	2: {"en": "Tuesday", "ru": "Вторник", "uz": "Seshanba"},
	3: {"en": "Wednesday", "ru": "Среда", "uz": "Chorshanba"},
	4: {"en": "Thursday", "ru": "Четверг", "uz": "Payshanba"},
	5: {"en": "Friday", "ru": "Пятница", "uz": "Juma"},
	6: {"en": "Saturday", "ru": "Суббота", "uz": "Shanba"},
	7: {"en": "Sunday", "ru": "Воскресенье", "uz": "Yakshanba"},
}

type SubscriptionDayModel struct {
	ID             uint           `gorm:"primaryKey;autoIncrement"`
	SubscriptionID uint           `gorm:"not null;index"`
	Day            int16          `gorm:"not null"`
	CreatedAt      time.Time      `gorm:"not null"`
	UpdatedAt      time.Time      `gorm:"not null"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func (SubscriptionDayModel) TableName() string {
	return "subscription_days"
}

func (m *SubscriptionDayModel) GetName(language string) string {
	return WeekDays[m.Day][language]
}
