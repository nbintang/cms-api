package config

import "go.uber.org/fx"


var Module = fx.Module(
	"Config",
	fx.Provide(NewEnv),
	fx.Provide(NewDatabase),
	fx.Provide(NewDBLogger),
	fx.Provide(NewAppLogger),
)