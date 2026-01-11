package category

import "github.com/gofiber/fiber/v2"

type CategoryHandler interface {
	GetAllCategories(c *fiber.Ctx) error
	GetCategoryByID(c *fiber.Ctx) error
	CreateCategory(c *fiber.Ctx) error
	UpdateCategoryByID(c *fiber.Ctx) error
	DeleteCategoryByID(c *fiber.Ctx) error
}
