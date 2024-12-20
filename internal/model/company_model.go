package model

import "time"

type CompanyModel struct {
	ID             uint      `gorm:"primaryKey;autoIncrement"`
	ImageID        int       `gorm:"column:image_id"`
	Name           string    `gorm:"column:name"`
	Description    string    `gorm:"column:description"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (CompanyModel) TableName() string {
	return "company"
}
