package app_test

import (
	"testing"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/app"
)

func Test_Redis(t *testing.T) {
	r := app.Redis()

	r2 := app.Redis()

	if r != r2 {
		t.Fatal("Redis() should return same instance")
	}
}
