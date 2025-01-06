package model

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type ImageModel struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	Path      string         `gorm:"type:varchar(255);not null"`
	Filename  string         `gorm:"type:varchar(255);not null"`
	Ext       string         `gorm:"type:varchar(255);not null"`
	HashUID   string         `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (ImageModel) TableName() string {
	return "images"
}

func (m *ImageModel) GetUrl() *string {
	var url *string
	u := fmt.Sprintf("http://api.qahvazor.uz/%s/%s.%s", m.Path, m.Filename, m.Ext)
	url = &u
	return url
}
