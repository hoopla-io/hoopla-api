package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/qahvazor/qahvazor/app/http/controller/api"
	_ "github.com/qahvazor/qahvazor/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewApiRoute(
	router *gin.Engine,
	AuthController *api.AuthController,
	CompanyController *api.CompanyController,
	ShopController *api.ShopController,
	SubscriptionController *api.SubscriptionController,
	UserSubscriptionController *api.UserSubscriptionController,
	QRCodeController *api.QRCodeController,
) {
	api_routes := router.Group("/api")
	{
		v1 := api_routes.Group("/v1")
		{

			auth := v1.Group("/auth")
			{
				auth.POST("/login", AuthController.Login)
				auth.POST("/confirm-sms", AuthController.ConfirmSms)
				auth.POST("/resend-sms", AuthController.ResendSms)
			}

			company := v1.Group("/company")
			{
				company.GET("/list", CompanyController.GetCompanyList)
				company.GET("/shops", CompanyController.GetCompanyShopsList)
			}

			shop := v1.Group("/shop")
			{
				shop.GET("/detail", ShopController.GetShopDetails)
			}

			subscriptions := v1.Group("/subscriptions")
			{
				subscriptions.GET("/", SubscriptionController.GetAllSubscriptions)
				subscriptions.GET("/:id", SubscriptionController.GetSubscriptionByID)
			}

			userSubscriptions := v1.Group("/user-subscriptions")
			{
				userSubscriptions.GET("/:user_id", UserSubscriptionController.GetUserActiveSubscription)
				userSubscriptions.POST("/", UserSubscriptionController.AssignSubscriptionToUser)
			}

			qr := v1.Group("/qr")
			{
				qr.GET("/generate", QRCodeController.GenerateQRCode)
				qr.GET("/decode", QRCodeController.DecodeQRCode)
			}

			v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}
	}
}
