package auth

import (
	"rest-fiber/config"
	"rest-fiber/internal/infra"

	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	authService AuthService
	validate    infra.Validator
	env         config.Env
}

func NewAuthHandler(authService AuthService, validate infra.Validator, env config.Env) AuthHandler {
	return &authHandler{authService, validate, env}
}

func (h *authHandler) Register(c *fiber.Ctx) error {
	var dto RegisterRequestDTO
	ctx := c.UserContext()
	if err := c.BodyParser(&dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := h.validate.Struct(dto); err != nil {
		return err
	}

	token, err := h.authService.Register(ctx, &dto)
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user created successfuly",
		"token":   token,
	})
}
