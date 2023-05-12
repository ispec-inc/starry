package config

var Sentry sentry

type sentry struct {
	DSN   string `env:"SENTRY_DSN"`
	Env   string `env:"SENTRY_ENVIRONMENT"`
	Debug bool   `env:"SENTRY_DEBUG"`
}
