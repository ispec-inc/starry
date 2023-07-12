package config

import (
	"time"
)

// Router APIのルーターの設定
var Router router

type router struct {
	Timeout      time.Duration `env:"ROUTER_TIMEOUT"`
	AllowOrigins []string      `env:"ROUTER_ALLOW_ORIGINS" envSeparator:","`
}
