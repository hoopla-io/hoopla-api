package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hoopla/hoopla-api/internal"
	"go.uber.org/fx"
)

// @title Hoopla | Api
// @version 1.0.0
// @contact.email davronbekov.o@itv.uz
// @host api.hoopla.uz
// @BasePath /api/v1
func main() {
	app := fx.New(
		internal.Modules,
		fx.Invoke(func(*gin.Engine) {}),
	)

	app.Run()
}
