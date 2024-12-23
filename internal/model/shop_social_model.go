package model

import "time"

type ShopSocialModel struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	ShopID    int       `gorm:"column:shop_id"`
	Platform  string    `gorm:"column:platform"`
	Url       string    `gorm:"column:url"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (ShopSocialModel) TableName() string {
	return "shop_socials"
}
