package category

import (
	"rest-fiber/pkg/httpx"
	"rest-fiber/utils/enums"

	"go.uber.org/fx"
)


var Module = fx.Module(
	"category",
	fx.Provide(
		NewCategoryRepository,
		NewCategoryService,
		NewCategoryHandler,
		httpx.ProvideRoute[CategoryRouteParams, httpx.ProtectedRoute](
			NewCategoryRoutes,
			enums.RouteProtected,
		),
	),
)