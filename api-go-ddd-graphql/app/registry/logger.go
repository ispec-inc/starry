package registry

import (
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/config"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/pkg/applog/logger"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/pkg/sentry"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/pkg/stdlog"
)

type Logger struct {
	lgr logger.Logger
}

func NewLogger() (Logger, func() error, error) {
	var (
		lgr     logger.Logger
		clenaup func() error
		err     error
	)

	switch config.Logger.Type {
	case config.LoggerTypeSentry:
		slgr, scleanup, serr := sentry.New(
			sentry.Config{
				Environment: config.Sentry.Env,
				DSN:         config.Sentry.DSN,
				Debug:       config.Sentry.Debug,
			},
		)
		lgr = slgr
		clenaup = func() error { scleanup(); return nil }
		err = serr
	default:
		lgr = stdlog.New()
		clenaup = func() error { return nil }
	}

	return Logger{lgr}, clenaup, err
}

func (l Logger) New() logger.Logger {
	return l.lgr
}
