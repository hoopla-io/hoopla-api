package model

import "time"

type PartnerTokenModel struct {
	ID          uint      `gorm:"primary_key"`
	ShopID      uint      `gorm:"index"`
	AccessToken string    `gorm:"type:text"`
	ExpiresAt   time.Time `gorm:"index"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (PartnerTokenModel) TableName() string {
	return "partner_tokens"
}
