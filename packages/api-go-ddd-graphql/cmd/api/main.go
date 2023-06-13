package main

import (
	"context"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/config"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/web"
)

func main() {
	if err := config.Init(); err != nil {
		panic(err)
	}

	api, err := web.NewAPI()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	api.Run(ctx)
}
