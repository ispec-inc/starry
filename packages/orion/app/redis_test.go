package app_test

import (
	"testing"

	"github.com/ispec-inc/starry/orion/app"
)

func Test_Redis(t *testing.T) {
	t.Parallel()

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
