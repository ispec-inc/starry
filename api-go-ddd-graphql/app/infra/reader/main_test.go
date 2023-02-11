package reader_test

import (
	"os"
	"testing"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/config"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/pkg/rdb"
)

func TestMain(m *testing.M) {
	err := rdb.InitTest(rdb.TestConfig{
		DBMS: rdb.DBMSMySQL,
		Conn: rdb.Connection{
			User:     config.MySQL.User,
			Password: config.MySQL.Password,
			Host:     config.MySQL.Host,
			Port:     config.MySQL.Port,
			Database: config.MySQL.Database,
		},
	})
	if err != nil {
		os.Exit(1)
	}
	m.Run()
}
