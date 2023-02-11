package main

import (
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/config"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/seed"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/pkg/rdb"
	"gorm.io/gorm"
)

func main() {
	db, err := getDB()
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	seeds := seed.Dev()
	err = db.Transaction(func(tx *gorm.DB) error {
		for _, s := range seeds {
			if err := tx.Create(s).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func getDB() (*gorm.DB, error) {
	return rdb.New(rdb.Config{
		DBMS: rdb.DBMSMySQL,
		Conn: rdb.Connection{
			User:     config.MySQL.User,
			Password: config.MySQL.Password,
			Host:     config.MySQL.Host,
			Port:     config.MySQL.Port,
			Database: config.MySQL.Database,
		},
		LogLevel:    rdb.LogLevelInfo,
		MaxIdleConn: config.MySQL.MaxIdleConn,
		MaxOpenConn: config.MySQL.MaxOpenConn,
		MaxLifetime: config.MySQL.MaxLifetime,
	})
}
