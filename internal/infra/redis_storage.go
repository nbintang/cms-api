package infra

import (
	"rest-fiber/config"

	redisStorage "github.com/gofiber/storage/redis"
)

func GetRedisStorage(env config.Env) *redisStorage.Storage {
	return redisStorage.New(redisStorage.ConfigDefault)
}