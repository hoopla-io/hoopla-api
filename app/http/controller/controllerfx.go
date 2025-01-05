package controller

import (
	"github.com/qahvazor/qahvazor/app/http/controller/api"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(api.NewAuthController),
	//fx.Provide(api.NewCompanyController),
	//fx.Provide(api.NewShopController),
	//fx.Provide(api.NewSubscriptionController),
	//fx.Provide(api.NewUserSubscriptionController),
)
