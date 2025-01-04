package routes

import (
	"github.com/gin-gonic/gin"
)

func NewDashboardRoute(
	router *gin.Engine,
	// CompanyController dashboard2.CompanyController,
	// ShopController dashboard2.ShopController,
	// CoffeeController dashboard2.CoffeeController,
	// SubscriptionController dashboard2.SubscriptionController,
) {
	_ = router.Group("/dashboard")
	{
		//v1 := dashboard.Group("/v1")
		//{
		//	company := v1.Group("/company")
		//	{
		//		company.POST("/store", CompanyController.Store)
		//		company.GET("/show/:company_id", CompanyController.Show)
		//		company.GET("/list", CompanyController.List)
		//		company.PUT("/edit", CompanyController.Edit)
		//
		//		social := company.Group("/social")
		//		{
		//			social.POST("/store", CompanyController.StoreCompanySocial)
		//			social.GET("/show/:social_id", CompanyController.ShowCompanySocial)
		//			social.POST("/list", CompanyController.ListCompanySocials)
		//			social.PUT("/edit", CompanyController.EditCompanySocial)
		//		}
		//	}
		//
		//	shop := v1.Group("/shop")
		//	{
		//		shop.POST("/store", ShopController.Store)
		//		shop.GET("/show/:shop_id", ShopController.Show)
		//		shop.GET("/list", ShopController.List)
		//		shop.PUT("/edit", ShopController.Edit)
		//
		//		worktime := shop.Group("/worktime")
		//		{
		//			worktime.POST("/store", ShopController.StoreShopWorktime)
		//			worktime.POST("/list", ShopController.ListShopWorktimes)
		//			worktime.GET("/show/:worktime_id", ShopController.ShowWorktime)
		//			worktime.PUT("/edit", ShopController.EditShopWorktime)
		//		}
		//
		//		phone := shop.Group("/phone")
		//		{
		//			phone.POST("/store", ShopController.StoreShopPhone)
		//			phone.GET("/show/:phone_id", ShopController.ShowShopPhone)
		//			phone.POST("/list", ShopController.ListShopPhones)
		//			phone.PUT("/edit", ShopController.EditShopPhone)
		//		}
		//	}
		//
		//	coffee := v1.Group("/coffee")
		//	{
		//		coffee.POST("/store", CoffeeController.Store)
		//		coffee.GET("/show/:coffee_id", CoffeeController.Show)
		//		coffee.GET("/list", CoffeeController.List)
		//		coffee.PUT("/edit", CoffeeController.Edit)
		//	}
		//
		//	subscription := v1.Group("/subscription")
		//	{
		//		subscription.POST("/store", SubscriptionController.Store)
		//		subscription.GET("/show/:subscription_id", SubscriptionController.Show)
		//		subscription.GET("/list", SubscriptionController.List)
		//		subscription.PUT("/edit", SubscriptionController.Edit)
		//	}
		//}
	}
}
