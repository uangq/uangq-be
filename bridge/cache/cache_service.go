package cache

import (
	"github.com/go-redis/redis/v7"
	"time"
	"uangq-be/common/logger"
	"uangq-be/configs"
)

var cacheService *redis.Client

func InitCache() {
	cacheService = redis.NewClient(&redis.Options{
		Addr: configs.RedisConfig.Address,
	})
	_, err := cacheService.Ping().Result()
	if err != nil {
		logger.Log("Fail to connect to redis.")
	}
}

func Get(key string) (string, error) {
	return cacheService.Get(key).Result()
}

func Set(key string, value interface{}, expiry time.Duration) {
	cacheService.Set(key, value, expiry)
}

func Del(key string) {
	cacheService.Del(key)
}
