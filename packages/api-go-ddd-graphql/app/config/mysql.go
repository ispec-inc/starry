package config

import (
	"time"

	"github.com/caarlos0/env/v6"
)

func init() {
	if err := env.Parse(&MySQL); err != nil {
		panic(err)
	}
}

var MySQL mysql

type mysql struct {
	User        string        `env:"MYSQL_USER,notEmpty"`
	Password    string        `env:"MYSQL_PASSWORD,notEmpty"`
	Database    string        `env:"MYSQL_DATABASE,notEmpty"`
	Host        string        `env:"MYSQL_HOST,notEmpty"`
	Port        string        `env:"MYSQL_PORT" envDefault:"3306"`
	ShowAllLog  bool          `env:"MYSQL_SHOW_ALL_LOG" envDefault:"false"`
	MaxIdleConn int           `env:"MYSQL_MAX_IDLE_CONN" envDefault:"25"`
	MaxOpenConn int           `env:"MYSQL_MAX_OPEN_CONN" envDefault:"25"`
	MaxLifetime time.Duration `env:"MYSQL_MAX_CONN_LIFETIME" envDefault:"25s"`
}
