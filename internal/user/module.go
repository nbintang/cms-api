package user

import "go.uber.org/fx"

var Module = fx.Module(
	"User",
	fx.Provide(
		NewUserRepository,
		NewUserService,
		NewUserHandler,
	),
	fx.Invoke(RegisterUserRoutes),
)
