package model

import "time"

type Subscription struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(255);not null"`
	CoffeeLimit int    `gorm:"not null"`
	Interval    int    `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Subscription) TableName() string {
	return "subscription"
}
