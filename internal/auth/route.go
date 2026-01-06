package auth

import (
	"rest-fiber/config"
	"rest-fiber/internal/contract"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"go.uber.org/fx"
)

type AuthRouteParams struct {
	fx.In
	H   AuthHandler
	Env config.Env
}

type authRouteImpl struct {
	h   AuthHandler
	env config.Env
}

func NewAuthRoute(p AuthRouteParams) contract.Route {
	return &authRouteImpl{h: p.H, env: p.Env}
}
func (r *authRouteImpl) RegisterRoute(api fiber.Router) {
	auth := api.Group("/auth")

	authLimiter := limiter.New(limiter.Config{
		Max:        5,
		Expiration: 1 * time.Minute,

		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},

		// pastiin request tetap dihitung walau status code error
		SkipFailedRequests:     false,
		SkipSuccessfulRequests: false,

		LimitReached: func(c *fiber.Ctx) error {
			return fiber.NewError(fiber.StatusBadRequest, "Too many requests")
		},
	})

	auth.Post("/register", authLimiter, r.h.Register)
	auth.Post("/verify", authLimiter, r.h.VerifyEmail)
	auth.Post("/login", authLimiter, r.h.Login)
	auth.Post("/refresh-token", r.h.RefreshToken)
}
