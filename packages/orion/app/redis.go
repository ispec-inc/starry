package app

import (
	"fmt"
	"sync"

	"github.com/ispec-inc/starry/orion/app/config"
	"github.com/redis/rueidis"
)

var (
	redisClient rueidis.Client
	redisOnce   sync.Once
)

// Redis はRedisのクライアントを返す
func Redis() (rueidis.Client, error) {
	var err error
	redisOnce.Do(func() {
		address := fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port)
		redisClient, err = rueidis.NewClient(rueidis.ClientOption{
			InitAddress: []string{address},
		})
	})

	return redisClient, err
}
