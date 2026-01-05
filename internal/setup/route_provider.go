package setup

import (
	"rest-fiber/internal/contract"

	"go.uber.org/fx"
)

type RouteConstructor[T any, R any] func(T) R

func RouteProvider[T any, R contract.Route](routeConstructor RouteConstructor[T, R]) any {
	return fx.Annotate(
		routeConstructor,
		fx.As(new(R)),
		fx.ResultTags(`group:"routes"`),
	)
}

func ProtectedRouteProvider[T any, R contract.ProtectedRoute](routeConstructor RouteConstructor[T, R]) any {
	return fx.Annotate(
		routeConstructor,
		fx.As(new(R)),
		fx.ResultTags(`group:"protected_routes"`),
	)
}
