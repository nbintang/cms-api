package auth

import (
	"rest-fiber/config"
	"rest-fiber/internal/infra"
	"rest-fiber/internal/setup"
	"time"

	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	authService AuthService
	validate    infra.Validator
	env         config.Env
	logger      *infra.AppLogger
}

func NewAuthHandler(authService AuthService, validate infra.Validator, env config.Env, logger *infra.AppLogger) AuthHandler {
	return &authHandler{authService, validate, env, logger}
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

	if err := h.authService.Register(ctx, &dto); err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "token sent successfully, please check your email",
	})
}

func (h *authHandler) VerifyEmail(c *fiber.Ctx) error {
	token := c.Query("token")
	ctx := c.UserContext()
	tokens, err := h.authService.VerifyEmailToken(ctx, token)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    tokens.RefreshToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Lax",
		Path:     "/",
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":       "email verified successfully",
		"access_token":  tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
	})
}

func (h *authHandler) Login(c *fiber.Ctx) error {
	var dto LoginRequestDTO
	ctx := c.UserContext()

	if err := c.BodyParser(&dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	tokens, err := h.authService.Login(ctx, &dto)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    tokens.RefreshToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   false,

		Path: "/",
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":       "login successful",
		"access_token":  tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
	})
}

func (h *authHandler) RefreshToken(c *fiber.Ctx) error {
	ctx := c.UserContext()
	refreshToken := c.Cookies("refresh_token")
	tokens, err := h.authService.RefreshToken(ctx, refreshToken)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "no token provide, please login!")
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":       "login successful",
		"access_token":  tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
	})
}

func (h *authHandler) Logout(c *fiber.Ctx) error {
	ctx := c.UserContext()
	refreshToken := c.Cookies("refresh_token")
	ok, err := h.authService.Logout(ctx, refreshToken)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "no token provide, please login!")
	}
	if ok {
		c.Cookie(&fiber.Cookie{
			Name:     "refresh_token",
			Value:    "",
			Expires:  time.Unix(0, 0),
			MaxAge:   -1,
			HTTPOnly: true,
			Secure:   false,

			Path: "/",
		})
		return c.Status(fiber.StatusOK).JSON(setup.NewHttpResponse(fiber.StatusOK,  nil, "logout successful",))
	}
	return c.Status(fiber.StatusBadRequest).JSON(setup.NewHttpResponse(fiber.StatusBadRequest, nil, "logout failed",))
}
