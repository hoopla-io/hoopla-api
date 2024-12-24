package model

import (
	"time"

	"gorm.io/gorm"
)

type CompanySocialModel struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	CompanyID int       `gorm:"column:company_id"`
	Platform  string    `gorm:"column:platform"`
	Url       string    `gorm:"column:url"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (CompanySocialModel) TableName() string {
	return "company_socials"
}
