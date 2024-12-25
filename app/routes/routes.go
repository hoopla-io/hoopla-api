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
				auth.POST("/login", controller.Api.AuthController.Login)
				auth.POST("/confirm-sms", controller.Api.AuthController.ConfirmSms)
				auth.POST("/resend-sms", controller.Api.AuthController.ResendSms)
			}

			company := v1.Group("/company") 
			{
				company.GET("/list", controller.Api.CompanyController.GetCompanyList)
				company.POST("/shops", controller.Api.CompanyController.GetCompanyShopsList)
			}

			shop := v1.Group("/shop")
			{
				shop.POST("/detail", controller.Api.ShopController.GetShopDetails)
			}
			
			// v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}
	}

	dashboard := router.Group("/dashboard") 
	{
		v1 := dashboard.Group("/v1")
		{
			company := v1.Group("/company")
			{
				company.POST("/store", controller.Dashboard.CompanyController.Store)
				company.GET("/show/:company_id", controller.Dashboard.CompanyController.Show)
				company.GET("/list", controller.Dashboard.CompanyController.List)
				company.PUT("/edit", controller.Dashboard.CompanyController.Edit)

				social := company.Group("/social")
				{
					social.POST("/store", controller.Dashboard.CompanyController.StoreCompanySocial)
					social.GET("/show/:social_id", controller.Dashboard.CompanyController.ShowCompanySocial)
					social.POST("/list", controller.Dashboard.CompanyController.ListCompanySocials)
					social.PUT("/edit", controller.Dashboard.CompanyController.EditCompanySocial)
				}
			}

			shop := v1.Group("/shop")
			{
				shop.POST("/store", controller.Dashboard.ShopController.Store)
				shop.GET("/show/:shop_id", controller.Dashboard.ShopController.Show)
				shop.GET("/list", controller.Dashboard.ShopController.List)
				shop.PUT("/edit", controller.Dashboard.ShopController.Edit)

				worktime := shop.Group("/worktime") 
				{
					worktime.POST("/store", controller.Dashboard.ShopController.StoreShopWorktime)
					worktime.POST("/list", controller.Dashboard.ShopController.ListShopWorktimes)
					worktime.GET("/show/:worktime_id", controller.Dashboard.ShopController.ShowWorktime)
					worktime.PUT("/edit", controller.Dashboard.ShopController.EditShopWorktime)
				}

				phone := shop.Group("/phone")
				{
					phone.POST("/store", controller.Dashboard.ShopController.StoreShopPhone)
					phone.GET("/show/:phone_id", controller.Dashboard.ShopController.ShowShopPhone)
					phone.POST("/list", controller.Dashboard.ShopController.ListShopPhones)
					phone.PUT("/edit", controller.Dashboard.ShopController.EditShopPhone)
				}
			}

			coffee := v1.Group("/coffee")
			{
				coffee.POST("/store", controller.Dashboard.CoffeeController.Store)
				coffee.GET("/show/:coffee_id", controller.Dashboard.CoffeeController.Show)
				coffee.GET("/list", controller.Dashboard.CoffeeController.List)
				coffee.PUT("/edit", controller.Dashboard.CoffeeController.Edit)
			}

			subscription := v1.Group("/subscription")
			{
				subscription.POST("/store", controller.Dashboard.SubscriptionController.Store)
				subscription.GET("/show/:subscription_id", controller.Dashboard.SubscriptionController.Show)
				subscription.GET("/list", controller.Dashboard.SubscriptionController.List)
				subscription.PUT("/edit", controller.Dashboard.SubscriptionController.Edit)
			}
		}
	}

	return router
}
