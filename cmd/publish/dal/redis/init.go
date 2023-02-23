package redis

import (
	"doushengV4/pkg/consts"
	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     consts.RedisAddr,
		Password: consts.RdisPassword,
		DB:       consts.RedisDatabase,
	})
}
