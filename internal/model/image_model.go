package model

import (
	"time"

	"github.com/google/uuid"
)

type ImageModel struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Path      string    `gorm:"type:varchar(255);not null"`
	Filename  string    `gorm:"type:varchar(255);not null"`
	Ext       string    `gorm:"type:varchar(255);not null"`
	HashUID   uuid.UUID `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (ImageModel) TableName() string {
	return "images"
}
