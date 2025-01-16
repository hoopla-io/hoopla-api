package db

import "go.uber.org/fx"

var Modules = fx.Options(
	fx.Provide(NewHooplaDB),
)
