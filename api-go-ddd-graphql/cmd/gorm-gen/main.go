package main

import (
	"log"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/config"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/pkg/rdb"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./api/app/infra/entity",
		ModelPkgPath:  "entity",
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	db, err := rdb.New(rdb.Config{
		DBMS: rdb.DBMSMySQL,
		Conn: rdb.Connection{
			User:     config.MySQL.User,
			Password: config.MySQL.Password,
			Host:     config.MySQL.Host,
			Port:     config.MySQL.Port,
			Database: config.MySQL.Database,
		},
		MaxIdleConn: config.MySQL.MaxIdleConn,
		MaxOpenConn: config.MySQL.MaxOpenConn,
		MaxLifetime: config.MySQL.MaxLifetime,
	})
	if err != nil {
		log.Fatal(err)
	}

	g.UseDB(db)
	g.GenerateAllTable()

	genClinic(g)
	genClinicalCase(g)

	g.Execute()
}
