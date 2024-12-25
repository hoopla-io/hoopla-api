package model

import (
	"time"

	"gorm.io/gorm"
)

type ShopCoffeeModel struct {
	gorm.Model
	ID             uint      `gorm:"primaryKey;autoIncrement"`
	ShopID         int       `gorm:"column:shop_id"`
	CoffeeID       int       `gorm:"column:coffee_id"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (ShopCoffeeModel) TableName() string {
	return "shop_coffees"
}
