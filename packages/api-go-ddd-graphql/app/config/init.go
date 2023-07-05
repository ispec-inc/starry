package config

import (
	"github.com/caarlos0/env/v6"
	"go.uber.org/multierr"
)

// Init 環境変数を読み込む
func Init() error {
	var err error
	if enverr := env.Parse(&MySQL); enverr != nil {
		err = multierr.Append(err, enverr)
	}

	if enverr := env.Parse(&Redis); enverr != nil {
		err = multierr.Append(err, enverr)
	}

	if enverr := env.Parse(&Router); enverr != nil {
		err = multierr.Append(err, enverr)
	}

	return err
}
