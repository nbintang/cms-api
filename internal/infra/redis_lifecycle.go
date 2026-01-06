package infra

import (
	"context" 

	"go.uber.org/fx"
)

func RegisterRedisLifecycle(
	lc fx.Lifecycle,
	redisSvc RedisService,
) {
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return redisSvc.Close()
		},
	})
}
