package controller

import "github.com/qahvazor/qahvazor/internal/service"

type Controller struct {
	AuthController
	CompanyController
	ShopController
}

func NewController(service *service.Service) *Controller {
	return &Controller{
		AuthController: *NewAuthController(service.AuthService),
		CompanyController: *NewCompanyController(service.CompanyService),
		ShopController: *NewShopController(service.ShopService),
	}
}
