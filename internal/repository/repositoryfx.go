package repository

import "go.uber.org/fx"

var Modules = fx.Options(
	fx.Provide(NewAuthRepository),
	fx.Provide(NewCoffeeRepository),
	fx.Provide(NewCompanyRepository),
	fx.Provide(NewCompanySocialRepository),
	fx.Provide(NewImageRepository),
	fx.Provide(NewShopCoffeeRepository),
	fx.Provide(NewShopPhoneRepository),
	fx.Provide(NewShopRepository),
	fx.Provide(NewShopWorkTimeRepository),
	fx.Provide(NewUserSubscriptionRepository),
	fx.Provide(NewSubscriptionRepository),
)
