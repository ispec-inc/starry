package api_test

import (
	"context"
	"testing"

	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir/api"
)

func TestAPI(t *testing.T) {
	api, err := api.New()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	api.Run(ctx)

	ctx.Done()
}
