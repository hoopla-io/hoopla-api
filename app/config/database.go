package config

import (
	"github.com/hoopla/hoopla-api/pkg/databasego"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

func NewMainDB() *gorm.DB {
	mainDB := &databasego.Config{
		DRIVER:   os.Getenv("DB_DRIVER"),
		USER:     os.Getenv("DB_USER"),
		DATABASE: os.Getenv("DB_DATABASE"),
		PASSWORD: os.Getenv("DB_PASSWORD"),
		HOST:     os.Getenv("DB_HOST"),
	}
	mainDB.PORT, _ = strconv.Atoi(os.Getenv("DB_PORT"))

	db, err := databasego.NewDatabase(*mainDB)
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
