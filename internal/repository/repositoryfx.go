package repository

import "go.uber.org/fx"

var Modules = fx.Options(
	fx.Provide(NewUserRepository),
	fx.Provide(NewPartnerRepository),
	fx.Provide(NewShopRepository),
	fx.Provide(NewPartnerDrinkRepository),
	fx.Provide(NewPartnerTokenRepository),
	fx.Provide(NewSubscriptionRepository),
	fx.Provide(NewUserOrderRepository),
	fx.Provide(NewUserSubscriptionRepository),
)
