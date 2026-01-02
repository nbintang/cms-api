package user

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type UserRepository interface {
	FindAll(ctx context.Context) ([]User, error)
}

type UserService interface {
	FindAllUsers(ctx context.Context) ([]UserResponse, error)
}

type UserHandler interface {
	GetAllUsers(c *fiber.Ctx) error
}
