package controller

import "github.com/qahvazor/qahvazor/internal/service"

type Controller struct {
	AuthController
}

func NewController(service *service.Service) *Controller {
	return &Controller{
		AuthController: *NewAuthController(service.AuthService),
	}
}
