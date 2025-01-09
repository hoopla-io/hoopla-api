package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/qahvazor/qahvazor/app/http/controller/api"
	"github.com/qahvazor/qahvazor/app/http/middleware"
	_ "github.com/qahvazor/qahvazor/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewApiRoute(
	router *gin.Engine,
	AuthController *api.AuthController,
	PartnerController *api.PartnerController,
	UserController *api.UserController,
	ShopController *api.ShopController,
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

			partners := v1.Group("/partners")
			{
				partners.GET("/", PartnerController.Partners)
				partners.GET("/partner", PartnerController.Partner)
			}

			shops := v1.Group("/shops")
			{
				shops.GET("/partner-shops", ShopController.PartnerShopList)
				shops.GET("/shop", ShopController.Shop)
			}

			user := v1.Group("/user")
			{
				user.GET("/get-me", middleware.JwtMiddleware(), UserController.GetMe)
				user.PATCH("/refresh-token", UserController.RefreshToken)
				user.POST("/logout", middleware.JwtMiddleware(), UserController.Logout)
				user.GET("/generate-qr-code", middleware.JwtMiddleware(), UserController.GenerateQRCode)
			}

			//company := v1.Group("/company")
			//{
			//	company.GET("/list", CompanyController.GetCompanyList)
			//	company.GET("/shops", CompanyController.GetCompanyShopsList)
			//}
			//
			//shop := v1.Group("/shop")
			//{
			//	shop.GET("/detail", ShopController.GetShopDetails)
			//}
			//
			//subscriptions := v1.Group("/subscriptions")
			//{
			//	subscriptions.GET("/", SubscriptionController.GetAllSubscriptions)
			//	subscriptions.GET("/:id", SubscriptionController.GetSubscriptionByID)
			//}
			//
			//userSubscriptions := v1.Group("/user-subscriptions")
			//{
			//	userSubscriptions.GET("/:user_id", UserSubscriptionController.GetUserActiveSubscription)
			//	userSubscriptions.POST("/", UserSubscriptionController.AssignSubscriptionToUser)
			//}

			v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}
	}
}
