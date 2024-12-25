package databasego

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresql(conf Config) *gorm.Dialector {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.HOST,
		conf.PORT,
		conf.USER,
		conf.PASSWORD,
		conf.DATABASE,
	)

	psql := postgres.Open(dsn)
	return &psql
}
