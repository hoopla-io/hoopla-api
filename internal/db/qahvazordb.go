package db

import (
	"github.com/qahvazor/qahvazor/app/config"
	"github.com/qahvazor/qahvazor/pkg/databasego"
	"gorm.io/gorm"
	"log"
)

func NewQahvazorDB() *gorm.DB {
	dbCfg := config.NewDatabaseConfig()
	db, err := databasego.NewDatabase(*dbCfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
