package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hoopla/hoopla-api/internal"
	"go.uber.org/fx"
)

// @title Hoopla | Api
// @version 1.0.0
// @contact.email davronbekov.o@itv.uz
// @host 127.0.0.1:8000
// @BasePath /api
func main() {
	app := fx.New(
		internal.Modules,
		fx.Invoke(func(*gin.Engine) {}),
	)

	app.Run()
}
