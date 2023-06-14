package app_test

import (
	"testing"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/app"
)

func Test_Redis(t *testing.T) {
	r, err := app.Redis()
	if err != nil {
		t.Fatal(err)
	}

	r2, nil := app.Redis()
	if err != nil {
		t.Fatal(err)
	}

	if r != r2 {
		t.Fatal("Redis() should return same instance")
	}
}
