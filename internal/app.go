package internal

import (
	"context"
	"log"
	"rest-fiber/config"
	"rest-fiber/pkg"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func NewApp() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: pkg.DefaultErrorHandler,
	})
	api := app.Group("api")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"hello_to": "world"})
	})
	return app
}

var RunApp = func(lc fx.Lifecycle, app *fiber.App, env config.Env) {
	addr := env.AppAddr
	if addr == "" {
		addr = ":8080"
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Printf("Fiber listening on %s", addr)
				if err := app.Listen(addr); err != nil {
					log.Printf("Fiber stopped: %s", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})
}
