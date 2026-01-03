package internal

import "go.uber.org/fx"

var Module = fx.Module(
	"App",
	fx.Provide(NewApp),
	fx.Invoke(RunApp),
)