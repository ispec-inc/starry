package registry

import (
	"gorm.io/gorm"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/config"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/domain/query"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/infra/reader"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/pkg/rdb"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository() (Repository, func() error, error) {
	var loglev rdb.LogLevel
	if config.MySQL.ShowAllLog {
		loglev = rdb.LogLevelInfo
	} else {
		loglev = rdb.LogLevelError
	}

	db, err := rdb.New(rdb.Config{
		DBMS: rdb.DBMSMySQL,
		Conn: rdb.Connection{
			User:     config.MySQL.User,
			Password: config.MySQL.Password,
			Host:     config.MySQL.Host,
			Port:     config.MySQL.Port,
			Database: config.MySQL.Database,
		},
		LogLevel:    loglev,
		MaxIdleConn: config.MySQL.MaxIdleConn,
		MaxOpenConn: config.MySQL.MaxOpenConn,
		MaxLifetime: config.MySQL.MaxLifetime,
	})
	if err != nil {
		return Repository{}, func() error { return nil }, err
	}

	repo := Repository{
		db: db,
	}
	cleanup := func() error { return nil }
	return repo, cleanup, nil
}

func (r Repository) NewClinicQuery() query.Clinic {
	return reader.NewClinic(r.db)
}
