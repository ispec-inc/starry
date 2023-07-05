package test

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	testDriver = "txdb"
)

// InitMySQL テスト用のMySQLの初期化
func InitMySQL() {
	dns := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		config.MySQL.User,
		config.MySQL.Password,
		config.MySQL.Host,
		config.MySQL.Port,
		config.MySQL.Database,
	)
	txdb.Register(testDriver, "mysql", dns)
}

// PrepareDB テスト用のDBの準備。シードデータを渡すとテスト用のDBにデータを投入する
func PrepareDB(t *testing.T, name string, seeds []interface{}) (*app.DB, func()) {
	t.Helper()

	dialector := mysql.New(mysql.Config{
		DriverName: testDriver,
		DSN:        name,
	})
	db, err := gorm.Open(dialector)
	if err != nil {
		t.Fatal(err)
	}

	for _, s := range seeds {
		if err := db.Create(s).Error; err != nil {
			t.Fatal(err)
		}
	}
	return app.NewDB(db), func() { sqldb, _ := db.DB(); sqldb.Close() } // nolint
}
