package main

import (
	"log"

	"github.com/qahvazor/qahvazor/app/config"
	"github.com/qahvazor/qahvazor/pkg/databasego"

	"github.com/qahvazor/qahvazor/app/http/controller"
	"github.com/qahvazor/qahvazor/app/routes"
	"github.com/qahvazor/qahvazor/cmd"
	"github.com/qahvazor/qahvazor/internal/repository"
	"github.com/qahvazor/qahvazor/internal/service"
)

// @title Qahvazor | Api
// @version 1.0.0
// @contact.email davronbekov.otabek@gmail.com
// @host api.qahvazor.uz
// @BasePath /api/v1
func main() {
	appCfg := config.NewAppConfig()
	dbCfg := config.NewDatabaseConfig()
	db, err := databasego.NewDatabase(*dbCfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	controller := controller.NewController(service)

	router := routes.NewRoute(controller)
	router.Static("/uploads", "./uploads")

	srv := new(cmd.Server)
	if err := srv.Run(appCfg, router); err != nil {
		log.Fatalf(`Error occured while running http server: %s`, err.Error())
	}
}
