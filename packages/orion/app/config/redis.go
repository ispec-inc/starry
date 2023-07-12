package config

// Redis Redisの設定
var Redis redis

type redis struct {
	Host string `env:"REDIS_MSGBS_HOST"`
	Port string `env:"REDIS_MSGBS_PORT"`
}
