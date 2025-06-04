package databasego

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func NewDatabase(conf Config) (*gorm.DB, error) {
	fmt.Println(conf)
	switch conf.DRIVER {
	case "postgresql":
		db := NewPostgresql(conf)
		return gorm.Open(*db, &gorm.Config{})
	default:
		return nil, errors.New("database driver not support")
	}
}
