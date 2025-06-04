package databasego

import (
	"errors"
	"gorm.io/gorm"
)

func NewDatabase(conf Config) (*gorm.DB, error) {
	switch conf.DRIVER {
	case "postgresql":
		db := NewPostgresql(conf)
		return gorm.Open(*db, &gorm.Config{})
	default:
		return nil, errors.New("database driver not support")
	}
}
