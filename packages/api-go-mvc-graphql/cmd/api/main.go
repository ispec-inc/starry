package main

import (
	"context"

	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir/api"
	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir/config"
)

func main() {
	if err := config.Init(); err != nil {
		panic(err)
	}

	api, err := api.New()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	api.Run(ctx)
}
