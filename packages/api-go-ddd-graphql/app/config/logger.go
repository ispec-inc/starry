package config

const (
	LoggerTypeStdlog = LoggerType("stdlog")
	LoggerTypeSentry = LoggerType("sentry")
)

type LoggerType string

var Logger logger

type logger struct {
	Type LoggerType
}
