package service

import "go.uber.org/fx"

var Modules = fx.Options(
	fx.Provide(NewUserService),
	fx.Provide(NewPartnerService),
	fx.Provide(NewShopService),
	fx.Provide(NewSubscriptionService),
)
