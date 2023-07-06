package app

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	mysqlDB   *gorm.DB
	mysqlOnce sync.Once
)

// DB データベースのコネクションを管理する構造体
type DB struct {
	// gorm gormのクライアント
	gorm *gorm.DB
}

// NewDB DBのコンストラクタ
func NewDB(gorm *gorm.DB) *DB {
	return &DB{gorm: gorm}
}

// Begin トランザクションを開始する
func (d *DB) Begin() *DB {
	return &DB{gorm: d.gorm.Begin()}
}

// Commit トランザクションをコミットする
func (d *DB) Commit() error {
	return d.gorm.Commit().Error
}

// Rollback トランザクションをロールバックする
func (d *DB) Rollback() {
	d.gorm.Rollback()
}

// Get gormのクライアントを返す
func (d *DB) Get() *gorm.DB {
	return d.gorm
}

// Transaction gormのトランザクションを実行する
func (d *DB) Transaction(f func(*gorm.DB) error) error {
	return d.gorm.Transaction(f)
}

// MySQL はMySQLのクライアントを返す
// sync.Onceを使うことで、複数回この関数が呼ばれても、クライアントは一度だけしか生成されないようにしている
func MySQL() (*DB, error) {
	var (
		err error
		db  *sql.DB
	)

	mysqlOnce.Do(func() {
		var loglev logger.LogLevel
		switch config.MySQL.LogLevel {
		case "info":
			loglev = logger.Info
		case "warn":
			loglev = logger.Warn
		default:
			loglev = logger.Error
		}
		dns := fmt.Sprintf(
			"%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=true",
			config.MySQL.User,
			config.MySQL.Password,
			config.MySQL.Host,
			config.MySQL.Port,
			config.MySQL.Database,
		)

		dialector := mysql.Open(dns)
		mysqlDB, err = gorm.Open(dialector, &gorm.Config{
			Logger: logger.Default.LogMode(loglev),
		})

		db, err = mysqlDB.DB()

		db.SetMaxIdleConns(config.MySQL.MaxIdleConn)
		db.SetMaxOpenConns(config.MySQL.MaxOpenConn)
		db.SetConnMaxLifetime(config.MySQL.MaxLifetime)
	})

	return &DB{gorm: mysqlDB}, err
}
