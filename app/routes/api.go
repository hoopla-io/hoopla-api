package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hoopla/hoopla-api/app/http/controller/api"
	api_user "github.com/hoopla/hoopla-api/app/http/controller/api/user"
	vendor_controllers "github.com/hoopla/hoopla-api/app/http/controller/vendors"
	"github.com/hoopla/hoopla-api/app/http/middleware"
	_ "github.com/hoopla/hoopla-api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewApiRoute(
	router *gin.Engine,
	AuthController *api.AuthController,
	PartnerController *api.PartnerController,
	UserController *api.UserController,
	ShopController *api.ShopController,
	SubscriptionController *api.SubscriptionController,
	UserOrderController *api_user.OrderController,
	PayController *api_user.PayController,

	IikoController *vendor_controllers.IikoController,
	PosterController *vendor_controllers.PosterController,
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

			user := v1.Group("/user")
			{
				user.GET("/get-me", middleware.JwtMiddleware(), UserController.GetMe)
				user.PATCH("/refresh-token", UserController.RefreshToken)
				user.POST("/logout", middleware.JwtMiddleware(), UserController.Logout)
				user.GET("/generate-qr-code", middleware.JwtMiddleware(), UserController.GenerateQRCode)
				user.DELETE("/deactivate", middleware.JwtMiddleware(), UserController.Deactivate)

				user_orders := user.Group("/orders")
				{
					user_orders.POST("/validate-order", middleware.JwtMiddleware(), UserOrderController.ValidateOrder)
					user_orders.POST("/create", middleware.JwtMiddleware(), UserOrderController.Create)
					user_orders.GET("/orders-list", middleware.JwtMiddleware(), UserOrderController.Orders)
					user_orders.GET("/drinks-stat", middleware.JwtMiddleware(), UserOrderController.DrinksStat)
				}

				pay_services := user.Group("/pay")
				{
					pay_services.GET("/services", PayController.Services)
					pay_services.GET("/top-up", PayController.TopUp)
				}
			}

			partners := v1.Group("/partners")
			{
				partners.GET("/", PartnerController.Partners)
				partners.GET("/partner", PartnerController.Partner)
			}

			shops := v1.Group("/shops")
			{
				shops.GET("/partner-shops", ShopController.PartnerShopList)
				shops.GET("/near-shops", ShopController.NearShops)
				shops.GET("/shop", ShopController.Shop)
			}

			subscriptions := v1.Group("/subscriptions")
			{
				subscriptions.GET("/", SubscriptionController.Subscriptions)
				subscriptions.POST("/buy", middleware.JwtMiddleware(), SubscriptionController.BuySubscription)
			}

			v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}

		vendors := api_routes.Group("/vendors")
		{
			iiko := vendors.Group("/iiko")
			{
				iiko.Match([]string{"GET", "POST"}, "/webhook", IikoController.Webhook)
			}

			poster := vendors.Group("/poster")
			{
				poster.GET("/oauth", PosterController.Oauth)
				poster.POST("/webhook", PosterController.Webhook)
			}
		}
	}
}
