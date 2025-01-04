package internal

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qahvazor/qahvazor/app/config"
	"github.com/qahvazor/qahvazor/app/http/controller"
	"github.com/qahvazor/qahvazor/app/routes"
	"github.com/qahvazor/qahvazor/internal/db"
	"github.com/qahvazor/qahvazor/internal/repository"
	"github.com/qahvazor/qahvazor/internal/service"
	"go.uber.org/fx"
	"net"
	"net/http"
	"time"
)

var Modules = fx.Options(
	fx.Provide(Server),
	routes.Modules,
	controller.Modules,
	service.Modules,
	repository.Modules,
	db.Modules,
)

func Server(lc fx.Lifecycle) *gin.Engine {
	appConf := config.NewAppConfig()

	router := gin.Default()

	srv := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", appConf.HOST, appConf.PORT),
		Handler:        router,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				fmt.Printf("[%s] Failed to Listen HTTP Server at %s\n", appConf.APP, srv.Addr)
				return err
			}
			go func() {
				err := srv.Serve(ln)
				if err != nil {
					fmt.Printf("[%s] Failed to Serve HTTP Server at %s\n", appConf.APP, srv.Addr)
				}
			}()
			fmt.Printf("[%s]Succeeded to start HTTP Server at %s\n", appConf.APP, srv.Addr)
			return nil

		},
		OnStop: func(ctx context.Context) error {
			err := srv.Shutdown(ctx)
			if err != nil {
				return err
			}
			fmt.Printf("[%s] HTTP Server is stopped", appConf.APP)
			return nil
		},
	})

	return router
}
