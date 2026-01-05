package internal

import (
	"rest-fiber/internal/auth"
	"rest-fiber/internal/user"

	"go.uber.org/fx"
)

var RunApp = func(lc fx.Lifecycle, a *App) { a.Run(lc) }
var Module = fx.Module(
	"app",
	user.Module,
	auth.Module,
	fx.Provide(NewApp),
	fx.Invoke(
		RegisterAllRoutes,
		RunApp,
	),
)
