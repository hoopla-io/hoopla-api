package model

import (
	"gorm.io/gorm"
	"time"
)

type UserOrderModel struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	PartnerID uint           `gorm:"not null;index"`
	ShopID    uint           `gorm:"not null;index"`
	UserID    uint           `gorm:"not null;index"`
	DrinkID   uint           `gorm:"index"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Status        string  `gorm:"index;default:pending"`
	Vendor        string  `gorm:"index'"`
	VendorOrderID string  `gorm:"index"`
	ProductPrice  float32 `gorm:"not null;default:0.0"`

	Partner *PartnerModel `gorm:"foreignKey:partner_id;references:id"`
	Shop    *ShopModel    `gorm:"foreignKey:shop_id;references:id"`
	Drink   *DrinkModel   `gorm:"foreignKey:drink_id;references:id"`
}

func (UserOrderModel) TableName() string {
	return "user_orders"
}
