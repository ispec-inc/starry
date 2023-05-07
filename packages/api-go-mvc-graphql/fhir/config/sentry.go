package config

var Sentry sentry

type sentry struct {
	DSN   string `env:"ARTICLE_SENTRY_DSN"`
	Env   string `env:"ARTICLE_SENTRY_ENVIRONMENT"`
	Debug bool   `env:"ARTICLE_SENTRY_DEBUG"`
}
