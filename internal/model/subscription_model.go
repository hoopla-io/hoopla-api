package model

import "time"

// Subscription представляет данные о подписке
type Subscription struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(255);not null"`
	CoffeeLimit int    `gorm:"not null"`
	Interval    int    `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// TableName задаёт имя таблицы в базе данных
func (Subscription) TableName() string {
	return "subscription"
}
