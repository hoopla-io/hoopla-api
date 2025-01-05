package repository

import "go.uber.org/fx"

var Modules = fx.Options(
	fx.Provide(NewUserRepository),
	fx.Provide(NewPartnerRepository),
)
