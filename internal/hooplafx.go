package internal

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hoopla/hoopla-api/app/config"
	"github.com/hoopla/hoopla-api/app/http/controller"
	"github.com/hoopla/hoopla-api/app/routes"
	"github.com/hoopla/hoopla-api/internal/repository"
	"github.com/hoopla/hoopla-api/internal/service"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(Server),
	routes.Modules,
	controller.Modules,
	service.Modules,
	repository.Modules,
	fx.Provide(config.NewMainDB),
)

func Server(lc fx.Lifecycle) *gin.Engine {
	config.LoadENV()
	appConf := config.NewAppConfig()

	router := gin.New()
	router.Use(gin.Recovery())

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
