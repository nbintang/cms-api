package auth

import "go.uber.org/fx"

var Module = fx.Module(
	"Auth",
	fx.Provide(
		NewAuthService,
		NewAuthHandler,
	),
	fx.Invoke(RegisterAuthRoutes),
)
