package fhir

import (
	"fmt"
	"sync"

	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir/config"
	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
	redisOnce   sync.Once
)

// Redis はRedisのクライアントを返す
func Redis() *redis.Client {
	redisOnce.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port),
			Password: "",
			DB:       0,
		})
	})

	return redisClient
}
