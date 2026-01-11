package category

import (
	"rest-fiber/pkg/httpx"

	"github.com/gofiber/fiber/v2"
)

type CategoryRouteParams struct {
	httpx.RouteParams
	CategoryHandler CategoryHandler
}

type categoryRouteImpl struct {
	categoryHandler CategoryHandler
}

func NewCategoryRoutes(params CategoryRouteParams) httpx.ProtectedRoute {
	return &categoryRouteImpl{categoryHandler: params.CategoryHandler}
}

func (r *categoryRouteImpl) RegisterProtectedRoute(route fiber.Router) {
	categories := route.Group("/categories")
	categories.Get("/", r.categoryHandler.GetAllCategories)
	categories.Get("/:id", r.categoryHandler.GetCategoryByID)
	categories.Post("/", r.categoryHandler.CreateCategory)
	categories.Patch("/:id", r.categoryHandler.UpdateCategoryByID)
	categories.Delete("/:id", r.categoryHandler.DeleteCategoryByID)
}
