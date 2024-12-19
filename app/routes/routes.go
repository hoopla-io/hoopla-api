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
				auth.POST("/confirm-sms", controller.AuthController.ConfirmSms)
				auth.POST("/resend-sms", controller.AuthController.ResendSms)
			}

			// v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}
	}

	return router
}
