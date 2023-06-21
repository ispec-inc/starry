package main

import (
	"fmt"
	"log"

	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir/config"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./api/fhir/query",
		ModelPkgPath:  "./api/fhir/model",
		WithUnitTest:  true,
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	dns := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		config.MySQL.User,
		config.MySQL.Password,
		config.MySQL.Host,
		config.MySQL.Port,
		config.MySQL.Database,
	)

	mysqlDB, err := gorm.Open(mysql.Open(dns))
	if err != nil {
		log.Fatal(err)
	}

	g.UseDB(mysqlDB)
	g.GenerateAllTable()

	genOrganization(g)
	g.Execute()
}
