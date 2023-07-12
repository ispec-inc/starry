package main

import (
	"context"

	"github.com/ispec-inc/starry/orion/api"
	"github.com/ispec-inc/starry/orion/app/config"
)

func main() {
	if err := config.Init(); err != nil {
		panic(err)
	}

	api, err := api.NewAPI()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	api.Run(ctx)
}
