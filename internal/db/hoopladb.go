package db

import (
	"log"

	"github.com/hoopla/hoopla-api/app/config"
	"github.com/hoopla/hoopla-api/pkg/databasego"
	"gorm.io/gorm"
)

func NewHooplaDB() *gorm.DB {
	dbCfg := config.NewDatabaseConfig()
	db, err := databasego.NewDatabase(*dbCfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
