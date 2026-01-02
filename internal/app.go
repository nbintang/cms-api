package internal

import (
	"context"
	"log"
	"rest-fiber/config"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func NewApp() *fiber.App {
	return fiber.New()

}

func RunApp(lc fx.Lifecycle, app *fiber.App, env config.Env) {
	addr := env.AppAddr
	if addr == "" {
		addr = ":3000"
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

var Module = fx.Module(
	"app",
	fx.Provide(NewApp),
	fx.Invoke(RunApp),
)
