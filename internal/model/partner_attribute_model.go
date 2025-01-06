package model

import (
	"gorm.io/gorm"
	"time"
)

type PartnerAttributeModel struct {
	ID             uint           `gorm:"primaryKey;autoIncrement"`
	PartnerID      uint           `gorm:"not null;index"`
	AttributeKey   string         `gorm:"not null;index"`
	AttributeValue string         `gorm:"not null"`
	CreatedAt      time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func (PartnerAttributeModel) TableName() string {
	return "partner_attributes"
}
