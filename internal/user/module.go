package user

import (
	"rest-fiber/internal/setup"

	"go.uber.org/fx"
)

var Module = fx.Module(
	"user",
	fx.Provide(
		NewUserRepository,
		NewUserService,
		NewUserHandler,
		setup.ProtectedRouteProvider[UserHandler](NewUserRoute),
	),
)
