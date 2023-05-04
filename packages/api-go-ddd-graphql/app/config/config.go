package config

import (
	"os"

	"github.com/caarlos0/env/v6"
	"go.uber.org/multierr"
)

func Init() error {
	var err error
	if enverr := env.Parse(&MySQL); enverr != nil {
		err = multierr.Append(err, enverr)
	}

	lgtype := os.Getenv("LOGGER_TYPE")
	switch lgtype {
	case string(LoggerTypeSentry):
		Logger.Type = LoggerTypeSentry
		if enverr := env.Parse(&Sentry); enverr != nil {
			err = multierr.Append(err, enverr)
		}
	default:
		Logger.Type = LoggerTypeStdlog
	}

	if enverr := env.Parse(&RedisMsgbs); enverr != nil {
		err = multierr.Append(err, enverr)
	}

	if enverr := env.Parse(&Router); enverr != nil {
		err = multierr.Append(err, enverr)
	}

	return err
}
