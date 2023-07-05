package main

import (
	"context"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/api"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/config"
)

func main() {
	if err := config.Init(); err != nil {
		panic(err)
	}

	if err := api.NewPubSub(); err != nil {
		panic(err)
	}

	api, err := api.NewAPI()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	api.Run(ctx)
}
