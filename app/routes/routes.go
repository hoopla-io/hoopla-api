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

			company := v1.Group("/company")
			{
				company.GET("/get", controller.CompanyController.GetCompany)
				company.GET("/get-list", controller.CompanyController.GetList)
			}

			subscriptions := v1.Group("/subscriptions")
			{
				subscriptions.GET("/", controller.SubscriptionController.GetAllSubscriptions)
				subscriptions.GET("/:id", controller.SubscriptionController.GetSubscriptionByID)
			}

			userSubscriptions := v1.Group("/user-subscriptions")
			{
				userSubscriptions.GET("/:user_id", controller.UserSubscriptionController.GetUserActiveSubscription)
				userSubscriptions.POST("/", controller.UserSubscriptionController.AssignSubscriptionToUser)
			}

			// shops := v1.Group("/shop")
			// {
			// 	// shops.POST("/create", controller.ShopController.CreateShop)
			// }

			// v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}
	}

	dashboard := api.Group("/dashboard")
	{
		v1 := api.Group("/v1")
		{
			company := dashboard.Group("/company")
			{
				company.POST("/create", controller.CompanyController.CreateCompany)
			}

			shop := v1.Group("/shop")
			{
				shop.POST("/create", controller.ShopController.CreateShop)

				worktime := shop.Group("/worktime")
				{
					worktime.POST("/create", controller.ShopController.CreateShopWorkTime)
				}

				phone := shop.Group("/phone")
				{
					phone.POST("/create", controller.ShopController.CreateShopPhone)
				}

				social := shop.Group("/social")
				{
					social.POST("/create", controller.ShopController.CreateShopSocial)
				}
			}

			coffee := v1.Group("/coffee")
			{
				coffee.POST("/create")
			}
		}
	}

	return router
}
