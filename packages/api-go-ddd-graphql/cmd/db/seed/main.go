package main

import (
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/service/clinic/seed"
	"gorm.io/gorm"
)

func main() {
	db, err := app.MySQL()

	if err != nil {
		panic(err)
	}

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
