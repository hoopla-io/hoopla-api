package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/qahvazor/qahvazor/app/http/controller"
)

func NewRoute(controller *controller.Controller) *gin.Engine {
	router := gin.New()

	// router.Use(middleware.GetLanguageMiddleware())

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{

			auth := v1.Group("/auth")
			{
				auth.POST("/login", controller.AuthController.Login)
			}
			// uzum := v1.Group("/payment_uzum")
			// {
			// 	uzum.POST("/check", controller.UzumController.Check)
			// 	uzum.POST("/create", controller.UzumController.Create)
			// 	uzum.POST("/confirm", controller.UzumController.Confirm)
			// 	uzum.POST("/reverse", controller.UzumController.Reverse)
			// 	uzum.POST("/status", controller.UzumController.Status)
			// }

			// v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}
	}

	return router
}
