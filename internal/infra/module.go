package infra

import "go.uber.org/fx"

var Module = fx.Module(
	"infra",
	fx.Provide(
		NewDatabase,
		NewLogger,
		NewDBLogger,
		NewValidator,
		NewTokenService,
		NewEmailService,
		NewRedisService,
	),
	fx.Invoke(RegisterRedisLifecycle),
)
