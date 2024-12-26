package controller

import (
	"github.com/qahvazor/qahvazor/app/http/controller/api/v1"
	"github.com/qahvazor/qahvazor/app/http/controller/dashboard/v1"
	"github.com/qahvazor/qahvazor/internal/service"
)

type Controller struct {
	Api
	Dashboard
}

type Api struct {
	AuthController    *api.AuthController
	CompanyController *api.CompanyController
	ShopController    *api.ShopController
	SubscriptionController *api.SubscriptionController
	UserSubscriptionController *api.UserSubscriptionController
}

type Dashboard struct {
	CompanyController      *dashboard.CompanyController
	ShopController         *dashboard.ShopController
	CoffeeController       *dashboard.CoffeeController
	SubscriptionController *dashboard.SubscriptionController
}

func NewController(service *service.Service) *Controller {
	return &Controller{
		Api: Api{
			AuthController:    api.NewAuthController(service.AuthService),
			CompanyController: api.NewCompanyController(service.CompanyService),
			ShopController:    api.NewShopController(service.ShopService),
			SubscriptionController: api.NewSubscriptionController(service.SubscriptionService),
			UserSubscriptionController: api.NewUserSubscriptionController(service.UserSubscriptionService),
		},
		Dashboard: Dashboard{
			CompanyController:      dashboard.NewCompanyController(service.CompanyService),
			ShopController:         dashboard.NewShopController(service.ShopService),
			CoffeeController:       dashboard.NewCoffeeController(service.CoffeeService),
			SubscriptionController: dashboard.NewSubscriptionController(service.SubscriptionService),
		},
	}
}
