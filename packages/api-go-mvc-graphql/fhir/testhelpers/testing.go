package testhelpers

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/go-playground/validator/v10"
	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir"
	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir/config"
	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir/query"
	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir/validation"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	testDriver = "txdb"
)

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

func PrepareDB(t *testing.T, name string, seeds []interface{}) (*gorm.DB, func()) {
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
	return db, func() { sqldb, _ := db.DB(); sqldb.Close() }
}

func NewRegistry(
	t *testing.T,
	name string,
	seeds []interface{},
) (*fhir.Registry, func()) {
	t.Helper()

	db, clnup := PrepareDB(t, name, seeds)

	cache := fhir.Redis()

	val := validator.New()
	val.RegisterValidation("kana", validation.ValidateKana)

	return &fhir.Registry{
		Q:         query.Use(db),
		Cache:     cache,
		Validator: val,
	}, clnup
}
