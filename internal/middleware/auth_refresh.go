package middleware
 
import (
	"rest-fiber/config"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func AuthJWTRefresh(c *fiber.Ctx) error {
	env, err := config.GetEnvs()
	if err != nil {
		return err
	}
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(env.JWTSecret)},
		ContextKey: "jwt",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message":     "unauthorized",
				"status_code": fiber.StatusUnauthorized,
				"error":       err.Error(),
			})
		},
	})(c)
}
