package infra

import (
	"context"
	"errors"
	"fmt"
	"rest-fiber/config"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisService interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value any, ttl time.Duration) error
	Del(ctx context.Context, keys ...string) error
	Ping(ctx context.Context) error
	Close() error
}

type redisServiceImpl struct {
	client *redis.Client
}

func NewRedisService(env config.Env) (RedisService, error) {
	address := fmt.Sprintf("%s:%s", env.RedisHost, env.RedisPort)
	password := env.RedisPassword
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB: 0,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, errors.New(err.Error())
	}
	return &redisServiceImpl{
		client: client,
	}, nil
}

func (r *redisServiceImpl) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *redisServiceImpl) Set(
	ctx context.Context,
	key string,
	value any,
	ttl time.Duration,
) error {
	return r.client.Set(ctx, key, value, ttl).Err()
}

func (r *redisServiceImpl) Del(ctx context.Context, keys ...string) error {
	return r.client.Del(ctx, keys...).Err()
}

func (r *redisServiceImpl) Ping(ctx context.Context) error {
	return r.client.Ping(ctx).Err()
}

func (r *redisServiceImpl) Close() error {
	return r.client.Close()
}
