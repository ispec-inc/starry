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

	api, sclnup, err := web.NewAPI()
	if err != nil {
		panic(err)
	}
	defer sclnup()

	ctx := context.Background()
	api.Run(ctx)
}
