package controller

import (
	"github.com/hoopla/hoopla-api/app/http/controller/api"
	api_user "github.com/hoopla/hoopla-api/app/http/controller/api/user"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(api.NewAuthController),
	fx.Provide(api.NewPartnerController),
	fx.Provide(api.NewUserController),
	fx.Provide(api.NewShopController),
	fx.Provide(api.NewSubscriptionController),
	fx.Provide(api_user.NewOrderController),
	fx.Provide(api_user.NewPayController),
)
