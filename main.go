package main

import (
	"log"

	"github.com/qahvazor/qahvazor/app/http/controller"
	"github.com/qahvazor/qahvazor/app/routes"
	"github.com/qahvazor/qahvazor/cmd"
	"github.com/qahvazor/qahvazor/config"
	"github.com/qahvazor/qahvazor/database/postgres"
	"github.com/qahvazor/qahvazor/internal/repository"
	"github.com/qahvazor/qahvazor/internal/service"
)

func main() {
	cfg := config.Load(".")
	db := postgres.NewPostgresDB(cfg)

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	controller := controller.NewController(service)

	router := routes.NewRoute(controller)

	srv := new(cmd.Server)
	if err := srv.Run(cfg.Port, router); err != nil {
		log.Fatalf(`Error occured while running http server: %s`, err.Error())
	}
}
